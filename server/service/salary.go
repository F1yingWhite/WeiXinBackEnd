package service

import (
	"weixin_backend/models"

	"github.com/gin-gonic/gin"
)

type CreateSalary struct {
	Company        string `form:"company" binding:"required"`
	City           string `form:"city" binding:"required"`
	Salary         string `form:"salary" binding:"required"`
	Major          string `form:"major" binding:"required"`
	Name           string `form:"name" binging:"required"`
	CategoryFirst  string `form:"category_first" binging:"required"`
	CategorySecond string `form:"category_second" binging:"required"`
}

func (createSalary *CreateSalary) Handle(c *gin.Context) (any, error) {
	authorization := c.Request.Header.Get("Authorization")
	salary := &models.Salary{
		Company:        createSalary.Company,
		City:           createSalary.City,
		Salary:         createSalary.Salary,
		Major:          createSalary.Major,
		Name:           createSalary.Name,
		CategoryFirst:  createSalary.CategoryFirst,
		CategorySecond: createSalary.CategorySecond,
		UserId:         authorization,
	}

	err := models.CreateSalary(salary)
	if err != nil {
		return nil, err
	}
	return map[string]interface{}{"msg": "success"}, nil
}

type GetSalary struct {
	Page     int `form:"page" binding:"required"`
	PageSize int `form:"page_size" binding:"required"`
}

func (getSalary *GetSalary) Handle(c *gin.Context) (any, error) {
	return models.GetSalaries(getSalary.Page, getSalary.PageSize)
}

type GetSalaryByCompany struct {
	Page     int    `form:"page" binding:"required"`
	PageSize int    `form:"page_size" binding:"required"`
	Company  string `form:"company" binding:"required"`
}

func (getSalaryByCompany *GetSalaryByCompany) Handle(c *gin.Context) (any, error) {
	return models.GetSalaryByCompany(getSalaryByCompany.Company, getSalaryByCompany.Page, getSalaryByCompany.PageSize)
}

type GetSalaryByCity struct {
	Page     int    `form:"page" binding:"required"`
	PageSize int    `form:"page_size" binding:"required"`
	City     string `form:"city" binding:"required"`
}

func (getSalaryByCity *GetSalaryByCity) Handle(c *gin.Context) (any, error) {
	return models.GetSalaryByCity(getSalaryByCity.City, getSalaryByCity.Page, getSalaryByCity.PageSize)
}

type GetSalaryById struct {
	Id uint `form:"id" binding:"required"`
}

func (getSalaryById *GetSalaryById) Handle(c *gin.Context) (any, error) {
	return models.GetSalaryById(getSalaryById.Id)
}

type GetSalariesByCompanyAndCity struct {
	Page     int    `form:"page" binding:"required"`
	PageSize int    `form:"page_size" binding:"required"`
	Company  string `form:"company" binding:"required"`
	City     string `form:"city" binding:"required"`
}

func (getSalariesByCompanyAndCity *GetSalariesByCompanyAndCity) Handle(c *gin.Context) (any, error) {
	return models.GetSalariesByCompanyAndCity(getSalariesByCompanyAndCity.Company, getSalariesByCompanyAndCity.City, getSalariesByCompanyAndCity.Page, getSalariesByCompanyAndCity.PageSize)
}

type GetSalaryByUserId struct {
	Page     int    `form:"page" binding:"required"`
	PageSize int    `form:"page_size" binding:"required"`
	UserId   string `form:"user_id" binding:"required"`
}

func (getSalaryByUserId *GetSalaryByUserId) Handle(c *gin.Context) (any, error) {
	return models.GetSalaryByUserId(getSalaryByUserId.UserId, getSalaryByUserId.Page, getSalaryByUserId.PageSize)
}
