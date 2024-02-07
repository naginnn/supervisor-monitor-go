package main

import (
	"embed"
	"github.com/abrander/go-supervisord"
	"github.com/gin-gonic/gin"
	cachecontrol "github.com/joeig/gin-cachecontrol"
	"github.com/spf13/viper"
	"html/template"
	"io/fs"
	"net/http"
	"supervisor-monitor/internal/supervisor"
	"supervisor-monitor/pkg/logger"
)

//go:embed static
var s embed.FS

//go:embed templates/*
var f embed.FS

func main() {

	gin.SetMode(gin.ReleaseMode)
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()

	sp := viper.GetStringMapString("supervisor")
	confPath := sp["config"]
	c, err := supervisord.NewClient(sp["url"])

	sr := viper.GetStringMapString("server")
	runStr := sr["host"] + ":" + sr["port"]

	if err != nil {
		logger.Log.Printf("Panic, config not found")
		panic(err.Error())
	}
	_, err = c.GetPID()
	if err != nil {
		logger.Log.Printf("Panic, supervisor dont access")
		panic(err.Error())
	}

	r := gin.Default()

	//t, err := loadTemplate()
	//if err != nil {
	//	panic(err)
	//}
	//r.SetHTMLTemplate(t)
	supervisor.RegisterRoutes(r, c, confPath)
	//r.Use(middlewares.JSONLogMiddleware())
	r.Use(cachecontrol.New(&cachecontrol.Config{
		MustRevalidate:       true,
		NoCache:              true,
		NoStore:              true,
		NoTransform:          false,
		Public:               false,
		Private:              false,
		ProxyRevalidate:      true,
		MaxAge:               nil,
		SMaxAge:              nil,
		Immutable:            false,
		StaleWhileRevalidate: nil,
		StaleIfError:         nil,
	}))

	r.NoRoute(func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/apps")
	})
	//r.Static("/static", "./static")

	staticFs, err := fs.Sub(s, "static")
	r.StaticFS("/static", http.FS(staticFs))

	templ := template.Must(template.New("").ParseFS(f, "templates/*"))
	r.SetHTMLTemplate(templ)
	//r.LoadHTMLGlob("./templates/*")
	err = r.Run(runStr)
	logger.Log.Printf("Success run on: " + "" + "")
	if err != nil {
		return
	}

	defer c.Close()
	//err = c.ClearLog()
	//if err != nil {
	//	panic(err.Error())
	//}

	//err = c.Restart()
	//if err != nil {
	//	panic(err.Error())
	//}
}
