package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	model "test_sat/model"
	"test_sat/pkg/auth"
	userRepo "test_sat/repository/user"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Login(c *gin.Context) {
	var body model.User
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
	if err := validator.New().Struct(body); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, model.View{
			Code:         "400",
			Message:      "Error",
			ErrorMessage: "Bad Request",
		})
		return
	}

	data, err := userRepo.ReadUserPass(c.Request.Context(), body)
	if err != nil {
		log.Println(err)

		c.JSON(http.StatusInternalServerError, model.View{
			Message: "failed to  login",
		})

		return
	}

	if data == nil {
		c.JSON(http.StatusBadRequest, model.View{
			Message:      "Error",
			ErrorMessage: "wrong email/password",
		})
		return
	}

	token, _ := auth.GenerateJWT(data.Username)

	c.JSON(http.StatusOK, model.View{
		Status: true,
		Data:   token,
	})
}

func Logout(c *gin.Context) {
	token := strings.Replace(c.GetHeader("Authorization"), "Bearer ", "", 1)
	auth.Logout(token)
	c.JSON(http.StatusOK, model.View{
		Status: true,
	})
}
