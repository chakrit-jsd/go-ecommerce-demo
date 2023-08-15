package services

import (
	"go-ecommerce-demo/core/models/model_gorm"
	"go-ecommerce-demo/core/models/model_io"
	"go-ecommerce-demo/core/ports"
	"go-ecommerce-demo/utils"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	userRepo ports.UserRepository
}

func NewUserService(userRepo ports.UserRepository) ports.UserService {

	return userService{userRepo}
}

func (s userService) Create(u *model_io.User) error {

	byte, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {

		return fiber.ErrUnprocessableEntity
	}

	user := model_gorm.User{
		UserName: u.UserName,
		Password: string(byte),
		Address:  u.Address,
	}

	err = s.userRepo.Create(&user)
	if err != nil {
		return utils.CusErrorDB(err)
	}

	return nil
}

func (s userService) Update(u *model_io.User, userId *int) error {

	user := model_gorm.User{
		Address: u.Address,
	}

	err := s.userRepo.Update(&user, userId)
	if err != nil {
		return utils.CusErrorDB(err)
	}

	return nil
}

func (s userService) Delete(userId *int) error {

	err := s.userRepo.Delete(userId)
	if err != nil {
		return utils.CusErrorDB(err)
	}

	return nil
}

func (s userService) GetUserById(userId *int) (*model_io.User, error) {

	user, err := s.userRepo.GetUserById(userId)
	if err != nil {
		return nil, utils.CusErrorDB(err)
	}

	return user, nil
}

func (s userService) GetUserByUserName(username *string) (*model_io.User, error) {

	user, err := s.userRepo.GetUserByUserName(username)
	if err != nil {
		return nil, utils.CusErrorDB(err)
	}

	return user, nil
}

func (s userService) Login(data *model_io.User) (*string, error) {
	user, err := s.userRepo.GetUserByUserName(&data.UserName)
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))
	if err != nil {
		return nil, fiber.NewError(fiber.StatusBadRequest, "Invalid username or password")
	}

	cliams := jwt.MapClaims{
		"iss": strconv.Itoa(user.ID),
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	token, err := jwtToken.SignedString([]byte(viper.GetString("jwt.secret")))
	if err != nil {
		return nil, fiber.NewError(fiber.StatusInternalServerError, "Something went wrong Please try again")
	}

	return &token, nil
}
