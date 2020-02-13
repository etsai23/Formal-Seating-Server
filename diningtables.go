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
		Table3    []int
	}

	//array to track empty and full tables
	tableFill := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	//used table array
	usedTables := []int{}
	//variable that helps define "person"
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

	//random seed
	rand.Seed(time.Now().UnixNano())

	for studentNum, student := range person {
		//random table number
		var table = rand.Intn(32)
		var table2 = rand.Intn(32)
		var table3 = rand.Intn(32)

		//table + 1, for each table seating number (1, 2, and 3)
		table++
		table2++
		table3++
		person[studentNum].Table = append(person[studentNum].Table, table)
		person[studentNum].Table2 = append(person[studentNum].Table2, table2)
		person[studentNum].Table3 = append(person[studentNum].Table3, table3)

		//merely here to declare "student" before it is used below
		if table == 34 {
			fmt.Println(student)
		}
		if table == 32 {
			//if the group is full, next table
			if tableFill[table] < 9 {
				//if Table 1 = Table 2, then find another table
				if person[studentNum].Table[0] == person[studentNum].Table2[0] {
					person[studentNum].Table2 = append(person[studentNum].Table2, rand.Intn(32))
					//gets rid of first element in array
					person[studentNum].Table2 = append(person[studentNum].Table2[:0], person[studentNum].Table2[1:]...)

				} else if person[studentNum].Table[0] == person[studentNum].Table3[0] {
					//if table 1 is the same as table 3, remove the element from the array
					person[studentNum].Table3 = append(person[studentNum].Table3, rand.Intn(32))
					person[studentNum].Table3 = append(person[studentNum].Table3[:0], person[studentNum].Table3[1:]...)

				} else if person[studentNum].Table2[0] == person[studentNum].Table3[0] {
					//if table 2 is the same as table 3, remove the element from the array
					person[studentNum].Table3 = append(person[studentNum].Table3, rand.Intn(32))
					person[studentNum].Table3 = append(person[studentNum].Table3[:0], person[studentNum].Table3[1:]...)

				} else if person[studentNum].Table[0] == 32 || person[studentNum].Table2[0] == 32 || person[studentNum].Table3[0] == 32 {
					fmt.Println(person[studentNum], "Kitchen Crew")
					tableFill[table]++
				} else {
					fmt.Println(person[studentNum])
					tableFill[table]++
				}
			} else {
				//if Table 1 = Table 2, then reshuffle
				if person[studentNum].Table[0] == person[studentNum].Table2[0] {
					person[studentNum].Table2 = append(person[studentNum].Table2, rand.Intn(32))
					//if table 1 is the same as table 2, remove the first element in the array
					person[studentNum].Table2 = append(person[studentNum].Table2[:0], person[studentNum].Table2[1:]...)
					//if table2 = 32, which is the kitchen crew table, assign as Kitchen Crew

				} else if person[studentNum].Table[0] == person[studentNum].Table3[0] {
					//if table 1 is the same as table 3, remove the element from the array
					person[studentNum].Table3 = append(person[studentNum].Table3, rand.Intn(32))
					person[studentNum].Table3 = append(person[studentNum].Table3[:0], person[studentNum].Table3[1:]...)

				} else if person[studentNum].Table2[0] == person[studentNum].Table3[0] {
					//if table 2 is the same as table 3, remove the element from the array
					person[studentNum].Table3 = append(person[studentNum].Table3, rand.Intn(32))
					person[studentNum].Table3 = append(person[studentNum].Table3[:0], person[studentNum].Table3[1:]...)

				} else if person[studentNum].Table[0] == 32 || person[studentNum].Table2[0] == 32 || person[studentNum].Table3[0] == 32 {
					fmt.Println(person[studentNum], "Kitchen Crew")
					tableFill[table]++
				} else {
					fmt.Println(person[studentNum])
					usedTables = append(usedTables, table)
				}
			}
		} else {
			if tableFill[table] < 9 {
				if person[studentNum].Table[0] == person[studentNum].Table2[0] {
					person[studentNum].Table2 = append(person[studentNum].Table2, rand.Intn(32))
					//if table 1 is the same as table 2, remove the element from the array
					person[studentNum].Table2 = append(person[studentNum].Table2[:0], person[studentNum].Table2[1:]...)

				} else if person[studentNum].Table[0] == person[studentNum].Table3[0] {
					//if table 1 is the same as table 3, remove the element from the array
					person[studentNum].Table3 = append(person[studentNum].Table3, rand.Intn(32))
					person[studentNum].Table3 = append(person[studentNum].Table3[:0], person[studentNum].Table3[1:]...)

				} else if person[studentNum].Table2[0] == person[studentNum].Table3[0] {
					//if table 2 is the same as table 3, remove the element from the array
					person[studentNum].Table3 = append(person[studentNum].Table3, rand.Intn(32))
					person[studentNum].Table3 = append(person[studentNum].Table3[:0], person[studentNum].Table3[1:]...)

				} else if person[studentNum].Table[0] == 32 || person[studentNum].Table2[0] == 32 || person[studentNum].Table3[0] == 32 {
					fmt.Println(person[studentNum], "Kitchen Crew")
					tableFill[table]++
				} else {
					fmt.Println(person[studentNum])
					tableFill[table]++
				}
			} else {
				if person[studentNum].Table[0] == person[studentNum].Table2[0] {
					person[studentNum].Table2 = append(person[studentNum].Table2, rand.Intn(32))
					//if table 1 is the same as table 2, remove the element from the array
					person[studentNum].Table2 = append(person[studentNum].Table2[:0], person[studentNum].Table2[1:]...)

				} else if person[studentNum].Table[0] == person[studentNum].Table3[0] {
					//if table 1 is the same as table 3, remove the  element from the array
					person[studentNum].Table3 = append(person[studentNum].Table3, rand.Intn(32))
					person[studentNum].Table3 = append(person[studentNum].Table3[:0], person[studentNum].Table3[1:]...)

				} else if person[studentNum].Table2[0] == person[studentNum].Table3[0] {
					//if table 2 is the same as table 3, remove the first element in the array
					person[studentNum].Table3 = append(person[studentNum].Table3, rand.Intn(32))
					person[studentNum].Table3 = append(person[studentNum].Table3[:0], person[studentNum].Table3[1:]...)
				} else {
					fmt.Println(person[studentNum], "Waiter")
					usedTables = append(usedTables, table)
				}
			}
		}
	}
}

// fmt.Println(person)
