package router

import (
	"gf-wps-web-office/handler/weboffice"
	"github.com/gogf/gf/v2/net/ghttp"
)

func RegisterRoute(group *ghttp.RouterGroup) {
	group.GET("/v3/3rd/files/:file_id", weboffice.GetFile)
	group.GET("/v3/3rd/files/:file_id/download", weboffice.GetFileDownload)
	group.GET("/v3/3rd/files/:file_id/permission", weboffice.GetFilePermission)
	group.GET("/v3/3rd/users", weboffice.GetUsers)
	group.GET("/v3/3rd/files/:file_id/watermark", weboffice.GetFileWatermark)
	//group.GET("/v3/3rd/files/:file_id/upload", weboffice.UpdateFile) // @TODO wps office delete already
	group.GET("/v3/3rd/files/:file_id/versions", weboffice.GetFileVersions)
	group.PUT("/v3/3rd/files/:file_id/name", weboffice.RenameFile)
	group.GET("/v3/3rd/files/:file_id/versions/:version", weboffice.GetFileVersion)
	group.GET("/v3/3rd/files/:file_id/versions/:version/download", weboffice.GetFileVersionDownload)

	// 三阶段提交更新
	group.GET("/v3/3rd/files/:file_id/upload/prepare", weboffice.UploadPrepare)
	group.POST("/v3/3rd/files/:file_id/upload/address", weboffice.UploadAddress)
	group.POST("/v3/3rd/files/:file_id/upload/complete", weboffice.UploadCallback)

	//group.GET("download", weboffice.DownloadFile)
	//group.PUT("/:file_id/upload_file", weboffice.UploadHandler)

}
