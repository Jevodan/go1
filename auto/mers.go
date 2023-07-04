package auto

const (
	MERS string = "mers"
)

var mers = mersCar{}

type mersCar struct {
	Car
	size Dimensions
}

func NewMers(model string, speed int, power int, sizes []float64) mersCar {
	mers.size = NewSizeCM(sizes)
	mers.Car = Car{MERS, model, speed, power}
	return mers
}

func (mers mersCar) Dimensions() Dimensions {
	return mers.size
}
func (mers mersCar) Brand() string {
	return MERS
}
func (mers mersCar) Model() string {
	return mers.Car.Model
}
func (mers mersCar) MaxSpeed() int {
	return mers.Car.MaxSpeed
}
func (mers mersCar) EnginePower() int {
	return mers.Car.EnginePower
}
