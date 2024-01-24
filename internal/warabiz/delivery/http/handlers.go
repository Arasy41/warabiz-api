package http

import (
	"fmt"
	"net/http"
	"strconv"
	"warabiz/api/config"
	"warabiz/api/internal/models/warabiz"
	"warabiz/api/internal/warabiz/usecase"
	"warabiz/api/pkg/http/exception"
	"warabiz/api/pkg/infra/logger"
	"warabiz/api/pkg/utils/getter"

	"github.com/gofiber/fiber/v2"
)

type WarabizHandler struct {
	usecase usecase.Usecase
	cfg     *config.Config
	logger  logger.Logger
}

func NewWarabizHandler(uc usecase.Usecase, cfg *config.Config, logger logger.Logger) WarabizHandler {
	return WarabizHandler{
		usecase: uc,
		cfg:     cfg,
		logger:  logger,
	}
}

func (h WarabizHandler) CreateWarabiz(c *fiber.Ctx) error {

	exc := exception.NewException(c, h.logger)

	//* Get Request
	req := new(warabiz.CreateWarabizRequest)
	if err := getter.GetFormDataRequest(c, req, h.logger); err != nil {
		return exc.WriteParseError(err)
	}

	//* Usecase
	id, err := h.usecase.CreateWarabiz(c.Context(), exc, req)
	if err != nil {
		return exc.WriteParseError(err)
	}

	//* Create success response
	return exc.WriteSuccessResponse(http.StatusOK, "sukses", id)
}

func (h WarabizHandler) GetWarabizByID(c *fiber.Ctx) error {

	exc := exception.NewException(c, h.logger)

	//* Get Request
	id, err := strconv.Atoi(fmt.Sprintf("%v", c.Params("id")))
	if err != nil {
		return exc.WriteErrorResponse(http.StatusBadRequest, "id tidak valid", nil)
	}

	//* Usecase
	warabiz, err := h.usecase.GetWarabizById(c.Context(), exc, int64(id))
	if err != nil {
		return exc.WriteParseError(err)
	}

	//* Create success response
	return exc.WriteSuccessResponse(http.StatusOK, "sukses", warabiz)
}

func (h WarabizHandler) GetAllWarabiz(c *fiber.Ctx) error {

	exc := exception.NewException(c, h.logger)

	//* Get Request
	req := new(warabiz.GetAllWarabizRequest)
	if err := c.QueryParser(req); err != nil {
		return exc.WriteErrorResponse(http.StatusBadRequest, "periksa kembali input anda !", err.Error())
	}

	//* Usecase
	warabizs, err := h.usecase.GetAllWarabiz(c.Context(), exc, req)
	if err != nil {
		return exc.WriteParseError(err)
	}

	//* Create succes response
	return exc.WriteSuccessResponse(http.StatusOK, "sukses", warabizs)
}

func (h WarabizHandler) UpdateWarabiz(c *fiber.Ctx) error {

	exc := exception.NewException(c, h.logger)

	//* Get Request
	req := new(warabiz.UpdateWarabizRequest)
	if err := getter.GetFormDataRequest(c, req, h.logger); err != nil {
		return exc.WriteParseError(err)
	}

	//* Usecase
	err := h.usecase.UpdateWarabiz(c.Context(), exc, req)
	if err != nil {
		return exc.WriteParseError(err)
	}

	//* Create success response
	return exc.WriteSuccessResponse(http.StatusOK, "sukses", nil)
}

func (h WarabizHandler) DeleteWarabiz(c *fiber.Ctx) error {

	exc := exception.NewException(c, h.logger)

	//* Get Request
	id, err := strconv.Atoi(fmt.Sprintf("%v", c.Params("id")))
	if err != nil {
		return exc.WriteErrorResponse(http.StatusBadRequest, "id tidak valid", nil)
	}

	//* Usecase
	err = h.usecase.DeleteWarabiz(c.Context(), exc, int64(id))
	if err != nil {
		return exc.WriteParseError(err)
	}

	//* Create success response
	return exc.WriteSuccessResponse(http.StatusOK, "sukses", nil)
}
