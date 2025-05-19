# Getting Started

1. Import the lib and your struct:

```go
import "github.com/mirkobrombin/go-struct-flags/v1/binder"

type Config struct {
    Verbose bool `flag:"verbose,bool"`
    Timeout int  `flag:"timeout,int"`
}
   ```
2. Create & bind:

```go
var cfg Config
binder, _ := binder.NewBinder(&cfg, "./backups", true)
binder.Run("verbose", []string{"true"})
binder.Run("timeout", []string{"30"})
```
3. Use the struct with the new values:

```go
fmt.Println(cfg.Verbose) // true
fmt.Println(cfg.Timeout) // 30
```