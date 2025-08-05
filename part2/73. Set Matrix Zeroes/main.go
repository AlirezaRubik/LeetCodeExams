package main

//import "fmt"

func setZeroes(matrix [][]int) {
	rows := len(matrix)
	cols := len(matrix[0])

	copyMatrixCopy := make([][]int, rows)
	for i := 0; i < rows; i++ {
		copyMatrixCopy[i] = make([]int, cols)
		copy(copyMatrixCopy[i], matrix[i])
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0 {
				for x := 0; x < rows; x++ {
					copyMatrixCopy[x][j] = 0
				}
				for a := 0; a < cols; a++ {
					copyMatrixCopy[i][a] = 0
				}

			}
		}
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			matrix[i][j] = copyMatrixCopy[i][j]
		}
	}
	// for _, row := range matrix {
	// 	fmt.Println(row)
	// }
}
func main() {
	arr := [][]int{
		{1, 1, 1},
		{1, 0, 1},
		{1, 1, 1},
	}
	setZeroes(arr)

}
