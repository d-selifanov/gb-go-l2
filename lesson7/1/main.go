/*
1. Написать функцию, которая принимает на вход структуру in (struct или кастомную struct) и
values map[string]interface{} (key - название поля структуры, которому нужно присвоить value
этой мапы). Необходимо по значениям из мапы изменить входящую структуру in с помощью
пакета reflect. Функция может возвращать только ошибку error. Написать к данной функции
тесты (чем больше, тем лучше - зачтется в плюс).
*/

package main

import (
	"fmt"
	"reflect"
)

func FillStruct(in interface{}, data map[string]interface{}) error {
	for key, value := range data {
		err := setField(in, key, value)
		if err != nil {
			return err
		}
	}
	return nil
}

func setField(obj interface{}, name string, value interface{}) error {
	if obj == nil {
		return fmt.Errorf("Empty input object")
	}
	rObjVal := reflect.ValueOf(obj)
	if rObjVal.Kind() == reflect.Ptr {
		rObjVal = rObjVal.Elem()
	}
	fldVal := rObjVal.FieldByName(name)
	if !fldVal.IsValid() {
		return fmt.Errorf("no such field %s in obj", name)
	}
	if !fldVal.CanSet() {
		return fmt.Errorf("cannot set %s field value", name)
	}
	rValVal := reflect.ValueOf(value)
	if fldVal.Type() == rValVal.Type() {
		fldVal.Set(rValVal)
		return nil
	}
	if m, ok := value.(map[string]interface{}); ok {
		if fldVal.Kind() == reflect.Struct {
			return FillStruct(fldVal.Addr().Interface(), m)
		}
		if fldVal.Kind() == reflect.Ptr && fldVal.Type().Elem().Kind() == reflect.Struct {
			if fldVal.IsNil() {
				fldVal.Set(reflect.New(fldVal.Type().Elem()))
			}
			return FillStruct(fldVal.Interface(), m)
		}
	}
	return fmt.Errorf("provided value type didn't match obj field type")
}
