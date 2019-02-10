package main

import (
	"flag"
	"github.com/labstack/echo"
	"log"
	"math/rand"
	"os"
	"time"

	"DuyStifler/GolangServer/keys"
	"github.com/google/logger"
)

func init()  {
	//get all param when run server
	flag.Parse()
	//create seed for random
	rand.Seed(time.Now().UnixNano())
}

type Handler struct {
	name string
}

func main() {
	testHandler := Handler{"Duy"}

	e := echo.New()
	echo.Context.Set("handler", Handler{})
}

func main2() {
	loggerFile, err := setupLogger()
	if err != nil {
		return
	}
	defer loggerFile.Close()
	defer logger.Init("LoggerServer", true, true, loggerFile).Close()
}

func setupLogger() (*os.File, error) {
	//setup logger
	logger.SetFlags(log.Llongfile)
	loggerFile, err := os.OpenFile(keys.LOGGER_FILE_PATH, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		logger.Error("[ERROR - main] create logger file")
		return nil, err
	}


	return loggerFile, nil
}
