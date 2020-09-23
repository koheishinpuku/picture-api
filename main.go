package main

import (
	// "fmt"

	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"local.packages/Controllers"
	"local.packages/Midd"
	"local.packages/Redd"
	// "time"
	// "os"
	// "bufio"
)

var secretKey string = os.Getenv("JWT_SECRET")

func testget() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello!!!!World!!!")
	}
}

func main() {
	e := echo.New()
	Redd.Init()
	e.Use(middleware.Recover()) //アプリケーションのどこかで予期せずにpanicを起こしてしまっても、サーバは落とさずにエラーレスポンスを返せるようにリカバリーするmiddleware
	e.Use(middleware.Logger())  //リクエスト単位にログを出してくれるmiddleware

	// e.POST("/image", Controllers.ImageProcess())

	e.POST("/signup", Controllers.Signup())
	e.POST("/login", Controllers.Login())
	v1 := e.Group("/v1")
	{
		v1.Use(middleware.JWTWithConfig(middleware.JWTConfig{
			SigningKey: []byte(secretKey),
			ContextKey: "rawToken",
		}))
		v1.Use(Midd.Authtest)
		v1.GET("/success", testget())
	}

	e.Logger.Fatal(e.Start(":5000"))

}
