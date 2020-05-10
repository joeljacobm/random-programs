package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"
	"sync"
	"time"
)

var counter int

func main() {

	problemFile := flag.String("csvfile", "problem.csv", "A csv file that has the question and answers")
	timer := flag.Int("timer", 5, "How much time do you need between each questions?")

	flag.Parse()

	file, err := os.Open(*problemFile)
	if err != nil {
		panic(err)
	}

	reader := csv.NewReader(file)

	op, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	QAMap := make(map[string]string, len(op))
	FAMap := make(map[string]string, len(op))

	for _, v := range op {

		QAMap[v[0]] = v[1]
	}

	answer := make(chan string)

	ticker := time.NewTicker(time.Duration(*timer) * time.Second)

	fmt.Println("-----------Welcome to command line quiz----------------")

	for k, _ := range QAMap {

		var ans string
		counter++
		fmt.Printf("Q%d: %s\n", counter, k)

		go func() {
			for {
				fmt.Scanf("%s", &ans)
				FAMap[k] = ans
				answer <- ans
			}
		}()

		select {
		case <-ticker.C:
			break
		case <-answer:
			break
		}
	}

	var fScore int
	for k, _ := range QAMap {
		if QAMap[k] == FAMap[k] {
			fScore++
		}

	}

	fmt.Printf("-------------Your Final Score is:%d------------- ", fScore)
}
