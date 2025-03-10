package config

import "fmt"

type EnvVarNotSetError struct {
	envVarName string
}

func newEnvVarNotSetError(envVarName string) *EnvVarNotSetError {
	return &EnvVarNotSetError{
		envVarName: envVarName,
	}
}

func (e *EnvVarNotSetError) Error() string {
	return fmt.Sprintf("error: %s env is not set", e.envVarName)
}

type InvalidPackSizeValueError struct {
	value string
}

func newInvalidPackSizeValueError(value string) *InvalidPackSizeValueError {
	return &InvalidPackSizeValueError{value: value}
}

func (e *InvalidPackSizeValueError) Error() string {
	return fmt.Sprintf("error: parsing package size value to int: %s", e.value)
}
