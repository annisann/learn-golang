package controllers

import (
	"08_rest_api_project/structs"
	"net/http"

	"github.com/gin-gonic/gin"
)

// get data by `id`
func (idb *InDB) GetPerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	id := c.Param("id")
	err := idb.DB.Where("id=?", id).First(&person).Error
	if err != nil {
		result = gin.H{
			"result": err.Error(),
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": person,
			"count":  1,
		}
	}
	c.JSON(http.StatusOK, result)
}

// get all data
func (idb *InDB) GetPersons(c *gin.Context) {
	var (
		persons []structs.Person
		result  gin.H
	)

	idb.DB.Find(&persons)
	if len(persons) <= 0 {
		result = gin.H{
			"result": nil,
			"count":  0,
		}
	} else {
		result = gin.H{
			"result": persons,
			"count":  len(persons),
		}
	}

	c.JSON(http.StatusOK, result)
}

// add new data to database
func (idb *InDB) CreatePerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")
	person.FirstName = first_name
	person.LastName = last_name

	idb.DB.Create(&person)
	result = gin.H{
		"result": person,
	}

	c.JSON(http.StatusOK, result)
}

// update data by `id`
func (idb *InDB) UpdatePerson(c *gin.Context) {
	id := c.Query("id")

	first_name := c.PostForm("first_name")
	last_name := c.PostForm("last_name")

	var (
		person    structs.Person
		newPerson structs.Person
		result    gin.H
	)

	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "Data not found",
		}
	}

	newPerson.FirstName = first_name
	newPerson.LastName = last_name

	err = idb.DB.Model(&person).Updates(newPerson).Error
	if err != nil {
		result = gin.H{
			"result": "Update failed",
		}
	} else {
		result = gin.H{
			"result": "Successfully updating data.",
		}
	}

	c.JSON(http.StatusOK, result)
}

// delete data by `id`
func (idb *InDB) DeletePerson(c *gin.Context) {
	var (
		person structs.Person
		result gin.H
	)

	id := c.Param("id")
	err := idb.DB.First(&person, id).Error
	if err != nil {
		result = gin.H{
			"result": "Failed deleting data",
		}
	} else {
		result = gin.H{
			"result": "Successfully deleting data",
		}
	}

	c.JSON(http.StatusOK, result)
}
