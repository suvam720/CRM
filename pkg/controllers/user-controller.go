package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/suvam720/crud-api/pkg/database"
	"github.com/suvam720/crud-api/pkg/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func FindUser(c *gin.Context) {
	var users []primitive.M
	cur, err := database.Collection.Find(c, bson.D{})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	for cur.Next(context.Background()) {
		var user bson.M
		if err := cur.Decode(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		users = append(users, user)
	}
	defer cur.Close(c)
	c.JSON(http.StatusOK, gin.H{"users": users})
}

func CreateUser(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	res, err := database.Collection.InsertOne(c, &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{"userId": res.InsertedID})
}

func DeleteUser(c *gin.Context) {
	Id := c.Param("id")
	fmt.Println(Id)
	id, err := primitive.ObjectIDFromHex(Id)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": id}

	res, err := database.Collection.DeleteOne(c, filter)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": res.DeletedCount})
}

func DeleteAllUser(c *gin.Context) {
	res, err := database.Collection.DeleteMany(c, bson.D{{}})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"data": res})
}

func UpdateUser(c *gin.Context) {
	var inputUser models.User
	Id := c.Param("id")
	id, _ := primitive.ObjectIDFromHex(Id)
	filter := bson.M{"_id": id}

	update := bson.M{"$set": bson.M{
		"name":    &inputUser.Name,
		"age":     &inputUser.Age,
		"gender":  &inputUser.Gender,
		"address": &inputUser.Address,
	}}

	if err := c.ShouldBindJSON(&inputUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		return
	}

	res, err := database.Collection.UpdateOne(
		c,
		filter,
		update,
	)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	fmt.Println(res)

	c.JSON(http.StatusOK, gin.H{"data": res.ModifiedCount})
}
