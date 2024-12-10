package weboffice

type GetFileReq struct {
	FileID string `json:"file_id"` // 文件id
}

type GetFileDownloadReq struct {
	FileID string `json:"file_id"` // 文件id
}

type GetFilePermissionReq struct {
	FileID string `json:"file_id"` // 文件id
}

type GetFileWatermarkReq struct {
	FileID string `json:"file_id"` // 文件id
}

type UploadPrepareReq struct {
	FileID string `json:"file_id"` // 文件id
}

type UploadFileReq struct {
	FileID string `form:"file_id" json:"file_id"`
}

type GetFileVersionsReq struct {
	FileID string `json:"file_id"` // 文件id
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
}

type GetFileVersionReq struct {
	FileID  string `json:"file_id"` // 文件id
	Version int32  `json:"version"`
}

type GetFileVersionDownloadReq struct {
	FileID  string `json:"file_id"` // 文件id
	Version int32  `json:"version"`
}
type DownloadFileReq struct {
	FileID  string `json:"file_id"` // 文件id
	Version int32  `json:"version"`
}

type RenameFileReq struct {
	FileID string `json:"file_id"` // 文件id
}

type GetUsersReq struct {
	UserIDs []string `form:"user_ids" json:"user_ids"`
}
