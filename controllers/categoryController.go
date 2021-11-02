package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	. "vatansoft-golang-case/database"
	. "vatansoft-golang-case/models"
)

func GetAllCategories(c echo.Context) error {
	var categories []Category
	DB.Find(&categories)
	return c.JSON(http.StatusOK, categories)
}

func GetCategoryById(c echo.Context) error {
	var category Category
	if err := DB.First(&category, c.Param("id")).Error; err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, "error")
	}
	return c.JSON(http.StatusOK, &category)
}

func SaveCategory(c echo.Context) error {
	var category Category
	if err := c.Bind(&category); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, "error")
	}
	if err := DB.Create(&category).Error; err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, "error")
	}
	return c.JSON(http.StatusCreated, &category)
}

func DeleteCategory(c echo.Context) error {
	var category Category
	if err := DB.First(&category, c.Param("id")).Error; err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, "error")
	}
	if err := DB.Delete(&category).Error; err != nil {
		fmt.Println(err)
		return c.JSON(
			http.StatusBadRequest,
			map[string]string{"message": "error: Cannot delete, there are products in category."},
		)
	}
	return c.JSON(http.StatusOK, "ok")
}

func UpdateCategory(c echo.Context) error {
	var category Category
	var newCategory Category
	if err := c.Bind(&newCategory); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, "error")
	}
	DB.First(&category, c.Param("id")).Updates(newCategory)
	return c.JSON(http.StatusOK, &category)
}