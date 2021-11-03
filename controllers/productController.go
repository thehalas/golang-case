package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	. "vatansoft-golang-case/database"
	. "vatansoft-golang-case/models"
)

func GetAllProductsNotDeleted(c echo.Context) error {
	var products []Product
	DB.
		Preload("ProductProperties.Property").
		Preload("Category").
		Where("is_deleted = 0").
		Find(&products)
	return c.JSON(http.StatusOK, &products)
}

func GetProductById(c echo.Context) error {
	var product Product
	if err := DB.
		Preload("ProductProperties.Property").
		Preload("Category").
		First(&product, c.Param("id")).Error; err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, "error")
	}
	return c.JSON(http.StatusOK, &product)
}

func SaveProduct(c echo.Context) error {
	var product Product
	if err := c.Bind(&product); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, "error")
	}
	if err := DB.Create(&product).Error; err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, "error")
	}
	if err := DB.Model(&product).
		Association("Category").
		Find(&product.Category); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, "error")
	}
	if err := DB.Model(&product).
		Preload("Property").
		Association("ProductProperties").
		Find(&product.ProductProperties); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, "error")
	}
	return c.JSON(http.StatusCreated, &product)
}

func UpdateProduct(c echo.Context) error {
	var product Product
	var newProduct Product
	if err := c.Bind(&newProduct); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, "error")
	}
	DB.Where("product_id = ?", c.Param("id")).Delete(&ProductProperty{})
	DB.First(&product, c.Param("id")).Updates(newProduct)
	if err := DB.Model(&product).Association("Category").Find(&product.Category); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, "error")
	}
	if err := DB.Model(&product).Association("ProductProperties").Append(newProduct.ProductProperties); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusInternalServerError, "error")
	}
	if err := DB.Model(&product).
		Preload("Property").
		Association("ProductProperties").
		Find(&product.ProductProperties); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, "error")
	}
	return c.JSON(http.StatusOK, &product)
}

func DeleteProduct(c echo.Context) error {
	var product Product
	if err := DB.First(&product, c.Param("id")).
		Update("is_deleted", true).Error; err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, "error")
	}
	return c.JSON(http.StatusOK, "ok")
}
