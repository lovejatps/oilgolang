package main

import (
	"net/http"
	"os"
	"os/signal"
	"log"
	"time"
	"fmt"
	"net/url"
)

var filepath string
var staticHandler http.Handler
func init()  {
	filepath ="/Users/huxiaoning/Downloads/app-alli-debug.apk"

	staticHandler = http.FileServer(http.Dir(filepath))
}



/**
提供APP下载
*/
func main() {

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	mux := http.NewServeMux()
	mux.HandleFunc("/downloadapp",APPdownload)


	server := http.Server{
		Addr:         ":9001",
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		Handler:mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Print("Server closed under requeset!!")
		} else {
			log.Fatal("Server closed unexpecteed!!")
		}

	}

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