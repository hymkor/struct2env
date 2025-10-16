//go:build example

package main

import (
	"fmt"

	"github.com/hymkor/struct2env"
)

type Env struct {
	Editor string `env:"EDITOR"`
}

func main() {
	var env Env
	struct2env.Bind(&env)

	fmt.Printf("EDITOR=%#v\n", env.Editor)
}
