package parser // Package parser ObjC

import (
	"container/list"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"log"
)

// STRUCTS
type Tree struct {
	charts.BaseConfiguration
}

type InfoType struct {
	Name string
	DataType string
	Scope string
}

type Arr struct {
	Name string
	Level int
}

type Flags struct {
	local int // local {...}
	declaratorSuffix bool // function parameter (...)
	typeSpecifier bool // for int long int char e.t.c.
	classInterface bool // @interface
	superclassName bool // ignore NSObject
	categoryName bool // ignore CategorizedComplex
	specifierQualifierList bool
	labeledStatement int // case & default
	pointer bool // *
	function bool // for functions
	iterationStatement bool
	instanceVariables bool
	methodType bool // for -(int) print
	methodSelector bool
	property bool // @property
	functionIn bool // for functions in {...}
	initDeclaratorList bool
	expression bool
	doWhile bool
	For bool
}

type QFlags struct {
	queue_while bool
	queue_multiplicative bool
	queue_initDeclarator bool
	queue_superclass bool
	queue_conditional bool
	queue_logical_or bool
	queue_logical_and bool
	queue_inclusive_or bool
	queue_exclusive_or bool
	queue_and bool
	queue_equality bool
	queue_relational bool
	queue_shift bool
	queue_additive bool
	queue_cast bool
	queue_unary bool
	queue_postfix bool
	queue_unary_operator bool
	queue_argument bool
	queue_assignment_operator bool
	queue_argument_expression_list bool
}

// BaseObjCListener is a complete listener for a parse tree produced by ObjCParser.
type BaseObjCListener struct {
	Tree Tree
	Root *opts.TreeData
	nodes []*opts.TreeData
	current *opts.TreeData
	Flags Flags
	QFlags QFlags
}

// VARS
var start = false
var debug = false // debug
var arrDeep []Arr
var _ ObjCListener = &BaseObjCListener{}
var globalHash = make(map[int]InfoType)
var visHash = make(map[int]InfoType)
var count = 0
var ne = 0
// Queue
var queue_while = list.New()
var queue_multiplicative = list.New()
var queue_initDeclarator = list.New()
var queue_superclass = list.New()
var queue_conditional = list.New()
var queue_logical_or = list.New()
var queue_logical_and = list.New()
var queue_inclusive_or = list.New()
var queue_exclusive_or = list.New()
var queue_and = list.New()
var queue_equality = list.New()
var queue_relational = list.New()
var queue_shift = list.New()
var queue_additive = list.New()
var queue_cast = list.New()
var queue_unary = list.New()
var queue_postfix = list.New()
var queue_unary_operator = list.New()
var queue_argument = list.New()
var queue_assignment_operator = list.New()
var queue_argument_expression_list =list.New()

var openscope = false
var openquotes = false

var closecopee = false
var ifelse = false

func NewBaseListener() *BaseObjCListener {
	l := BaseObjCListener {
		Tree:  Tree{},
		nodes: []*opts.TreeData{},
		Flags: Flags{},
	}
	return &l
}

func NewGlobalInfo() map[int]InfoType {
	return globalHash
}

// gluing string
func gluing() string {
	result := ""
	for key, v := range arrDeep {
		if key == 0 {
			result += v.Name
			continue
		}
		result += "." + v.Name
	}
	return result
}

func typeSpecifier(s *BaseObjCListener) {
	if s.Flags.typeSpecifier == true {
		ne = ne + 2
	} else if ne > 0 && !s.Flags.pointer {
		ne--
	}
}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseObjCListener) VisitTerminal(node antlr.TerminalNode) {
	log.Println(node.GetText(), node.GetSymbol().GetTokenType())
	//log.Println("GLOBAL: ", globalHash)
	//log.Println("LOCAL: ", visHash)
	typeSpecifier(s)
	if node.GetSymbol().GetTokenType() == 82 || node.GetSymbol().GetTokenType() == 81 ||
		node.GetSymbol().GetTokenType() == 88 || node.GetSymbol().GetTokenType() == 89 {
		queue_while.PushBack(node.GetText())
	}
	if  s.Flags.initDeclaratorList && node.GetSymbol().GetTokenType() == 80 {
		queue_initDeclarator.PushBack(node.GetText())
	}
	if node.GetSymbol().GetTokenType() == 86 {
		queue_superclass.PushBack(node.GetText())
	}

	// assignment
	if s.Flags.expression {
		// conditional_expression
		if node.GetSymbol().GetTokenType() == 85 || node.GetSymbol().GetTokenType() == 86 {
			queue_conditional.PushBack(node.GetText())
			//log.Println(queue_conditional)
		}
		// logical_or_expression
		if node.GetSymbol().GetTokenType() == 92 {
			queue_logical_or.PushBack(node.GetText())
			//log.Println(queue_logical_or)
		}
		// logical_and_expression
		if node.GetSymbol().GetTokenType() == 91 {
			queue_logical_and.PushBack(node.GetText())
			//log.Println(queue_logical_and)
		}
		// inclusive_or_expression
		if node.GetSymbol().GetTokenType() == 100 {
			queue_inclusive_or.PushBack(node.GetText())
			//log.Println(queue_inclusive_or)
		}
		// exclusive_or_expression
		if node.GetSymbol().GetTokenType() == 101 {
			queue_exclusive_or.PushBack(node.GetText())
			//log.Println(queue_exclusive_or)
		}
		// and_expression
		if node.GetSymbol().GetTokenType() == 99 {
			queue_and.PushBack(node.GetText())
			//log.Println(queue_and)
		}
		// equality_expression
		if node.GetSymbol().GetTokenType() == 90 || node.GetSymbol().GetTokenType() == 87 {
			queue_equality.PushBack(node.GetText())
			//log.Println(queue_equality)
		}
		// relational_expression
		if node.GetSymbol().GetTokenType() == 82 || node.GetSymbol().GetTokenType() == 81 ||
			node.GetSymbol().GetTokenType() == 88 || node.GetSymbol().GetTokenType() == 89 {
			queue_relational.PushBack(node.GetText())
			//log.Println(queue_relational)
		}
		// shift_expression
		if node.GetSymbol().GetTokenType() == 104 || node.GetSymbol().GetTokenType() == 103 {
			queue_shift.PushBack(node.GetText())
			//log.Println(queue_shift)
		}
		// additive_expression
		if node.GetSymbol().GetTokenType() == 95 || node.GetSymbol().GetTokenType() == 96 {
			queue_additive.PushBack(node.GetText())
			//log.Println(queue_additive)
		}
		// multiplicative_expression
		if node.GetSymbol().GetTokenType() == 97 || node.GetSymbol().GetTokenType() == 98 || node.GetSymbol().GetTokenType() == 102 {
			queue_multiplicative.PushBack(node.GetText())
			//log.Println(queue_multiplicative)
		}
		// unary_expression
		if node.GetSymbol().GetTokenType() == 93 || node.GetSymbol().GetTokenType() == 94 {
			queue_unary.PushBack(node.GetText())
			//log.Println(queue_unary.Front())
		}
		// postfix_expression
		if node.GetSymbol().GetTokenType() == 73 || node.GetSymbol().GetTokenType() == 74 ||
			node.GetSymbol().GetTokenType() == 69 || node.GetSymbol().GetTokenType() == 70 ||
			node.GetSymbol().GetTokenType() == 77 || node.GetSymbol().GetTokenType() == 78 {
			queue_postfix.PushBack(node.GetText())
			//log.Println(queue_postfix)
		}
		// unary_operator
		if node.GetSymbol().GetTokenType() == 99 || node.GetSymbol().GetTokenType() == 97 ||
			node.GetSymbol().GetTokenType() == 96 || node.GetSymbol().GetTokenType() == 84 ||
			node.GetSymbol().GetTokenType() == 83 {
			queue_unary_operator.PushBack(node.GetText())
			//log.Println(queue_unary_operator)
		}
		// assignment_operator
		if node.GetSymbol().GetTokenType() == 80 || node.GetSymbol().GetTokenType() == 107 ||
			node.GetSymbol().GetTokenType() == 108 || node.GetSymbol().GetTokenType() == 112 ||
			node.GetSymbol().GetTokenType() == 105 || node.GetSymbol().GetTokenType() == 106 ||
			node.GetSymbol().GetTokenType() == 113 || node.GetSymbol().GetTokenType() == 114 ||
			node.GetSymbol().GetTokenType() == 109 || node.GetSymbol().GetTokenType() == 111 ||
			node.GetSymbol().GetTokenType() == 110 {
			queue_assignment_operator.PushBack(node.GetText())
			//log.Println(queue_assignment_operator)
		}
		// argument_expression_list
		if node.GetSymbol().GetTokenType() == 76 {
			queue_argument_expression_list.PushBack(node.GetText())
			//log.Println(queue_argument_expression_list)
		}
	}

	// for exp (...)
	if node.GetSymbol().GetTokenType() == 69 {
		openscope = true
	} else if node.GetSymbol().GetTokenType() == 70 {
		openscope = false
		if s.Flags.For && !openscope {
			s.Flags.For = false
			log.Println("EEEEEEEEEEDDDDDDDDDD")
			closecopee = true
		}
	}
	// if exp
	if node.GetSymbol().GetTokenType() == 37 {
		ifelse = true
	}
	// for exp [...]
	if node.GetSymbol().GetTokenType() == 73 {
		openquotes = true
	} else if node.GetSymbol().GetTokenType() == 74 {
		openquotes = false
	}

	// while_statement | do_statement | for_statement | for_in_statement | if | switch case;
	if node.GetSymbol().GetTokenType() == 43 || node.GetSymbol().GetTokenType() == 58 ||
		node.GetSymbol().GetTokenType() == 37 || node.GetSymbol().GetTokenType() == 64 || node.GetSymbol().GetTokenType() == 35 {
		arrDeep = append(arrDeep, Arr{node.GetText(), s.Flags.local})
	}


	var m = NewGlobalInfo()
	e, ok := m[count]
	if !ok {
		e = InfoType{}
	}
	hash, ok := visHash[count]
	if !ok {
		hash = InfoType{}
	}

	// GLOBAL VARS & CLASS
	if s.Flags.local == 0 && !s.Flags.superclassName && !s.Flags.categoryName && !s.Flags.classInterface && node.GetSymbol().GetTokenType() != 67 {
		if node.GetSymbol().GetTokenType() == 125 {
			if s.Flags.declaratorSuffix {
				e.Scope = arrDeep[0].Name  + "(FunctionParameter)"
				hash.Scope = arrDeep[0].Name
			} else if s.Flags.classInterface {
				e.Scope = "classInterface"
				hash.Scope = gluing()
			} else {
				e.Scope = "global"
				hash.Scope = "global" // функции всегда глобальны
			}
			e.Name = node.GetText()
			hash.Name = node.GetText()

			m[count] = e
			visHash[count] = hash
			count++
		} else if node.GetSymbol().GetTokenType() == 69 && !s.Flags.iterationStatement {
			e, ok := m[count-1]
			if !ok {
				log.Fatal("No match key!")
			}
			hash, ok := visHash[count-1]
			if !ok {
				log.Fatal("No match key in hash!")
			}
			e.DataType = "function"
			hash.DataType = "function"
			m[count-1] = e
			visHash[count-1] = hash
		} else if s.Flags.typeSpecifier {
			e.DataType = node.GetText()
			hash.DataType = node.GetText()
			m[count] = e
			visHash[count] = hash
		}
	} else if s.Flags.local == 0 && !s.Flags.superclassName && s.Flags.classInterface && !s.Flags.categoryName {
		if node.GetSymbol().GetTokenType() > 0 && node.GetSymbol().GetTokenType() < 22 && node.GetSymbol().GetTokenType() != 2 &&
			node.GetSymbol().GetTokenType() != 7 && node.GetSymbol().GetTokenType() != 21 && node.GetSymbol().GetTokenType() != 18 &&
			node.GetSymbol().GetTokenType() != 19  {
			e.DataType = node.GetText()
			hash.DataType = node.GetText()
			m[count] = e
			visHash[count] = hash
		} else if s.Flags.specifierQualifierList {
			e.DataType = node.GetText()
			hash.DataType = node.GetText()
			m[count] = e
			visHash[count] = hash
		} else if s.Flags.typeSpecifier {
			e.DataType = node.GetText()
			hash.DataType = node.GetText()
			m[count] = e
			visHash[count] = hash
		} else if node.GetSymbol().GetTokenType() == 125 {
			if s.Flags.classInterface && !s.Flags.methodSelector && !s.Flags.property {
				e.Scope = "global class"
				hash.Scope = "global"
			} else if s.Flags.methodSelector {
				e.Scope = arrDeep[0].Name + "(Method)"
				hash.Scope = arrDeep[0].Name
			} else if s.Flags.classInterface && s.Flags.property {
				e.Scope = arrDeep[0].Name + "(Property)"
				hash.Scope = arrDeep[0].Name
				s.Flags.property = false
			}
			e.Name = node.GetText()
			hash.Name = node.GetText()
			m[count] = e
			visHash[count] = hash
			count++
		}
	// LOCAL VARS & CLASS
	} else if s.Flags.local > 0 && !s.Flags.superclassName && !s.Flags.categoryName && !s.Flags.classInterface && (ne > 0 || s.Flags.functionIn) {
		if node.GetSymbol().GetTokenType() == 125 && s.Flags.typeSpecifier {
			e.DataType = node.GetText()
			hash.DataType = node.GetText()
			m[count] = e
			visHash[count] = hash
		} else if s.Flags.function { // for func1(); func2(); func3()...
			e.Name = e.DataType
			e.DataType = "function"
			e.Scope = gluing()
			hash.Name = e.DataType
			hash.DataType = "function"
			hash.Scope = gluing()
			m[count] = e
			visHash[count] = hash
			count++
		} else if node.GetSymbol().GetTokenType() == 125 && !s.Flags.typeSpecifier && !s.Flags.function && !s.Flags.functionIn {
			e.Scope = gluing()
			e.Name = node.GetText()
			hash.Scope = gluing()
			hash.Name = node.GetText()
			m[count] = e
			visHash[count] = hash
			count++
		} else if node.GetSymbol().GetTokenType() == 125 && !s.Flags.typeSpecifier && s.Flags.functionIn {
			e.Scope = gluing()
			e.Name = node.GetText()
			e.DataType = "function"
			hash.Scope = gluing()
			hash.Name = node.GetText()
			hash.DataType = "function"
			m[count] = e
			visHash[count] = hash
			count++
			s.Flags.functionIn = false
			arrDeep = append(arrDeep[:len(arrDeep)-1])
		} else if node.GetSymbol().GetTokenType() == 69 && !s.Flags.iterationStatement {
			e, ok := m[count-1]
			if !ok {
				log.Fatal("No match key in global hash!")
			}
			hash, ok := visHash[count-1]
			if !ok {
				log.Fatal("No match key in visHash !")
			}
			e.DataType = "function"
			hash.DataType = "function"
			m[count-1] = e
			visHash[count-1] = hash
		} else if s.Flags.typeSpecifier {
			e.DataType = node.GetText()
			hash.DataType = node.GetText()
			m[count] = e
			visHash[count] = hash
		}
	} else if s.Flags.local > 0 && !s.Flags.superclassName && !s.Flags.categoryName && s.Flags.classInterface && ne > 0  {
		if node.GetSymbol().GetTokenType() == 125 && s.Flags.typeSpecifier {
			e.DataType = node.GetText()
			hash.DataType = node.GetText()
			m[count] = e
			visHash[count] = hash
		} else if node.GetSymbol().GetTokenType() == 125 && !s.Flags.typeSpecifier && !s.Flags.function {
			e.Scope = gluing()
			e.Name = node.GetText()
			hash.Scope = gluing()
			hash.Name = node.GetText()
			m[count] = e
			visHash[count] = hash
			count++
		} else if node.GetSymbol().GetTokenType() == 69 && !s.Flags.iterationStatement {
			e, ok := m[count-1]
			if !ok {
				log.Fatal("No match key!")
			}
			hash, ok := visHash[count-1]
			if !ok {
				log.Fatal("No match key!")
			}
			e.DataType = "function"
			hash.DataType = "function"
			m[count-1] = e
			visHash[count-1] = hash
		} else if s.Flags.typeSpecifier {
			e.DataType = node.GetText()
			hash.DataType = node.GetText()
			m[count] = e
			visHash[count] = hash
		}
	}

	// ERROR NODE
	if debug && node.GetSymbol().GetTokenType() == 125 && !s.Flags.superclassName && !s.Flags.typeSpecifier && !s.Flags.initDeclaratorList &&
		node.GetSymbol().GetText() != "alloc" && node.GetSymbol().GetText() != "init" && node.GetSymbol().GetText() != "drain" &&
		node.GetSymbol().GetText() != "complexWithRe" && node.GetSymbol().GetText() != "NSMutableArray" {
		er := false
		for v, _ := range visHash {
			if node.GetText() == visHash[v].Name {
				er = true
				break
			}
		}
		if !er {
			log.Fatal("Fatal error compilation in: " + node.GetText())
		}
	}
}

