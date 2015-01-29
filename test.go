package main

import "fmt"

<<<<<<< HEAD
func main(){
    sli := make([]int, 5)
    fmt.Println(sli)
    fmt.Println(len(sli))
    fmt.Println(cap(sli))
	myMap := make(map[string] int, 20)
    fmt.Println(myMap)

    fmt.Println(len(myMap))
=======
var t int

func main() {
	//fmt.Println(bson.M{"name": "mario"})
	// var m = make(map[string]int)
	// fmt.Println(m)
	// m["name"] = 25
	// if a, b := m["name"]; true {
	// 	fmt.Println(a)
	// 	fmt.Println(b)
	// }
	// rand.Seed(time.Now().Unix())
	// fmt.Println(rand.Int())
	fmt.Println("------------")

	fmt.Println(rand.Reader) // &{0 {0 0}}
	b := make([]byte, 32)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		fmt.Println("err isnot nil")
	}

	fmt.Println(b)
	fmt.Println(base64.URLEncoding.EncodeToString(b))

	fmt.Println("------------")
	fmt.Println(t)
>>>>>>> 1857c995ba84ed662a488db024b87916a2beb2e7
}

