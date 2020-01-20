package main

import (
	"fmt"
	"sync"
)

// 一个对象只有一个实例
type singletonObject struct {
	Map map[string]string
}
//1. use sync.Once to implement object only once
var (
	objectInstance singletonObject
	once sync.Once
)

func NewObject() *singletonObject {
	once.Do(func(){
		objectInstance = singletonObject{Map: map[string]string{}}
	})
	return &objectInstance
}

func main() {
	object1 := NewObject()
	object1.Map["hello"] = "world"
	fmt.Println(object1.Map)
	object2 := NewObject()
	object2.Map["你好"]= "世界"
	fmt.Println(object2.Map)
	object1.Map["Kon ni chi wa"]= "Sei Kai"
	fmt.Println(object1.Map)
}
/*
map[hello:world]
map[hello:world 你好:世界]
map[Kon ni chi wa:Sei Kai hello:world 你好:世界]
*/


