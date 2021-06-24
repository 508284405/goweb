package main

import (
	"fmt"
	"reflect"
)

func main() {
	var s = student{age: 15,name: "zhangsan"}
	testReflect(s)
}

type student struct {
	name string
	age int
}

func testReflect(b interface{}) {
	//fmt.Println(reflect.TypeOf(b))
	//rval := reflect.ValueOf(b)
	//fmt.Println(rval)
	//fmt.Printf("type := %T", rval)
	//fmt.Println()
	//fmt.Printf("type := %T", rval)
	//fmt.Println()
	//i := rval.Int()
	//i +=1
	//fmt.Println("type := %T", i)
	fmt.Println(reflect.TypeOf(b))
	rval := reflect.ValueOf(b)

	fmt.Printf("%T",rval)
}
