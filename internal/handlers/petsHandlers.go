package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/zepyrshut/pet-clinic/internal/models"
)

func (m *Repository) InsertNewPet(c *gin.Context) {
	pet := models.Pet{}

	err := c.ShouldBind(&pet)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "malformed request",
			"error":  err.Error(),
		})
		return
	}

	err = m.DB.NewPet(pet)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "created",
		"pet":    pet,
	})
}

func (m *Repository) GetOnePet(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "malformed id",
			"error":  err.Error(),
		})
		return
	}

	pet, err := m.DB.OnePet(id)
	if err != nil {
		m.App.InfoLog.Println(err)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status": "not found",
			"error":  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"pet":    pet,
	})
}
