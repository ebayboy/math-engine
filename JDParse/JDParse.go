/**
 * @Author: fanpengfei
 * @Description:
 * @File:  JDParse
 * @Version: 1.0.0
 * @Date: 2020/5/26 17:13
 */

package JDParse

import (
	"fmt"
	"strconv"
	"strings"
)

/* TODO:
	- 表达式变量拆分：实现Parse
	- 表达式中间变量重组： 实现Compose
    - 表达式转换： vGID_vSNM_vIP_count_404_5m/vGID_vSNM_vIP_count_5m > 0.5 => GT(vGID_vSNM_vIP_count_404_5m/vGID_vSNM_vIP_count_5m, 0.5)
*/

type JDParser interface {
	//vGID_vSNM_vIP_count_404_5m/vGID_vSNM_vIP_count_5m > 0.5
	//GT(vGID_vSNM_vIP_count_404_5m/vGID_vSNM_vIP_count_5m, 0.5)
	//前缀匹配确认变量vGID_xxx => %v
	//GT(vGID_vSNM_vIP_count_404_5m/vGID_vSNM_vIP_count_5m, 0.5) => GT(%v/%v, 0.5)
	//float64切片数组按照index顺序存储变量
	//vGID[0] = vGID_vSNM_vIP_count_404_5m
	//vGID[1] = vGID_vSNM_vIP_count_5m

	parser()
	compose()
}

//变量拆分字符特征
var g_SplitSeps = " +-*/><=(),"

type JDParse struct {
	expOri     string             // ori
	ExpParse   string             // after parse
	vars       map[string]float64 //store var key：value
	filiterVar []string
	varsNames  []string //store var position in expression
}

//格式转换 vGID_vSNM_vIP_count_404_5m/vGID_vSNM_vIP_count_5m > 0.5 => GT(vGID_vSNM_vIP_count_404_5m/vGID_vSNM_vIP_count_5m, 0.5)
//可以根据运算符号拆，也可以根据
func (J *JDParse) Convert() error {

	return nil
}

func splitSep(r rune) bool {
	for _, sep := range g_SplitSeps {
		if r == sep {
			return true
		}
	}

	return false
}

//Split string with multi seps
func (J *JDParse) splitVars() []string {
	return strings.FieldsFunc(J.expOri, splitSep)
}

func IsNum(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}

/* @in
GT(vGID_vSNM_vIP_count_404_5m/vGID_vSNM_vIP_count_5m, 0.5)
IsNumber: 0.5
i:[0] s:[GD]  过滤关键字
i:[1] s:[vGID_vSNM_vIP_count_404_5m]
i:[2] s:[vGID_vSNM_vIP_count_5m]
i:[3] s:[0.5]  过滤数字
*/
func (J *JDParse) Parse() error {
	fmt.Println("Parse:")

	names := J.splitVars()

	//Parse out vars
	var addFlag bool
	for _, name := range names {
		addFlag = true
		//filter func  words and number
		for _, s := range J.filiterVar {
			//filter func key words && number
			if name == s || IsNum(name) {
				addFlag = false
				break
			}
		}
		if addFlag {
			J.varsNames = append(J.varsNames, name)
			J.vars[name] = 0
		}
	}

	return nil
}

func (J *JDParse) Show() {
	fmt.Println("Show:", " expOri:", J.expOri, " ExpParse:", J.ExpParse)

	for key := range J.vars {
		fmt.Printf("key:[%v] value[%v]\n", key, J.vars[key])
	}

	for i, name := range J.varsNames {
		fmt.Printf("i:[%v] name:[%v]\n", i, name)
	}
}

func (J *JDParse) Compose() error {
	J.ExpParse = J.expOri
	fmt.Println("J.ExpParse:", J.ExpParse)
	for _, name := range J.varsNames {
		fmt.Println("Start J.ExpParse:", J.ExpParse, "name:", name)
		J.ExpParse = strings.Replace(J.ExpParse, name, "%v", 1)
		fmt.Println("Replace J.ExpParse:", J.ExpParse)
		J.ExpParse = fmt.Sprintf(J.ExpParse, J.vars[name])
		fmt.Println("Sprintf J.ExpParse:", J.ExpParse)
	}

	return nil
}

func NewJDParse(expOri string) *JDParse {
	if expOri == "" {
		fmt.Println("Error: expConv:%v", expOri)
		return nil
	}
	jdp := new(JDParse)
	jdp.expOri = expOri
	jdp.vars = make(map[string]float64)
	jdp.filiterVar = []string{"GT", "LT", "EQ", "GE", "LE"}

	return jdp
}
