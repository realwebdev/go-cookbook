package main

import "fmt"

type Car interface {
	getCar() string
}

type SedanType struct {
	Name string
}

func getNewSedan() *SedanType {
	return &SedanType{}
}

func (c SedanType) getCar() string {
	return "Honda City"
}

type HatchBackType struct {
	Name string
}

func getNewHatchBack() *HatchBackType {
	return &HatchBackType{}
}

func (c HatchBackType) getCar() string {
	return "Vitz"
}

func main() {
	getCarFactory(1)
	getCarFactory(2)
}

func getCarFactory(cartype int) {
	var car Car
	if cartype == 1 {
		car = getNewHatchBack()

	} else {
		car = getNewSedan()
	}

	carInfo := car.getCar()
	fmt.Println(carInfo)
}
