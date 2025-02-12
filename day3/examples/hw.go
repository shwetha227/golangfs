package main
import "fmt"
type student struct{
	Name string
	Rgno int
	Dept string
}
func main(){
   st := student{Name: "Student1", Rgno:12,Dept:"CS"}
   fmt.Println("Name:",st.Name,"\nRegister number:",st.Rgno,"\ndepartment:",st.Dept)
}
func ifelseDemo(){
	var a int
	var b int
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	if a == b {
		fmt.Println("two numbers are Equal")
	}else if a < b{
		fmt.Println("a is less")
	}else{
		fmt.Println("a is greater")
	}
}
func forThreeVarDemo(){
	sum := 0
	var i int
	for i=0 ; i<=5 ;i++{
		sum = sum+i
	}
	fmt.Println(sum)
}
func forCondiDemo(){
	n := 1;
    for n < 5{
		n*=2
	}
	fmt.Println(n)
}
func forPythonStyle(){
	strings := []string{"hello","world","golang","NIE"}
	for i,s := range strings{
		fmt.Println(i,s)
	}
}
