package errormessage

import (
	"app/infra/logger"
	"encoding/xml"
	"os"
)

const (
	BadRequest  = 1000
	FailAuth    = 1001
	SystemError = 9999
)

type EInfos struct {
	XMLName   xml.Name `xml:"errors"`
	ErrorInfo []EInfo  `xml:"error"`
}

type EInfo struct {
	Code    int    `xml:"code"`
	Status  int    `xml:"status"`
	Message string `xml:"message"`
}

func SettingError() map[int]EInfo {
	infomap := map[int]EInfo{}
	d, err := os.ReadFile("/go/src/app/infra/errormessage/error.xml")
	if err != nil {
		logger.Error("エラーファイルの読み込みに失敗しました")
		return infomap
	}

	var infos EInfos
	err = xml.Unmarshal(d, &infos)
	if err != nil {
		logger.Error("XMLファイルから構造体への変換に失敗しました")
		return infomap
	}

	for _, v := range infos.ErrorInfo {
		infomap[v.Code] = v
	}

	return infomap
}