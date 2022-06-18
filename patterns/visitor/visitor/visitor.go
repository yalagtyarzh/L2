package visitor

import "fmt"

// Shape - интерфейс, представляющий собой фигуру
type Shape interface {
	GetType() string
	Accept(Visitor)
}

// Square - класс, представляющий собой конкретную фигуру квадрата
type Square struct {
	Side int
}

// Accept - метод, позволяющий осуществить логику посетителя для квадрата
func (s *Square) Accept(v Visitor) {
	v.VisitSquare(s)
}

// GetType возвращает тип фигуры в письменном виде
func (s *Square) GetType() string {
	return "Square"
}

// Circle - класс, представляющий собой конкретную фигуру круга
type Circle struct {
	Radius int
}

// Accept - метод, позволяющий осуществить логику посетителя для круга
func (c *Circle) Accept(v Visitor) {
	v.VisitCircle(c)
}

// GetType возвращает тип фигуры в письменном виде
func (c *Circle) GetType() string {
	return "Circle"
}

// Rectangle - класс, представляющий собой конкретную фигуру круга
type Rectangle struct {
	L int
	B int
}

// Accept - метод, позволяющий осуществить логику посетителя для прямоугольника
func (r *Rectangle) Accept(v Visitor) {
	v.VisitRectangle(r)
}

// GetType возвращает тип фигуры в письменном виде
func (r *Rectangle) GetType() string {
	return "Rectangle"
}

// Visitor - интерфейс, представляющий собой посетителя
type Visitor interface {
	VisitSquare(s *Square)
	VisitCircle(c *Circle)
	VisitRectangle(r *Rectangle)
}

// AreaCalculator - класс, представляющий собой конкретного посетителя, который умеет вычислять площадь фигуры
type AreaCalculator struct {
	Area int
}

func (a *AreaCalculator) VisitSquare(s *Square) {
	fmt.Println("Calculating area for square")
}

func (a *AreaCalculator) VisitCircle(s *Circle) {
	fmt.Println("Calculating area for circle")
}

func (a *AreaCalculator) VisitRectangle(s *Rectangle) {
	fmt.Println("Calculating area for rectangle")
}

// MiddleCoordinates - класс, представляющий собой конкретного посетителя, который умеет вычислять координаты центральной точки фигуры
type MiddleCoordinates struct {
	X int
	Y int
}

func (a *MiddleCoordinates) VisitSquare(s *Square) {
	fmt.Println("Calculating middle point coordinates for square")
}

func (a *MiddleCoordinates) VisitCircle(s *Circle) {
	fmt.Println("Calculating middle point coordinates for circle")
}

func (a *MiddleCoordinates) VisitRectangle(s *Rectangle) {
	fmt.Println("Calculating middle point coordinates for rectangle")
}
