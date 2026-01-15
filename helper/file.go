package helper

import (
	"mime"
	"net/http"
	"strings"
)

func GetImgExtFromBytes(imgData []byte) string {
	typeL := getExtsFromBytes(imgData)
	if len(typeL) > 0 {
		return typeL[0]
	}

	return ""
}

func GetImgExtMapFromBytes(imgData []byte) map[string]bool {
	typeL := getExtsFromBytes(imgData)
	extMap := map[string]bool{}
	if len(typeL) > 0 {
		for _, ext := range typeL {
			extMap[ext] = true
		}
	}

	return extMap
}

func getExtsFromBytes(imgData []byte) []string {
	if len(imgData) == 0 {
		return []string{}
	}

	sniff := string(imgData[:min(100, len(imgData))])

	if strings.Contains(sniff, "<svg") {
		return []string{".svg"}
	}

	if len(imgData) >= 12 {
		boxType := string(imgData[4:12])
		if strings.Contains(boxType, "ftypheic") || strings.Contains(boxType, "ftypheix") || strings.Contains(boxType, "ftypmif1") {
			return []string{".heic"}
		}
	}

	mimeType := http.DetectContentType(imgData)
	typeL, _ := mime.ExtensionsByType(mimeType)
	if len(typeL) > 0 {
		return typeL
	}

	return []string{}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func AllowImageType(imgData []byte, allowedExts []string) bool {
	extMap := GetImgExtMapFromBytes(imgData)
	if len(extMap) == 0 {
		return false
	}

	for _, allowed := range allowedExts {
		if extMap["."+strings.ToLower(allowed)] {
			return true
		}
	}
	return false
}
