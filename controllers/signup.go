package Controllers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"local.packages/DB"
	"local.packages/Models"
)

func Signup() echo.HandlerFunc {
	return func(c echo.Context) error {

		var userdata []Models.Userdata

		db := DB.GormConnect()

		hash, err := bcrypt.GenerateFromPassword([]byte(c.FormValue("password")), 12)
		if err != nil {
			return c.JSON(http.StatusOK, `message:"Passwordのハッシュ化に失敗しました"`)
		}

		newTable := Models.Userdata{
			Id:       0,
			Username: c.FormValue("username"),
			Email:    c.FormValue("email"),
			Password: string(hash),
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
