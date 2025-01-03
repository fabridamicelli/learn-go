package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func processOnes(lines chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Create("./ones.txt")
	if err != nil {
		panic(err)
	}
	for line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			panic(err)
		}
	}

}

func processTwos(lines chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Create("./twos.txt")
	if err != nil {
		panic(err)
	}
	for line := range lines {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			panic(err)
		}
	}

}

func main() {
	file, err := os.Open("./data.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	ones := make(chan string)
	twos := make(chan string)

	var wg sync.WaitGroup
	wg.Add(2)
	go processOnes(ones, &wg)
	go processTwos(twos, &wg)

	s := bufio.NewScanner(file)
	for s.Scan() {
		line := s.Text()
		switch string(line[0]) {
		case "1":
			ones <- line
		case "2":
			twos <- line
		default:
			fmt.Println("dont know")
		}

	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	close(ones)
	close(twos)
	wg.Wait()
}
