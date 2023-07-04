package main

import (
	"car/auto"
	"fmt"
)

func main() {
	fmt.Println("----------------Cars-------------")
	sizeBMW := []float64{432.4, 186.4, 130.4}
	sizeMers := []float64{460.6, 198.4, 196.9}
	sizeDodge := []float64{475.5 / auto.D, 174 / auto.D, 134.6 / auto.D}
	bmw := auto.NewBMW("BMW Z4", 280, 200, sizeBMW)
	mers := auto.NewMers("G-Class", 230, 300, sizeMers)
	dodge := auto.NewDodge("Avenger", 180, 140, sizeDodge)

	cars := []auto.Auto{bmw, mers, dodge}

	for _, val := range cars {
		printCharacteristics(val)
	}

}

func printCharacteristics(car auto.Auto) {
	fmt.Println("Model: ", car.Model(), "     Speed: ", car.MaxSpeed())
	fmt.Println("Brand: ", car.Brand(), "        Power: ", car.EnginePower())
	fmt.Println("---------------Gabarits-------------")
	fmt.Println("Length: ", car.Dimensions().Length().Value, " ", car.Dimensions().Width().T)
	fmt.Println("Width: ", car.Dimensions().Width().Value, " ", car.Dimensions().Width().T)
	fmt.Println("Height: ", car.Dimensions().Height().Value, " ", car.Dimensions().Height().T)
	fmt.Println("-------------------------------------")
}
