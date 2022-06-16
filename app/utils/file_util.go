package utils

import (
	"fmt"
	"github.com/google/uuid"
	"io/ioutil"
	"net/http"
	"path"
	"path/filepath"
)

//文件下载
func DownloadFile(url string, tempDir string) string {
	_, _fileName := filepath.Split(url)
	fmt.Printf("需要下载的文件名:%s\n", _fileName)
	_ext := path.Ext(url)
	_u4 := uuid.New()
	_downloadFilePath := fmt.Sprintf("%s\\%s%s", tempDir, _u4, _ext)
	_resp, _err := http.Get(url)
	if _err != nil {
		panic(_err)
	}
	defer _resp.Body.Close()
	_data, _err := ioutil.ReadAll(_resp.Body)
	if _err != nil {
		panic(_err)
	}
	ioutil.WriteFile(_downloadFilePath, _data, 0644)
	return _downloadFilePath
}
