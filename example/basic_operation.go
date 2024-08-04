package main

import (
	"fmt"
	bitcask "kv_project"
)

func main() {
	opts := bitcask.DefaultOptions
	opts.DirPath = "D:\\golang_pro\\kv_project\\tmp"
	db, err := bitcask.Open(opts)
	if err != nil {
		panic(err)
	}

	err = db.Put([]byte("name"), []byte("bitCask"))
	if err != nil {
		panic(err)
	}

	val, err := db.Get([]byte("name"))
	if err != nil {
		panic(err)
	}
	fmt.Println("Val = " + string(val))

	err = db.Delete([]byte("name"))
	if err != nil {
		panic(err)
	}
}
