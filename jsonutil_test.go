package goutil

import (
	"testing"
	"fmt"
	"encoding/json"
)

func TestSetValue(t *testing.T)  {
	s := `
	{
		"name": "gopher",
		"age": {
			"name": "gopher",
			"age": {
				"name": "gopher",
				"age": 7
			}
		},
		"array": [1, 2, 3, 4, 5]
	}
	`

	ss := `
	{
		"name":"test"
	}
	`

	var u interface{}
	json.Unmarshal([]byte(s), &u)
	json.Unmarshal([]byte(ss), &u)
	fmt.Printf("%+v\n", u) // {Name:gopher Age:7}
	JsonSetValue(u, "data", "add data")
	fmt.Printf("%+v\n", u) // {Name:gopher Age:7}

	JsonSetValue(u, "age.age.name", "change data")
	fmt.Printf("%+v\n", u) // {Name:gopher Age:7}

	JsonSetValue(u, "age.age", JsonNewObject() )
	fmt.Printf("%+v\n", u) // {Name:gopher Age:7}

	JsonSetValue(u, "age.age.name", "change data")
	fmt.Printf("%+v\n", u) // {Name:gopher Age:7}

	result, _ := json.Marshal(u)
	fmt.Printf("%+v\n", string(result))
}


func TestGetValue(t *testing.T)  {
	s := `
	{
		"name": "gopher",
		"age": {
			"name": "gopher",
			"age": {
				"name": "gopher",
				"age": 7
			},
			"array": [
				{
					"name": "test1"
				},
				{
					"name": "test2"
				},
				{
					"name": "test3"
				}
			]
		},
		"array": [1, 2, 3, 4, 5]
	}
	`

	var u interface{}
	json.Unmarshal([]byte(s), &u)

	fmt.Printf("%+v\n", JsonGetValue(u, "age")) // {Name:gopher Age:7}

	fmt.Printf("%+v\n", JsonGetValue(u, "name")) // {Name:gopher Age:7}

	fmt.Printf("%+v\n", JsonGetValue(u, "age.name")) // {Name:gopher Age:7}

	fmt.Printf("%+v\n", JsonGetValue(u, "age.age.age")) // {Name:gopher Age:7}

	fmt.Printf("%+v\n", JsonGetValue(u, "array[1]")) // {Name:gopher Age:7}

	fmt.Printf("%+v\n", JsonGetValue(u, "age.array[1]")) // {Name:gopher Age:7}
	fmt.Printf("%+v\n", JsonGetValue(u, "age.array[0].name")) // {Name:gopher Age:7}
	fmt.Printf("%+v\n", JsonGetValue(u, "age.array[2].name")) // {Name:gopher Age:7}
}