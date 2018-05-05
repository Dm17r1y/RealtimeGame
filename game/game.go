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

	for i := 0; i < len(objects); i++ {
		object := objects[i]
		object.Move()
	}

	newObjects := make([]IGameObject, 0)
	for i := 0; i < len(objects); i++ {
		object := objects[i]
		newObject := object.CreateNewObject()
		if newObject != nil {
			newObjects = append(newObjects, newObject)
		}
	}

	for i := 0; i < len(newObjects); i++ {
		objects = append(objects, newObjects[i])
	}

	for i := 0; i < len(objects); i++ {
		object := objects[i]
		for j := 0; j < len(game.GameObjects); j++ {
			if i == j {
				continue
			}
			otherObject := game.GameObjects[j]
			if game.isCollide(object, otherObject) {
				object.SolveCollision(otherObject)
			}
		}
	}

	newObjects = make([]IGameObject, 0)
	for i := 0; i < len(objects); i++ {
		object := objects[i]
		if !object.IsDead() {
			newObjects = append(newObjects, objects[i])
		}
	}

	game.GameObjects = newObjects
}
