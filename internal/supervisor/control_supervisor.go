package supervisor

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/exec"
)

type requestBody struct {
	Config string `json:"config"`
}

func (h handler) ControlSupervisor(c *gin.Context) {
	cmd := c.Param("cmd")
	var err error
	switch cmd {
	case "shutdown":
		err = h.SP.Shutdown()
	case "restart":
		err = h.SP.Restart()
	case "reload_config":
		_, _, _, err = h.SP.ReloadConfig()
	case "clear_log":
		err = h.SP.ClearLog()
	case "shutdown_apply_config":
		err = h.SP.Shutdown()
		cmd := exec.Command("supervisord", "-c", h.conf)
		err = cmd.Run()
	case "kill_all_python_processes":
		command := "ps aux | grep python | grep -v \"grep python\" | awk '{print $2}' | xargs kill -9"
		cmd := exec.Command("bash", "-c", command)
		err = cmd.Run()
	case "save_config":
		var reqBody requestBody
		err = c.BindJSON(&reqBody)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(reqBody.Config)
			file, err := os.Create(h.conf)

			if err != nil {
				fmt.Println("Unable to create file:", err)
				os.Exit(1)
			}
			defer file.Close()
			file.WriteString(reqBody.Config)
		}
	}
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotModified, nil)
		return
	}
	c.JSON(http.StatusOK, nil)
}
