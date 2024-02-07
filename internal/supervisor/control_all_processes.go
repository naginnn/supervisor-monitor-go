package supervisor

import (
	"github.com/abrander/go-supervisord"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func (h handler) ControlProcesses(c *gin.Context) {
	cmd := c.Param("cmd")
	var proc []supervisord.ProcessInfo
	var err error
	switch cmd {
	case "start":
		proc, err = h.SP.StartAllProcesses(true)
	case "stop":
		proc, err = h.SP.StopAllProcesses(true)
	case "restart":
		_, err = h.SP.StopAllProcesses(true)
		proc, err = h.SP.StartAllProcesses(true)
	}
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusNotModified, nil)
		return
	}
	c.JSON(http.StatusOK, proc)
}

//func (h handler) StopAllProcesses(c *gin.Context) {
//	proc, err := h.SP.StopAllProcesses(true)
//	if err != nil {
//		log.Println(err)
//		c.JSON(http.StatusNotModified, nil)
//		return
//	}
//	c.JSON(http.StatusOK, proc)
//}
//
//func (h handler) StartAllProcesses(c *gin.Context) {
//	proc, err := h.SP.StartAllProcesses(true)
//	if err != nil {
//		log.Println(err)
//		c.JSON(http.StatusNotModified, nil)
//		return
//	}
//	c.JSON(http.StatusOK, proc)
//}
//
//func (h handler) RestartAllProcesses(c *gin.Context) {
//	_, err := h.SP.StopAllProcesses(true)
//	if err != nil {
//		log.Println(err)
//	}
//	proc, err := h.SP.StartAllProcesses(true)
//	if err != nil {
//		log.Println(err)
//		c.JSON(http.StatusNotModified, nil)
//		return
//	}
//	c.JSON(http.StatusOK, proc)
//}
