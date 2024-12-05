package generic

import (
	"errors"
	"reflect"
)

type callStruct struct {
	reflectedFunc   reflect.Value
	funcType        reflect.Type
	params          []interface{}
	reflectedParams []reflect.Value
}

const MessageError = "[generic.CallMethod] the number of params is out of index"

func CallMethod(instance interface{}, funcName string, params ...interface{}) (result []interface{}, err error) {
	reflectedFunc := reflect.ValueOf(instance).MethodByName(funcName)
	funcType := reflectedFunc.Type()

	if len(params) != funcType.NumIn() {
		err = errors.New(MessageError)
		return
	}
	
	call := callStruct{
		reflectedFunc: reflectedFunc,
		funcType:      funcType,
		params:        params,
	}

	reflectedParams := call.extractParams()

	call.reflectedParams = reflectedParams
	return call.invoke(), nil
}

func (call callStruct) extractParams() []reflect.Value {
	reflectedParams := make([]reflect.Value, len(call.params))

	for k, param := range call.params {
		if param != nil {
			reflectedParams[k] = reflect.ValueOf(param)
			continue
		}

		expectedType := call.funcType.In(k)
		reflectedParams[k] = reflect.New(expectedType).Elem()
	}

	return reflectedParams
}

func (call callStruct) invoke() (result []interface{}) {
	res := call.reflectedFunc.Call(call.reflectedParams)

	for _, val := range res {
		result = append(result, val.Interface())
	}

	return
}
