package env

import (
	"os"
	"strconv"

	"github.com/pkg/errors"
)

type EnvReader func() error

func ReadEnvExec(readers ...EnvReader) error {
	for _, reader := range readers {
		if err := reader(); err != nil {
			return err
		}
	}
	return nil
}

func ReadInt(envName string, data *int) EnvReader {
	return func() error {
		strValue, ok := os.LookupEnv(envName)
		if !ok {
			return errors.Errorf("Missing env %q", envName)
		}
		intValue, err := strconv.Atoi(strValue)
		if err != nil {
			return errors.Errorf("Env %q is not an integer", envName)
		}

		*data = intValue
		return nil
	}
}

func ReadString(envName string, data *string) EnvReader {
	return func() error {
		value, ok := os.LookupEnv(envName)
		if !ok {
			return errors.Errorf("Missing env %q", envName)
		}

		*data = value
		return nil
	}
}
