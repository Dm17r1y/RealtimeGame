package game

import (
	"math"
	"fmt"
)

type Point struct {
	X float64
	Y float64
}

func (point *Point) GetDistance(otherPoint *Point) float64 {
	return math.Sqrt(float64((point.X-otherPoint.X)*(point.X-otherPoint.X) +
		(point.Y-otherPoint.Y)*(point.Y-otherPoint.Y)))
}

func (point *Point) AddVector(vector *Vector) *Point {
	p := &Point{point.X + float64(vector.GetX()), point.Y + float64(vector.GetY())}
	return p
}

func (point Point) String() string {
	return fmt.Sprintf("(%.1f, %.1f)", point.X, point.Y)
}

type Vector struct {
	AngleInRadian float64
	Length        float64
}

func NewVector(x float64, y float64) *Vector {
	length := math.Sqrt(x*x + y*y)
	if length == 0.0 {
		return &Vector{0, 0}
	}
	return &Vector{math.Atan2(y, x), length}
}

func Round(val float64) int {
	if val < 0 {
		return int(val - 0.5)
	}
	return int(val + 0.5)
}

func (vector *Vector) GetX() float64 {
	return math.Cos(vector.AngleInRadian) * vector.Length
}

func (vector *Vector) GetY() float64 {
	return math.Sin(vector.AngleInRadian) * vector.Length
}

func (vector *Vector) MultiplyByScalar(scalar float64) *Vector {
	return &Vector{vector.AngleInRadian, vector.Length * scalar}
}

func (vector Vector) String() string {
	return fmt.Sprintf("Vector(%.1f, %.1f)", vector.GetX(), vector.GetY())
}

func (vector *Vector) Normalize() *Vector {
	if vector.Length == 0 {
		return NewVector(0, 0)
	}
	return &Vector{vector.AngleInRadian, 1}
}

func (vector *Vector) Add(secondVector *Vector) *Vector {
	return NewVector(vector.GetX()+secondVector.GetX(), vector.GetY()+secondVector.GetY())
}
