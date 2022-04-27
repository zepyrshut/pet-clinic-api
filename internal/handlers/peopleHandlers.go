package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/zepyrshut/pet-clinic/internal/models"
)

func (m *Repository) InsertNewPerson(c *gin.Context) {
	person := models.Person{}

	err := c.ShouldBind(&person)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "malformed request",
			"error":  err.Error(),
		})
		return
	}

	err = m.DB.NewPerson(person)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status": "created",
		"person": person,
	})
}

func (m *Repository) BindPetWithOwner(c *gin.Context) {
	err := c.Request.ParseForm()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "malformed request",
			"error":  err.Error(),
		})
		return
	}

	petID, errPet := strconv.Atoi(c.Request.FormValue("pet_id"))
	personID, errPer := strconv.Atoi(c.Request.FormValue("person_id"))

	m.App.InfoLog.Println(petID, personID)

	if errPet != nil || errPer != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "malformed id",
			"error":  err.Error(),
		})
		return
	}

	err = m.DB.BindPetWithOwner(petID, personID)
	if err != nil {
		c.Error(err)
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
	})
}

func (m *Repository) GetOnePerson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status": "malformed id",
			"error":  err.Error(),
		})
		return
	}

	person, err := m.DB.OnePerson(id)
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
		"person": person,
	})
}
