package game

type AlgaeCoral struct {
	AlgaeAutoProcessorCount   int
	AlgaeAutoNetCount         int
	AlgaeTeleopProcessorCount int
	AlgaeTeleopNetCount       int
	CoralAutoCount            [4]int
	CoralTeleopCount          [4]int
}

type Level int

const (
	levelOne Level = iota
	levelTwo
	levelThree
	levelFour
	levelCount
)

var autoCoralPoints = map[Level]int{
	levelOne:   3,
	levelTwo:   4,
	levelThree: 6,
	levelFour:  7,
}

var teleopCoralPoints = map[Level]int{
	levelOne:   2,
	levelTwo:   3,
	levelThree: 4,
	levelFour:  5,
}

func (algaeCoral *AlgaeCoral) AutoGamePiecePoints() int {
	points := 0
	for level := levelOne; level < levelCount; level++ {
		points += algaeCoral.CoralAutoCount[level] * autoCoralPoints[level]
	}
	points += algaeCoral.AlgaeAutoNetCount*4 + algaeCoral.AlgaeAutoProcessorCount*6
	return points
}

func (algaeCoral *AlgaeCoral) TeleopGamePiecePoints() int {
	points := 0
	for level := levelOne; level < levelCount; level++ {
		points += algaeCoral.CoralTeleopCount[level] * teleopCoralPoints[level]
	}
	points += algaeCoral.AlgaeTeleopNetCount*4 + algaeCoral.AlgaeTeleopProcessorCount*6
	return points
}

func (algaeCoral *AlgaeCoral) TotalCoral() []int {
	ret := make([]int, 4)
	for level := 0; level < 4; level++ {
		ret[level] = algaeCoral.CoralAutoCount[level] + algaeCoral.CoralTeleopCount[level]
	}
    return ret
}

func (algaeCoral *AlgaeCoral) IsCoopertitionThresholdAchieved() bool {
	return (algaeCoral.AlgaeAutoProcessorCount + algaeCoral.AlgaeTeleopNetCount) > 2
}
