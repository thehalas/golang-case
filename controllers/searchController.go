package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	. "vatansoft-golang-case/database"
	. "vatansoft-golang-case/models"
)

func FilterProducts(c echo.Context) error {
	var products []Product
	hideDeleted := c.QueryParam("showDeleted") != "true"
	hideSoldOut := c.QueryParam("showSoldOut") != "true"
	props := c.QueryParams()["prop"]
	category := c.QueryParam("category")
	base := DB.
		Preload("ProductProperties.Property").
		Preload("Category").
		Distinct("products.id, products.name, products.count, products.category_id, is_deleted")

	if category != "" {
		base = base.Where("category = ?", category)
	}
	if hideDeleted {
		base = base.Where("is_deleted = 0")
	}
	if hideSoldOut {
		base = base.Where("count > 0")
	}
	if len(props) > 0 {
		base = base.Joins("JOIN product_properties ON products.id = product_properties.product_id")
		ors := DB
		flag := true
		for _, kv := range props {
			if i := strings.Index(kv, "|"); i > -1 {
				propId := kv[:i]
				propValue := kv[i+1:]
				fmt.Println(propId, propValue)
				if flag {
					ors = base.Where("property_id = ? AND value = ?", propId, propValue)
					flag = false
				} else {
					ors = base.Or("property_id = ? AND value = ?", propId, propValue)

				}
			}
		}
		base = base.Where(ors)
	}
	base.Find(&products)
	return c.JSON(http.StatusOK, products)
}
