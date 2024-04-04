package models

import (
	"fmt"
	"log"

	"gorm.io/gorm"
)

type Salary struct {
	gorm.Model
	UserId    string `json:"user_id"`                   // 创建的用户ID
	Company   string `json:"company" gorm:"not null"`   // 公司名称
	City      string `json:"city" gorm:"not null"`      // 城市
	Position  string `json:"position" gorm:"not null"`  // 岗位
	Salary    string `json:"salary" gorm:"not null"`    // 薪资描述
	YearEnd   string `json:"year_end"`                  // 年终换算范围
	Type      string `json:"type"`                      // 类型（校招/实习）
	Education string `json:"education" gorm:"not null"` // 学历
	Industry  string `json:"industry" `                 // 行业
	Remark    string `json:"remark"`                    // 备注
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

func GetSalaryByUserId(user_id string) ([]Salary, error) {
	logPrefix := fmt.Sprintf("[models/Salary]: GetSalaryByUserId(user_id: %s)", user_id)

	log.Printf("%s: 正在查询...", logPrefix)
	var salaries []Salary
	if err := DB.Where("user_id = ?", user_id).Find(&salaries).Error; err != nil {
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
