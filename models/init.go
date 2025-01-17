package models

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"weixin_backend/utils"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func sqliteDB(dsn string, config *gorm.Config) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dsn+"?_pragma=foreign_keys(1)"), config)
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxOpenConns(1)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func InitDB() {
	var db *gorm.DB
	var err error
	db, err = sqliteDB("weixin_DB.db", &gorm.Config{})

	if err != nil {
		log.Panicf("无法连接数据库，%s", err)
	}

	DB = db
	Migrate()

	//创建系统管理员xyh
	_, err = GetUserById("xyh")
	if err != nil {
		CreateUser("xyh")
		ToDataBase()
	}

}

func ToDataBase() {
	// 读取分类
	categoryMap, err := utils.BindJSONToMap("categories.json")
	if err != nil {
		panic(err)
	}
	// 保存数据
	for i := 1; i <= 8; i++ {
		response, err := utils.ReadJSONFromFile(getData(i))
		if err != nil {
			panic(err)
		}
		for _, job := range response.Data.Jobs {
			err := CreateSalary(&Salary{ // Pass a pointer to models.Salary
				UserId:         "xyh",
				CategoryFirst:  categoryMap[job.CategoryFirst],
				CategorySecond: categoryMap[job.CategorySecond],
				Company:        job.CompanyName,
				City:           job.City,
				Salary:         job.Salary,
				Major:          job.Major,
				Name:           job.Name,
			})
			if err != nil {
				panic(err)
			}
		}
	}
}

func Migrate() {
	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Salary{})
}

func getData(page int) []byte {
	// 目标URL
	url := fmt.Sprintf("https://offershow.cn/api/od/search_job?size=50&page=%d", page)
	resp, err := http.Post(url, "", nil)
	if err != nil {
		log.Fatalf("请求失败：%v", err)
	}

	defer resp.Body.Close()

	bodyBytes, _ := io.ReadAll(resp.Body)
	return bodyBytes
}
