package general

//* Constans file size
const (
	ImageMaxSize  int64 = 1024000
	FileMaxSize   int64 = 1024000
	MultiPartSize int64 = 1024000
)

//* Constants file format
const (
	MimeTypeImage string = "image"
	MimeTypeVideo string = "video"

	ImageTypeJPEG string = "image/jpeg"
	ImageTypePNG  string = "image/png"

	VideoTypeFLV         string = "video/x-flv"
	VideoTypeMP4         string = "video/mp4"
	VideoTypeMPEGURL     string = "application/x-mpegURL"
	VideoTypeMP2T        string = "video/MP2T"
	VideoType3gpp        string = "video/3gpp"
	VideoTypeQuicktime   string = "video/quicktime"
	VideoTypeMSVideo     string = "video/x-msvideo"
	VideoTypeWMV         string = "video/x-ms-wmv"
	VideoTypeOctetStream string = "application/octet-stream"
)
