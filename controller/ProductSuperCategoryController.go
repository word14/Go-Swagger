package controller

import (
	"errors"
	"sync"

	"../database"
	"../model"
	"github.com/gin-gonic/gin"
)

func test1(wait *sync.WaitGroup) <-chan []model.ProductSuperCategory {
	channel := make(chan []model.ProductSuperCategory)
	wait.Add(1)
	go func() {
		var result []model.ProductSuperCategory
		// time.Sleep(time.Second * 1)
		database.DB.Find(&result)
		wait.Done()
		channel <- result
	}()
	return channel
}

func test2(wait *sync.WaitGroup) chan model.ProductSuperCategory {
	channel := make(chan model.ProductSuperCategory)
	wait.Add(1)
	go func() {
		var result model.ProductSuperCategory
		database.DB.Where(model.ProductSuperCategory{Slug: "food"}).First(&result)
		wait.Done()
		channel <- result
	}()
	return channel
}

// @Tags ProductSuperCategory
// @Summary Get Product Super Category
// @Description Get Product Super Category
// @Accept  json
// @Produce  json
// @Success 200 {array} model.ProductSuperCategory "Success"
// @Router /test [post]
// @Param body body model.ProductSuperCategory true "Request Body"
func ProductSuperCategoryList(context *gin.Context) {
	// var result []model.ProductSuperCategory
	var wait sync.WaitGroup
	result := test1(&wait)
	result2 := test2(&wait)
	// test2(&wait)
	// wait.Wait()
	// database.DB.Find(&result)
	context.JSON(200, gin.H{
		"result1": <-result,
		"result2": <-result2})
}

func ProductSuperCategoryGet(context *gin.Context) {
	slug := context.Param("slug")
	var result model.ProductSuperCategory
	db := database.DB.Where(&model.ProductSuperCategory{Slug: slug}).First(&result)

	if !db.RecordNotFound() {
		context.JSON(200, result)
	} else {
		context.Error(errors.New("Not Found"))
	}
}
