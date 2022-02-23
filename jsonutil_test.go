package goutil

import (
	"encoding/json"
	"reflect"
	"testing"
	//"reflect"
)

func TestSetValue(t *testing.T) {
	// s := `
	// {
	// 	"name": "gopher",
	// 	"age": {
	// 		"name": "gopher",
	// 		"age": {
	// 			"name": "gopher",
	// 			"age": 7
	// 		}
	// 	},
	// 	"array": [1, 2, 3, 4, 5]
	// }
	// `

	ss := `{"name":"test"}`

	var src, dst interface{}
	var sd string
	//json.Unmarshal([]byte(s), &u)
	json.Unmarshal([]byte(ss), &src)
	//fmt.Printf("%+v\n", src)

	JsonSetValue(src, "data", "add data")
	sd = `{"name":"test","data":"add data"}`
	json.Unmarshal([]byte(sd), &dst)
	if !reflect.DeepEqual(src, dst) {
		t.Error(`fail to add key-value - JsonSetValue(src, "data", "add data")`)
	}

	JsonSetValue(src, "age.age.name", "change data")
	sd = `{"age":{"age":{"name":"change data"}},"data":"add data","name":"test"}`
	json.Unmarshal([]byte(sd), &dst)
	if !reflect.DeepEqual(src, dst) {
		t.Error(`fail to add multi depth data - JsonSetValue(src, "age.age.name", "change data")`)
	}

	JsonSetValue(src, "age.age", JsonNewObject())
	sd = `{"age":{"age":{}},"data":"add data","name":"test"}`
	json.Unmarshal([]byte(sd), &dst)
	if !reflect.DeepEqual(src, dst) {
		t.Error(`fail to replace object - JsonSetValue(src, "age.age", JsonNewObject())`)
	}

	JsonSetValue(src, "age.age.name", "change data")
	sd = `{"age":{"age":{"name":"change data"}},"data":"add data","name":"test"}`
	json.Unmarshal([]byte(sd), &dst)
	if !reflect.DeepEqual(src, dst) {
		t.Error(`fail to re-add key-value - JsonSetValue(src, "age.age.name", "change data")`)
	}

	// result, _ := json.Marshal(src)
	// fmt.Printf("%+v\n", string(result))
}

