package main

import "errors"

type ProductType string

const (
	Smartphone ProductType = "смартфон"
	Computer   ProductType = "компьютер"
	Television ProductType = "телевизор"
)

type Product struct {
	Type     ProductType
	Price    float64
	Duration int
}

func CalculateTotal(p Product) (float64, error) {
	validDurations := map[ProductType][]int{
		Smartphone: {3, 6, 9},
		Computer:   {3, 6, 9, 12},
		Television: {3, 6, 9, 12, 18},
	}

	interestRates := map[ProductType]float64{
		Smartphone: 0.03,
		Computer:   0.04,
		Television: 0.05,
	}

	durations, ok := validDurations[p.Type]
	if !ok {
		return 0, errors.New("неизвестный тип продукта")
	}

	isValid := false
	for _, d := range durations {
		if d == p.Duration {
			isValid = true
			break
		}
	}
	if !isValid {
		return 0, errors.New("недопустимый срок рассрочки")
	}

	steps := (p.Duration - 3) / 3
	rate := interestRates[p.Type]
	total := p.Price + p.Price*rate*float64(steps)

	return total, nil
}
