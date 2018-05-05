package game

type Command struct {
	Shoot *Shoot
}

type Shoot struct {
	Vector *Vector
}
