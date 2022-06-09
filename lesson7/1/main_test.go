package main

import (
	"fmt"
	"reflect"
	"testing"
)

type Contacts struct {
	Phone      string
	Email      string
	OtherEmail map[string]string
}

type Person struct {
	Name     string
	Age      int64
	Contacts Contacts
}

func buildData() map[string]interface{} {
	data := make(map[string]interface{})
	data["Name"] = "Danil"
	data["Age"] = int64(35)
	contacts := make(map[string]interface{})
	contacts["Phone"] = "71234567890"
	contacts["Email"] = "danil@danil.loc"
	data["Contacts"] = contacts

	return data
}

func TestFillStruct(t *testing.T) {
	data := buildData()
	valid := &Person{
		Name: "Danil",
		Age:  35,
		Contacts: Contacts{
			Phone: "71234567890",
			Email: "danil@danil.loc",
		},
	}
	result := &Person{}
	err := FillStruct(result, data)
	if err != nil {
		t.Errorf("results not match\nGot:\n%+v\nExpected:\n%+v", err, valid)
	}
	if !reflect.DeepEqual(result, valid) {
		t.Errorf("results not match\nGot:\n%+v\nExpected:\n%+v", result, valid)
	}
}

func TestFieldNotExists(t *testing.T) {
	data := buildData()
	result := &Person{}

	data["SomeField"] = "SomeField"

	validError := fmt.Errorf("no such field %s in obj", "SomeField")
	err := FillStruct(result, data)
	if err == nil {
		t.Errorf("results not match\nGot:\n%+v\nExpected:\n%+v", err, validError)
	}
	if !reflect.DeepEqual(err, validError) {
		t.Errorf("results not match\nGot:\n%+v\nExpected:\n%+v", err, validError)
	}
}
func TestInvalidType(t *testing.T) {
	data := buildData()
	result := &Person{}

	validError := fmt.Errorf("provided value type didn't match obj field type")
	data["Age"] = float64(35.44)
	err := FillStruct(result, data)
	if err == nil {
		t.Errorf("results not match\nGot:\n%+v\nExpected:\n%+v", err, validError)
	}
	if !reflect.DeepEqual(err, validError) {
		t.Errorf("results not match\nGot:\n%+v\nExpected:\n%+v", err, validError)
	}
}
