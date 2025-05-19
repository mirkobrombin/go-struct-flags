package binder

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// NewBinder creates a Binder for dst and optionally writes a JSON backup to backupDir.
func NewBinder(dst any, backupDir string, autobackup bool) (*Binder, error) {
	if autobackup {
		if err := os.MkdirAll(backupDir, 0o755); err != nil {
			return nil, err
		}
		backupPath := filepath.Join(backupDir, fmt.Sprintf("backup-%d.json", time.Now().Unix()))
		f, err := os.Create(backupPath)
		if err != nil {
			return nil, err
		}
		defer f.Close()
		if err := json.NewEncoder(f).Encode(dst); err != nil {
			return nil, err
		}
	}
	b := &Binder{
		dst:      dst,
		handlers: make(map[string]HandlerFunc),
		hooks:    newHooks(),
	}
	autoDiscover(b)
	return b, nil
}

// AddBool overrides the default handler for key with a custom bool→error function.
func (b *Binder) AddBool(key string, fn func(bool) error) {
	b.handlers[key] = wrapBool(fn)
}

// AddInt overrides the default handler for key with a custom int→error function.
func (b *Binder) AddInt(key string, fn func(int64) error) {
	b.handlers[key] = wrapInt(fn)
}

// AddDuration overrides the default handler for key with a custom duration→error function.
func (b *Binder) AddDuration(key string, fn func(time.Duration) error) {
	b.handlers[key] = wrapDuration(fn)
}

// AddEnum overrides the default handler for key with a custom enum handler enforcing choices.
func (b *Binder) AddEnum(key string, choices []string, fn func(string) error) {
	b.handlers[key] = wrapEnum(choices, fn)
}

// AddStrings overrides the default handler for key with a custom []string→error function.
func (b *Binder) AddStrings(key string, fn func([]string) error) {
	b.handlers[key] = fn
}

// BeforeHook registers a hook called before setting a field.
func (b *Binder) BeforeHook(key string, fn HookFunc) {
	b.hooks.addBefore(key, fn)
}

// AfterHook registers a hook called after setting a field.
func (b *Binder) AfterHook(key string, fn HookFunc) {
	b.hooks.addAfter(key, fn)
}

// Run applies handler for key with args, invoking before and after hooks.
func (b *Binder) Run(key string, args []string) error {
	b.hooks.runBefore(key, args)
	h, ok := b.handlers[key]
	if !ok {
		return fmt.Errorf("no handler for key %s", key)
	}
	err := h(args)
	b.hooks.runAfter(key, args)
	return err
}

// RunAll applies multiple key→args mappings in sequence.
func (b *Binder) RunAll(batch map[string][]string) error {
	for k, args := range batch {
		if err := b.Run(k, args); err != nil {
			return err
		}
	}
	return nil
}
