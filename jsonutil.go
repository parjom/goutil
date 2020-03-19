package goutil

import (
	"strings"
)

func JsonGetValue(obj interface{}, key string) interface{} {
	keys := strings.Split(key, ".")
	var localObj interface{} = obj
	for _, k := range keys {
		switch localObj.(type) {
		case map[string]interface{}:
			localObj = localObj.(map[string]interface{})[k]
		default:
			localObj = nil
		}
	}
	return localObj
}
