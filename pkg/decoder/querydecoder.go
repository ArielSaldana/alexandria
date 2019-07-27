package decoder

import (
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type QueryDecoder struct {
}

func (q QueryDecoder) Decode(ref interface{}, values url.Values) {
	elem := reflect.ValueOf(ref).Elem()
	//elemType := elem.Type()
	elemNumFields := elem.Type().NumField()
	//fmt.Println(elem , elemType, elemNumFields)

	for i := 0; i < elemNumFields; i++ {
		nameOfField := reflect.ValueOf(ref).Elem().Type().Field(i).Name
		nameOfFieldLowerCase := strings.ToLower(nameOfField)
		value := values.Get(nameOfFieldLowerCase)
		assignDynamicVariable(ref, nameOfField, value)
	}
}

func assignDynamicVariable(ref interface{}, key string, value interface{}) {
	var dynamicVariable reflect.Value

	if reflect.ValueOf(ref).Kind() == reflect.Ptr {
		dynamicVariable = reflect.ValueOf(ref).Elem().FieldByName(key)
	} else {
		dynamicVariable = reflect.ValueOf(ref).FieldByName(key)
	}

	if !dynamicVariable.IsValid() || !dynamicVariable.CanSet() {
		// handle err
		fmt.Println("error 1")
		return
	}
	if dynamicVariable.Type() != reflect.TypeOf(value) {
		// handle err

		// instead of handling the error, we'll manually convert it to the correct format.
		switch dynamicVariable.Type().String() {
		case "uint":
			newVal, err := strconv.ParseUint(reflect.ValueOf(value).String(), 10, 64)

			if err != nil {
				// handle err
				return
			}

			newUintVal := uint(newVal)
			dynamicVariable.Set(reflect.ValueOf(newUintVal))
			break
		}

		return
	}

	dynamicVariable.Set(reflect.ValueOf(value))
}
