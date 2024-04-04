package service

import "github.com/gin-gonic/gin"

type Search struct {
	Company string `form:"company"`
	City    string `form:"city"`
}

func (s *Search) Handle(c *gin.Context) (any, error) {
	return s, nil
}
