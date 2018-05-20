package game

import (
	"fmt"
	"math"
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

func GetDistanceToLineSegment(point *Point, startLine *Point, endLine *Point) float64 {
	isPerpendicularInLineSegment := func(point *Point, startLine *Point, endLine *Point) bool {
		firstVector := NewVector(point.X-startLine.X, point.Y-startLine.Y)
		secondVector := NewVector(endLine.X-startLine.X, endLine.Y-startLine.Y)
		if GetAngleBetweenVectors(firstVector, secondVector) > math.Pi/2 {
			return false
		}
		firstVector = NewVector(point.X-endLine.X, point.Y-endLine.Y)
		secondVector = NewVector(startLine.X-endLine.X, startLine.Y-endLine.Y)
		if GetAngleBetweenVectors(firstVector, secondVector) > math.Pi/2 {
			return false
		}
		return true
	}

	if isPerpendicularInLineSegment(point, startLine, endLine) {
		x0 := point.X
		y0 := point.Y
		x1 := startLine.X
		y1 := startLine.Y
		x2 := endLine.X
		y2 := endLine.Y
		return math.Abs(((y2-y1)*x0 - (x2-x1)*y0 + x2*y1 - y2*x1) /
			math.Sqrt((y2-y1)*(y2-y1)+(x2-x1)*(x2-x1)))
	}
	return math.Min(point.GetDistance(startLine), point.GetDistance(endLine))
}

func GetAngleBetweenVectors(vector1 *Vector, vector2 *Vector) float64 {
	if vector1.Length == 0 || vector2.Length == 0 {
		return math.Pi / 2
	}
	return math.Acos(GetScalarProduct(vector1, vector2) / (vector1.Length * vector2.Length))
}

func GetScalarProduct(vector1 *Vector, vector2 *Vector) float64 {
	return vector1.GetX()*vector2.GetX() + vector1.GetY()*vector2.GetY()
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
