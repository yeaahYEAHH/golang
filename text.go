package main

import (
	"bufio"
	"fmt"
	"io"
	"os"

	"time"
)

const (
	DATA = "data.txt"
	LINES = "lines.txt"
	OUTPUT = "output.txt"
	LOG = "log.txt"
	FILE1 = "filename1"
	FILE2 = "filename2"
	FILE3 = "filename3"
)

var FILES = []string{FILE1, FILE2, FILE3}

func task1(filename string) string {

	file, err := os.Open(filename)
	if err != nil {
		return "Ошибка при открытии файла"
	}
	defer file.Close()

	content, err := io.ReadAll(file)

	if err != nil {
		return "Ошибка при чтение файла"
	}

	return string(content)
}

func task2(filename string) string{
	var result string

	file, err := os.Open(filename)
	if err != nil {
		return "Ошибка при открытии файла"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var i int = 1
	for scanner.Scan() {
		result += fmt.Sprintf("%d: %s\n", i, scanner.Text())
		i++
	}

	return result
}

func task3(filename string) string{
	var result string

	file, err := os.Create(filename)
	if err != nil {
		return "Ошибка при открытии файла"
	}
	defer file.Close()

	for i := 1; i <= 20; i++ {
		if i == 20 {
			result += fmt.Sprintf("%d. строчки строченьки добавились!\n", i)
			break
		}
		result += fmt.Sprintf("%d. строчка строчеьнка добавься\n", i)
	}

	file.WriteString(result)

	return result + "\nДобавлено в файл"
}

func task4(filename string) string{
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return "Ошибка при открытии файла"
	}
	defer file.Close()

	now := time.Now().UTC().Local() // можно просто time.Now()
   formatted := now.Format("2006.01.02 15:04:05")
   write := fmt.Sprintf("%s: строченька была записана\n", formatted)

	file.WriteString(write)

	return write + "Была записана"
}

func task5(filename string) string{
	var result string

	info, err := os.Stat(filename)
	if err != nil {
		return "Ошибка чтение файла"
	}

	result = fmt.Sprintf("%.2f\tbytes\n", float64(info.Size()))
	result += fmt.Sprintf("%.2f\tKbytes\n", float64(info.Size()) / 1024)
	result += fmt.Sprintf("%.2f\tMbytes\n", float64(info.Size()) / (1024*1024))

	return result
}

func task6(filenames ...string) string{
	fileOutput, err := os.OpenFile("output.txt",  os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return "Ошибка чтение файла: output.txt"
	}
	defer fileOutput.Close()

	for _, filename := range filenames {
		file, err := os.ReadFile(filename)
		if err != nil {
			return "Ошибка чтение файла: " + filename + err.Error()
		}

		_, err = fileOutput.WriteString(string(file) + "\n")
		if err != nil {
			return "Ошибка записи в файл: output.txt " + err.Error()
		}
	}

	
	return "Контент успешно добавлен в output.txt"
}

func main() {
	fmt.Println("task #1\n" + task1(DATA) + "\n")
	fmt.Println("task #2\n" + task2(LINES) + "\n")
	fmt.Println("task #3\n" + task3(OUTPUT) + "\n")
	fmt.Println("task #4\n" + task4(LOG) + "\n")
	fmt.Println("task #5\n" + task5(DATA) + "\n")
	fmt.Println("task #6\n" + task6(FILES...))
}