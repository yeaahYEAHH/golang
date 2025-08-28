package main

import (
	"fmt"
	"strings"
)

var queryList = map[string]bool{
	"DELETE": true,
	"CREATE": true,
	"SELECT_ALL": true,
	"SEARCH": true,
	"INSERT": true,
	"REMOVE": true,
	"UPDATE": true,
}

func parseQuery(query string) (string, string, []string, error) {
	query = strings.TrimSpace(query)
	parts := strings.Fields(query)

	if len(parts) == 0 {
		return "", "", nil, fmt.Errorf("пустой запрос")
	}

	action := strings.ToUpper(parts[0])
	var table string
	var args []string

	switch action {
	case "DELETE":
		if len(parts) < 2 {
			return "", "", nil, fmt.Errorf("ожидалось имя таблицы для %s", action)
		}
		table = parts[1]
	case "CREATE":
		if len(parts) < 3 {
			return "", "", nil, fmt.Errorf("ожидалось имя таблицы и минимум одно поле для %s", action)
		}
		table = parts[1]
		args = parts[2:]
	case "SELECTALL":
		if len(parts) < 2 {
			return "", "", nil, fmt.Errorf("ожидалось имя таблицы для %s", action)
		}
		table = parts[1]
		if len(parts) > 2 {
			args = parts[2:]
		}
	case "SEARCH", "REMOVE":	
		if len(parts) != 3{
			return "", "", nil, fmt.Errorf("неправильное количество агрументов для %s", action)
		}
		table = parts[1]
		args = parts[2:]
	case "INSERT", "UPDATE":
		if len(parts) < 3{
			return "", "", nil, fmt.Errorf("неправильное количество агрументов для %s", action)
		}
		table = parts[1]
		args = parts[2:]
	default:
		return "", "", nil, fmt.Errorf("неизвестная команда: %s", action)
	}

	return action, table, args, nil
}

func handleQuery(requestType string, table string, arguments []string) (string, error) {
	switch requestType {
	case "DELETE":
		if err := deleteTable(table); 
			err != nil {
				return  "", err
		}
	case "CREATE":
		if err := createTable(table, arguments); 
			err != nil {
				return "", err
		}		
	case "SEARCH":
		if search, err := selectRecord(table, arguments[0]); 
			err != nil {
				return "", err
		}else {
			return search, nil
		}

	case "SELECTALL":
		if search, err := selectAll(table); 
			err != nil {
				return "", err
		}else {
			return search, nil
		}

	case "INSERT":
		if err := insertRecord(table, arguments);
			err != nil{
				return "", err
		}

	case "REMOVE":
		if err := removeRecord(table, arguments[0]);
			err != nil{
				return "", err
		}

	case "UPDATE":
		if err := updateRecord(table, arguments[0], arguments[1:]);
			err != nil{
			return "", err
		}
	}

	return "", nil
}