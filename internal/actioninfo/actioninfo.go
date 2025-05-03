package actioninfo

import (
	"log"
)

type DataParser interface {
	Parse(string) error
	ActionInfo() (info string, err error)
}

func Info(dataset []string, dp DataParser) {
	if len(dataset) == 0 {
		log.Printf("Empty ")
	}
	for _, data := range dataset {
		// поиск ошибок парсинга даты
		err := dp.Parse(data)
		if err != nil {
			log.Printf("Error Parse ,%s : %v", data, err)
			continue
		} else {
			log.Printf("processed test data\n")
		}
		// Acctioninfo
		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("error ActionInfo'%s': %v", data, err)
			continue
		}

		// вывод на экран
		log.Println(info)
	}
}
