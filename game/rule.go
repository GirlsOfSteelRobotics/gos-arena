// Copyright 2020 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)
//
// Model of a game-specific rule.

package game

type Rule struct {
	Id             int
	RuleNumber     string
	IsTechnical    bool
	IsRankingPoint bool
	Description    string
}

// All rules from the 2022 game that carry point penalties.
var rules = []*Rule{
	{1, "G210", true, false, "Don't expect to gain by doing others harm. A strategy not consistent with standard gameplay and clearly aimed at forcing the opponent ALLIANCE to violate a rule is not in the spirit of FIRST Robotics Competition and not allowed."},
	{2, "G301", true, false, "Be prompt. A DRIVE TEAM member may not cause significant delays to the start of their MATCH."},
	{3, "G401", false, false, "Behind the lines. In AUTO, a DRIVE TEAM member staged behind a HUMAN STARTING LINE may not contact anything in front of that HUMAN STARTING LINE, unless for personal or equipment safety, topress the E-Stop or A-Stop, or granted permission by a Head REFEREE or FTA"},
	{4, "G402", false, false, "Let the ROBOT do its thing. In AUTO, a DRIVE TEAM member may not directly or indirectly interact with a ROBOT or an OPERATOR CONSOLE unless for personal safety, OPERATOR CONSOLE safety, or pressing an E-Stop or A-Stop. A HUMAN PLAYER feeding CORAL to a ROBOT is an exception to this rule."},
	{5, "G403", true, false, "Limited AUTO opponent interaction. In AUTO, a ROBOT whose BUMPERS are completely across the BARGE ZONE (i.e. to the opposite side of the BARGE ZONE from its ROBOT STARTING LINE) may not contact an opponent ROBOT (either directly or transitively through a SCORING ELEMENT CONTROLLED by either ROBOT and regardless of who initiates contact). "},
	{6, "G404", false, false, "No throwing in AUTO. In AUTO, a HUMAN PLAYER may not enter ALGAE onto the field."},
	{7, "G405", false, false, "No opponents CAGES in AUTO. In AUTO, a ROBOT may not contact the opposing ALLIANCE’s CAGES."},
	{8, "G406", true, false, "ROBOTS: use SCORING ELEMENTS as directed. A ROBOT may not deliberately use a SCORING ELEMENT in an attempt to ease or amplify a challenge associated with a FIELD element"},
	{9, "G407", false, false, "A ROBOT may not intentionally eject a SCORING ELEMENT from the FIELD (either directly or by bouncing off a FIELD element or other ROBOT) other than through a PROCESSOR"},
	{10, "G407", true, false, "Repeated: A ROBOT may not intentionally eject a SCORING ELEMENT from the FIELD (either directly or by bouncing off a FIELD element or other ROBOT) other than through a PROCESSOR"},
	{11, "G408", true, false, "Repeated: Neither a ROBOT nor a HUMAN PLAYER may damage a SCORING ELEMENT."},
	{12, "G409", false, false, "1 of each at a time. A ROBOT may not simultaneously CONTROL more than 1 CORAL and 1 ALGAE either directly or transitively through other objects"},
	{13, "G410", true, true, "No de-scoring. A ROBOT may not de-score a CORAL scored on the opponent’s REEF"},
	{14, "G411", true, false, "Don’t put ALGAE on their REEF. A ROBOT may not deliberately put ALGAE on their opponent’s REEF"},
	{15, "G412", true, false, "Only throw CORAL if in your REEF ZONE. A ROBOT may not launch CORAL unless their BUMPERS are partially in their REEF ZONE."},
	{16, "G414", false, false, "Keep your BUMPERS low. BUMPERS must be in the BUMPER ZONE (see R405)."},
	{17, "G415", false, false, "Expansion limits. A ROBOT may not extend more than 1 ft. 6 in. (~45 cm) beyond the vertical projection of its ROBOT PERIMETER. "},
	{18, "G415", true, false, "For strategic benefit: Expansion limits. A ROBOT may not extend more than 1 ft. 6 in. (~45 cm) beyond the vertical projection of its ROBOT PERIMETER. "},
	{19, "G417", true, false, "Watch your FIELD interaction. A ROBOT is prohibited from the following interactions with FIELD elements with the exception of CAGES"},
	{20, "G418", true, true, "An Opponent’s CAGES are off-limits in TELEOP. In TELEOP, A ROBOT may not contact an opponent’s CAGE."},
	{21, "G419", true, false, "ANCHORS are off-limits. A ROBOT may not contact the ANCHORS. Exceptions are granted for actions that are MOMENTARY and inconsequential."},
	{22, "G420", true, false, "NET and contents are off-limits. A ROBOT may not contact either NET or any ALGAE scored in a NET. "},
	{23, "G421", false, false, "1 defender at a time. No more than 1 ROBOT may be on the opponent’s side of the FIELD (i.e.containing the opponent REEF) with its BUMPERS fully outside and beyond the BARGE ZONES. "},
	{24, "G421", true, false, "3 seconds not corrected: 1 defender at a time. No more than 1 ROBOT may be on the opponent’s side of the FIELD (i.e. containing the opponent REEF) with its BUMPERS fully outside and beyond the BARGE ZONES. "},
	{25, "G422", false, false, "Stay out of other ROBOTS. A ROBOT may not use a COMPONENT outside its ROBOT PERIMETER (except its BUMPERS) to initiate contact with an opponent ROBOT inside the vertical projection of the opponent’s ROBOT PERIMETER."},
	{26, "G423", true, false, "This isn’t combat robotics. A ROBOT may not damage or functionally impair an opponent ROBOT in either of the following ways:"},
	{27, "G424", true, false, "Don’t tip or entangle. A ROBOT may not deliberately, attach to, tip, or entangle with an opponentROBOT."},
	{28, "G425", false, false, "There’s a 3-count on PINS. A ROBOT may not PIN an opponent’s ROBOT for more than 3 seconds. A ROBOT is PINNING if it is preventing the movement of an opponent ROBOT by contact, either direct or transitive (such as against a FIELD element)"},
	{29, "G425", true, false, "Additional: There’s a 3-count on PINS. A ROBOT may not PIN an opponent’s ROBOT for more than 3 seconds. A ROBOT is PINNING if it is preventing the movement of an opponent ROBOT by contact, either direct or transitive (such as against a FIELD element)"},
	{30, "G426", true, false, "Don’t collude with your partners to shut down major parts of game play. 2 or more ROBOTS that appear to a REFEREE to be working together may not isolate or close off any major element of MATCHplay."},
	{31, "G427", true, false, "ZONE protection. A ROBOT may not contact, directly or transitively through a SCORING ELEMENT, an opponent ROBOT partially inside the opponent’s BARGE ZONE or REEF ZONE regardless of who initiates contact. "},
	{32, "G428", true, true, "CAGE protection. A ROBOT may not contact, directly or transitively through a SCORING ELEMENT, an opponent ROBOT in contact with an opponent CAGE during the last 20 seconds regardless of who initiates contact. "},
	{33, "G429", false, false, "No wandering. A DRIVE TEAM member must remain in their designated area"},
	{34, "G430", true, false, "*COACHES and other teams: hands off the controls. A ROBOT shall be operated only by the DRIVERS and/or HUMAN PLAYERS of that team. A COACH activating their E-Stop or A-Stop is the exception to this rule."},
	{35, "G431", false, false, "DRIVE TEAMS, watch your reach. A DRIVE TEAM member may not extend into the CHUTE. "},
	{36, "G432", true, false, "Humans: use SCORING ELEMENTS as directed. A DRIVE TEAM member may not deliberately use a SCORING ELEMENT in an attempt to ease or amplify a challenge associated with a FIELD element."},
	{37, "G433", true, false, "SCORING ELEMENT delivery. SCORING ELEMENTS may only be entered onto the FIELD in certain ways"},
	{38, "G434", false, false, "COACHES, SCORING ELEMENTS are off limits. COACHES may not touch SCORING ELEMENTS, unless for safety purposes."},
	{39, "G435", true, false, "The PROCESSOR AREA has a storage limit. HUMAN PLAYERS may not store more than 4 ALGAE in the PROCESSOR AREA (up to 3 in the holders on top of the PROCESSOR and no more than 1 at the end of the PROCESSOR exit ramp). HUMAN PLAYERS making a good-faith effort to immediately enter additional ALGAE is an exception to this rule. "},
}
var ruleMap map[int]*Rule

// Returns the rule having the given ID, or nil if no such rule exists.
func GetRuleById(id int) *Rule {
	return GetAllRules()[id]
}

// Returns a slice of all defined rules that carry point penalties.
func GetAllRules() map[int]*Rule {
	if ruleMap == nil {
		ruleMap = make(map[int]*Rule, len(rules))
		for _, rule := range rules {
			ruleMap[rule.Id] = rule
		}
	}
	return ruleMap
}
