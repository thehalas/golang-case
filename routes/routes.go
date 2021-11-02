package routes

import (
	"github.com/labstack/echo/v4"
	. "vatansoft-golang-case/controllers"
)

func Setup(e *echo.Echo) {
	// product
	e.GET("/stocks", GetAllProductsNotDeleted)
	e.GET("/stocks/filter", nil)
	e.POST("/stock/insert", SaveProduct)
	e.PUT("/stock/:id/update", UpdateProduct)
	e.DELETE("/stock/:id/delete", DeleteProduct)
	e.GET("/stock/:id", GetProductById)
	// Category
	e.GET("/categories", GetAllCategories)
	e.GET("/category/:id", GetCategoryById)
	e.POST("/category/insert", SaveCategory)
	e.DELETE("/category/:id/delete", DeleteCategory)
	e.PUT("/category/:id/update", UpdateCategory)
	// Property
	e.GET("/properties", GetAllProperties)
	e.GET("/property/:id", GetPropertyById)
	e.POST("/property/insert", SaveProperty)
	e.DELETE("/property/:id/delete", DeleteProperty)
	e.PUT("/property/:id/update", UpdateProperty)
	// Invoice
	e.GET("/invoices", GetAllInvoices)
	e.GET("/invoice/:id", GetInvoiceById)
	e.POST("/invoice/insert", SaveInvoice)
	e.DELETE("/invoice/:id/delete", DeleteInvoice)
	e.PUT("/invoice/:id/update", UpdateInvoice)

}
