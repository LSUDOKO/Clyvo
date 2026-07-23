package controller

import (
	"context"
	"gin"
	"log"
	"net/http"
	"time"

	"github.com/LSUDOKO/Clyvo/models"
	"go.mongodb.org/mongo-driver/bson"
)

func HashPassword(password string) string {

}

func VerifyPassword(userPassword string, givenPassword string) (bool, string) {

}

func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		var user models.User

		if err := c.BindJSON(&user); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		ValidationErr := Validate.Struct(user)
		if ValidationErr != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": ValidationErr.Error()})
			return
		}

		count, err := UserCollection.CountDocuments(ctx, bson.M{"email": user.Email})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while checking for the email"})
			return
		}
		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "This email already exists"})
			return
		}

		count, err = UserCollection.CountDocuments(ctx, bson.M{"phone": user.Phone})
		if err != nil {
			log.Panic(err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error occurred while checking for the phone number"})
			return
		}
		if count > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "This phone number already exists"})
			return
		}
	}
}

func Login() gin.HandlerFunc {

}

func ProductViewerAdmin() gin.HandlerFunc {

}

func searchProductByQuery() gin.HandlerFunc {

}

func searchProduct() gin.HandlerFunc {

}
