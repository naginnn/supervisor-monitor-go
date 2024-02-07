package supervisor

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"supervisor-monitor/middlewares"
	"supervisor-monitor/pkg/logger"
)

func (h handler) GetAuthPage(c *gin.Context) {
	fmt.Println()
	c.HTML(http.StatusOK, "auth.html", nil)
}

func (h handler) GetToken(c *gin.Context) {
	header := c.GetHeader("Authorization")
	contain := strings.Split(header, " ")
	token, err := middlewares.CheckBasicAuth(contain[1])
	if err != nil {
		logger.Log.Printf("unauthorized")
		c.JSON(http.StatusUnauthorized, nil)
		return
	}
	c.SetCookie("token", token, 3600, "/", "localhost", false, true)
	logger.Log.Printf("success log in")
	c.JSON(http.StatusOK, gin.H{"redirect": "/apps"})
}

func (h handler) GetAppsPage(c *gin.Context) {
	c.HTML(http.StatusOK, "apps.html", gin.H{
		"title": "Main website",
	})
}

func getLogByPath(path string) string {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	command := "tail -n 100 " + path
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return ""
	}
	logger.Log.Printf("get log ok")
	return stdout.String()
}

func (h handler) GetProcessLogPage(c *gin.Context) {
	name := c.Param("name")
	proc, err := h.SP.GetProcessInfo(name)
	if err != nil {

	}
	log := getLogByPath(proc.StdoutLogfile)
	logger.Log.Printf("get StdoutLogfile ok")
	c.HTML(http.StatusOK, "logs.html", gin.H{
		"log": log,
	})
}

func (h handler) GetProcessErrLogPage(c *gin.Context) {
	name := c.Param("name")
	proc, err := h.SP.GetProcessInfo(name)
	if err != nil {

	}
	log := getLogByPath(proc.StderrLogfile)
	c.HTML(http.StatusOK, "logs.html", gin.H{
		"log": log,
	})
}

func (h handler) GetConfig(c *gin.Context) {
	var names []string
	pid, err := h.SP.GetPID()
	if err != nil {
		log.Println(err)
		pid = 0
	}

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	path := h.conf
	cmd := exec.Command("cat", path)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Run()
	if err != nil {
	}
	config := stdout.String()

	//if err != nil {
	//	log.Println(err)
	//} else {
	//	for _, value := range config {
	//		names = append(names, value.Name)
	//	}
	//}
	c.HTML(http.StatusOK, "config.html", gin.H{
		"Names": names, "Config": config, "Pid": pid,
	})
}

func (h handler) GetAppPage(c *gin.Context) {
	name := c.Param("name")
	proc, err := h.SP.GetProcessInfo(name)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotModified, nil)
		return
	}
	c.HTML(http.StatusOK, "app.html", proc)
}
