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
	"github.com/ebayboy/math-engine/JDParse"
	"github.com/ebayboy/math-engine/engine"
)

func main() {
	jf := JDFunc.NewJDFunc()
	jf.ShowNames()
	var err error
	var r float64
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

	//waf demo: vGID_vSNM_vIP_count_404_5m/vGID_vSNM_vIP_count_5m > 0.5
	//TODO: 拆出vGID_vSNM_vIP_count_404_5m/vGID_vSNM_vIP_count_5m > 0.5中的变量拆出 vGID_vSNM_vIP_count_404_5m 和 vGID_vSNM_vIP_count_5m
	fmt.Println("waf demo test:")
	vGID_vSNM_vIP_count_404_5m := 100
	vGID_vSNM_vIP_count_5m := 199

	//tmp := "vGID_vSNM_vIP_count_404_5m/vGID_vSNM_vIP_count_5m > 0.5"

	exp = fmt.Sprintf("GT(%v/%v, 0.5)", vGID_vSNM_vIP_count_404_5m, vGID_vSNM_vIP_count_5m)
	fmt.Println("exp:", exp)
	r, err = engine.ParseAndExec(exp)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v = %d\n", exp, int(r))

	fmt.Println("test JDParse")
	jdParser := JDParse.NewJDParse("GT(vGID_vSNM_vIP_count_404_5m/vGID_vSNM_vIP_count_5m,0.5)")
	fmt.Println("before parse")
	jdParser.Show()

	jdParser.Parse()
	fmt.Println("after parse")
	jdParser.Show()

	jdParser.Compose()
	fmt.Println("after compose")
	jdParser.Show()

	r, err = engine.ParseAndExec(jdParser.ExpParse)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v = %d\n", jdParser.ExpParse, int(r))
}
