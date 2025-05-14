package main

import "fmt"

func show(m map[string]float64){
    var i int = 1
    for key, _ := range m {
        fmt.Printf("%d. %s\n", i, key)
        i++
    }
}

func main() {
    var valute = make(map[int]string)
    var rates = map[string]float64{
            "USD": 1.0,   // Доллар США
            "EUR": 0.92,  // Евро
            "RUB": 90.0,  // Российский рубль
            "JPY": 157.0, // Японская иена
            "CNY": 7.25,  // Китайский юань
            "GBP": 0.78,  // Британский фунт
            "KZT": 460.0, // Казахстанский тенге
            "TRY": 32.5,  // Турецкая лира
            "INR": 83.0,  // Индийская рупия
            "BRL": 5.12,  // Бразильский реал
            "AUD": 1.50,  // Австралийский доллар
            "CAD": 1.36,  // Канадский доллар
            "CHF": 0.89,  // Швейцарский франк
            "SEK": 10.8,  // Шведская крона
            "NOK": 10.5,  // Норвежская крона
    }


    fmt.Println("Добро пожаловать в конвертор валют")
    fmt.Println("Список доступных валют для конвертации")

    var i int = 1
    
    for key, _ := range rates {
        fmt.Printf("%d %s\n", i, key)
        valute[i] = key
        i++
    }


    for {
        var money, rate int

        fmt.Printf("Введите сумму в USD: ")
        fmt.Scan(&money)
        
        if money < 0 {
            fmt.Println("Сумма должна превыщать 0")
        }else {
            fmt.Println("Введите номер для конвертации валют из списка выше:")
            fmt.Scan(&rate)

             if rates[valute[rate]] != 0 {
                fmt.Println(rates[valute[rate]] * float64(money))
            }else {
                fmt.Println("Неправильный выбор валют")
            }
        }
    }
}
