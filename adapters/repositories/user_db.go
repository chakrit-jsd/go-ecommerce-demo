package repositories

import (
	"errors"
	"fmt"
	"go-ecommerce-demo/core/models/model_gorm"
	"go-ecommerce-demo/core/models/model_io"
	"go-ecommerce-demo/core/ports"

	"gorm.io/gorm"
)

type userRepositoryDB struct {
	db *gorm.DB
}

func NewUserRepositoryDB(db *gorm.DB) ports.UserRepository {

	// Create table from model
	// db.AutoMigrate(
	// 	&model_gorm.User{},
	// 	&model_gorm.Category{},
	// 	&model_gorm.Product{},
	// 	&model_gorm.Cart{},
	// 	&model_gorm.CartDetail{},
	// 	&model_gorm.Order{},
	// 	&model_gorm.OrderDetail{},
	// )

	return userRepositoryDB{db}
}

func (r userRepositoryDB) Create(user *model_gorm.User) (err error) {

	// Mock data table user
	// users := []model_gorm.User{}

	// for x := 6; x <= 1000; x++ {
	// 	b := []byte(fmt.Sprintf("test%v", x))
	// 	byte, err := bcrypt.GenerateFromPassword(b, bcrypt.DefaultCost)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	users = append(users, model_gorm.User{
	// 		UserName: fmt.Sprintf("test%v", x),
	// 		Password: string(byte),
	// 		Address:  fmt.Sprintf("%v a%v%v%v xyz", x, x, x, x),
	// 	})
	// }

	err = r.db.Model(&model_gorm.User{}).Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r userRepositoryDB) Update(user *model_gorm.User, userId *int) (err error) {

	res := r.db.Model(&model_gorm.User{}).Where("id = ?", userId).Updates(&user)

	if res.Error != nil {
		return res.Error
	}
	if res.RowsAffected == 0 {
		return errors.New("update fail id not found")
	}

	return nil
}

func (r userRepositoryDB) Delete(userId *int) error {

	res := r.db.Model(&model_gorm.User{}).Delete("id = ?", userId)
	if res.Error != nil {
		return res.Error
	}

	if res.RowsAffected == 0 {
		return errors.New("delete fail id not found")
	}

	return nil
}

func (r userRepositoryDB) GetUserById(userId *int) (user *model_io.User, err error) {

	err = r.db.Omit("password").Model(&model_gorm.User{}).First(&user, userId).Error
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r userRepositoryDB) GetUserByUserName(username *string) (user *model_io.User, err error) {

	err = r.db.Model(&model_gorm.User{}).Where("user_name = ?", username).First(&user).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return user, nil
}
