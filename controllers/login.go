package Controllers

import (
	// "fmt"

	"fmt"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"

	"local.packages/DB"
	"local.packages/Models"
	"local.packages/Redd"
)

var secretKey string = os.Getenv("JWT_SECRET")

type ResponseLogin struct {
	Status string
}

//DBに情報があった場合にのみトークンを発行する処理
func Login() echo.HandlerFunc {
	return func(c echo.Context) error {
		var userdata Models.Userdata

		db := DB.GormConnect()
		res := ResponseLogin{
			Status: "ERROR",
		}

		user := new(Models.Userdata)
		if err := c.Bind(user); err != nil {
			res.Status = "Cannot get Response"
			c.JSON(http.StatusInternalServerError, res)
		}

		db.Where("username = ?", user.Username).Find(&userdata)

		//ハッシュ化したパスワードを確認
		err := bcrypt.CompareHashAndPassword([]byte(userdata.Password), []byte(user.Password))
		if err != nil {
			res.Status = "Password is incorrect"
			return c.JSON(http.StatusOK, res)
		}

		token := jwt.New(jwt.GetSigningMethod("HS256"))

		token.Claims = jwt.MapClaims{
			"username": user.Username,
			"email":    user.Email,
			"password": user.Password,
		}

		tokenString, err := token.SignedString([]byte(secretKey))
		if err != nil {
			res.Status = "CannotGetToken"
			c.JSON(http.StatusInternalServerError, res)
		}

		err = RedCreateToken(fmt.Sprintf("%s/%s", user.Username, Redd.CacheLoginToken), fmt.Sprintf("%s", user.Username))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, res)
		}
		res.Status = tokenString

		return c.JSON(http.StatusOK, res)
	}
}

//redisにセッション記録する
func RedCreateToken(key, value string) error {

	err := Redd.Rdb.Set(Redd.Ctx, key, value, 500*time.Second).Err()

	if err != nil {
		panic(err)
	}
	return err
}
