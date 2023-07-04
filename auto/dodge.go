package auto

const (
	DODGE string = "dodge"
)

var dodge = dodgeCar{}

type dodgeCar struct {
	Car
	size Dimensions
}

func NewDodge(model string, speed int, power int, sizes []float64) dodgeCar {
	dodge.size = NewSizeInch(sizes)
	dodge.Car = Car{DODGE, model, speed, power}
	return dodge
}

func (dodge dodgeCar) Dimensions() Dimensions {
	return dodge.size
}

func (dodge dodgeCar) Brand() string {
	return DODGE
}
func (dodge dodgeCar) Model() string {
	return dodge.Car.Model
}
func (dodge dodgeCar) MaxSpeed() int {
	return dodge.Car.MaxSpeed
}
func (dodge dodgeCar) EnginePower() int {
	return dodge.Car.EnginePower
}
