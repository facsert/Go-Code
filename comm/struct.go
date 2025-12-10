package comm

import (
	"reflect"
	"strings"
)

type TagType string
var (
	Json TagType = "json"
	Insert TagType = "insert"
	Update TagType = "update"
)

func parseJsonTag(field reflect.StructField, tag TagType) string {
	jsonTag := field.Tag.Get(string(tag))
	if  jsonTag == "-" {
		return ""
	}
	if idx := strings.Index(jsonTag, ","); idx != -1 {
		return jsonTag[:idx]	
	}
	return field.Name
}


func StructKeys(m any, tag TagType) []string {
	if m == nil {
		return nil
	}
	t := reflect.TypeOf(m)
	if t.Kind() == reflect.Pointer {
		t = t.Elem()
	}
	keys := make([]string, 0, t.NumField())
	for index := 0; index < t.NumField(); index++ {
		if key := parseJsonTag(t.Field(index), tag); key != "" {
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

func StructMap(m any, tag TagType) map[string]any {
	if m == nil {
		return nil
	}
	
	t := reflect.TypeOf(m)
	v := reflect.ValueOf(m)

	if v.Kind() == reflect.Pointer {
		t, v = t.Elem(), v.Elem()
	}

	mMap := make(map[string]any, t.NumField())
	for index := 0; index < t.NumField(); index++ {
		if key := parseJsonTag(t.Field(index), tag); key != "" {
			mMap[key] = v.Field(index).Interface()
		}
	}
	return mMap
}