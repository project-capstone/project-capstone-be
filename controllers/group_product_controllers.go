package controllers

import (
	"final-project/lib/databases"
	"final-project/middlewares"
	"final-project/models"
	response "final-project/responses"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

// controller untuk menambahkan user (registrasi)
func CreateGroupProductControllers(c echo.Context) error {
	new_group := models.GroupProduct{}
	id := c.Param("id_products")
	c.Bind(&new_group)

	id_user, role := middlewares.ExtractTokenId(c)
	if role == "admin" {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Access Forbidden"))
	}

	id_product, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}

	duration := time.Now().AddDate(0, 0, 14)
	//mengambil banyaknya jumlah data yang ada pada group
	_, len_group, _ := databases.GetAllGroupProduct()
	name, price, limit, _ := databases.GetDataProduct(int(id_product))
	fee := 5000

	new_group.UsersID = uint(id_user)
	new_group.ProductsID = uint(id_product)
	new_group.NameGroupProduct = name + "-" + strconv.Itoa(len_group+1)
	new_group.CapacityGroupProduct = 0
	new_group.AdminFee = fee
	new_group.TotalPrice = (price / limit) + fee
	new_group.DurationGroup = duration.Format("02-01-2006")
	new_group.Status = "Available"

	d, er := databases.CreateGroupProduct(&new_group, new_group.ProductsID)
	if er != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	if name == "" || price == 0 {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Id Product Not Found"))
	}

	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", d))
}

func GetByIdGroupProductControllers(c echo.Context) error {
	id := c.Param("id_group")
	id_group_product, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}
	data, e := databases.GetGroupProductById(id_group_product)
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	if data == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", data))
}
func GetByIdProductsGroupProductControllers(c echo.Context) error {
	id := c.Param("id_products")
	id_products, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}
	data, e := databases.GetGroupProductByIdProducts(id_products)
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	if data == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", data))
}

// controller untuk menampilkan seluruh data users
func GetAllGroupProductControllers(c echo.Context) error {
	data, _, err := databases.GetAllGroupProduct()
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	if data == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", data))
}

func GetAvailableGroupProductControllers(c echo.Context) error {
	status := c.Param("status")
	if status != "available" {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Param"))
	}
	data, e := databases.GetGroupProductByAvailable(status)
	if e != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Bad Request"))
	}
	if data == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	return c.JSON(http.StatusOK, response.SuccessResponseData("Success Operation", data))
}

func DeleteGroupProductControllers(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id_group"))
	logged, role := middlewares.ExtractTokenId(c) // check token
	if logged != id && role != "admin" {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Access Forbidden"))
	}
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Invalid Id"))
	}
	data_order, _, _ := databases.GetOrderByIdGroup(id)
	if data_order != nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Access is denied ID data is in the order"))
	}
	data, _ := databases.GetGroupProductById(id)
	if data == nil {
		return c.JSON(http.StatusBadRequest, response.BadRequestResponse("Data Not Found"))
	}
	databases.DeleteGroupProduct(id)

	return c.JSON(http.StatusOK, response.SuccessResponseNonData("Success Operation"))
}
