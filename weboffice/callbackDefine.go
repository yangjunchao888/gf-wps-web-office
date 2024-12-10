package weboffice

type CallbackDigest struct {
	Sha1 string `json:"sha1"`
}
type CallbackMustRequest struct {
	FileID   string         `json:"file_id"`
	Name     string         `json:"name"`
	Size     int            `json:"size"`
	Digest   CallbackDigest `json:"digest"`
	IsManual bool           `json:"is_manual"`
}
type CallbackMustHeaders struct {
	Etag string `json:"etag"`
}
type CallbackMustResponse struct {
	StatusCode int                 `json:"status_code"`
	Headers    CallbackMustHeaders `json:"headers"`
}
type CallbackSendBackParams struct {
	Foo string `json:"foo"`
}
