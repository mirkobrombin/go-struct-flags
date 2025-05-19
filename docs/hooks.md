# Hooks

You can register logic to run **before** or **after** any key:

```go
binder.BeforeHook("timeout", func(key string, args []string) {
  log.Printf("About to set %s=%v", key, args)
})
binder.AfterHook("timeout", func(key string, args []string) {
  log.Printf("Finished setting %s", key)
})
```

Imagine a scenario where you want to trigger a function before or after setting
a value. For example, you might want to invoke a function of your own before
setting a value, or you might want to log the value after it has been set.
