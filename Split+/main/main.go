package main

import (
	"fmt"

	student ".."
)

func main() {

	str := "<<==123==>>In<<==123==>>1820,<<==123==>>Thomas<<==123==>>de<<==123==>>Colmar<<==123==>>launched<<==123==>>the<<==123==>>mechanical<<==123==>>calculator<<==123==>>industry[note<<==123==>>1]<<==123==>>when<<==123==>>he<<==123==>>released<<==123==>>his<<==123==>>simplified<<==123==>>arithmometer,"
	str2 := "<<==123==>>"
	fmt.Println(student.Split(str, str2))
}
