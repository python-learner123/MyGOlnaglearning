package main

import (
	"log"
	"os"
)

func main() {
	//log.Println("Hello world!")  //default print to std err
	file, err := os.OpenFile("logfile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Not able to opne Log file")
	}
	log.SetOutput(file)
	defer file.Close()
	log.Println("Hello world!")
}

//op: 
//2020/07/12 11:36:57 Hello world!
