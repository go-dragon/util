package util

// map functions

// get map keys
func GetMapKeys(data map[string]interface{}) []string {
	keys := make([]string, 0)
	for key := range data {
		keys = append(keys, key)
	}
	return keys
}
