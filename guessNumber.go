package main

import (
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/widget"
    "fmt"
    "math/rand"
    "strconv"
)

const (
    MAX = 300
    ATTEMPT = 10
)

func main() {
    myApp := app.New()
    myWindow := myApp.NewWindow("GuessNumber!")

    var random, attempt int

    computerGuessLabel := widget.NewLabel("Нажмите кнопку ниже, чтобы компьютер загадал число между 0 и 300!")

    startGameButton := widget.NewButton("Загадать новое число!", func() {
        random = rand.Intn(MAX)
        attempt = ATTEMPT
        fmt.Println(random)
    })


    guessDisplay := widget.NewLabel("Введите ваше первое число!")

    userGuessInput := widget.NewEntry()
    userGuessInput.SetPlaceHolder("Введите ваше число: ")

    tryGuessButton := widget.NewButton("Попробовать", func() {
        
        var str string
        userInput, err := strconv.Atoi(userGuessInput.Text)

        if err != nil {
            guessDisplay.SetText("Ошибка")
        }

        if userInput <= 0 && userInput >= MAX {
            guessDisplay.SetText("Неправильное число! Введите число между 0 и 300")
        }

        switch {
            
            case userInput > random:
                str = fmt.Sprintf("Загаданное число меньше, Попытки: %d", attempt)
                attempt--
            case userInput < random:
                str = fmt.Sprintf("Загаданное число больше, Попытки: %d", attempt)
                attempt--
            case userInput == random:
                str = fmt.Sprintf("Вы выиграли! :) Попытки: %d", attempt)        
            case attempt >= 0:
                str = fmt.Sprintf("Вы проиграли! :( за количество попыток: %d", attempt)
        }

        guessDisplay.SetText(str)
    })

    myWindow.SetContent(
        container.NewVBox(
            computerGuessLabel,
            startGameButton,
            guessDisplay,
            userGuessInput,
            tryGuessButton,
        ),
    )

    myWindow.Resize(fyne.NewSize(300, 200))
    myWindow.ShowAndRun()
}
