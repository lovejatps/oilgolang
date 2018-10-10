package test

import (
	"fmt"

)



func T_slice(){
	fmt.Printf("%s\n","切片slice 学习练习。。。 ")
	var numbers [] int
	printSlice(numbers)
	numbers = append(numbers,1)  //向切片中增加一个元素
	numbers =append(numbers,2,3,4,5,6)  //向切片中同时增加多个元素
	printSlice(numbers)
	numberstwo := make([]int,len(numbers),(cap(numbers))*2)
	printSlice(numberstwo)
	copy(numberstwo,numbers)		///* 拷贝 numbers 的内容到 numberstwo */
	printSlice(numberstwo)

}


func T_slicetow () {
	x := []int {1,2,3,4,5,6,7}
	printSlice(x)
	x1 := make([]int ,len(x),(cap(x)*2))
	printSlice(x1)
	copy(x1,x)
	printSlice(x1)

}


func printSlice(x []int ){
	fmt.Printf("len=%d cap=%d slice=%v\n",len(x),cap(x),x)
	var sum = 0
	for num := range x{
		sum = sum + num

	}
	fmt.Printf("合计=%d\n",sum)

	for i,v := range x{
		fmt.Printf("%d -> %d\n",i,v)
	}


}


func test3(x int){
	fmt.Println(100/x )
}


func Test_iota(){

	var a ,b = 1,3
	swap2(&a,&b)
	fmt.Printf("a=%d , b=%d\n",a,b)

	defer fmt.Println("aaaaaaaaaaaaaa")
	defer fmt.Println("vvvvvvvvvvvvv")
	//defer test3(0)
	defer fmt.Println("44444444444")


}

func swap2(c1,c2 *int){
	*c1,*c2 = *c2,*c1;
}


