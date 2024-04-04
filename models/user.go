package models

import (
	"fmt"
	"log"
)

type User struct {
	UserId    string `json:"user_id" gorm:"primaryKey"`             // 用户ID
	Username  string `json:"username" gorm:"default:热心用户;not null"` // 用户名
	Signature string `json:"signature" gorm:"default:我很帅我知道;"`      // 用户个性签名
	//一个用户可以创建多个薪资信息
	Salaries []Salary `json:"salaries" gorm:"foreignKey:UserId;references:UserId"`
}

func CreateUser(user_id string) (*User, error) {
	logPrefix := fmt.Sprintf("[models/user]: CreateUser(username: %s)", user_id)

	log.Printf("%s: 正在创建...", logPrefix)
	user := User{UserId: user_id}

	if err := DB.Create(&user).Error; err != nil {
		log.Printf("%s: 创建失败(%s)", logPrefix, err)
		return nil, err
	}
	log.Printf("%s: 创建成功(id = %s)", logPrefix, user.UserId)
	return &user, nil
}

func GetUserById(user_id string) (*User, error) {
	logPrefix := fmt.Sprintf("[models/user]: GetUserById(user_id: %s)", user_id)

	log.Printf("%s: 正在查询...", logPrefix)
	var user User
	if err := DB.Where("user_id = ?", user_id).First(&user).Error; err != nil {
		log.Printf("%s: 查询失败(%s)", logPrefix, err)
		return nil, err
	}
	log.Printf("%s: 查询成功", logPrefix)
	return &user, nil
}

func (user *User) UpdateUser() error {
	logPrefix := fmt.Sprintf("[models/user]: UpdateUser(user_id: %s)", user.UserId)

	log.Printf("%s: 正在更新...", logPrefix)
	if err := DB.Save(user).Error; err != nil {
		log.Printf("%s: 更新失败(%s)", logPrefix, err)
		return err
	}
	log.Printf("%s: 更新成功", logPrefix)
	return nil
}
