package main

import (
	"DuyStifler/GolangServer/models"
	"encoding/json"
	"fmt"
)

func main() {
	serverConfig := models.ServerConfig{}
	str, err := json.Marshal(serverConfig)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(str))
	}
}