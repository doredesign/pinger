package main

import (
	// "fmt"
	// "log"
	"github.com/gin-gonic/gin"
	"os/exec"
	"regexp"
	"strconv"
	"errors"
	"time"
)

var TimeRegex *regexp.Regexp = regexp.MustCompile(`time=(\d+\.\d+)`)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{})
	})
	r.GET("/pinger", func(c *gin.Context) {
		c.HTML(200, "pinger.tmpl", gin.H{})
	})

	r.GET("/ping", func(c *gin.Context) {
		errorMsg := ""
		time, err := ping()
		// time, err := fakeTime()

		if err != nil {
			errorMsg = err.Error()
		}

		c.JSON(200, gin.H{
			"time": time,
			"error": errorMsg,
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func ping() (float64, error) {
	cmd := exec.Command("ping", "-c", "1", "8.8.8.8")
	stdoutStderr, err := cmd.CombinedOutput()
	result := 0.0
	if err != nil {
		return result, err
	}

	arrayResult := TimeRegex.FindStringSubmatch(string(stdoutStderr))
	err = errors.New("Ping unsuccessful")
	if len(arrayResult) > 1 {
		result, err = strconv.ParseFloat(arrayResult[1], 64)
	}

	return result, err
}

func fakeTime() (float64, error) {
	second := time.Now().Second()
	return float64(second * 10), nil
}
