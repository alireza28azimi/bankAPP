package uservalidator

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"main.go/dto"
	"regexp"
)

func (v Validator) ValidatorRegisterRequest(req dto.RegisterRequest) (map[string]string, error) {
	if err := validation.ValidateStruct(&req,

		validation.Field(&req.Name,
			validation.Required,
			validation.Length(3, 50)),

		validation.Field(&req.Password,
			validation.Required,
			validation.Match(regexp.MustCompile(`^[A-Za-z0-9!@#%^&*]{8,}$`))),
		validation.Field(&req.Email,
			validation.Required,
			validation.Match(regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`))),
		validation.Field(&req.PhoneNumber,
			validation.Required,
			validation.Match(regexp.MustCompile(phoneNumberRegex))),
	); err != nil {
		return nil, err
	}
	return nil, nil
}

func (v Validator) checkPhoneNumberUniqueness(value interface{}) error {
	phoneNumber := value.(string)

	if isUnique, err := v.repo.IsPhoneNumberUnique(phoneNumber); err != nil || !isUnique {
		if err != nil {
			return err
		}

		if !isUnique {
			return fmt.Errorf("phone number is not unique", err)
		}
		return nil
	}
	return nil
}
