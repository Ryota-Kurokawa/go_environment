// package main

// import "github.com/gin-gonic/gin"

// func main() {
// 	router := gin.Default()

// 	router.GET("/", func(c *gin.Context) {
// 		c.JSON(200, gin.H{
// 			"message": "Hello World",
// 		})
// 	})

//		router.Run(":3306")
//	}
package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1.engineを作成します


	import (
		"log"

		"github.com/gin-gonic/gin"
	)

	func main() {
		engine := infra.DBInit()

		// 2.サービスを作成します
		factory := service.NewService(engine)

		// 3.最後にengineを閉めます
		defer func() {
			log.Println("engine closed")
			engine.Close()
		}()

		// 4.Ginを作成します
		g := gin.Default()

		// 5.サービス層のMiddlewareを作成します
		g.Use(service.ServiceFactoryMiddleware(factory))

		// 6.v1というグループを作成します
		routes := g.Group("/v1")

		// 7.ルーティングです
		{
			routes.POST("/users", handler.Create)
			routes.GET("/users", handler.GetAll)
			routes.GET("/users/:user-id", handler.GetOne)
			routes.PUT("/users/:user-id", handler.Update)
			routes.DELETE("/users/:user-id", handler.Delete)
		}

		// 8.デフォルトは8080です
		g.Run(":3000")
	}

	// 5.サービス層のMiddlewareを作成します
	g.Use(service.ServiceFactoryMiddleware(factory))

	// 6.v1というグループを作成します
	routes := g.Group("/v1")

	// 7.ルーティングです
	{
		routes.POST("/users", handler.Create)
		routes.GET("/users", handler.GetAll)
		routes.GET("/users/:user-id", handler.GetOne)
		routes.PUT("/users/:user-id", handler.Update)
		routes.DELETE("/users/:user-id", handler.Delete)
	}
	// 8.デフォルトは8080です
	g.Run(":3000")
}
