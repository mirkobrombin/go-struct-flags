package binder

import (
	"fmt"
	"reflect"
	"strings"
)

// AutoDiscover reflects on the destination struct to register default handlers.
func (b *Binder) AutoDiscover() error {
	v := reflect.ValueOf(b.dst)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("binder: destination must be a pointer to a struct")
	}

	elem := v.Elem()
	typ := elem.Type()

	for i := 0; i < elem.NumField(); i++ {
		field := elem.Field(i)
		fieldType := typ.Field(i)

		tag := fieldType.Tag.Get("flag")
		if tag == "" {
			continue
		}

		params := ParseTag(tag)
		long := params["long"]
		if long == "" {
			long = strings.ToLower(fieldType.Name)
		}

		// Register default handler if not already registered
		if _, ok := b.handlers[long]; !ok {
			b.RegisterDefaultHandler(field, long)
		}
	}

	return nil
}

func (b *Binder) RegisterDefaultHandler(field reflect.Value, name string) {
	switch field.Kind() {
	case reflect.Bool:
		b.handlers[name] = func(args []string) error {
			if len(args) > 0 {
				var val bool
				fmt.Sscanf(args[0], "%t", &val)
				field.SetBool(val)
			} else {
				field.SetBool(true)
			}
			return nil
		}
	case reflect.String:
		b.handlers[name] = func(args []string) error {
			if len(args) > 0 {
				field.SetString(args[0])
			}
			return nil
		}
	case reflect.Int, reflect.Int64:
		b.handlers[name] = func(args []string) error {
			if len(args) > 0 {
				var val int64
				fmt.Sscanf(args[0], "%d", &val)
				field.SetInt(val)
			}
			return nil
		}
	}
}

func ParseTag(tag string) map[string]string {
	params := make(map[string]string)
	parts := strings.Split(tag, ",")
	for _, part := range parts {
		kv := strings.Split(part, ":")
		if len(kv) == 2 {
			params[strings.TrimSpace(kv[0])] = strings.TrimSpace(kv[1])
		} else {
			// Support flags without keys as long names (e.g. `flag:"verbose"`)
			key := strings.TrimSpace(part)
			if !strings.Contains(tag, ":") {
				params["long"] = key
			}
		}
	}
	return params
}
