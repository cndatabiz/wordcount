package main

import (
	"fmt"
	"github.com/dataz.cn/wordcount/tools"
	"io/ioutil"
	"os"

)

func main()  {
	filename := os.Args[1]

	contents , err := ioutil.ReadFile(filename)
	if err != nil{
		fmt.Println(err)
		return
	}

	text := string(contents)
	cnt := tools.WCount(text)

	fmt.Println(text)
	fmt. Printf("There are %d words in your text. \n", cnt)
}
