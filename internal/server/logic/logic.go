package logic

var pastFourScores = []float64{5.0, 4.0, 2.0, 1.0}

func GetSize() float64 {
	if pastFourScores[0]+pastFourScores[1] < pastFourScores[2]+pastFourScores[3] {
		return 600.0
	}
	return 100.0
}

func Average() float64 {
	var sum float64
	for _, item := range pastFourScores {
		sum += item
	}
	average := sum / float64(len(pastFourScores))
	return average
}
