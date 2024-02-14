package AutoProperties

type UnitType string

const (
	Inch UnitType = "inch"
	CM   UnitType = "cm"
)

type Unit struct {
	Value float64
	T     UnitType
}

func (u Unit) Get(t UnitType) float64 {
	value := u.Value

	if t != u.T {
		if t == Inch {
			value *= 0.393701
		} else {
			value /= 0.393701
		}
	}
	return value
}

type Dimensions interface {
	Length() Unit
	Width() Unit
	Height() Unit
}
type Auto interface {
	Brand() string
	Model() string
	Dimensions() Dimensions
	MaxSpeed() int
	EnginePower() int
}

type DimensionsType struct {
	length Unit
	width  Unit
	height Unit
}

func (d DimensionsType) Length() Unit { return d.length }
func (d DimensionsType) Width() Unit  { return d.width }
func (d DimensionsType) Height() Unit { return d.height }
func NewUnitCM(value float64) Unit {
	return Unit{Value: value, T: CM}
}
func NewUnitInch(value float64) Unit {
	return Unit{Value: value, T: Inch}
}

type baseAuto struct {
	model       string
	maxSpeed    int
	enginePower int
	dimensions  Dimensions
}

type BMW struct{ baseAuto }

func (auto BMW) Dimensions() Dimensions { return auto.dimensions }
func (auto BMW) Brand() string          { return "BMW" }
func (auto BMW) Model() string          { return auto.model }
func (auto BMW) MaxSpeed() int          { return auto.maxSpeed }
func (auto BMW) EnginePower() int       { return auto.enginePower }

func NewBMW(model string, maxSpeed, enginePower int, height, width, length Unit) BMW {
	return BMW{baseAuto{
		model:       model,
		maxSpeed:    maxSpeed,
		enginePower: enginePower,
		dimensions: DimensionsType{
			height: height,
			width:  width,
			length: length,
		},
	}}
}
