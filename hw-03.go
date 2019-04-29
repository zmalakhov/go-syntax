package main

/**
* Syntax Go. Homework 3
* Zaur Malakhov, dated Apr 29, 2019
*/

import (
	"fmt"
)

// общие характеристики
type Car struct {
	Model string
	BodyType string
	Engine string
	Year int
	Transmission string
	Volume int
	IsStarted bool
	IsOpenWindows bool
	Fill int
}

// легковой автомобиль
type PassengerCar struct {
	Car
}

// грузовой автомобиль
type TruckCar struct {
	Car
	Carrying int 
	CabinType string
	Wheels string
}

func main() {


	car1 := PassengerCar{
		Car: Car{
			Model: "LADA",
			BodyType: "sedan",
			Engine: "diesel",
			Year: 2005,
			Transmission: "auto",
			Volume: 3,
			IsStarted: false,
			IsOpenWindows: false,
			Fill: 0,
		},
	}

	car2 := PassengerCar{
		Car: Car{
			Model: "Ford",
			BodyType: "sedan",
			Engine: "gasoline",
			Year: 2003,
			Transmission: "auto",
			Volume: 3,
			IsStarted: false,
			IsOpenWindows: false,
			Fill: 0,
		},
	}

	car3 := TruckCar{
		Car: Car{
			Model: "IVECO",
			BodyType: "truck",
			Engine: "diesel",
			Year: 2010,
			Transmission: "auto",
			Volume: 50,
			IsStarted: false,
			IsOpenWindows: false,
			Fill: 0,
		},
		Carrying: 5000,
		CabinType: "2x2",
		Wheels: "6x4",
	}
	
	car4 := TruckCar{
		Car: Car{
			Model: "MAN",
			BodyType: "truck",
			Engine: "diesel",
			Year: 2015,
			Transmission: "auto",
			Volume: 50,
			IsStarted: false,
			IsOpenWindows: false,
			Fill: 0,
		},
		Carrying: 10000,
		CabinType: "6x2",
		Wheels: "8x8",
	}

	car1.Fill = 1
	car2.IsOpenWindows = true
	car3.Fill = 1500
	car4.IsStarted = true

	fmt.Println(car4.IsStarted)
	fmt.Println(car3.Car)


	fmt.Println(car1)
	fmt.Println(car2)
	fmt.Println(car3)
	fmt.Println(car4)


}