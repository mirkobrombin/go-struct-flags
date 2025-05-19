# Struct Tags

Every field you want to bind needs a `flag:"<name>,<type>[,<choices>]"` tag.

- `<name>`: the key you’ll pass to `Run()`.  
- `<type>`: one of `bool`, `int`, `duration`, `enum`, `strings`.  
- `<choices>` (only for `enum`): pipe-separated allowed values.

Example:

```go
type Opts struct {
  Mode   string   `flag:"mode,enum,dev|prod"`
  Extras []string `flag:"extras,strings"`
}
```
The above struct will accept:

```go
binder.Run("mode", []string{"dev"})
binder.Run("extras", []string{"foo", "bar"})
```

The `Mode` field will only accept `dev` or `prod`, while `Extras` can accept any number of strings.
The `Run()` method will return an error if the value is not valid for the field type or if the value is not in the allowed choices for `enum` types.