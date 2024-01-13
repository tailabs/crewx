package validator

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/mitchellh/mapstructure"
)

func MapStruct(m any, i any) error {
	if err := mapstructure.Decode(m, &i); err != nil {
		return fmt.Errorf("failed to decode map to struct: %w", err)
	}

	v := validator.New(validator.WithRequiredStructEnabled())
	if err := v.Struct(i); err != nil {
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
