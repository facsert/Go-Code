package comm

import (
	"reflect"
	"strings"
)

type TagType string

var (
	JsonTag   TagType = "json"
	InsertTag TagType = "insert"
	UpdateTag TagType = "update"
)

func parseTag(field reflect.StructField, tag TagType) string {
	jsonTag := field.Tag.Get(string(tag))
	if jsonTag == "-" {
		return ""
	}
	if idx := strings.Index(jsonTag, ","); idx != -1 {
		return jsonTag[:idx]
	}

	if jsonTag != "" {
		return jsonTag
	}

	return field.Name
}

func StructKeys(m any, tags ...TagType) []string {
	if m == nil {
		return nil
	}
	t := reflect.TypeOf(m)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}

	tag := JsonTag
	if len(tags) > 0 {
		tag = tags[0]
	}

	keys := make([]string, 0, t.NumField())
	for index := 0; index < t.NumField(); index++ {
		if key := parseTag(t.Field(index), tag); key != "" {
			keys = append(keys, key)
		}
	}
	return keys
}

func StructValue(m any, key string) any {
	if m == nil {
		return nil
	}

	t := reflect.TypeOf(m)
	v := reflect.ValueOf(m)

	if v.Kind() == reflect.Pointer {
		t, v = t.Elem(), v.Elem()
	}
	for index := 0; index < t.NumField(); index++ {
		if key == t.Field(index).Name {
			return v.Field(index).Interface()
		}
	}
	return nil
}

func StructMap(m any, tags ...TagType) map[string]any {
	if m == nil {
		return nil
	}

	t := reflect.TypeOf(m)
	v := reflect.ValueOf(m)

	if v.Kind() == reflect.Pointer {
		t, v = t.Elem(), v.Elem()
	}

	tag := JsonTag
	if len(tags) > 0 {
		tag = tags[0]
	}

	mMap := make(map[string]any, t.NumField())
	for index := 0; index < t.NumField(); index++ {
		if key := parseTag(t.Field(index), tag); key != "" {
			mMap[key] = v.Field(index).Interface()
		}
	}
	return mMap
}
