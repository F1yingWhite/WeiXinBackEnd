package service

import (
	"log"
	"weixin_backend/models"

	"github.com/gin-gonic/gin"
)

type UserInfo struct {
}

func (userInfo *UserInfo) Handle(c *gin.Context) (any, error) {
	authorization := c.Request.Header.Get("Authorization")
	user, err := models.GetUserById(authorization)
	if err != nil {
		return nil, err
	}
	return user, nil
}

type UpdateUserInfo struct {
	Username  string `form:"username"`
	Signature string `form:"signature"`
}

func (updateUserInfo *UpdateUserInfo) Handle(c *gin.Context) (any, error) {
	authorization := c.Request.Header.Get("Authorization")
	log.Printf("username: %s, signature: %s", updateUserInfo.Username, updateUserInfo.Signature)
	user, err := models.GetUserById(authorization)
		if err != nil {
			return nil, err
		}
		if updateUserInfo.Username != "" {
			user.Username = updateUserInfo.Username
		}
		if updateUserInfo.Signature != "" {
			user.Signature = updateUserInfo.Signature
		}
		err = user.UpdateUser()
		if err != nil {
			return nil, err
		}
		return map[string]interface{}{"msg": "success"}, nil
	}
