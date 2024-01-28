package http

import (
	"fmt"
	"net/http"
	"strconv"
	"warabiz/api/config"
	waraCareer "warabiz/api/internal/models/wara_career"
	"warabiz/api/internal/wara_career/usecase"
	"warabiz/api/pkg/http/exception"
	"warabiz/api/pkg/infra/logger"

	"github.com/gofiber/fiber/v2"
)

type WaraCareerHandler struct {
	usecase usecase.Usecase
	cfg     *config.Config
	logger  logger.Logger
}

func NewWaraCareerHandler(uc usecase.Usecase, cfg *config.Config, logger logger.Logger) WaraCareerHandler {
	return WaraCareerHandler{
		usecase: uc,
		cfg:     cfg,
		logger:  logger,
	}
}

func (h WaraCareerHandler) CreateWaraCareer(c *fiber.Ctx) error {

	exc := exception.NewException(c, h.logger)

	//* Get Request
	req := new(waraCareer.CreateWaraCareerRequest)
	if err := c.BodyParser(req); err != nil {
		return exc.WriteErrorResponse(http.StatusBadRequest, "periksa kembali input anda !", err.Error())
	}

	//* Usecase
	id, err := h.usecase.CreateWaraCareer(c.Context(), exc, req)
	if err != nil {
		return exc.WriteParseError(err)
	}

	//* Create success response
	return exc.WriteSuccessResponse(http.StatusOK, "sukses", id)
}

func (h WaraCareerHandler) GetWaraCareerByID(c *fiber.Ctx) error {

	exc := exception.NewException(c, h.logger)

	//* Get Request
	id, err := strconv.Atoi(fmt.Sprintf("%v", c.Params("id")))
	if err != nil {
		return exc.WriteErrorResponse(http.StatusBadRequest, "id tidak valid", nil)
	}

	//* Usecase
	WaraCareer, err := h.usecase.GetWaraCareerById(c.Context(), exc, int64(id))
	if err != nil {
		return exc.WriteParseError(err)
	}

	//* Create success response
	return exc.WriteSuccessResponse(http.StatusOK, "sukses", WaraCareer)
}

func (h WaraCareerHandler) GetAllWaraCareer(c *fiber.Ctx) error {

	exc := exception.NewException(c, h.logger)

	//* Get Request
	req := new(waraCareer.GetAllWaraCareerRequest)
	if err := c.QueryParser(req); err != nil {
		return exc.WriteErrorResponse(http.StatusBadRequest, "periksa kembali input anda !", err.Error())
	}

	//* Usecase
	waraCareer, err := h.usecase.GetAllWaraCareer(c.Context(), exc, req)
	if err != nil {
		return exc.WriteParseError(err)
	}

	//* Create succes response
	return exc.WriteSuccessResponse(http.StatusOK, "sukses", waraCareer)
}

func (h WaraCareerHandler) UpdateWaraCareer(c *fiber.Ctx) error {

	exc := exception.NewException(c, h.logger)

	//* Get Request
	req := new(waraCareer.UpdateWaraCareerRequest)
	if err := c.BodyParser(req); err != nil {
		return exc.WriteErrorResponse(http.StatusBadRequest, "periksa kembali input anda !", err.Error())
	}

	//* Usecase
	err := h.usecase.UpdateWaraCareer(c.Context(), exc, req)
	if err != nil {
		return exc.WriteParseError(err)
	}

	//* Create success response
	return exc.WriteSuccessResponse(http.StatusOK, "sukses", nil)
}

func (h WaraCareerHandler) DeleteWaraCareer(c *fiber.Ctx) error {

	exc := exception.NewException(c, h.logger)

	//* Get Request
	id, err := strconv.Atoi(fmt.Sprintf("%v", c.Params("id")))
	if err != nil {
		return exc.WriteErrorResponse(http.StatusBadRequest, "id tidak valid", nil)
	}

	//* Usecase
	err = h.usecase.DeleteWaraCareer(c.Context(), exc, int64(id))
	if err != nil {
		return exc.WriteParseError(err)
	}

	//* Create success response
	return exc.WriteSuccessResponse(http.StatusOK, "sukses", nil)
}

// func (h WaraCareerHandler) GetWaraCareerDetails(c *fiber.Ctx) error {

// 	exc := exception.NewException(c, h.logger)

// 	//* Get Request
// 	id, err := strconv.Atoi(fmt.Sprintf("%v", c.Params("id")))
// 	if err != nil {
// 		return exc.WriteErrorResponse(http.StatusBadRequest, "id tidak valid", nil)
// 	}

// 	//* Usecase
// 	WaraCareer, err := h.usecase.GetWaraCareerDetailById(c.Context(), exc, int64(id))
// 	if err != nil {
// 		return exc.WriteParseError(err)
// 	}

// 	//* Create success response
// 	return exc.WriteSuccessResponse(http.StatusOK, "sukses", WaraCareer)
// }
