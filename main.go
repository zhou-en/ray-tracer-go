package main

import (
	"fmt"
	"github.com/zhou-en/ray-tracing-by-go/pkg/projectile"
	"github.com/zhou-en/ray-tracing-by-go/pkg/tuple"
)

func main() {
	fmt.Println("Hello, world!")
	fmt.Println(tuple.NewVector(1, 2, 3).X)
	projectile.SimulateProjectile()
}
