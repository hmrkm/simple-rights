package main

import (
	"github.com/hmrkm/simple-rights/adapter"
	"github.com/labstack/echo/v4"
)

func Router(e *echo.Echo, ra adapter.Rights) {
	e.POST("/v1/rights", func(c echo.Context) error {
		req := adapter.RequestRights{}
		if err := c.Bind(&req); err != nil {
			return c.JSON(400, nil)
		}

		err := ra.Verify(req)

		if err != nil {
			return ErrorHandler(c, err)
		}

		return c.JSON(200, nil)
	})
}
