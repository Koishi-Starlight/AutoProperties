package main

import (
	"fmt"
	AutoProperties "main.go/autoProperties"
)

func PrintAutoParams(a AutoProperties.Auto) {
	fmt.Printf("Brand: %v\nModel: %v\nEnginePower: %v\nMaxSpeed: %v\n", a.Brand(), a.Model(), a.EnginePower(), a.MaxSpeed())
	dim := a.Dimensions()
	fmt.Printf("Dimensions: Length %v, Width %v, Height %v", dim.Length(), dim.Width(), dim.Height())
}
func main() {
	bmw1 := AutoProperties.NewBMW(
		"RX7",
		120,
		40,
		AutoProperties.NewUnitCM(100),
		AutoProperties.NewUnitCM(80),
		AutoProperties.NewUnitCM(150),
	)
	PrintAutoParams(bmw1)
}
