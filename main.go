package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type company struct {
	name         string
	domain       string
	yearFounded  string
	currentEmpoy int
}

type readFunc func(data []string)

func readCompanies(comcompany *[]company) readFunc {
	return func(record []string) {
		cm := company{}
		cm.name = record[1]
		cm.domain = record[2]

		*comcompany = append(*comcompany, cm)
	}
}

func main() {
	// open files
	var data []company
	err := readCSV("dataset/companies_sorted.csv", 5, readCompanies(&data))
	if err != nil {
		log.Fatal(err)
	}
}

func readCSV(filesUrl string, limit int, funcRead readFunc) error {
	files, err := os.Open(filesUrl)
	if err != nil {
		return err
	}
	i := 0

	r := csv.NewReader(files)
	r.Comma = ','
	now := time.Now()

	for {
		if i == limit && limit != 0 {
			break
		}
		record, err := r.Read()
		if i == 0 {
			i++
			continue
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("got an error %d: %v", i, err)
		}
		funcRead(record)
		i++
	}
	fmt.Printf("total time %q: %d \n", filesUrl, time.Since(now).Milliseconds())
	return nil
}
