package Midd

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"local.packages/Redd"
)

func Authtest(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("rawToken").(*jwt.Token)
		fmt.Println(user)
		claims := user.Claims.(jwt.MapClaims)

		if claims["username"] == nil {
			return c.JSON(http.StatusUnauthorized, `"status": "USER_UNAUTHORIZED"`)
		}
		uname := claims["username"].(string)

		findResult := findToken(uname)

		if findResult == true {
			return next(c)
			// return c.String(http.StatusOK, uname+"は登録されているユーザーです")
		} else {
			return c.String(http.StatusUnauthorized, `"status": "USER_UNAUTHORIZED"`)
		}

	}
}

func findToken(uid string) bool {
	val, err := Redd.Rdb.Get(Redd.Ctx, fmt.Sprintf("%s/%s", uid, Redd.CacheLoginToken)).Result()
	if err != nil {
		// panic(err)
		fmt.Println("redisにdataにありませんでした")
		return false
	}
	fmt.Println(val)
	return true
}
