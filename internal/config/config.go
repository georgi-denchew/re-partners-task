package config

import (
	"os"
	"strconv"
	"strings"
)

// Config is a type into which environment variables are parsed
type Config struct {
	// needs to be ordered
	PackSizes []int
}

const (
	envVarNamePackSizes = "PACK_SIZES"
)

// LoadConfig is a function which attempts to read environment variables and parse them into a Config variable.
// An error is returned if parsing encounters failures.
func LoadConfig() (*Config, error) {
	packSizesStr := os.Getenv(envVarNamePackSizes)
	if packSizesStr == "" {
		return nil, newEnvVarNotSetError(envVarNamePackSizes)
	}

	config := &Config{
		PackSizes: make([]int, 0),
	}

	iteratorFunc := strings.SplitSeq(packSizesStr, ",")

	var err error

	iteratorFunc(func(packSizeStr string) bool {
		packSizeParsed, nestedErr := strconv.Atoi(packSizeStr)
		if nestedErr != nil || packSizeParsed <= 0 {
			err = newInvalidPackSizeValueError(packSizeStr)
			return false
		}

		config.PackSizes = append(config.PackSizes, packSizeParsed)
		return true
	})

	if err != nil {
		return nil, err
	}

	return config, nil
}
