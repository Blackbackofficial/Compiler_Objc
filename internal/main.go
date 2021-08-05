package main

import (
	"CompilerAntlr/internal/parser"
	"flag"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/go-echarts/go-echarts/opts"
	"log"
	"os"
)

var (
	inputFlag   string
)

func InitFlags() {
	flag.StringVar(&inputFlag, "i", "", "path to input file")
	flag.Parse()
}


func main() {
	InitFlags()

	is, err := antlr.NewFileStream(inputFlag)
	if err != nil {
		fmt.Printf("no input file provided")
	}

	//TODO: сделать информационную таблицу
	lexer := parser.NewCLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parser.NewCParser(stream)
	listener := &parser.BaseCListener{}
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Function() )

	root := listener.Root
	graph := charts.NewTree()
	graph.AddSeries("main", []opts.TreeData{root}, charts.WithTreeOpts(opts.TreeChart{
		Orient: "TB", Roam: true, Layout: "radial",
	}), charts.WithLabelOpts(opts.Label{Show: true, Position: "top", Color: "Black"}))

	f, err := os.Create("graph.html")
	if err != nil {
		log.Println(err)
	}
	graph.Render(f)
}
