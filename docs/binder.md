# Binder Package

The `binder` package is the core engine of **Go Struct Flags**. While `structflags.Bind()` provides an easy way to map standard library flags to structs, the `binder` package can be used directly for manual control, custom handlers, and lifecycle hooks.

## Installation

```bash
import "github.com/mirkobrombin/go-struct-flags/v2/pkg/binder"
```

## Basic Usage

The binder reflects on a struct to discover fields tagged with `flag` and registers default handlers for them.

```go
type Config struct {
    Port int `flag:"port"`
}

cfg := &Config{}
b, _ := binder.NewBinder(cfg)

// Manually run a binding
b.Run("port", []string{"8080"})
fmt.Println(cfg.Port) // 8080
```

## Custom Handlers

You can override or add custom logic for specific keys:

```go
b.AddInt("port", func(val int64) error {
    if val < 1024 {
        return fmt.Errorf("privileged ports not allowed")
    }
    cfg.Port = int(val)
    return nil
})

// Enums
b.AddEnum("mode", []string{"dev", "prod"}, func(val string) error {
    cfg.Mode = val
    return nil
})

// Durations
b.AddDuration("timeout", func(val time.Duration) error {
    cfg.Timeout = val
    return nil
})
```

## Lifecycle Hooks

Hooks allow you to execute logic before or after a value is bound to a field.

```go
b.BeforeHook("port", func(key string, args []string) {
    fmt.Printf("Attempting to bind %s with values %v\n", key, args)
})

b.AfterHook("port", func(key string, args []string) {
    fmt.Printf("Successfully bound %s\n", key)
})
```

## Manual Dispatching

The binder is transport-agnostic. You can use it to bind data from any source (CLI arguments, API requests, configuration files):

```go
// Example: binding from a custom source
data := map[string][]string{
    "port": {"9090"},
}

for k, v := range data {
    if err := b.Run(k, v); err != nil {
        log.Printf("Failed to bind %s: %v", k, err)
    }
}
```
