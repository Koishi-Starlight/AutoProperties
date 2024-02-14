package main

import (
	"fmt"
	AutoProperties "main.go/autoProperties"
)

func PrintAutoParams(a AutoProperties.Auto) {
	fmt.Printf("Brand: %v\nModel: %v\nEnginePower: %v\nMaxSpeed: %v\n", a.Brand(), a.Model(), a.EnginePower(), a.MaxSpeed())
	fmt.Printf("Dimensions: Length, Width, Height %v", a.Dimensions())
}
func main() {
	bmw1 := AutoProperties.NewBMW(
		"RX7",
		120,
		40,
		120,
		90,
		100)
	PrintAutoParams(bmw1)
}
