package main

import (
	"net/http"
	"github.com/tssaini/golang-webserver/controllers"
	"fmt"
)

const (
	// FIVE should be 0
	FIVE = iota
	//SECOND shold be 1
	SECOND = iota
)

func main() {
	// test()
	controllers.RegisterControllers()
	http.ListenAndServe(":3000", nil)

}



func startWebServer(port, retry int) (int, error) {
	fmt.Println("Starting webserver on port:", port)

	fmt.Println("Webserver started", retry)
	// return errors.New("Something went wrong")
	return port, nil
}

func test() {
	// const pi int64 = 3
	fmt.Println(SECOND*7)

	firstname := "taran"
	fmt.Println(firstname)
	ptr := &firstname
	fmt.Println(ptr, *ptr)

	//array
	arr := [3]int{1,2,3}
	fmt.Println(arr)

	//slices
	slice := []int{1,2,3}
	fmt.Println(slice)

	slice = append(slice, 4)
	fmt.Println(slice)

	s2:= slice[1:]
	fmt.Println(s2)

	//map
	m := map[string]int{"a":1, "b":4}
	m["b"]=10
	fmt.Println(m)
	delete(m, "b")
	fmt.Println(m)

	//struct
	type user struct {
		ID int
		FirstName string
		LastName string
	}

	var u user
	u.ID = 1
	u.FirstName = "Arthur"
	u.LastName = "Dent"
	fmt.Println(u)

	u2 := user{ID: 1, FirstName: "Arthur", LastName: "Google"}
	fmt.Println(u2)

	//functions
	port:= 3000
	_, err := startWebServer(port, 3)
	fmt.Println(err)
}
