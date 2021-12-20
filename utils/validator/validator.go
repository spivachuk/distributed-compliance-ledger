// Copyright 2020 DSR Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validator

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	uni *ut.UniversalTranslator
	vl  *validator.Validate
)

//nolint:wrapcheck,errcheck
func Validate(s interface{}) error {
	en := en.New()
	uni = ut.New(en, en)

	trans, _ := uni.GetTranslator("en")

	vl = validator.New()

	_ = vl.RegisterValidation("address", validateAddress)
	_ = en_translations.RegisterDefaultTranslations(vl, trans)

	_ = vl.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is a required field", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())

		return t
	})

	_ = vl.RegisterTranslation("required_with", trans, func(ut ut.Translator) error {
		return ut.Add("required_with", "{0} is a required field", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required_with", fe.Field())

		return t
	})
	_ = vl.RegisterTranslation("address", trans, func(ut ut.Translator) error {
		return ut.Add("address", "Field {0} : {1} is not a valid address", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("address", fe.Field(), fmt.Sprintf("%v", fe.Value()))

		return t
	})

	vl.RegisterTranslation("max", trans, func(ut ut.Translator) error {
		return ut.Add("max", "maximum length for {0} allowed is {1}", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("max", fe.Field(), fe.Param())

		return t
	})

	vl.RegisterTranslation("url", trans, func(ut ut.Translator) error {
		return ut.Add("url", "Field {0} : {1} is not a valid url", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("url", fe.Field(), fmt.Sprintf("%v", fe.Value()))

		return t
	})

	vl.RegisterTranslation("startsnotwith", trans, func(ut ut.Translator) error {
		return ut.Add("startsnotwith", "Field {0} : {1} is not a valid url", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("startsnotwith", fe.Field(), fmt.Sprintf("%v", fe.Value()))
		return t
	})

	errs := vl.Struct(s)

	if errs != nil {
		for _, e := range errs.(validator.ValidationErrors) {
			if e.Tag() == "max" {
				return sdkerrors.Wrap(ErrFieldMaxLengthExceeded, e.Translate(trans))
			}

			if e.Tag() == "required" || e.Tag() == "required_with" {
				return sdkerrors.Wrap(ErrRequiredFieldMissing, e.Translate(trans))
			}

			if e.Tag() == "url" || e.Tag() == "startsnotwith" || e.Tag() == "address" {
				return sdkerrors.Wrap(ErrFieldNotValid, e.Translate(trans))
			}

		}
	}

	return nil
}

func validateAddress(fl validator.FieldLevel) bool {
	if sdk.VerifyAddressFormat(fl.Field().Bytes()) != nil {
		return false
	} else {
		return true
	}
}
