// Copyright 2014 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)
//
// Web handlers for scoring interface.

package web

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/Team254/cheesy-arena/field"
	"github.com/Team254/cheesy-arena/game"
	"github.com/Team254/cheesy-arena/model"
	"github.com/Team254/cheesy-arena/websocket"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
)

// Renders the scoring interface which enables input of scores in real-time.
func (web *Web) scoringPanelHandler(w http.ResponseWriter, r *http.Request) {
	if !web.userIsAdmin(w, r) {
		return
	}

	vars := mux.Vars(r)
	alliance := vars["alliance"]
	if alliance != "red" && alliance != "blue" {
		handleWebErr(w, fmt.Errorf("Invalid alliance '%s'.", alliance))
		return
	}

	template, err := web.parseFiles("templates/scoring_panel.html", "templates/base.html")
	if err != nil {
		handleWebErr(w, err)
		return
	}
	data := struct {
		*model.EventSettings
		PlcIsEnabled        bool
		Alliance            string
		ValidGridNodeStates map[game.Row]map[int]map[game.NodeState]string
	}{web.arena.EventSettings, web.arena.Plc.IsEnabled(), alliance, game.ValidGridNodeStates()}
	err = template.ExecuteTemplate(w, "base_no_navbar", data)
	if err != nil {
		handleWebErr(w, err)
		return
	}
}

// The websocket endpoint for the scoring interface client to send control commands and receive status updates.
func (web *Web) scoringPanelWebsocketHandler(w http.ResponseWriter, r *http.Request) {
	if !web.userIsAdmin(w, r) {
		return
	}

	vars := mux.Vars(r)
	alliance := vars["alliance"]
	if alliance != "red" && alliance != "blue" {
		handleWebErr(w, fmt.Errorf("Invalid alliance '%s'.", alliance))
		return
	}

	var realtimeScore **field.RealtimeScore
	if alliance == "red" {
		realtimeScore = &web.arena.RedRealtimeScore
	} else {
		realtimeScore = &web.arena.BlueRealtimeScore
	}

	ws, err := websocket.NewWebsocket(w, r)
	if err != nil {
		handleWebErr(w, err)
		return
	}
	defer ws.Close()
	web.arena.ScoringPanelRegistry.RegisterPanel(alliance, ws)
	web.arena.ScoringStatusNotifier.Notify()
	defer web.arena.ScoringStatusNotifier.Notify()
	defer web.arena.ScoringPanelRegistry.UnregisterPanel(alliance, ws)

	// Subscribe the websocket to the notifiers whose messages will be passed on to the client, in a separate goroutine.
	go ws.HandleNotifiers(web.arena.MatchLoadNotifier, web.arena.MatchTimeNotifier, web.arena.RealtimeScoreNotifier,
		web.arena.ReloadDisplaysNotifier)

	// Loop, waiting for commands and responding to them, until the client closes the connection.
	for {
		command, data, err := ws.Read()
		if err != nil {
			if err == io.EOF {
				// Client has closed the connection; nothing to do here.
				return
			}
			log.Println(err)
			return
		}
		score := &(*realtimeScore).CurrentScore
		scoreChanged := false

		if command == "commitMatch" {
			if web.arena.MatchState != field.PostMatch {
				// Don't allow committing the score until the match is over.
				ws.WriteError("Cannot commit score: Match is not over.")
				continue
			}
			web.arena.ScoringPanelRegistry.SetScoreCommitted(alliance, ws)
			web.arena.ScoringStatusNotifier.Notify()
		} else {
			args := struct {
				TeamPosition int
				GridRow      int
				GridNode     int
				NodeState    game.NodeState
			}{}
			err = mapstructure.Decode(data, &args)
			if err != nil {
				ws.WriteError(err.Error())
				continue
			}

			switch command {
			case "leaveStatus":
				if args.TeamPosition >= 1 && args.TeamPosition <= 3 {
					score.LeaveStatuses[args.TeamPosition-1] = !score.LeaveStatuses[args.TeamPosition-1]
					scoreChanged = true
				}
			case "endgameStatus":
				if args.TeamPosition >= 1 && args.TeamPosition <= 3 {
					score.EndgameStatuses[args.TeamPosition-1]++
					if score.EndgameStatuses[args.TeamPosition-1] > 3 {
						score.EndgameStatuses[args.TeamPosition-1] = 0
					}
					scoreChanged = true
				}
			default:
				auto, level, down := parseCommand(command)
				if auto {
					current := score.AlgaeCoral.CoralAutoCount[level]
					if down {
						if current > 0 {
							score.AlgaeCoral.CoralAutoCount[level]--
							scoreChanged = true
						}
					} else {
						if current < game.NumCorals {
							score.AlgaeCoral.CoralAutoCount[level]++
							scoreChanged = true
						}
					}
				} else {
					current := score.AlgaeCoral.CoralTeleopCount[level]
					if down {
						if current > 0 {
							score.AlgaeCoral.CoralTeleopCount[level]--
							scoreChanged = true
						}
					} else {
						if current < game.NumCorals {
							score.AlgaeCoral.CoralTeleopCount[level]++
							scoreChanged = true
						}
					}
				}
				// case "autoCoral1Plus":
				// 	if score.AlgaeCoral.CoralAutoCount[0] < 12 {
				// 		score.AlgaeCoral.CoralAutoCount[0]++
				// 	}
				// 	scoreChanged = true
				// case "gridAutoScoring":
				// 	if args.GridRow >= 0 && args.GridRow <= 2 && args.GridNode >= 0 && args.GridNode <= 8 {
				// 		score.Grid.AutoScoring[args.GridRow][args.GridNode] =
				// 			!score.Grid.AutoScoring[args.GridRow][args.GridNode]
				// 		scoreChanged = true
				// 	}
				// case "gridNode":
				// 	if args.GridRow >= 0 && args.GridRow <= 2 && args.GridNode >= 0 && args.GridNode <= 8 {
				// 		currentState := score.Grid.Nodes[args.GridRow][args.GridNode]
				// 		if currentState == args.NodeState {
				// 			score.Grid.Nodes[args.GridRow][args.GridNode] = game.Empty
				// 			if web.arena.MatchState == field.AutoPeriod || web.arena.MatchState == field.PausePeriod {
				// 				score.Grid.AutoScoring[args.GridRow][args.GridNode] = false
				// 			}
				// 		} else {
				// 			score.Grid.Nodes[args.GridRow][args.GridNode] = args.NodeState
				// 			if web.arena.MatchState == field.AutoPeriod || web.arena.MatchState == field.PausePeriod {
				// 				score.Grid.AutoScoring[args.GridRow][args.GridNode] = true
				// 			}
				// 		}
				// 		scoreChanged = true
				// 	}
			}

			if scoreChanged {
				web.arena.RealtimeScoreNotifier.Notify()
			}
		}
	}
}

// commands are of the format autoCoral1Minus, or teleopCoral3Plus
func parseCommand(command string) (auto bool, level game.Level, down bool) {
	if strings.Contains(command, "auto") {
		auto = true
	}
	if strings.Contains(command, "Minus") {
		down = true
	}
	if strings.Contains(command, "1") {
		level = game.Level(0)
	} else if strings.Contains(command, "2") {
		level = game.Level(1)
	} else if strings.Contains(command, "3") {
		level = game.Level(2)
	} else if strings.Contains(command, "4") {
		level = game.Level(3)
	}
	return
}
