package structflags

import (
	"flag"
	"fmt"
	"reflect"
	"strings"

	"github.com/mirkobrombin/go-struct-flags/v2/pkg/binder"
)

// Bind maps command-line flags to the fields of a pointer to a struct.
func Bind(target any) {
	b, err := binder.NewBinder(target)
	if err != nil {
		panic(fmt.Sprintf("structflags.Bind: %v", err))
	}

	v := reflect.ValueOf(target).Elem()
	typ := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldType := typ.Field(i)

		tag := fieldType.Tag.Get("flag")
		if tag == "" {
			continue
		}

		params := binder.ParseTag(tag)
		short := params["short"]
		long := params["long"]
		name := params["name"]
		def := params["default"]

		if long == "" {
			long = strings.ToLower(fieldType.Name)
		}

		if name == "" {
			name = fieldType.Name
		}

		isBool := field.Kind() == reflect.Bool

		// Register the long flag
		registerFlagWithBinder(b, long, long, name, isBool)

		// Register the short flag if provided
		if short != "" {
			if h, ok := b.Handlers()[long]; ok {
				b.Handlers()[short] = h
			}
			registerFlagWithBinder(b, short, long, name, isBool)
		}

		// Apply default value via binder if present
		if def != "" {
			b.Run(long, []string{def})
		}
	}
}

// binderValue implements flag.Value and flag.boolFlag
type binderValue struct {
	b      *binder.Binder
	key    string
	isBool bool
}

func (v *binderValue) Set(s string) error {
	return v.b.Run(v.key, []string{s})
}

func (v *binderValue) String() string {
	return ""
}

func (v *binderValue) IsBoolFlag() bool {
	return v.isBool
}

func registerFlagWithBinder(b *binder.Binder, flagName, binderKey, usage string, isBool bool) {
	val := &binderValue{
		b:      b,
		key:    binderKey,
		isBool: isBool,
	}
	flag.Var(val, flagName, usage)
}