func TestGetValue(t *testing.T) {
	ss := `{"age":{"age":{"age":7,"name":"gopher3"},"array":[{"name":"test1"},{"name":"test2"},{"name":"test3"}],"name":"gopher2"},"array":[1,2,3,4,5],"name":"gopher1"}`

	var sd string
	var src, dst, res interface{}
	json.Unmarshal([]byte(ss), &src)

	res = JsonGetValue(src, "age")
	sd = `{"age":{"age":7,"name":"gopher3"},"array":[{"name":"test1"},{"name":"test2"},{"name":"test3"}],"name":"gopher2"}`
	json.Unmarshal([]byte(sd), &dst)
	if !reflect.DeepEqual(res, dst) {
		t.Error(`fail to get object - JsonGetValue(src, "age")`)
	}
	// temp, _ := json.Marshal(res)
	// fmt.Printf("%+v\n", string(temp))

	// 키 호출 테스트
	res = JsonGetValue(src, "name")
	if res != "gopher1" {
		t.Error(`fail to 키 호출 테스트 - JsonGetValue(src, "name")`)
	}

	// 조합키 호출테스트
	res = JsonGetValue(src, "age.name")
	if res != "gopher2" {
		t.Error(`fail to 조합키 호출테스트1 - JsonGetValue(src,"age.name")`)
	}
	res = JsonGetValue(src, "age.age.age")
	if res != 7.0 {
		t.Error(`fail to 조합키 호출테스트2 - JsonGetValue(src,"age.age.age")`)
	}

	// 배열 인덱스 호출테스트
	res = JsonGetValue(src, "array[1]")
	if res != 2.0 {
		t.Error(`fail to 배열 인덱스 호출테스트1 - JsonGetValue(src,"array[1]")`)
	}
	res = JsonGetValue(src, "age.array[1]")
	sd = `{"name":"test2"}`
	json.Unmarshal([]byte(sd), &dst)
	if !reflect.DeepEqual(res, dst) {
		t.Error(`fail to 배열 인덱스 호출테스트2 - JsonGetValue(src, "age.array[1]")`)
	}
	res = JsonGetValue(src, "age.array[0].name")
	if res != "test1" {
		t.Error(`fail to 배열 인덱스 호출테스트3 - JsonGetValue(src, "age.array[0].name")`)
	}
	res = JsonGetValue(src, "age.array[2].name")
	if res != "test3" {
		t.Error(`fail to 배열 인덱스 호출테스트3 - JsonGetValue(src, "age.array[2].name")`)
	}

	// 존재하지 않는 키 호출 테스트
	res = JsonGetValue(src, "age.age.age.name")
	if res != nil {
		t.Error(`fail to 존재하지 않는 키 호출 테스트1 - JsonGetValue(src,"age.age.age.name")`)
	}
	res = JsonGetValue(src, "age.array[3].name")
	if res != nil {
		t.Error(`fail to 존재하지 않는 키 호출 테스트2 - JsonGetValue(src,"age.array[3].name")`)
	}
	res = JsonGetValue(src, "array[5]")
	if res != nil {
		t.Error(`fail to 존재하지 않는 키 호출 테스트3 - JsonGetValue(src,"array[5]")`)
	}

	// 데이터가 없을경우 기본값을 반환하는 펑션 테스트
	res = JsonGetValueDefault(src, "array[5]", "test")
	if res != "test" {
		t.Error(`fail to 데이터가 없을경우 기본값을 반환하는 펑션 테스트 - JsonGetValueDefault(src, "array[5]", "test")`)
	}
}

func TestCheckNulls(t *testing.T) {
	ss := `{"age":{"age":{"age":7,"name":"gopher3"},"array":[{"name":"test1"},{"name":"test2"},{"name":"test3"}],"name":"gopher2"},"array":[1,2,3,4,5],"name":"gopher1"}`

	var err error
	var src interface{}
	json.Unmarshal([]byte(ss), &src)

	// Json객체에서 모든 키가 존재하는지 확인하는 함수 테스트
	err = CheckNulls(src, []string{"name", "age", "age.name", "array[3]"})
	if err != nil {
		t.Error(`fail to Json객체에서 모든 키가 존재하는지 확인하는 함수 테스트1 - CheckNulls(src, []string{"name", "age", "age.name", "array[3]"})`)
	}
	err = CheckNulls(src, []string{"name.age", "age", "age.name", "array[7]"})
	if err == nil {
		t.Error(`fail to Json객체에서 모든 키가 존재하는지 확인하는 함수 테스트2 - CheckNulls(src, []string{"name.age", "age", "age.name", "array[7]"})`)
	}
}


func TestDeepCopy(t *testing.T) {
	ss := `{"age":{"age":{"age":7,"name":"gopher3"},"array":[{"name":"test1"},{"name":"test2"},{"name":"test3"}],"name":"gopher2"},"array":[1,2,3,4,5],"name":"gopher1"}`

	var err error
	var src, dst, res interface{}
	json.Unmarshal([]byte(ss), &src)
	json.Unmarshal([]byte(ss), &dst)

	err = JsonDeepCopy(&res, src)
	if err != nil {
		t.Error(`fail to Json 객체 DeepCopy 테스트1 - JsonDeepCopy(&dst, src) - ` + err.Error())
	} else {
		if !reflect.DeepEqual(res, dst) {
			t.Error(`fail to Json 객체 DeepCopy 테스트1 - JsonDeepCopy(&dst, src) - ` + "객체가 서로 다름")
		}
	}
}

