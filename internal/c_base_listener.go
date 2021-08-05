// Generated from C.g4 by ANTLR 4.7.

package internal // C

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/go-echarts/go-echarts/charts"
	"github.com/go-echarts/go-echarts/opts"
	"log"
)

type Tree struct {
	charts.BaseConfiguration
}


func (c *Tree) AddSeries(name string, data []opts.TreeData, options ...charts.SeriesOpts) *Tree {
	return c
}

// BaseCListener is a complete listener for a parse tree produced by CParser.
type BaseCListener struct{
	Tree Tree
	Root opts.TreeData
	nodes []*opts.TreeData
	current *opts.TreeData
}

func NewBaseListener() *BaseCListener {
	l := BaseCListener{
		Tree:  Tree{},
		nodes: []*opts.TreeData{},
	}
	return &l
}

var _ CListener = &BaseCListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCListener) EnterEveryRule(ctx antlr.ParserRuleContext) {

}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFunction is called when production function is entered.
func (s *BaseCListener) EnterFunction(ctx *FunctionContext) {
	s.Root = opts.TreeData{Name: fmt.Sprintf("Function: %s", ctx.GetChild(1))}
	s.current = &s.Root
	s.nodes = append(s.nodes, &s.Root)
}

// ExitFunction is called when production function is exited.
func (s *BaseCListener) ExitFunction(ctx *FunctionContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseCListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseCListener) ExitDeclaration(ctx *DeclarationContext) {
	node := opts.TreeData{Name: "Declaration"}
	left := opts.TreeData{Name: fmt.Sprintf("Type: %s", ctx.GetChild(0).GetChild(0))}
	right := opts.TreeData{Name: fmt.Sprintf("ID: %s", ctx.GetChild(1))}
	node.Children = []*opts.TreeData{&left, &right}
	s.current.Children = append(s.current.Children, &node)

}

// EnterArgList is called when production argList is entered.
func (s *BaseCListener) EnterArgList(ctx *ArgListContext) {}

// ExitArgList is called when production argList is exited.
func (s *BaseCListener) ExitArgList(ctx *ArgListContext) {}

// EnterArg is called when production arg is entered.
func (s *BaseCListener) EnterArg(ctx *ArgContext) {}

// ExitArg is called when production arg is exited.
func (s *BaseCListener) ExitArg(ctx *ArgContext) {
}

// EnterTypeId is called when production typeId is entered.
func (s *BaseCListener) EnterTypeId(ctx *TypeIdContext) {}

// ExitTypeId is called when production typeId is exited.
func (s *BaseCListener) ExitTypeId(ctx *TypeIdContext) {}

// EnterIDList is called when production iDList is entered.
func (s *BaseCListener) EnterIDList(ctx *IDListContext) {}

// ExitIDList is called when production iDList is exited.
func (s *BaseCListener) ExitIDList(ctx *IDListContext) {}

// EnterCompoundStatement is called when production compoundStatement is entered.
func (s *BaseCListener) EnterCompoundStatement(ctx *CompoundStatementContext) {}

// ExitCompoundStatement is called when production compoundStatement is exited.
func (s *BaseCListener) ExitCompoundStatement(ctx *CompoundStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseCListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseCListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterStatementList is called when production statementList is entered.
func (s *BaseCListener) EnterStatementList(ctx *StatementListContext) {}

// ExitStatementList is called when production statementList is exited.
func (s *BaseCListener) ExitStatementList(ctx *StatementListContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseCListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseCListener) ExitStatement(ctx *StatementContext) {
}

// EnterForStatement is called when production forStatement is entered.
func (s *BaseCListener) EnterForStatement(ctx *ForStatementContext) {
	node := opts.TreeData{Name: "FOR"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitForStatement is called when production forStatement is exited.
func (s *BaseCListener) ExitForStatement(ctx *ForStatementContext) {

	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterOptionExpression is called when production optionExpression is entered.
func (s *BaseCListener) EnterOptionExpression(ctx *OptionExpressionContext) {}

// ExitOptionExpression is called when production optionExpression is exited.
func (s *BaseCListener) ExitOptionExpression(ctx *OptionExpressionContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BaseCListener) EnterWhileStatement(ctx *WhileStatementContext) {
	node := opts.TreeData{Name: "While Loop"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
	//s.Root = node
}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BaseCListener) ExitWhileStatement(ctx *WhileStatementContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseCListener) EnterIfStatement(ctx *IfStatementContext) {
	node := opts.TreeData{Name: "IF statement"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)

}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseCListener) ExitIfStatement(ctx *IfStatementContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterElseStatement is called when production elseStatement is entered.
func (s *BaseCListener) EnterElseStatement(ctx *ElseStatementContext) {
	node := opts.TreeData{Name: "ELSE part"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitElseStatement is called when production elseStatement is exited.
func (s *BaseCListener) ExitElseStatement(ctx *ElseStatementContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterExpression is called when production expression is entered.
func (s *BaseCListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseCListener) ExitExpression(ctx *ExpressionContext) {
	log.Print(ctx.GetText())
	childs := ctx.GetChildCount()
	if childs == 3 {
		md := opts.TreeData{Name: fmt.Sprintf("Assigment: %s", ctx.GetChild(1))}
		left := opts.TreeData{Name: fmt.Sprintf("LValue: %s", ctx.GetChild(0))}
		right := opts.TreeData{Name: fmt.Sprintf("RValue: %s",
			ctx.GetChild(2).GetChild(0))}

		md.Children = append(md.Children, &left)
		md.Children = append(md.Children, &right)
		s.current.Children = append(s.current.Children, &md)

	}
}

// EnterRValue is called when production rValue is entered.
func (s *BaseCListener) EnterRValue(ctx *RValueContext) {}

// ExitRValue is called when production rValue is exited.
func (s *BaseCListener) ExitRValue(ctx *RValueContext) {}

// EnterCompare is called when production compare is entered.
func (s *BaseCListener) EnterCompare(ctx *CompareContext) {}

// ExitCompare is called when production compare is exited.
func (s *BaseCListener) ExitCompare(ctx *CompareContext) {}

// EnterMagnitude is called when production magnitude is entered.
func (s *BaseCListener) EnterMagnitude(ctx *MagnitudeContext) {}

// ExitMagnitude is called when production magnitude is exited.
func (s *BaseCListener) ExitMagnitude(ctx *MagnitudeContext) {}

// EnterFactor is called when production factor is entered.
func (s *BaseCListener) EnterFactor(ctx *FactorContext) {}

// ExitFactor is called when production factor is exited.
func (s *BaseCListener) ExitFactor(ctx *FactorContext) {}

// EnterTerm is called when production term is entered.
func (s *BaseCListener) EnterTerm(ctx *TermContext) {}

// ExitTerm is called when production term is exited.
func (s *BaseCListener) ExitTerm(ctx *TermContext) {}