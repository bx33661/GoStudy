package main

import "fmt"

type Phone struct {
	band    string
	name    string
	country string
	price   int
}

func main() {
	//直接打印内容
	fmt.Println(Phone{"Huawei", "P70", "China", 9999})

	//variable_name := structure_variable_type {value1, value2...valuen}
	Bx_Phone := Phone{
		band:    "HUAWEI",
		name:    "Mete70",
		country: "CN",
		price:   6999,
	}
	fmt.Println(Bx_Phone)

	//访问结构体成员
	// 使用.字符访问
	dx_phone := Phone{}
	dx_phone.name = "Aplle"
	dx_phone.country = "USA"
	dx_phone.band = "MAC"
	fmt.Println(dx_phone)
	fmt.Println(dx_phone.price)
	fmt.Println(dx_phone.name)
	fmt.Println(dx_phone.country)
	fmt.Println(dx_phone.band)
}
