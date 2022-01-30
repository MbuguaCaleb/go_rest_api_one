package routes

import (
	"github.com/MbuguaCaleb/go_rest_api_one/database"
	"github.com/MbuguaCaleb/go_rest_api_one/models"
	"github.com/gofiber/fiber/v2"
)

//good practice--->custom way to represent my datatypes in GO
//custom datatype in your routes
type UserSerializer struct{
	//this is not the model User, see this as the serializer
	ID uint `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
}

//it will take in my UserModel and return my serializer
//Serializer  return data into simple readbale format
//It will be used as a helper for my  endpoints
func CreateResponseUser(userModel models.User) UserSerializer {
	 return UserSerializer{ID:userModel.ID,FirstName: userModel.FirstName,LastName: userModel.LastName}
}


//Endpoints
func CreateUser (c *fiber.Ctx) error {
	//I am creating an empty model instance
	var user models.User

	//binds my request body to my Struct
	if err := c.BodyParser(&user); err != nil{
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&user)

	responseUser := CreateResponseUser(user)
	return c.Status(200).JSON(responseUser)

}

func GetUsers(c *fiber.Ctx) error {
	users := []models.User{}

	database.Database.Db.Find(&users)
	responseUsers := []UserSerializer{}

	for _, user := range users {
		responseUser := CreateResponseUser(user)
		responseUsers = append(responseUsers, responseUser)
	}

	return c.Status(200).JSON(responseUsers)
}