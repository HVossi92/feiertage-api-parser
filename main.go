package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	if err := run(); err != nil {
		log.Fatalln(err)
	}
}

type CliConfig struct {
	startYear int
	endYear   int
}

type Holiday struct {
	Name string
	Date string `json:"datum"`
	Info string `json:"hinweis"`
}

func run() error {
	var cliConfig CliConfig
	flag.IntVar(&cliConfig.startYear, "start", 2021, "start year")
	flag.IntVar(&cliConfig.endYear, "end", 2021, "end year")
	flag.Parse()

	url := fmt.Sprint("https://feiertage-api.de/api/?jahr=", cliConfig.startYear)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {
		decoder := json.NewDecoder(resp.Body)

		var data map[string]interface{}
		err = decoder.Decode(&data)
		if err != nil {
			return err
		}
		jsonData, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			return err
		}

		os.WriteFile("holidays-"+fmt.Sprint(2019)+".json", jsonData, 0644)
		file, err := os.Open("holidays-" + fmt.Sprint(2019) + ".json")
		if err != nil {
			return err
		}
		var lines []string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := strings.Trim(scanner.Text(), " ")
			line = strings.ReplaceAll(line, "\"", "")
			if len(line) > 7 && !strings.Contains(line, "NATIONAL") {
				lines = append(lines, line)
			}
		}

		var holidays []Holiday
		for i, line := range lines {
			if strings.Contains(line, "{") {
				var holiday Holiday
				holiday.Name = strings.Split(strings.ReplaceAll(line, "{", ""), ":")[0]
				holiday.Date = strings.Trim(strings.ReplaceAll(strings.Split(lines[i+1], ":")[1], ",", ""), " ")
				holiday.Info = strings.Trim(strings.Split(lines[i+2], ":")[1], " ")
				if !holidaysContainsDate(holidays, holiday.Date) {
					holidays = append(holidays, holiday)
				}
			}
		}

		json, err := json.MarshalIndent(holidays, "", "  ")
		if err != nil {
			return err
		}
		os.WriteFile("parsed-holidays.json", []byte(json), 0644)
		fmt.Println("Done")
	}

	return nil
}

func holidaysContainsDate(holidays []Holiday, date string) bool {
	for _, holiday := range holidays {
		if holiday.Date == date {
			return true
		}
	}
	return false
}
