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
}

type GlobalState struct {
	states   []*State
	position *game.Point
}

func GetGlobalState(g *game.Game, model *game.PlayerModel) *GlobalState {
	states := make([]*State, 0)
	for i := 0; i < len(g.GameObjects); i++ {
		object := g.GameObjects[i]
		states = append(states, &State{point: object.GetPosition(), direction: object.GetVector(),
			objectType: object.GetTypeName()})
	}
	var position *game.Point
	if model == nil {
		position = &game.Point{X: 0, Y: 0}
	} else {
		position = model.GetPosition()
	}
	return &GlobalState{states, position}
}

func (state *GlobalState) ToJsonMap() map[string]interface{} {
	jsonMap := make(map[string]interface{})
	jsonMap["myLocation"] = map[string]int{"x": int(state.position.X), "y": int(state.position.Y)}
	data := make([]map[string]interface{}, 0)
	for i := 0; i < len(state.states); i++ {
		currentState := state.states[i]
		stateMap := make(map[string]interface{})
		stateMap["objectType"] = currentState.objectType
		stateMap["point"] = map[string]int{"x": int(currentState.point.X), "y": int(currentState.point.Y)}
		stateMap["direction"] = map[string]int{"x": int(currentState.direction.GetX()),
			"y": int(currentState.direction.GetY())}
		data = append(data, stateMap)
	}
	jsonMap["render"] = data
	return jsonMap
}
