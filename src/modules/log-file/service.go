package log_file

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

func Logging(c *gin.Context) {
	file := makeLogFile()

	logSessionId(c)

	logRequest(c)

	logRequestBody(c)

	c.Next()

	// response log will be manually put in every API before they return the response

	// close the file after the request is done
	defer file.Close()
}

func makeLogFile() *os.File {
	// make date as the name
	logFile := time.Now().Format("2006-01-02") + ".log"

	// create the file if not exist
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	// set the log output to the file
	log.SetOutput(file)

	// set the log format
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	return file
}

func logSessionId(c *gin.Context) {
	sessionId := c.GetHeader("X-Session-Id")
	log.Println("=== Session ID:", sessionId)
}

func logRequest(c *gin.Context) {
	log.Println(c.Request.URL.String())
	log.Println(c.Request.Header)
}

func logRequestBody(c *gin.Context) {
	// because the body can only be read once, we need to read it first
	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println(err)
	}

	// set the body back to the request
	// so that the request can be read again later
	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	// log the body as string instead of bytes
	log.Println(string(bodyBytes))

}

func SaveErrorLog(c *gin.Context, err *error) {
	if err == nil {
		return
	}

	log.Println((*err).Error())
}

func SaveResponseLog(c *gin.Context, response *gin.H) {
	// format the response to be more readable
	formattedResponse, err := json.MarshalIndent(response, "", "  ")
	if err != nil {
		log.Println(err)
	}

	log.Println("Response:", string(formattedResponse))
}

func LogErrorAndResponse(c *gin.Context, err *error, response *gin.H, statusCode int) {
	SaveErrorLog(c, err)
	SaveResponseLog(c, response)

	c.JSON(statusCode, response)
}
