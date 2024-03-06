package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// структура локации
type Location struct {
	ID      string   `json:"id"`
	Text    string   `json:"text"`
	Options []Option `json:"options"`
}

// структура выбора
type Option struct {
	Text   string `json:"text"`
	NextID string `json:"nextID"`
}

func main() {
	//загрузка
	quest, err := loadQuest("quest.json")
	if err != nil {
		log.Fatal(err)
	}

	//начало
	startLocationID := "start"
	currentLocation := findLocationByID(quest, startLocationID)

	for {
		//вывод текста
		fmt.Println(currentLocation.Text)

		//вывод вариантов
		for i, option := range currentLocation.Options {
			fmt.Printf("%d. %s\n", i+1, option.Text)
		}

		//вывод от пользователя
		fmt.Print("Выберите действие: ")
		var choice int
		fmt.Scanln(&choice)

		//проверка выбора (доступности)
		if choice < 1 || choice > len(currentLocation.Options) {
			fmt.Println("Неверный выбор. Попробуйте еще раз.")
			continue
		}

		//обновление локации
		nextLocationID := currentLocation.Options[choice-1].NextID
		currentLocation = findLocationByID(quest, nextLocationID)

		//проверка на завершение
		if nextLocationID == "end" {
			fmt.Println("Отлично! Вы завершили квест.")
			break
		}
	}
}

// загрузка квеста
func loadQuest(filename string) ([]Location, error) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var quest []Location
	err = json.Unmarshal(file, &quest)
	if err != nil {
		return nil, err
	}

	return quest, nil
}

// поиск локации по id
func findLocationByID(quest []Location, id string) Location {
	for _, loc := range quest {
		if loc.ID == id {
			return loc
		}
	}
	return Location{} //возврат пустой локации
}
