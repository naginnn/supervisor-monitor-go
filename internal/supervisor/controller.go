package supervisor

import (
	"github.com/abrander/go-supervisord"
	"github.com/gin-gonic/gin"
	m "supervisor-monitor/middlewares"
)

type handler struct {
	SP   *supervisord.Client
	conf string
}

func RegisterRoutes(r *gin.Engine, c *supervisord.Client, confPath string) {
	h := &handler{
		SP:   c,
		conf: confPath,
	}
	routes := r.Group("/")
	// views
	routes.GET("/auth", h.GetAuthPage)
	routes.GET("/token", h.GetToken)

	routes.GET("/apps", m.SessionChecker(h.GetAppsPage))
	routes.GET("/apps/:name", m.SessionChecker(h.GetAppPage))
	routes.GET("/config", m.SessionChecker(h.GetConfig))

	// endpoint
	routes.POST("/control/processes/:cmd", m.SessionChecker(h.ControlProcesses))
	routes.POST("/control/processes/process/:name/:cmd", m.SessionChecker(h.ControlProcess))
	routes.POST("/control/config/:cmd", m.SessionChecker(h.ControlSupervisor))
	routes.POST("/logs/clear/:name", m.SessionChecker(h.ClearLogFile))

	routes.GET("/logs/:name", m.SessionChecker(h.GetProcessLogPage))
	routes.GET("/logs/err/:name", m.SessionChecker(h.GetProcessErrLogPage))
	routes.GET("/processes", m.SessionChecker(h.GetAllProcesses))
	routes.GET("/processes/:name", m.SessionChecker(h.GetProcess))

	//routes.GET("/logs/:name", h.GetProcessLog)

	//routes.GET("/error_logs/:log", h.GetProcessLog)

	//routes.GET("/start_all", h.StartAllProcesses)
	//routes.GET("/restart_all", h.RestartAllProcesses)
	//
	//routes.GET("/:name/stop", h.RestartAllProcesses)
	//routes.GET("/:name/start", h.RestartAllProcesses)
	//routes.GET("/:name/restart", h.RestartAllProcesses)
	//
	//routes.GET("/config/:cmd", h.RestartAllProcesses)

	//routes.DELETE("/delete/:hdfsName", h.GetHClusters)
	//routes.POST("/add", h.AddHCluster)
}
