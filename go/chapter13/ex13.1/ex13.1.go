package main

import "fmt"

type House struct {
	Address  string
	Size     int
	Price    float64
	Category string
}

func main() {
	var house House
	house.Address = "서울시 강남구 ..."
	house.Size = 28
	house.Price = 10
	house.Category = "아파트"

	fmt.Println(house)
	fmt.Printf("%v \n", house)
	fmt.Printf("주소:%s 사이즈:%d평, 가격:%v억원, 종류:%s \n", house.Address, house.Size, int(house.Price), house.Category)
}
