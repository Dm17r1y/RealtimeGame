package game

type Weapon struct {
	GetBullet     func(model *PlayerModel) IGameObject
	reloadTime    int
	reloadProcess int
}

func NewWeapon(getBullet func(model *PlayerModel) IGameObject, reloadTime int) *Weapon {
	return &Weapon{GetBullet: getBullet, reloadTime: reloadTime, reloadProcess: 0}
}

func (weapon *Weapon) CanShoot() bool {
	return weapon.reloadProcess == 0
}

func (weapon *Weapon) Shoot(model *PlayerModel) IGameObject {
	weapon.reloadProcess = weapon.reloadTime
	return weapon.GetBullet(model)
}

func (weapon *Weapon) Reload() {
	if weapon.reloadProcess > 0 {
		weapon.reloadProcess -= 1
	}
}
