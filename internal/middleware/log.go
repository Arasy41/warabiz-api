package middleware

import (
	"encoding/json"
	"net/http"
	"time"

	cg "warabiz/api/pkg/constants/general"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

func (mw *MiddlewareManager) RequestLog() fiber.Handler {
	return func(c *fiber.Ctx) error {

		start := time.Now()

		err := c.Next()

		if fields, ok := c.Locals("logfields").(logrus.Fields); !ok {
			mw.logger.Error("failed to get log data")
			return err
		} else {

			statusCode := c.Response().StatusCode()

			logfileds := logrus.Fields{
				"request_id":      c.Locals(cg.CtxRequestID),
				"ip_address":      c.IP(),
				"request_method":  c.Method(),
				"url":             c.OriginalURL(),
				"status_code":     statusCode,
				"path":            c.Path(),
				"user_agent":      c.Get("User-Agent", "unkown"),
				"response_time":   time.Since(start).String(),
				"request_payload": getJsonString(c.Body()),
				"response_data":   getJsonString(c.Response().Body()),
			}
			for key, val := range fields {
				logfileds[key] = val
			}
			msg := fields["msg"]
			delete(logfileds, "msg")

			entry := mw.logger.WithFields(logfileds)

			if statusCode == http.StatusOK || statusCode == http.StatusCreated {
				entry.Info(msg)
			} else if statusCode == http.StatusBadRequest || statusCode == http.StatusUnauthorized || statusCode == http.StatusNotFound {
				entry.Warn(msg)
			} else {
				entry.Error(msg)
			}
		}

		return err
	}
}

func getJsonString(bytes []byte) map[string]interface{} {
	if bytes == nil || len(bytes) == 0 {
		return nil
	}
	var output map[string]interface{}
	err := json.Unmarshal(bytes, &output)
	if err != nil {
		return nil
	}
	return output
}
