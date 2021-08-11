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
	is, err := antlr.NewFileStream("./test/factorial.txt")
	if err != nil {
		fmt.Printf("No input file provided")
	}

	//TODO: сделать информационную таблицу
	lexer := parser.NewObjCLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewObjCParser(stream)
	listener := parser.NewBaseListener()
	global := parser.NewGlobalInfo()
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Translation_unit() )
	log.Printf("%v\n", global)

	root := listener.Root
	graph := charts.NewTree()
	graph.AddSeries("main", []opts.TreeData{root}, charts.WithTreeOpts(opts.TreeChart{
		Orient: "TB", Roam: true, Layout: "orthogonal", Left: "0%", Right: "0%",
	}), charts.WithLabelOpts(opts.Label{Show: true, Position: "top", Color: "Black"}))

	maps := charts.NewMap()
	maps.AddSeries("Maps", []opts.MapData{global})

	f, err := os.Create("graph.html")
	if err != nil {
		log.Println(err)
	}
	err = graph.Render(f)
	if err != nil {
		return
	}
}
