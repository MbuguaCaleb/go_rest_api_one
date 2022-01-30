package routes

import (
	"errors"

	"github.com/MbuguaCaleb/go_rest_api_one/database"
	"github.com/MbuguaCaleb/go_rest_api_one/models"
	"github.com/gofiber/fiber/v2"
)

type ProductSerializer struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	SerialNumber string `json:"serial_number"`
}

//handling responses from the Product
func CreateResponseProduct(productModel models.Product) ProductSerializer{
	return ProductSerializer{ID:productModel.ID, Name: productModel.Name,SerialNumber: productModel.SerialNumber}
}

//create Product
func CreateProduct(c *fiber.Ctx) error {
	var product models.Product

	if err := c.BodyParser(&product); err != nil{
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&product)
	responseProduct := CreateResponseProduct(product)
	return c.Status(200).JSON(responseProduct)

}


//Get Products
func GetProducts(c *fiber.Ctx) error  {
	products :=  []models.Product{}
	database.Database.Db.Find(&products)
	
	responseProducts := []ProductSerializer{}
    
	for _,product := range products{
		responseProduct:= CreateResponseProduct(product)
		responseProducts = append(responseProducts,responseProduct)
	}

	return c.Status(200).JSON(responseProducts)	
}

//find Product
func findProduct (id int, product *models.Product) error{
	database.Database.Db.Find(&product,"id = ?",id)
	if product.ID == 0 {
		return errors.New("product does not exist")
	}
	return nil

}

//Get Product
func GetProduct(c *fiber.Ctx) error{

	//because i am passing an Integer Parameter
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil{
		return c.Status(400).JSON("Please ensure that :id is an Integer")
	}

	if err := findProduct(id,&product); err!=nil{
		return c.Status(400).JSON(err.Error())
	}

	responseProduct :=CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)

}


func UpdateProduct (c *fiber.Ctx) error{
	
	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil{
		return c.Status(400).JSON("Please ensure that :id is an Integer")
	}

	if err := findProduct(id,&product); err!=nil{
		return c.Status(400).JSON(err.Error())
	}

	//Update User Struct
	type UpdateProduct struct {
		Name string `json:"name"`
		SerialNumber string `json:"serial_number"`
	}

	var updateData UpdateProduct

	//c handles the request Body
	if err :=c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	product.Name = updateData.Name
	product.SerialNumber = updateData.SerialNumber

	database.Database.Db.Save(&product)

	responseProduct := CreateResponseProduct(product)

	return c.Status(200).JSON(responseProduct)
	
}

func DeleteProduct(c *fiber.Ctx) error{

	id, err := c.ParamsInt("id")

	var product models.Product

	if err != nil{
		return c.Status(400).JSON("Please ensure that :id is an Integer")
	}

	if err := findProduct(id,&product); err!=nil{
		return c.Status(400).JSON(err.Error())
	}


	if err := database.Database.Db.Delete(&product).Error; err != nil {
		return c.Status(404).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfully Deleted Product")

}

