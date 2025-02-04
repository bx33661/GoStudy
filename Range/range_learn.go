package main

import (
	"fmt"
	"math"
)

func main() {
	pow := []int{1, 2, 3, 4, 5, 6}
	for i, v := range pow {
		//fmt.Println(i, 2*pow[v-1])
		result := math.Pow(2, float64(v))
		fmt.Println(i, result)
	}
	map1 := make(map[int]float32)

	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0
	map1[4] = 4.0
	for k := range map1 {
		fmt.Println(k, map1[k])
	}
	for _, v := range map1 {
		fmt.Println(v)
	}

}
