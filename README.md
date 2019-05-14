# Coding Exercise by Wenzel Lowe

## Coding Problem

### Original

Please implement the following stories.

- Read in the included file named customer_purchases.csv
- Find the average number of items purchased per customer
- Write the results to a new CSV file.

### Assumptions

These are the following assumptions made:

- The output file should:
  - Have unique users as the first column (not like the original input CSV where there can be different Orders per person)
  - Other columns has total number of orders and total number of items
  - Final column is the requested average of items per order of that specific customer (not average number of items bought per order of ALL customers)
- The input and output file names should be easily changeable
- The input should scale accordingly and still be performant

## Wenzel's Implementation in Go

### Concept

The program takes in the CSV file. A map of all possible customers is created, and per line of the CSV, creates a new Customer instance. If the Customer is already in the map, then we update the map with the number of orders, items, and the average.

### Running the code

Run this code with Go (`go1.11.5 darwin/amd64`). Ensure the proper file is in the same directory `customer_purchases.csv`.

`go run main.go`

### Insights

The complexity of this solution is initially O(n+c), linear or O(n), where n is the number of orders in the CSV, and c is the number of customers.

If I had more time on this, I would add a whole benchmark test suite to evaluate performance, including other inputs to run against.
