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

func testSelf() {
	engine.RegFunction("contain", 1, func(expr ...engine.ExprAST) float64 {
		// you can use the index value directly according to the number of parameters
		// without worrying about crossing the boundary.
		// use ExprASTResult to get the result of the ExprAST structure.
		return engine.ExprASTResult(expr[0]) * 2
	})

	exp := "double(6) + 2"
	r, err := engine.ParseAndExec(exp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("double(6) + 2 = %f\n", r) // will print ： double(6) + 2 = 14.000000
}

func main() {
	/*
		s := "1 + 2 * 6 / 4 + (456 - 8 * 9.2) - (2 + 4 ^ 5)"
		// call top level function
		r, err := engine.ParseAndExec(s)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s = %v\n", s, r)

		testSelf()
	*/
	testSelf2()
}
