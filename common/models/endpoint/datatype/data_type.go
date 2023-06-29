package datatype

type DataType string

const (
	ApplicationJSON       DataType = "application/json"
	ApplicationJavaScript DataType = "application/javascript"
	ApplicationOgg        DataType = "application/ogg"
	ApplicationZip        DataType = "application/zip"
	ApplicationGZip       DataType = "application/gzip"
	ApplicationXML        DataType = "application/xml"
	ApplicationMSWord     DataType = "application/msword"
)

const (
	AudioACC   DataType = "audio/acc"
	AudioBasic DataType = "audio/basic"
	AudioL24   DataType = "audio/L24"
	AudioMP4   DataType = "audio/mp4"
	AudioMPEG  DataType = "audio/mpeg"
	AudiOgg    DataType = "audio/ogg"
	AudiVorbis DataType = "audio/vorbis"
	AudiWebM   DataType = "audio/webm"
)

const (
	ImageGif    DataType = "image/gif"
	ImageJPEG   DataType = "image/jpeg"
	ImagePJPEG  DataType = "image/pjpeg"
	ImagePNG    DataType = "image/png"
	ImageSVGXML DataType = "image/svg+xml"
	ImageTIFF   DataType = "image/tiff"
	ImageWebP   DataType = "image/webp"
)

const (
	MessageHTTP DataType = "message/http"
)

const (
	ModelExample DataType = "model/example"
	ModelIGS     DataType = "model/iges"
	ModelMSH     DataType = "model/mesh"
	ModelVRML    DataType = "model/vrml"
)

const (
	TextCMD        DataType = "text/cmd"
	TextCSS        DataType = "text/css"
	TextCSV        DataType = "text/csv"
	TextHTML       DataType = "text/html"
	TextJavaScript DataType = "text/javascript"
	TextPlain      DataType = "text/plain"
	TextPHP        DataType = "text/php"
	TextXML        DataType = "text/xml"
	TextMarkDown   DataType = "text/markdown"
)

const (
	VideoMPEG  DataType = "video/mpeg"
	VideoMP4   DataType = "video/mp4"
	VideoOGG   DataType = "video/ogg"
	VideoWebM  DataType = "video/webm"
	Video3GPP  DataType = "video/3gpp"
	Video3GPP2 DataType = "video/3gpp2"
)

func (dt DataType) ToString() string {
	return string(dt)
}
