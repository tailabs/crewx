package validator

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/mitchellh/mapstructure"
)

// MapStruct decodes a map into a struct using mapstructure and validates the struct using the validator package.
// It returns an error if decoding or validation fails.
func MapStruct(m any, i any) error {
	// Decode the map to the struct using mapstructure
	if err := mapstructure.Decode(m, &i); err != nil {
		return fmt.Errorf("failed to decode map to struct: %w", err)
	}

	// Create a new validator instance with support for required fields
	v := validator.New(validator.WithRequiredStructEnabled())

	// Validate the struct using the validator
	if err := v.Struct(i); err != nil {
		// Handle validation errors
		var errs validator.ValidationErrors
		if errors.As(err, &errs) {
			errFields := make([]string, 0, len(errs))
			for _, fieldError := range errs {
				errFields = append(errFields, fieldError.StructField())
			}
			return fmt.Errorf("fields %+v validate failed", errFields)
		}
		return err
	}

	return nil
}
