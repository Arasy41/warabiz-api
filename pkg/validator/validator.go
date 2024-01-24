package validator

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	cg "warabiz/api/pkg/constants/general"
	"warabiz/api/pkg/http/exception"

	"golang.org/x/exp/slices"
)

const (
	minimalPasswordLength = 8
)

func DefaultDate(value string) bool {
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	return re.MatchString(value)
}

func DefaultPhoneNumber(value string) bool {
	re := regexp.MustCompile(`^(\+62|62)?[\s-]?0?8[1-9]{1}\d{1}[\s-]?\d{4}[\s-]?\d{2,5}$`)
	return re.MatchString(value)
}

func DefaultCompareString(value1 string, value2 string) bool {
	return value1 == value2
}

func DefautlMinimalLengthPassword(value string) bool {
	if len(value) < minimalPasswordLength {
		return false
	}
	return true
}

func PasswordValidator(password string, defaultLength int, owner ...string) error {

	// Regex pattern untuk memeriksa adanya huruf kecil
	lowercaseRegex := regexp.MustCompile(`[a-z]`)
	// Regex pattern untuk memeriksa adanya huruf besar
	uppercaseRegex := regexp.MustCompile(`[A-Z]`)

	isValid := true

	if len(password) < defaultLength {
		isValid = false
	}
	if !lowercaseRegex.MatchString(password) {
		isValid = false
	}
	if !uppercaseRegex.MatchString(password) {
		isValid = false
	}

	common := cg.CommonPasswords
	for i := 0; i < len(owner); i++ {
		owner[i] = strings.ToLower(owner[i])
		common = append(common, owner[i])
		common = append(common, owner[i]+"123")
		common = append(common, owner[i]+"1234")
	}
	if len(owner) == 2 {
		common = append(common, strings.ToLower(owner[0]+strings.ToLower(owner[1])))
		common = append(common, strings.ToLower(owner[0]+strings.ToLower(owner[1])+"123"))
		common = append(common, strings.ToLower(owner[0]+strings.ToLower(owner[1])+"1234"))
	}
	if !isValid {
		return fmt.Errorf("Mohon gunakan %v atau lebih karakter dengan kombinasi huruf kapital, huruf kecil, angka dan simbol khusus", defaultLength)
	}
	if slices.Contains(common, strings.ToLower(password)) {
		return errors.New("password terlalu umum")
	}
	return nil
}

func DefaultUsernameValidator(username string) exception.MapErrors {

	if username == "" {
		return nil
	}

	length := len(username)
	firstChar := username[0:1]
	lastChar := username[length-1 : length]
	errMap := exception.NewMapErr()

	rule1 := regexp.MustCompile(`^[a-zA-Z0-9._-]*$`)
	if !rule1.MatchString(username) {
		errMap.AppendErrors("username", "username hanya boleh alfabet, titik( . ), strip( - ), and garis bawah ( _ )")
	}
	rule2 := regexp.MustCompile(`^[a-zA-Z]*$`)
	if !rule2.MatchString(firstChar) {
		errMap.AppendErrors("username", "username tidak boleh diawali dengan spesial karakter")
	}
	if !rule2.MatchString(lastChar) {
		errMap.AppendErrors("username", "username tidak boleh diakhiri dengan spesial karakter")
	}
	if strings.Contains(username, "._") || strings.Contains(username, "_.") || strings.Contains(username, "..") || strings.Contains(username, "__") {
		errMap.AppendErrors("username", "username tidak boleh berisi urutan spesial karakter")
	}
	if length < 6 {
		errMap.AppendErrors("username", "username tidak boleh kurang dari 6 karakter")
	}
	if length > 30 {
		errMap.AppendErrors("username", "username tidak boleh lebih dari 30 karakter")
	}
	return errMap
}

func DateRangeValidation(dateStart, dateEnd string) (validDate, validRange bool) {
	validDate = true
	validRange = false
	loc, _ := time.LoadLocation("Asia/Jakarta")
	dateStartT, err := time.ParseInLocation(cg.LayoutDateOnly, dateStart, loc)
	if err != nil {
		validDate = false
	}
	dateEndT, err := time.ParseInLocation(cg.LayoutDateOnly, dateEnd, loc)
	if err != nil {
		validDate = false
	}
	if dateEndT.After(dateStartT) {
		validRange = true
	}
	return validDate, validRange
}

func DateValidationYYYY_MM_DD(date string) (validDate bool) {
	validDate = true
	loc, _ := time.LoadLocation("Asia/Jakarta")
	_, err := time.ParseInLocation(cg.LayoutDateOnly, date, loc)
	if err != nil {
		validDate = false
	}
	return validDate
}

func IsNilMap(m map[string]string) bool {
	if m == nil {
		return true
	}
	if len(m) == 0 {
		return true
	}
	return false
}
