package main

import (
"fmt"
"reflect"
)

type MyStruct struct {


	Myfield string `myTag:"blah"`
	Myfield2 int `myTag:"blah2"`
}
func DoReflect2(i interface{}){
	// get a ptr to some instance of some struct
	// gets the type of the struct which the ptr is pointing to
	theType :=reflect.ValueOf(i).Type().Elem()
	fmt.Println(theType.Field(0))
	reflect.ValueOf(i).Elem().Field(0).Set(reflect.ValueOf("HELLO WOLRD"))

	fmt.Println(theType)

}

func DoReflect(i interface{}) {
	// gets the ptr to some list of some type
	fmt.Println(reflect.ValueOf(i).Type())
	fmt.Println(reflect.ValueOf(i).Kind())

	listStructType := reflect.ValueOf(i).Type().Elem().Elem() // gets the struct type that the slice contains first elem refers to the ptr , second elem refer to the element which slice contains
	fmt.Println(listStructType)


	newElement := reflect.New(listStructType) // make a new instance of the struct
	fmt.Println(newElement)
	fmt.Println(reflect.ValueOf(i).Type().Elem())
	mySlice := reflect.MakeSlice(reflect.ValueOf(i).Type().Elem(),0,10) // make new slice which is the same type as what the ptr is pointing to.
	fmt.Println(mySlice)
	DoReflect2(newElement.Interface()) // pass the interface instead of the new element because newelement is a value which is NOT the same as the interface, it contains the interface, but not hte same.
	fmt.Println(newElement)
	mySlice = reflect.Append(mySlice, newElement.Elem()) // appends the newly created element to the slice, notice the newElement is a ptr, thus need to call elem to get the actual value.
	fmt.Println(mySlice)
	reflect.ValueOf(i).Elem().Set(mySlice) // replace original slice with new slice.


	//fmt.Println(listStructType)
	//fmt.Println(newElement)
}

func DoOtherthings(a *[]MyStruct){

	*a = append(*a, MyStruct{Myfield:"asdf"})
}
func main(){

	somestructs:= []MyStruct{}

	DoOtherthings(&somestructs)
	//fmt.Println(somestructs)
	DoReflect(&somestructs)
	fmt.Println(somestructs)
	fmt.Println("working")
}
