package car

import (
	"fmt"
	"github.com/go-errors/errors"
)

type Speed float64

const (
	MPH Speed = 1
	KPH       = 1.60934
	GPH       = 3
)

type Color string

const (
	BlueColor  Color = "blue"
	GreenColor       = "green"
	RedColor         = "red"
)

type Wheels string
type Breaker string

const (
	SportsWheels Wheels = "sports wheels"
	SteelWheels         = "steel wheels"
	GoodBreaker Breaker = "goodBreaker"
)


type Builder interface {
	Color(Color) Builder
	Wheels(Wheels) Builder
	TopSpeed(Speed) Builder
	Breaker ( Breaker) Builder
	Build() Interface
}

type Interface interface {
	Drive() error
	Stop() error
}

type car struct {
	color  Color
	wheels Wheels
	topSpeed Speed
	breaker Breaker
}


func (c car) Build() Interface {
	return c
}

func NewBuilder() Builder {
	return &car{}
}



func (c car) Color(color Color) Builder {
	c.color = color
	return c
}

func (c car) Wheels(wheels Wheels) Builder {
	c.wheels = wheels
	return c
}

func (c car) TopSpeed(topSpeed Speed) Builder {
	c.topSpeed = topSpeed
	return c
}


func (c car) Breaker(breaker Breaker) Builder {
	c.breaker = breaker
	return c
}

func (c car) Drive() error {
	fmt.Printf("You are Driving a %v car with %v  at %v\n",c.color,c.wheels,c.topSpeed)
	if c.topSpeed <= 2 {
		return nil
	} else {
		return errors.New("overdrive !!!, police is watching you")
	}
}

func (c car) Stop() error {
	if c.breaker == GoodBreaker {
		fmt.Println("You successful stop the car")
		return nil
	}else {
		return errors.New("Oh my God !!! the breaker is useless")
	}
}