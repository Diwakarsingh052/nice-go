package main

import "fmt"

type Conf struct {
	db string // unexported field
}

func NewConf(dataSourceName string) Conf {
	return Conf{db: dataSourceName}
}

func main() {
	c := NewConf("postgres")
	fmt.Println(c)
}
