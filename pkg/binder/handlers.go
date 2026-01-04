package binder

import (
	"fmt"
	"slices"
	"time"
)

// AddDuration registers a custom handler for a duration field.
func (b *Binder) AddDuration(key string, fn func(time.Duration) error) {
	b.handlers[key] = func(args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("missing value for %s", key)
		}
		val, err := time.ParseDuration(args[0])
		if err != nil {
			return err
		}
		return fn(val)
	}
}

// AddEnum registers a handler that validates against a set of allowed values.
func (b *Binder) AddEnum(key string, choices []string, fn func(string) error) {
	b.handlers[key] = func(args []string) error {
		if len(args) == 0 {
			return fmt.Errorf("missing value for %s", key)
		}
		val := args[0]
		found := slices.Contains(choices, val)
		if !found {
			return fmt.Errorf("invalid value %s for %s, allowed: %v", val, key, choices)
		}
		return fn(val)
	}
}
