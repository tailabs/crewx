package validator

import "testing"

func Test_MapStruct(t *testing.T) {
	type Struct struct {
		AccessToken string `mapstructure:"access_token" validate:"required"`
		Secret      string `validate:"required"`
	}

	mapData := map[string]string{
		"access_token": "this_is_access_token",
		"secret":       "this_is_secret",
	}

	s := Struct{}

	if err := MapStruct(mapData, s); err != nil {
		t.Fatal(err)
	}
}
