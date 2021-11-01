package util

import (
	"reflect"
)

//struct to map
func Struct2MapJson(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get("json") == "-" {
			continue
		}
		switch v.Field(i).Kind() {
		case reflect.String:
			if v.Field(i).String() != "" {
				data[string(t.Field(i).Tag.Get("json"))] = v.Field(i).String()
			}
		case reflect.Int, reflect.Int64:
			if v.Field(i).Int() > 0 {
				data[string(t.Field(i).Tag.Get("json"))] = v.Field(i).Int()
			}
		case reflect.Int32:
			if v.Field(i).Int() > 0 {
				data[string(t.Field(i).Tag.Get("json"))] = int32(v.Field(i).Int())
			}
		}
	}
	return data
}

type Paged interface {
	GetPagination() interface{}
}

func Paged2MapJson(obj Paged) map[string]interface{} {
	pagination := obj.GetPagination()
	t := reflect.TypeOf(obj)
	pt := reflect.TypeOf(pagination)
	v := reflect.ValueOf(obj)
	pv := reflect.ValueOf(pagination)

	var data = make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		if t.Field(i).Tag.Get("json") == "-" {
			continue
		}
		switch v.Field(i).Kind() {
		case reflect.String:
			if v.Field(i).String() != "" {
				data[string(t.Field(i).Tag.Get("json"))] = v.Field(i).String()
			}
		case reflect.Int, reflect.Int64:
			if v.Field(i).Int() > 0 {
				data[string(t.Field(i).Tag.Get("json"))] = v.Field(i).Int()
			}
		case reflect.Int32:
			if v.Field(i).Int() > 0 {
				data[string(t.Field(i).Tag.Get("json"))] = int32(v.Field(i).Int())
			}
		}
	}
	for i := 0; i < pt.NumField(); i++ {
		if pt.Field(i).Tag.Get("json") == "-" {
			continue
		}
		switch pv.Field(i).Kind() {
		case reflect.String:
			if pv.Field(i).String() != "" {
				data[string(pt.Field(i).Tag.Get("json"))] = pv.Field(i).String()
			}
		case reflect.Int, reflect.Int64:
			if pv.Field(i).Int() > 0 {
				data[string(pt.Field(i).Tag.Get("json"))] = pv.Field(i).Int()
			}
		case reflect.Int32:
			if pv.Field(i).Int() > 0 {
				data[string(pt.Field(i).Tag.Get("json"))] = int32(pv.Field(i).Int())
			}
		}
	}
	return data
}
