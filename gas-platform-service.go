package main

import (
	"fmt"
	api "gas-platform-service/internal/adapters/api/gas-platform-service"
	"gas-platform-service/internal/adapters/repository/mongoDB"
	"gas-platform-service/internal/core/helper"
	services "gas-platform-service/internal/core/services/gas-platform-service"
	"gas-platform-service/internal/core/shared"
	"gas-platform-service/internal/ports"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func main() {
	helper.InitializeLogDir()
	service_address, service_port, dbtype, mongodb_port, _, mongodb_DBHost, dbName, _ := helper.LoadConfig()
	mongoURL := fmt.Sprintf("%s://%s:%s", dbtype, mongodb_DBHost, mongodb_port)
	DBRepository := ConnectToMongo(mongoURL, dbName)
	platformService := services.New(DBRepository)

	// Managing the handler and routes
	platformHandler := api.NewHTTPHandler(platformService)
	router := gin.Default()
	router.Use(helper.LogRequest)

	// Category routes
	router.POST("/:api-gate-way/platform/categories", platformHandler.CreateCategory())
	router.PUT("/:api-gate-way/platform/categories/reference/:category-reference", platformHandler.UpdateCategory())
	router.GET("/:api-gate-way/platform/categories/reference/:category-reference", platformHandler.GetCategoryByReference())
	router.GET("/:api-gate-way/platform/categories/name/:name", platformHandler.GetCategoryByName())
	router.GET("/:api-gate-way/platform/categories/list/page/:page", platformHandler.GetCategoryList())
	router.DELETE("/:api-gate-way/platform/categories/reference/:category-reference", platformHandler.DeleteCategoryByReference())

	// SubCategory routes
	router.POST("/:api-gate-way/platform/sub-categories", platformHandler.CreateSubCategory())
	router.PUT("/:api-gate-way/platform/sub-categories/:sub-category-reference", platformHandler.UpdateSubCategory())
	router.GET("/:api-gate-way/platform/sub-categories/reference/:sub-category-reference", platformHandler.GetSubCategoryByReference())
	router.GET("/:api-gate-way/platform/sub-categories/name/:name", platformHandler.GetSubCategoryByName())
	router.GET("/:api-gate-way/platform/sub-categories/list/page/:page", platformHandler.GetSubCategoryList())
	router.DELETE("/:api-gate-way/platform/sub-categories/reference/:sub-category-reference", platformHandler.DeleteSubCategoryByReference())

	// State routes
	router.POST("/:api-gate-way/platform/states", platformHandler.CreateState())
	router.PUT("/:api-gate-way/platform/states/:state-reference", platformHandler.UpdateState())
	router.GET("/:api-gate-way/platform/states/:page", platformHandler.GetStateList())
	router.DELETE("/:api-gate-way/platform/states/:state-reference", platformHandler.DeleteStateByReference())

	// No route
	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404, helper.PrintErrorMessage("404", shared.NORESOURCEFOUND))
	})

	fmt.Println("service running on" + service_address + ":" + service_port)
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
