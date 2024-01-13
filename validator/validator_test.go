package validator

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type TestStruct struct {
	Name  string `validate:"required,min=3,max=20"`
	Email string `validate:"required,email"`
}

func TestMapStruct(t *testing.T) {
	tests := []struct {
		name        string
		inputMap    map[string]interface{}
		expectedErr bool
	}{
		{
			name: "Valid Input",
			inputMap: map[string]interface{}{
				"Name":  "John Doe",
				"Email": "john@example.com",
			},
			expectedErr: false,
		},
		{
			name: "Invalid Input - Missing Required Field",
			inputMap: map[string]interface{}{
				"Name": "John Doe",
				// Missing "Email" field
			},
			expectedErr: true,
		},
		{
			name: "Invalid Input - Invalid Email",
			inputMap: map[string]interface{}{
				"Name":  "John Doe",
				"Email": "invalid-email",
			},
			expectedErr: true,
		},
		{
			name: "Invalid Input - Decode Failure",
			inputMap: map[string]interface{}{
				"InvalidField": "value",
			},
			expectedErr: true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var example TestStruct
			err := MapStruct(tc.inputMap, &example)

			if tc.expectedErr {
				require.Error(t, err, "Expected an error but got nil")
			} else {
				require.NoError(t, err, "Expected no error but got %v", err)
			}
		})
	}
}
