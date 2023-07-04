package auto

type UnitType string

const (
	Inch UnitType = "inch"
	CM   UnitType = "cm"
	D    float64  = 2.54
)

type Unit struct {
	Value float64
	T     UnitType
}

func (u Unit) Get(t UnitType) float64 {
	value := u.Value
	if t != u.T {
		switch u.T {
		case CM:
			value = u.Value / D
		case Inch:
			value = u.Value * D
		default:
			panic("Неопознанная единица измерения")
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
