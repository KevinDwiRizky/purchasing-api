package controllers

import (
	"purchasing-api/config"
	"purchasing-api/models"
	"purchasing-api/utils"
	"strings"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return utils.ResponseError(c, 400, "Request tidak valid")
	}

	if user.Username == "" {
		return utils.ResponseError(c, 400, "Username tidak boleh kosong")
	}

	if user.Password == "" {
		return utils.ResponseError(c, 400, "Password tidak boleh kosong")
	}

	var count int64
	config.DB.Model(&models.User{}).Where("username = ?", user.Username).Count(&count)
	if count > 0 {
		return utils.ResponseError(c, 409, "Username sudah terdaftar")
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	user.Password = string(hash)

	if err := config.DB.Create(&user).Error; err != nil {
		if strings.Contains(err.Error(), "duplicate key") || strings.Contains(err.Error(), "idx_users_username") {
			return utils.ResponseError(c, 409, "Username sudah terdaftar")
		}
		return utils.ResponseError(c, 500, "Gagal membuat akun")
	}

	userData := fiber.Map{
		"id":       user.ID,
		"username": user.Username,
		"role":     user.Role,
	}

	return utils.ResponseSuccess(c, 201, "Registrasi berhasil", userData)
}

func Login(c *fiber.Ctx) error {
	var request struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&request); err != nil {
		return utils.ResponseError(c, 400, "Request tidak valid")
	}

	if request.Username == "" {
		return utils.ResponseError(c, 400, "Username tidak boleh kosong")
	}

	if request.Password == "" {
		return utils.ResponseError(c, 400, "Password tidak boleh kosong")
	}

	var user models.User
	if err := config.DB.Where("username = ?", request.Username).First(&user).Error; err != nil {
		return utils.ResponseError(c, 401, "Username atau password salah")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return utils.ResponseError(c, 401, "Username atau password salah")
	}

	token, err := utils.GenerateToken(user.ID)
	if err != nil {
		return utils.ResponseError(c, 500, "Gagal membuat token")
	}

	loginData := fiber.Map{
		"id":       user.ID,
		"username": user.Username,
		"role":     user.Role,
		"token":    token,
	}

	return utils.ResponseSuccess(c, 200, "Login berhasil", loginData)
}
