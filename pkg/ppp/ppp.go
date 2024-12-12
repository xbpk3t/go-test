package ppp

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

// CountryPPP 代表国家购买力平价信息
type CountryPPP struct {
	Range       string `json:"range"`
	CountryCode string `json:"countryCode"`
	CountryName string `json:"countryName"`
}

// Discount 代表折扣信息
type Discount struct {
	Code     string  `json:"code"`
	Discount float64 `json:"discount"`
}

// Result 合并国家购买力信息和折扣信息
type Result struct {
	Range        string  `json:"range"`
	CountryCode  string  `json:"countryCode"`
	CountryName  string  `json:"countryName"`
	DiscountCode string  `json:"discountCode"`
	Discount     float64 `json:"discount"`
}

// GetDiscountInfo 根据国家代码获取购买力平价和折扣信息
// purchasing-power-parity
func GetDiscountInfo(countryCode string) (*Result, error) {
	// 假设 flatppp 是一个全局变量，包含所有国家的购买力数据
	var flatppp []CountryPPP

	// 这里需要加载或定义你的购买力数据
	// 例如，可以从文件或数据库中加载
	// 为了简化，这里假设 flatppp 已经被填充

	// 查找国家
	countryPPP := findCountry(countryCode, flatppp)

	if countryPPP == nil {
		return nil, fmt.Errorf("country not found")
	}

	// 获取环境变量中的折扣信息
	discount := getDiscount(os.Getenv("CLOUDFLARE_ENV"), countryPPP.Range)

	// 合并结果
	result := mergeDiscountResult(countryPPP, discount)
	return result, nil
}

// findCountry 在购买力列表中找到对应国家
func findCountry(countryCode string, flatppp []CountryPPP) *CountryPPP {
	for _, deal := range flatppp {
		if deal.CountryCode == countryCode {
			return &deal
		}
	}
	return nil
}

// getDiscount 根据购买力水平，在环境变量里找到配置的折扣信息
func getDiscount(env, rangeValue string) Discount {
	switch rangeValue {
	case "0.0-0.1":
		return Discount{Code: os.Getenv("level0_1"), Discount: parseInt(os.Getenv("level0_1_discount"), 0)}
	case "0.1-0.2":
		return Discount{Code: os.Getenv("level1_2"), Discount: parseInt(os.Getenv("level1_2_discount"), 0)}
	// ... 其他case
	default:
		return Discount{Code: "", Discount: 0}
	}
}

// parseInt 尝试将字符串解析为整数，如果失败则返回默认值
func parseInt(value string, defaultValue int) int {
	if v, err := strconv.Atoi(value); err == nil {
		return v
	}
	return defaultValue
}

// mergeDiscountResult 合并国家购买力信息和折扣信息
func mergeDiscountResult(countryPPP *CountryPPP, discount Discount) *Result {
	return &Result{
		Range:        countryPPP.Range,
		CountryCode:  countryPPP.CountryCode,
		CountryName:  countryPPP.CountryName,
		DiscountCode: discount.Code,
		Discount:     discount.Discount,
	}
}

// ToJSON 将Result转换为JSON字符串
func (r *Result) ToJSON() string {
	jsonData, err := json.Marshal(r)
	if err != nil {
		log.Fatal(err)
	}
	return string(jsonData)
}
