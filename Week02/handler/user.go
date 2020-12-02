package handler

import (
	"database/sql"
	"errors"
	"github.com/EndlessIdea/Go-000/Week02/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserInfo(c *gin.Context) {
	paramID := c.Param("id")
	userID, err := strconv.Atoi(paramID)
	if err != nil {
		log.Printf("invalid user id %s", paramID)
		c.JSON(http.StatusBadRequest, gin.H{"err_msg": "Invalid user id"})
		return
	}

	user, err := service.UserService.GetUserByID(c, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			log.Printf("Non-existent user %d: %+v", userID, err)
			c.JSON(http.StatusNotFound, gin.H{"err_msg": "Non-existent user"})
			return
		}
		log.Printf("failed to get user info %d: %+v", userID, err)
		c.JSON(http.StatusInternalServerError, gin.H{"err_msg": "service error"})
		return
	}

	c.JSON(http.StatusOK, user)
}
