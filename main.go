package main

import (
	"Compiler_Objc/internal/parser"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"log"
	"os"
)

func main()  {
	is, err := antlr.NewFileStream("./test/2.txt")
	if err != nil {
		fmt.Printf("no input file provided")
	}

	//TODO: сделать информационную таблицу
	lexer := parser.NewObjCLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewObjCParser(stream)
	listener := &parser.BaseObjCListener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Translation_unit() )

	root := listener.Root
	graph := charts.NewTree()
	graph.AddSeries("main", []opts.TreeData{root}, charts.WithTreeOpts(opts.TreeChart{
		Orient: "TB", Roam: true, Layout: "orthogonal",
	}), charts.WithLabelOpts(opts.Label{Show: true, Position: "top", Color: "Black"}))

	f, err := os.Create("graph.html")
	if err != nil {
		log.Println(err)
	}
	err = graph.Render(f)
	if err != nil {
		return 
	}
}
