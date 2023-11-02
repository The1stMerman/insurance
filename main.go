package main

import "fmt"

type carInsurance struct {
	brand string
	model string
}

type travelInsurance struct {
	country string
	day     int
}

type healthInsurance struct {
	sex string
	age int
	job string
}

func main() {
	car := carInsurance{"Honda", "Civic"}
	travel := travelInsurance{"Italy", 7}
	health := healthInsurance{"Male", 60, "Driver"}

	fmt.Printf("ประกันรถ : %.2f\n", insurance(car))
	fmt.Printf("ประกันการเดินทาง : %.2f\n", insurance(travel))
	fmt.Printf("ประกันอุบัติเหตุ : %.2f\n", insurance(health))
}

func insurance(i interface{}) float64 {
	switch t := i.(type) {
	case carInsurance:
		carRates := map[string]map[string]float64{
			"Toyota": {
				"Yaris": 1000.00,
				"Altis": 1500.00,
				"Camry": 2000.00,
			},
			"Honda": {
				"Accord": 1000.00,
				"City":   1500.00,
				"Civic":  2000.00,
			},
		}
		return carRates[t.brand][t.model]

	case travelInsurance:
		countryRates := map[string]float64{
			"Japan":   1500.00,
			"England": 2000.00,
			"US":      2500.00,
			"Italy":   3000.00,
		}
		return countryRates[t.country] * float64(t.day)

	case healthInsurance:
		sexRates := map[string]float64{
			"Male":   2,
			"Female": 1.5,
		}
		jobRates := map[string]float64{
			"Driver": 300.00,
			"HR":     150.00,
			"SE":     100.00,
		}
		return sexRates[t.sex]*jobRates[t.job] + float64(t.age)
	}

	return 0.0
}
