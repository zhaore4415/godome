package main

import (
	"fmt"
	"mycrm/dbmysql"
	"mycrm/calutils"
)

func main() {
	fmt.Println("test02")
	dbmysql.GetConn()
  bbb :=calutils.Add(2,3)
	fmt.Println(bbb)
}
