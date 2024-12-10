package main

import (
	"gf-wps-web-office/boot"
	"gf-wps-web-office/config"
	"gf-wps-web-office/handler/weboffice"
	"gf-wps-web-office/router"
	"github.com/gogf/gf/v2/frame/g"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	s := g.Server()
	boot.Init()

	apiGroup := s.Group("/api/v1")
	wps := apiGroup.Group("/weboffice")
	router.RegisterRoute(wps)

	apiGroup.GET("/download", weboffice.DownloadFile)
	apiGroup.PUT("/:file_id/upload_file", weboffice.UploadHandler)

	port := config.GetConfig().Port
	s.SetPort(port)
	s.Run()
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
