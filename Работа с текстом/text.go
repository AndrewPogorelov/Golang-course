//Данная задача в основном ориентирована на изучение типа bufio.Reader, поскольку этот тип позволяет считывать
//данные постепенно.

//В тестовом файле, который вы можете скачать из нашего репозитория на github.com, содержится длинный ряд чисел,
//разделенных символом ";". Требуется найти, на какой позиции находится число 0 и указать её в качестве ответа.
//Требуется вывести именно позицию числа, а не индекс (то-есть порядковый номер, нумерация с 1).

//Например:  12;234;6;0;78 :
//Правильный ответ будет порядковый номер числа: 4

package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strings"
)

func readCsvFile(filePath string) []string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	arrayStrings := strings.Split(records[0][0], ";")

	return arrayStrings
}

func searchNull(data []string) int {
	for i, el := range data {
		if el == "0" {
			return i + 1
		}
	}
	return -1
}

func main() {
	records := readCsvFile("C:\\Users\\PogorelovJS\\Desktop\\Golang\\stepik-go-master\\work_with_files\\task_sep_1\\task.txt")
	fmt.Println(searchNull(records))
}
