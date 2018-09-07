package myapp

import (
	"fmt"
	"net/http"
	"os"
	"time"
	"net/url"
)
var staticHandler http.Handler
var filepath string

func init()  {
	filepath ="/Users/huxiaoning/Downloads/app-alli-debug.apk"

	staticHandler = http.FileServer(http.Dir(filepath))
}

/**
下载文件处理
 */

func APPdownload(w http.ResponseWriter, r *http.Request){

	fmt.Print(w, "开始下载...")
	//w.Header().Set("Content-Type", "application/force-download")
	//w.Header().Set("Content-Type","application/octet-stream")
	w.Header().Set("Content-Disposition:", "attachment;filename="+url.QueryEscape(time.Now().Format("20060102150405")+".apk"))
	w.Header().Set("Content-Description","File Transfer")
	w.Header().Set("Content-Transfer-Encoding","binary")
	w.Header().Set("Cache-Control","must-revalidate")



	streamBytes, err := os.Open(filepath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Unable to open and read file : %v", err), 500)
		return
	}
	defer streamBytes.Close()
	http.ServeContent(w, r, filepath, time.Time{}, streamBytes)
}