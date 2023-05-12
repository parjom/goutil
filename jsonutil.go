package goutil

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"
)

func CheckNulls(data interface{}, keys []string) error {
	var nullKeys []string = []string{}
	for _, key := range keys {
		if JsonGetValue(data, key) == nil {
			nullKeys = append(nullKeys, key)
		}
	}
	if len(nullKeys) > 0 {
		return errors.New("miss parameter is [ " + strings.Join(nullKeys, ", ") + " ]")
	}
	return nil
}

func JsonGetValueDefault(obj interface{}, key string, defaultValue interface{}) interface{} {
	value := JsonGetValue(obj, key)
	if value == nil {
		return defaultValue
	}
	return value
}

func JsonGetValue(obj interface{}, key string) interface{} {
	keys := strings.Split(key, ".")
	var localObj interface{} = obj
	for _, k := range keys {
		k2, idx, err := checkArray(k)
		if err != nil {
			return nil
		}
		switch v := localObj.(type) {
		case map[string]interface{}:
			if idx == -1 {
				localObj = v[k]
			} else {
				if v[k2] == nil {
					localObj = nil
				} else {
					switch v[k2].(type) {
					case []interface{}:
						tmpObj := v[k2].([]interface{})
						if len(tmpObj) <= idx {
							// 인덱스 값보다 데이터 수가 적은경우
							localObj = nil
						} else {
							localObj = tmpObj[idx]
						}
					default:
						localObj = nil
					}
				}
			}
		default:
			localObj = nil
		}
	}
	return localObj
}
func checkArray(key string) (string, int, error) {
	if len(key) == 0 {
		return "", 0, errors.New("invaild key")
	}
	if string(key[len(key)-1]) == "]" {
		e := len(key) - 1
		for i := 0; i < e; i++ {
			if string(key[i]) == "[" {
				if i == 0 || i+1 == e { //  key의 패턴이 [...] 이거나 ...[] 인경우, 에러 반환
					return "", 0, errors.New("invaild key")
				} else {
					retKey := key[0:i]
					retIndex, err := strconv.Atoi(key[i+1 : e])
					if err != nil {
						return "", 0, err
					} else {
						return retKey, retIndex, nil
					}
				}
			}
		}
		return "", 0, errors.New("invaild key")
	}
	// 문자열이 ] 으로 끝나지 않는다면, 배열키가 아니라고 생각하고 반환함
	return key, -1, nil
}

func JsonSetValue(obj interface{}, key string, value interface{}) bool {
	keys := strings.Split(key, ".")
	var localObj interface{} = obj
	e := len(keys)
	for i := 0; i < e; i++ {
		switch v := localObj.(type) {
		case map[string]interface{}:
			if i == (e - 1) {
				v[keys[i]] = value
				return true
			} else {
				// 없는 해시키맵을 만들어서 넣어야 한다.
				if v[keys[i]] == nil {
					v[keys[i]] = make(map[string]interface{})
				}
				localObj = v[keys[i]]
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
	if err != nil {
		return nil, err
	}
	return u, nil
}

func JsonDeepCopy(dst interface{}, src interface{}) error {
	if dst == nil {
		return errors.New("dst cannot be nil")
	}
	if src == nil {
		return errors.New("src cannot be nil")
	}
	bytes, err := json.Marshal(src)
	if err != nil {
		return errors.New("Unable to marshal src: " + err.Error())
	}
	err = json.Unmarshal(bytes, dst)
	if err != nil {
		return errors.New("Unable to unmarshal into dst: " + err.Error())
	}
	return nil
}
