package test

import (
	"fmt"
	"os/exec"
	"bufio"
	"io"
	"io/ioutil"
	"time"
)

func Tprgo() {

	var a int
	var ptr *int
	var pptr **int

	a = 3000

	/* 指针 ptr 地址 */
	ptr = &a

	/* 指向指针 ptr 地址 */
	pptr = &ptr

	/* 获取 pptr 的值 */
	fmt.Printf("变量 a = %d\n", a )
	fmt.Printf("指针变量 *ptr = %d\n", ptr )
	fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)

}


/**
c执行shell命令 查看日志
 */

func Redshell(){
	fmt.Print("执行shell命令...")

	tail := "cat /Users/huxiaoning/Downloads/nginx.conf"
	cmd := exec.Command("/bin/bash", "-c", tail)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println(err)
		return
	}
	cmd.Start()
	//创建一个流来读取管道内内容，这里逻辑是通过一行一行的读取的
	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}

	//阻塞直到该命令执行完成，该命令必须是被Start方法开始执行的
	cmd.Wait()

	return

}
func run() {
	cmd := exec.Command("/bin/sh", "-c", "ping 127.0.0.1")
	_, err := cmd.Output()
	if err != nil {
		panic(err.Error())
	}

	if err := cmd.Start(); err != nil {
		panic(err.Error())
	}

	if err := cmd.Wait(); err != nil {
		panic(err.Error())
	}
}

func Tialgo(){
	for {
		go run()
		time.Sleep(1e9)

		cmd := exec.Command("/bin/sh", "-c", `ps -ef | grep -v "grep" | grep "ping"`)
		stdout, err := cmd.StdoutPipe()
		if err != nil {
			fmt.Println("StdoutPipe: " + err.Error())
			return
		}

		stderr, err := cmd.StderrPipe()
		if err != nil {
			fmt.Println("StderrPipe: ", err.Error())
			return
		}

		if err := cmd.Start(); err != nil {
			fmt.Println("Start: ", err.Error())
			return
		}

		bytesErr, err := ioutil.ReadAll(stderr)
		if err != nil {
			fmt.Println("ReadAll stderr: ", err.Error())
			return
		}

		if len(bytesErr) != 0 {
			fmt.Printf("stderr is not nil: %s", bytesErr)
			return
		}

		bytes, err := ioutil.ReadAll(stdout)
		if err != nil {
			fmt.Println("ReadAll stdout: ", err.Error())
			return
		}

		if err := cmd.Wait(); err != nil {
			fmt.Println("Wait: ", err.Error())
			return
		}

		fmt.Printf("stdout: %s", bytes)
	}
}


