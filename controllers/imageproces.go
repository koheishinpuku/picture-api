package Controllers

import (
	"io"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func ImageProcess() echo.HandlerFunc {
	return func(c echo.Context) error {
		file, err := c.FormFile("upload")

		if err != nil {
			panic(err)
		}
		src, _ := file.Open()
		// img, _, _ := image.Decode(src)

		f, err := os.Create("./image/test.jpg")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		io.Copy(f, src)

		return c.JSON(http.StatusOK, `"message":"アップロードしました"`)
	}
}
