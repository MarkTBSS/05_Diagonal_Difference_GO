package main

import (
	"fmt"
)

func main() {
	var n = 3
	var arr [][]int32
	arr = [][]int32{
		{11, 2, 4},
		{4, 5, 6},
		{10, 8, -12}}

	var leftDiagonalSum int32
	var rightDiagonalSum int32
	var notAbsoluteOfLeftAndRight int32
	var absoluteOfLeftAndRight int32

	for i := 0; i < n; i++ {
		/* for j := 0; j < n; j++ {
			fmt.Printf("%d ", arr[i][j])
		} */
		leftDiagonalSum += arr[i][i]
		rightDiagonalSum += arr[i][n-1-i]
	}
	notAbsoluteOfLeftAndRight = leftDiagonalSum - rightDiagonalSum
	//absoluteOfLeftAndRight = int32(math.Abs(float64(notAbsoluteOfLeftAndRight)))
	if notAbsoluteOfLeftAndRight < 0 {
		absoluteOfLeftAndRight = -notAbsoluteOfLeftAndRight
	} else {
		absoluteOfLeftAndRight = notAbsoluteOfLeftAndRight
	}

	fmt.Printf("%d", absoluteOfLeftAndRight)
}
