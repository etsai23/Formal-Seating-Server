package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

var table = 1
var assignment = ""
var waiterTable = 1

type Person struct {
	Name       string `json:"name`
	Assignment string `json:"assignment"`
}

func contains(s []int, e int) bool {
	//searches for an integer through a slice. Used to check if a table is already filled before assigning.
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

var seating = []Person{}
var seatingJson = []byte{}

func handlerPer(w http.ResponseWriter, r *http.Request) {
	// info
	fmt.Println("Handler for People - Incoming Request: ")
	fmt.Println("Method: ", r.Method, " ", r.URL)

	fmt.Println("Name", seating)

	name := strings.Split(r.URL.Path, "/")[2]
	//var error: error
	name, _ = url.PathUnescape(name)

	var currentTable = ""
	//seperate name from URL - "Ethan Ng"

	//search through array "seating" and find struct with firstname Ethan and lastname Ng
	for index, _ := range seating {
		if seating[index].Name == name {
			currentTable = seating[index].Assignment
			break
		}
	}

	//save corresponding assignment to a variable
	//search through array "seating" and find any structs with assignment of that variable
	var currentSeating = []string{}

	for index, _ := range seating {
		if seating[index].Assignment == currentTable {
			currentSeating = append(currentSeating, seating[index].Name)
		} else if index == 290 {
			break
		}
	}

	// var finalArray = [][]string{}
	// finalArray = append.(finalArray, currentSeating)

	//append each firstname and lastname into a new array, and print that using the bottom line of code
	seatingJson, _ = json.Marshal(currentSeating)

	tableJSON := "{\"Table\": \"" + currentTable + "\", \"Names\": " + string(seatingJson) + "}"

	// Answer the Client request
	fmt.Fprintf(w, tableJSON)
}

func handlerTab(w http.ResponseWriter, r *http.Request) {
	// Let's print the info
	fmt.Println("Hander Incoming Request: ")
	fmt.Println("Method: ", r.Method, " ", r.URL)

	table := strings.Split(r.URL.Path, "/")[2]
	table, _ = url.PathUnescape(table)

	//seperate assignment from URL - "2"
	var currentSeating = []string{}

	//search through array "seating" and find any structs with same assignment
	for index, _ := range seating {
		if seating[index].Assignment == table {
			currentSeating = append(currentSeating, seating[index].Name)
		} else if index == 290 {
			break
		}
	}

	//append each firstname and lastname into a new array, and print usung the bottom line of code
	seatingJson, _ = json.Marshal(currentSeating)

	tableJSON := "{\"Table\": \"" + table + "\", \"Names\": " + string(seatingJson) + "}"

	// Answer the Client request
	fmt.Fprintf(w, tableJSON)
}

func findTable() {
	//find a new table assignment number
	table = rand.Intn(32)
	//gets rid of table 0
	table++

}

func main() {
	//seeding so that it is random every time
	rand.Seed(time.Now().UnixNano())

	//tableFill keeps track of how full each table is
	tableFill := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	//usedTables keeps track of tables that are full
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

		//picks a random number to assign the current person to a table
		findTable()

		for {
			//continues to find a new number until it finds a table or assignment that isn't filled up yet.
			if contains(usedTables, table) {
				findTable()
			} else {
				break
			}
		}

		if table < 32 {
			//continues to fill table until there are 8 people.
			if tableFill[table] < 9 {
				tableFill[table]++
			} else {
				tableFill[table]++
				//once 8 people are at the table, adds one more and adds the table to usedTables to ignore.
				usedTables = append(usedTables, table)
			}

			//assign the first person at each table to be the waiter at that table
			if tableFill[table] == 1 {
				assignment = "Waiter - " + strconv.Itoa(table)
			} else {
				assignment = strconv.Itoa(table)
			}

		} else if table == 32 {
			//table 32 is kitchen crew, so same deal as above but with 6 people instead of 8.
			if tableFill[table] < 6 {
				tableFill[table]++
			} else {
				tableFill[table]++
				//again, once hits 6 then add one more and close it off.
				usedTables = append(usedTables, table)
			}
			assignment = "Kitchen Crew"
		}

		seating = append(seating, Person{
			Name:       line[1] + " " + line[0],
			Assignment: assignment,
		})

	}
	http.HandleFunc("/Person/", handlerPer)
	http.HandleFunc("/Table/", handlerTab)
	log.Fatal(http.ListenAndServe(":80", nil))
}
