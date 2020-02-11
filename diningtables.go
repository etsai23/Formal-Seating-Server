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

	type Person struct {
		Firstname string
		Lastname  string
		Table     []int
		Table2    []int
	}

	//Seeding that randomized seating assignments every time
	rand.Seed(time.Now().UnixNano())

	//slice that keeps track of how many people are at each table
	tableFill := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	//slice for keeping track of which tables are full
	usedTables := []int{}

	var person []Person

	csvFile, _ := os.Open("Dining.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}
		person = append(person, Person{
			Firstname: line[0],
			Lastname:  line[1],
		})
	}

	for studentNum, student := range person {

		//print table number without 0. In order to do this, add 1 to each random int table value.
		var table = rand.Intn(33)
		var table2 = rand.Intn(33)

		if table == 35 {
			fmt.Println(student)
		}

		table++
		table2++
		person[studentNum].Table = append(person[studentNum].Table, table)
		person[studentNum].Table2 = append(person[studentNum].Table2, table2)

		if table == 32 {
			//if the table is not full (full being 8 people), print this group of eight students as kitchen group.
			if tableFill[table] < 8 {
				//if statement to prevent second seating from being the same as the first
				if person[studentNum].Table[0] == person[studentNum].Table2[0] {
					person[studentNum].Table2 = append(person[studentNum].Table2, rand.Intn(33))
					fmt.Println(person[studentNum], "Kitchen Crew")
					tableFill[table]++
				} else {
					fmt.Println(person[studentNum], "Kitchen Crew")
					tableFill[table]++
				}
			} else {
				if person[studentNum].Table[0] == person[studentNum].Table2[0] {
					person[studentNum].Table2 = append(person[studentNum].Table2, rand.Intn(33))
					fmt.Println(person[studentNum], "Kitchen Crew")
					tableFill[table]++
				} else {
					fmt.Println(person[studentNum], "Kitchen Crew")

					//remove from available tables/groups
					usedTables = append(usedTables, table)
				}

			}
		} else if table == 33 {
			//if waiter group is fewer than 30 individuals, assign student to waiter group if table = 33
			if tableFill[table] < 30 {
				if person[studentNum].Table[0] == person[studentNum].Table2[0] {
					person[studentNum].Table2 = append(person[studentNum].Table2, rand.Intn(33))
					fmt.Println(person[studentNum], "Waiter")
					tableFill[table]++
				} else {
					fmt.Println(person[studentNum], "Waiter")
					tableFill[table]++
				}
			} else {
				if person[studentNum].Table[0] == person[studentNum].Table2[0] {
					person[studentNum].Table2 = append(person[studentNum].Table2, rand.Intn(33))
					fmt.Println(person[studentNum], "Kitchen Crew")
					tableFill[table]++
				} else {
					fmt.Println(person[studentNum], "Kitchen Crew")
					//remove from availble tables/groups
					usedTables = append(usedTables, table)

				}
			}
		} else {
			//if table we are looking at has fewer than eight people, add them to the tableFill and assign them to the table
			if tableFill[table] < 8 {
				if person[studentNum].Table[0] == person[studentNum].Table2[0] {
					person[studentNum].Table2 = append(person[studentNum].Table2, rand.Intn(33))
					fmt.Println(person[studentNum])
					tableFill[table]++
				} else {
					fmt.Println(person[studentNum])
					tableFill[table]++
				}
			} else {
				if person[studentNum].Table[0] == person[studentNum].Table2[0] {
					person[studentNum].Table2 = append(person[studentNum].Table2, rand.Intn(33))
					fmt.Println(person[studentNum])
					tableFill[table]++
				} else {
					fmt.Println(person[studentNum])
					//remove from availble tables/groups
					usedTables = append(usedTables, table)

				}

			}

		}

	}
}
