package exception

import (
	"fmt"
	"net/http"

	cg "warabiz/api/pkg/constants/general"
	"warabiz/api/pkg/infra/logger"
	"warabiz/api/pkg/utils/converter"
	"warabiz/api/pkg/utils/generator"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type Exception interface {
	Error() string
	StatusCode() int
	Detail() interface{}
	IsError() bool
	ExportLogFields() logrus.Fields
	NewRestError(statusCode int, msg string, detail interface{}) Exception
	SetLog(key string, value interface{}) Exception
	WriteErrorResponse(statusCode int, msg string, detail interface{}) error
	WriteSuccessResponse(statusCode int, msg string, data interface{}) error
	WriteParseError(restErr error) error
	ParseError(restErr error) Exception
}

type exception struct {
	ctx       *fiber.Ctx
	isError   bool
	errStatus int
	errMsg    string
	errDetail interface{}
	logger    logger.Logger
	logfields logrus.Fields
}

func NewException(c *fiber.Ctx, log logger.Logger) *exception {
	return &exception{ctx: c, logger: log, logfields: logrus.Fields{}}
}

// func NewExceptionFromContext(ctx context.Context, log logger.Logger) *exception {
// 	excRaw := ctx.Value("exception")
// 	if exc, ok := excRaw.(*exception); ok {
// 		return exc
// 	}
// 	return NewException(log)
// }

// func (e *exception) ExportInContext(ctx context.Context) context.Context {
// 	return context.WithValue(ctx, "exception", e)
// }

func (e *exception) Error() string {
	return e.errMsg
}

func (e *exception) StatusCode() int {
	return e.errStatus
}

func (e *exception) Detail() interface{} {
	return e.errDetail
}

func (e *exception) IsError() bool {
	return e.isError
}

func (e *exception) ExportLogFields() logrus.Fields {
	return e.logfields
}

func (e *exception) NewRestError(statusCode int, msg string, detail interface{}) Exception {
	e.isError = true
	e.errStatus = statusCode
	e.errMsg = msg
	e.errDetail = detail
	return e
}

func (e *exception) SetLog(key string, value interface{}) Exception {
	e.logfields[key] = value
	return e
}

type Response struct {
	Status    Status      `json:"status"`
	RequestID string      `json:"request_id"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Errors    interface{} `json:"errors,omitempty"`
	TimeStamp string      `json:"time_stamp"`
}

type Status struct {
	Code int    `json:"code"`
	Name string `json:"name"`
}

func (e *exception) WriteErrorResponse(statusCode int, msg string, detail interface{}) error {

	requestID := fmt.Sprintf("%v", e.ctx.Locals(cg.CtxRequestID))
	timeLocStr := fmt.Sprintf("%v", e.ctx.Locals(cg.CtxTimeZone))
	timestamp := generator.GenerateTimeNowLocal(converter.GetTimeLocation(timeLocStr)).Format(cg.FullTimeFormat)

	logfields := logrus.Fields{
		"time_stamp":  timestamp,
		"status_code": statusCode,
		"msg":         msg,
	}
	for key, val := range e.logfields {
		logfields[key] = val
	}

	e.ctx.Locals("logfields", logfields)

	return e.ctx.Status(statusCode).JSON(&Response{
		Status: Status{
			Code: statusCode,
			Name: http.StatusText(statusCode),
		},
		RequestID: requestID,
		Message:   msg,
		Errors:    detail,
		TimeStamp: timestamp,
	})
}

func (e *exception) WriteSuccessResponse(statusCode int, msg string, data interface{}) error {

	requestID := fmt.Sprintf("%v", e.ctx.Locals(cg.CtxRequestID))
	timeLocStr := fmt.Sprintf("%v", e.ctx.Locals(cg.CtxTimeZone))
	timestamp := generator.GenerateTimeNowLocal(converter.GetTimeLocation(timeLocStr)).Format(cg.FullTimeFormat)

	logfields := logrus.Fields{
		"time_stamp":  timestamp,
		"status_code": statusCode,
		"msg":         msg,
	}
	for key, val := range e.logfields {
		logfields[key] = val
	}

	e.ctx.Locals("logfields", logfields)

	return e.ctx.Status(statusCode).JSON(&Response{
		Status: Status{
			Code: statusCode,
			Name: http.StatusText(statusCode),
		},
		RequestID: requestID,
		Message:   msg,
		Data:      data,
		TimeStamp: timestamp,
	})
}

func (e *exception) WriteParseError(restErr error) error {
	exc := e.ParseError(restErr)
	e.logfields = exc.ExportLogFields()
	return e.WriteErrorResponse(exc.StatusCode(), exc.Error(), exc.Detail())
}

func (e *exception) ParseError(restErr error) Exception {
	if exc, ok := restErr.(Exception); ok {
		return exc
	}
	return e.NewRestError(http.StatusInternalServerError, "failed to parse error", nil)
}
