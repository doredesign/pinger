package main

import (
	"errors"
	"os/exec"
	"path"
	"regexp"
	"runtime"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var timeRegex *regexp.Regexp = regexp.MustCompile(`time=(\d+\.\d+)`)

func main() {
	r := gin.Default()
	_, filename, _, _ := runtime.Caller(0)
	templatesPath := path.Join(path.Dir(filename), "./templates/*")
	r.LoadHTMLGlob(templatesPath)
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.tmpl", gin.H{})
	})
	r.GET("/pinger", func(c *gin.Context) {
		c.HTML(200, "pinger.tmpl", gin.H{})
	})

	r.GET("/ping", func(c *gin.Context) {
		errorMsg := ""
		time, err := ping()
		// time, err := longTime()

		if err != nil {
			errorMsg = err.Error()
		}

		c.JSON(200, gin.H{
			"time":  time,
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

	arrayResult := timeRegex.FindStringSubmatch(string(stdoutStderr))
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

func longTime() (float64, error) {
	second := time.Now().Second()
	sleepTime := (50 / 3) * second
	time.Sleep(time.Duration(sleepTime) * time.Millisecond)
	return float64(sleepTime), nil
}
