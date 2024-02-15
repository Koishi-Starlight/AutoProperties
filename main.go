package main

import (
	"fmt"
	AutoProperties "main.go/autoProperties"
)

func PrintAutoParams(a AutoProperties.Auto) {
	fmt.Printf("Brand: %v\nModel: %v\nEnginePower: %v\nMaxSpeed: %v\n", a.Brand(), a.Model(), a.EnginePower(), a.MaxSpeed())
	dim := a.Dimensions()
	fmt.Printf("Dimensions: Length %v, Width %v, Height %v\n", dim.Length(), dim.Width(), dim.Height())
}
func DisplayInch(a AutoProperties.Auto) {
	fmt.Printf("%v's dimensions:\n", a.Model())
	fmt.Printf("Lenght: %.1f inch\t", a.Dimensions().Length().Get(AutoProperties.Inch))
	fmt.Printf("Width: %.1f inch\t", a.Dimensions().Width().Get(AutoProperties.Inch))
	fmt.Printf("Height: %.1f inch\n", a.Dimensions().Height().Get(AutoProperties.Inch))
}
func main() {

	bmw1 := AutoProperties.NewBMW(
		"M8 Competition Gran Coupe",
		305,
		617,
		AutoProperties.NewUnitCM(510.286),
		AutoProperties.NewUnitCM(194.31),
		AutoProperties.NewUnitCM(141.986),
	)
	dodge1 := AutoProperties.NewDodge(
		"Charger",
		326,
		807,
		AutoProperties.NewUnitInch(198.4),
		AutoProperties.NewUnitInch(75),
		AutoProperties.NewUnitInch(57.8),
	)
	mercedes1 := AutoProperties.NewMercedes(
		"AMG SL Roadster",
		315,
		585,
		AutoProperties.NewUnitCM(407.5),
		AutoProperties.NewUnitCM(191.5),
		AutoProperties.NewUnitCM(135.9),
	)
	dodge2 := AutoProperties.NewDodge(
		"Charger(CM)",
		326,
		807,
		AutoProperties.NewUnitCM(503.9),
		AutoProperties.NewUnitCM(190.5),
		AutoProperties.NewUnitCM(146.8),
	)
	PrintAutoParams(bmw1)
	PrintAutoParams(dodge1)
	PrintAutoParams(mercedes1)
	DisplayInch(dodge1)
	DisplayInch(dodge2)
	DisplayInch(bmw1)
	DisplayInch(mercedes1)
}
