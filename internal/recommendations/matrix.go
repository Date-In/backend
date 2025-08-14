package recommendations

func DatingGoalCompatibilityMatrix() map[uint]map[uint]float64 {
	return map[uint]map[uint]float64{
		1: {1: 1.0, 2: 0.2, 3: 0.1, 4: 0.0, 5: 0.8},
		2: {1: 0.2, 2: 1.0, 3: 0.9, 4: -1.0, 5: 0.2},
		3: {1: 0.1, 2: 0.9, 3: 1.0, 4: -1.0, 5: 0.1},
		4: {1: 0.0, 2: -1.0, 3: -1.0, 4: 1.0, 5: 0.4},
		5: {1: 0.8, 2: 0.2, 3: 0.1, 4: 0.4, 5: 1.0},
	}
}

func WorldviewCompatibilityMatrix() map[uint]map[uint]float64 {
	matrix := make(map[uint]map[uint]float64)
	for i := uint(1); i <= 8; i++ {
		matrix[i] = make(map[uint]float64)
		for j := uint(1); j <= 8; j++ {
			if i == j {
				matrix[i][j] = 1.0
			} else {
				matrix[i][j] = 0.1
			}
		}
	}
	matrix[6][7], matrix[7][6] = 0.8, 0.8
	matrix[1][6], matrix[6][1] = 0.7, 0.7
	matrix[1][7], matrix[7][1] = 0.7, 0.7
	return matrix
}

func EducationCompatibilityMatrix() map[uint]map[uint]float64 {
	matrix := make(map[uint]map[uint]float64)
	numLevels := uint(6)

	for i := uint(1); i <= numLevels; i++ {
		matrix[i] = make(map[uint]float64)
		for j := uint(1); j <= numLevels; j++ {
			diff := float64(i) - float64(j)
			if diff < 0 {
				diff = -diff
			}
			matrix[i][j] = 1.0 - (diff / float64(numLevels-1))
		}
	}
	return matrix
}

func AttitudeCompatibilityMatrix() map[uint]map[uint]float64 {
	return map[uint]map[uint]float64{
		1: {1: 1.0, 2: 0.7, 3: 0.1, 4: 0.0, 5: -1.0},
		2: {1: 0.7, 2: 1.0, 3: 0.5, 4: 0.2, 5: -0.8},
		3: {1: 0.1, 2: 0.5, 3: 1.0, 4: 0.8, 5: 0.0},
		4: {1: 0.0, 2: 0.2, 3: 0.8, 4: 1.0, 5: 0.4},
		5: {1: -1.0, 2: -0.8, 3: 0.0, 4: 0.4, 5: 1.0},
	}
}

func calculateCompatibilityScore(
	userValue, candidateValue *uint,
	baseWeight float64,
	matrix map[uint]map[uint]float64,
) float64 {
	if userValue == nil || candidateValue == nil {
		return 0
	}

	userOptionID := *userValue
	candidateOptionID := *candidateValue

	if compatMap, ok := matrix[userOptionID]; ok {
		if compatibility, ok := compatMap[candidateOptionID]; ok {
			return baseWeight * compatibility
		}
	}
	return 0
}
