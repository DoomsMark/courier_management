package main

import ("fmt"
		"example.com/golang_project_courier_management/db_conn"
		"log")

func main(){
	conn, err := db_conn.Connection()
	if err != nil{
		log.Fatal(err)
		return
	}
	fmt.Println("Connected with Database")
}