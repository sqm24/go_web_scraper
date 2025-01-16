package main

import (
	"fmt"
	"bufio"
	"os"
	"net/http"
	"io/ioutil"
	"log"
	"time"
	"strings"
	"regexp"
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

	var re = regexp.MustCompile(`href="(https?://[^"]+)"`)
	html_str := string(html)
	matches := re.FindAllStringSubmatch(html_str, -1)

	for _, match := range matches {
		fmt.Println("Found link:", match[1])
	}
}

// asks user what URL you want to test
func user_input() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the website you want to parse: ")

	rawURL, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input: ", err)
		return "ERROR IDIOT"
	}

	cleanURL := strings.ReplaceAll(rawURL, "\n", "")
	cleanURL = strings.ReplaceAll(cleanURL, "\r", "")
	cleanURL = strings.ReplaceAll(cleanURL, "\t", "")

	return cleanURL
}
