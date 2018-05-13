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
	return &PlayerModel{GameObject {
		position: point,
		movementDirection: NewVector(0, 0),
		direction: NewVector(0, 0),
		areaRadius: PLAYER_RADIUS,
		isDead: false,
	}, currentGame, &defaultCommand}
}

func (model *PlayerModel) SetCommand(command *Command) {
	model.command = command
}

func (model *PlayerModel) SetMovement(direction *Vector) {
	model.movementDirection = direction
}

func (model *PlayerModel) SetDirection(direction *Vector) {
	model.direction = direction
}

func (model *PlayerModel) Move() {
	model.position = Bound(model.position.AddVector(model.movementDirection.MultiplyByScalar(float64(PLAYER_SPEED))))
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

func (model *PlayerModel) GetMovementDirection() *Vector {
	return model.movementDirection.MultiplyByScalar(PLAYER_SPEED)
}

func (model *PlayerModel) CreateNewObject() IGameObject {
	shoot := model.command.Shoot
	if shoot != nil {
		model.command.Shoot = nil
		return NewBullet(model.position, model.direction, model)
	}
	return nil
}
