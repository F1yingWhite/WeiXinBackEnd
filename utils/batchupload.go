package utils

import (
	"encoding/json"
	"io"
	"os"
)

type Category struct {
	Value    int        `json:"value"`
	Label    string     `json:"label"`
	Children []Category `json:"children"`
}

type CategoryResponse struct {
	Data []Category `json:"data"`
}

func BindJSONToMap(path string) (map[int]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var response CategoryResponse
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, err
	}

	categoryMap := make(map[int]string)
	for _, category := range response.Data {
		categoryMap[category.Value] = category.Label
		for _, subCategory := range category.Children {
			categoryMap[subCategory.Value] = subCategory.Label
		}
	}

	return categoryMap, nil
}

type Salary struct {
	CompanyName    string `json:"company_name" `   // 公司名称
	City           string `json:"city" `           // 城市
	Salary         string `json:"salary" `         // 薪资描述
	Major          string `json:"major" `          // 学历
	Name           string `json:"name"`            // 名称
	CategoryFirst  int    `json:"category_first"`  //大分类
	CategorySecond int    `json:"category_second"` //细分
}

type JobResponse struct {
	Data struct {
		Jobs []Salary `json:"jobs"`
	} `json:"data"`
}

func BindJSONToStruct(jsonData []byte) (JobResponse, error) {
	var response JobResponse
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return JobResponse{}, err
	}
	return response, nil
}

func ReadJSONFromFile(filename string) (JobResponse, error) {
	file, err := os.Open(filename)
	if err != nil {
		return JobResponse{}, err
	}
	defer file.Close()

	bytes, err := io.ReadAll(file)
	if err != nil {
		return JobResponse{}, err
	}

	return BindJSONToStruct(bytes)
}
