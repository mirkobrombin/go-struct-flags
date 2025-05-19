// File: handlers.go
package flagbinder

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// wrapBool adapts fn to HandlerFunc.
func wrapBool(fn func(bool) error) HandlerFunc {
	return func(args []string) error {
		v, err := strconv.ParseBool(args[0])
		if err != nil {
			return err
		}
		return fn(v)
	}
}

// defaultBool returns a HandlerFunc that sets a reflect.Bool field.
func defaultBool(field reflect.Value) HandlerFunc {
	return func(args []string) error {
		v, err := strconv.ParseBool(args[0])
		if err != nil {
			return err
		}
		field.SetBool(v)
		return nil
	}
}

// wrapInt adapts fn to HandlerFunc.
func wrapInt(fn func(int64) error) HandlerFunc {
	return func(args []string) error {
		v, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return err
		}
		return fn(v)
	}
}

// defaultInt returns a HandlerFunc that sets a reflect.Int field.
func defaultInt(field reflect.Value) HandlerFunc {
	return func(args []string) error {
		v, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			return err
		}
		field.SetInt(v)
		return nil
	}
}

// wrapDuration adapts fn to HandlerFunc.
func wrapDuration(fn func(time.Duration) error) HandlerFunc {
	return func(args []string) error {
		d, err := time.ParseDuration(args[0])
		if err != nil {
			return err
		}
		return fn(d)
	}
}

// defaultDuration returns a HandlerFunc that sets a duration-stored-as-int64 field.
func defaultDuration(field reflect.Value) HandlerFunc {
	return func(args []string) error {
		d, err := time.ParseDuration(args[0])
		if err != nil {
			return err
		}
		field.SetInt(int64(d))
		return nil
	}
}

// wrapEnum enforces allowed choices then calls fn.
func wrapEnum(choices []string, fn func(string) error) HandlerFunc {
	return func(args []string) error {
		v := args[0]
		for _, c := range choices {
			if v == c {
				return fn(v)
			}
		}
		return fmt.Errorf("invalid value %q, must be one of %v", v, choices)
	}
}

// defaultEnum returns a HandlerFunc that sets a string field with validation.
func defaultEnum(field reflect.Value, choices []string) HandlerFunc {
	return func(args []string) error {
		v := args[0]
		for _, c := range choices {
			if v == c {
				field.SetString(v)
				return nil
			}
		}
		return fmt.Errorf("invalid value %q, must be one of %v", v, choices)
	}
}

// defaultStrings returns a HandlerFunc that sets a []string field.
func defaultStrings(field reflect.Value) HandlerFunc {
	return func(args []string) error {
		var parts []string
		if len(args) == 1 {
			parts = strings.Split(args[0], ":")
		} else {
			parts = args
		}
		slice := reflect.MakeSlice(field.Type(), len(parts), len(parts))
		for i, s := range parts {
			slice.Index(i).SetString(s)
		}
		field.Set(slice)
		return nil
	}
}
