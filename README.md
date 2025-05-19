# Go Struct Flags

A lightweight and flexible library for binding command‑line arguments to Go 
structs. Go Struct Flags provides automatic discovery of struct tags, default 
handlers for common types, custom overrides, before/after hooks, and optional 
JSON backup.

## Features

* **Auto‑Discovery**: Read struct fields tagged with `flag:"name,type[,choices]"` and bind accordingly.
* **Default Handlers**: Built‑in support for:

  * `bool` → `strconv.ParseBool`
  * `int` → `strconv.ParseInt`
  * `duration` → `time.ParseDuration`
  * `enum` → string validation against allowed choices
  * `strings` → split colon‑separated values into `[]string`
* **Custom Handlers**: Override any binding with `AddBool`, `AddInt`, `AddDuration`, `AddEnum`, or `AddStrings`.
* **Hooks**: Register `BeforeHook` and `AfterHook` callbacks around each binding operation.
* **Batch Binding**: Apply multiple key→args pairs at once with `RunAll`.
* **Optional Backup**: Write a JSON snapshot of your struct on `NewBinder(..., autobackup=true)`.

## Getting Started

### Installation

```bash
go get github.com/mirkobrombin/go-struct-flags
```

### Basic Usage

```go
package main

import (
  "fmt"
  "time"

  "github.com/mirkobrombin/go-struct-flags/v1/binder"
)

type Config struct {
  Verbose bool          `flag:"verbose,bool"`
  Timeout time.Duration `flag:"timeout,duration"`
  Mode    string        `flag:"mode,enum,dev|prod"`
  Tags    []string      `flag:"tags,strings"`
}

func main() {
  var cfg Config
  binder, err := binder.NewBinder(&cfg, "./_backups", true)
  if err != nil {
    panic(err)
  }

  // Bind individual flags
  binder.Run("verbose", []string{"true"})
  binder.Run("timeout", []string{"30s"})
  binder.Run("mode",    []string{"prod"})
  binder.Run("tags",    []string{"alpha:beta:gamma"})

  // Or in batch
  settings := map[string][]string{
    "verbose": {"true"},
    "timeout": {"45s"},
    "mode":    {"dev"},
    "tags":    {"x:y:z"},
  }
  binder.RunAll(settings)

  fmt.Printf("Loaded config: %+v\n", cfg)
}
```

## Documentation

* [Getting Started](docs/getting-started.md)
* [Struct Tags](docs/struct-tags.md)
* [Default Handlers](docs/handlers.md)
* [Custom Hooks](docs/hooks.md)

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) 
file for details.
