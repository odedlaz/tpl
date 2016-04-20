package enumerable

// MapKeys returns the Keys of a map
func MapKeys(m map[string]interface{}) []string {
	// maybe in the future reflect.ValueOf(abc).MapKeys()
	keys := []string{}
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
