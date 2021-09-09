package main

import (
	"errors"
	"fmt"
	"reflect"
)

type A struct {
	Color string `json:"color,omitempty"`
}

func (a *A) fa() {
	fmt.Println("A方法...")
}

type B struct {
	A      `json:"a"`
	Size   int     `json:"size,omitempty"`
	Width  float32 `json:"width,omitempty"`
	Height float32 `json:"height,omitempty"`
}

type CommonError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Cause   string `json:"cause,omitempty"`
	Path    string `json:"path,omitempty"`
}

func (error *CommonError) Error() string {
	return fmt.Sprintf("code: %d,message: %s,path: %s", error.Code, error.Message, error.Path)
}

func (b *B) Fb() {
	fmt.Println("B方法...")
}

//遍历结构体属性
func LookupField(v reflect.Value) (err error) {
	defer func() {
		if e := recover(); e != nil {
			if err == nil {
				err = errors.New(e.(string))
			}
		}
	}()
	//panic("panic")
	if (v.Kind() == reflect.Interface || v.Kind() == reflect.Ptr) && v.Elem().Kind() == reflect.Struct {
		v = v.Elem()
	} else if v.Kind() != reflect.Struct {
		return errors.New("Type is not a struct")
	}
	for i := 0; i < v.NumField(); i++ {
		//如果字段是匿名字段且是结构体类型，则进行递归遍历属性
		field := v.Field(i)
		if v.Type().Field(i).Anonymous && (field.Kind() == reflect.Struct || ((field.Kind() == reflect.Interface || field.Kind() == reflect.Ptr) && field.Elem().Kind() == reflect.Struct)) {
			LookupField(field)
		} else {
			fmt.Printf("%s=%v\n", v.Type().Field(i).Name, field.Interface())
		}
	}
	return nil
}

func appendEl(slice *[]int) {
	for i := 0; i < 100; i++ {
		*slice = append(*slice, i)
	}
}

func fc(a []int, str string) {
	fmt.Println("a=", a)
	fmt.Println("str=", str)
}

func main() {
	//b := &B{
	//	A{"red"},// 8B
	//	100, // 8B
	//	20, // 4B
	//	5, // 4B
	//}
	//fmt.Printf("%v\n",b)
	//rtb := reflect.TypeOf(b)
	//for i := 0; i < rtb.Elem().NumField(); i++ {
	//	fmt.Println(rtb.Elem().Field(i).Tag.Get("json"))
	//}
	//rvb := reflect.ValueOf(b)
	//
	//if err := LookupField(rvb);err != nil {
	//	fmt.Println(err)
	//}
	//
	//bytes := make([]byte, 1, 10)
	//bytes = append(bytes, 1,2,3)

	//fmt.Printf("序列化前对象地址：%p",b)
	//if bytes, err := json.Marshal(b); err == nil {
	//	var c B
	//	fmt.Println(string(bytes))
	//	if err := json.Unmarshal(bytes, &c); err != nil {
	//		fmt.Println(err)
	//	}
	//	fmt.Printf("反序列化后的地址：%p\n",b)
	//	fmt.Println("反序列化结果：",c)
	//} else {
	//	fmt.Println(err)
	//}

	//fmt.Println(unsafe.Sizeof("aaaa3422123456789fgefjiehafieoawhefiefe"))

	//var err error = &CommonError{
	//	Code: 400,
	//	Message: "error!",
	//	Cause: "error!",
	//	Path: "reflect/main.go",
	//}
	//
	//fmt.Println(err)
	//
	//slice := make([]int,0,10)
	//appendEl(&slice)
	//fmt.Println(slice)

	//fmt.Println(appendEl)
	rv := reflect.ValueOf(fc)
	rt := reflect.TypeOf(fc)
	values := make([]reflect.Value, 0, rt.NumIn())
	var value reflect.Value
	for i := 0; i < rt.NumIn(); i++ {
		in := rt.In(i)
		if in.Kind() == reflect.String {
			value = reflect.New(in)
			value.Elem().Set(reflect.ValueOf("value"))
			values = append(values, value.Elem())
		} else if in.Kind() == reflect.Slice {
			value = reflect.MakeSlice(in, 3, 3)
			for j := 0; j < value.Len(); j++ {
				el := value.Index(j)
				if el.Kind() == reflect.Int {
					el.Set(reflect.ValueOf(j))
				}
			}
			values = append(values, value)
		}
	}
	rv.Call(values)
}
