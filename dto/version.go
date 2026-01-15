package dto

type VersionInfo struct {
	Version     string `json:"version"`     // 版本
	ProjectName string `json:"projectName"` // 專案名稱
	Environment string `json:"environment"` // 啟動環境
}
