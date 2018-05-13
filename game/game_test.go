package game

import (
	"testing"
	"fmt"
)

func TestVector(t *testing.T) {
	v := NewVector(40, 0)
	if v.Length != 40.0 {
		t.Error("Wrong length")
	}
	if v.GetX() != 40 || v.GetY() != 0 {
		t.Error(fmt.Sprintf("Wrong vector: %s, expected Vector(40.0, 0.0)", v))
	}
}

func TestMultiplyByScalar(t *testing.T) {
	v := NewVector(1, 0)
	v2 := v.MultiplyByScalar(40)
	if v2.GetX() != 40 || v2.GetY() != 0 {
		t.Error(fmt.Sprintf("Wrong vector: %s, expected Vector(40.0, 0.0)", v))
	}
}

func TestAddVectorToPoint(t *testing.T) {
	v := NewVector(40, 0)
	p := Point{0, 0}
	p2 := p.AddVector(v)
	if p2.X != 40 || p2.Y != 0 {
		t.Error(fmt.Sprintf("Wrong point: %s, expected (40, 0)", p2))
	}
}

func TestBulletMovement(t *testing.T) {
	g := NewGame()
	bullet := NewBullet(&Point{0, 0}, NewVector(1, 0), nil)
	g.AddGameObject(bullet)
	g.MakeTurn()
	if bullet.GetPosition().X != BULLET_SPEED || bullet.GetPosition().Y != 0 {
		t.Error(fmt.Sprintf("Expected (%.1f, 0.0), got %s", float64(BULLET_SPEED), bullet.GetPosition()))
	}
}

func TestPlayerMovement(t *testing.T) {
	g := NewGame()
	player := NewPlayer(&Point{0, 0}, g)
	g.AddGameObject(player)
	player.movementDirection = NewVector(1, 0)
	g.MakeTurn()
	if player.GetPosition().X != PLAYER_SPEED || player.GetPosition().Y != 0 {
		t.Error(fmt.Sprintf("Expected (%.1f, 0.0), got %s", float64(BULLET_SPEED), player.GetPosition()))
	}
}

func TestPlayerShoot(t *testing.T) {
	g := NewGame()
	player := NewPlayer(&Point{0, 0}, g)
	player.command = &Command{Shoot: &Shoot{NewVector(1, 0)}}
	g.AddGameObject(player)
	g.MakeTurn()
	if len(g.GameObjects) != 2 {
		t.Error(fmt.Sprintf("Wrong game objects length: %d", len(g.GameObjects)))
		return
	}
	if !contains(g.GameObjects, player) {
		t.Error("Player not in objects")
		return
	}
	bullet := g.GameObjects[1].(*Bullet)
	vector := bullet.GetMovementDirection()
	if vector.GetX() != BULLET_SPEED || vector.GetY() != 0 {
		t.Error(fmt.Sprintf("Expected bullet position Vector(%.1f, 0.0), got %s", float64(BULLET_SPEED), vector))
		return
	}
	if bullet.owner != player {
		t.Error("Wrong bullet owner")
		return
	}
	if bullet.position != player.position {
		t.Error(fmt.Sprintf("Expected bullet position %s, got %s", player.position, bullet.position))
		return
	}
}

func contains(gameObjects []IGameObject, object IGameObject) bool {
	for i := 0; i < len(gameObjects); i++ {
		if object == gameObjects[i] {
			return true
		}
	}
	return false
}

func TestBulletKillPlayer(t *testing.T) {
	g := NewGame()
	player := NewPlayer(&Point{0, 0}, g)
	g.AddGameObject(player)
	bullet := NewBullet(&Point{0, 0}, NewVector(1, 0), player)
	g.AddGameObject(bullet)
	g.MakeTurn()
	if !contains(g.GameObjects, player) {
		t.Error("Bullet kills your own player")
		return
	}
	bullet2 := NewBullet(&Point{0, 0}, NewVector(1, 0), nil)
	g.AddGameObject(bullet2)
	g.MakeTurn()
	if contains(g.GameObjects, player) {
		t.Error("Bullet should kill players")
	}
}
