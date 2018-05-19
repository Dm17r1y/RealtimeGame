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
	owner *PlayerModel
}

func NewFastBullet(position *Point, shootDirection *Vector, owner *PlayerModel) *FastBullet {
	return &FastBullet{
		position: position,
		shootDirection: shootDirection.Normalize(),
		lifeTime: FAST_BULLET_LIVE_TIME,
		owner: owner,
	}
}

func (bullet *FastBullet) Move() {
	bullet.lifeTime--
}

func (bullet *FastBullet) SolveCollision(otherObject IGameObject) {
	switch player := otherObject.(type) {
	case *PlayerModel:
		if player == bullet.owner {
			return
		}
		if bullet.lifeTime != FAST_BULLET_LIVE_TIME {
			return
		}
		startPoint := bullet.position
		endPoint := bullet.position.AddVector(bullet.GetDirection().MultiplyByScalar(FAST_BULLET_SHOOT_RANGE))
		if GetDistanceToLineSegment(player.GetPosition(), startPoint, endPoint) <= player.GetAreaRadius() {
			player.Die()
		}
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

func (bullet *FastBullet) Die() {

}