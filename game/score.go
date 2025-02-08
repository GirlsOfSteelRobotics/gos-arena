// Copyright 2023 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)
//
// Model representing the instantaneous score of a match.

package game

type Score struct {
	LeaveStatuses   [3]bool
	AlgaeCoral      AlgaeCoral
	EndgameStatuses [3]EndgameStatus
	Fouls           []Foul
	PlayoffDq       bool
}

var CoralPerLevelThreshold = 5
var CoralNumLevelsThresholdWithoutCoop = 4
var CoralNumLevelsThresholdWithCoop = 3
var BargePointsThreshold = 14

const NumCorals = 12

// Represents the state of a robot at the end of the match.
type EndgameStatus int

const (
	EndgameNone EndgameStatus = iota
	EndgameParked
	EndgameShallow
	EndgameDeep
)

// Calculates and returns the summary fields used for ranking and display.
func (score *Score) Summarize(opponentScore *Score) *ScoreSummary {
	summary := new(ScoreSummary)

	// Leave the score at zero if the alliance was disqualified.
	if score.PlayoffDq {
		return summary
	}

	// Calculate autonomous period points.
	for _, leave := range score.LeaveStatuses {
		if leave {
			summary.LeavePoints += 3
		}
	}
	autoAlgaeCoralPoints := score.AlgaeCoral.AutoGamePiecePoints()

	summary.AutoPoints = summary.LeavePoints + autoAlgaeCoralPoints

	// Calculate teleoperated period points.
	teleopAlgaeCoralPoints := score.AlgaeCoral.TeleopGamePiecePoints()

	for i := 0; i < 3; i++ {
		switch score.EndgameStatuses[i] {
		case EndgameParked:
			summary.EndgamePoints += 2
		case EndgameShallow:
			summary.EndgamePoints += 6
		case EndgameDeep:
			summary.EndgamePoints += 12
		}
	}

	summary.AlgaeCoralPoints = autoAlgaeCoralPoints + teleopAlgaeCoralPoints
	summary.AlgaePoints = score.AlgaeCoral.TotalAlgaePoints()
	summary.CoralPoints = score.AlgaeCoral.TotalCoralPoints()
	summary.MatchPoints = summary.LeavePoints + summary.AlgaeCoralPoints + summary.EndgamePoints

	// Calculate penalty points.
	for _, foul := range opponentScore.Fouls {
		summary.FoulPoints += foul.PointValue()
		// Store the number of tech fouls since it is used to break ties in playoffs.
		if foul.IsTechnical {
			summary.NumOpponentTechFouls++
		}

		rule := foul.Rule()
		if rule != nil {
			// Check for the opponent fouls that automatically trigger a ranking point.
			if rule.IsRankingPoint {
				summary.BargeRankingPoint = true // TODO - check which RP
			}
		}
	}

	summary.Score = summary.MatchPoints + summary.FoulPoints

	// Calculate bonus ranking points.
	allRobotsLeft := true
	for i := 0; i < 3; i++ {
		allRobotsLeft = allRobotsLeft && score.LeaveStatuses[i]
	}
	anyCoralScored := false
	for level := levelOne; level < levelCount; level++ {
		if score.AlgaeCoral.CoralAutoCount[level] > 0 {
			anyCoralScored = true
		}
	}
	// An AutoRankingPoint is achieved if all robots leave nd there's
	summary.AutoRankingPoint = allRobotsLeft && anyCoralScored

	summary.CoopertitionBonus = score.AlgaeCoral.IsCoopertitionThresholdAchieved() &&
		opponentScore.AlgaeCoral.IsCoopertitionThresholdAchieved()

	levelsAboveThreshold := 0
	allLevels := score.AlgaeCoral.TotalCoral()
	for level := levelOne; level < levelCount; level++ {
		if allLevels[level] > CoralPerLevelThreshold {
			levelsAboveThreshold++
		}
	}

	enoughLevels := levelsAboveThreshold >= CoralNumLevelsThresholdWithoutCoop
	if summary.CoopertitionBonus {
		enoughLevels = levelsAboveThreshold >= CoralNumLevelsThresholdWithCoop
	}

	summary.CoralRankingPoint = enoughLevels

	summary.BargeRankingPoint = summary.EndgamePoints >= BargePointsThreshold

	if summary.AutoRankingPoint {
		summary.BonusRankingPoints++
	}
	if summary.CoralRankingPoint {
		summary.BonusRankingPoints++
	}
	if summary.BargeRankingPoint {
		summary.BonusRankingPoints++
	}

	return summary
}

// Returns true if and only if all fields of the two scores are equal.
func (score *Score) Equals(other *Score) bool {
	if score.LeaveStatuses != other.LeaveStatuses ||
		score.AlgaeCoral != other.AlgaeCoral ||
		score.EndgameStatuses != other.EndgameStatuses ||
		score.PlayoffDq != other.PlayoffDq ||
		len(score.Fouls) != len(other.Fouls) {
		return false
	}

	for i, foul := range score.Fouls {
		if foul != other.Fouls[i] {
			return false
		}
	}

	return true
}
