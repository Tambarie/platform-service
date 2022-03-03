package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	api "platform-service/internal/adapters/api/gas-platform-service"
	"platform-service/internal/adapters/repository/mongoDB"
	"platform-service/internal/core/helper"
	services "platform-service/internal/core/services/gas-platform-service"
	"platform-service/internal/core/shared"
	"platform-service/internal/ports"
	"time"
)

func main() {
	helper.InitializeLogDir()
	service_address, service_port, dbtype, mongodb_port, _, mongodb_DBHost, dbName, _ := helper.LoadConfig()
	mongoURL := fmt.Sprintf("%s://%s:%s", dbtype, mongodb_DBHost, mongodb_port)
	fmt.Println(mongoURL)
	fmt.Println(dbName)
	DBRepository := ConnectToMongo(mongoURL, dbName)
	platformService := services.New(DBRepository)

	// Managing the handler and routes
	platformHandler := api.NewHTTPHandler(platformService)
	router := gin.Default()
	router.Use(helper.LogRequest)

	// routes
	router.POST("/:api-gate-way/platform/categories", platformHandler.CreatePlatform())
	router.PUT("/:api-gate-way/platform/categories/reference/:category-reference", platformHandler.UpdatePlatform())
	router.GET("/:api-gate-way/platform/categories/reference/:category-reference", platformHandler.GetCategoryByReference())
	router.GET("/:api-gate-way/platform/categories/name/:name", platformHandler.GetCategoryByName())
	router.GET("/:api-gate-way/platform/categories/list/page/:page", platformHandler.GetPlatformPage())
	router.DELETE("/:api-gate-way/platform/categories/reference/:category-reference", platformHandler.DeleteCategoryByReference())

	// No route
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, helper.PrintErrorMessage("404", shared.NORESOURCEFOUND))
	})

	fmt.Println("sevice running on" + service_address + ":" + service_port)
	helper.LogEvent("info", fmt.Sprintf("started platform service on "+service_address+":"+service_port+"in"+time.Since(time.Now()).String()))
	_ = router.Run(":" + service_port)

}

func ConnectToMongo(mongoURL string, DBName string) ports.PlatformRepository {
	repo, err := mongoDB.NewMongoRepository(mongoURL, DBName, 2000)
	if err != nil {
		_ = helper.PrintErrorMessage("500", err.Error())
		log.Fatal(err)
	}

	return services.New(repo)
}
