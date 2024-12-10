package weboffice

import (
	"gf-wps-web-office/handler/weboffice/internal/demo"
	"gf-wps-web-office/utils"
	"gf-wps-web-office/weboffice"
	"github.com/gogf/gf/v2/net/ghttp"
	"log"
)

type TestProvider struct {
	Demo *demo.Provider
}

func (t *TestProvider) GetFile(ctx weboffice.Context, fileID string) (*weboffice.GetFileReply, error) {
	return t.Demo.GetFile(ctx, fileID)
}

func (t *TestProvider) RenameFile(ctx weboffice.Context, fileID string, args *weboffice.RenameFileArgs) error {
	return t.Demo.RenameFile(ctx, fileID, args)
}

func (t *TestProvider) UpdateFile(ctx weboffice.Context, fileID string, args *weboffice.UpdateFile1PhaseArgs) (*weboffice.GetFileReply, error) {
	return t.Demo.UpdateFile(ctx, fileID, args)
}

func (t *TestProvider) UploadFile(ctx weboffice.Context, fileID string, content []byte) error {
	return t.Demo.UploadFile(ctx, fileID, content)
}

func (t *TestProvider) GetFileDownload(ctx weboffice.Context, fileID string) (*weboffice.GetFileDownloadReply, error) {
	return t.Demo.GetFileDownload(ctx, fileID)
}

func (t *TestProvider) GetUsers(ctx weboffice.Context, userIDs []string) ([]*weboffice.User, error) {
	return t.Demo.GetUsers(ctx, userIDs)
}

func (t *TestProvider) GetFileWatermark(_ weboffice.Context, _ string) (*weboffice.GetWatermarkReply, error) {
	return &weboffice.GetWatermarkReply{
		Type:       1,
		Value:      "weboffice-go-sdk",
		FillStyle:  "rgba(192,192,192,0.6)",
		Font:       "bold 20px Serif",
		Rotate:     0.5,
		Horizontal: 50,
		Vertical:   50,
	}, nil
}

func (t *TestProvider) GetFilePermission(ctx weboffice.Context, fileID string) (*weboffice.GetFilePermissionReply, error) {
	return t.Demo.GetFilePermission(ctx, fileID)
}

func (t *TestProvider) GetFileVersions(ctx weboffice.Context, fileID string, offset, limit int) ([]*weboffice.GetFileReply, error) {
	return t.Demo.GetFileVersions(ctx, fileID, offset, limit)
}

func (t *TestProvider) GetFileVersion(ctx weboffice.Context, fileID string, version int32) (*weboffice.GetFileReply, error) {
	return t.Demo.GetFileVersion(ctx, fileID, version)
}

func (t *TestProvider) GetFileVersionDownload(ctx weboffice.Context, fileID string, version int32) (*weboffice.GetFileDownloadReply, error) {
	return t.Demo.GetFileVersionDownload(ctx, fileID, version)
}

func (t *TestProvider) UploadPrepare(ctx weboffice.Context) *weboffice.UploadFilePrepareReply {
	return &weboffice.UploadFilePrepareReply{
		DigestTypes: []string{"sha1"},
	}
}

func (t *TestProvider) DownloadFile(r *ghttp.Request) {
	t.Demo.DownloadFile(r)
}

var (
	// 示例使用
	minioEndpoint   = "http://202.105.141.102:9000"
	accessKeyID     = "mrV4hDVOyeZ4bZGUEwuN"
	secretAccessKey = "6EzIP3NJLp3jKgrJ2YDEwxZtao3mRjBpLBSMhpCD"
	bucketName      = "test"
)

// minio场景使用
func (t *TestProvider) UploadAddressMinio(ctx weboffice.Context) (*weboffice.UploadAddressReply, error) {
	editProvider, err := utils.NewMinioEditProvider(minioEndpoint, accessKeyID, secretAccessKey, bucketName)
	if err != nil {
		log.Println("Error creating MinioEditProvider:", err)
		return nil, err
	}

	urlInfo, err := editProvider.UploadAddress(ctx, "wps_web_office/test.xlsx")
	if err != nil {
		log.Println("Error getting upload address:", err)
		return nil, err
	}

	return &weboffice.UploadAddressReply{
		Method: "PUT",
		URL:    urlInfo.String(),
	}, nil
}

// 先按demo用法存储到内存中
func (t *TestProvider) UploadAddress(_ weboffice.Context, args *weboffice.UploadAddressArgs) (*weboffice.UploadAddressReply, error) {
	return t.Demo.UploadAddress(args.FileID)
}

func (t *TestProvider) UploadCallback(ctx weboffice.Context, args *weboffice.UploadCallbackReq) (*weboffice.UploadCallbackReply, error) {
	fileInfo, err := t.Demo.GetFileLatestVersion(ctx, args.Request.FileID)
	if err != nil {
		return nil, err
	}

	return &weboffice.UploadCallbackReply{
		CreateTime: fileInfo.CreateTime,
		CreatorID:  fileInfo.CreatorId,
		ID:         fileInfo.ID,
		ModifierID: fileInfo.ModifierId,
		ModifyTime: fileInfo.ModifyTime,
		Name:       fileInfo.Name,
		Size:       fileInfo.Size,
		Version:    fileInfo.Version,
	}, nil
}

var (
	provider = NewProvider(&TestProvider{
		Demo: demo.NewProvider(),
	})
)

func NewProvider(demo *TestProvider) *weboffice.FullProvider {
	return &weboffice.FullProvider{
		BaseProvider:      demo,
		UserProvider:      demo,
		WatermarkProvider: demo,
		EditProvider:      demo,
		VersionProvider:   demo,
		DownloadProvider:  demo,
	}
}
