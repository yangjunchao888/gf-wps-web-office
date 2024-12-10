package weboffice

import (
	"gf-wps-web-office/log"
	"gf-wps-web-office/model/weboffice"
	webofficeTool "gf-wps-web-office/weboffice"
	"github.com/gogf/gf/v2/net/ghttp"
	"io/ioutil"
	"net/http"
)

func GetFile(r *ghttp.Request) {
	req := &weboffice.GetFileReq{}
	if err := r.Parse(req); err != nil {
		log.Errorf("req err: %v", err)
		r.Response.WriteStatus(http.StatusNotFound, "file_id not found")
		return
	}

	log.Infof("GetFile req: %+v\n", req)

	ctx := webofficeTool.ParseContext(r.Request)
	statusCode, data := JsonConstruct(provider.GetFile(ctx, req.FileID))

	r.Response.Header().Set("Content-Type", "application/json")
	r.Response.WriteStatus(statusCode, data)
	log.Infof("GetFile resp: statusCode %d, data %+v\n", statusCode, data)
}

func GetFileDownload(r *ghttp.Request) {
	req := &weboffice.GetFileDownloadReq{}
	if err := r.Parse(req); err != nil {
		log.Errorf("req err: %v", err)
		r.Response.WriteStatus(http.StatusNotFound, "file_id not found")
		return
	}

	log.Infof("GetFileDownload req: %+v\n", req)

	ctx := webofficeTool.ParseContext(r.Request)
	statusCode, data := JsonConstruct(provider.GetFileDownload(ctx, req.FileID))

	r.Response.Header().Set("Content-Type", "application/json")
	r.Response.WriteStatus(statusCode, data)
	log.Infof("GetFileDownload resp: statusCode %d, data %+v\n", statusCode, data)
}

func GetFilePermission(r *ghttp.Request) {
	req := &weboffice.GetFilePermissionReq{}
	if err := r.Parse(req); err != nil {
		log.Errorf("req err: %v", err)
		r.Response.WriteStatus(http.StatusNotFound, "file_id not found")
		return
	}

	log.Infof("GetFilePermission req: %+v\n", req)

	ctx := webofficeTool.ParseContext(r.Request)
	statusCode, data := JsonConstruct(provider.GetFilePermission(ctx, req.FileID))

	// 设置响应头为JSON类型
	r.Response.Header().Set("Content-Type", "application/json")
	r.Response.WriteStatus(statusCode, data)
	log.Infof("GetFilePermission resp: statusCode %d, data %+v\n", statusCode, data)
}

func GetUsers(r *ghttp.Request) {
	req := &weboffice.GetUsersReq{}
	if err := r.Parse(req); err != nil {
		log.Errorf("req err: %v", err)
		r.Response.WriteStatus(http.StatusNotFound, "file_id not found")
		return
	}
	queryCache := r.Request.URL.Query()
	req.UserIDs, _ = queryCache["user_ids"]

	log.Infof("GetUsers req: %+v\n", req)

	ctx := webofficeTool.ParseContext(r.Request)
	statusCode, data := JsonConstruct(provider.GetUsers(ctx, req.UserIDs))

	r.Response.Header().Set("Content-Type", "application/json")
	r.Response.WriteStatus(statusCode, data)
	log.Infof("GetUsers resp: statusCode %d, data %+v\n", statusCode, data)
}

func GetFileWatermark(r *ghttp.Request) {
	req := &weboffice.GetFileWatermarkReq{}
	if err := r.Parse(req); err != nil {
		log.Errorf("req err: %v", err)
		r.Response.WriteStatus(http.StatusNotFound, "file_id not found")
		return
	}

	log.Infof("GetFileWatermark req: %+v\n", req)

	ctx := webofficeTool.ParseContext(r.Request)
	statusCode, data := JsonConstruct(provider.GetFileWatermark(ctx, req.FileID))

	r.Response.Header().Set("Content-Type", "application/json")
	r.Response.WriteStatus(statusCode, data)
	log.Infof("GetFileWatermark resp: statusCode %d, data %+v\n", statusCode, data)
}

func GetFileVersions(r *ghttp.Request) {
	req := &weboffice.GetFileVersionsReq{}
	if err := r.Parse(req); err != nil {
		log.Errorf("req err: %v", err)
		r.Response.WriteStatus(http.StatusNotFound, "file_id not found")
		return
	}

	log.Infof("GetFileVersions req: %+v\n", req)

	ctx := webofficeTool.ParseContext(r.Request)
	statusCode, data := JsonConstruct(provider.GetFileVersions(ctx, req.FileID, req.Offset, req.Limit))

	r.Response.Header().Set("Content-Type", "application/json")
	r.Response.WriteStatus(statusCode, data)
	log.Infof("GetFileVersions resp: statusCode %d, data %+v\n", statusCode, data)
}

