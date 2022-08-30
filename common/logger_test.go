package common

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type structA struct {
	name  *string
	name1 string
	age   byte
	b     *structB
}

type structB struct {
	bname *string
}

func TestP(t *testing.T) {
	name := "tom"
	data := []struct {
		data   interface{}
		expect string
	}{
		{"a", "a"},
		{1, "1"},
		{int32(1), "1"},
		{uint32(1), "1"},
		{float32(1), "1"},
		{[]string{"a", "b"}, "[a, b]"},
		{[]int{1, 2}, "[1, 2]"},
		{[]interface{}{"a", 1}, "[a, 1]"},
		{map[string]int{"a": 1, "b": 2}, "{a: 1, b: 2}"},
		{
			structA{&name, "tom1", 123, &structB{&name}},
			"common.structA({name: tom, name1: tom1, age: 123, b: common.structB({bname: tom})})",
		},
		{
			struct {
				name  *string
				name1 string
				age   byte
				inter interface{}
			}{&name, "tom1", 123, float32(2)},
			"struct { name *string; name1 string; age uint8; inter interface {} }({name: tom, name1: tom1, age: 123, inter: 2})",
		},
	}
	for i, value := range data {
		fmt.Printf("test:%d\n", i)
		assert.Equal(t, P(value.data), value.expect)
	}
}
