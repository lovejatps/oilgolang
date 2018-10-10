package test

import "fmt"

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
