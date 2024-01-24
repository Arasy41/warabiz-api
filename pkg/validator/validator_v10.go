package validator

import (
	"fmt"
	"net/mail"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"

	"warabiz/api/pkg/http/exception"
	"warabiz/api/pkg/utils/converter"
)

const (
	ValTag             = "validator_v10"
	requiredTag        = "required"
	requiredIfEmptyTag = "requiredIfEmpty"
	emailTag           = "email"
	dateTag            = "date"
	phoneNumberTag     = "phone"
	pwdCompareTag      = "pwd_compare"
	minTag             = "min"
	maxTag             = "max"
	mediaTag           = "media"
	usernameTag        = "username"
	lengthTag          = "len"
)

func ValidateStruct(req interface{}) exception.MapErrors {

	val := validator.New()
	mapErr := exception.NewMapErr()
	var err error

	val.RegisterValidation(dateTag, dateValidator)
	val.RegisterValidation(phoneNumberTag, phoneValidator)
	val.RegisterValidation(pwdCompareTag, pwdCompareValidator)
	val.RegisterValidation(mediaTag, mediaValidator)
	val.RegisterValidation(requiredIfEmptyTag, validateRequiredIfOtherFieldEmpty)

	err = val.Struct(req)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			mapErr.AppendErrors(ValTag, fmt.Sprintf("error: %v", err.Error()))
			return mapErr
		}
		for _, err := range err.(validator.ValidationErrors) {

			field := converter.ConvertToSnakeCase(err.Field())
			parrent := converter.ConvertToSnakeCase(err.Namespace())
			if parrent != "" {
				field = parrent
			}
			part := strings.Split(field, ".")
			field = part[len(part)-1]
			field = strings.TrimPrefix(strings.TrimSuffix(field, "_"), "_")

			switch err.Tag() {
			case requiredTag:
				mapErr.AppendErrors(field, "wajib diisi")
			case requiredIfEmptyTag:
				mapErr.AppendErrors(field, "wajib diisi")
			case emailTag:
				mapErr.AppendErrors(field, "email tidak valid")
			case dateTag:
				mapErr.AppendErrors(field, "format tanggal tidak valid")
			case phoneNumberTag:
				mapErr.AppendErrors(field, "nomor telepon tidak valid")
			case pwdCompareTag:
				mapErr.AppendErrors("password", "tidak sama")
				mapErr.AppendErrors("confirm_password", "tidak sama")
			case minTag:
				mapErr.AppendErrors(field, fmt.Sprintf("%s tidak boleh kurang dari %v karakter", field, err.Param()))
			case maxTag:
				mapErr.AppendErrors(field, fmt.Sprintf("%s tidak boleh lebih dari %v karakter", field, err.Param()))
			case mediaTag:
				mapErr.AppendErrors(field, fmt.Sprintf("format %s tidak valid", field))
			case lengthTag:
				mapErr.AppendErrors(field, fmt.Sprintf("panjang %s harus %v karakter", field, err.Param()))
			default:
				mapErr.AppendErrors(field, fmt.Sprintf("error: %v", err.Error()))
			}
		}
		return mapErr
	}
	return nil
}

func dateValidator(field validator.FieldLevel) bool {
	value := field.Field().String()
	return DefaultDate(value)
}

func phoneValidator(field validator.FieldLevel) bool {
	value := field.Field().String()
	return DefaultPhoneNumber(value)
}

func pwdCompareValidator(field validator.FieldLevel) bool {
	value, _, _, ok := field.GetStructFieldOK2()
	if !ok {
		return false
	}
	comparator := field.Field().Interface()
	compared := value.Interface()
	return comparator == compared
}

func emailValidator(field validator.FieldLevel) bool {
	value := field.Field().String()
	_, err := mail.ParseAddress(value)
	if err != nil {
		return false
	}
	return true
}

func mediaValidator(field validator.FieldLevel) bool {
	if emailValidator(field) || phoneValidator(field) {
		return true
	}
	return false
}

func validateRequiredIfOtherFieldEmpty(fl validator.FieldLevel) bool {

	field := fl.Field()
	structValue := reflect.ValueOf(fl.Parent())
	structType := structValue.Type()

	for i := 0; i < structType.NumField(); i++ {
		otherFieldValue, _, found := fl.GetStructFieldOK()
		if found && reflect.DeepEqual(otherFieldValue.Interface(), reflect.Zero(otherFieldValue.Type()).Interface()) {
			return !reflect.DeepEqual(field.Interface(), reflect.Zero(field.Type()).Interface())
		}
	}
	return true
}
