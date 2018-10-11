package test

import (
	"fmt"
	"unicode"
	"sort"
	"math/rand"
	"time"
)

func HomeworkA(n int){


	for i := 1 ; i <=(n+1)/2; i++{
		for j := 1 ;j <= (n+1)/2-i;j++{
			fmt.Print(" ")
		}
		for j:=1;j<= i*2-1;j++{
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i :=(n+1)/2-1;i>=1;i--{
		for j:=1;j<=(n+1)/2-i;j++{
			fmt.Print(" ")
		}
		for j:=(n+1)/2-i;j<=(n+1)/2-2+i;j++{
			fmt.Print("*")
		}

		fmt.Println()
	}


}

func HomeworkB(s string){
	var b = []rune(s)
	var i,j,x ,y int
	for _,v := range b{
		if unicode.IsLetter(v){
			i += 1
		}else if unicode.IsDigit(v){
			j +=1
		}else if unicode.IsSpace(v){
			x += 1
		}else {
			y += 1
		}
	}

	fmt.Printf("字母个数：%d\n数字个数：%d\n空白符个数：%d\n其他字符个数:%d\n",i,j,x,y)

}

type students  struct{
	Name 	string
	sex  	string
	No 		string
	Age 	int
	Results int

}

type sorting [] students



func HomeworkC(){
	st := make(map[string] students,10)
	sv :=make([]students,10)
	var i = 0
	s1 :=students{"huxn1","男","0001",35,75,}
	s2 :=students{"huxn2","男","0002",35,100,}
	s3 :=students{"huxn3","男","0003",35,95,}
	s4 :=students{"huxn4","男","0004",35,45,}
	s5 :=students{"huxn5","男","0005",35,65,}
	s6 :=students{"huxn6","男","0006",35,75,}
	s7 :=students{"huxn7","男","0007",35,85,}
	s8 :=students{"huxn8","男","0008",35,95,}
	s9 :=students{"huxn9","男","0009",35,98,}
	s10 :=students{"huxn10","男","00010",35,87,}

	st["0001"]=s1
	st["0002"]=s2
	st["0003"]=s3
	st["0004"]=s4
	st["0005"]=s5
	st["0006"]=s6
	st["0007"]=s7
	st["0008"]=s8
	st["0009"]=s9
	st["00010"]=s10

	for _,v := range st{
		sv[i] = v
		i++
	}

	sort.Sort(sorting(sv))



	for _, v := range sv{
		fmt.Printf("学号：%s  姓名：%s  年龄%d   成绩：%d\n",v.No,v.Name,v.Age,v.Results)
	}


}

func (I sorting) Len() int {
	return len(I)
}
func (I sorting) Less(i, j int) bool {
	return I[i].Results > I[j].Results
}
func (I sorting) Swap(i, j int) {
	I[i], I[j] = I[j], I[i]
}




func HomeworkD(key,address string){

	if cat(key){
		navigation(address)
	}


}

func cat(key string)(b bool){
	if key =="12345"{
		fmt.Println(" 汽车已启动......您好我的主人！")
		return true
	}
	fmt.Println("您忘记拿钥匙了...")
	return false
}

func navigation(address string){
	fmt.Println("您本次导航的目的址：",address)
	fmt.Println("开始导航....")

	for i:=0; i < 10 ; i++  {
		turnTo(rand.Intn(10))
		time.Sleep(1000 * time.Millisecond)
	}
	fmt.Println("目的址位于道路右边，本次导航结束....")

	

}

func turnTo(a int )  {

	switch  {
	case a <= 3:
		fmt.Println("前方向右转....")
	case 3 < a  && a <= 5:
		fmt.Println("前方向左转....")
	case 5 < a  && a <= 7:
		fmt.Println("前方向右后方转....")
	case 7 < a  && a <= 9:
		fmt.Println("前方向左后方转转....")
	default:
		fmt.Println("前方路口直行....")

	}
	
}