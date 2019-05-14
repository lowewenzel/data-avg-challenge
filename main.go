package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

// Customer is a single customer struct
type Customer struct {
	ID          int
	TotalOrders int
	Orders      []int
	SumItems    int
	Average     float32
}

// Files contains the CSV and file structures
type Files struct {
	buffer *csv.Reader
	writer *csv.Writer
	file   *os.File
	data   *os.File
}

// Main functionality: Read CSV, calculate average, write CSV
func main() {
	files := openFiles("customer-purchases.CSV", "customers.csv")
	writer := files.writer
	file := files.file
	data := files.data
	defer file.Close()
	defer data.Close()
	defer writer.Flush()

	customers := readInput(files)

	// testPrint(customers)

	writeToCSV(customers, writer)
}

// Opens the input CSV file, preps the output CSV file
func openFiles(fileOpen string, fileWrite string) Files {
	// Read CSV input file
	data, err := os.Open(fileOpen)
	if err != nil {
		log.Fatal(err)
	}
	buffer := csv.NewReader(bufio.NewReader(data))

	// Write CSV output file
	file, err := os.Create(fileWrite)
	if err != nil {
		log.Fatal(err)
	}
	writer := csv.NewWriter(file)

	return Files{buffer, writer, file, data}
}

// This function loops through each line of input CSV and creates a slice of Customers
func readInput(files Files) map[int]Customer {
	customers := make(map[int]Customer)
	buffer := files.buffer
	for {
		line, err := buffer.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		id, err := strconv.Atoi(line[0])
		if err != nil {
			continue
		}
		intItems, _ := strconv.Atoi(line[2])
		orderID, _ := strconv.Atoi(line[1])
		floatItems := float32(intItems)

		// Create customer instance
		customer := Customer{
			id,
			1,
			[]int{orderID},
			intItems,
			floatItems,
		}

		// Update customer in slice if found, otherwise assign it
		if found, ok := customers[id]; ok {
			found.TotalOrders++
			found.Orders = append(found.Orders, orderID)
			found.SumItems += intItems
			found.Average = float32(found.SumItems) / float32(found.TotalOrders)
			customers[id] = found
		} else {
			customers[id] = customer
		}
	}
	return customers
}

// This function writes the in-memory customers data read from the CSV file to another CSV file with averages
func writeToCSV(customers map[int]Customer, writer *csv.Writer) {
	header := []string{"customer_id", "total_orders", "total_items", "average_order"}
	writer.Write(header)

	for _, c := range customers {
		writer.Write([]string{strconv.Itoa(c.ID), strconv.Itoa(c.TotalOrders), strconv.Itoa(c.SumItems), fmt.Sprintf("%f", c.Average)})
	}
}

// This function outputs the data that would be in CSV to the standard output.
func testPrint(customers map[int]Customer) {
	for _, v := range customers {
		fmt.Println("Customer", v.ID)
		fmt.Println("Orders:", v.Orders)
		fmt.Println("Total orders", v.TotalOrders)
		fmt.Println("Total items purchased", v.SumItems)
		fmt.Println("Average per purchase", v.Average)
	}
}
