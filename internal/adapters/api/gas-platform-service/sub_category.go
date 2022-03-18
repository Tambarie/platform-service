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

func (h *HTTPHandler) CreateSubCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		platform := &domain.SubCategory{}
		platform.Reference = uuid.New().String()

		err := ctx.BindJSON(&platform)
		if err != nil {
			log.Println(err)
		}

		dbplatform, err := h.platformService.CreateSubCategory(platform)
		if err != nil {
			ctx.AbortWithStatusJSON(500, err)
			return
		}

		ctx.JSON(201, gin.H{
			"reference": dbplatform.Reference,
		})
	}
}

func (h *HTTPHandler) UpdateSubCategory() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		subCategoryReference := ctx.Param("sub-category-reference")
		platform := &domain.SubCategory{}
		err := ctx.BindJSON(platform)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "could not bind json")
			return
		}

		// checking if category reference exists in the database
		reference, err := h.platformService.GetSubCategoryByReference(subCategoryReference)
		if err != nil {
			ctx.AbortWithStatusJSON(500, err)
			return
		}
		if len(reference) != 0 {
			_, err = h.platformService.UpdateSubCategory(subCategoryReference, platform)
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

func (h *HTTPHandler) GetSubCategoryByReference() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reference := ctx.Param("sub-category-reference")
		platform, err := h.platformService.GetSubCategoryByReference(reference)
		if err != nil {
			ctx.AbortWithStatusJSON(500, err)
			return
		}
		ctx.JSON(200, gin.H{
			"platform": platform,
		})

	}
}

func (h *HTTPHandler) GetSubCategoryByName() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		categoryName := ctx.Param("name")
		platform, err := h.platformService.GetSubCategoryByName(categoryName)
		if err != nil {
			ctx.AbortWithStatusJSON(500, err)
			return
		}
		ctx.JSON(200, gin.H{
			"platform": platform,
		})
	}
}

func (h *HTTPHandler) GetSubCategoryList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page := ctx.Param("page")

		pageNumber, err := strconv.Atoi(page)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		platform, err, count := h.platformService.GetSubCategoryList(int64(pageNumber))
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

func (h *HTTPHandler) DeleteSubCategoryByReference() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		SubCategoryReference := ctx.Param("sub-category-reference")

		// checking if sub-category reference exists in the database
		reference, err := h.platformService.GetSubCategoryByReference(SubCategoryReference)
		if err != nil {
			ctx.AbortWithStatusJSON(500, err)
			return
		}
		if len(reference) != 0 {
			platform, err := h.platformService.DeleteSubCategoryByReference(SubCategoryReference)
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
