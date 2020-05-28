package _map

import "fmt"

func UsageOfMap() {

	m1 := map[int]int{}
	m1[0] = 100
	m1[1] = 200
	fmt.Printf("%v \n", m1)

	m2 := make(map[float64]float64)
	m2[0.1] = 17.7
	m2[1.0] = 0.06
	fmt.Printf("%v \n", m2)
}
