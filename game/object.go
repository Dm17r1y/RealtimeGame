package game

type IGameObject interface {
	Move()
	SolveCollision(otherObject IGameObject)
	IsDead() bool
	GetPosition() *Point
	GetDirection() *Vector
	GetVector() *Vector
	GetAreaRadius() float64
	CreateNewObject() IGameObject
	GetTypeName() string
}

type GameObject struct {
	position   *Point
	direction  *Vector
	areaRadius float64
	isDead     bool
}

func (object *GameObject) IsDead() bool {
	return object.isDead
}

func (object *GameObject) GetPosition() *Point {
	return object.position
}

func (object *GameObject) GetDirection() *Vector {
	return object.direction
}

func (object *GameObject) GetAreaRadius() float64 {
	return object.areaRadius
}
