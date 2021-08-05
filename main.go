package main

import (
	"Compiler_Objc/internal/parser"
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
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
}
