package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"time"
)

var aaa string = "加藤";

var (
	a = "aaa"
	b = "bbb"
	c = "ccc"
	d = "ddd"
	e = 11
)

func main() {
	var  aaa  string = "hhh"
	var sss string = "hhh"
	sss = "fff"
	fmt.Println("Hello, 世界")
	fmt.Println(aaa)
	fmt.Println(e);
	fmt.Println(sss);


	a, b := 10, 100
	if a > b {
		fmt.Println("a is larger than b")
	} else if a < b {
		fmt.Println("a is smaller than b")
	} else {
		fmt.Println("a equals b")
	}


		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}

	for f:=8; f<100 ;f--  {
		fmt.Println(f)
		if f == 0 {
			fmt.Println("終わり")
			break
		}



	}


	n := 10
	switch n {
	case 15:
		fmt.Println("FizzBuzz")
	case 5, 10:
		fmt.Println("Buzz")
	case 3, 6, 9:
		fmt.Println("Fizz")
	default:
		fmt.Println(n)
	}
	wer,rrrt := swap(2,6)
	fmt.Println(wer,rrrt)
	type Person struct {
		ID int
		Name string
		Email string
		Age int
		Address string
		memo string
	}

	defer func() {
		err := recover()
		if err != nil {
			// runtime error: index out of range
			log.Fatal(err)
		}
	}()


	person := &Person{
		ID: 1,
		Name: "Gopher",
		Email: "gopher@example.org",
		Age: 5,
		Address: "aaa",
		memo: "golang lover",
	}
	fsfffsss, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}

	say("world")
	go saya("hello")


	fmt.Println(string(fsfffsss)) // 文字列に変換

	http.HandleFunc("/", IndexHandler)
	http.ListenAndServe(":3000", nil)

}

func swap(i, j int) (int, int) {
	return j, i
}

func IndexHandler(w http.ResponseWriter,
	r *http.Request) {

	fmt.Fprint(w, "hello world")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func saya(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(s)
	}
}



