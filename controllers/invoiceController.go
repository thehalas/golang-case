package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	. "vatansoft-golang-case/database"
	. "vatansoft-golang-case/models"
)

func GetAllInvoices(c echo.Context) error {
	var invoices []Invoice
	DB.Find(&invoices)
	return c.JSON(http.StatusOK, invoices)
}

func GetInvoiceById(c echo.Context) error {
	var invoice Invoice
	if err := DB.First(&invoice, c.Param("id")).Error; err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, "error")
	}
	return c.JSON(http.StatusOK, &invoice)
}

func SaveInvoice(c echo.Context) error {
	var invoice Invoice
	if err := c.Bind(&invoice); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, "error")
	}
	if err := DB.Create(&invoice).Error; err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, "error")
	}
	return c.JSON(http.StatusCreated, &invoice)
}

func DeleteInvoice(c echo.Context) error {
	var invoice Invoice
	if err := DB.First(&invoice, c.Param("id")).Error; err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusNotFound, "error")
	}
	if err := DB.Delete(&invoice).Error; err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, "error",
		)
	}
	return c.JSON(http.StatusOK, "ok")
}

func UpdateInvoice(c echo.Context) error {
	var invoice Invoice
	var newInvoice Invoice
	if err := c.Bind(&newInvoice); err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, "error")
	}
	DB.First(&invoice, c.Param("id")).Updates(newInvoice)
	return c.JSON(http.StatusOK, &invoice)
}
