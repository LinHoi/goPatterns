package main

import  (
	"../car"
	"fmt"
	"github.com/go-acme/lego/log"
)

func main() {
	blueCar := car.NewBuilder().Color(car.BlueColor)
	sportsCar := blueCar.Wheels(car.SportsWheels).TopSpeed(car.KPH).Breaker(car.GoodBreaker).Build()
	if err := sportsCar.Drive(); err != nil {
		log.Println(err)
	}

	if err := sportsCar.Stop(); err != nil {
		log.Println(err)
	}
	fmt.Println()

	steelCar := blueCar.Wheels(car.SteelWheels).TopSpeed(car.GPH).Build()
	if err := steelCar.Drive(); err != nil {
		log.Println(err)
		if err := steelCar.Stop(); err != nil {
			log.Println(err, ",You Dei")
		}
	}

}


