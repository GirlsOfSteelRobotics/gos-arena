// Copyright 2017 Team 254. All Rights Reserved.
// Author: pat@patfairbank.com (Patrick Fairbank)
//
// Helper methods for use in tests in this package and others.

package game

func TestScore1() *Score {
	fouls := []Foul{
		{true, 25, 13},
		{false, 1868, 14},
		{true, 25, 15},
	}
	return &Score{
		LeaveStatuses:   [3]bool{true, true, false},
		EndgameStatuses: [3]EndgameStatus{EndgameParked, EndgameNone, EndgameDeep},
		Fouls:           fouls,
		PlayoffDq:       false,

		AlgaeCoral: AlgaeCoral{
			AlgaeAutoProcessorCount:   2,
			AlgaeAutoNetCount:         0,
			AlgaeTeleopProcessorCount: 0,
			AlgaeTeleopNetCount:       0,
			CoralAutoCount:            [4]int{1, 1, 0, 1},
			CoralTeleopCount:          [4]int{4, 4, 0, 4},
		},
	}
}

func TestScore2() *Score {
	return &Score{
		LeaveStatuses:   [3]bool{true, true, true},
		EndgameStatuses: [3]EndgameStatus{EndgameParked, EndgameShallow, EndgameShallow},
		Fouls:           []Foul{},
		PlayoffDq:       false,

		AlgaeCoral: AlgaeCoral{
			AlgaeAutoProcessorCount:   0,
			AlgaeAutoNetCount:         1,
			AlgaeTeleopProcessorCount: 2,
			AlgaeTeleopNetCount:       2,
			CoralAutoCount:            [4]int{5, 0, 0, 1},
			CoralTeleopCount:          [4]int{0, 5, 0, 4},
		},
	}
}

func TestRanking1() *Ranking {
	return &Ranking{254, 1, 0, RankingFields{20, 1, 625, 90, 554, 0.254, 3, 2, 1, 0, 10}}
}

func TestRanking2() *Ranking {
	return &Ranking{1114, 2, 1, RankingFields{18, 1, 700, 625, 90, 0.1114, 1, 3, 2, 0, 10}}
}
