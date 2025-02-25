package main

import "fmt"

type Car struct {
	Brand string
	Model string
	Year  int
	Price float64
}

type CarBuilder interface {
	SetBrand(brand string) CarBuilder
	SetModel(model string) CarBuilder
	SetYear(year int) CarBuilder
	SetPrice(price float64) CarBuilder
	Build() *Car
}

type carBuilder struct {
	brand string
	model string
	year  int
	price float64
}

func NewCarBuilder() CarBuilder {
	return &carBuilder{}
}

func (c *carBuilder) SetBrand(brand string) CarBuilder {
	c.brand = brand
	return c
}

func (c *carBuilder) SetModel(model string) CarBuilder {
	c.model = model
	return c
}

func (c *carBuilder) SetYear(year int) CarBuilder {
	c.year = year
	return c
}

func (c *carBuilder) SetPrice(price float64) CarBuilder {
	c.price = price
	return c
}

func (c *carBuilder) Build() *Car {
	return &Car{
		Brand: c.brand,
		Model: c.model,
		Year:  c.year,
		Price: c.price,
	}
}

func main() {
	car := NewCarBuilder().
		SetBrand("Toyota").
		SetModel("Corolla").
		SetYear(2021).
		SetPrice(20000.00).
		Build()

	fmt.Printf("Car: %+v\n", car)
}
