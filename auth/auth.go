package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ipreferwater/doko-graphql/db"
)

type inputLogin struct {
	Username     string `json:"username"`
	Password string `json:"password"`
}

func Login(c *gin.Context) {

	var input inputLogin
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't parse the payload to requiered model", "log": err.Error()})
		return
	}

	id, err := db.PostRepository.GetUserIdByUsernamePassword(input.Username, input.Password)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "can't check credentials", "log": err.Error()})
		return
	}

	if id == 0 {
		c.JSON(http.StatusNotFound, gin.H{"response": "credentials not found"})
		return
	}

	tokenString, err := CreateToken(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token": tokenString,
		"token_type":   "Bearer",
	})
}
