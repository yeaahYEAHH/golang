package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

const (
	PEOPLE = "people.csv"
	PRODUCTS = "products.csv"
	RECORDS = "records.csv"
	TRANSACTIONS = "transactions.csv"
	FILE1 = "file1.csv"
	FILE2 = "file2.csv"
	MULTIPLICATION = "multiplication_table.csv"
	EMPLOYEES = "employees.csv"
)

var FILES = []string{FILE1,FILE2}

func task1(filename string) (string, error){
	var result string
	file, err := os.Open(filename)
	if err != nil {
		return "", fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	
	_, err = reader.Read()
	if err != nil {
		return "",fmt.Errorf("ошибка чтения заголовка: %w", err)
	}

	records, err := reader.ReadAll()
	if err != nil {
		return "",fmt.Errorf("ошибка чтения CSV: %w", err)
	}

	for _, record := range records {
		result += fmt.Sprintf("Имя: %s, Возраст: %s, Город: %s\n", record[0], record[1], record[2])
	}

	return result, nil
}

func task2(filename string) error{
	var data = [][]string {
		{"Имя", "Возраст", "Город"},
		{"Анна", "28", "Москва"},
		{"Дмитрий", "35", "Санкт-Петербург"},
		{"Елена", "22", "Новосибирск"},
		{"Артём", "31", "Екатеринбург"},
		{"Мария", "40", "Казань"},
	}
	
	file, err := os.Create(filename)

	if err != nil {
		return fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	defer writer.Flush() 
	
	if err := writer.WriteAll(data); err != nil {
		return fmt.Errorf("ошибка записи CSV: %w", err)
	}

	return nil
}

func task3(filename string) error {
	var data = [][]string {
		{"Мария", "22", "Киев"},
		{"Алексей", "35", "Лондон"},
		{"Светлана", "29", "Бристоль"},
	}

	file, err := os.OpenFile(filename, os.O_WRONLY | os.O_APPEND, 0644 )
	if err != nil {
		return fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	if err := writer.WriteAll(data); err != nil {
		return fmt.Errorf("ошибка записи CSV: %w", err)
	}
	defer writer.Flush()
	
	return nil
}

func task4(filename string) (int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return 0, fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	_, err = reader.Read()
	if err != nil {
		return 0, fmt.Errorf("ошибка чтения заголовка: %w", err)
	}

	records, err := reader.ReadAll()
	if err != nil {
		return 0, fmt.Errorf("ошибка чтения CSV: %w", err)
	}

	return len(records), nil
}

func task5(filenames ...string) error {
	fileMerged, err := os.OpenFile("merged.csv", os.O_RDONLY | os.O_WRONLY | os.O_APPEND, 0644)
	if err != nil {
		return fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer fileMerged.Close()

	writerMerged := csv.NewWriter(fileMerged)
	defer writerMerged.Flush()

	for _, filename := range filenames {
		file, err := os.Open(filename)
		if err != nil {
			return fmt.Errorf("ошибка открытия файла: %w", err)
		}
		defer file.Close()

		reader := csv.NewReader(file)
		_, err = reader.Read()
		if err != nil {
        return fmt.Errorf("ошибка чтения заголовка: %w", err)
    	}

		records, err := reader.ReadAll()

		if err := writerMerged.WriteAll(records); err != nil {
			return fmt.Errorf("ошибка записи CSV: %w", err)
		}
	}

	return nil
}

func task6(filename string, n int) error{
	multi := func () [][]string {
		var result [][]string
		for i := 1; i <= n; i++ {
			var row []string
			for j := 1; j <= n; j++ {
				str := fmt.Sprintf("%d", i * j)
				row = append(row, str)
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

	writer := csv.NewWriter(file)
	if err := writer.WriteAll(multi()); err != nil {
		return fmt.Errorf("ошибка записи CSV: %w", err)
	}

	return nil
}

func task7(filename string, search string) (string, error) {
	var result string

	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		return "", fmt.Errorf("ошибка открытия файла: %w", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	head, err := reader.Read()
	if err != nil {
		return "", fmt.Errorf("ошибка чтения заголовка: %w", err)
	}

	recordsArray, err := reader.ReadAll()

	var index int
	var found bool
	for i, records := range recordsArray {
		for _, record := range records {
			if record == search {
				index, found = i, true
				break
			}
		}
	}

	fmt.Println(head, index)
	if !found {
		return "", nil
	}

	for key, field := range head  {
		result += fmt.Sprintf("%s: %s, ", field, recordsArray[index][key])
	}

	return result, nil
}

func main() {
	if msg, err := task1(PEOPLE); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Задача #1: Чтение CSV файла")
		fmt.Println(msg)
	}

	if err := task2(PRODUCTS); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Задача #2: Запись данных в новый CSV файл")
	}

	if err := task3(RECORDS); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Задача #3: Добавление данных в CSV файл")
	}
	
	if msg, err := task4(TRANSACTIONS); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Задача #4: Подсчёт строк в CSV файле \nCOUNT: ", msg)
	}

	if err := task5("file1.csv", "file2.csv"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Задача #5: Объединение двух CSV файлов")
	}

	if err := task6(MULTIPLICATION, 10); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Задача #6: Генерация CSV файла с таблицей умножения")
	}

	if msg, err := task7(EMPLOYEES, "Инженер"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Задача #7: Поиск данных в CSV файле \nSearch", msg)
	}
}