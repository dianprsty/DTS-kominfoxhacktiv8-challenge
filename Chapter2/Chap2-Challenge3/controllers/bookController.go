package controllers

import (
	"Chap2-Challenge3/models"
	"Chap2-Challenge3/repository"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

var BookDatas = []models.Book{}

func CreateBook(ctx *gin.Context) {
	var book models.Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	b, err := repository.InsertBook(book)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"book": b,
	})
}

func UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	// condition := false
	var updatedBook models.Book

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err := repository.Update(updatedBook, id)

	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Updated success",
		"id":      id,
		"data":    updatedBook,
	})
}

func GetBook(ctx *gin.Context) {
	id := ctx.Param("id")

	book, err := repository.GetById(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": fmt.Sprintf("Book with id %s Not Found", id),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": book,
	})

}

func DeleteBook(ctx *gin.Context) {
	id := ctx.Param(("id"))

	err := repository.Delete(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": fmt.Sprintf("Book with id %s Not Found", id),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete Success",
	})

}

func GetAllBook(ctx *gin.Context) {

	books, err := repository.GetAll()

	if err != nil {
		fmt.Println(err)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": books,
	})
}
