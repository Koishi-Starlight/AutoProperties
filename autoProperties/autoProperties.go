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
			u.T = Inch
		} else {
			value /= 0.393701
			u.T = CM
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

type DimensionsInch struct {
	length Unit
	width  Unit
	height Unit
}

func (d DimensionsInch) Length() Unit { return d.length }
func (d DimensionsInch) Width() Unit  { return d.width }
func (d DimensionsInch) Height() Unit { return d.height }

type DimensionsCM struct {
	length Unit
	width  Unit
	height Unit
}

func (d DimensionsCM) Length() Unit { return d.length }
func (d DimensionsCM) Width() Unit  { return d.width }
func (d DimensionsCM) Height() Unit { return d.height }

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

func NewBMW(model string, maxSpeed, enginePower int, height, width, length float64) BMW {
	return BMW{baseAuto{
		model:       model,
		maxSpeed:    maxSpeed,
		enginePower: enginePower,
		dimensions: DimensionsCM{
			height: Unit{Value: height, T: CM},
			width:  Unit{Value: width, T: CM},
			length: Unit{Value: length, T: CM},
		},
	}}
}
