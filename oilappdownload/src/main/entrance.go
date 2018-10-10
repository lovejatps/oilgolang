package main

import (
	"net/http"
	"os"
	"time"
	"fmt"
	"net/url"
	"os/exec"
	"test"
)

var filepath string
var staticHandler http.Handler
func init()  {
	filepath ="/usr/local/image/oil/myapp/app-myapp-release.apk"


	staticHandler = http.FileServer(http.Dir(filepath))
}


/**
提供APP下载
*/
func main() {
	//test.T_slice()
	//test.T_slicetow()

	test.HomeworkA(7)
	//
	//quit := make(chan os.Signal)
	//signal.Notify(quit, os.Interrupt)
	//
	//
	//
	//mux := http.NewServeMux()
	//mux.HandleFunc("/downloadapp",APPdownload)
	//mux.HandleFunc("/shellmac",test.Tialgo)
	//
	//
	//server := http.Server{
	//	Addr:         ":9001",
	//	ReadTimeout:  60 * time.Second,
	//	WriteTimeout: 60 * time.Second,
	//	Handler:mux,
	//}
	//
	//err := server.ListenAndServe()
	//if err != nil {
	//	if err == http.ErrServerClosed {
	//		log.Print("Server closed under requeset!!")
	//	} else {
	//		log.Fatal("Server closed unexpecteed!!")
	//	}
	//
	//}

}


/**
c执行shell命令 查看日志
 */

func Redshell(w http.ResponseWriter, r *http.Request){
	fmt.Print("执行shell命令...")

	tail := "tail -f -n2000 /Users/huxiaoning/Downloads/nginx.conf"
	cmd := exec.Command("/bin/bash", "-c", tail)
	output, err := cmd.Output()
	if err != nil{
		fmt.Print(w,err)
	}

	w.Write(output)

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