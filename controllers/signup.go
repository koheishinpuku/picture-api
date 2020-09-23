package Controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"local.packages/DB"
	"local.packages/Models"
)

var userdata []Models.Userdata

func Signup() echo.HandlerFunc {
	return func(c echo.Context) error {

		db := DB.GormConnect()

		newTable := Models.Userdata{
			Id:       0,
			Username: c.FormValue("username"),
			Email:    c.FormValue("email"),
			Password: c.FormValue("password"),
		}

		db.Where("username = ?", newTable.Username).Find(&userdata)
		if len(userdata) != 0 {
			return c.JSON(http.StatusOK, `message:"そのUsernameは使われている"`)
		}
		db.Where("username = ?", newTable.Email).Find(&userdata)
		if len(userdata) != 0 {
			return c.JSON(http.StatusOK, `message:"そのEmailは使われている"`)
		}

		db.Create(&newTable)

		return c.JSON(http.StatusOK, `message:"データベースに追加しました"`)
	}
}
