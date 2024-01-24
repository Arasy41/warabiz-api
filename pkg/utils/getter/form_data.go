package getter

import (
	"mime/multipart"
	"net/http"
	"reflect"
	"strconv"
	"strings"
	"warabiz/api/pkg/http/exception"
	"warabiz/api/pkg/infra/logger"

	"github.com/gofiber/fiber/v2"
)

func GetFormDataRequest(c *fiber.Ctx, req interface{}, log logger.Logger) error {

	exc := exception.NewException(c, log)
	v := reflect.ValueOf(req)
	typeOfT := v.Elem().Type()
	errMap := exception.NewMapErr()

	if v.Kind() != reflect.Ptr {
		return exc.NewRestError(http.StatusInternalServerError, "must pass a pointer, not a value, to RequestScan destination", nil)
	}

	for i := 0; i < v.Elem().NumField(); i++ {

		fieldValue := v.Elem().Field(i)
		fieldType := typeOfT.Field(i).Type
		jsonTag := typeOfT.Field(i).Tag.Get("json")
		validateTag := typeOfT.Field(i).Tag.Get("validate")
		isPtr := false

		if fieldValue.IsValid() && fieldValue.CanSet() {

			if isFile(fieldType) {
				val, err := c.FormFile(jsonTag)
				if val != nil {
					if err != nil {
						return exc.NewRestError(http.StatusInternalServerError, "gagal memuat file", err.Error())
					}
				}
				v.Elem().Field(i).Set(reflect.ValueOf(val))
			} else if isSlicesOfFile(fieldType) {
				form, err := c.MultipartForm()
				if err != nil {
					return exc.NewRestError(http.StatusInternalServerError, "gagal memuat multipart form data", err.Error())
				}
				v.Elem().Field(i).Set(reflect.ValueOf(form.File[jsonTag]))
			} else {
				if fieldType.Kind() == reflect.Ptr {
					fieldType = typeOfT.Field(i).Type.Elem()
					isPtr = true
				}

				if fieldType.Kind() == reflect.Int || fieldType.Kind() == reflect.Int16 || fieldType.Kind() == reflect.Int32 || fieldType.Kind() == reflect.Int64 {
					val := c.FormValue(jsonTag)
					if !strings.Contains(validateTag, "required") {
						if val == "" {
							continue
						}
					}
					intVal, err := strconv.Atoi(val)
					if err != nil {
						errMap.AppendErrors(jsonTag, "format tidak valid")
						continue
					}

					if isPtr {
						v.Elem().Field(i).Set(reflect.New(typeOfT.Field(i).Type.Elem()))
						v.Elem().Field(i).Elem().SetInt(int64(intVal))
					} else {
						v.Elem().Field(i).SetInt(int64(intVal))
					}
				} else if fieldType.Kind() == reflect.Uint || fieldType.Kind() == reflect.Uint16 || fieldType.Kind() == reflect.Uint32 || fieldType.Kind() == reflect.Uint64 {
					val := c.FormValue(jsonTag)
					if !strings.Contains(validateTag, "required") {
						if val == "" {
							continue
						}
					}
					uintVal, err := strconv.ParseUint(val, 10, 64)
					if err != nil {
						errMap.AppendErrors(jsonTag, "format tidak valid")
						continue
					}

					if isPtr {
						v.Elem().Field(i).Set(reflect.New(typeOfT.Field(i).Type.Elem()))
						v.Elem().Field(i).Elem().SetUint(uintVal)
					} else {
						v.Elem().Field(i).SetUint(uintVal)
					}
				} else if fieldType.Kind() == reflect.Float32 || fieldType.Kind() == reflect.Float64 {
					val := c.FormValue(jsonTag)
					if !strings.Contains(validateTag, "required") {
						if val == "" {
							continue
						}
					}
					floatVal, err := strconv.ParseFloat(val, 64)
					if err != nil {
						errMap.AppendErrors(jsonTag, "format tidak valid")
						continue
					}

					if isPtr {
						v.Elem().Field(i).Set(reflect.New(typeOfT.Field(i).Type.Elem()))
						v.Elem().Field(i).Elem().SetFloat(floatVal)
					} else {
						v.Elem().Field(i).SetFloat(floatVal)
					}
				} else if fieldType.Kind() == reflect.Bool {
					val := c.FormValue(jsonTag)
					if !strings.Contains(validateTag, "required") {
						if val == "" {
							continue
						}
					}
					boolVal, err := strconv.ParseBool(val)
					if err != nil {
						errMap.AppendErrors(jsonTag, "format tidak valid")
						continue
					}

					if isPtr {
						v.Elem().Field(i).Set(reflect.New(typeOfT.Field(i).Type.Elem()))
						v.Elem().Field(i).Elem().SetBool(boolVal)
					} else {
						v.Elem().Field(i).SetBool(boolVal)
					}
				} else if fieldType.Kind() == reflect.String {
					if isPtr {
						v.Elem().Field(i).Set(reflect.New(typeOfT.Field(i).Type.Elem()))
						v.Elem().Field(i).Elem().SetString(c.FormValue(jsonTag))
					} else {
						v.Elem().Field(i).SetString(c.FormValue(jsonTag))
					}
				} else {
					continue
				}
			}
		} else {
			return exc.NewRestError(http.StatusInternalServerError, "could not set data request", nil)
		}
	}

	if errMap.IsNotEmpty() {
		return exc.NewRestError(http.StatusBadRequest, "validation error", errMap)
	}
	return nil
}

func isFile(t reflect.Type) bool {
	var compared *multipart.FileHeader
	if t == reflect.TypeOf(compared) {
		return true
	}
	return false
}

func isSlicesOfFile(t reflect.Type) bool {
	var comparedSlices []*multipart.FileHeader
	if t == reflect.TypeOf(comparedSlices) {
		return true
	}
	return false
}
