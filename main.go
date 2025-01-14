package main

import (
	"fmt"
	"bufio"
	"os"
	"net/http"
	"io/ioutil"
	"log"
	"time"
)

func main(){
	url := user_input()
	if  len(url) == 0 {
		fmt.Println("URL is empty. Exiting...")
		os.Exit(1);
	}
	go get_website_info(url)
	time.Sleep(2 * time.Second)
}


func get_website_info(url string){
	resp,err := http.Get(url)
	if err != nil {
	  log.Fatal(err)
	}

	defer resp.Body.Close()

	html,err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(html))
}

// asks user what URL you want to test
func user_input() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the website you want to parse: ")

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input: ", err)
		return "ERROR IDIOT"
	}

	return input
}
