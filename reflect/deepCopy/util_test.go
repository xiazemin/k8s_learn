package main

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
)

// benchamark测试
func BenchmarkReflect(b *testing.B) {
	p1 := NewPlayer()
	for i := 0; i < b.N; i++ {
		Copy(p1)
	}
}

func TestCopy(t *testing.T) {
	type args struct {
		in interface{}
	}
	type testStruct struct {
		A string
		B []string
		C map[string]interface{}
		D int32
		E *string
		F [4]int16
	}
	str1 := "Abc$,asdsfs"
	str2 := "Abc ,asdsfs"
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
		{
			name: "map",

			args: args{
				in: map[string]interface{}{
					"abc": []int{123, 456},
					"def": "a$a$$$",
				}, //interface{}
			},
			want: map[string]interface{}{
				"abc": []int{123, 456},
				"def": "a a   ",
			}, //interface{},
		},
		{
			name: "struct",
			args: args{
				in: testStruct{
					A: "ab$c",
					B: []string{"abc", "abc$"},
					C: map[string]interface{}{
						"C": 123,
						"d": "abc$",
					},
					D: 5,
					E: &str1,
					F: [4]int16{1, 2, 3, 4},
				},
			},
			want: testStruct{
				A: "ab c",
				B: []string{"abc", "abc "},
				C: map[string]interface{}{
					"C": 123,
					"d": "abc ",
				},
				D: 5,
				E: &str2,
				F: [4]int16{1, 2, 3, 4},
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		t.Run(tt.name, func(t *testing.T) {
			if got := Copy(tt.args.in); !reflect.DeepEqual(got, tt.want) {
				go1, _ := json.Marshal(got)
				w1, _ := json.Marshal(tt.want)
				if string(go1) != string(w1) {
					t.Errorf("filterInvalidString() = %#v, want %#v ,diff: %s --- %s", got, tt.want, string(go1), string(w1))
				}
			}
		})
	}
}
