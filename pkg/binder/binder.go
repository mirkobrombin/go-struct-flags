package binder

import (
	"fmt"
)

// HandlerFunc is a function that handles binding a set of string arguments to a field.
type HandlerFunc func(args []string) error

// Binder handles mapping and execution of field bindings.
type Binder struct {
	dst      any
	handlers map[string]HandlerFunc
	hooks    *hooks
}

// NewBinder creates a new binder for the destination object.
func NewBinder(dst any) (*Binder, error) {
	b := &Binder{
		dst:      dst,
		handlers: make(map[string]HandlerFunc),
		hooks:    newHooks(),
	}

	if err := b.AutoDiscover(); err != nil {
		return nil, err
	}

	return b, nil
}

// Handlers returns the registered handlers.
func (b *Binder) Handlers() map[string]HandlerFunc {
	return b.handlers
}

// AddBool registers a custom handler for a boolean field.
func (b *Binder) AddBool(key string, fn func(bool) error) {
	b.handlers[key] = func(args []string) error {
		var val bool
		if len(args) > 0 {
			fmt.Sscanf(args[0], "%t", &val)
		} else {
			val = true
		}
		return fn(val)
	}
}

// AddInt registers a custom handler for an integer field.
func (b *Binder) AddInt(key string, fn func(int64) error) {
	b.handlers[key] = func(args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("missing value for %s", key)
		}
		var val int64
		if _, err := fmt.Sscanf(args[0], "%d", &val); err != nil {
			return err
		}
		return fn(val)
	}
}

// AddStrings registers a custom handler for a string slice field.
func (b *Binder) AddStrings(key string, fn func([]string) error) {
	b.handlers[key] = fn
}

// Run executes the handler for the given key with the provided arguments.
func (b *Binder) Run(key string, args []string) error {
	b.hooks.runBefore(key, args)

	h, ok := b.handlers[key]
	if !ok {
		return fmt.Errorf("no handler for key: %s", key)
	}

	if err := h(args); err != nil {
		return err
	}

	b.hooks.runAfter(key, args)
	return nil
}

// BeforeHook registers a function to be called before a specific key is handled.
func (b *Binder) BeforeHook(key string, fn HookFunc) {
	b.hooks.addBefore(key, fn)
}

// AfterHook registers a function to be called after a specific key is handled.
func (b *Binder) AfterHook(key string, fn HookFunc) {
	b.hooks.addAfter(key, fn)
}
