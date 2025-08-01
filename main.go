package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var productType string
	var price float64
	var duration int
	var phone string

	fmt.Print("Введите тип продукта (смартфон, компьютер, телевизор): ")
	fmt.Scanln(&productType)
	productType = strings.ToLower(productType)

	allowedTypes := map[string]bool{
		"смартфон":  true,
		"компьютер": true,
		"телевизор": true,
	}

	if !allowedTypes[productType] {
		fmt.Println("Ошибка: неизвестный тип продукта")
		return
	}

	fmt.Print("Введите цену продукта: ")
	var input string
	fmt.Scanln(&input)
	priceParsed, err := strconv.ParseFloat(input, 64)
	if err != nil || priceParsed <= 0 {
		fmt.Println("Ошибка: введите корректную положительную цену")
		return
	}
	price = priceParsed

	fmt.Print("Введите срок рассрочки (в месяцах): ")
	fmt.Scanln(&input)

	durationParsed, err := strconv.Atoi(input)
	if err != nil || durationParsed <= 0 {
		fmt.Println("Ошибка: введите корректный срок рассрочки")
		return
	}
	duration = durationParsed

	valid := false
	switch productType {
	case "смартфон":
		valid = duration == 3 || duration == 6 || duration == 9
	case "компьютер":
		valid = duration == 3 || duration == 6 || duration == 9 || duration == 12
	case "телевизор":
		valid = duration == 3 || duration == 6 || duration == 9 || duration == 12 || duration == 15 || duration == 18
	default:
		fmt.Println("Ошибка: неизвестный тип продукта")
		return
	}

	if !valid {
		fmt.Println("Ошибка: неверный срок рассрочки для выбранного продукта")
		return
	}

	fmt.Print("Введите номер телефона: ")
	fmt.Scanln(&phone)

	if len(phone) != 13 {
		fmt.Println("Ошибка: номер должен содержать ровно 13 символов")
		return
	}
	if !strings.HasPrefix(phone, "+992") {
		fmt.Println("Ошибка: номер должен начинаться с +992")
		return
	}
	for _, ch := range phone[4:] {
		if ch < '0' || ch > '9' {
			fmt.Println("Ошибка: номер должен содержать только цифры после +992")
			return
		}
	}
	product := Product{
		Type:     ProductType(productType),
		Price:    price,
		Duration: duration,
	}

	total, err := CalculateTotal(product)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	message := fmt.Sprintf("Вы купили %s на %d мес. Сумма: %.2f сомони", product.Type, product.Duration, total)
	SendSMS(phone, message)
	fmt.Println("Платёж успешно завершён.")
}
