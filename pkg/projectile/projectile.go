package projectile

import (
	"fmt"

	"github.com/zhou-en/ray-tracing-by-go/pkg/tuple"
)

type Projectile struct {
	Position tuple.Point
	Velocity tuple.Vector
}

type Environment struct {
	Gravity tuple.Vector
	Wind    tuple.Vector
}

func Tick(env Environment, proj Projectile) Projectile {
	position := proj.Position.AddVector(proj.Velocity)
	withGrav := proj.Velocity.AddVector(env.Gravity)
	velocity := withGrav.AddVector(env.Wind)

	return Projectile{position, velocity}
}

func SimulateProjectile() {

	v := tuple.NewVector(0, .0, 0)
	p := Projectile{
		Position: tuple.NewPoint(0, 0, 0),
		Velocity: v.Norm(),
	}
	e := Environment{
		Gravity: tuple.NewVector(0, -0.1, 0),
		Wind:    tuple.NewVector(-0.01, 0, 0),
	}
	step := 0
	for {
		step += 1
		p = Tick(e, p)
		fmt.Printf("%d: %f\n", step, p.Position.Y)
		if p.Position.Y <= 0.0 {
			break
		}
	}

}
