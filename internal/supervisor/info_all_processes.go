package supervisor

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

func (h handler) GetAllProcesses(c *gin.Context) {
	proc, err := h.SP.GetAllProcessInfo()
	if err != nil {
		log.Println(err)
		//c.JSON(http.StatusNotModified, nil)
		//return
	}
	pid, err := h.SP.GetPID()
	var mes string
	if err != nil {
		log.Println(err)
		mes = "FATAL PID NOT FOUND"
	} else {
		mes = strconv.Itoa(pid)
	}
	c.JSON(http.StatusOK, gin.H{"apps": proc, "state": mes})
}
