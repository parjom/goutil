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

func JsonSetValue(obj interface{}, key string, value interface{}) bool {
	keys := strings.Split(key, ".")
	var localObj interface{} = obj
	e:=len(keys)
	for i := 0; i < e; i++ {
		switch localObj.(type) {
		case map[string]interface{}:
			if i == (e-1) {
				localObj.(map[string]interface{})[keys[i]] = value
				return true
			} else {
				localObj = localObj.(map[string]interface{})[keys[i]]
			}
		default:
			localObj = nil
		}
	}
	return false
}