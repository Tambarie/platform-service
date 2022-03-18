package api

import (
	"fmt"
	domain "gas-platform-service/internal/core/domain/gas-platform-service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"strconv"
)

func (h *HTTPHandler) CreateCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		platform := &domain.Category{}
		platform.Reference = uuid.New().String()

		err := ctx.BindJSON(&platform)
		if err != nil {
			log.Println(err)
		}

		dbplatform, err := h.platformService.CreateCategory(platform)
		if err != nil {
			ctx.AbortWithStatusJSON(500, err)
			return
		}

		ctx.JSON(201, gin.H{
			"reference": dbplatform.Reference,
		})
	}
}

func (h *HTTPHandler) UpdateCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		categoryReference := ctx.Param("category-reference")
		platform := &domain.Category{}
		err := ctx.BindJSON(platform)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "could not bind json")
			return
		}

		// checking if category reference exists in the database
		reference, err := h.platformService.GetCategoryByReference(categoryReference)
		if err != nil {
			ctx.AbortWithStatusJSON(500, err)
			return
		}
		if len(reference) != 0 {
			_, err = h.platformService.UpdateCategory(categoryReference, platform)
			if err != nil {
				ctx.AbortWithStatusJSON(500, err)
				return
			}
			ctx.JSON(200, nil)
		} else {
			ctx.AbortWithStatusJSON(500, err)
		}

	}
}

func (h *HTTPHandler) GetCategoryByReference() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reference := ctx.Param("category-reference")
		platform, err := h.platformService.GetCategoryByReference(reference)
		if err != nil {
			ctx.AbortWithStatusJSON(500, err)
			return
		}
		ctx.JSON(200, gin.H{
			"platform": platform,
		})

	}
}

func (h *HTTPHandler) GetCategoryByName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		categoryName := ctx.Param("name")
		platform, err := h.platformService.GetCategoryByName(categoryName)
		if err != nil {
			ctx.AbortWithStatusJSON(500, err)
			return
		}
		ctx.JSON(200, gin.H{
			"platform": platform,
		})

	}
}

func (h *HTTPHandler) GetCategoryList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page := ctx.Param("page")

		pageNumber, err := strconv.Atoi(page)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		platform, err, count := h.platformService.GetCategoryList(int64(pageNumber))
		fmt.Println(count)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		ctx.JSON(200, gin.H{
			"platform": platform,
			"count":    count,
		})
	}
}

func (h *HTTPHandler) DeleteCategoryByReference() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		categoryReference := ctx.Param("category-reference")

		// checking if category reference exists in the database
		reference, err := h.platformService.GetCategoryByReference(categoryReference)
		if err != nil {
			ctx.AbortWithStatusJSON(500, err)
			return
		}
		if len(reference) != 0 {
			platform, err := h.platformService.DeleteCategoryByReference(categoryReference)
			if err != nil {
				ctx.AbortWithStatusJSON(500, err)
				return
			}
			ctx.JSON(200, gin.H{
				"reference": platform,
			})
		} else {
			ctx.AbortWithStatusJSON(500, err)
		}

	}
}
