package main

import (
	"fmt"
	"gopkg.in/mgo.v2"      // go get gopkg.in/mgo.v2
	"gopkg.in/mgo.v2/bson" // go get gopkg.in/mgo.v2/bson
)

type Person struct {
	Name  string
	Phone string
}

func main() {
	session, err := mgo.Dial("123.57.46.103") // 创建dialogue
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)

	c := session.DB("test").C("people")
	err = c.Insert(&Person{"ale", "18223292888"})
	if err != nil {
		panic(err)
	}

	result := Person{}
	err = c.Find(bson.M{"name": "ale"}).One(&result)
	if err != nil {
		panic(err)
	}
	fmt.Println("Phone:", result.Phone)
	fmt.Println(err)
}
