package utility

import (
	"fmt"
	"log"
	"os"
)

type CustomError struct {
	code    int
	Message string
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Error %d: %s",e.code,e.Message)
}

func NewCustomError(code int,message string) error{
	return &CustomError{
		code: code,
		Message: message,
	}
}

var(
	InfoLogger *log.Logger
	ErrorLogger *log.Logger
	DebugLogger *log.Logger
)

func InitLogger(){
	file, err := os.OpenFile("application.log", os.O_CREATE | os.O_WRONLY | os.O_APPEND,0666)
	if err !=nil{
		log.Fatal("Failed to open log file: %v",err)
	}

	InfoLogger = log.New(file,"INFO:",log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file,"ERROR:",log.Ldate|log.Ltime|log.Lshortfile)
	DebugLogger = log.New(file,"DEBUG:",log.Ldate|log.Ltime|log.Lshortfile)
}