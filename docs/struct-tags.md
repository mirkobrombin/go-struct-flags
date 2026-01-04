# Struct Tags

Go Struct Flags uses the `flag` tag to map struct fields to command-line flags.

## Tag Format

The tag values are separated by commas and use the `key:value` format.

```go
type Options struct {
    Verbose bool `flag:"short:v, long:verbose, name:Enable Verbose Output, default:false"`
}
```

## Available Keys

| Key | Description | Example |
|-----|-------------|---------|
| `short` | The short name for the flag (e.g., `-v`). | `short:v` |
| `long` | The long name for the flag (e.g., `--verbose`). Defaults to the field name in lowercase. | `long:verbose` |
| `name` | The description/usage text for the flag. | `name:Enable verbose output` |
| `default` | The default value for the flag. | `default:json` |

## Supported Types

The library supports the following Go types:
- `bool`
- `string`
- `int`
- `int64`
- `uint`
- `uint64`
- `float64`
