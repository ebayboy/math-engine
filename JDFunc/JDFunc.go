/**
 * @Author: fanpengfei
 * @Description: 支持数据计算， 不能支持字符串的匹配和查找
 * @File:  JDFunc
 * @Version: 1.0.0
 * @Date: 2020/5/26 14:29
 */

package JDFunc

import (
	"fmt"
	"github.com/ebayboy/math-engine/engine"
)

/* 扩展数据计算（不能处理字符串过滤以及计算）
TODO:
- 支持大于、小于、等于  OK
- 支持字符串查找以及正则匹配 ？？
- 拆出vGID_vSNM_vIP_count_404_5m/vGID_vSNM_vIP_count_5m > 0.5中的变量拆出 vGID_vSNM_vIP_count_404_5m 和 vGID_vSNM_vIP_count_5m
*/

type JDFunc struct {
	regNames []string //registe func names
}

func NewJDFunc() *JDFunc {
	jf := new(JDFunc)
	jf.regNames = make([]string, 0, 10)
	jf.RegisteAllJDFunc()
	return jf
}

func (e JDFunc) ShowNames() {
	for i := 0; i < len(e.regNames); i++ {
		fmt.Println("name:", e.regNames[i])
	}
}

func (e JDFunc) registe(name string, argc int, fun func(...engine.ExprAST) float64) error {
	fmt.Println("registe func: ", name)
	err := engine.RegFunction(name, argc, fun)
	if err != nil {
		return err
	}

	e.regNames = append(e.regNames, name)

	fmt.Println("regNames:", e.regNames)

	return nil
}

/* > */
func GT(expr ...engine.ExprAST) float64 {
	a := engine.ExprASTResult(expr[0])
	b := engine.ExprASTResult(expr[1])

	fmt.Printf("a:%v b:%v\n", a, b)
	if a > b {
		return 1
	}

	return 0
}

/* < */
func LT(expr ...engine.ExprAST) float64 {
	a := engine.ExprASTResult(expr[0])
	b := engine.ExprASTResult(expr[1])

	fmt.Printf("a:%v b:%v\n", a, b)
	if a < b {
		return 1
	}

	return 0
}

/* == */
func EQ(expr ...engine.ExprAST) float64 {
	a := engine.ExprASTResult(expr[0])
	b := engine.ExprASTResult(expr[1])

	fmt.Printf("a:%v b:%v\n", a, b)
	if a == b {
		return 1
	}

	return 0
}

/* >= */
func GE(expr ...engine.ExprAST) float64 {
	a := engine.ExprASTResult(expr[0])
	b := engine.ExprASTResult(expr[1])

	fmt.Printf("a:%v b:%v\n", a, b)
	if a >= b {
		return 1
	}

	return 0
}

func LE(expr ...engine.ExprAST) float64 {
	a := engine.ExprASTResult(expr[0])
	b := engine.ExprASTResult(expr[1])

	fmt.Printf("a:%v b:%v\n", a, b)
	if a <= b {
		return 1
	}

	return 0
}

func init() {
}

func (e JDFunc) RegisteAllJDFunc() error {
	var err error = nil

	err = e.registe("GT", 2, GT)
	if err != nil {
		return err
	}

	err = e.registe("LT", 2, LT)
	if err != nil {
		return err
	}

	err = e.registe("EQ", 2, EQ)
	if err != nil {
		return err
	}

	err = e.registe("GE", 2, GE)
	if err != nil {
		return err
	}

	err = e.registe("LE", 2, LE)
	if err != nil {
		return err
	}

	return err
}
