struct2env
==========

`struct2env` binds environment variables to struct fields according to their `env` tags.

It’s designed to work with [encoding/json] and [struct2flag] for layered configuration handling.

[encoding/json]: https://pkg.go.dev/encoding/json@go1.25.3
[struct2flag]: https://github.com/hymkor/struct2flag

`example.go`

```example.go
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
```

`go run example.go`

```go run example.go |
EDITOR="C:/Users/hymko/scoop/apps/vim/current/gvim.exe"
```
