package AutoProperties

type UnitType string

const (
	Inch          UnitType = "inch"
	CM            UnitType = "cm"
	BrandBWM      string   = "BMW"
	BrandDodge             = "Dodge"
	BrandMercedes          = "Mercedes"
)

type Unit struct {
	Value float64
	T     UnitType
}

// Get function takes dimensions, displays them in Inches or CMs
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

// DimensionsType structure
type DimensionsType struct {
	length Unit
	width  Unit
	height Unit
}

func (d DimensionsType) Length() Unit { return d.length }
func (d DimensionsType) Width() Unit  { return d.width }
func (d DimensionsType) Height() Unit { return d.height }

// NewUnitCM function to create auto dimension in CMs
func NewUnitCM(value float64) Unit {
	return Unit{Value: value, T: CM}
}

// NewUnitInch function to create auto dimension in Inches
func NewUnitInch(value float64) Unit {
	return Unit{Value: value, T: Inch}
}

//type DimensionsInch struct {
//	length Unit
//	width  Unit
//	height Unit
//}
//
//func (d DimensionsInch) Length() Unit { return d.length }
//func (d DimensionsInch) Width() Unit  { return d.width }
//func (d DimensionsInch) Height() Unit { return d.height }

// baseAuto structure for other autos.
type baseAuto struct {
	model       string
	maxSpeed    int
	enginePower int
	dimensions  Dimensions
}

// BMW structure template
type BMW struct{ baseAuto }

func (auto BMW) Dimensions() Dimensions { return auto.dimensions }
func (auto BMW) Brand() string          { return BrandBWM }
func (auto BMW) Model() string          { return auto.model }
func (auto BMW) MaxSpeed() int          { return auto.maxSpeed }
func (auto BMW) EnginePower() int       { return auto.enginePower }

func NewBMW(model string, maxSpeed, enginePower int, length, width, height Unit) BMW {
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

// Dodge structure template
type Dodge struct{ baseAuto }

func (auto Dodge) Dimensions() Dimensions { return auto.dimensions }
func (auto Dodge) Brand() string          { return BrandDodge }
func (auto Dodge) Model() string          { return auto.model }
func (auto Dodge) MaxSpeed() int          { return auto.maxSpeed }
func (auto Dodge) EnginePower() int       { return auto.enginePower }

func NewDodge(model string, maxSpeed, enginePower int, height, width, length Unit) Dodge {
	return Dodge{baseAuto{
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

type Mercedes struct{ baseAuto }

func (auto Mercedes) Dimensions() Dimensions { return auto.dimensions }
func (auto Mercedes) Brand() string          { return BrandMercedes }
func (auto Mercedes) Model() string          { return auto.model }
func (auto Mercedes) MaxSpeed() int          { return auto.maxSpeed }
func (auto Mercedes) EnginePower() int       { return auto.enginePower }

func NewMercedes(model string, maxSpeed, enginePower int, height, width, length Unit) Mercedes {
	return Mercedes{baseAuto{
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
