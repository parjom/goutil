package goutil

import (
	"strings"
	"encoding/json"
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
				// 없는 해시키맵을 만들어서 넣어야 한다.
				if (localObj.(map[string]interface{})[keys[i]] == nil) {
					localObj.(map[string]interface{})[keys[i]] = make(map[string]interface{})
				}
				localObj = localObj.(map[string]interface{})[keys[i]]
			}
		default:
			localObj = nil
		}
	}
	return false
}

func JsonNewObject() interface{} {
	return make(map[string]interface{})
}

func JsonEncoding(jsonString string) (interface{}, error) {
	var u interface{}
	err := json.Unmarshal([]byte(jsonString), &u)
	if (err != nil) {
		return nil, err
	}
	return u, nil
}