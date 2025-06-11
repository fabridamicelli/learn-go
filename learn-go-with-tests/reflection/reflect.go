package main

import (
	"reflect"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		val = val.Elem()
	}

	return val

}

func walk(x interface{}, fn func(inp string)) {

	val := getValue(x)
	walkVal := func(val reflect.Value) {
		walk(val.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkVal(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkVal(val.Index(i))
		}
	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkVal(val.MapIndex(key))
		}
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walkVal(v)
			} else {
				break
			}
		}

	case reflect.Func:
		out := val.Call(nil)
		for _, v := range out {
			walkVal(v)
		}

	}

}
