package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/albert-wang/tracederror"
)

type Config struct {
	Debug bool
	Port  int

	PostgresConnectionURL string
	RedisHost             string
	RedisPassword         string
}

// Sets fields in a struct from environment variables, according to their `env` tag.
// Only sets values with a non-empty environment variable, leaves them unmodified otherwise.
// Only supports string, int and bool variables.
func LoadConfigurationFromEnvironmentVariables(cfg interface{}) error {
	val := reflect.ValueOf(cfg)
	if val.Kind() != reflect.Ptr {
		return fmt.Errorf("Input not a pointer to struct")
	}

	derefed := reflect.Indirect(val)
	if derefed.Kind() != reflect.Struct {
		return fmt.Errorf("Input not a pointer to struct")
	}

	derefedType := derefed.Type()
	for i := 0; i < derefedType.NumField(); i++ {
		typeField := derefedType.Field(i)

		env := typeField.Tag.Get("env")
		if len(env) == 0 {
			continue
		}

		value := os.Getenv(env)
		value = strings.TrimSpace(value)
		if len(value) == 0 {
			continue
		}

		field := derefed.Field(i)
		switch field.Kind() {
		case reflect.String:
			{
				field.SetString(value)
				break
			}
		case reflect.Bool:
			{
				if value == "false" || value == "0" {
					field.SetBool(false)
				} else if value == "true" || value == "1" {
					field.SetBool(true)
				} else {
					return fmt.Errorf("Invalid boolean format for environment variable %s", env)
				}

				break
			}
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			{
				parsed, err := strconv.ParseInt(value, 10, 64)
				if err != nil {
					return err
				}

				field.SetInt(parsed)
				break
			}
		default:
			return fmt.Errorf("Unsupported type in struct at %s", typeField.Name)
		}
	}

	return nil
}

func LoadConfigurationFromFileAndEnvironment(file string, cfg interface{}) error {
	bytes, err := ioutil.ReadFile(file)
	if err != nil {
		return tracederror.New(err)
	}

	err = json.Unmarshal(bytes, cfg)

	if err != nil {
		return tracederror.New(err)
	}

	err = LoadConfigurationFromEnvironmentVariables(cfg)
	return tracederror.New(err)
}
