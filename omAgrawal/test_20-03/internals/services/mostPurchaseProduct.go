package services

import (
	"fmt"
	//"inventory_Management/product"
	"strconv"
)

func MostPurchasedProduct(records [][]string) {
	MostPurchasedProductMap := make(map[string]int)
	qnt := 0
	var str string
	for _, record := range records {
		qnty, _ := strconv.Atoi(record[4])
		MostPurchasedProductMap[record[2]] += qnty
		if MostPurchasedProductMap[record[2]] > qnt {
			qnt = MostPurchasedProductMap[record[2]]
			str = record[2]
		}
	}
	fmt.Println("most purchased Product is :", str, "and Quantity ", qnt)
	//fmt.Println(MostPurchasedProductMap)
}
