package main

import (
	"fmt"

	"builder/creation"
)

func main() {
	// Создаем директора и строителя
	director := creation.ManufacturingDirector{}
	cb := creation.CarBuilder{}
	// Отныне директор отвечает за строителя автомобиля
	director.SetBuilder(&cb)
	// Строим автомобиль
	director.Construct()

	// Получаем собранную тачку и выводим на экран
	car := cb.GetVehicle()
	fmt.Println(car)

	// Создаем строителя мотоциклов
	bb := creation.BikeBuilder{}
	// Меняем строителя у директора и строим мопед
	director.SetBuilder(&bb)
	director.Construct()

	// Получаем собранный KTM и выводим на экран
	bike := cb.GetVehicle()
	fmt.Println(bike)
}
