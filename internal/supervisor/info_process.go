package supervisor

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"os/exec"
)

func (h handler) GetProcess(c *gin.Context) {
	name := c.Param("name")
	proc, err := h.SP.GetProcessInfo(name)
	if err != nil {
		log.Error().Err(err)
		c.JSON(http.StatusNotModified, nil)
		return
	}
	c.JSON(http.StatusOK, proc)
}

func (h handler) ClearLogFile(c *gin.Context) {
	name := c.Param("name")
	proc, err := h.SP.GetProcessInfo(name)
	command := "cat /dev/null > " + proc.StdoutLogfile + ";" + "cat /dev/null > " + proc.StderrLogfile
	cmd := exec.Command("bash", "-c", command)
	err = cmd.Run()
	if err != nil {
		log.Error().Err(err)
	}
	c.JSON(http.StatusOK, nil)
}

func (h handler) GetProcessErrLog(c *gin.Context) {
	name := c.Param("name")
	proc, err := h.SP.GetProcessInfo(name)
	if err != nil {
		log.Error().Err(err)
		c.JSON(http.StatusNotModified, nil)
		return
	}
	c.JSON(http.StatusOK, proc)
}
