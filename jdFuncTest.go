/**
 * @Author: fanpengfei
 * @Description:
 * @File:  jdFuncTest
 * @Version: 1.0.0
 * @Date: 2020/5/26 15:22
 */

package main

import (
	"fmt"
	"github.com/ebayboy/math-engine/JDFunc"
	"github.com/ebayboy/math-engine/engine"
)

func main() {
	jf := JDFunc.NewJDFunc()
	jf.ShowNames()
	var err error
	var r float64
	r, err = engine.ParseAndExec("double(6) + 2")
	if err != nil {
		panic(err)
	}
	fmt.Printf("double(6) + 2 = %f\n", r)

	var exp string

	//简单大于、小于、等于计算测试
	fmt.Println("Simple test:")
	exp = "GT(4, 3)"
	r, err = engine.ParseAndExec(exp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v = %d\n", exp, int(r))

	exp = "LT(3, 4)"
	r, err = engine.ParseAndExec(exp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v = %d\n", exp, int(r))

	exp = "EQ(3, 3)"
	r, err = engine.ParseAndExec(exp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v = %d\n", exp, int(r))

	exp = "GE(5, 4)"
	r, err = engine.ParseAndExec(exp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v = %d\n", exp, int(r))

	exp = "LE(3, 4)"
	r, err = engine.ParseAndExec(exp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v = %d\n", exp, int(r))

	//复杂 大于、小于、等于测试
	fmt.Println("Complex test:")
	exp = "GE((2+2)*2 - 1, 7)"
	r, err = engine.ParseAndExec(exp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v = %d\n", exp, int(r))

	exp = "GE((2+2)*2 - 1, 6)"
	r, err = engine.ParseAndExec(exp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v = %d\n", exp, int(r))

	exp = "GE((2+2)*2 - 1, 8)"
	r, err = engine.ParseAndExec(exp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v = %d\n", exp, int(r))
}
