package golang_leetcode

import (
	"fmt"
	"testing"
)

func TestMouseEscape(t *testing.T) {
	n := 10
	result := [][]float64{
		{1.0, 0.0, 0.0},
		{0.0, 0.0, 0.0},
		{0.0, 0.0, 0.0},
	}  // 初始
	for i:=0; i<n; i++ {
		//v = []float64{0.5 * v[1] + 0.5 * v[2], 0.5 * v[0], 0.5 * v[0], 0.5 * v[1] + 0.5 * v[2] + v[3]}
		result = [][]float64{
			{result[0][1]*1/3 + result[1][0]*1/3, result[0][0]*0.5 + result[0][2]*0.5 + result[1][1]*0.25, result[0][1]*1/3 + result[1][2]*1/3},
			{result[0][0]*0.5 + result[1][1]*0.25 + result[2][0]*0.5, result[0][1]*1/3 + result[1][0]*1/3 + result[1][2]*1/3 + result[2][1]*1/3, result[0][2]*0.5 + result[1][1]*0.25},
			{result[1][0]*1/3 + result[2][1]*1/3, result[2][0]*0.5 + result[1][1]*0.25, result[2][1]*1/3 + result[1][2]*1/3 + result[2][2]},
		}
		fmt.Println(result)
	}
}