// VisitErrorNode is called when an error node is visited.
func (s *BaseObjCListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseObjCListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseObjCListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTranslation_unit is called when production translation_unit is entered.
func (s *BaseObjCListener) EnterTranslation_unit(ctx *Translation_unitContext) {}

// ExitTranslation_unit is called when production translation_unit is exited.
func (s *BaseObjCListener) ExitTranslation_unit(ctx *Translation_unitContext) {}

// EnterExternal_declaration is called when production external_declaration is entered. вход
func (s *BaseObjCListener) EnterExternal_declaration(ctx *External_declarationContext) {
	if !start {
		s.Root = &opts.TreeData{Name: "Start"}
		start = true
	}
	s.current = s.Root
	s.nodes = append(s.nodes, s.Root)
}

// ExitExternal_declaration is called when production external_declaration is exited. и выход
func (s *BaseObjCListener) ExitExternal_declaration(ctx *External_declarationContext) {}

// EnterPreprocessor_declaration is called when production preprocessor_declaration is entered.
func (s *BaseObjCListener) EnterPreprocessor_declaration(ctx *Preprocessor_declarationContext) {
	node := opts.TreeData{Name: "#import"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitPreprocessor_declaration is called when production preprocessor_declaration is exited.
func (s *BaseObjCListener) ExitPreprocessor_declaration(ctx *Preprocessor_declarationContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterClass_interface is called when production class_interface is entered.
func (s *BaseObjCListener) EnterClass_interface(ctx *Class_interfaceContext) {
	start := opts.TreeData{Name: ctx.GetStart().GetText()}
	node := opts.TreeData{Name: "Class_interface"}
	end := opts.TreeData{Name: ctx.GetStop().GetText()}
	s.current.Children = append(s.current.Children, &start)
	s.current.Children = append(s.current.Children, &node)
	s.current.Children = append(s.current.Children, &end)
	s.current = &node
	s.nodes = append(s.nodes, &node)
	s.Flags.classInterface = true
}

// ExitClass_interface is called when production class_interface is exited.
func (s *BaseObjCListener) ExitClass_interface(ctx *Class_interfaceContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
	s.Flags.classInterface = false
	arrDeep = nil
}

// EnterCategory_interface is called when production category_interface is entered.
func (s *BaseObjCListener) EnterCategory_interface(ctx *Category_interfaceContext) {
	node := opts.TreeData{Name: ctx.GetStart().GetText()}
	//next := opts.TreeData{Name: "Category_interface"}
	s.current.Children = append(s.current.Children, &node)
	//s.current.Children = append(s.current.Children, &next)
	//s.current = &next
	//s.nodes = append(s.nodes, &next)
	s.Flags.classInterface = true
}

// ExitCategory_interface is called when production category_interface is exited.
func (s *BaseObjCListener) ExitCategory_interface(ctx *Category_interfaceContext) {
	//s.nodes = s.nodes[:len(s.nodes)-1]
	//s.current = s.nodes[len(s.nodes)-1]
	end := opts.TreeData{Name: "@end"}
	s.current.Children = append(s.current.Children, &end)
	s.Flags.classInterface = false
	arrDeep = nil
}

// EnterClass_implementation is called when production class_implementation is entered.
func (s *BaseObjCListener) EnterClass_implementation(ctx *Class_implementationContext) {
	start := opts.TreeData{Name: ctx.GetStart().GetText()}
	node := opts.TreeData{Name: "Class_implementation"}
	end := opts.TreeData{Name: ctx.GetStop().GetText()}
	s.current.Children = append(s.current.Children, &start)
	s.current.Children = append(s.current.Children, &node)
	s.current.Children = append(s.current.Children, &end)
	s.current = &node
	s.nodes = append(s.nodes, &node)
	s.Flags.classInterface = true
}

// ExitClass_implementation is called when production class_implementation is exited.
func (s *BaseObjCListener) ExitClass_implementation(ctx *Class_implementationContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
	s.Flags.classInterface = false
	arrDeep = nil
}

// EnterCategory_implementation is called when production category_implementation is entered.
func (s *BaseObjCListener) EnterCategory_implementation(ctx *Category_implementationContext) {
	node := opts.TreeData{Name: "@implementation"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
	s.Flags.classInterface = true
}

// ExitCategory_implementation is called when production category_implementation is exited.
func (s *BaseObjCListener) ExitCategory_implementation(ctx *Category_implementationContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
	s.Flags.classInterface = false
	//arrDeep = nil
}

// EnterProtocol_declaration is called when production protocol_declaration is entered.
func (s *BaseObjCListener) EnterProtocol_declaration(ctx *Protocol_declarationContext) {
	node := opts.TreeData{Name: "@protocol"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
	s.Flags.classInterface = true
}

// ExitProtocol_declaration is called when production protocol_declaration is exited.
func (s *BaseObjCListener) ExitProtocol_declaration(ctx *Protocol_declarationContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
	s.Flags.classInterface = false
}

// EnterProtocol_declaration_list is called when production protocol_declaration_list is entered.
func (s *BaseObjCListener) EnterProtocol_declaration_list(ctx *Protocol_declaration_listContext) {
	node := opts.TreeData{Name: "@protocol"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
	s.Flags.classInterface = true
}

// ExitProtocol_declaration_list is called when production protocol_declaration_list is exited.
func (s *BaseObjCListener) ExitProtocol_declaration_list(ctx *Protocol_declaration_listContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
	s.Flags.classInterface = false
}

// EnterClass_declaration_list is called when production class_declaration_list is entered.
func (s *BaseObjCListener) EnterClass_declaration_list(ctx *Class_declaration_listContext) {
	start := opts.TreeData{Name: ctx.GetStart().GetText()}
	node := opts.TreeData{Name: "Class_declaration"}
	end := opts.TreeData{Name: ctx.GetStop().GetText()}
	s.current.Children = append(s.current.Children, &start)
	s.current.Children = append(s.current.Children, &node)
	s.current.Children = append(s.current.Children, &end)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitClass_declaration_list is called when production class_declaration_list is exited.
func (s *BaseObjCListener) ExitClass_declaration_list(ctx *Class_declaration_listContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterClass_list is called when production class_list is entered.
func (s *BaseObjCListener) EnterClass_list(ctx *Class_listContext) {
	node := opts.TreeData{Name: "Class_list"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitClass_list is called when production class_list is exited.
func (s *BaseObjCListener) ExitClass_list(ctx *Class_listContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterProtocol_reference_list is called when production protocol_reference_list is entered.
func (s *BaseObjCListener) EnterProtocol_reference_list(ctx *Protocol_reference_listContext) {
	node := opts.TreeData{Name: "Protocol_reference_list"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitProtocol_reference_list is called when production protocol_reference_list is exited.
func (s *BaseObjCListener) ExitProtocol_reference_list(ctx *Protocol_reference_listContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterProtocol_list is called when production protocol_list is entered.
func (s *BaseObjCListener) EnterProtocol_list(ctx *Protocol_listContext) {
	node := opts.TreeData{Name: "Protocol_list"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitProtocol_list is called when production protocol_list is exited.
func (s *BaseObjCListener) ExitProtocol_list(ctx *Protocol_listContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterProperty_declaration is called when production property_declaration is entered.
func (s *BaseObjCListener) EnterProperty_declaration(ctx *Property_declarationContext) {
	prop := opts.TreeData{Name: ctx.GetStart().GetText()}
	//node := opts.TreeData{Name: "Property_declaration"}
	s.current.Children = append(s.current.Children, &prop)
	//s.current.Children = append(s.current.Children, &node)
	//s.current = &node
	//s.nodes = append(s.nodes, &node)
	s.Flags.property = true
}

// ExitProperty_declaration is called when production property_declaration is exited.
func (s *BaseObjCListener) ExitProperty_declaration(ctx *Property_declarationContext) {
	//s.nodes = s.nodes[:len(s.nodes)-1]
	//s.current = s.nodes[len(s.nodes)-1]
}

// EnterProperty_attributes_declaration is called when production property_attributes_declaration is entered.
func (s *BaseObjCListener) EnterProperty_attributes_declaration(ctx *Property_attributes_declarationContext) {
	//node := opts.TreeData{Name: "Property_attributes_declaration"}
	//s.current.Children = append(s.current.Children, &node)
	//s.current = &node
	//s.nodes = append(s.nodes, &node)
}

// ExitProperty_attributes_declaration is called when production property_attributes_declaration is exited.
func (s *BaseObjCListener) ExitProperty_attributes_declaration(ctx *Property_attributes_declarationContext) {
	//s.nodes = s.nodes[:len(s.nodes)-1]
	//s.current = s.nodes[len(s.nodes)-1]
}

// EnterProperty_attributes_list is called when production property_attributes_list is entered.
func (s *BaseObjCListener) EnterProperty_attributes_list(ctx *Property_attributes_listContext) {
	start := opts.TreeData{Name: "("}
	node := opts.TreeData{Name: "Property_attributes_list"}
	end := opts.TreeData{Name: ")"}
	s.current.Children = append(s.current.Children, &start)
	s.current.Children = append(s.current.Children, &node)
	s.current.Children = append(s.current.Children, &end)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitProperty_attributes_list is called when production property_attributes_list is exited.
func (s *BaseObjCListener) ExitProperty_attributes_list(ctx *Property_attributes_listContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterProperty_attribute is called when production property_attribute is entered.
func (s *BaseObjCListener) EnterProperty_attribute(ctx *Property_attributeContext) {
	node := opts.TreeData{Name: ctx.GetStart().GetText()}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitProperty_attribute is called when production property_attribute is exited.
func (s *BaseObjCListener) ExitProperty_attribute(ctx *Property_attributeContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterClass_name is called when production class_name is entered.
func (s *BaseObjCListener) EnterClass_name(ctx *Class_nameContext) {
	node := opts.TreeData{Name: ctx.GetStart().GetText()}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
	if ctx.GetStart().GetTokenType() == 125 && s.Flags.classInterface && !s.Flags.methodType && !s.Flags.typeSpecifier{
		arrDeep = append(arrDeep, Arr{ctx.GetText(), s.Flags.local})
	}
}

// ExitClass_name is called when production class_name is exited.
func (s *BaseObjCListener) ExitClass_name(ctx *Class_nameContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterSuperclass_name is called when production superclass_name is entered.
func (s *BaseObjCListener) EnterSuperclass_name(ctx *Superclass_nameContext) {
	if queue_superclass.Len() > 0 {
		n := opts.TreeData{Name: queue_superclass.Front().Value.(string)}
		s.current.Children = append(s.current.Children, &n)
		queue_superclass.Remove(queue_superclass.Front())
	}
	node := opts.TreeData{Name: "Superclass"}
	s.current.Children = append(s.current.Children, &node)
	n := opts.TreeData{Name: ctx.GetStart().GetText()}
	s.current = &node
	node.Children = append(s.current.Children, &n)
	s.nodes = append(s.nodes, &node)
	s.Flags.superclassName = true
}

// ExitSuperclass_name is called when production superclass_name is exited.
func (s *BaseObjCListener) ExitSuperclass_name(ctx *Superclass_nameContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
	s.Flags.superclassName = false
}

// EnterCategory_name is called when production category_name is entered.
func (s *BaseObjCListener) EnterCategory_name(ctx *Category_nameContext) {
	node := opts.TreeData{Name: "Category_name: " + ctx.GetStart().GetText()}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
	s.Flags.categoryName = true
}

// ExitCategory_name is called when production category_name is exited.
func (s *BaseObjCListener) ExitCategory_name(ctx *Category_nameContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
	s.Flags.categoryName = false
}

// EnterProtocol_name is called when production protocol_name is entered.
func (s *BaseObjCListener) EnterProtocol_name(ctx *Protocol_nameContext) {
	node := opts.TreeData{Name: "Protocol_name"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitProtocol_name is called when production protocol_name is exited.
func (s *BaseObjCListener) ExitProtocol_name(ctx *Protocol_nameContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterInstance_variables is called when production instance_variables is entered.
func (s *BaseObjCListener) EnterInstance_variables(ctx *Instance_variablesContext) {
	//node := opts.TreeData{Name: "Instance_variables"}
	//s.current.Children = append(s.current.Children, &node)
	//s.current = &node
	//s.nodes = append(s.nodes, &node)
	s.Flags.local++
	s.Flags.instanceVariables = true
}

// ExitInstance_variables is called when production instance_variables is exited.
func (s *BaseObjCListener) ExitInstance_variables(ctx *Instance_variablesContext) {
	//s.nodes = s.nodes[:len(s.nodes)-1]
	//s.current = s.nodes[len(s.nodes)-1]
	s.Flags.local--
	s.Flags.instanceVariables = false
	if ctx.GetStop().GetTokenType() == 72 && arrDeep[len(arrDeep)-1].Level == s.Flags.local && !s.Flags.classInterface {
		arrDeep = append(arrDeep[:len(arrDeep)-1])
		var last = 0
		for e := count; 0 <= e; e-- {
			_, ok := visHash[e]
			if ok {
				last = e
				break
			}
		}

		var str = visHash[last].Scope
		for i := last; visHash[i].Scope == str; i-- {
			delete(visHash, i)
		}
	}
}

// EnterVisibility_specification is called when production visibility_specification is entered.
func (s *BaseObjCListener) EnterVisibility_specification(ctx *Visibility_specificationContext) {
	node := opts.TreeData{Name: "Visibility_specification"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitVisibility_specification is called when production visibility_specification is exited.
func (s *BaseObjCListener) ExitVisibility_specification(ctx *Visibility_specificationContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterInterface_declaration_list is called when production interface_declaration_list is entered.
func (s *BaseObjCListener) EnterInterface_declaration_list(ctx *Interface_declaration_listContext) {
	node := opts.TreeData{Name: "Interface_declaration_list"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitInterface_declaration_list is called when production interface_declaration_list is exited.
func (s *BaseObjCListener) ExitInterface_declaration_list(ctx *Interface_declaration_listContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterClass_method_declaration is called when production class_method_declaration is entered.
func (s *BaseObjCListener) EnterClass_method_declaration(ctx *Class_method_declarationContext) {
	node := opts.TreeData{Name: "+"}
	//node := opts.TreeData{Name: "'+'"}
	s.current.Children = append(s.current.Children, &node)
	//s.current = &node
	//s.nodes = append(s.nodes, &node)
}

// ExitClass_method_declaration is called when production class_method_declaration is exited.
func (s *BaseObjCListener) ExitClass_method_declaration(ctx *Class_method_declarationContext) {
	//s.nodes = s.nodes[:len(s.nodes)-1]
	//s.current = s.nodes[len(s.nodes)-1]
}

// EnterInstance_method_declaration is called when production instance_method_declaration is entered.
func (s *BaseObjCListener) EnterInstance_method_declaration(ctx *Instance_method_declarationContext) {
	node := opts.TreeData{Name: "-"}
	s.current.Children = append(s.current.Children, &node)
	//s.current = &node
	//s.nodes = append(s.nodes, &node)
}

// ExitInstance_method_declaration is called when production instance_method_declaration is exited.
func (s *BaseObjCListener) ExitInstance_method_declaration(ctx *Instance_method_declarationContext) {
	//s.nodes = s.nodes[:len(s.nodes)-1]
	//s.current = s.nodes[len(s.nodes)-1]
}

// EnterMethod_declaration is called when production method_declaration is entered.
func (s *BaseObjCListener) EnterMethod_declaration(ctx *Method_declarationContext) {
	node := opts.TreeData{Name: "Method_declaration"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitMethod_declaration is called when production method_declaration is exited.
func (s *BaseObjCListener) ExitMethod_declaration(ctx *Method_declarationContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterImplementation_definition_list is called when production implementation_definition_list is entered.
func (s *BaseObjCListener) EnterImplementation_definition_list(ctx *Implementation_definition_listContext) {
	node := opts.TreeData{Name: "Implementation_definition"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitImplementation_definition_list is called when production implementation_definition_list is exited.
func (s *BaseObjCListener) ExitImplementation_definition_list(ctx *Implementation_definition_listContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterClass_method_definition is called when production class_method_definition is entered.
func (s *BaseObjCListener) EnterClass_method_definition(ctx *Class_method_definitionContext) {
	node := opts.TreeData{Name: "Class_method_definition"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitClass_method_definition is called when production class_method_definition is exited.
func (s *BaseObjCListener) ExitClass_method_definition(ctx *Class_method_definitionContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterInstance_method_definition is called when production instance_method_definition is entered.
func (s *BaseObjCListener) EnterInstance_method_definition(ctx *Instance_method_definitionContext) {
	//node := opts.TreeData{Name: "Instance_method_definition"}
	//s.current.Children = append(s.current.Children, &node)
	//s.current = &node
	//s.nodes = append(s.nodes, &node)
}

// ExitInstance_method_definition is called when production instance_method_definition is exited.
func (s *BaseObjCListener) ExitInstance_method_definition(ctx *Instance_method_definitionContext) {
	//s.nodes = s.nodes[:len(s.nodes)-1]
	//s.current = s.nodes[len(s.nodes)-1]
}

// EnterMethod_definition is called when production method_definition is entered.
func (s *BaseObjCListener) EnterMethod_definition(ctx *Method_definitionContext) {
	node := opts.TreeData{Name: "Method_definition"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitMethod_definition is called when production method_definition is exited.
func (s *BaseObjCListener) ExitMethod_definition(ctx *Method_definitionContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterMethod_selector is called when production method_selector is entered.
func (s *BaseObjCListener) EnterMethod_selector(ctx *Method_selectorContext) {
	node := opts.TreeData{Name: "Method_selector"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
	s.Flags.methodSelector = true
	if ctx.GetStart().GetTokenType() == 125 && s.Flags.classInterface {
		arrDeep = append(arrDeep, Arr{ctx.GetText(), s.Flags.local})
	}
}

// ExitMethod_selector is called when production method_selector is exited.
func (s *BaseObjCListener) ExitMethod_selector(ctx *Method_selectorContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
	s.Flags.methodSelector = false
}

// EnterKeyword_declarator is called when production keyword_declarator is entered.
func (s *BaseObjCListener) EnterKeyword_declarator(ctx *Keyword_declaratorContext) {
	node := opts.TreeData{Name: "Keyword_declarator"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitKeyword_declarator is called when production keyword_declarator is exited.
func (s *BaseObjCListener) ExitKeyword_declarator(ctx *Keyword_declaratorContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterSelector is called when production selector is entered.
func (s *BaseObjCListener) EnterSelector(ctx *SelectorContext) {
	node := opts.TreeData{Name: ctx.GetStart().GetText()}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitSelector is called when production selector is exited.
func (s *BaseObjCListener) ExitSelector(ctx *SelectorContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterMethod_type is called when production method_type is entered.
func (s *BaseObjCListener) EnterMethod_type(ctx *Method_typeContext) {
	start := opts.TreeData{Name: "("}
	node := opts.TreeData{Name: "Method_type"}
	end := opts.TreeData{Name: ")"}
	s.current.Children = append(s.current.Children, &start)
	s.current.Children = append(s.current.Children, &node)
	s.current.Children = append(s.current.Children, &end)
	s.current = &node
	s.nodes = append(s.nodes, &node)
	s.Flags.methodType = true
}

// ExitMethod_type is called when production method_type is exited.
func (s *BaseObjCListener) ExitMethod_type(ctx *Method_typeContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
	s.Flags.methodType = false
}

// EnterProperty_implementation is called when production property_implementation is entered.
func (s *BaseObjCListener) EnterProperty_implementation(ctx *Property_implementationContext) {
	node := opts.TreeData{Name: ctx.GetStart().GetText()}
	s.current.Children = append(s.current.Children, &node)
	//s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitProperty_implementation is called when production property_implementation is exited.
func (s *BaseObjCListener) ExitProperty_implementation(ctx *Property_implementationContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterProperty_synthesize_list is called when production property_synthesize_list is entered.
func (s *BaseObjCListener) EnterProperty_synthesize_list(ctx *Property_synthesize_listContext) {
	node := opts.TreeData{Name: "Property"}//????????
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitProperty_synthesize_list is called when production property_synthesize_list is exited.
func (s *BaseObjCListener) ExitProperty_synthesize_list(ctx *Property_synthesize_listContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterProperty_synthesize_item is called when production property_synthesize_item is entered.
func (s *BaseObjCListener) EnterProperty_synthesize_item(ctx *Property_synthesize_itemContext) {
	node := opts.TreeData{Name: ctx.GetStart().GetText()}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitProperty_synthesize_item is called when production property_synthesize_item is exited.
func (s *BaseObjCListener) ExitProperty_synthesize_item(ctx *Property_synthesize_itemContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterBlock_type is called when production block_type is entered.
func (s *BaseObjCListener) EnterBlock_type(ctx *Block_typeContext) {
	node := opts.TreeData{Name: "Block_type"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitBlock_type is called when production block_type is exited.
func (s *BaseObjCListener) ExitBlock_type(ctx *Block_typeContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterType_specifier is called when production type_specifier is entered.
func (s *BaseObjCListener) EnterType_specifier(ctx *Type_specifierContext) {
	node := opts.TreeData{Name: ctx.GetStart().GetText()}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
	s.Flags.typeSpecifier = true
}

// ExitType_specifier is called when production type_specifier is exited.
func (s *BaseObjCListener) ExitType_specifier(ctx *Type_specifierContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
	s.Flags.typeSpecifier = false
}

// EnterType_qualifier is called when production type_qualifier is entered.
func (s *BaseObjCListener) EnterType_qualifier(ctx *Type_qualifierContext) {
	node := opts.TreeData{Name: ctx.GetStart().GetText()}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitType_qualifier is called when production type_qualifier is exited.
func (s *BaseObjCListener) ExitType_qualifier(ctx *Type_qualifierContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterProtocol_qualifier is called when production protocol_qualifier is entered.
func (s *BaseObjCListener) EnterProtocol_qualifier(ctx *Protocol_qualifierContext) {
	node := opts.TreeData{Name: "Protocol_qualifier"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitProtocol_qualifier is called when production protocol_qualifier is exited.
func (s *BaseObjCListener) ExitProtocol_qualifier(ctx *Protocol_qualifierContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterPrimary_expression is called when production primary_expression is entered.
func (s *BaseObjCListener) EnterPrimary_expression(ctx *Primary_expressionContext) {
	q := opts.TreeData{Name: "Primary_expression"}
	s.current.Children = append(s.current.Children, &q)
	s.current = &q
	if ctx.GetStart().GetTokenType() == 125 {
		//identifier := opts.TreeData{Name: "Identifier"}
		//s.current.Children = append(s.current.Children, &identifier)
		//s.current = &identifier
		node := opts.TreeData{Name: ctx.GetStart().GetText()}
		s.current.Children = append(s.current.Children, &node)
		s.nodes = append(s.nodes, &node)
	} else if ctx.GetStart().GetTokenType() == 127 {
		//str := opts.TreeData{Name: "String"}
		//s.current.Children = append(s.current.Children, &str)
		//s.current = &str
		node := opts.TreeData{Name: ctx.GetStart().GetText()}
		s.current.Children = append(s.current.Children, &node)
		s.nodes = append(s.nodes, &node)
	} else {
		// TODO: нужно обработать
		node := opts.TreeData{Name: ctx.GetStart().GetText()}
		s.current.Children = append(s.current.Children, &node)
		s.nodes = append(s.nodes, &q)
	}
}

// ExitPrimary_expression is called when production primary_expression is exited.
func (s *BaseObjCListener) ExitPrimary_expression(ctx *Primary_expressionContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterDictionary_pair is called when production dictionary_pair is entered.
func (s *BaseObjCListener) EnterDictionary_pair(ctx *Dictionary_pairContext) {
	node := opts.TreeData{Name: "Dictionary_pair"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitDictionary_pair is called when production dictionary_pair is exited.
func (s *BaseObjCListener) ExitDictionary_pair(ctx *Dictionary_pairContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterDictionary_expression is called when production dictionary_expression is entered.
func (s *BaseObjCListener) EnterDictionary_expression(ctx *Dictionary_expressionContext) {
	start := opts.TreeData{Name: "{"}
	node := opts.TreeData{Name: "Dictionary"}
	end := opts.TreeData{Name: "}"}
	s.current.Children = append(s.current.Children, &start)
	s.current.Children = append(s.current.Children, &node)
	s.current.Children = append(s.current.Children, &end)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitDictionary_expression is called when production dictionary_expression is exited.
func (s *BaseObjCListener) ExitDictionary_expression(ctx *Dictionary_expressionContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterArray_expression is called when production array_expression is entered.
func (s *BaseObjCListener) EnterArray_expression(ctx *Array_expressionContext) {
	node := opts.TreeData{Name: "Array_expression"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitArray_expression is called when production array_expression is exited.
func (s *BaseObjCListener) ExitArray_expression(ctx *Array_expressionContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterBox_expression is called when production box_expression is entered.
func (s *BaseObjCListener) EnterBox_expression(ctx *Box_expressionContext) {
	node := opts.TreeData{Name: "Box_expressio"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitBox_expression is called when production box_expression is exited.
func (s *BaseObjCListener) ExitBox_expression(ctx *Box_expressionContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterBlock_parameters is called when production block_parameters is entered.
func (s *BaseObjCListener) EnterBlock_parameters(ctx *Block_parametersContext) {
	node := opts.TreeData{Name: "Block_parameters"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitBlock_parameters is called when production block_parameters is exited.
func (s *BaseObjCListener) ExitBlock_parameters(ctx *Block_parametersContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterBlock_expression is called when production block_expression is entered.
func (s *BaseObjCListener) EnterBlock_expression(ctx *Block_expressionContext) {
	node := opts.TreeData{Name: "Block_expression"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitBlock_expression is called when production block_expression is exited.
func (s *BaseObjCListener) ExitBlock_expression(ctx *Block_expressionContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterMessage_expression is called when production message_expression is entered.
func (s *BaseObjCListener) EnterMessage_expression(ctx *Message_expressionContext) {
	node := opts.TreeData{Name: ctx.GetStart().GetText()}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitMessage_expression is called when production message_expression is exited.
func (s *BaseObjCListener) ExitMessage_expression(ctx *Message_expressionContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterReceiver is called when production receiver is entered.
func (s *BaseObjCListener) EnterReceiver(ctx *ReceiverContext) {
	node := opts.TreeData{Name: "Receiver"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitReceiver is called when production receiver is exited.
func (s *BaseObjCListener) ExitReceiver(ctx *ReceiverContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterMessage_selector is called when production message_selector is entered.
func (s *BaseObjCListener) EnterMessage_selector(ctx *Message_selectorContext) {
	node := opts.TreeData{Name: "Message_selector"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitMessage_selector is called when production message_selector is exited.
func (s *BaseObjCListener) ExitMessage_selector(ctx *Message_selectorContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterKeyword_argument is called when production keyword_argument is entered.
func (s *BaseObjCListener) EnterKeyword_argument(ctx *Keyword_argumentContext) {
	node := opts.TreeData{Name: "Keyword_argument"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitKeyword_argument is called when production keyword_argument is exited.
func (s *BaseObjCListener) ExitKeyword_argument(ctx *Keyword_argumentContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterSelector_expression is called when production selector_expression is entered.
func (s *BaseObjCListener) EnterSelector_expression(ctx *Selector_expressionContext) {
	node := opts.TreeData{Name: "Selector_expression"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitSelector_expression is called when production selector_expression is exited.
func (s *BaseObjCListener) ExitSelector_expression(ctx *Selector_expressionContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterSelector_name is called when production selector_name is entered.
func (s *BaseObjCListener) EnterSelector_name(ctx *Selector_nameContext) {
	node := opts.TreeData{Name: "Selector_name"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitSelector_name is called when production selector_name is exited.
func (s *BaseObjCListener) ExitSelector_name(ctx *Selector_nameContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterProtocol_expression is called when production protocol_expression is entered.
func (s *BaseObjCListener) EnterProtocol_expression(ctx *Protocol_expressionContext) {
	node := opts.TreeData{Name: "Protocol_expression"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitProtocol_expression is called when production protocol_expression is exited.
func (s *BaseObjCListener) ExitProtocol_expression(ctx *Protocol_expressionContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterEncode_expression is called when production encode_expression is entered.
func (s *BaseObjCListener) EnterEncode_expression(ctx *Encode_expressionContext) {
	//node := opts.TreeData{Name: "Encode_expression"}
	//s.current.Children = append(s.current.Children, &node)
	//s.current = &node
	//s.nodes = append(s.nodes, &node)
}

// ExitEncode_expression is called when production encode_expression is exited.
func (s *BaseObjCListener) ExitEncode_expression(ctx *Encode_expressionContext) {
	//s.nodes = s.nodes[:len(s.nodes)-1]
	//s.current = s.nodes[len(s.nodes)-1]
}

// EnterType_variable_declarator is called when production type_variable_declarator is entered.
func (s *BaseObjCListener) EnterType_variable_declarator(ctx *Type_variable_declaratorContext) {
	node := opts.TreeData{Name: "Type_variable_declarator"}
	in := opts.TreeData{Name: "in"}
	s.current.Children = append(s.current.Children, &node)
	s.current.Children = append(s.current.Children, &in)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitType_variable_declarator is called when production type_variable_declarator is exited.
func (s *BaseObjCListener) ExitType_variable_declarator(ctx *Type_variable_declaratorContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterTry_statement is called when production try_statement is entered.
func (s *BaseObjCListener) EnterTry_statement(ctx *Try_statementContext) {
	node := opts.TreeData{Name: "Try_statement"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitTry_statement is called when production try_statement is exited.
func (s *BaseObjCListener) ExitTry_statement(ctx *Try_statementContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterCatch_statement is called when production catch_statement is entered.
func (s *BaseObjCListener) EnterCatch_statement(ctx *Catch_statementContext) {
	node := opts.TreeData{Name: "Catch_statement"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitCatch_statement is called when production catch_statement is exited.
func (s *BaseObjCListener) ExitCatch_statement(ctx *Catch_statementContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterFinally_statement is called when production finally_statement is entered.
func (s *BaseObjCListener) EnterFinally_statement(ctx *Finally_statementContext) {
	node := opts.TreeData{Name: "Finally_statement"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitFinally_statement is called when production finally_statement is exited.
func (s *BaseObjCListener) ExitFinally_statement(ctx *Finally_statementContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterThrow_statement is called when production throw_statement is entered.
func (s *BaseObjCListener) EnterThrow_statement(ctx *Throw_statementContext) {
	node := opts.TreeData{Name: "Throw_statemen"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitThrow_statement is called when production throw_statement is exited.
func (s *BaseObjCListener) ExitThrow_statement(ctx *Throw_statementContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterTry_block is called when production try_block is entered.
func (s *BaseObjCListener) EnterTry_block(ctx *Try_blockContext) {
	node := opts.TreeData{Name: "Try_block"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitTry_block is called when production try_block is exited.
func (s *BaseObjCListener) ExitTry_block(ctx *Try_blockContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterSynchronized_statement is called when production synchronized_statement is entered.
func (s *BaseObjCListener) EnterSynchronized_statement(ctx *Synchronized_statementContext) {
	node := opts.TreeData{Name: ctx.GetStart().GetText()}
	n := opts.TreeData{Name: "Body_Sync"}
	s.current.Children = append(s.current.Children, &node)
	s.current.Children = append(s.current.Children, &n)
	s.current = &n
	s.nodes = append(s.nodes, &n)
}

// ExitSynchronized_statement is called when production synchronized_statement is exited.
func (s *BaseObjCListener) ExitSynchronized_statement(ctx *Synchronized_statementContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterAutorelease_statement is called when production autorelease_statement is entered.
func (s *BaseObjCListener) EnterAutorelease_statement(ctx *Autorelease_statementContext) {
	node := opts.TreeData{Name: "Autorelease_statement"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitAutorelease_statement is called when production autorelease_statement is exited.
func (s *BaseObjCListener) ExitAutorelease_statement(ctx *Autorelease_statementContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterFunction_definition is called when production function_definition is entered.
func (s *BaseObjCListener) EnterFunction_definition(ctx *Function_definitionContext) {
	node := opts.TreeData{Name: "Function_definition"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitFunction_definition is called when production function_definition is exited.
func (s *BaseObjCListener) ExitFunction_definition(ctx *Function_definitionContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
	arrDeep = nil
}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseObjCListener) EnterDeclaration(ctx *DeclarationContext) {
	//node := opts.TreeData{Name: "Declaration"}
	//s.current.Children = append(s.current.Children, &node)
	//s.current = &node
	//s.nodes = append(s.nodes, &node)
}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseObjCListener) ExitDeclaration(ctx *DeclarationContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterDeclaration_specifiers is called when production declaration_specifiers is entered.
func (s *BaseObjCListener) EnterDeclaration_specifiers(ctx *Declaration_specifiersContext) {
	//node := opts.TreeData{Name: "Declaration_specifier"}
	//s.current.Children = append(s.current.Children, &node)
	//s.current = &node
	//s.nodes = append(s.nodes, &node)
}

// ExitDeclaration_specifiers is called when production declaration_specifiers is exited.
func (s *BaseObjCListener) ExitDeclaration_specifiers(ctx *Declaration_specifiersContext) {
	//s.nodes = s.nodes[:len(s.nodes)-1]
	//s.current = s.nodes[len(s.nodes)-1]
}

// EnterArc_behaviour_specifier is called when production arc_behaviour_specifier is entered.
func (s *BaseObjCListener) EnterArc_behaviour_specifier(ctx *Arc_behaviour_specifierContext) {
	node := opts.TreeData{Name: "Arc_behaviour_specifier"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitArc_behaviour_specifier is called when production arc_behaviour_specifier is exited.
func (s *BaseObjCListener) ExitArc_behaviour_specifier(ctx *Arc_behaviour_specifierContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterStorage_class_specifier is called when production storage_class_specifier is entered.
func (s *BaseObjCListener) EnterStorage_class_specifier(ctx *Storage_class_specifierContext) {
	node := opts.TreeData{Name: "Storage_class_specifier"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitStorage_class_specifier is called when production storage_class_specifier is exited.
func (s *BaseObjCListener) ExitStorage_class_specifier(ctx *Storage_class_specifierContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterInit_declarator_list is called when production init_declarator_list is entered.
func (s *BaseObjCListener) EnterInit_declarator_list(ctx *Init_declarator_listContext) {
	node := opts.TreeData{Name: "Init_declarator_list"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
	s.Flags.initDeclaratorList = true
}

// ExitInit_declarator_list is called when production init_declarator_list is exited.
func (s *BaseObjCListener) ExitInit_declarator_list(ctx *Init_declarator_listContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
	s.Flags.initDeclaratorList = false
}

// EnterInit_declarator is called when production init_declarator is entered.
func (s *BaseObjCListener) EnterInit_declarator(ctx *Init_declaratorContext) {
	//node := opts.TreeData{Name: "Init_declarator"}
	//sep := opts.TreeData{Name: ";"}
	//s.current.Children = append(s.current.Children, &node)
	//s.current.Children = append(s.current.Children, &sep)
	//s.nodes = append(s.nodes, &node)
}

// ExitInit_declarator is called when production init_declarator is exited.
func (s *BaseObjCListener) ExitInit_declarator(ctx *Init_declaratorContext) {
	//s.nodes = s.nodes[:len(s.nodes)-1]
	//s.current = s.nodes[len(s.nodes)-1]
}

// EnterStruct_or_union_specifier is called when production struct_or_union_specifier is entered.
func (s *BaseObjCListener) EnterStruct_or_union_specifier(ctx *Struct_or_union_specifierContext) {
	node := opts.TreeData{Name: "Struct_or_union_specifier"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitStruct_or_union_specifier is called when production struct_or_union_specifier is exited.
func (s *BaseObjCListener) ExitStruct_or_union_specifier(ctx *Struct_or_union_specifierContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterStruct_declaration is called when production struct_declaration is entered.
func (s *BaseObjCListener) EnterStruct_declaration(ctx *Struct_declarationContext) {
	//node := opts.TreeData{Name: "Struct_declaration"}
	//s.current.Children = append(s.current.Children, &node)
	//s.current = &node
	//s.nodes = append(s.nodes, &node)
}

// ExitStruct_declaration is called when production struct_declaration is exited.
func (s *BaseObjCListener) ExitStruct_declaration(ctx *Struct_declarationContext) {
	//s.nodes = s.nodes[:len(s.nodes)-1]
	//s.current = s.nodes[len(s.nodes)-1]
}

// EnterSpecifier_qualifier_list is called when production specifier_qualifier_list is entered.
func (s *BaseObjCListener) EnterSpecifier_qualifier_list(ctx *Specifier_qualifier_listContext) {
	//node := opts.TreeData{Name: "Specifier_qualifier_list"}
	//s.current.Children = append(s.current.Children, &node)
	//s.current = &node
	//s.nodes = append(s.nodes, &node)
	s.Flags.specifierQualifierList = true
}

// ExitSpecifier_qualifier_list is called when production specifier_qualifier_list is exited.
func (s *BaseObjCListener) ExitSpecifier_qualifier_list(ctx *Specifier_qualifier_listContext) {
	//s.nodes = s.nodes[:len(s.nodes)-1]
	//s.current = s.nodes[len(s.nodes)-1]
	s.Flags.specifierQualifierList = false
}

// EnterStruct_declarator_list is called when production struct_declarator_list is entered.
func (s *BaseObjCListener) EnterStruct_declarator_list(ctx *Struct_declarator_listContext) {
	//node := opts.TreeData{Name: ";"}
	//s.current.Children = append(s.current.Children, &node)
	////s.current = &node
	//s.nodes = append(s.nodes, &node)
}

// ExitStruct_declarator_list is called when production struct_declarator_list is exited.
func (s *BaseObjCListener) ExitStruct_declarator_list(ctx *Struct_declarator_listContext) {
	//node := opts.TreeData{Name: ";"}
	//s.current.Children = append(s.current.Children, &node)
	//s.nodes = append(s.nodes, &node)
	//s.nodes = s.nodes[:len(s.nodes)-1]
	//s.current = s.nodes[len(s.nodes)-1]
}

// EnterStruct_declarator is called when production struct_declarator is entered.
func (s *BaseObjCListener) EnterStruct_declarator(ctx *Struct_declaratorContext) {
	node := opts.TreeData{Name: "Struct_declarator"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitStruct_declarator is called when production struct_declarator is exited.
func (s *BaseObjCListener) ExitStruct_declarator(ctx *Struct_declaratorContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterEnum_specifier is called when production enum_specifier is entered.
func (s *BaseObjCListener) EnterEnum_specifier(ctx *Enum_specifierContext) {
	node := opts.TreeData{Name: "Enum_specifie"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitEnum_specifier is called when production enum_specifier is exited.
func (s *BaseObjCListener) ExitEnum_specifier(ctx *Enum_specifierContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterEnumerator_list is called when production enumerator_list is entered.
func (s *BaseObjCListener) EnterEnumerator_list(ctx *Enumerator_listContext) {
	node := opts.TreeData{Name: "Enumerator_list"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitEnumerator_list is called when production enumerator_list is exited.
func (s *BaseObjCListener) ExitEnumerator_list(ctx *Enumerator_listContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterEnumerator is called when production enumerator is entered.
func (s *BaseObjCListener) EnterEnumerator(ctx *EnumeratorContext) {
	node := opts.TreeData{Name: "Enumerator"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitEnumerator is called when production enumerator is exited.
func (s *BaseObjCListener) ExitEnumerator(ctx *EnumeratorContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterPointer is called when production pointer is entered.
func (s *BaseObjCListener) EnterPointer(ctx *PointerContext) {
	node := opts.TreeData{Name: ctx.GetStart().GetText()}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
	s.Flags.pointer = true
}

// ExitPointer is called when production pointer is exited.
func (s *BaseObjCListener) ExitPointer(ctx *PointerContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
	s.Flags.pointer = false
}

// EnterDeclarator is called when production declarator is entered.
func (s *BaseObjCListener) EnterDeclarator(ctx *DeclaratorContext) {
	//node := opts.TreeData{Name: "Declarator"}
	//s.current.Children = append(s.current.Children, &node)
	//s.current = &node
	//s.nodes = append(s.nodes, &node)

	// for function
	if ctx.GetStop().GetTokenType() == 70 && !(ctx.GetStart().GetTokenType() == 69) {
		arrDeep = append(arrDeep, Arr{ctx.GetStart().GetText(), s.Flags.local})
	} else if ctx.GetStop().GetTokenType() == 70 && ctx.GetStart().GetTokenType() == 69 {
		s.Flags.function = true
	}
}

// ExitDeclarator is called when production declarator is exited.
func (s *BaseObjCListener) ExitDeclarator(ctx *DeclaratorContext) {
	//s.nodes = s.nodes[:len(s.nodes)-1]
	//s.current = s.nodes[len(s.nodes)-1]
	if s.Flags.function {
		s.Flags.function = false
	}
}

// EnterDirect_declarator is called when production direct_declarator is entered.
func (s *BaseObjCListener) EnterDirect_declarator(ctx *Direct_declaratorContext) {
	node := opts.TreeData{Name: ctx.GetStart().GetText()}
	s.current.Children = append(s.current.Children, &node)
	////s.current = &node
	//s.nodes = append(s.nodes, &node)
}

// ExitDirect_declarator is called when production direct_declarator is exited.
func (s *BaseObjCListener) ExitDirect_declarator(ctx *Direct_declaratorContext) {
	//s.nodes = s.nodes[:len(s.nodes)-1]
	//s.current = s.nodes[len(s.nodes)-1]
}

// EnterDeclarator_suffix is called when production declarator_suffix is entered.
func (s *BaseObjCListener) EnterDeclarator_suffix(ctx *Declarator_suffixContext) {
	if ctx.GetStart().GetTokenType() == 69 {
		start := opts.TreeData{Name: "("}
		node := opts.TreeData{Name: "Declarator_suffix"}
		end := opts.TreeData{Name: ")"}
		s.current.Children = append(s.current.Children, &start)
		s.current.Children = append(s.current.Children, &node)
		s.current.Children = append(s.current.Children, &end)
		s.current = &node
		s.nodes = append(s.nodes, &node)
	} else if ctx.GetStart().GetTokenType() == 73 {
		start := opts.TreeData{Name: "["}
		node := opts.TreeData{Name: "Declarator_suffix"}
		end := opts.TreeData{Name: "]"}
		s.current.Children = append(s.current.Children, &start)
		s.current.Children = append(s.current.Children, &node)
		s.current.Children = append(s.current.Children, &end)
		s.current = &node
		s.nodes = append(s.nodes, &node)
	}
}

// ExitDeclarator_suffix is called when production declarator_suffix is exited.
func (s *BaseObjCListener) ExitDeclarator_suffix(ctx *Declarator_suffixContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterParameter_list is called when production parameter_list is entered.
func (s *BaseObjCListener) EnterParameter_list(ctx *Parameter_listContext) {
	//node := opts.TreeData{Name: "Parameter_list"}
	//s.current.Children = append(s.current.Children, &node)
	//s.current = &node
	//s.nodes = append(s.nodes, &node)
}

// ExitParameter_list is called when production parameter_list is exited.
func (s *BaseObjCListener) ExitParameter_list(ctx *Parameter_listContext) {
	//s.nodes = s.nodes[:len(s.nodes)-1]
	//s.current = s.nodes[len(s.nodes)-1]
}

// EnterParameter_declaration is called when production parameter_declaration is entered.
func (s *BaseObjCListener) EnterParameter_declaration(ctx *Parameter_declarationContext) {
	node := opts.TreeData{Name: "Parameter_declaration"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitParameter_declaration is called when production parameter_declaration is exited.
func (s *BaseObjCListener) ExitParameter_declaration(ctx *Parameter_declarationContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterInitializer is called when production initializer is entered.
func (s *BaseObjCListener) EnterInitializer(ctx *InitializerContext) {
	if queue_initDeclarator.Len() > 0 {
		n := opts.TreeData{Name: queue_initDeclarator.Front().Value.(string)}
		s.current.Children = append(s.current.Children, &n)
		s.nodes = append(s.nodes, &n)
		queue_initDeclarator.Remove(queue_initDeclarator.Front())
	}
	//node := opts.TreeData{Name: "Initializer"}
	//s.current.Children = append(s.current.Children, &node)
	//s.current = &node
	//s.nodes = append(s.nodes, &node)
}

// ExitInitializer is called when production initializer is exited.
func (s *BaseObjCListener) ExitInitializer(ctx *InitializerContext) {
	// TODO: осторожно!
	//s.nodes = s.nodes[:len(s.nodes)-1]
	//s.current = s.nodes[len(s.nodes)-1]
}

// EnterType_name is called when production type_name is entered.
func (s *BaseObjCListener) EnterType_name(ctx *Type_nameContext) {
	//node := opts.TreeData{Name: "Type_name"}
	//s.current.Children = append(s.current.Children, &node)
	//s.current = &node
	//s.nodes = append(s.nodes, &node)
}

// ExitType_name is called when production type_name is exited.
func (s *BaseObjCListener) ExitType_name(ctx *Type_nameContext) {
	//s.nodes = s.nodes[:len(s.nodes)-1]
	//s.current = s.nodes[len(s.nodes)-1]
}

// EnterAbstract_declarator is called when production abstract_declarator is entered.
func (s *BaseObjCListener) EnterAbstract_declarator(ctx *Abstract_declaratorContext) {
	//node := opts.TreeData{Name: "Abstract_declarator"}
	//s.current.Children = append(s.current.Children, &node)
	//s.current = &node
	//s.nodes = append(s.nodes, &node)
}

// ExitAbstract_declarator is called when production abstract_declarator is exited.
func (s *BaseObjCListener) ExitAbstract_declarator(ctx *Abstract_declaratorContext) {
	//s.nodes = s.nodes[:len(s.nodes)-1]
	//s.current = s.nodes[len(s.nodes)-1]
}

// EnterAbstract_declarator_suffix is called when production abstract_declarator_suffix is entered.
func (s *BaseObjCListener) EnterAbstract_declarator_suffix(ctx *Abstract_declarator_suffixContext) {
	node := opts.TreeData{Name: "Abstract_declarator_suffi"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitAbstract_declarator_suffix is called when production abstract_declarator_suffix is exited.
func (s *BaseObjCListener) ExitAbstract_declarator_suffix(ctx *Abstract_declarator_suffixContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterParameter_declaration_list is called when production parameter_declaration_list is entered.
func (s *BaseObjCListener) EnterParameter_declaration_list(ctx *Parameter_declaration_listContext) {
	//node := opts.TreeData{Name: "Parameter_declaration_list"}
	//s.current.Children = append(s.current.Children, &node)
	//s.current = &node
	//s.nodes = append(s.nodes, &node)
	s.Flags.declaratorSuffix = true
}

// ExitParameter_declaration_list is called when production parameter_declaration_list is exited.
func (s *BaseObjCListener) ExitParameter_declaration_list(ctx *Parameter_declaration_listContext) {
	//s.nodes = s.nodes[:len(s.nodes)-1]
	//s.current = s.nodes[len(s.nodes)-1]
	s.Flags.declaratorSuffix = false
}

// EnterStatement_list is called when production statement_list is entered.
func (s *BaseObjCListener) EnterStatement_list(ctx *Statement_listContext) {
	//node := opts.TreeData{Name: "Statement_list"}
	//s.current.Children = append(s.current.Children, &node)
	//s.current = &node
	//s.nodes = append(s.nodes, &node)
}

// ExitStatement_list is called when production statement_list is exited.
func (s *BaseObjCListener) ExitStatement_list(ctx *Statement_listContext) {
	//s.nodes = s.nodes[:len(s.nodes)-1]
	//s.current = s.nodes[len(s.nodes)-1]
}

// EnterStatement is called when production statement is entered.
func (s *BaseObjCListener) EnterStatement(ctx *StatementContext) {
	if closecopee {
		node := opts.TreeData{Name: ")"}
		s.current.Children = append(s.current.Children, &node)
		closecopee = false
	}
}

// ExitStatement is called when production statement is exited.
func (s *BaseObjCListener) ExitStatement(ctx *StatementContext) {
	//s.nodes = s.nodes[:len(s.nodes)-1]
	//s.current = s.nodes[len(s.nodes)-1]

	if ctx.GetStart().GetTokenType() == 52 || ctx.GetStart().GetTokenType() == 46 ||
		ctx.GetStart().GetTokenType() == 33 || ctx.GetStart().GetTokenType() == 27 {
		node := opts.TreeData{Name: ctx.GetStop().GetText()}
		s.current.Children = append(s.current.Children, &node)
		//s.current = &node
		s.nodes = append(s.nodes, &node)
	}
}

// EnterLabeled_statement is called when production labeled_statement is entered.
func (s *BaseObjCListener) EnterLabeled_statement(ctx *Labeled_statementContext) {
	node := opts.TreeData{Name: "Labeled_statement"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
	if ctx.GetStart().GetTokenType() == 30 {
		s.Flags.labeledStatement = ctx.GetStart().GetTokenType()
	} else if ctx.GetStart().GetTokenType() == 34 {
		arrDeep = append(arrDeep, Arr{"default", s.Flags.local})
	}
}

// ExitLabeled_statement is called when production labeled_statement is exited.
func (s *BaseObjCListener) ExitLabeled_statement(ctx *Labeled_statementContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
	if ctx.GetStart().GetTokenType() == 30 {
		s.Flags.labeledStatement = ctx.GetStart().GetTokenType()
	} else if ctx.GetStart().GetTokenType() == 34 {
		arrDeep = append(arrDeep[:len(arrDeep)-1])
		var last = 0
		for e := count; 0 <= e; e-- {
			_, ok := visHash[e]
			if ok {
				last = e
				break
			}
		}

		var str = visHash[last].Scope
		for i := last; visHash[i].Scope == str; i-- {
			delete(visHash, i)
		}
		s.Flags.labeledStatement = 0
	}
}

// EnterCompound_statement is called when production compound_statement is entered.
func (s *BaseObjCListener) EnterCompound_statement(ctx *Compound_statementContext) {
	if ifelse {
		strelse := opts.TreeData{Name: "else"}
		s.current.Children = append(s.current.Children, &strelse)
		ifelse = false
	}
	start := opts.TreeData{Name: "{"}
	node := opts.TreeData{Name: "Compound_statement"}
	end := opts.TreeData{Name: "}"}
	s.current.Children = append(s.current.Children, &start)
	s.current.Children = append(s.current.Children, &node)
	s.current.Children = append(s.current.Children, &end)
	if s.Flags.doWhile {
		while := opts.TreeData{Name: "while"}
		s.current.Children = append(s.current.Children, &while)
	}
	s.current = &node
	s.nodes = append(s.nodes, &node)
	s.Flags.local++
}

// ExitCompound_statement is called when production compound_statement is exited.
func (s *BaseObjCListener) ExitCompound_statement(ctx *Compound_statementContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
	s.Flags.local--
	// DROP LEVEL
	// TODO: править
	if ctx.GetStop().GetTokenType() == 72 && arrDeep[len(arrDeep)-1].Level == s.Flags.local{
		arrDeep = append(arrDeep[:len(arrDeep)-1])
		var last = 0
		for e := count; 0 <= e; e-- {
			_, ok := visHash[e]
			if ok {
				last = e
				break
			}
		}

		var str = visHash[last].Scope
		for i := last; visHash[i].Scope == str; i-- {
			delete(visHash, i)
		}
	}
}

// EnterSelection_statement is called when production selection_statement is entered.
func (s *BaseObjCListener) EnterSelection_statement(ctx *Selection_statementContext) {
	start := opts.TreeData{Name: ctx.GetStart().GetText()}
	node := opts.TreeData{Name: "Selection_statement"}
	if ctx.GetStart().GetTokenType() == 43 {
		node = opts.TreeData{Name: "Selection_statement"}
	} else if ctx.GetStart().GetTokenType() == 58 {
		node = opts.TreeData{Name: "Selection_statement"}
	}
	s.current.Children = append(s.current.Children, &start)
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
	s.Flags.local++
}

// ExitSelection_statement is called when production selection_statement is exited.
func (s *BaseObjCListener) ExitSelection_statement(ctx *Selection_statementContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
	s.Flags.local--
}

// EnterFor_in_statement is called when production for_in_statement is entered.
func (s *BaseObjCListener) EnterFor_in_statement(ctx *For_in_statementContext) {
	start := opts.TreeData{Name: ctx.GetStart().GetText()}
	open := opts.TreeData{Name: "("}
	node := opts.TreeData{Name: "For_in_statement"}
	//close := opts.TreeData{Name: ")"}
	s.current.Children = append(s.current.Children, &start)
	s.current.Children = append(s.current.Children, &open)
	s.current.Children = append(s.current.Children, &node)
	//s.current.Children = append(s.current.Children, &close)
	s.current = &node
	s.nodes = append(s.nodes, &node)
	if ctx.GetStart().GetTokenType() == 41 {
		arrDeep = append(arrDeep, Arr{ctx.GetStart().GetText(), s.Flags.local})
	}
	s.Flags.For = true
}

// ExitFor_in_statement is called when production for_in_statement is exited.
func (s *BaseObjCListener) ExitFor_in_statement(ctx *For_in_statementContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterFor_statement is called when production for_statement is entered.
func (s *BaseObjCListener) EnterFor_statement(ctx *For_statementContext) {
	start := opts.TreeData{Name: ctx.GetStart().GetText()}
	open := opts.TreeData{Name: "("}
	node := opts.TreeData{Name: "For_statement"}
	//close := opts.TreeData{Name: ")"}
	s.current.Children = append(s.current.Children, &start)
	s.current.Children = append(s.current.Children, &open)
	s.current.Children = append(s.current.Children, &node)
	//s.current.Children = append(s.current.Children, &close)
	s.current = &node
	s.nodes = append(s.nodes, &node)
	if ctx.GetStart().GetTokenType() == 41 {
		arrDeep = append(arrDeep, Arr{ctx.GetStart().GetText(), s.Flags.local})
	}
	s.Flags.For = true
}

// ExitFor_statement is called when production for_statement is exited.
func (s *BaseObjCListener) ExitFor_statement(ctx *For_statementContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
	s.Flags.For = false
}

// EnterWhile_statement is called when production while_statement is entered.
func (s *BaseObjCListener) EnterWhile_statement(ctx *While_statementContext) {
	start := opts.TreeData{Name: ctx.GetStart().GetText()}
	node := opts.TreeData{Name: "While_statement"}
	s.current.Children = append(s.current.Children, &start)
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitWhile_statement is called when production while_statement is exited.
func (s *BaseObjCListener) ExitWhile_statement(ctx *While_statementContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterDo_statement is called when production do_statement is entered.
func (s *BaseObjCListener) EnterDo_statement(ctx *Do_statementContext) {
	start := opts.TreeData{Name: ctx.GetStart().GetText()}
	node := opts.TreeData{Name: "Do_statement"}
	end := opts.TreeData{Name: ctx.GetStop().GetText()}
	s.current.Children = append(s.current.Children, &start)
	s.current.Children = append(s.current.Children, &node)
	s.current.Children = append(s.current.Children, &end)
	s.current = &node
	s.nodes = append(s.nodes, &node)
	s.Flags.doWhile = true
}

// ExitDo_statement is called when production do_statement is exited.
func (s *BaseObjCListener) ExitDo_statement(ctx *Do_statementContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
	s.Flags.doWhile = false
}

// EnterIteration_statement is called when production iteration_statement is entered.
func (s *BaseObjCListener) EnterIteration_statement(ctx *Iteration_statementContext) {
	//node := opts.TreeData{Name: "Iteration_statement"}
	//s.current.Children = append(s.current.Children, &node)
	//s.current = &node
	//s.nodes = append(s.nodes, &node)
	s.Flags.iterationStatement = true
}

// ExitIteration_statement is called when production iteration_statement is exited.
func (s *BaseObjCListener) ExitIteration_statement(ctx *Iteration_statementContext) {
	//s.nodes = s.nodes[:len(s.nodes)-1]
	//s.current = s.nodes[len(s.nodes)-1]
	s.Flags.iterationStatement = false
}

// EnterJump_statement is called when production jump_statement is entered.
func (s *BaseObjCListener) EnterJump_statement(ctx *Jump_statementContext) {
	node := opts.TreeData{Name: ctx.GetStart().GetText()}
	s.current.Children = append(s.current.Children, &node)
	//s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitJump_statement is called when production jump_statement is exited.
func (s *BaseObjCListener) ExitJump_statement(ctx *Jump_statementContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterExpression is called when production expression is entered.
func (s *BaseObjCListener) EnterExpression(ctx *ExpressionContext) {
// TODO: закомментировать добавить ";"
	node := opts.TreeData{Name: "Expression"}
	s.current.Children = append(s.current.Children, &node)
	q := opts.TreeData{Name: ";"}
	s.current.Children = append(s.current.Children, &q)
	s.current = &node
	s.nodes = append(s.nodes, &node)
	s.Flags.expression = true
}

// ExitExpression is called when production expression is exited.
func (s *BaseObjCListener) ExitExpression(ctx *ExpressionContext) {
	// TODO: перенести expression
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
	s.Flags.expression = false
}

// EnterAssignment_expression is called when production assignment_expression is entered.
func (s *BaseObjCListener) EnterAssignment_expression(ctx *Assignment_expressionContext) {
	node := opts.TreeData{Name: "Assignment_exp"}
	if openscope && !s.Flags.For {
		scopeopen := opts.TreeData{Name: "("}
		s.current.Children = append(s.current.Children, &scopeopen)
		s.current.Children = append(s.current.Children, &node)
		scopeclose := opts.TreeData{Name: ")"}
		s.current.Children = append(s.current.Children, &scopeclose)
		//s.current = &node
	} else if openquotes && !s.Flags.For {
		quotesopen := opts.TreeData{Name: "["}
		s.current.Children = append(s.current.Children, &quotesopen)
		s.current.Children = append(s.current.Children, &node)
		closequotes := opts.TreeData{Name: "]"}
		s.current.Children = append(s.current.Children, &closequotes)
		//s.current = &node
	} else {
		s.current.Children = append(s.current.Children, &node)
		s.current = &node
	}
	s.Flags.expression = true
	s.nodes = append(s.nodes, &node)
}

// ExitAssignment_expression is called when production assignment_expression is exited.
func (s *BaseObjCListener) ExitAssignment_expression(ctx *Assignment_expressionContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterAssignment_operator is called when production assignment_operator is entered.
func (s *BaseObjCListener) EnterAssignment_operator(ctx *Assignment_operatorContext) {
	q := opts.TreeData{Name: "Assignment_operator"}
	s.current.Children = append(s.current.Children, &q)
	s.current = &q
	node := opts.TreeData{Name: ctx.GetText()}
	s.current.Children = append(s.current.Children, &node)
	//s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitAssignment_operator is called when production assignment_operator is exited.
func (s *BaseObjCListener) ExitAssignment_operator(ctx *Assignment_operatorContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
}

// EnterConditional_expression is called when production conditional_expression is entered.
func (s *BaseObjCListener) EnterConditional_expression(ctx *Conditional_expressionContext) {
	//if queue_assignment_operator.Len() > 0 {
	//	n := opts.TreeData{Name: queue_assignment_operator.Front().Value.(string)}
	//	s.current.Children = append(s.current.Children, &n)
	//	s.nodes = append(s.nodes, &n)
	//	queue_assignment_operator.Remove(queue_assignment_operator.Front())
	//	s.QFlags.queue_assignment_operator = true
	//}
}

// ExitConditional_expression is called when production conditional_expression is exited.
func (s *BaseObjCListener) ExitConditional_expression(ctx *Conditional_expressionContext) {
	//if s.QFlags.queue_assignment_operator {
	//	s.nodes = s.nodes[:len(s.nodes)-1]
	//	s.current = s.nodes[len(s.nodes)-1]
	//	s.QFlags.queue_assignment_operator = false
	//}
}

// EnterConstant_expression is called when production constant_expression is entered.
func (s *BaseObjCListener) EnterConstant_expression(ctx *Constant_expressionContext) {
	node := opts.TreeData{Name: "Constant_expression"}
	s.current.Children = append(s.current.Children, &node)
	//s.current = &node
	s.nodes = append(s.nodes, &node)
	if s.Flags.labeledStatement == 30 {
		arrDeep = append(arrDeep, Arr{"case "+ ctx.GetStart().GetText(), s.Flags.local})
	}
}

// ExitConstant_expression is called when production constant_expression is exited.
func (s *BaseObjCListener) ExitConstant_expression(ctx *Constant_expressionContext) {
	s.nodes = s.nodes[:len(s.nodes)-1]
	s.current = s.nodes[len(s.nodes)-1]
	if s.Flags.labeledStatement > 0 && arrDeep[len(arrDeep)-1].Level == s.Flags.local {
		arrDeep = append(arrDeep[:len(arrDeep)-1])
		var last = 0
		for e := count; 0 <= e; e-- {
			_, ok := visHash[e]
			if ok {
				last = e
				break
			}
		}

		var str = visHash[last].Scope
		for i := last; visHash[i].Scope == str; i-- {
			delete(visHash, i)
		}
	}
}

// EnterLogical_or_expression is called when production logical_or_expression is entered.
func (s *BaseObjCListener) EnterLogical_or_expression(ctx *Logical_or_expressionContext) {
	if queue_conditional.Len() > 0 {
		q := opts.TreeData{Name: "Conditional_expression"}
		s.current.Children = append(s.current.Children, &q)
		s.current = &q
		n := opts.TreeData{Name: queue_conditional.Front().Value.(string)}
		s.current.Children = append(s.current.Children, &n)
		s.nodes = append(s.nodes, &n)
		queue_conditional.Remove(queue_conditional.Front())
		s.QFlags.queue_conditional = true
	}
}

// ExitLogical_or_expression is called when production logical_or_expression is exited.
func (s *BaseObjCListener) ExitLogical_or_expression(ctx *Logical_or_expressionContext) {
	if s.QFlags.queue_conditional {
		s.nodes = s.nodes[:len(s.nodes)-1]
		s.current = s.nodes[len(s.nodes)-1]
		s.QFlags.queue_conditional = false
	}
}

// EnterLogical_and_expression is called when production logical_and_expression is entered.
func (s *BaseObjCListener) EnterLogical_and_expression(ctx *Logical_and_expressionContext) {
	if queue_logical_or.Len() > 0 {
		q := opts.TreeData{Name: "Logical_or_expression"}
		s.current.Children = append(s.current.Children, &q)
		s.current = &q
		n := opts.TreeData{Name: queue_logical_or.Front().Value.(string)}
		s.current.Children = append(s.current.Children, &n)
		s.nodes = append(s.nodes, &n)
		queue_logical_or.Remove(queue_logical_or.Front())
		s.QFlags.queue_logical_or = true
	}
}

// ExitLogical_and_expression is called when production logical_and_expression is exited.
func (s *BaseObjCListener) ExitLogical_and_expression(ctx *Logical_and_expressionContext) {
	if s.QFlags.queue_logical_or {
		s.nodes = s.nodes[:len(s.nodes)-1]
		s.current = s.nodes[len(s.nodes)-1]
		s.QFlags.queue_logical_or = false
	}
}

// EnterInclusive_or_expression is called when production inclusive_or_expression is entered.
func (s *BaseObjCListener) EnterInclusive_or_expression(ctx *Inclusive_or_expressionContext) {
	log.Println("HHHHEEEEEELLLLLLLPPPPPP:" + ctx.GetText() + " " + ctx.GetStart().GetText())
	if queue_logical_and.Len() > 0 {
		q := opts.TreeData{Name: "Logical_and_expression"}
		s.current.Children = append(s.current.Children, &q)
		s.current = &q
		n := opts.TreeData{Name: queue_logical_and.Front().Value.(string)}
		s.current.Children = append(s.current.Children, &n)
		s.nodes = append(s.nodes, &n)
		queue_logical_and.Remove(queue_logical_and.Front())
		s.QFlags.queue_logical_and = true
	}
}

// ExitInclusive_or_expression is called when production inclusive_or_expression is exited.
func (s *BaseObjCListener) ExitInclusive_or_expression(ctx *Inclusive_or_expressionContext) {
	if s.QFlags.queue_logical_and {
		s.nodes = s.nodes[:len(s.nodes)-1]
		s.current = s.nodes[len(s.nodes)-1]
		s.QFlags.queue_logical_and = false
	}
}

// EnterExclusive_or_expression is called when production exclusive_or_expression is entered.
func (s *BaseObjCListener) EnterExclusive_or_expression(ctx *Exclusive_or_expressionContext) {
	if queue_inclusive_or.Len() > 0 {
		q := opts.TreeData{Name: "inclusive_or_expression"}
		s.current.Children = append(s.current.Children, &q)
		s.current = &q
		n := opts.TreeData{Name: queue_inclusive_or.Front().Value.(string)}
		s.current.Children = append(s.current.Children, &n)
		s.nodes = append(s.nodes, &n)
		queue_inclusive_or.Remove(queue_inclusive_or.Front())
		s.QFlags.queue_inclusive_or = true
	}
}

// ExitExclusive_or_expression is called when production exclusive_or_expression is exited.
func (s *BaseObjCListener) ExitExclusive_or_expression(ctx *Exclusive_or_expressionContext) {
	if s.QFlags.queue_inclusive_or {
		s.nodes = s.nodes[:len(s.nodes)-1]
		s.current = s.nodes[len(s.nodes)-1]
		s.QFlags.queue_inclusive_or = false
	}
}

// EnterAnd_expression is called when production and_expression is entered.
func (s *BaseObjCListener) EnterAnd_expression(ctx *And_expressionContext) {
	if queue_exclusive_or.Len() > 0 {
		q := opts.TreeData{Name: "exclusive_or_expression"}
		s.current.Children = append(s.current.Children, &q)
		s.current = &q
		n := opts.TreeData{Name: queue_exclusive_or.Front().Value.(string)}
		s.current.Children = append(s.current.Children, &n)
		s.nodes = append(s.nodes, &n)
		queue_exclusive_or.Remove(queue_exclusive_or.Front())
		s.QFlags.queue_exclusive_or = true
	}
	if ctx.GetStart().GetTokenType() == 125 && ctx.GetStop().GetTokenType() == 70 {
		s.Flags.functionIn = true
		arrDeep = append(arrDeep, Arr{ctx.GetStart().GetText(), s.Flags.local})
	}
}

// ExitAnd_expression is called when production and_expression is exited.
func (s *BaseObjCListener) ExitAnd_expression(ctx *And_expressionContext) {
	if s.QFlags.queue_exclusive_or {
		s.nodes = s.nodes[:len(s.nodes)-1]
		s.current = s.nodes[len(s.nodes)-1]
		s.QFlags.queue_exclusive_or = false
	}
}

// EnterEquality_expression is called when production equality_expression is entered.
func (s *BaseObjCListener) EnterEquality_expression(ctx *Equality_expressionContext) {
	if queue_and.Len() > 0 {
		q := opts.TreeData{Name: "And_expression"}
		s.current.Children = append(s.current.Children, &q)
		s.current = &q
		n := opts.TreeData{Name: queue_and.Front().Value.(string)}
		s.current.Children = append(s.current.Children, &n)
		s.nodes = append(s.nodes, &n)
		queue_and.Remove(queue_and.Front())
		s.QFlags.queue_and = true
	}
}

// ExitEquality_expression is called when production equality_expression is exited.
func (s *BaseObjCListener) ExitEquality_expression(ctx *Equality_expressionContext) {
	if s.QFlags.queue_and {
		s.nodes = s.nodes[:len(s.nodes)-1]
		s.current = s.nodes[len(s.nodes)-1]
		s.QFlags.queue_and = false
	}
}

// EnterRelational_expression is called when production relational_expression is entered.
func (s *BaseObjCListener) EnterRelational_expression(ctx *Relational_expressionContext) {
	if queue_equality.Len() > 0 {
		q := opts.TreeData{Name: "Equality_expression"}
		s.current.Children = append(s.current.Children, &q)
		s.current = &q
		n := opts.TreeData{Name: queue_equality.Front().Value.(string)}
		s.current.Children = append(s.current.Children, &n)
		s.nodes = append(s.nodes, &n)
		queue_equality.Remove(queue_equality.Front())
		s.QFlags.queue_equality = true
	}
	//node := opts.TreeData{Name: "Relational_expression "+ ctx.GetStart().GetText()}
	//log.Println(ctx.GetStart().GetText(), ctx.GetStart().GetTokenType())
	//log.Println(ctx.GetStop().GetText(), ctx.GetStop().GetTokenType())
	//s.current.Children = append(s.current.Children, &node)
	//s.current = &node
	//s.nodes = append(s.nodes, &node)
}

// ExitRelational_expression is called when production relational_expression is exited.
func (s *BaseObjCListener) ExitRelational_expression(ctx *Relational_expressionContext) {
	if s.QFlags.queue_equality {
		s.nodes = s.nodes[:len(s.nodes)-1]
		s.current = s.nodes[len(s.nodes)-1]
		s.QFlags.queue_equality = false
	}
}

// EnterShift_expression is called when production shift_expression is entered.
func (s *BaseObjCListener) EnterShift_expression(ctx *Shift_expressionContext) {
	if queue_relational.Len() > 0 {
		q := opts.TreeData{Name: "Relational_expression"}
		s.current.Children = append(s.current.Children, &q)
		s.current = &q
		n := opts.TreeData{Name: queue_relational.Front().Value.(string)}
		s.current.Children = append(s.current.Children, &n)
		s.nodes = append(s.nodes, &n)
		queue_relational.Remove(queue_relational.Front())
		s.QFlags.queue_relational = true
	}
	//node := opts.TreeData{Name: ctx.GetStart().GetText()}
	//s.current.Children = append(s.current.Children, &node)
	//s.current = &node
	//s.nodes = append(s.nodes, &node)
}

// ExitShift_expression is called when production shift_expression is exited.
func (s *BaseObjCListener) ExitShift_expression(ctx *Shift_expressionContext) {
	if s.QFlags.queue_relational {
		s.nodes = s.nodes[:len(s.nodes)-1]
		s.current = s.nodes[len(s.nodes)-1]
		s.QFlags.queue_relational = false
	}
}

// EnterAdditive_expression is called when production additive_expression is entered.
func (s *BaseObjCListener) EnterAdditive_expression(ctx *Additive_expressionContext) {
	if queue_additive.Len() > 0 {
		n := opts.TreeData{Name: queue_additive.Front().Value.(string)}
		s.current.Children = append(s.current.Children, &n)
		s.nodes = append(s.nodes, &n)
		queue_additive.Remove(queue_additive.Front())
		s.QFlags.queue_additive = true
	} else if ctx.GetStart().GetTokenType() == 95 || ctx.GetStart().GetTokenType() == 96 {
		n := opts.TreeData{Name: ctx.GetStart().GetText()}
		s.current.Children = append(s.current.Children, &n)
		s.nodes = append(s.nodes, &n)
	}
	//node := opts.TreeData{Name: ctx.GetStart().GetText()}
	//s.current.Children = append(s.current.Children, &node)
	//s.current = &node
	//s.nodes = append(s.nodes, &node)
}

// ExitAdditive_expression is called when production additive_expression is exited.
func (s *BaseObjCListener) ExitAdditive_expression(ctx *Additive_expressionContext) {
	//s.nodes = s.nodes[:len(s.nodes)-1]
	//s.current = s.nodes[len(s.nodes)-1]
	if s.QFlags.queue_shift {
		s.nodes = s.nodes[:len(s.nodes)-1]
		s.current = s.nodes[len(s.nodes)-1]
		s.QFlags.queue_shift = false
	}
}

// EnterMultiplicative_expression is called when production multiplicative_expression is entered.
func (s *BaseObjCListener) EnterMultiplicative_expression(ctx *Multiplicative_expressionContext) {
	if queue_additive.Len() > 0 {
		q := opts.TreeData{Name: "Additive_expression"}
		s.current.Children = append(s.current.Children, &q)
		s.current = &q
		n := opts.TreeData{Name: queue_additive.Front().Value.(string)}
		s.current.Children = append(s.current.Children, &n)
		s.nodes = append(s.nodes, &n)
		queue_additive.Remove(queue_additive.Front())
		s.QFlags.queue_additive = true
	}
}

// ExitMultiplicative_expression is called when production multiplicative_expression is exited.
func (s *BaseObjCListener) ExitMultiplicative_expression(ctx *Multiplicative_expressionContext) {
	//s.nodes = s.nodes[:len(s.nodes)-1]
	//s.current = s.nodes[len(s.nodes)-1]
	if s.QFlags.queue_additive {
		s.nodes = s.nodes[:len(s.nodes)-1]
		s.current = s.nodes[len(s.nodes)-1]
		s.QFlags.queue_additive = false
	}
}

// EnterCast_expression is called when production cast_expression is entered.
func (s *BaseObjCListener) EnterCast_expression(ctx *Cast_expressionContext) {
	if queue_multiplicative.Len() > 0 {
		q := opts.TreeData{Name: "Multiplicative_expression"}
		s.current.Children = append(s.current.Children, &q)
		s.current = &q
		n := opts.TreeData{Name: queue_multiplicative.Front().Value.(string)}
		s.current.Children = append(s.current.Children, &n)
		s.nodes = append(s.nodes, &n)
		queue_multiplicative.Remove(queue_multiplicative.Front())
		s.QFlags.queue_multiplicative = true
	}
	if queue_unary_operator.Len() > 0 {
		q := opts.TreeData{Name: "Unary_operator"}
		s.current.Children = append(s.current.Children, &q)
		s.current = &q // TODO: если нужно сделать параллельно
		n := opts.TreeData{Name: queue_unary_operator.Front().Value.(string)}
		s.current.Children = append(s.current.Children, &n)
		s.nodes = append(s.nodes, &q)
		queue_unary_operator.Remove(queue_unary_operator.Front())
		s.QFlags.queue_unary_operator = true
	}
}

// ExitCast_expression is called when production cast_expression is exited.
func (s *BaseObjCListener) ExitCast_expression(ctx *Cast_expressionContext) {
	//s.nodes = s.nodes[:len(s.nodes)-1]
	//s.current = s.nodes[len(s.nodes)-1]
	if s.QFlags.queue_multiplicative {
		s.nodes = s.nodes[:len(s.nodes)-1]
		s.current = s.nodes[len(s.nodes)-1]
		s.QFlags.queue_multiplicative = false
	}
	if s.QFlags.queue_unary_operator {
		s.nodes = s.nodes[:len(s.nodes)-1]
		s.current = s.nodes[len(s.nodes)-1]
		s.QFlags.queue_unary_operator = false
	}
}

// EnterUnary_expression is called when production unary_expression is entered.
func (s *BaseObjCListener) EnterUnary_expression(ctx *Unary_expressionContext) {
	if ctx.GetStart().GetTokenType() == 93 || ctx.GetStart().GetTokenType() == 94 {
		q := opts.TreeData{Name: "Unary_expression"}
		s.current.Children = append(s.current.Children, &q)
		s.current = &q
		n := opts.TreeData{Name: ctx.GetStart().GetText()}
		s.current.Children = append(s.current.Children, &n)
		s.nodes = append(s.nodes, &n)
		s.QFlags.queue_cast = true
	}
}

// ExitUnary_expression is called when production unary_expression is exited.
func (s *BaseObjCListener) ExitUnary_expression(ctx *Unary_expressionContext) {
	//if ctx.GetStart().GetTokenType() == 93 || ctx.GetStart().GetTokenType() == 94 || ctx.GetStart().GetTokenType() == 97 || ctx.GetStart().GetTokenType() == 99 ||
	//	ctx.GetStart().GetTokenType() == 96 || ctx.GetStart().GetTokenType() == 84 || ctx.GetStart().GetTokenType() == 83 {
	//	s.nodes = s.nodes[:len(s.nodes)-1]
	//	s.current = s.nodes[len(s.nodes)-1]
	//}
	//s.nodes = s.nodes[:len(s.nodes)-1]
	//s.current = s.nodes[len(s.nodes)-1]
	if s.QFlags.queue_cast {
		s.nodes = s.nodes[:len(s.nodes)-1]
		s.current = s.nodes[len(s.nodes)-1]
		s.QFlags.queue_cast = false
	}
}

// EnterUnary_operator is called when production unary_operator is entered.
func (s *BaseObjCListener) EnterUnary_operator(ctx *Unary_operatorContext) {
	//if ctx.GetStart().GetTokenType() == 97 || ctx.GetStart().GetTokenType() == 99 ||
	//	ctx.GetStart().GetTokenType() == 96 || ctx.GetStart().GetTokenType() == 84 || ctx.GetStart().GetTokenType() == 83 {
	//	q := opts.TreeData{Name: "Unary_operator"}
	//	s.current.Children = append(s.current.Children, &q)
	//	s.current = &q
	//	n := opts.TreeData{Name: ctx.GetText()}
	//	s.current.Children = append(s.current.Children, &n)
	//	s.nodes = append(s.nodes, &n)
	//	s.QFlags.queue_unary_operator = true
	//}
	if queue_unary.Len() > 0 {
		q := opts.TreeData{Name: "Unary_operator"}
		s.current.Children = append(s.current.Children, &q)
		s.current = &q
		n := opts.TreeData{Name: queue_unary.Front().Value.(string)}
		s.current.Children = append(s.current.Children, &n)
		s.nodes = append(s.nodes, &n)
		queue_unary.Remove(queue_unary.Front())
		s.QFlags.queue_unary = true
	}
}

// ExitUnary_operator is called when production unary_operator is exited.
func (s *BaseObjCListener) ExitUnary_operator(ctx *Unary_operatorContext) {
	//s.nodes = s.nodes[:len(s.nodes)-1]
	//s.current = s.nodes[len(s.nodes)-1]
	if s.QFlags.queue_unary {
		s.nodes = s.nodes[:len(s.nodes)-1]
		s.current = s.nodes[len(s.nodes)-1]
		s.QFlags.queue_unary = false
	}
}

// EnterPostfix_expression is called when production postfix_expression is entered.
func (s *BaseObjCListener) EnterPostfix_expression(ctx *Postfix_expressionContext) {
	if ctx.GetStop().GetTokenType() == 93 || ctx.GetStop().GetTokenType() == 94 ||
		ctx.GetStop().GetTokenType() == 78 || ctx.GetStop().GetTokenType() == 77 {
		q := opts.TreeData{Name: "Postfix_expression"}
		s.current.Children = append(s.current.Children, &q)
		s.current = &q
		s.nodes = append(s.nodes, &q)
		s.QFlags.queue_postfix = true
	}
	//if ctx.GetStop().GetTokenType() == 93 || ctx.GetStop().GetTokenType() == 94 {
	//	q := opts.TreeData{Name: "Postfix_expression"}
	//	s.current.Children = append(s.current.Children, &q)
	//	s.current = &q
	//	n := opts.TreeData{Name: ctx.GetStop().GetText() + "(post)"}
	//	s.current.Children = append(s.current.Children, &n)
	//	s.nodes = append(s.nodes, &n)
	//}
}

// ExitPostfix_expression is called when production postfix_expression is exited.
func (s *BaseObjCListener) ExitPostfix_expression(ctx *Postfix_expressionContext) {
	if s.QFlags.queue_postfix {
		n := opts.TreeData{Name: ctx.GetStop().GetText()}
		s.current.Children = append(s.current.Children, &n)
		s.nodes = append(s.nodes, &n)
		s.QFlags.queue_postfix = false
	}
	//if ctx.GetStop().GetTokenType() == 93 || ctx.GetStop().GetTokenType() == 94 {
	//	s.nodes = s.nodes[:len(s.nodes)-1]
	//	s.current = s.nodes[len(s.nodes)-1]
	//}
	//if s.QFlags.queue_unary_operator {
	//	s.nodes = s.nodes[:len(s.nodes)-1]
	//	s.current = s.nodes[len(s.nodes)-1]
	//	s.QFlags.queue_unary_operator = false
	//}
}

// EnterArgument_expression_list is called when production argument_expression_list is entered.
func (s *BaseObjCListener) EnterArgument_expression_list(ctx *Argument_expression_listContext) {
	if queue_postfix.Len() > 0 {
		n := opts.TreeData{Name: queue_postfix.Front().Value.(string)}
		s.current.Children = append(s.current.Children, &n)
		s.nodes = append(s.nodes, &n)
		queue_postfix.Remove(queue_postfix.Front())
		s.QFlags.queue_postfix = true
	}
}

// ExitArgument_expression_list is called when production argument_expression_list is exited.
func (s *BaseObjCListener) ExitArgument_expression_list(ctx *Argument_expression_listContext) {
	if s.QFlags.queue_postfix {
		s.nodes = s.nodes[:len(s.nodes)-1]
		s.current = s.nodes[len(s.nodes)-1]
		s.QFlags.queue_postfix = false
	}
}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseObjCListener) EnterIdentifier(ctx *IdentifierContext) {
	if queue_argument_expression_list.Len() > 0 {
		n := opts.TreeData{Name: queue_argument_expression_list.Front().Value.(string)}
		s.current.Children = append(s.current.Children, &n)
		s.nodes = append(s.nodes, &n)
		queue_argument_expression_list.Remove(queue_argument_expression_list.Front())
		s.QFlags.queue_argument_expression_list = true
	}
}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseObjCListener) ExitIdentifier(ctx *IdentifierContext) {
	if s.QFlags.queue_argument_expression_list {
		s.nodes = s.nodes[:len(s.nodes)-1]
		s.current = s.nodes[len(s.nodes)-1]
		s.QFlags.queue_argument_expression_list = false
	}
}

// EnterConstant is called when production constant is entered.
func (s *BaseObjCListener) EnterConstant(ctx *ConstantContext) {
	if ctx.GetStart().GetTokenType() == 129 || ctx.GetStart().GetTokenType() == 128 ||
		ctx.GetStart().GetTokenType() == 130 || ctx.GetStart().GetTokenType() == 131 {
		//q := opts.TreeData{Name: "Constant"}
		//s.current.Children = append(s.current.Children, &q)
		//s.current = &q
		if ctx.GetStart().GetTokenType() == 129 {
			qw := opts.TreeData{Name: "Decimal"}
			s.current.Children = append(s.current.Children, &qw)
			s.current = &qw
		} else if ctx.GetStart().GetTokenType() == 128 {
			qw := opts.TreeData{Name: "Hex"}
			s.current.Children = append(s.current.Children, &qw)
			s.current = &qw
		} else if ctx.GetStart().GetTokenType() == 130 {
			qw := opts.TreeData{Name: "Octal"}
			s.current.Children = append(s.current.Children, &qw)
			s.current = &qw
		} else if ctx.GetStart().GetTokenType() == 131 {
			qw := opts.TreeData{Name: "Floating"}
			s.current.Children = append(s.current.Children, &qw)
			s.current = &qw
		}
		qw := opts.TreeData{Name: ctx.GetText()}
		s.current.Children = append(s.current.Children, &qw)
		s.current = &qw
	}

	if queue_relational.Len() > 0 {
		n := opts.TreeData{Name: queue_relational.Front().Value.(string)}
		s.current.Children = append(s.current.Children, &n)
		s.nodes = append(s.nodes, &n)
		queue_relational.Remove(queue_relational.Front())
		s.QFlags.queue_relational = true
	}
}

// ExitConstant is called when production constant is exited.
func (s *BaseObjCListener) ExitConstant(ctx *ConstantContext) {
	//s.nodes = s.nodes[:len(s.nodes)-1]
	//s.current = s.nodes[len(s.nodes)-1]
	if s.QFlags.queue_shift {
		s.nodes = s.nodes[:len(s.nodes)-1]
		s.current = s.nodes[len(s.nodes)-1]
		s.QFlags.queue_shift = false
	}
}
