package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	for {
		fmt.Printf("SQL:\\ ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		
		str := scanner.Text()
		
		action, table, args, err := parseQuery(str)
		if err != nil {
			fmt.Println("ошибка: %w", err)
		}
		
		msg, err := handleQuery(action, table, args)
		if err != nil {
			fmt.Println("ошибка: %w", err)
		}

		if msg == "" {
			fmt.Printf("успешно: %s %s\n", action, table)
		}else {
			fmt.Println(msg)
		}
	}
}