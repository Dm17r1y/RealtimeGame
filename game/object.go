package game

type IGameObject interface {
	Move()
	SolveCollision(otherObject IGameObject)
	IsDead() bool
	GetPosition() *Point
	GetMovementDirection() *Vector
	GetDirection() *Vector
	GetAreaRadius() float64
	CreateNewObject() IGameObject
}

type GameObject struct {
	position          *Point
	movementDirection *Vector
	direction         *Vector
	areaRadius        float64
	isDead            bool
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