func GetFileVersion(r *ghttp.Request) {
	req := &weboffice.GetFileVersionReq{}
	if err := r.Parse(req); err != nil {
		log.Errorf("req err: %v", err)
		r.Response.WriteStatus(http.StatusNotFound, "file_id not found")
		return
	}

	log.Infof("GetFileVersion req: %+v\n", req)

	ctx := webofficeTool.ParseContext(r.Request)
	statusCode, data := JsonConstruct(provider.GetFileVersion(ctx, req.FileID, req.Version))

	r.Response.Header().Set("Content-Type", "application/json")
	r.Response.WriteStatus(statusCode, data)
	log.Infof("GetFileVersion resp: statusCode %d, data %+v\n", statusCode, data)
}

func GetFileVersionDownload(r *ghttp.Request) {
	req := &weboffice.GetFileVersionDownloadReq{}
	if err := r.Parse(req); err != nil {
		log.Errorf("req err: %v", err)
		r.Response.WriteStatus(http.StatusNotFound, "file_id not found")
		return
	}

	log.Infof("GetFileVersionDownload req: %+v\n", req)

	ctx := webofficeTool.ParseContext(r.Request)
	statusCode, data := JsonConstruct(provider.GetFileVersionDownload(ctx, req.FileID, req.Version))

	r.Response.Header().Set("Content-Type", "application/json")
	r.Response.WriteStatus(statusCode, data)
	log.Infof("GetFileVersionDownload resp: statusCode %d, data %+v\n", statusCode, data)
}

func DownloadFile(r *ghttp.Request) {
	provider.DownloadFile(r)
	return
}

func RenameFile(r *ghttp.Request) {
	args := &webofficeTool.RenameFileArgs{}
	fileID := r.GetQuery("file_id").String()
	if err := r.Parse(args); err != nil {
		log.Errorf("req err: %v", err)
		r.Response.WriteStatus(http.StatusNotFound, "file_id not found")
		return
	}

	log.Infof("RenameFile req: %s\n", fileID)

	ctx := webofficeTool.ParseContext(r.Request)
	err := provider.RenameFile(ctx, fileID, args)
	if err != nil {
		r.Response.WriteStatus(JsonConstruct(nil, err))
		return
	}

	r.Response.Header().Set("Content-Type", "application/json")
	r.Response.WriteStatus(JsonConstruct(&webofficeTool.Empty{}, err))
}

func UploadPrepare(r *ghttp.Request) {
	args := &weboffice.UploadPrepareReq{}
	if err := r.Parse(args); err != nil {
		log.Errorf("req err: %v", err)
		r.Response.WriteStatus(http.StatusNotFound, "file_id not found")
		return
	}

	log.Infof("UploadPrepare req: %+v\n", args)

	ctx := webofficeTool.ParseContext(r.Request)
	data := provider.UploadPrepare(ctx)

	r.Response.Header().Set("Content-Type", "application/json")
	r.Response.WriteStatus(JsonConstruct(data, nil))
	log.Infof("UploadPrepare resp: data %+v\n", data)
}

func UploadAddress(r *ghttp.Request) {
	args := &webofficeTool.UploadAddressArgs{}
	if err := r.Parse(args); err != nil {
		log.Errorf("req err: %v", err)
		r.Response.WriteStatus(http.StatusNotFound, "file_id not found")
		return
	}

	log.Infof("UploadAddress req: %+v\n", args)

	ctx := webofficeTool.ParseContext(r.Request)
	statusCode, data := JsonConstruct(provider.UploadAddress(ctx, args))

	r.Response.Header().Set("Content-Type", "application/json")
	r.Response.WriteStatus(statusCode, data)
	log.Infof("UploadAddress resp: statusCode %d, data %+v\n", statusCode, data)
}

func UploadCallback(r *ghttp.Request) {
	args := &webofficeTool.UploadCallbackReq{}
	if err := r.Parse(args); err != nil {
		log.Errorf("req err: %v", err)
		r.Response.WriteStatus(http.StatusNotFound, "file_id not found")
		return
	}

	ctx := webofficeTool.ParseContext(r.Request)
	statusCode, data := JsonConstruct(provider.UploadCallback(ctx, args))

	r.Response.Header().Set("Content-Type", "application/json")
	r.Response.WriteStatus(statusCode, data)
}

// UploadHandler 处理文件更新请求
func UploadHandler(r *ghttp.Request) {
	req := &weboffice.UploadFileReq{}
	if err := r.Parse(req); err != nil {
		log.Errorf("req err: %v", err)
		r.Response.WriteStatus(http.StatusNotFound, "file_id not found")
		return
	}

	// 读取请求体中的文件内容
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		r.Response.WriteStatus(http.StatusInternalServerError, err.Error())
		return
	}

	ctx := webofficeTool.ParseContext(r.Request)
	log.Infof("UploadHandler req: %+v\n", req)
	// 更新文件
	err = provider.UploadFile(ctx, req.FileID, content)
	if err != nil {
		r.Response.WriteStatus(http.StatusInternalServerError, err.Error())
		return
	}
	r.Response.Header().Set("Content-Type", "application/json")
	r.Response.WriteStatus(200, "ok")
}
