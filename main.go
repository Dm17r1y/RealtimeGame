package main

import (
	"./game"
	"math"
	"sync"
	"time"
)

var DELAY = time.Millisecond * 1000 / 60

func main() {
	g := game.NewGame()
	go startServer()

	ticker := time.NewTicker(DELAY)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ticker.C:
				newControllers := ReadFromConnections(Controllers, g)
				UpdateConnections(newControllers)
				g.MakeTurn()
				WriteToConnections(Controllers, g)
			}
		}
	}()
	wg.Wait()
}

func ReadFromConnections(controllers []*Controller, game *game.Game) []*Controller {
	aliveConnections := make([]*Controller, 0)
	for i := 0; i < len(controllers); i++ {
		controller := controllers[i]
		breakFlag := false
		for !breakFlag {
			select {
			case message, ok := <-controller.readChannel:
				if !ok {
					breakFlag = true
					break
				}
				ApplyCommand(controller, game, message)
			default:
				aliveConnections = append(aliveConnections, controller)
				breakFlag = true
			}
		}
	}

	return aliveConnections
}

func ApplyCommand(controller *Controller, g *game.Game, data *MessageData) {
	if controller.model == nil || controller.model.IsDead() {
		weapon := game.NewWeapon(func(model *game.PlayerModel) game.IGameObject {
			return game.NewFastBullet(model.GetPosition(), model.GetDirection(), model)
		}, 60)
		controller.model = game.NewPlayer(g.GetRandomPoint(), g, weapon)
		g.AddGameObject(controller.model)
	}
	direction := game.NewVector(data.MouseLocation.X-controller.model.GetPosition().X,
		data.MouseLocation.Y-controller.model.GetPosition().Y).Normalize()
	controller.model.SetDirection(direction)
	// Old
	//up := &game.Vector{AngleInRadian: direction.AngleInRadian, Length: direction.Length}
	//down := &game.Vector{AngleInRadian: direction.AngleInRadian + math.Pi, Length: direction.Length}
	//right := &game.Vector{AngleInRadian: direction.AngleInRadian + math.Pi/2, Length: direction.Length}
	//left := &game.Vector{AngleInRadian: direction.AngleInRadian + 3*math.Pi/2, Length: direction.Length}
	up := &game.Vector{AngleInRadian: -math.Pi / 2, Length: direction.Length}
	down := &game.Vector{AngleInRadian: math.Pi / 2, Length: direction.Length}
	left := &game.Vector{AngleInRadian: math.Pi, Length: direction.Length}
	right := &game.Vector{AngleInRadian: 0, Length: direction.Length}
	vector := &game.Vector{AngleInRadian: 0, Length: 0}
	if data.Movement.Up {
		vector = vector.Add(up)
	}
	if data.Movement.Down {
		vector = vector.Add(down)
	}
	if data.Movement.Left {
		vector = vector.Add(left)
	}
	if data.Movement.Right {
		vector = vector.Add(right)
	}
	controller.model.SetMovement(vector.Normalize())

	if data.Shoot {
		controller.model.SetCommand(&game.Command{&game.Shoot{data.ClientTick}})
	}
}

func WriteToConnections(controllers []*Controller, g *game.Game) {
	states := GetStates(g)
	for i := 0; i < len(controllers); i++ {
		controller := controllers[i]
		select {
		case controller.writeChannel <- GetGlobalState(states, controller.model, g.Tick):
		default:
		}
	}
}
