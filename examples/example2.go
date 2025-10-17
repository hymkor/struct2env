//go:build example

package main

import (
	"os"
	"fmt"
	"encoding/json"
	"flag"

	"github.com/hymkor/struct2env"
	"github.com/hymkor/struct2flag"
)

type Env struct {
	Editor string `env:"EDITOR" flag:"e,specify text editor" "json:"editor"`
}

func main() {
	var env Env

	if jsonText, err := os.ReadFile("example2.json"); err == nil {
		if err := json.Unmarshal(jsonText, &env); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return
		}
	}
	struct2env.Bind(&env)
	struct2flag.BindDefault(&env)
	flag.Parse()

	fmt.Printf("EDITOR=%#v\n", env.Editor)
}
