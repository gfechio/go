package main

import "fmt"

func average(rc ...float64) float64 {
	fmt.Println(rc)
	fmt.Printf("%T \n", rc)

	var total float64
	for _, v := range rc {
		total = +v
	}
	return total / float64(len(rc))
}

func main() {
	n := average(42, 54, 89, 90, 93, 56)
	fmt.Println(n)
}
