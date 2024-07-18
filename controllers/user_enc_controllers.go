package controllers

import (
	"io"
	"fmt"
	"time"
	"bytes"
	"context"
	"net/http"
	"encoding/json"
	"encoding/base64"
	"gin-mongo-api/models"
	"gin-mongo-api/configs"
	"gin-mongo-api/responses"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	//"github.com/go-playground/validator/v10"
)

var userCol *mongo.Collection = configs.GetCollection(configs.DB, "users")
// var valid = validator.New()

func GetCvesFrom() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		_, VI_URI := configs.EnvMongoURI()

		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			fmt.Println(err)
		}
		defer c.Request.Body.Close()

		JSONpayload := make(map[string]interface{})

		if err := json.Unmarshal(body, &JSONpayload); err != nil {
			fmt.Println(err)
		}

		var user models.InsUser

		payload := map[string]string{
			"vendor":  JSONpayload["vendor"].(string),
			"product": JSONpayload["product"].(string),
			"version": JSONpayload["version"].(string),
		}

		marshalPayload, _ := json.Marshal(payload)

		filter := bson.M{
			"username" : "harshit",
		}

		userCol.FindOne(ctx, filter).Decode(&user)

		req, err := http.NewRequest(http.MethodPost, VI_URI+"/api/v1/fetchproductdata", bytes.NewReader(marshalPayload))
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status: http.StatusInternalServerError,
				Message: "error",
				Data: map[string]interface{}{"data":err.Error()},
			})
			return
		}

		encodedUsername := base64.StdEncoding.EncodeToString(user.EncUsername)
		encodedPassword := base64.StdEncoding.EncodeToString(user.Password)

		req.Header.Add("username", encodedUsername)
		req.Header.Add("password", encodedPassword)
		req.Header.Add("key", string(user.Key))
		req.Header.Add("token", user.Token)

		client := &http.Client{}

		response, err := client.Do(req)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.UserResponse{
				Status: http.StatusInternalServerError,
				Message: "error",
				Data: map[string]interface{}{"data":err.Error()},
			})
			return
		}

		var softVulDetail models.ResponsesofVulDetails

		resbody, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}

		if err = json.Unmarshal([]byte(resbody), &softVulDetail); err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, responses.UserResponse{
			Status: http.StatusOK,
			Message: "success",
			Data: map[string]interface{}{"data": softVulDetail},
		})
	}
}
