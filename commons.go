package framework

import (
	"errors"
	"reflect"
)

//VariableType tipo de dato para tipo de variable
type VariableType int

const (
	//TypeInt Tipo de dato int
	TypeInt VariableType = iota
	//TypeBool Tipo de dato bool
	TypeBool
	//TypeString Tipo de dato string
	TypeString
	//TypeNone Tipo de dato no identificado
	TypeNone
	//ErrorNil error cuando la interface de entrada es nil
	ErrorNil = "in data is null"
	//ErrorTypeNone error con el tipo de dato
	ErrorTypeNone = "type data not supported"
)

//ResolveType retorna tipo de dato
func ResolveType(data interface{}) (VariableType, error) {

	var typeData VariableType //variable de retorno con el tipo de dato -> numero

	if data == nil {
		return TypeNone, errors.New(ErrorNil)
	}

	stringType := reflect.TypeOf(data).String()
	switch stringType {
	case "bool":
		typeData = TypeBool
	case "string":
		typeData = TypeString
	case "int":
		typeData = TypeInt
	default:
		return TypeNone, errors.New(ErrorTypeNone)
	}
	return typeData, nil
}

//MapValues validate map
func MapValues(values map[string]interface{}, entries []ConfigEntry) map[string]interface{} {
	if values == nil || entries == nil {
		panic("config not exists")
	}
	for _, entry := range entries {
		_, exists := values[entry.VariableName]
		if !exists {
			panic(entry.VariableName + " config not exists")
		}

	}
	return values
}
