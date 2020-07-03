package main

import(
	"fmt"
	"os"
	"encoding/json"
)
var colorReset = "\033[0m"
var colorRed = "\033[31m"
var colorGreen = "\033[32m"
var colorYellow = "\033[33m"
var colorBlue = "\033[34m"
var colorPurple = "\033[35m"
var colorCyan = "\033[36m"
var colorWhite = "\033[37m"

type Configuration struct {
	Namespace string

}
func printMenu() {
	fmt.Printf("Welcom on auto diag tools\n")
	fmt.Printf("Which test could you init?\n")
	fmt.Printf("1- full test (node,pod,deploy,service and ingress)\n")
	fmt.Printf("2- node test \n")
	fmt.Printf("3- pod test\n")
	fmt.Printf("4- deploy test\n")
	fmt.Printf("5- service and ingress  test\n")
}
func loadConfig() {
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	if err != nil {
		  fmt.Println("error:", err)
	}
	
	fmt.Println(string(colorYellow)+"Configuration loaded"+string(colorReset))
	fmt.Println(string(colorGreen)+"[DEBUG]"+configuration.Namespace+string(colorReset))

}
func main () {
	loadConfig()
	printMenu()
}