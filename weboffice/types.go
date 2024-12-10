package weboffice

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"io"
	"net/http"
	"net/url"
)

type Context interface {
	context.Context

	AppID() string
	Token() string
	Query() url.Values
	RequestID() string
}

type userContext struct {
	context.Context

	appID     string
	token     string
	query     url.Values
	requestID string
}

func (uc *userContext) AppID() string {
	return uc.appID
}
func (uc *userContext) Token() string {
	return uc.token
}
func (uc *userContext) Query() url.Values {
	return uc.query
}
func (uc *userContext) RequestID() string {
	return uc.requestID
}

func ParseContext(req *http.Request) Context {
	uc := &userContext{
		Context:   req.Context(),
		appID:     req.Header.Get("X-App-ID"),
		token:     req.Header.Get("X-WebOffice-Token"),
		requestID: req.Header.Get("X-Request-ID"),
	}
	if v, err := url.ParseQuery(req.Header.Get("X-User-Query")); err == nil {
		uc.query = v
	} else {
		uc.query = url.Values{}
	}
	return uc
}

type Reply struct {
	Code    Code   `json:"code"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data"`
}

type Empty struct {
}

type GetFileReply struct {
	CreateTime int64  `json:"create_time"`
	CreatorId  string `json:"creator_id"`
	ID         string `json:"id"`
	ModifierId string `json:"modifier_id"`
	ModifyTime int64  `json:"modify_time"`
	Name       string `json:"name"`
	Size       int64  `json:"size"`
	Version    int32  `json:"version"`
}

type GetFileDownloadReply struct {
	URL        string            `json:"url"`
	Digest     string            `json:"digest"`
	DigestType string            `json:"digest_type"`
	Headers    map[string]string `json:"headers"`
}

type UploadFilePrepareReply struct {
	DigestTypes []string `json:"digest_types"`
}

type UploadAddressReply struct {
	Method string `json:"method"`
	URL    string `json:"url"`
}

// same as fileInfo>?
type UploadCallbackReply struct {
	CreateTime int64  `json:"create_time"`
	CreatorID  string `json:"creator_id"`
	ID         string `json:"id"`
	ModifierID string `json:"modifier_id"`
	ModifyTime int64  `json:"modify_time"`
	Name       string `json:"name"`
	Size       int64  `json:"size"`
	Version    int32  `json:"version"`
}

type GetFilePermissionReply struct {
	Comment  int    `json:"comment"`
	Copy     int    `json:"copy"`
	Download int    `json:"download"`
	History  int    `json:"history"`
	Print    int    `json:"print"`
	Read     int    `json:"read"`
	Rename   int    `json:"rename"`
	SaveAs   int    `json:"saveas"`
	Update   int    `json:"update"`
	UserId   string `json:"user_id"`
}

type BaseProvider interface {
	GetFile(ctx Context, fileID string) (*GetFileReply, error)
	GetFileDownload(ctx Context, fileID string) (*GetFileDownloadReply, error)
	GetFilePermission(ctx Context, fileID string) (*GetFilePermissionReply, error)
}

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
	Logined   bool   `json:"logined"`
}

type UserProvider interface {
	GetUsers(ctx Context, userIDs []string) ([]*User, error)
}

type GetWatermarkReply struct {
	Type       int     `json:"type"`
	Value      string  `json:"value"`
	FillStyle  string  `json:"fill_style"`
	Font       string  `json:"font"`
	Rotate     float64 `json:"rotate"`
	Horizontal int     `json:"horizontal"`
	Vertical   int     `json:"vertical"`
}

type WatermarkProvider interface {
	GetFileWatermark(ctx Context, fileID string) (*GetWatermarkReply, error)
}

type UpdateFile1PhaseArgs struct {
	Name     string
	Size     int64
	SHA1     string
	IsManual bool
	Content  io.Reader
}

type UploadAddressArgs struct {
	FileID string `json:"file_id"`
	Name   string `json:"name"`
	Size   int    `json:"size"`
	Digest struct {
		Sha1 string `json:"sha1"`
	} `json:"digest"`
	IsManual bool `json:"is_manual"`
}

type UploadCallbackReq struct {
	Request        CallbackMustRequest    `json:"request"`
	Response       CallbackMustResponse   `json:"response"`
	SendBackParams CallbackSendBackParams `json:"send_back_params"`
}

type EditProvider interface {
	UploadPrepare(ctx Context) *UploadFilePrepareReply
	UploadAddress(ctx Context, args *UploadAddressArgs) (*UploadAddressReply, error)
	UploadCallback(ctx Context, args *UploadCallbackReq) (*UploadCallbackReply, error)
	UpdateFile(ctx Context, fileID string, args *UpdateFile1PhaseArgs) (*GetFileReply, error) // 单阶段提交,官方已停用，建议走三阶段
	RenameFile(ctx Context, fileID string, args *RenameFileArgs) error
	UploadFile(ctx Context, fileID string, content []byte) error
}

type RenameFileArgs struct {
	Name string `json:"name"`
}

type VersionProvider interface {
	GetFileVersions(ctx Context, fileID string, offset, limit int) ([]*GetFileReply, error)
	GetFileVersion(ctx Context, fileID string, version int32) (*GetFileReply, error)
	GetFileVersionDownload(ctx Context, fileID string, version int32) (*GetFileDownloadReply, error)
}

type DownloadProvider interface {
	DownloadFile(r *ghttp.Request)
}

type FullProvider struct {
	BaseProvider
	UserProvider
	WatermarkProvider
	EditProvider
	VersionProvider
	DownloadProvider
}
