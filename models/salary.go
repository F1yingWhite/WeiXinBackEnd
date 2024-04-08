package models

import (
	"fmt"
	"log"
	"time"
)

type Salary struct {
	ID             uint       `gorm:"primary_key" json:"id"`
	CreatedAt      time.Time  `json:"-"`
	UpdatedAt      time.Time  `json:"-"`
	DeletedAt      *time.Time `sql:"index" json:"-"`
	UserId         string     `json:"-"`                       // 创建的用户ID
	Company        string     `json:"company" gorm:"not null"` // 公司名称
	City           string     `json:"city" gorm:"not null"`    // 城市
	Salary         string     `json:"salary" gorm:"not null"`  // 薪资描述
	Major          string     `json:"major" gorm:"not null"`   // 学历
	Name           string     `json:"name"`                    // 名称
	CategoryFirst  string     `json:"category_first"`          //大分类
	CategorySecond string     `json:"category_second"`         //细分
}

func CreateSalary(salary *Salary) error {
	logPrefix := fmt.Sprintf("[models/Salary]: CreateSalary(createSalary.Company: %s)", salary.Company)

	log.Printf("%s: 正在创建...", logPrefix)
	if err := DB.Create(&salary).Error; err != nil {
		log.Printf("%s: 创建失败(%s)", logPrefix, err)
		return err
	}
	log.Printf("%s: 创建成功(id = %s)", logPrefix, salary.Company)
	return nil
}

func DeteleSalaryById(id uint) error {
	logPrefix := fmt.Sprintf("[models/Salary]: DeteleSalaryById(id: %d)", id)

	log.Printf("%s: 正在删除...", logPrefix)
	if err := DB.Where("id = ?", id).Delete(&Salary{}).Error; err != nil {
		log.Printf("%s: 删除失败(%s)", logPrefix, err)
		return err
	}
	log.Printf("%s: 删除成功", logPrefix)
	return nil
}

func GetSalaries(page, pageSize int) ([]Salary, error) {
	logPrefix := "[models/Salary]: GetSalaries()"

	log.Printf("%s: 正在查询...", logPrefix)
	var salaries []Salary
	if err := DB.Offset((page - 1) * pageSize).Limit(pageSize).Find(&salaries).Error; err != nil {
		log.Printf("%s: 查询失败(%s)", logPrefix, err)
		return nil, err
	}
	log.Printf("%s: 查询成功", logPrefix)
	return salaries, nil
}

func GetSalaryById(id uint) (*Salary, error) {
	logPrefix := fmt.Sprintf("[models/Salary]: GetSalaryById(id: %d)", id)

	log.Printf("%s: 正在查询...", logPrefix)
	var salary Salary
	if err := DB.Where("id = ?", id).First(&salary).Error; err != nil {
		log.Printf("%s: 查询失败(%s)", logPrefix, err)
		return nil, err
	}
	log.Printf("%s: 查询成功", logPrefix)
	return &salary, nil
}

func GetSalaryByCompany(company string, page, pageSize int) ([]Salary, error) {
	logPrefix := fmt.Sprintf("[models/Salary]: GetSalaryByCompany(company: %s)", company)

	log.Printf("%s: 正在查询...", logPrefix)
	var salaries []Salary
	//修改为模糊查询
	if err := DB.Where("company LIKE ?", "%"+company+"%").Offset((page - 1) * pageSize).Limit(pageSize).Find(&salaries).Error; err != nil {
		log.Printf("%s: 查询失败(%s)", logPrefix, err)
		return nil, err
	}
	log.Printf("%s: 查询成功", logPrefix)
	return salaries, nil
}

func GetSalaryByCity(city string, page, pageSize int) ([]Salary, error) {
	logPrefix := fmt.Sprintf("[models/Salary]: GetSalaryByCity(city: %s)", city)

	log.Printf("%s: 正在查询...", logPrefix)
	var salaries []Salary
	if err := DB.Where("city = ?", city).Offset((page - 1) * pageSize).Limit(pageSize).Find(&salaries).Error; err != nil {
		log.Printf("%s: 查询失败(%s)", logPrefix, err)
		return nil, err
	}
	log.Printf("%s: 查询成功", logPrefix)
	return salaries, nil
}

func GetSalariesByCompanyAndCity(company, city string, page, pageSize int) ([]Salary, error) {
	logPrefix := fmt.Sprintf("[models/Salary]: GetSalariesByCompanyAndCity(company: %s, city: %s)", company, city)

	log.Printf("%s: 正在查询...", logPrefix)
	var salaries []Salary
	if err := DB.Where("company LIKE ? AND city = ?", "%"+company+"%", city).Offset((page - 1) * pageSize).Limit(pageSize).Find(&salaries).Error; err != nil {
		log.Printf("%s: 查询失败(%s)", logPrefix, err)
		return nil, err
	}
	log.Printf("%s: 查询成功", logPrefix)
	return salaries, nil
}

func GetSalaryByUserId(user_id string, page, pageSize int) ([]Salary, error) {
	logPrefix := fmt.Sprintf("[models/Salary]: GetSalaryByUserId(user_id: %s)", user_id)

	log.Printf("%s: 正在查询...", logPrefix)
	var salaries []Salary
	if err := DB.Where("user_id = ?", user_id).Offset((page - 1) * pageSize).Limit(pageSize).Find(&salaries).Error; err != nil {
		log.Printf("%s: 查询失败(%s)", logPrefix, err)
		return nil, err
	}
	log.Printf("%s: 查询成功", logPrefix)
	return salaries, nil
}

func (salary *Salary) UpdateSalary() error {
	logPrefix := fmt.Sprintf("[models/Salary]: UpdateSalary(id: %d)", salary.ID)

	log.Printf("%s: 正在更新...", logPrefix)
	if err := DB.Save(salary).Error; err != nil {
		log.Printf("%s: 更新失败(%s)", logPrefix, err)
		return err
	}
	log.Printf("%s: 更新成功", logPrefix)
	return nil
}
