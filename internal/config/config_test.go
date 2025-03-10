package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfigErrors(t *testing.T) {
	// Arrange
	testCases := []struct {
		name            string
		packSizesString string
		expectedError   error
	}{
		{
			name:            "PACK_SIZES not set",
			packSizesString: "",
			expectedError:   newEnvVarNotSetError(envVarNamePackSizes),
		},
		{
			name:            "PACK_SIZES contains invalid characters",
			packSizesString: "100,200a,300",
			expectedError:   newInvalidPackSizeValueError("200a"),
		},
		{
			name:            "PACK_SIZES contains negative values",
			packSizesString: "100,-200,300",
			expectedError:   newInvalidPackSizeValueError("-200"),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			os.Setenv(envVarNamePackSizes, testCase.packSizesString)

			// Act
			config, err := LoadConfig()

			// Assert
			assert.Nil(t, config)
			assert.EqualError(t, err, testCase.expectedError.Error())
		})
	}
}
