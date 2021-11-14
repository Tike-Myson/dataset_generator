package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"time"
)

type Player struct {
	FullName string
	PositionID int
	Shots float64
	Makes float64
	Pass float64
	Accuracy float64
	PersonalRating float64
	Goals int
	Assists int
	Fouls int
	YellowCards int
	RedCards int
	YearsPro int
	TeamValue string
}

func main() {
	matrix := getMatrix()

	csvFile, err := os.Create("csv/scoring.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter := csv.NewWriter(csvFile)

	for _, empRow := range matrix {
		_ = csvwriter.Write(empRow)
	}
	csvwriter.Flush()
	csvFile.Close()
}

func getMatrix() [][]string {
	var matrix [][]string
	arr := []string {"Player_Name", "Position_ID", "Shots", "Makes", "Pass", "Accuracy", "Personal_Rating", "Goals", "Assists", "Fouls", "Yellow_Card", "Red_Card", "Years_Pro", "Team_Value"}
	matrix = append(matrix, arr)
	var Player Player
	//teamValue := []string {"future", "rotation", "main", "superstar"}
	for i := 0; i < 95; i++ {
		arr = nil
		//полное имя игрока
		Player.FullName = GetFullName()
		arr = append(arr, Player.FullName)
		//позиция
		Player.PositionID = rand.Intn(11)
		arr = append(arr, fmt.Sprintf("%d", Player.PositionID))
		//сколько ударов наносит по воротам
		Player.Shots = GetFloat(0.5, 5.5)
		arr = append(arr, fmt.Sprintf("%f", Player.Shots))
		//сколько ударов завершились голом
		Player.Makes = GetFloat(0.2, 0.8)
		arr = append(arr, fmt.Sprintf("%f", Player.Makes))
		//сколько раз пасует игрок за игру
		Player.Pass = GetFloat(10.0, 110.0)
		arr = append(arr, fmt.Sprintf("%f", Player.Pass))
		//процент точности пасов
		Player.Accuracy = GetFloat(32.0, 99.0)
		arr = append(arr, fmt.Sprintf("%f", Player.Accuracy))
		//средняя оценка за весь сезон
		Player.PersonalRating = GetFloat(4.0, 9.9)
		arr = append(arr, fmt.Sprintf("%f", Player.PersonalRating))
		//сколько голов забил игрок
		Player.Goals = rand.Intn(25)
		arr = append(arr, fmt.Sprintf("%d", Player.Goals))
		//сколько ассистов отдал игрок
		Player.Assists = rand.Intn(25)
		arr = append(arr, fmt.Sprintf("%d", Player.Assists))
		//сколько фолов совершил игрок
		Player.Fouls = rand.Intn(70)
		arr = append(arr, fmt.Sprintf("%d", Player.Fouls))
		//сколько желтых карточек получил игрок
		Player.YellowCards = rand.Intn(15)
		arr = append(arr, fmt.Sprintf("%d", Player.YellowCards))
		//сколько красных карточек получил игрок
		Player.RedCards = rand.Intn(5)
		arr = append(arr, fmt.Sprintf("%d", Player.RedCards))
		//сколько играет игрок на профессиональном уровне
		Player.YearsPro = rand.Intn(15)
		arr = append(arr, fmt.Sprintf("%d", Player.YearsPro))
		//значение игрока для клуба
		//Player.SetTeamValue()
		//arr = append(arr, Player.TeamValue)
		matrix = append(matrix, arr)
	}
	return matrix
}

func (p *Player) SetTeamValue() {
	if p.Goals > 17 && p.Assists > 15 && p.PersonalRating > 7.0 && p.YearsPro > 2 {
		p.TeamValue = "superstar"
	} else if p.Goals > 8 && p.Goals < 17 && p.Assists > 8 && p.Assists < 15 && p.PersonalRating > 5.0 && p.PersonalRating < 7.0 {
		p.TeamValue = "main"
	} else if p.YearsPro == 0 {
		p.TeamValue = "future"
	} else {
		p.TeamValue = "rotation"
	}
}

func GetFloat(min, max float64) float64 {
	rand.Seed(time.Now().UnixNano())
	res := min + rand.Float64() * (max - min)
	return math.Ceil(res * 100)/100
}

func GetFullName() string {
	var names []string
	var surnames []string
	var kazakhSurnames []string
	var kazakhNames []string
	var russianSurnames []string
	var russianNames []string
	var englishNames []string

	file, err := os.Open("parametr/english_names.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		englishNames = append(englishNames, scanner.Text())
	}

	file, err = os.Open("parametr/kazakh_names.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		kazakhNames = append(kazakhNames, scanner.Text())
	}

	file, err = os.Open("parametr/kazakh_surnames.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		kazakhSurnames = append(kazakhSurnames, scanner.Text())
	}

	file, err = os.Open("parametr/russian_names.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		russianNames = append(russianNames, scanner.Text())
	}

	file, err = os.Open("parametr/russian_surnames.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		russianSurnames = append(russianSurnames, scanner.Text())
	}
	randomLanguage := rand.Intn(3)
	if randomLanguage == 0 {
		names = kazakhNames
		surnames = kazakhSurnames
	}
	if randomLanguage == 1 {
		names = russianNames
		surnames = russianSurnames
	}
	if randomLanguage == 2 {
		names = englishNames
		surnames = englishNames
	}
	randomNameID := rand.Intn(len(names))
	randomSurnameID := rand.Intn(len(surnames))

	return names[randomNameID] + " " + surnames[randomSurnameID]
}