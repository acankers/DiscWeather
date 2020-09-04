package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Location struct {
	lat string
	lon string
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var loc Location

	fmt.Println("Latitude?")
	text, _ := reader.ReadString('\n')
	fmt.Println("longitude")
	text2, _ := reader.ReadString('\n')

	loc.lat = strings.TrimRight(text, "\r\n")
	loc.lon = strings.TrimRight(text2, "\r\n")

	var response = parseWeatherData(loc)

	fmt.Println(response)
}

func parseWeatherData(loc Location) string {
	var apikey = ""
	var apiURL = "https://api.openweathermap.org/data/2.5/onecall?lat=" + loc.lat + "&lon=" + loc.lon + "&exclude=&appid=" + apikey

	return apiURL
}
