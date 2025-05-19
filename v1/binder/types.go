package binder

type tagMeta struct {
	Name    string
	Type    string
	Choices []string
}

// Binder manages mapping of keys to handlers, auto-discovery from struct tags,
// before/after hooks, and optional JSON backup.
type Binder struct {
	dst      any
	handlers map[string]HandlerFunc
	hooks    *hooks
}

// HandlerFunc processes raw args for a field.
type HandlerFunc func(args []string) error

// HookFunc defines a hook signature.
type HookFunc func(key string, args []string)

// hooks stores before/after hook functions.
type hooks struct {
	before map[string][]HookFunc
	after  map[string][]HookFunc
}
