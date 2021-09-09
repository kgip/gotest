package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
)

type User struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

func (user *User) F(i int) error{
	if i>10 {
		return nil
	}
	return errors.New("参数太小了")
}

func F(i int) error{
	if i>10 {
		return nil
	}
	return errors.New("参数太小了")
}

//序列化
func (user  *User) MarshalJSON() (data []byte,err error) {
	return json.Marshal(user)
}
//反序列化
func (user *User) UnmarshalJSON(data []byte) error  {
	fmt.Println(string(data))
	json.Unmarshal(data,user)
	return nil
}

//为反射创建的map添加值
func put(key reflect.Value,value reflect.Value,m reflect.Value) (err error) {
	defer func() {
		if e := recover(); e!= nil{
			err = errors.New(e.(string))
		}
	}()
	if m.Kind() != reflect.Map {
		panic("传入参数不是一个map")
	}
	m.SetMapIndex(key,value)
	return
}

//对方法的某个参数进行解析并赋值
func resolve(in reflect.Type,args map[string]interface{}) reflect.Value{
	switch in.Kind() {
		case reflect.Struct:
			//根据参数类型创建结构体对象
			value := reflect.New(in)
			//遍历结构体对象的属性，并赋值
			for i := 0; i < value.Elem().NumField(); i++ {
				vf := value.Elem().Field(i)
				vf.Set(resolve(vf.Type(), args))
			}
		case reflect.Ptr:
			resolve(in.Elem(),args)
		case reflect.Map:
			////创建一个空map
			//newMap := reflect.MakeMap(in)
			//if args != nil && len(args) >= 1 {
			//	for k, v := range args {
			//		put(reflect.ValueOf(k),reflect.ValueOf(v),newMap)
			//		//reflect.TypeOf(v).ConvertibleTo()
			//		//return newMap
			//	}
			//}
			if in.String() == reflect.TypeOf(args).String() {
				return reflect.ValueOf(args)
			}
		case reflect.Chan,reflect.Interface,reflect.Func,reflect.Slice,reflect.Array:
			return reflect.ValueOf(nil)
		default:   //基本数据类型
			in.FieldAlign()
			//for k, v := range args {
			//	if k == in.Name()
			//}
	}
	return reflect.ValueOf(nil)
}

//方法参数解析
func ResolveMethodArgs(mt reflect.Type,args map[string]interface{}) ([]reflect.Value,error) {
	if mt.Kind() != reflect.Func {
		return nil,errors.New("参数类型不是一个方法")
	}
	argsV := make([]reflect.Value,mt.NumIn())
	for i := 0; i < mt.NumIn(); i++ {
		in := mt.In(i)
		value := resolve(in, args)
		argsV = append(argsV, value)
	}
	return argsV,nil
}

//func main() {
//	//if jsonStr, err := json.Marshal(User{
//	//	Username: "aaa",
//	//	Password: "2ih2r",
//	//}); err != nil {
//	//	fmt.Println(err)
//	//} else {
//	//	fmt.Println(string(jsonStr))
//	//	u := &User{}
//	//	json.Unmarshal(jsonStr,u)
//	//}
//	//rv := reflect.ValueOf(&User{})
//	//method := rv.Method(0)
//	//mt := method.Type()
//	//params := make([]reflect.Value,mt.NumIn())
//	//for i := 0; i < mt.NumIn(); i++ {
//	//
//	//}
//	//rt := reflect.TypeOf(F)
//	//value := reflect.New(rt)
//	//value.Elem().Call([]reflect.Value{reflect.ValueOf(0)})
//
//	//m := make(map[string]int32)
//	//rv := reflect.ValueOf(m)
//	//rv.SetMapIndex(reflect.ValueOf("2"),reflect.ValueOf(int32(5)))
//	//rv.SetMapIndex(reflect.ValueOf("3"),reflect.ValueOf(int32(7)))
//	//rv.Type().String()
//	//fmt.Println(rv.Type().String())
//	//fmt.Println(m)
//}
