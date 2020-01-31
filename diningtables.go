package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	tableFill := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	usedTables := []int{}

	csvFile, _ := os.Open("Dining.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		//print table number without 0.
		var table = rand.Intn(33)
		table++

		if table == 32 {
			if tableFill[table] < 8 {
				fmt.Println(line[0], line[1], "Kitchen Crew")
				tableFill[table]++
			} else {
				tableFill[table]++
				fmt.Println(line[0], line[1], "Kitchen Crew")
				tableFill[table]++
			}

		} else if table == 33 {
			if tableFill[table] < 30 {
				fmt.Println(line[0], line[1], "Waiter")
				tableFill[table]++
			} else {
				tableFill[table]++
				fmt.Println(line[0], line[1], "Waiter")
				tableFill[table]++
			}

		} else {
			//if table we are looking at has fewer than eight people, add them to the tableFill and assign them to the table
			if tableFill[table] < 8 {
				fmt.Println(line[0], line[1], table)
				tableFill[table]++
			} else {
				//else, add the integer assigned to the table to the usedTables slice.
				tableFill[table]++
				fmt.Println(line[0], line[1], table)
				usedTables = append(usedTables, table)
			}

		}

	}

}
