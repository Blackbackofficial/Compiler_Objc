package main

import (
	"Compiler_Objc/internal/parser"
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

type InfoType struct {
	Name string
	DataType string
	Scope string
}

var global = parser.NewGlobalInfo()

func main()  {
	is, err := antlr.NewFileStream("./test/functions.m")
	if err != nil {
		fmt.Printf("No input file provided")
	}

	lexer := parser.NewObjCLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewObjCParser(stream)
	listener := parser.NewBaseListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Translation_unit())
	//deleteUnused("")

	// Graph
	root := listener.Root
	graph := charts.NewTree()
	graph.AddSeries("main", []opts.TreeData{root}, charts.WithTreeOpts(opts.TreeChart{
		Orient: "TB", Roam: true, Layout: "orthogonal", Left: "0%", Right: "0%",
	}), charts.WithLabelOpts(opts.Label{Show: true, Position: "top", Color: "Black"}))

	f, err := os.Create("graph.html")
	if err != nil {
		log.Println(err)
	}
	err = graph.Render(f)
	if err != nil {
		log.Println(err)
	}

	// sort key
	keys := make([]int, 0, len(global))
	for name := range global {
		keys = append(keys, name)
	}
	sort.Ints(keys) //sort by key
	var tableGlobal = make(map[int]InfoType)
	for q := 0; q < len(global); q++ {
		tableGlobal[q] = InfoType(global[keys[q]])
	}

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

//func deleteUnused(dataType string) {
//	for key, t := range global {
//		if t.DataType == dataType {
//			delete(global, key)
//		}
//	}
//}
