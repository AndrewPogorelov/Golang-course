//Данная задача ориентирована на реальную работу с данными в формате JSON.
//Для решения мы будем использовать справочник ОКВЭД (Общероссийский
//классификатор видов экономической деятельности), который был размещен
//на web-портале data.gov.ru.

//Необходимая вам информация о структуре данных содержится в файле
//structure-20190514T0000.json, а сами данные, которые вам потребуется
//декодировать, содержатся в файле data-20190514T0100.json.
//Файлы размещены в нашем репозитории на github.com.

//Для того, чтобы показать, что вы действительно смогли декодировать
//документ вам необходимо в качестве ответа записать сумму полей
//global_id всех элементов, закодированных в файле.

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type YouStruct struct {
	Id int `json:"global_id"`
}

func readTextFile(filePath string) int {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	dec := json.NewDecoder(f)
	dst := []YouStruct{}
	dec.Decode(&dst)

	sum := 0
	for _, el := range dst {
		sum += el.Id
	}

	return sum
}

func main() {
	fmt.Println(readTextFile("C:\\Users\\PogorelovJS\\Desktop\\Golang\\stepik-go-master\\work_with_json\\data-20190514T0100.json"))
}
