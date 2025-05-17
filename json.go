package main

import (
	"encoding/json"
	"fmt"
	"os"
	"encoding/csv"
	"io"
)

const (
	PEOPLE = "people.json"
	PRODUCTS = "products.json"
	RECORDS = "records.json"
	TRANSACTIONS = "transactions.json"
	FILE1 = "file1.json"
	FILE2 = "file2.json"
	MULTIPLICATION = "multiplication_table.json"
	EMPLOYEES = "employees.json"
	DATA = "data.json"
)

var FILES = []string{FILE1, FILE2}

func task1(filename string) (string, error){
	var result string

	file, err := os.OpenFile(filename, os.O_RDONLY, 0664)
	if err != nil {
		return "", fmt.Errorf("Ошибка чтения файла %w", err)
	}
	defer file.Close()

	var data []map[string]interface{}
	
	reader := json.NewDecoder(file)
	err = reader.Decode(&data)
	if err != nil {
		return "", fmt.Errorf("Ошибка декодирования JSON %w", err)
	}

	for _, record := range data {
		result += fmt.Sprintf("Имя: %s, Возраст: %.0f, Город: %s\n", record["name"], record["age"], record["city"])
  	}

	return result, nil
}

func task2(filename string) error{
	data := []map[string]interface{}{
		{"name": "Молоко", "quantity": 2, "price": 100},
		{"name": "Хлеб", "quantity": 1, "price": 40},
		{"name": "Сыр", "quantity": 1, "price": 300},
		{"name": "Кофе", "quantity": 3, "price": 500},
		{"name": "Шоколад", "quantity": 10, "price": 150},
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("Ошибка чтения файла %w", err)
	}
	defer file.Close()

	writer := json.NewEncoder(file)
	err = writer.Encode(data)

	if err != nil {
		return fmt.Errorf("Ошибка записи в JSON файл %w", err)
	}

	return nil
}

func task3(filename string) error{
	var data []map[string]interface{}

	file, err := os.OpenFile(filename, os.O_RDWR | os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("Ошибка чтения файла %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err = decoder.Decode(&data); err != nil {
		return fmt.Errorf("Ошибка декодирования JSON %w", err)
	}

	data = append(
		data, 
		[]map[string]interface{}{
			{"name": "Мария", "age": 22, "city": "Киев"},
			{"name": "Алексей", "age": 35, "city": "Лондон"},
			{"name": "Светлана", "age": 29, "city": "Бристоль"},
		}...
	)

	if _, err := file.Seek(0, io.SeekStart); err != nil {
		return fmt.Errorf("Ошибка перемещения курсора: %w", err)
	}
	if err := file.Truncate(0); err != nil {
		return fmt.Errorf("Ошибка очистки файла: %w", err)
	}

	if err = json.NewEncoder(file).Encode(data); err != nil {
		return fmt.Errorf("Ошибка записи в JSON файл %w", err)
  	}

	return nil
}

func task4(filename string) (int, error) {
	var data []map[string]interface{}
	
	file, err := os.OpenFile(filename, os.O_RDONLY, 0664)
	if err != nil {
		return -1, fmt.Errorf("Ошибка чтения файла %w", err)
	}
	defer file.Close()

	if err = json.NewDecoder(file).Decode(&data); err != nil {
		return -1, fmt.Errorf("Ошибка записи в JSON файл %w", err)
	}

	return len(data), nil
}

func task5(filenames ...string) error {
	var newData []map[string]interface{}

	for _, filename := range filenames {
		var data []map[string]interface{}
		file, err := os.OpenFile(filename, os.O_RDONLY, 0664)
		if err != nil {
			return fmt.Errorf("Ошибка чтения файла %w", err)
		}
		defer file.Close()

		if err = json.NewDecoder(file).Decode(&data); err != nil {
			return fmt.Errorf("Ошибка записи в JSON файл %w", err)
		}

		newData = append(newData, data...) 
	}

	file, err := os.Create("merged.json")
	if err != nil {
		return fmt.Errorf("Ошибка чтения файла %w", err)
	}
	defer file.Close()

	if err = json.NewEncoder(file).Encode(newData); err != nil {
		return fmt.Errorf("Ошибка записи в JSON файл %w", err)
	}

	return nil
}

func task6(filename string, n int) error {
	multi := func () [][]int {
		var result [][]int
		for i := 1; i <= n; i++ {
			var row []int
			for j := 1; j <= n; j++ {
				row = append(row, i*j)
			}
			result = append(result, row)
		}

		return result
	}

	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer file.Close()

	if err = json.NewEncoder(file).Encode(multi()); err != nil {
		return fmt.Errorf("Ошибка записи в JSON файл %w", err)
	}

	return nil
}

func task7(filename string, search any) (string, error) {
	var result string
	var data []map[string]interface{}

	file, err := os.OpenFile(filename, os.O_RDONLY, 0664)
	if err != nil {
		return "", fmt.Errorf("Ошибка чтения файла %w", err)
	}
	defer file.Close()

	if err = json.NewDecoder(file).Decode(&data); err != nil {
		return "", fmt.Errorf("Ошибка чтение в JSON файл %w", err)
	}

	var index int
	var found bool
	for i, row := range data {
		for _, val := range row {
			if val == search {
				index, found = i, true
				break
			}
		}	
	}

	if !found {
		return "", nil
	}

	for key, value := range data[index]  {
		fmt.Println(key, value)
		result += fmt.Sprintf("%s: %v, ", key, value)
	}

	return result, nil
}

func task8(filename string) error{
	var data []map[string]interface{}
	var table [][]string
	var head []string

	file, err := os.OpenFile(filename, os.O_RDONLY, 0664)
	if err != nil {
		return fmt.Errorf("Ошибка чтения файла %w", err)
	}
	defer file.Close()

	if err = json.NewDecoder(file).Decode(&data); err != nil {
		return fmt.Errorf("Ошибка чтение в JSON файл %w", err)
	}

	for key := range data[0] {
		head = append(head, key)
	}

	table = append(table, head)

	for _, row := range data {
		var addRow []string
		for _, val := range row {
			str := fmt.Sprintf("%v", val)
			addRow = append(addRow, str)
		}
		table = append(table, addRow)
	}

	fileCSV, err := os.Create("data.csv")
	if err != nil {
		return fmt.Errorf("Ошибка чтения файла %w", err)
	}
	defer fileCSV.Close()

	if err := csv.NewWriter(fileCSV).WriteAll(table); err != nil {
		return fmt.Errorf("Ошибка записи файла %w", err)
	}

	return nil
}

func main() {
	if msg, err := task1(PEOPLE); err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Задача #1: Чтение JSON файла")
		fmt.Println(msg)
	}

	if err := task2(PRODUCTS); err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Задача #2: Запись данных в новый JSON файл")
	}

	if err := task3(RECORDS); err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Задача #3: Добавление данных в JSON файл")
	}

	if length, err := task4(TRANSACTIONS); err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Задача #4: Подсчёт объектов в JSON файле \nКоличество транзакций:", length)
	}

	if err := task5(FILES...); err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Задача #5: Объединение двух JSON файлов")
	}

	if err := task6(MULTIPLICATION, 10); err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Задача #6: Генерация JSON файла с таблицей умножения")
	}

	if search, err := task7(EMPLOYEES, 50000.0); err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Задача #7: Поиск данных в JSON файле\n",search)
	}

	if  err := task8(DATA); err != nil {
		fmt.Println(err)
	}else{
		fmt.Println("Задача #8: Преобразование JSON в CSV")
	}
}