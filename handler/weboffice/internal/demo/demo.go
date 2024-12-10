package demo

import (
	"bytes"
	"crypto/sha1"
	"embed"
	"encoding/hex"
	"errors"
	"fmt"
	"gf-wps-web-office/weboffice"
	"github.com/gogf/gf/v2/net/ghttp"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"
)

//go:embed test.xlsx
//go:embed test.docx
//go:embed test.pptx
var _fs embed.FS

func sha1hex(src []byte) string {
	h := sha1.New()
	h.Write(src)
	digest := h.Sum(nil)
	return hex.EncodeToString(digest)
}

type File struct {
	info     FileInfo
	versions []FileInfo
}

type FileInfo struct {
	id       string
	creator  string
	modifier string

	name       string
	createTime time.Time
	modifyTime time.Time
	version    int32
	size       int64

	content []byte
	digest  string
}

func (file *FileInfo) ToFileInfo() *weboffice.GetFileReply {
	return &weboffice.GetFileReply{
		CreateTime: file.createTime.Unix(),
		CreatorId:  file.creator,
		ID:         file.id,
		ModifierId: file.modifier,
		ModifyTime: file.modifyTime.Unix(),
		Name:       file.name,
		Size:       file.size,
		Version:    file.version,
	}
}

type Provider struct {
	gateway string

	mutex sync.Mutex
	files map[string]*File
}

func NewProvider() *Provider {
	gateway := os.Getenv("PROVIDER_GATEWAY")
	if gateway == "" {
		gateway = "http://localhost:8180/api/v1"
		//gateway = "http://localhost:8180/api/v1/weboffice"
	}

	db := &Provider{
		gateway: gateway,
		files:   map[string]*File{},
	}
	db.createFile("1", "test.xlsx")
	db.createFile("2", "test.docx")
	db.createFile("3", "test.pptx")
	return db
}

func (db *Provider) UploadAddress(fileID string) (*weboffice.UploadAddressReply, error) {
	return &weboffice.UploadAddressReply{
		Method: "PUT",
		URL:    fmt.Sprintf("%s/%s/upload_file", db.gateway, fileID),
	}, nil
}

func (db *Provider) createFile(id string, name string) *Provider {
	content, err := _fs.ReadFile(name)
	if err != nil {
		log.Fatal(err)
	}

	info := FileInfo{
		id:         id,
		creator:    "owner",
		modifier:   "owner",
		name:       name,
		createTime: time.Now(),
		modifyTime: time.Now(),
		version:    1,
		size:       int64(len(content)),
		content:    content,
		digest:     sha1hex(content),
	}
	file := &File{info: info}
	file.versions = append(file.versions, file.info)

	db.files[file.info.id] = file
	return db
}

// updateFile 更新一个文件
func (db *Provider) updateFile(id string, content []byte) error {
	file, ok := db.files[id]
	if !ok {
		log.Printf("File not found: id %s\n", id)
		return errors.New("file not found")
	}

	info := FileInfo{
		id:         id,
		creator:    file.info.creator,
		modifier:   "owner",
		name:       file.info.name,
		createTime: file.info.createTime,
		modifyTime: time.Now(),
		version:    file.info.version + 1,
		size:       int64(len(content)),
		content:    content,
		digest:     sha1hex(content),
	}
	file.info = info
	file.versions = append(file.versions, info)

	return nil
}

// UploadFile 处理文件更新请求
func (db *Provider) UploadFile(_ weboffice.Context, fileID string, content []byte) error {
	// 更新文件
	return db.updateFile(fileID, content)
}

func (db *Provider) GetFile(_ weboffice.Context, fileID string) (*weboffice.GetFileReply, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	file, ok := db.files[fileID]
	if !ok {
		return nil, weboffice.ErrFileNotExists
	}
	return file.info.ToFileInfo(), nil
}

func (db *Provider) DownloadFile(r *ghttp.Request) {
	//if r.GetHeader("Referer") == "https://solution.wps.cn" {
	//	r.Response.WriteStatus(http.StatusForbidden, "invalid referer")
	//	return
	//}

	fileID := r.GetQuery("file_id").String()
	versionID := r.GetQuery("version").String()

	db.mutex.Lock()
	defer db.mutex.Unlock()

	file, ok := db.files[fileID]
	if !ok {
		r.Response.WriteStatus(http.StatusNotFound, "file not found")
		return
	}

	log.Printf("DownloadFile req fileID: %s, version: %s\n", fileID, versionID)

	if versionID == "" {
		r.Response.ServeContent(file.info.name, file.info.modifyTime, bytes.NewReader(file.info.content))
		return
	}
	for _, v := range file.versions {
		if strconv.Itoa(int(v.version)) == versionID {
			r.Response.ServeContent(file.info.name, file.info.modifyTime, bytes.NewReader(file.info.content))
			return
		}
	}
}

func (db *Provider) RenameFile(_ weboffice.Context, fileID string, args *weboffice.RenameFileArgs) error {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	file, ok := db.files[fileID]
	if !ok {
		return weboffice.ErrFileNotExists
	}
	file.info.name = args.Name
	return nil
}

