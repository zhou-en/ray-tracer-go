package main

import (
	"fmt"
	"github.com/zhou-en/ray-tracing-by-go/pkg/projectile"
	"github.com/zhou-en/ray-tracing-by-go/pkg/vector"
)

func main() {
	fmt.Println("Hello, world!")
	fmt.Println(vector.New(1, 2, 3).X)
	projectile.SimulateProjectile()
}
