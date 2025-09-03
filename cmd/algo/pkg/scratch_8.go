package main

import "fmt"

// // 车辆的接口
// type Car interface {
// 	Drive()
// }
//
// // 具体车辆类型 - 轿车
// type Sedan struct{}
//
// func (s Sedan) Drive() {
// 	fmt.Println("Sedan is driving on the road.")
// }
//
// // 具体车辆类型 - 摩托车
// type Motorcycle struct{}
//
// func (m Motorcycle) Drive() {
// 	fmt.Println("Motorcycle is driving on the road.")
// }
//
// // 工厂方法，创建车辆
// func CreateCar(carType string) Car {
// 	switch carType {
// 	case "Sedan":
// 		return Sedan{}
// 	case "Motorcycle":
// 		return Motorcycle{}
// 	default:
// 		panic("Unknown car type")
// 	}
// }
//
// func main() {
// 	// 创建车辆并驾驶
// 	car := CreateCar("Sedan")
// 	car.Drive()
// }

type Car[T any] interface {
	Drive()
	// Read()
}

type Moto[T any] struct {
}

type Sedan[T any] struct {
}

func (t Moto[T]) Drive() {
	fmt.Println("Driving Moto")
}

func (t Sedan[T]) Drive() {
	fmt.Println("Driving Sedan")
}

func (t Sedan[T]) Read() {
	fmt.Println("Driving Sedan")
}

func NewCar[T any](carType string) Car[T] {
	switch carType {
	case "moto":
		return Moto[T]{}
	case "sedan":
		return Sedan[T]{}
	default:
		return Moto[T]{}
	}
}

func main() {
	// car := NewCar("moto").Drive
	// fmt.Println(car)

	NewCar[int]("moto").Drive()
	NewCar[int]("sedan").Drive()
	NewCar[string]("sedan").Drive()
}
