package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type Icon struct {
	Category   string
	CategoryCN string
}

func main() {
	icons := []Icon{
		{"Base", "基础"},
		{"Safe", "安全 & 防护"},
		{"Office", "办公文档"},
		{"Edit", "编辑"},
		{"Emoji", "表情"},
		{"Measurement", "测量 & 试验"},
		{"Abstract", "抽象图形"},
		{"Money", "电商财产"},
		{"Animals", "动物"},
		{"Music", "多媒体音乐"},
		{"Clothes", "服饰"},
		{"Character", "符号标识"},
		{"Industry", "工业"},
		{"Makeups", "化妆美妆"},
		{"Graphics", "几何图形"},
		{"Build", "建筑"},
		{"Arrows", "箭头方向"},
		{"Communicate", "交流沟通"},
		{"Travel", "交通旅游"},
		{"Components", "界面组件"},
		{"Connect", "链接"},
		{"Operate", "美颜调整"},
		{"Baby", "母婴儿童"},
		{"Energy", "能源 & 生命"},
		{"Brand", "品牌"},
		{"Life", "生活"},
		{"Time", "时间日期"},
		{"Foods", "食品"},
		{"Hands", "手势动作"},
		{"Datas", "数据"},
		{"Charts", "数据图表"},
		{"Sports", "体育运动"},
		{"Weather", "天气"},
		{"Constellation", "星座"},
		{"Health", "医疗健康"},
		{"Hardware", "硬件"},
		{"Peoples", "用户人名"},
		{"Game", "游戏"},
		{"Others", "其它"},
	}
	for _, icon := range icons {
		funcNames, _ := listVariables(icon.Category)
		for _, funcName := range funcNames {
			f := icon.Category + "." + funcName + "()"
			fmt.Printf("{%s, \"%s\", \"%s\"},\n", f, icon.CategoryCN, f)
		}
	}
}

func listVariables(pkgName string) ([]string, error) {
	fset := token.NewFileSet()
	srcFile := "./iconpark/" + pkgName + "/icons.go"
	f, err := parser.ParseFile(fset, srcFile, nil, 0)
	if err != nil {
		return nil, err
	}
	variables := make([]string, 0)
	ast.Inspect(f, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.FuncDecl:
			if x.Name.IsExported() {
				name := x.Name.Name
				variables = append(variables, name)
			}
		}
		return true
	})
	return variables, nil
}
