package main

import (
	"fmt"
)

func main() {

	plantCapacities := []float64{30, 30, 30, 60, 100}
	activePlants := []int{0, 1}
	gridLoad := 75.

	fmt.Println("1) Generate Power Plan Report")
	fmt.Println("2) Generate Power Grid Report")
	fmt.Println("Please choose an option: ")

	var option string

	fmt.Scanln(&option)

	switch option {
	case "1":
		GeneratePlantCapacityReport(plantCapacities...)

	case "2":
		GeneratePowerGridReport(activePlants, plantCapacities, gridLoad)
	}
}

func GeneratePlantCapacityReport(plantCapacities ...float64) {
	for idx, v := range plantCapacities {
		fmt.Printf("Plant %v Capacity: %v\n", idx, v)

	}
}
func GeneratePowerGridReport(activePlants []int, plantCapacities []float64, gridLoad float64) {
	capacity := 0.
	for _, plantId := range activePlants {
		capacity += plantCapacities[plantId]
	}
	fmt.Println("Capacity: ", capacity)
	fmt.Println("Load: ", gridLoad)
	fmt.Println("Utilisation: ", gridLoad/capacity*100)
}
