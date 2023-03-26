package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Book struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var BookDatas = []Book{}

func CreateBook(ctx *gin.Context) {
	var book Book

	if err := ctx.ShouldBindJSON(&book); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var id int
	if len(BookDatas) > 0 {
		last := BookDatas[len(BookDatas)-1]
		id, _ = strconv.Atoi(last.Id)
	}
	book.Id = fmt.Sprintf("%d", id+1)
	BookDatas = append(BookDatas, book)

	ctx.JSON(http.StatusCreated, gin.H{
		"book": book,
	})
}

func UpdateBook(ctx *gin.Context) {
	id := ctx.Param("id")
	condition := false
	var updatedBook Book

	if err := ctx.ShouldBindJSON(&updatedBook); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	updatedBook.Id = id
	for i, book := range BookDatas {
		if book.Id == id {
			condition = true
			BookDatas[i] = updatedBook
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": fmt.Sprintf("Book with id %s Not Found", id),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Updated",
		"dfd":     "dhgkd",
	})
}

func GetBook(ctx *gin.Context) {
	id := ctx.Param("id")
	condition := false
	var bookData Book

	for i, book := range BookDatas {
		if book.Id == id {
			condition = true
			BookDatas[i] = bookData
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": fmt.Sprintf("Book with id %s Not Found", id),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": bookData,
	})

}

func DeleteBook(ctx *gin.Context) {
	id := ctx.Param(("id"))
	condition := false
	var index int

	for i, book := range BookDatas {
		if book.Id == id {
			condition = true
			index = i
			break
		}
	}

	if !condition {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": fmt.Sprintf("Book with id %s Not Found", id),
		})
		return
	}

	copy(BookDatas[index:], BookDatas[index+1:])
	BookDatas[len(BookDatas)-1] = Book{}
	BookDatas = BookDatas[:len(BookDatas)-1]

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Delete Success",
	})

}

func GetAllBook(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"data": BookDatas,
	})
}
