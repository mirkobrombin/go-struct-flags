package flagbinder

import (
	"reflect"
	"strings"
)

// parseTag splits a tag like "socketX11,bool" or "mode,enum,dev|prod" into tagMeta.
func parseTag(tag string) tagMeta {
	parts := strings.Split(tag, ",")
	meta := tagMeta{Name: parts[0], Type: parts[1]}
	if meta.Type == "enum" && len(parts) > 2 {
		meta.Choices = strings.Split(parts[2], "|")
	}
	return meta
}

// autoDiscover inspects dst struct tags and registers default handlers.
func autoDiscover(b *Binder) {
	v := reflect.ValueOf(b.dst).Elem()
	t := v.Type()
	for i := range t.NumField() {
		f := t.Field(i)
		tag := f.Tag.Get("flag")
		if tag == "" {
			continue
		}
		meta := parseTag(tag)
		field := v.Field(i)
		switch meta.Type {
		case "bool":
			b.handlers[meta.Name] = defaultBool(field)
		case "int":
			b.handlers[meta.Name] = defaultInt(field)
		case "duration":
			b.handlers[meta.Name] = defaultDuration(field)
		case "enum":
			b.handlers[meta.Name] = defaultEnum(field, meta.Choices)
		case "strings":
			b.handlers[meta.Name] = defaultStrings(field)
		}
	}
}
