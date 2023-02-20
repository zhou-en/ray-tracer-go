package projectile

import (
	"fmt"
	"github.com/zhou-en/ray-tracing-by-go/pkg/point"
	"github.com/zhou-en/ray-tracing-by-go/pkg/tuple"
	"github.com/zhou-en/ray-tracing-by-go/pkg/vector"
)

type Projectile struct {
	Position tuple.Tuple
	Velocity tuple.Tuple
}

type Environment struct {
	Gravity tuple.Tuple
	Wind    tuple.Tuple
}

func Tick(env Environment, proj Projectile) Projectile {
	position := proj.Velocity.Add(proj.Velocity)
	velocity := proj.Velocity.Add(env.Gravity).Add(env.Wind)

	return Projectile{position, velocity}
}

func SimulateProjectile() {
	p := Projectile{
		Position: point.New(0, 1.0, 0.0),
		Velocity: vector.New(1.0, 3.0, 0).Normal(),
	}
	e := Environment{
		Gravity: vector.New(0, -0.1, 0),
		Wind:    vector.New(-0.01, 0, 0),
	}
	step := 0
	for {
		fmt.Println(step)
		step += 1
		fmt.Println(p.Position.Y)
		p = Tick(e, p)
		if p.Position.Y <= 0.0 {
			break
		}
	}

}
