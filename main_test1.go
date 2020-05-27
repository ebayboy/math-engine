/**
 * @Author: fanpengfei
 * @Description:
 * @File:  main_test1
 * @Version: 1.0.0
 * @Date: 2020/5/26 13:07
 */

package main

import (
	"fmt"
	"github.com/ebayboy/math-engine/engine"
)

func main() {

	s := "0/0 + 2"
	// call top level function
	r, err := engine.ParseAndExec(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s = %v\n", s, r)

}
