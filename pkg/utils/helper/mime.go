package helper

import "strings"

func GetMimetype(path string) string {
	// text types
	if strings.HasSuffix(path, ".txt") {
		return "text/plain"
	}

	if strings.HasSuffix(path, ".html") || strings.HasSuffix(path, ".htm") {
		return "text/html"
	}

	if strings.HasSuffix(path, ".js") {
		return "text/javascript"
	}

	if strings.HasSuffix(path, ".css") {
		return "text/css"
	}

	// image types
	if strings.HasSuffix(path, ".png") {
		return "image/png"
	}

	if strings.HasSuffix(path, ".gif") {
		return "image/gif"
	}

	if strings.HasSuffix(path, ".jpg") || strings.HasSuffix(path, ".jpeg") {
		return "image/jpeg"
	}

	if strings.HasSuffix(path, ".svg") {
		return "image/svg+xml"
	}

	if strings.HasSuffix(path, ".webp") {
		return "image/webp"
	}

	// audio types
	if strings.HasSuffix(path, ".wav") || strings.HasSuffix(path, ".wave"){
		return "audio/wav"
	}

	if strings.HasSuffix(path, ".ogg") {
		return "audio/ogg"
	}

	if strings.HasSuffix(path, ".mp3") {
		return "audio/mpeg3"
	}

	// video types
	if strings.HasSuffix(path, ".mp4") {
		return "video/mp4"
	}

	if strings.HasSuffix(path, ".mov") {
		return "video/quicktime"
	}

	if strings.HasSuffix(path, ".avi") {
		return "video/x-msvideo"
	}

	if strings.HasSuffix(path, ".wmv") {
		return "video/x-ms-wmv"
	}

	// application types
	if strings.HasSuffix(path, ".pdf") {
		return "application/pdf"
	}

	if strings.HasSuffix(path, ".xml") {
		return "application/xml"
	}

	if strings.HasSuffix(path, ".json") {
		return "application/json"
	}

	// default type
	return "application/octet-stream"
}
