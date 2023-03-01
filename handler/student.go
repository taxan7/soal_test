package handler

import (
	"encoding/json"
	"log"
	"net/http"

	model "test_sat/model"

	repo "test_sat/repository/student"

	"github.com/gin-gonic/gin"
)

func GetStudent(c *gin.Context) {
	var body model.Student
	err := json.NewDecoder(c.Request.Body).Decode(&body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.View{
			Code:         "400",
			Message:      "Error",
			ErrorMessage: "bad request",
		})
		return
	}
	repo.Insert(c.Request.Context(), body)
	c.JSON(http.StatusOK, model.View{
		Status: true,
		Data:   "",
	})
}

func CreateStudent(c *gin.Context) {
	var body model.Student
	err := json.NewDecoder(c.Request.Body).Decode(&body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.View{
			Code:         "500",
			ErrorMessage: "internal error",
		})
		return
	}
	err = repo.Insert(c.Request.Context(), body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.View{
			Code:         "400",
			Message:      "Error",
			ErrorMessage: "bad request",
		})
		return
	}
	c.JSON(http.StatusOK, model.View{
		Status:  true,
		Message: "success",
	})

}

func FindStudent(c *gin.Context) {

	data, err := repo.Read(c.Request.Context())
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.View{
			Code:         "400",
			Message:      "Error",
			ErrorMessage: "bad request",
		})
		return
	}

	var res model.View
	res.Message = "data found"
	res.Data = data
	if data == nil {
		res.Message = "data not found"
	}

	c.JSON(http.StatusOK, res)
}

func FindStudentDetail(c *gin.Context) {
	id := c.Param("id")
	data, err := repo.ReadDetail(c.Request.Context(), id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.View{
			Code:         "400",
			Message:      "Error",
			ErrorMessage: "bad request",
		})
		return
	}

	var res model.View
	res.Message = "data found"
	res.Data = data
	if data == nil {
		res.Message = "data not found"
	}

	c.JSON(http.StatusOK, res)
}

func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var body model.Student
	err := json.NewDecoder(c.Request.Body).Decode(&body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.View{
			Code:         "500",
			ErrorMessage: "internal error",
		})
		return
	}
	err = repo.Update(c.Request.Context(), id, body)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.View{
			Code:         "400",
			Message:      "Error",
			ErrorMessage: "bad request",
		})
		return
	}
	c.JSON(http.StatusOK, model.View{
		Status:  true,
		Message: "success",
	})

}

func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	err := repo.Delete(c.Request.Context(), id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, model.View{
			Code:         "400",
			Message:      "Error",
			ErrorMessage: "bad request",
		})
		return
	}

	c.JSON(http.StatusOK, model.View{
		Code:    "200",
		Message: "Success",
	})
}
