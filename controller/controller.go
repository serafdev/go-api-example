package controller

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
)

// GetItems godoc
// @Summary      GetItems
// @Description  get items
// @Tags         items
// @Accept       json
// @Produce      json
// @Success      200  {array}   Item
// @Router       /items [get]
func GetItems(c echo.Context) error {
	return c.JSON(http.StatusOK, []Item{
		{
			Information:     "An Item",
			MoreInformation: 42,
		},
		{
			Information:     "Another Item",
			MoreInformation: 90,
		},
	})
}
