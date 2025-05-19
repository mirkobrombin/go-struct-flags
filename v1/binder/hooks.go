package binder

// newHooks creates an empty hooks registry.
func newHooks() *hooks {
	return &hooks{
		before: make(map[string][]HookFunc),
		after:  make(map[string][]HookFunc),
	}
}

// addBefore registers a before-hook for a key.
func (h *hooks) addBefore(key string, fn HookFunc) {
	h.before[key] = append(h.before[key], fn)
}

// addAfter registers an after-hook for a key.
func (h *hooks) addAfter(key string, fn HookFunc) {
	h.after[key] = append(h.after[key], fn)
}

// runBefore executes all before-hooks for a key.
func (h *hooks) runBefore(key string, args []string) {
	for _, fn := range h.before[key] {
		fn(key, args)
	}
}

// runAfter executes all after-hooks for a key.
func (h *hooks) runAfter(key string, args []string) {
	for _, fn := range h.after[key] {
		fn(key, args)
	}
}
