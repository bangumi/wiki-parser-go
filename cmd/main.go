package main

import (
	"fmt"

	"github.com/bangumi/wiki-parser-go"
)

func main() {
	_, err := wiki.Parse(`{{Infobox animanga/Manga
|中文名= 坐在旁边的家伙用这种眼神看着我
|别名={邻座的家伙用瑟瑟的眼神盯着我的故事
}
|作者= mmk
|作画= 
}}`)
	if err == nil {
		panic("expecting error")
	}

	fmt.Println(err.Error())

	se := err.(*wiki.SyntaxError)

	fmt.Println(se.ReadableError())
}
