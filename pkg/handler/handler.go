package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"random-generator-API/models"
	"random-generator-API/pkg"
)

type Handler struct {
	useCase pkg.UseCase
}

func NewHandler(useCase pkg.UseCase) *Handler {
	return &Handler{
		useCase: useCase,
	}
}

// Generate godoc
// @Summary Generate amount of random values
// @Tags generate
// @Description Generate amount of random values
// @ID generate
// @Accept  json
// @Produce  json
// @Param input body models.Amount true "amount info"
// @Success 200 {object} statusResponse
// @Failure 400,404 {object} Error
// @Failure 500 {object} Error
// @Failure default {object} Error
// @Router /generator/generate [post]

func (h *Handler) Generate(c *gin.Context) {
	input := new(models.Amount)
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	ch := make(chan string, input.Amount)
	err := h.useCase.Generate(input, ch)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, statusResponse{
		Status: "ok",
	})

}

// Result godoc
// @Summary Get Last Result
// @Tags result
// @Description get last generator result
// @ID get-result
// @Accept  json
// @Produce  json
// @Success 200 {object} models.RandomItem
// @Failure 400,404 {object} Error
// @Failure 500	{object} Error
// @Failure default	{object} Error
// @Router generator/result [get]

func (h *Handler) Result(c *gin.Context) {
	res, err := h.useCase.GetLastOutput()

	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}
