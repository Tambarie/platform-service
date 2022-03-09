package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	domain "platform-service/internal/core/domain/gas-platform-service"
	"platform-service/internal/core/helper"
	"platform-service/internal/core/shared"
	"strconv"
)

func (h *HTTPHandler) CreateState() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		state := &domain.State{}
		state.Reference = uuid.New().String()

		err := ctx.BindJSON(&state)
		if err != nil {
			log.Println(err)
		}

		dbState, err := h.platformService.CreateState(state)
		if err != nil {
			ctx.AbortWithStatusJSON(500, err)
			return
		}

		ctx.JSON(201, gin.H{
			"reference": dbState.Reference,
		})
	}
}

func (h *HTTPHandler) UpdateState() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		stateReference := ctx.Param("state-reference")
		state := &domain.State{}
		err := ctx.BindJSON(state)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, "could not bind json")
			return
		}

		// checking if category reference exists in the database
		reference, err := h.platformService.GetStateByReference(stateReference)
		if err != nil {
			ctx.JSON(404, helper.PrintErrorMessage("404", shared.REQUEST_NOT_FOUND))
			return
		}
		if len(reference) != 0 {
			_, err = h.platformService.UpdateState(stateReference, state)
			if err != nil {
				ctx.JSON(404, helper.PrintErrorMessage("404", shared.REQUEST_NOT_FOUND))
				return
			}
			ctx.JSON(200, nil)
		} else {
			ctx.JSON(404, helper.PrintErrorMessage("404", shared.REQUEST_NOT_FOUND))
		}

	}
}

func (h *HTTPHandler) GetStateList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		page := ctx.Param("page")

		pageNumber, err := strconv.Atoi(page)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		state, err, count := h.platformService.GetStateList(int64(pageNumber))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		ctx.JSON(200, gin.H{
			"platform": state,
			"count":    count,
		})
	}
}

func (h *HTTPHandler) DeleteStateByReference() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		categoryReference := ctx.Param("state-reference")

		// checking if category reference exists in the database
		reference, err := h.platformService.GetStateByReference(categoryReference)
		log.Println(reference)
		if err != nil {
			ctx.AbortWithStatusJSON(500, err)
			return
		}
		if len(reference) != 0 {
			state, err := h.platformService.DeleteStateByReference(categoryReference)
			if err != nil {
				ctx.AbortWithStatusJSON(500, err)
				return
			}
			ctx.JSON(200, gin.H{
				"reference": state,
			})
		} else {
			ctx.JSON(404, helper.PrintErrorMessage("404", shared.REQUEST_NOT_FOUND))
		}

	}
}
