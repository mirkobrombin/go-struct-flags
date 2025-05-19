# Handlers

Out-of-the-box, binder supports:

- **bool** → `strconv.ParseBool` → `field.SetBool`  
- **int** → `strconv.ParseInt` → `field.SetInt`  
- **duration** → `time.ParseDuration` → `field.SetInt(int64(d))`  
- **enum** → string match → `field.SetString`  
- **strings** → split by `:` → `field.Set([]string)`  

You don’t need to call `AddXxx` unless you want to override.

## Custom Handlers

> Before diving into custom handlers, check if you can achieve your goal using
> the hook system. It might be more efficient and easier to maintain.

Override a default binding with your own logic:

```go
binder.AddBool("dryrun", func(v bool) error {
  if v {
    fmt.Println("Running in dry-run mode")
  }
  return nil
})
```

Available override methods:

- AddBool(key, fn func(bool) error)
- AddInt(key, fn func(int64) error)
- AddDuration(key, fn func(time.Duration) error)
- AddEnum(key, choices, fn func(string) error)
- AddStrings(key, fn func([]string) error)
