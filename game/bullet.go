package game

type Bullet struct {
	GameObject
	owner *PlayerModel
}

const AREA_RADIUS = 5
const BULLET_SPEED = 15

func NewBullet(position *Point, direction *Vector, owner *PlayerModel) *Bullet {
	return &Bullet{GameObject{
		position: position,
		movementDirection: direction,
		direction: direction,
		areaRadius: AREA_RADIUS,
		isDead: false,
		}, owner}
}

func (bullet *Bullet) Move() {
	bullet.position = bullet.position.AddVector(bullet.movementDirection.MultiplyByScalar(BULLET_SPEED))
}

func (bullet *Bullet) SolveCollision(object IGameObject) {
	if !IsInsideTheMap(bullet.position) {
		bullet.isDead = true
	}
}

func (bullet *Bullet) GetMovementDirection() *Vector {
	return bullet.movementDirection.MultiplyByScalar(BULLET_SPEED)
}

func (bullet *Bullet) CreateNewObject() IGameObject {
	return nil
}
