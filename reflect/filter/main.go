package main

import (
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func reflectFilter(rv reflect.Value) reflect.Value {
	typ := rv.Type()
	switch typ.Kind() {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		return rv
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.UnsafePointer:
		return rv
	case reflect.Ptr:
		val := reflectFilter(rv.Elem())
		rv.Set(reflect.New(rv.Type().Elem()))
		rv.Elem().Set(val)
		return rv
	case reflect.Array:
		arr := reflect.Indirect(reflect.New(typ))
		for i := 0; i < rv.Len(); i++ {
			val := reflectFilter(rv.Index(i))
			arr.Index(i).Set(val)
		}
		return arr
	case reflect.Slice:
		rSlice := reflect.MakeSlice(typ, 0, rv.Cap())
		for i := 0; i < rv.Len(); i++ {
			val := reflectFilter(rv.Index(i))
			rSlice = reflect.Append(rSlice, val)
		}
		return rSlice
	case reflect.Map:
		m1 := reflect.MakeMap(typ)
		iter := rv.MapRange()
		for iter.Next() {
			key := iter.Key()
			val := reflectFilter(iter.Value())
			m1.SetMapIndex(key, val)
		}
		return m1
	case reflect.String:
		res := strings.ReplaceAll(rv.String(), "$", " ")
		fmt.Println("xiazemin", res)
		//rv.SetString(res)
		return reflect.ValueOf(res)
	case reflect.Struct:
		st := reflect.New(typ)
		for i := 0; i < rv.NumField(); i++ {
			val := reflectFilter(rv.Field(i))
			st.Elem().FieldByName(typ.Field(i).Name).Set(val)
		}
		return st
	}
	return rv
}

func filterInvalidStringType(in interface{}) interface{} {
	return reflectFilter(reflect.ValueOf(in)).Interface()
}

func filterInvalidString(in interface{}) interface{} {
	typ := reflect.TypeOf(in)
	rv := reflect.ValueOf(in)
	switch typ.Kind() {
	case reflect.Bool, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr, reflect.Float32, reflect.Float64, reflect.Complex64, reflect.Complex128:
		return in
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.UnsafePointer:
		return in
	case reflect.Ptr:
		val := filterInvalidString(rv.Elem().Interface())
		ptr := reflect.NewAt(typ.Elem(), unsafe.Pointer(&val))
		return reflect.ValueOf(ptr).Interface()
	case reflect.Array:
		arr := reflect.Indirect(reflect.New(typ))
		for i := 0; i < rv.Len(); i++ {
			val := filterInvalidString(rv.Index(i).Interface())
			val1, ok := val.(reflect.Value)
			if ok {
				arr.Index(i).Set(val1)
			} else if reflect.TypeOf(val).Kind() == reflect.Ptr {
				arr.Index(i).Set(reflect.ValueOf(val).Elem())
			} else {
				arr.Index(i).Set(reflect.ValueOf(val))
			}
		}
		return arr.Interface()
	case reflect.Slice:
		rSlice := reflect.MakeSlice(typ, 0, rv.Cap())
		for i := 0; i < rv.Len(); i++ {
			val := filterInvalidString(rv.Index(i).Interface())
			val1, ok := val.(reflect.Value)
			if ok {
				rSlice = reflect.Append(rSlice, val1)
			} else {
				rSlice = reflect.Append(rSlice, reflect.ValueOf(val))
			}
		}
		return rSlice.Interface()
	case reflect.Map:
		m1 := reflect.MakeMap(typ)
		iter := rv.MapRange()
		for iter.Next() {
			key := iter.Key()
			val := filterInvalidString(iter.Value().Interface())
			val1, ok := val.(reflect.Value)
			if ok {
				m1.SetMapIndex(key, val1)
			} else {
				m1.SetMapIndex(key, reflect.ValueOf(val))
			}
		}
		return m1.Interface()
	case reflect.String:
		return strings.ReplaceAll(rv.String(), "$", " ")
	case reflect.Struct:
		st := reflect.New(typ)
		for i := 0; i < rv.NumField(); i++ {
			val := filterInvalidString(rv.Field(i).Interface())
			val1, ok := val.(reflect.Value)
			if ok {
				st.Elem().FieldByName(typ.Field(i).Name).Set(val1)
			} else {
				st.Elem().FieldByName(typ.Field(i).Name).Set(reflect.ValueOf(val))
			}
		}
		return st.Interface()
	}
	return in
	// src, _ := json.Marshal(in)
	// data := strings.ReplaceAll(string(src), "$", " ")
	// var dst interface{}
	// _ = json.Unmarshal([]byte(data), &dst)
	// return dst
}

func main() {
	fmt.Println(filterInvalidStringType(map[string]interface{}{
		"abc": []int{123, 456},
		"def": "a$a$$$",
	}))
	fmt.Println(filterInvalidString(map[string]interface{}{
		"abc": []int{123, 456},
		"def": "a$a$$$",
	}))
}
