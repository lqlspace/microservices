package parser

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type School struct {
	Name string `json:"name"`
	Grade int64 `json:"grade"`
}

type Student struct {
	Name string `json:"name" db:"name"`
	Age int64 `json:"age"`
	Man bool `json:"man"`
	Score float64 `json:"score"`
	School School `json:"school"`
}

type S struct {
	Name string `json:"name" db:"name"`
	Age int8 `json:"age"`
	Man bool `json:"man"`
	Score float32 `json:"score"`
	School School `json:"school"`
}


func TestParseMapToStruct(t *testing.T) {
	s := &Student{
		Name: "allen",
		Age: 20,
		Man: true,
		Score: 11.2,
		School: School{
			Name:  "nuaa",
			Grade: 3,
		},
	}

	bytes, err := json.Marshal(&s)
	assert.Nil(t, err)

	var m map[string]interface{}
	err = json.Unmarshal(bytes, &m)
	assert.Nil(t, err)

	fmt.Println(m["birthday"])

	var stu S

	// 传值、传指针都可以
	parser := NewParser("json")
	parser.ParseEntry(m, &stu)
	assert.Equal(t, "allen", stu.Name)
	//反序列成map后，整型的都会改为float64类型
	fmt.Println(stu.Age)
	assert.True(t, stu.Man)
	fmt.Println(stu.Score)
	fmt.Println(stu.School)


}
