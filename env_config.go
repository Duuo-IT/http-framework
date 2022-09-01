package framework

import (
	"errors"
	"github.com/spf13/viper"
	"reflect"
	"strconv"
	"time"
)

//ConfigEntry entrada de configuración
type ConfigEntry struct {
	VariableName string
	Description  string
	Shortcut     string
	DefaultValue interface{}
}

//Configure method
func Configure(entries []ConfigEntry) (map[string]interface{}, error) {

	if len(entries) == 0 {
		return nil, errors.New("no entries given")
	}

	//Configuration for environment variables on the system
	viper.AutomaticEnv()

	values := make(map[string]interface{})

	for _, entry := range entries {
		valType, err := validateEntry(entry)
		if err != nil {
			return nil, errors.New("error on ResolveType")
		}

		name := entry.VariableName
		val := viperConfiguration(name, entry, valType)

		switch valType {
		case TypeInt:
			values[name] = val.(int)
		case TypeBool:
			values[name] = val.(bool)
		case TypeString:
			values[name] = val.(string)
		}
	}
	res := MapValues(values, entries)
	return res, nil
}

/*
Sets the value by getting the viper value
If there's no viper value, sets the default value
*/
func viperConfiguration(name string, entry ConfigEntry, valType VariableType) interface{} {
	val := viper.Get(name)
	if val == nil {
		val = entry.DefaultValue
	}

	if strType := reflect.TypeOf(val).String(); valType != TypeString && strType == "string" {
		strValue := val.(string)
		if intValue, err := strconv.Atoi(strValue); err == nil {
			return intValue
		}
		if boolValue, err := strconv.ParseBool(strValue); err == nil {
			return boolValue
		}
	}

	//checks for bool/int values that are set as string by enviroment variables
	return val
}

/*
Validates if the entry is accepted by the typeResolver
Returns type and an error
*/
func validateEntry(entry ConfigEntry) (VariableType, error) {
	val, err := ResolveType(entry.DefaultValue)

	if err != nil {
		return TypeNone, errors.New("error on ResolveType")
	}
	return val, nil
}

//BaseVariables base variables for API
type BaseVariables struct {
	Port         string
	LoggingLevel string
	Timeout      time.Duration
	URIPrefix    string
}

//GetBaseVariables configuration for base variables
func GetBaseVariables(cfg map[string]interface{}) *BaseVariables {
	return &BaseVariables{
		Port:         cfg["port"].(string),
		LoggingLevel: cfg["logging_level"].(string),
		Timeout:      time.Duration(cfg["timeout"].(int)) * time.Second,
		URIPrefix:    cfg["uri_prefix"].(string),
	}
}
