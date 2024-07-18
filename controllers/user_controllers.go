package controllers

import (
	"fmt"
	"time"
	"context"
	"net/http"
	"gin-mongo-api/utils"
	"gin-mongo-api/models"
	"gin-mongo-api/configs"
	"gin-mongo-api/responses"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()

		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{
				Status: http.StatusBadRequest,
				Message: "error",
				Data: map[string]interface{}{"data":err.Error()},
			})
			return
		}

		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.UserResponse{
				Status: http.StatusBadRequest,
				Message: "error",
				Data: map[string]interface{}{"data":validationErr.Error()},
			})
			return
		}

		key := []byte("erthgrferghntrrt")
		enc_username := []byte(user.Username)
		enc_password := []byte(user.Password)

		encrypted_username, err := utils.Encrypt(key, enc_username)
		if err != nil {
			fmt.Println(err)
		}
		encrypted_password, err := utils.Encrypt(key, enc_password)
		if err != nil {
			fmt.Println(err)
		}

		newUser := models.InsUser{
			Username: user.Username,
			EncUsername: []byte(encrypted_username),
			Password: []byte(encrypted_password),
			Key: key,
			Token: user.Token,
		}

		result, err := userCollection.InsertOne(ctx, newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status: http.StatusInternalServerError,
				Message: "error",
				Data: map[string]interface{}{"data":err.Error()},
			})
			return
		}

		c.JSON(http.StatusCreated, responses.UserResponse{
			Status: http.StatusCreated,
			Message: "success",
			Data: map[string]interface{}{"data":result},
		})
	}
}

func GetAUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx , cancel := context.WithTimeout(context.Background(), 10*time.Second)
		userId := c.Param("userId")

		var user models.User
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(userId)

		err := userCollection.FindOne(ctx, bson.M{
			"id"  : objId,
		}).Decode(&user)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status: http.StatusInternalServerError,
				Message: "error",
				Data: map[string]interface{}{"data": err.Error()},

			})
		}

		c.JSON(http.StatusOK, responses.UserResponse{
			Status: http.StatusOK,
			Message: "success",
			Data: map[string]interface{}{"data":user},
		})

	}
}

func GetAllUsers() gin.HandlerFunc{
	return func(c *gin.Context) {
		ctx, cancel:= context.WithTimeout(context.Background(), 10*time.Second)
		var users []models.User
		defer cancel()

		results, err := userCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status: http.StatusInternalServerError,
				Message: "error",
				Data: map[string]interface{}{"data":err.Error()},
			})
			return
		}

		defer results.Close(ctx)
		for results.Next(ctx){
			var singleUser models.User
			if err = results.Decode(&singleUser); err != nil {
				c.JSON(http.StatusInternalServerError, responses.UserResponse{
					Status: http.StatusInternalServerError,
					Message: "error",
					Data: map[string]interface{}{"data":err.Error()},
				})

				users = append(users, singleUser)
			}

			c.JSON(http.StatusOK, responses.UserResponse{
				Status: http.StatusOK,
				Message: "success",
				Data: map[string]interface{}{"data":users},
			})
		}
	}
}