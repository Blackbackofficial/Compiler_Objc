package main
// debug false

import (
	"Compiler_Objc/internal/parser"
	"Compiler_Objc/internal/visualization"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"html/template"
	"log"
	"net/http"
	"os"
	"sort"
)

var global = parser.NewGlobalInfo()

func main()  {
	is, err := antlr.NewFileStream("./test/0.m")
	if err != nil {
		fmt.Printf("No input file provided")
	}

	lexer := parser.NewObjCLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewObjCParser(stream)
	listener := parser.NewBaseListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Translation_unit())
	deleteUnused("")


	// Tree
	root := listener.Root
	tree := charts.NewTree()
	tree.AddSeries("main", []opts.TreeData{root}, charts.WithTreeOpts(opts.TreeChart{
		Orient: "TB", Roam: true, Layout: "orthogonal",  Left: "0%", Right: "0%", ExpandAndCollapse: false, InitialTreeDepth: -1, // поменять глубину если нужно
	}), charts.WithLabelOpts(opts.Label{Show: true, Position: "top", Color: "Black"}))

	f, err := os.Create("./result/tree.html")
	if err != nil {
		log.Println(err)
	}
	err = tree.Render(f)
	if err != nil {
		log.Println(err)
	}


	// sort key
	keys := make([]int, 0, len(global))
	for name := range global {
		keys = append(keys, name)
	}
	sort.Ints(keys) //sort by key
	var tableGlobal = make(map[int]parser.InfoType)
	for q := 0; q < len(global); q++ {
		tableGlobal[q] = global[keys[q]]
	}


	// function call graph
	globalFunc := make(map[int]parser.InfoType)
	allFuncArr := make(map[int]parser.InfoType)
	for key, t := range global {
		if t.DataType == "function" {
			allFuncArr[key] = parser.InfoType{Name: t.Name, DataType: t.DataType, Scope: t.Scope}
		}
		if  t.DataType == "function" && t.Scope == "global" {
			globalFunc[key] = parser.InfoType{Name: t.Name, DataType: t.DataType, Scope: t.Scope}
		}
	}

	visualization.ParseGraph(globalFunc, allFuncArr)


	// Table global & local
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl, _ := template.ParseFiles("templates/index.html")
		err := tmpl.Execute(w, global)
		if err != nil {
			log.Println(err)
		}
	})

	fmt.Println("Server is listening at http://localhost:8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Println(err)
	}
}

func deleteUnused(dataType string) {
	for key, t := range global {
		if t.DataType == dataType || (t.Name == "" && t.Scope == "") {
			delete(global, key)
		}
	}
}