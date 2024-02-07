package supervisor

import (
	"errors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os/exec"
	"strconv"
)

// ControlProcess

func (h handler) ControlProcess(c *gin.Context) {
	name := c.Param("name")
	cmd := c.Param("cmd")
	var err error
	switch cmd {
	case "start":
		err = h.SP.StartProcess(name, true)
	case "stop":
		err = h.SP.StopProcess(name, true)
	case "restart":
		err = h.SP.StopProcess(name, true)
		err = h.SP.StartProcess(name, true)
	case "kill":
		proc, _ := h.SP.GetProcessInfo(name)
		command := "kill -9 " + strconv.Itoa(proc.Pid)
		if proc.Pid != 0 {
			cmd := exec.Command("bash", "-c", command)
			_ = cmd.Run()
		} else {
			err = errors.New("pid not found")
		}

	}
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotModified, nil)
		return
	}
	c.JSON(http.StatusOK, nil)
}

//func (h handler) StopProcess(c *gin.Context) {
//	name := c.Param("name")
//	err := h.SP.StopProcess(name, true)
//	if err != nil {
//		log.Println(err)
//		c.JSON(http.StatusNotModified, nil)
//		return
//	}
//	c.JSON(http.StatusOK, nil)
//}
//
//func (h handler) StartProcess(c *gin.Context) {
//	name := c.Param("name")
//	err := h.SP.StartProcess(name, true)
//	if err != nil {
//		log.Println(err)
//		c.JSON(http.StatusNotModified, nil)
//		return
//	}
//	c.JSON(http.StatusOK, nil)
//}
//
//func (h handler) RestartProcess(c *gin.Context) {
//	name := c.Param("name")
//	err := h.SP.StartProcess(name, true)
//	if err != nil {
//		log.Println(err)
//		c.JSON(http.StatusNotModified, nil)
//		return
//	}
//	c.JSON(http.StatusOK, nil)
//}
