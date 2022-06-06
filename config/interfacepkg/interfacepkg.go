package interfacepkg

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func InArray(val, array interface{}) (exists bool, index int) {
	exists = false
	index = -1

	switch reflect.TypeOf(array).Kind() {
	case reflect.Slice:
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) {
				index = i
				exists = true
				return
			}
		}
	}

	return
}

func Exist(val, array interface{}) (exists bool) {
	exists = false

	if reflect.TypeOf(array).Kind() == reflect.Slice {
		s := reflect.ValueOf(array)

		for i := 0; i < s.Len(); i++ {
			if reflect.DeepEqual(val, s.Index(i).Interface()) {
				exists = true
				return
			}
		}
	}

	return
}

func InterfaceArrayToString(data []interface{}) (res string) {
	for i, v := range data {
		if i != 0 {
			res = res + ","
		}
		res = res + fmt.Sprintf("%v", v)
	}

	return res
}

func Marshall(data interface{}) (res string) {
	name, err := json.Marshal(data)
	if err != nil {
		return res
	}
	res = string(name)

	return res
}

func Unmarshall(data string) (res interface{}) {
	json.Unmarshal([]byte(data), &res)

	return res
}

func UnmarshallCb(data string, res interface{}) {
	json.Unmarshal([]byte(data), &res)
}

func UnmarshallCbInterface(data, res interface{}) {
	dataString := Marshall(data)
	json.Unmarshal([]byte(dataString), &res)
}

func MarshallMap(data map[string]interface{}) (res string) {
	name, err := json.Marshal(data)
	if err != nil {
		return res
	}
	res = string(name)

	return res
}

func UnmarshallMap(data string) (res map[string]interface{}) {
	json.Unmarshal([]byte(data), &res)

	return res
}

func Convert(data, cb interface{}) (err error) {
	dataString := Marshall(data)
	err = json.Unmarshal([]byte(dataString), &cb)

	return err
}
