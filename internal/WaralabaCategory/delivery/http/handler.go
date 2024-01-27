package http

import (
	"fmt"
	"net/http"
	"strconv"
	"warabiz/api/config"
	"warabiz/api/internal/WaralabaCategory/usecase"
	newscategories "warabiz/api/internal/models/category"
	"warabiz/api/pkg/http/exception"
	"warabiz/api/pkg/infra/logger"

	"github.com/gofiber/fiber/v2"
)

type CategoryHandler struct {
	usecase usecase.Usecase
	cfg     *config.Config
	logger  logger.Logger
}

func NewCategoryHandler(uc usecase.Usecase, cfg *config.Config, logger logger.Logger) CategoryHandler {
	return CategoryHandler{
		usecase: uc,
		cfg:     cfg,
		logger:  logger,
	}
}

func (h CategoryHandler) CreateCategory(c *fiber.Ctx) error {

	exc := exception.NewException(c, h.logger)

	//* Get Request
	req := new(newscategories.CreateCategoryRequest)
	if err := c.BodyParser(req); err != nil {
		return exc.WriteErrorResponse(http.StatusBadRequest, "periksa kembali input anda !", err.Error())
	}

	//* Usecase
	id, err := h.usecase.CreateCategory(c.Context(), exc, req)
	if err != nil {
		return exc.WriteParseError(err)
	}

	//* Create success response
	return exc.WriteSuccessResponse(http.StatusOK, "sukses", id)
}

func (h CategoryHandler) GetCategoryByID(c *fiber.Ctx) error {

	exc := exception.NewException(c, h.logger)

	//* Get Request
	id, err := strconv.Atoi(fmt.Sprintf("%v", c.Params("id")))
	if err != nil {
		return exc.WriteErrorResponse(http.StatusBadRequest, "id tidak valid", nil)
	}

	//* Usecase
	newsCategory, err := h.usecase.GetCategoryById(c.Context(), exc, int64(id))
	if err != nil {
		return exc.WriteParseError(err)
	}

	//* Create success response
	return exc.WriteSuccessResponse(http.StatusOK, "sukses", newsCategory)
}

func (h CategoryHandler) GetAllCategory(c *fiber.Ctx) error {

	exc := exception.NewException(c, h.logger)

	//* Get Request
	req := new(newscategories.GetAllCategoryRequest)
	if err := c.QueryParser(req); err != nil {
		return exc.WriteErrorResponse(http.StatusBadRequest, "periksa kembali input anda !", err.Error())
	}

	//* Usecase
	news, err := h.usecase.GetAllCategory(c.Context(), exc, req)
	if err != nil {
		return exc.WriteParseError(err)
	}

	//* Create succes response
	return exc.WriteSuccessResponse(http.StatusOK, "sukses", news)
}

func (h CategoryHandler) UpdateCategory(c *fiber.Ctx) error {

	exc := exception.NewException(c, h.logger)

	//* Get Request
	req := new(newscategories.UpdateCategoryRequest)
	if err := c.BodyParser(req); err != nil {
		return exc.WriteErrorResponse(http.StatusBadRequest, "periksa kembali input anda !", err.Error())
	}

	//* Usecase
	err := h.usecase.UpdateCategory(c.Context(), exc, req)
	if err != nil {
		return exc.WriteParseError(err)
	}

	//* Create success response
	return exc.WriteSuccessResponse(http.StatusOK, "sukses", nil)
}

func (h CategoryHandler) DeleteCategory(c *fiber.Ctx) error {

	exc := exception.NewException(c, h.logger)

	//* Get Request
	id, err := strconv.Atoi(fmt.Sprintf("%v", c.Params("id")))
	if err != nil {
		return exc.WriteErrorResponse(http.StatusBadRequest, "id tidak valid", nil)
	}

	//* Usecase
	err = h.usecase.DeleteCategory(c.Context(), exc, int64(id))
	if err != nil {
		return exc.WriteParseError(err)
	}

	//* Create success response
	return exc.WriteSuccessResponse(http.StatusOK, "sukses", nil)
}

// func (h CategoryHandler) GetCategoryDetails(c *fiber.Ctx) error {

// 	exc := exception.NewException(c, h.logger)

// 	//* Get Request
// 	id, err := strconv.Atoi(fmt.Sprintf("%v", c.Params("id")))
// 	if err != nil {
// 		return exc.WriteErrorResponse(http.StatusBadRequest, "id tidak valid", nil)
// 	}

// 	//* Usecase
// 	newsCategory, err := h.usecase.GetCategoryDetailById(c.Context(), exc, int64(id))
// 	if err != nil {
// 		return exc.WriteParseError(err)
// 	}

// 	//* Create success response
// 	return exc.WriteSuccessResponse(http.StatusOK, "sukses", newsCategory)
// }
