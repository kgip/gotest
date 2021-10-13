package main

import (
	"fmt"
	"unsafe"
)

/**
type slice struct {
	array unsafe.Pointer // 元素指针
	len int // 长度
	cap int // 容量
}
*/
/**
type hmap struct {
	count int
	flags uint8
	B uint8
	noverflow uint16
	hash0 uint32
	buckets unsafe.Pointer
	oldbuckets unsafe.Pointer
	nevacuate uintptr
	extra *mapextra
}
*/
func main() {
	arr := [...]int{1, 2, 5, 6, 9, 0}
	pointer := unsafe.Pointer(&arr)
	fmt.Printf("arr:%p,pointer:%p\n", &arr, pointer)
	i := *(*int)(unsafe.Pointer(uintptr(pointer) + 8))
	fmt.Println(i)
	////fmt.Printf("%T",arr)
	//fmt.Printf("数组首地址：%p\n",&arr[0])
	////strings := make([]string, 0, 20)
	//strings := arr[:]
	//fmt.Printf("slice首地址：%p\n",&strings[0])
	//strings = append(strings, "a", "b", "c")
	//fmt.Println(strings)
	//fmt.Printf("%v\n",uintptr(unsafe.Pointer(&strings)))
	//sp := *(unsafe.Pointer(&strings))
	//fmt.Println(sp)

	//arr := [10]int{1,2,2,4,5,6}
	//for i := 0; i < len(arr); i++ {
	//	n := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&arr)) + uintptr(8*i)))
	//	fmt.Println(n)
	//}
	//m := make(map[string]string, 10)
	//m["key"] = "value1"
	//fmt.Printf("%p\n",m)
	//pointer := unsafe.Pointer(&m)
	////hmap的字段值
	//fmt.Println(**(**int)(pointer))
	//fmt.Println(**(**uint8(unsafe.Pointer())))
	//i := len(m)
	//fmt.Println(unsafe.Pointer(&i))
	//fmt.Println(uintptr(unsafe.Pointer(&i)))

	//str := "hello"

	//slice1 := []int{1,2,3}
	//sliceheader1 := (*reflect.SliceHeader)(unsafe.Pointer(&slice1))
	//sliceheader2 := &reflect.SliceHeader{
	//	Data: sliceheader1.Data,
	//	Len: sliceheader1.Len,
	//	Cap: sliceheader1.Cap << 1,
	//}
	//
	//sliceheader1.Data = uintptr(unsafe.Pointer(&arr))
	//sliceheader1.Len = len(arr)
	//sliceheader1.Cap = sliceheader1.Len << 1
	//
	//slice2 := *(*[]int)(unsafe.Pointer(sliceheader2))
	//fmt.Printf("数组地址：%p\n",&arr)
	//fmt.Printf("%v;%p,len=%d,cap=%d\n",slice1,slice1,len(slice1),cap(slice1))
	//fmt.Printf("%v;%p,len=%d,cap=%d",slice2,slice2,len(slice2),cap(slice2))
}
