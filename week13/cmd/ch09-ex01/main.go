package main

import "fmt"

type Meters float64
type Kilometers float64
type Miles float64

func (l Kilometers) ToMiles() Miles {
	return Miles(l * 0.621371)
}
func (m Meters) ToMiles() Miles {
	return Miles(m * 0.000621371)
}

func main() {
	kmph := Kilometers(150)
	fmt.Printf("%0.2f kilo meter per hour equals %0.2f mile per hour\n", kmph, kmph.ToMiles())
	meter := Meters(150000)
	fmt.Printf("%0.3f meter %0.3f mile\n", meter, meter.ToMiles())
}
