package structs

import "fmt"

type Rectangle struct {
	length float64
}

type Square struct {
	length float64
}

type Area interface {
	CalculateArea() float64
}

func (r Rectangle) CalculateArea() float64 {
	return r.length * r.length
}

func (s Square) CalculateArea() float64 {
	return s.length * s.length
}

type Product struct {
	Name   string `json:"name"`
	Review int    `json:"review"`
	IsUS   bool   `json:"is"`
}

type ProductInterface interface {
	SingleLine() string
	DoubleLine() string
}

func (userInfo *Product) SingleLine() string {
	return fmt.Sprintf("Name: %v Review %v IsUS %v", userInfo.Name, userInfo.Review, userInfo.IsUS)
}
