package jsoner

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func mapper(v any, path string, filters *filter) (any, error) {
	value := reflect.Indirect(reflect.ValueOf(v))
	if !value.IsValid() {
		return nil, nil
	}

	switch value.Kind() {
	case reflect.Array, reflect.Slice:
		return mapSlice(value, path, filters)
	case reflect.Map:
		return mapMap(value, path, filters)
	case reflect.Struct:
		return mapStruct(value, path, filters)
	default:
		return v, nil
	}
}

func mapSlice(value reflect.Value, path string, filters *filter) ([]any, error) {
	result := make([]any, value.Len())
	for i := 0; i < value.Len(); i++ {
		res, err := mapper(value.Index(i).Interface(), path, filters)
		if err != nil {
			return nil, err
		}
		result[i] = res
	}
	return result, nil
}

func mapMap(value reflect.Value, path string, filters *filter) (map[string]any, error) {
	result := make(map[string]any)
	for _, k := range value.MapKeys() {
		key := fmt.Sprint(k.Interface())
		fullPath := pathJoiner(path, key)
		if filters.shouldSkip(fullPath) {
			continue
		}

		res, err := mapper(value.MapIndex(k).Interface(), fullPath, filters)
		if err != nil {
			return nil, err
		}
		result[key] = res
	}
	return result, nil
}

func mapStruct(value reflect.Value, path string, filters *filter) (any, error) {
	encoded, err := json.Marshal(value.Interface())
	if err != nil {
		return nil, err
	}

	var decoded map[string]any
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, err
	}

	return mapper(decoded, path, filters)
}

func pathJoiner(root, name string) string {
	if root == "" {
		return name
	}
	return root + "." + name
}
