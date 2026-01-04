# Go Struct Flags

**Go Struct Flags** is a lightweight, low-level library for mapping command-line flags to Go structs using struct tags.

## Installation

```bash
go get github.com/mirkobrombin/go-struct-flags
```

## Features

- **Simple API**: Just one call to `Bind()` and you are done.
- **Auto-binding**: Automatically maps fields to flags using the standard `flag` package.
- **Short & Long flags**: Supports both short (e.g., `-v`) and long (e.g., `--verbose`) flag names.
- **Type-safe**: Supports `bool`, `string`, `int`, `uint`, `float64` and more.

## Usage

Define your options struct with `flag` tags:

```go
type Options struct {
    Verbose bool   `flag:"short:v, long:verbose, name:Enable Verbose Output"`
    Output  string `flag:"short:o, long:output, default:json"`
    Port    int    `flag:"short:p, long:port, default:8080"`
}

func main() {
    opts := &Options{}
    
    // Bind flags to the struct
    structflags.Bind(opts) 
    
    // Parse standard flags
    flag.Parse()
    
    fmt.Printf("Verbose: %v, Port: %d\n", opts.Verbose, opts.Port)
}
```

## Documentation

- [Struct Tags](docs/struct-tags.md)
- [Binder Package](docs/binder.md)

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
