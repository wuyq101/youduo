package model

import (
	"fmt"
	"testing"
)

func TestLoadFromFile(t *testing.T) {
	dao := &GameDao{}
	file := "/Users/wuyingqiang/Downloads/徐友多/超人10.19.xlsx"
	list, err := dao.LoadFromExcel(file)
	if err != nil {
		t.Fatal(err)
	}
	for _, g := range list {
		fmt.Printf("%+v\n", g)
	}
}
