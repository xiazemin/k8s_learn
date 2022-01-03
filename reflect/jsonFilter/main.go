package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func filter(in interface{}) interface{} {
	//return reflectFilter(reflect.ValueOf(in)).Interface()
	src, _ := json.Marshal(in)
	data := strings.ReplaceAll(string(src), "$", " ")
	var dst interface{}
	_ = json.Unmarshal([]byte(data), &dst)
	return dst
}

func main() {
	fmt.Println(filter(map[string]interface{}{
		"abc": []int{123, 456},
		"def": "a$a$$$",
	}))
}
