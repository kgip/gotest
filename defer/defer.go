package main

import (
	"fmt"
	"reflect"
)

func TestDefer() (r []int)  {
	//i := 1
	r = make([]int,0,10)
	fmt.Println("r=",r)
	defer func() {
		r = append(r, 5)
		fmt.Println("r=",r)
	}()
	r = append(r,1,2)
    //i++
	return
}

type Stu struct {
	a int8
	b float32
	c bool
}


func main() {
	//fmt.Println(TestDefer())

	//m := make(map[string]string,10)

	//r := make([]int,3)
	//fmt.Printf("%p,%d,%d\n",r,len(r),cap(r))
	//r = append(r, 4)
	//fmt.Printf("%p,%d,%d",r,len(r),cap(r))

	//hash := maphash.Hash{}
	//hash.Write([]byte("hellfheaifeheawiefheiheweafewa"))
	//sum64 := hash.Sum64()
	//fmt.Println(sum64)

	//m := map[string]string{"key1":"value1","key2":"value2"}
	//
	//for key, value := range m {
	//	fmt.Println("key=",key,"value=",value)
	//}
	//

	//stu := Stu{
	//	1,
	//	3,
	//	false,
	//}
	//
	//fmt.Printf("结构体大小：%d,a地址：%p,b地址：%p,c地址：%p",unsafe.Sizeof(stu),&stu.a,&stu.b,&stu.c)


	//r := make([]int,2,10)
	//
	//fmt.Printf("%p\n",r)
	//r = append(r, 5)
	//fmt.Printf("%p",r)
	var i interface{} = 1
	to := reflect.TypeOf(i).ConvertibleTo(reflect.TypeOf([]byte{}))
	fmt.Println(to)
}
