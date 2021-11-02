package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	. "vatansoft-golang-case/database"
	. "vatansoft-golang-case/models"
)

func GetAllProperties(c echo.Context) error {
	var properties []Property
	DB.Find(&properties)
	return c.JSON(http.StatusOK, properties)
}

func GetPropertyById(c echo.Context) error {
	var property Property
	if err := DB.First(&property, c.Param("id")).Error; err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, "error")
	}
	return c.JSON(http.StatusOK, &property)
}

func SaveProperty(c echo.Context) error {
	var property Property
	if err := c.Bind(&property); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, "error")
	}
	if err := DB.Create(&property).Error; err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, "error")
	}
	return c.JSON(http.StatusCreated, &property)
}

func DeleteProperty(c echo.Context) error {
	var property Property
	if err := DB.First(&property, c.Param("id")).Error; err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, "error")
	}
	if err := DB.Delete(&property).Error; err != nil {
		fmt.Println(err)
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{"message": "error: Cannot delete, there are products with this property."},
		)
	}
	return c.JSON(http.StatusOK, "ok")
}

func UpdateProperty(c echo.Context) error {
	var property Property
	var newProperty Property
	if err := c.Bind(&newProperty); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, "error")
	}
	DB.First(&property, c.Param("id")).Updates(newProperty)
	return c.JSON(http.StatusOK, &property)
}

