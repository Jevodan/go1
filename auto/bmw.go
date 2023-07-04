package auto

const (
	BMW string = "BMW"
)

var bmw = bmvCar{}

type bmvCar struct {
	Car
	size Dimensions
}

func NewBMW(model string, speed int, power int, sizes []float64) bmvCar {
	bmw.size = NewSizeCM(sizes)
	bmw.Car = Car{BMW, model, speed, power}
	return bmw
}

func (bmv bmvCar) Dimensions() Dimensions {
	return bmv.size
}

func (bmv bmvCar) Brand() string {
	return BMW
}
func (bmv bmvCar) Model() string {
	return bmv.Car.Model
}
func (bmv bmvCar) MaxSpeed() int {
	return bmv.Car.MaxSpeed
}
func (bmv bmvCar) EnginePower() int {
	return bmv.Car.EnginePower
}
