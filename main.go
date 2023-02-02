// package main
// Example of go-api
package main

import (
	echo "github.com/labstack/echo/v4"
	"github.com/serafdev/go-api-example/controller"
	_ "github.com/serafdev/go-api-example/docs"
	swagger "github.com/swaggo/echo-swagger"
)

// @contact.name   Seraf
// @contact.url    http://seraf.dev
// @contact.email  seraf@example.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
func main() {

	e := echo.New()
	e.GET("/swagger/*", swagger.WrapHandler)
	e.GET("/items", controller.GetItems)
	e.Logger.Fatal(e.Start(":1323"))
}
