package Controllers

import (
	// "fmt"

	"fmt"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	"local.packages/Models"
	"local.packages/Redd"
)

var secretKey string = os.Getenv("JWT_SECRET")

type ResponseLogin struct {
	Status string
}

func Login() echo.HandlerFunc {
	return func(c echo.Context) error {

		res := ResponseLogin{
			Status: "ERROR",
		}

		userdata := new(Models.Userdata)
		if err := c.Bind(userdata); err != nil {
			res.Status = "Cannot get Response"
			c.JSON(http.StatusInternalServerError, res)
		}

		token := jwt.New(jwt.GetSigningMethod("HS256"))

		token.Claims = jwt.MapClaims{
			"username": userdata.Username,
			"email":    userdata.Email,
			"password": userdata.Password,
		}

		tokenString, err := token.SignedString([]byte(secretKey))
		if err != nil {
			res.Status = "CannotGetToken"
			c.JSON(http.StatusInternalServerError, res)
		}

		err = RedCreateToken(fmt.Sprintf("%s/%s", userdata.Username, Redd.CacheLoginToken), fmt.Sprintf("%s", userdata.Username))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, res)
		}
		res.Status = tokenString

		return c.JSON(http.StatusOK, res)
	}
}

func RedCreateToken(key, value string) error {

	err := Redd.Rdb.Set(Redd.Ctx, key, value, 300*time.Second).Err()

	if err != nil {
		panic(err)
	}
	return err
}
