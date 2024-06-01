package UserController

import (
	"fmt"
	"log"
	"os"
	"portalapp/Database"
	usermodel "portalapp/Model/UserModel"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func AddUser(ctx *fiber.Ctx) error {

	var body struct {
		UserName string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
		RoleID   uint   `json:"role_id"`
	}

	err := ctx.BodyParser(&body)
	fmt.Println(body)

	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "Invalid Data"})
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": "Failed to hash password."})
	}

	user := usermodel.User{Email: body.Email, RoleID: body.RoleID, UserName: body.UserName, Password: string(hash), CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	result := Database.DB.Create(&user)

	if result.Error != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": "Failed to create user."})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "User created sucessfully",
	})
}

func Login(ctx *fiber.Ctx) error {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := ctx.BodyParser(&body)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "Invalid Data."})
	}

	var user usermodel.User
	Database.DB.First(&user, "email = ?", body.Email)
	if user.ID == 0 {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "User not found."})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"success": false, "message": "Password Mismatch."})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"sub": user.ID, "exp": time.Now().Add(time.Hour * 24 * 30).Unix()})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"success": false, "message": "Failed to create token."})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": tokenString,
	})
}

func GetLogInForm(ctx *fiber.Ctx) error {
	log.Println("Providing log in form.")
	return ctx.Render("login", fiber.Map{})
}

func ProcessLogInData(ctx *fiber.Ctx) error {
	log.Println("Processing log in form.")
	email := ctx.FormValue("email")
	password := ctx.FormValue("password")

	var user usermodel.User
	Database.DB.First(&user, "email = ?", email)
	if user.ID == 0 {
		return ctx.Render("login_error", fiber.Map{"message": "Failed to log in."})
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return ctx.Render("login_error", fiber.Map{"message": "Incorrect password."})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"sub": user.ID, "exp": time.Now().Add(time.Hour * 24 * 30).Unix()})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return ctx.Render("login_error", fiber.Map{"message": "Server error.Failed to create token."})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    tokenString,
		Expires:  time.Now().Add(time.Hour * 24 * 30),
		HTTPOnly: true,
		Secure:   true,
	}

	ctx.Cookie(&cookie)

	return ctx.Render("userhome", fiber.Map{})
}
