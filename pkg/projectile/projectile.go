package projectile

import (
	"fmt"
	"github.com/zhou-en/ray-tracing-by-go/pkg/canvas"
	"github.com/zhou-en/ray-tracing-by-go/pkg/color"
	"github.com/zhou-en/ray-tracing-by-go/pkg/ppm"
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
	withGrav := proj.Velocity.Add(env.Gravity)
	velocity := withGrav.Add(env.Wind)

	return Projectile{position, velocity}
}

func SimulateProjectile() {
	// Initial velocity: Horizontal and vertical components
	v := tuple.NewVector(6.5, 11.0, 0)

	// Initial position
	p := Projectile{
		Position: tuple.NewPoint(0, 1, 0),
		Velocity: v,
	}

	// Environment with gravity and wind
	e := Environment{
		Gravity: tuple.NewVector(0, -0.15, 0), // Steeper gravity for faster descent
		Wind:    tuple.NewVector(-0.03, 0, 0), // Slight leftward force
	}

	// Canvas for trajectory
	c := canvas.New(900, 550)

	// Simulate and plot the trajectory
	for {
		p = Tick(e, p)

		// Add pixel to canvas
		if int(p.Position.X) >= 0 && int(p.Position.X) < c.Width &&
			c.Height-int(p.Position.Y) >= 0 && c.Height-int(p.Position.Y) < c.Height {
			c.AddPixel(int(p.Position.X), c.Height-int(p.Position.Y), color.CreateRGBByName("red"))
			//c.AddPixel(int(p.Position.X), int(p.Position.Y), color.CreateRGBByName("red"))
		}

		// Stop simulation when projectile hits the ground or exits canvas
		if p.Position.Y <= 0.0 || p.Position.X >= float64(c.Width) {
			break
		}
	}

	// Write to PPM file
	ppmData := ppm.New(c)
	if err := ppmData.WritePPMToFile(ppmData, "output.ppm"); err != nil {
		fmt.Printf("Error writing PPM file: %v\n", err)
	}
}
