package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var data [][]string
var ID =  make(map[int]bool)

func readTable(tableName string) error{ 	
	var fileName string = fmt.Sprintf("%s.csv", tableName);

	file, err := os.OpenFile(fileName, os.O_RDONLY, 0644)
	if err != nil {
		return fmt.Errorf("не существует таблицы %s \nДоп.: %w", tableName, err)
	}
	defer file.Close()

	data, err = csv.NewReader(file).ReadAll()
	if err != nil{
		return  fmt.Errorf("проверьте формат таблицы %s \nДоп.: %w",  tableName, err)
	}

	return nil
}

func writeTable(tableName string) error{
	var fileName string = fmt.Sprintf("%s.csv", tableName);

	file, err := os.OpenFile(fileName, os.O_TRUNC | os.O_RDWR, 0644)
	if err != nil {
		return fmt.Errorf("не существует таблицы %s \nДоп.: %w", tableName, err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err != writer.WriteAll(data){
		return  fmt.Errorf("не удалось записать в таблицу %s\nДоп.: %w",tableName, err)
	}

	return nil	
}

func createTable(tableName string, fieldNames []string) error {
	var fileName string = fmt.Sprintf("%s.csv", tableName);

	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("не удалось создать таблицу TABLE=%s\nДоп.:%v",tableName, err)
	}
	defer file.Close()

	data = append(data, append([]string{"ID"}, fieldNames...))

	if err = writeTable(tableName); 
		err != nil{
		return err
	}

	return nil
}

func deleteTable(tableName string) error {
	var fileName string = fmt.Sprintf("%s.csv", tableName);

	if err := os.Remove(fileName); 
		err != nil {
			return fmt.Errorf("не удалось удалить таблицу %s\nДоп.:%w", tableName, err)
	}

	data = data[:0]
	
	return nil
}

func selectRecord(tableName string, id string) (string, error) {
	err := readTable(tableName)
	if err != nil {
		return "", err
	}

	head := data[0]
	for i, _ := range data {
		if data[i][0] == id {
			return strings.Join(head, "|") + "\n" + strings.Join(data[i], "|"), nil
		}
	}

	return "not found", fmt.Errorf("не удалось найти запись ID=%s TABLE=%s", id, tableName)
}


func insertRecord(tableName string, fieldValues []string) error {
	id, err := getNextID(tableName)
	if err != nil {
		return err
	}

	var idStr = strconv.Itoa(id)
	data = append(data, append([]string{idStr}, fieldValues...))

	if err := writeTable(tableName);
		err != nil {
		return err
	}

	return nil
}

func updateRecord(tableName string, id string, fieldValues []string) error {
	err := readTable(tableName)
	if err != nil {
		return err
	}

	for i := 0; i < len(data); i++ {
		if data[i][0] == id {
			data[i] = append([]string{id}, fieldValues...)

			if err := writeTable(tableName); err != nil {
				return fmt.Errorf("не удалось сохранить изменения: %w", err)
			}

			return nil
		}
	}

	return fmt.Errorf("не удалось обновить запись ID=%s TABLE=%s", id, tableName)
}


func removeRecord(tableName string, id string) error {
	err := readTable(tableName)
	if err != nil {
		return err
	}

	for i := 1; i < len(data); i++ {
		if data[i][0] == id {
			data = append(data[:i], data[i+1:]...)
			if err := writeTable(tableName); err != nil {
				return fmt.Errorf("не удалось сохранить изменения: %w", err)
			}

			return nil
		}
	}

	return fmt.Errorf("не удалось удалить запись ID=%s TABLE=%s", id, tableName)
}

func selectAll(tableName string) (string, error) {
	var result []string

	err := readTable(tableName)
	if err != nil {
		return "", err
	}

	for i := 0; i < len(data); i++ {
		result = append(result, strings.Join(data[i], "|"))
	}

	return strings.Join(result, "\n"), nil
}


func getNextID(tableName string) (int, error) {
	if err := readTable(tableName);
		err != nil {
		return -1, err
	}

	if len(ID) == 0 {
		for i, row := range data {
			if i == 0 {
				continue
			}
			
			id, _ := strconv.Atoi(row[0])
			if !ID[id] {
				ID[id] = true
			}
		}
	}

	for id := 1; ;id++ {
		if !ID[id] {
			return id, nil
		}
	}
}