func (db *Provider) GetFileDownload(_ weboffice.Context, fileID string) (*weboffice.GetFileDownloadReply, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	file, ok := db.files[fileID]
	if !ok {
		return nil, weboffice.ErrFileNotExists
	}

	return &weboffice.GetFileDownloadReply{
		URL:        db.gateway + "/download?file_id=" + fileID,
		Digest:     file.info.digest,
		DigestType: "sha1",
		Headers: map[string]string{
			"Referer": "https://solution.wps.cn",
		},
	}, nil
}

func (db *Provider) GetFilePermission(ctx weboffice.Context, fileID string) (*weboffice.GetFilePermissionReply, error) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	_, ok := db.files[fileID]
	if !ok {
		return nil, weboffice.ErrFileNotExists
	}

	if ctx.Token() != "" {
		return &weboffice.GetFilePermissionReply{
			Comment:  1,
			Copy:     1,
			Download: 1,
			History:  1,
			Print:    1,
			Read:     1,
			Rename:   0,
			SaveAs:   1,
			Update:   1,
			UserId:   ctx.Token(),
		}, nil
	} else {
		// @TODO 仅调试使用，匿名用户给所有文件操作权限
		return &weboffice.GetFilePermissionReply{
			Comment:  1,
			Copy:     1,
			Download: 1,
			History:  1,
			Print:    1,
			Read:     1,
			Rename:   1,
			SaveAs:   1,
			Update:   1,
			UserId:   "anonymous",
		}, nil
	}
}

func (db *Provider) UpdateFile(ctx weboffice.Context, fileID string, args *weboffice.UpdateFile1PhaseArgs) (*weboffice.GetFileReply, error) {
	content, err := io.ReadAll(args.Content)
	if err != nil {
		return nil, weboffice.ErrInternalError.WithMessage(err.Error())
	}
	db.mutex.Lock()
	defer db.mutex.Unlock()

	file, ok := db.files[fileID]
	if !ok {
		return nil, weboffice.ErrFileNotExists
	}

	file.versions = append(file.versions, file.info)

	file.info.version++
	file.info.size = args.Size
	file.info.modifier = ctx.Token()
	file.info.modifyTime = time.Now()
	file.info.content = content

	return file.info.ToFileInfo(), nil
}

func (db *Provider) GetUsers(_ weboffice.Context, userIDs []string) ([]*weboffice.User, error) {
	var users []*weboffice.User
	for _, id := range userIDs {
		if id == "anonymous" {
			users = append(users, &weboffice.User{
				ID:        id,
				Name:      "anonymous",
				AvatarURL: "",
				Logined:   false,
			})
		} else {
			users = append(users, &weboffice.User{
				ID:        id,
				Name:      "user_" + id,
				AvatarURL: "",
				Logined:   true,
			})
		}
	}
	return users, nil
}

func (db *Provider) GetFileVersions(_ weboffice.Context, fileID string, offset, limit int) ([]*weboffice.GetFileReply, error) {
	file, ok := db.files[fileID]
	if !ok {
		return nil, weboffice.ErrFileNotExists
	}

	var reply []*weboffice.GetFileReply
	if offset < len(file.versions) {
		if remain := len(file.versions) - offset; remain < limit {
			limit = remain
		}
		for _, v := range file.versions[offset : offset+limit] {
			reply = append(reply, v.ToFileInfo())
		}
	}
	return reply, nil
}

func (db *Provider) GetFileVersion(_ weboffice.Context, fileID string, version int32) (*weboffice.GetFileReply, error) {
	file, ok := db.files[fileID]
	if !ok {
		return nil, weboffice.ErrFileNotExists
	}

	for _, v := range file.versions {
		if v.version == version {
			return v.ToFileInfo(), nil
		}
	}
	return nil, weboffice.ErrFileVersionNotExists
}

func (db *Provider) GetFileLatestVersion(_ weboffice.Context, fileID string) (*weboffice.GetFileReply, error) {
	file, ok := db.files[fileID]
	if !ok {
		return nil, weboffice.ErrFileNotExists
	}

	versionLen := len(file.versions)
	if versionLen == 0 {
		return nil, weboffice.ErrFileVersionNotExists
	}
	return file.versions[versionLen-1].ToFileInfo(), nil
}

func (db *Provider) GetFileVersionDownload(_ weboffice.Context, fileID string, version int32) (*weboffice.GetFileDownloadReply, error) {
	file, ok := db.files[fileID]
	if !ok {
		return nil, weboffice.ErrFileNotExists
	}

	for _, v := range file.versions {
		if v.version == version {
			return &weboffice.GetFileDownloadReply{
				URL:        fmt.Sprintf("%s/download?file_id=%s&version=%d", db.gateway, fileID, version),
				Digest:     v.digest,
				DigestType: "sha1",
				Headers: map[string]string{
					"Referer": "https://solution.wps.cn",
				},
			}, nil
		}
	}
	return nil, weboffice.ErrFileVersionNotExists
}
