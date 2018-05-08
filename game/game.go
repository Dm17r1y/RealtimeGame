package game

import (
	"math/rand"
	"time"
)

type Game struct {
	GameObjects []IGameObject
}

const MAX_WIDTH = 1000
const MAX_HEIGHT = 500

func NewGame() *Game {
	return &Game{GameObjects: make([]IGameObject, 0)}
}

func (game *Game) AddGameObject(gameObject IGameObject) {
	game.GameObjects = append(game.GameObjects, gameObject)
}

func (game *Game) isCollide(object IGameObject, otherObject IGameObject) bool {
	return object.GetPosition().GetDistance(otherObject.GetPosition()) <= object.GetAreaRadius()
}

func (game *Game) GetRandomPoint() *Point {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(MAX_WIDTH)
	y := rand.Intn(MAX_HEIGHT)
	return &Point{float64(x), float64(y)}
}

func IsInsideTheMap(point *Point) bool {
	return point.X <= MAX_WIDTH && point.X >= 0 && point.Y <= MAX_HEIGHT && point.Y >= 0
}

func Bound(point *Point) *Point {
	return &Point{normallize(point.X, 0, MAX_WIDTH), normallize(point.Y, 0, MAX_HEIGHT)}
}

func normallize(num float64, min float64, max float64) float64 {
	if num < min {
		return min
	}
	if num > max {
		return max
	}
	return num
}

func (game *Game) MakeTurn() {

	objects := make([]IGameObject, len(game.GameObjects))
	copy(objects, game.GameObjects)

	for _, object := range objects {
		object.Move()
	}

	newObjects := make([]IGameObject, 0)

	for _, object := range objects {
		newObject := object.CreateNewObject()
		if newObject != nil {
			newObjects = append(newObjects, newObject)
		}
	}

	for _, object := range newObjects {
		objects = append(objects, object)
	}

	for i, object := range objects {
		for j, otherObject := range objects {
			if i == j {
				continue
			}
			if game.isCollide(object, otherObject) {
				object.SolveCollision(otherObject)
			}
		}
	}

	newObjects = make([]IGameObject, 0)
	for _, object := range objects {
		if !object.IsDead() {
			newObjects = append(newObjects, object)
		}
	}

	game.GameObjects = newObjects
}
