package algorithm

const defaultDotMatrixSize = 2

// GenerateDotScreentone outputs byte slice by
//
//	size = dotSize.Size*dotSize.Size*4.
//
// And side size of matrix.
//
// This byte values are threshold values for black. If {pixel_value} < {output_slice[index]} than that pixel is black
func GenerateDotScreentone(dotSize int) ([]byte, int) {
	regularDot := generateDot(dotSize, false)
	invertedDot := generateDot(dotSize, true)

	dotMatricies := make([][]byte, 0, defaultDotMatrixSize*defaultDotMatrixSize)

	for y := 0; y < 2; y++ {
		for x := 0; x < 2; x++ {
			matrix := make([]byte, dotSize*dotSize)

			if (x+y)%2 == 0 {
				copy(matrix, regularDot)
			} else {
				copy(matrix, invertedDot)
			}

			dotMatricies = append(dotMatricies, matrix)
		}
	}

	return mergeMatricies(defaultDotMatrixSize, dotSize, dotMatricies), dotSize * defaultDotMatrixSize
}

// mergeMatricies merge matricies value in one big matrix. dotMatrixSize should be equal of sqrt(len(matricies)).
// Matricies len should be equal to dotSize*dotSize.
//
// Matricies order:
//
//	1 2 3
//	4 5 6
//	7 8 9 etc.
func mergeMatricies(dotMatrixSize int, dotSize int, matricies [][]byte) []byte {
	if dotMatrixSize*dotMatrixSize != len(matricies) {
		panic("dotMatrixSize is not right")
	}

	resultMatrixSize := dotSize * dotMatrixSize
	resultMatrix := make([]byte, 0, resultMatrixSize*resultMatrixSize)

	// NOTE(shadream): algorithm copies values in result matrix by lines in each matrix (copy entire Y line of small matrix instead one element at time).
	for y := 0; y < resultMatrixSize; y++ {
		matrixY := y / dotSize
		insideMatrixIndex := (y % dotSize) * dotSize
		for matrixXIndex := 0; matrixXIndex < dotMatrixSize; matrixXIndex++ {
			matrixIndex := (matrixY * dotMatrixSize) + matrixXIndex

			values := matricies[matrixIndex][insideMatrixIndex : insideMatrixIndex+dotSize]

			resultMatrix = append(resultMatrix, values...)
		}
	}

	return resultMatrix
}
