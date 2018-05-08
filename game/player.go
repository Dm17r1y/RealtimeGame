package game

type PlayerModel struct {
	GameObject
	currentGame *Game
	command     *Command
}

const PLAYER_RADIUS = 25
const PLAYER_SPEED = 3

var defaultCommand = Command{nil}

func NewPlayer(point *Point, currentGame *Game) *PlayerModel {
	return &PlayerModel{GameObject{point, NewVector(0, 0),
		PLAYER_RADIUS, false}, currentGame, &defaultCommand}
}

func (model *PlayerModel) SetCommand(command *Command) {
	model.command = command
}

func (model *PlayerModel) SetDirection(direction *Vector) {
	model.direction = direction
}

func (model *PlayerModel) Move() {
	model.position = Bound(model.position.AddVector(model.direction.MultiplyByScalar(float64(PLAYER_SPEED))))
}

func (model *PlayerModel) SolveCollision(object IGameObject) {
	switch object.(type) {
	case *Bullet:
		bullet := object.(*Bullet)
		if bullet.owner != model {
			model.isDead = true
		}
	}
}

func (model *PlayerModel) GetVector() *Vector {
	return model.GetDirection().MultiplyByScalar(PLAYER_SPEED)
}

func (model *PlayerModel) CreateNewObject() IGameObject {
	shoot := model.command.Shoot
	if shoot != nil {
		model.command.Shoot = nil
		return NewBullet(model.position, shoot.Vector, model)
	}
	return nil
}
