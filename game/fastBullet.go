package game

import (
	"math"
)

const FAST_BULLET_LIVE_TIME = 30
const FAST_BULLET_SHOOT_RANGE = 2000

type FastBullet struct {
	position *Point
	shootDirection *Vector
	lifeTime int
}

func NewFastBullet(position *Point, shootDirection *Vector) *FastBullet {
	return &FastBullet{
		position: position,
		shootDirection: shootDirection.Normalize(),
		lifeTime: FAST_BULLET_LIVE_TIME,
	}
}

func (bullet *FastBullet) Move() {
	bullet.lifeTime--
}

func (bullet *FastBullet) SolveCollision(otherObject IGameObject) {
	if bullet.lifeTime != FAST_BULLET_LIVE_TIME - 1 {
		return
	}
	startPoint := bullet.position
	endPoint := bullet.position.AddVector(bullet.GetDirection().MultiplyByScalar(FAST_BULLET_SHOOT_RANGE))
	if GetDistanceToLineSegment(otherObject.GetPosition(), startPoint, endPoint) <= otherObject.GetAreaRadius() {
		otherObject.Die()
	}
}

func (bullet *FastBullet) IsDead() bool {
	return bullet.lifeTime <= 0
}

func (bullet *FastBullet) GetPosition() *Point {
	return bullet.position
}

func (bullet *FastBullet) GetMovementDirection() *Vector {
	return &Vector{0, 0}
}

func (bullet *FastBullet) GetDirection() *Vector {
	return bullet.shootDirection
}

func (bullet *FastBullet) GetAreaRadius() float64 {
	return math.MaxFloat64
}

func (bullet *FastBullet) CreateNewObject() IGameObject {
	return nil
}

func (bullet *FastBullet) Die() {

}