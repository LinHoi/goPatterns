package main

import "fmt"

//工厂模式(Factory Method)
/*
抽象工厂负责定义产品公共接口
具体工厂负责 生成具体 产品
延迟产品的实例化
*/

//1. Define Factory Method
type IFactoryMethod interface {
	Method1()
	Method2()
}

//2. Define Product
type ProductType int
const (
	ProductType1 ProductType  =  1  << iota
	ProductType2
)
type Product1 struct {

}
type Product2 struct {

}


//3. Generate Product by ConCreteFactory
func ConcreteFactory(productType ProductType) IFactoryMethod {
	switch productType{
	case ProductType1:
		return &Product1{}
	case ProductType2:
		return &Product2{}
	default:
		return nil
	}

}

//4. Implement Factory Method in Every Product
func (p Product1) Method1() {
	fmt.Println("implement method1 in Product1")
}
func (p Product1) Method2() {
	fmt.Println("implement method2 in Product1")
}

func (p Product2) Method1() {
	fmt.Println("implement method1 in Product2")
}
func (p Product2) Method2() {
	fmt.Println("implement method2 in Product2")
}


func main() {
	product1 := ConcreteFactory(ProductType1)
	product2 := ConcreteFactory(ProductType2)
	product1.Method1()
	product1.Method2()
	product2.Method1()
	product2.Method2()
}
/*
implement method1 in Product1
implement method2 in Product1
implement method1 in Product2
implement method2 in Product2
*/