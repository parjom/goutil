package goutil

import (
	"testing"
	"fmt"
	"encoding/json"
)

func TestSetValue(t *testing.T)  {
	s := `{"name":"gopher","age":{"name":"gopher","age":{"name":"gopher","age":7}}}`

	var u interface{}
	var uu interface{}
	json.Unmarshal([]byte(s), &u)
	json.Unmarshal([]byte(s), &uu)
	fmt.Printf("%+v\n", u) // {Name:gopher Age:7}
	JsonSetValue(u, "age.age.name", uu)
	fmt.Printf("%+v\n", u) // {Name:gopher Age:7}
}