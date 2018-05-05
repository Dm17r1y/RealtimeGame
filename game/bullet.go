package game

type Bullet struct {
	GameObject
	owner *PlayerModel
}

const AREA_RADIUS = 5
const BULLET_SPEED = 15

func NewBullet(position *Point, direction *Vector, owner *PlayerModel) *Bullet {
	return &Bullet{GameObject{position, direction, AREA_RADIUS,
		false}, owner}
}

func (bullet *Bullet) Move() {
	bullet.position = bullet.position.AddVector(bullet.direction.MultiplyByScalar(BULLET_SPEED))
}

func (bullet *Bullet) SolveCollision(object IGameObject) {
	if !IsInsideTheMap(bullet.position) {
		bullet.isDead = true
	}
}

func (bullet *Bullet) GetVector() *Vector {
	return bullet.GetDirection().MultiplyByScalar(BULLET_SPEED)
}

func (bullet *Bullet) CreateNewObject() IGameObject {
	return nil
}

func (bullet *Bullet) GetTypeName() string {
	return "Bullet"
}
