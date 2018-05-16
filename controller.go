package main

import (
	"./game"
	"sync"
)

type Controller struct {
	model        *game.PlayerModel
	readChannel  chan *MessageData
	writeChannel chan *GlobalState
}

var ControllersLocker sync.Mutex
var Controllers = make([]*Controller, 0)

func UpdateConnections(newControllers []*Controller) {
	ControllersLocker.Lock()
	defer ControllersLocker.Unlock()
	Controllers = newControllers
}

func AddController(readChannel chan *MessageData, writeChannel chan *GlobalState) {
	ControllersLocker.Lock()
	defer ControllersLocker.Unlock()
	Controllers = append(Controllers, &Controller{nil, readChannel, writeChannel})
}

type MovementData struct {
	Up    bool
	Down  bool
	Left  bool
	Right bool
}

type MessageData struct {
	Movement      *MovementData
	MouseLocation *game.Point
	Shoot         bool
}

func NewMessageData(jsonData map[string]interface{}) *MessageData {
	movementJSON := jsonData["movement"].(map[string]interface{})
	shoot := jsonData["shoot"].(bool)
	mouseLocationJSON := jsonData["mouseLocation"].(map[string]interface{})

	mouseLocation := game.Point{X: mouseLocationJSON["x"].(float64), Y: mouseLocationJSON["y"].(float64)}

	return &MessageData{&MovementData{Up: movementJSON["up"].(bool), Down: movementJSON["down"].(bool),
		Left: movementJSON["left"].(bool), Right: movementJSON["right"].(bool)}, &mouseLocation,
		shoot}
}

type State struct {
	objectType string
	point      *game.Point
	direction  *game.Vector
	object 	   game.IGameObject
}

type GlobalState struct {
	states       []*State
	indexInArray int
	position     *game.Point
	tickNumber   int
}

func GetStates(g *game.Game) []*State {

	states := make([]*State, 0)
	for i := 0; i < len(g.GameObjects); i++ {
		object := g.GameObjects[i]
		var objectType string

		switch object.(type) {
		case *game.Bullet:
			objectType = "Bullet"
		case *game.FastBullet:
			objectType = "FastBullet"
		case *game.PlayerModel:
			objectType = "Player"
		}
		currentState := &State{
			point: object.GetPosition(),
			direction: object.GetDirection(),
			objectType: objectType,
			object: object,
		}
		states = append(states, currentState)
	}
	return states
}

func GetGlobalState(states []*State, model *game.PlayerModel, tickNumber int) *GlobalState {
	indexInArray := -1
	for i := 0; i < len(states); i++ {
		if states[i].object == model {
			indexInArray = i
			break
		}
	}
	var position *game.Point
	if model == nil {
		position = &game.Point{X: 0, Y: 0}
	} else {
		position = model.GetPosition()
	}
	return &GlobalState{states, indexInArray, position, tickNumber}
}

func (state *GlobalState) ToJsonMap() map[string]interface{} {
	jsonMap := make(map[string]interface{})
	jsonMap["stateIndex"] = state.indexInArray
	data := make([]map[string]interface{}, 0)
	for i := 0; i < len(state.states); i++ {
		currentState := state.states[i]
		stateMap := make(map[string]interface{})
		stateMap["objectType"] = currentState.objectType
		stateMap["position"] = map[string]int{
			"x": int(currentState.point.X),
			"y": int(currentState.point.Y),
		}
		stateMap["direction"] = currentState.direction.AngleInRadian
		data = append(data, stateMap)
	}
	jsonMap["render"] = data
	jsonMap["tick"] = state.tickNumber
	return jsonMap
}
