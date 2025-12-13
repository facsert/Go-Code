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

type fieldTag struct {
	name string
	ignore bool
	omitZero bool
}

// 解析 Tag
func parseTag(field reflect.StructField, tag TagType) *fieldTag {
	jsonTag := field.Tag.Get(string(tag))
	if jsonTag == "-" {
		return &fieldTag{"", true, false}
	}
	name, opts, _ := strings.Cut(jsonTag, ",")
	if strings.TrimSpace(name) == "" {
		name = field.Name
	}
	return &fieldTag{name, false, strings.Contains(opts, "omitzero")}
}

// 确保传入的是结构体类型
func ensureStruct(m any) (reflect.Type, reflect.Value, bool){
	t, v := reflect.TypeOf(m), reflect.ValueOf(m)
	if t.Kind() == reflect.Pointer {
		if v.IsNil() {
			return nil, reflect.Value{}, false
		}
		t, v = t.Elem(), v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil, reflect.Value{}, false
	}
	return t, v, true
}

func StructKeys(m any, tags ...TagType) []string {
	t, v, isStruct := ensureStruct(m)
	if !isStruct {
		return nil
	}
	
	tag := JsonTag
	if len(tags) > 0 {
		tag = tags[0]
	}

	keys := make([]string, 0, t.NumField())
	for index := 0; index < t.NumField(); index++ {
		ft := parseTag(t.Field(index), tag)
		if ft.ignore {
			continue
		}
		if ft.omitZero && v.Field(index).IsZero() {
			continue
		}
		keys = append(keys, ft.name)
	}
	return keys
}

func StructValue(m any, field string) any {
    _, v, isStruct := ensureStruct(m)
    if !isStruct {
        return nil
    }
	return v.FieldByName(field).Interface()
}

func StructMap(m any, tags ...TagType) map[string]any {
	t, v, isStruct := ensureStruct(m)
    if !isStruct {
        return nil
    }

	tag := JsonTag
	if len(tags) > 0 {
		tag = tags[0]
	}

	structMap := make(map[string]any, t.NumField())
	for index := 0; index < t.NumField(); index++ {
		ft := parseTag(t.Field(index), tag)
		if ft.ignore {
			continue
		}
		if ft.omitZero && v.Field(index).IsZero() {
			continue
		}
		structMap[ft.name] = v.Field(index).Interface()
	}
	return structMap
}
