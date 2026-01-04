package binder

// HookFunc is a function called before or after a value is bound.
type HookFunc func(key string, args []string)

type hooks struct {
	before map[string][]HookFunc
	after  map[string][]HookFunc
}

func newHooks() *hooks {
	return &hooks{
		before: make(map[string][]HookFunc),
		after:  make(map[string][]HookFunc),
	}
}

func (h *hooks) addBefore(key string, fn HookFunc) {
	h.before[key] = append(h.before[key], fn)
}

func (h *hooks) addAfter(key string, fn HookFunc) {
	h.after[key] = append(h.after[key], fn)
}

func (h *hooks) runBefore(key string, args []string) {
	for _, fn := range h.before[key] {
		fn(key, args)
	}
}

func (h *hooks) runAfter(key string, args []string) {
	for _, fn := range h.after[key] {
		fn(key, args)
	}
}
