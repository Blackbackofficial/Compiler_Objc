package parser // ObjC

import (
	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type Tree struct {
	charts.BaseConfiguration
}

// BaseObjCListener is a complete listener for a parse tree produced by ObjCParser.
type BaseObjCListener struct {
	Tree Tree
	Root opts.TreeData
	nodes []*opts.TreeData
	current *opts.TreeData
}

func NewBaseListener() *BaseObjCListener {
	l := BaseObjCListener{
		Tree:  Tree{},
		nodes: []*opts.TreeData{},
	}
	return &l
}

var _ ObjCListener = &BaseObjCListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseObjCListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseObjCListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseObjCListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	//node := opts.TreeData{Name: "EveryRule"}
	//s.current.Children = append(s.current.Children, &node)
	//s.current = &node
	//s.nodes = append(s.nodes, &node)
}

// ExitEveryRule is called when any rule is exited.
func (s *BaseObjCListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTranslation_unit is called when production translation_unit is entered.
func (s *BaseObjCListener) EnterTranslation_unit(ctx *Translation_unitContext) {
	//node := opts.TreeData{Name: "Translation_unit"}
	//s.current.Children = append(s.current.Children, &node)
	//s.current = &node
	//s.nodes = append(s.nodes, &node)
}

// ExitTranslation_unit is called when production translation_unit is exited.
func (s *BaseObjCListener) ExitTranslation_unit(ctx *Translation_unitContext) {}

// EnterExternal_declaration is called when production external_declaration is entered. вход
func (s *BaseObjCListener) EnterExternal_declaration(ctx *External_declarationContext) {
	s.Root = opts.TreeData{Name: "Start"}
	s.current = &s.Root
	s.nodes = append(s.nodes, &s.Root)
}

// ExitExternal_declaration is called when production external_declaration is exited. и выход
func (s *BaseObjCListener) ExitExternal_declaration(ctx *External_declarationContext) {}

// EnterPreprocessor_declaration is called when production preprocessor_declaration is entered.
func (s *BaseObjCListener) EnterPreprocessor_declaration(ctx *Preprocessor_declarationContext) {
	node := opts.TreeData{Name: "Preprocessor_declaration"}
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
	node := opts.TreeData{Name: "Class_interface"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitClass_interface is called when production class_interface is exited.
func (s *BaseObjCListener) ExitClass_interface(ctx *Class_interfaceContext) {}

// EnterCategory_interface is called when production category_interface is entered.
func (s *BaseObjCListener) EnterCategory_interface(ctx *Category_interfaceContext) {
	node := opts.TreeData{Name: "Category_interface"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitCategory_interface is called when production category_interface is exited.
func (s *BaseObjCListener) ExitCategory_interface(ctx *Category_interfaceContext) {}

// EnterClass_implementation is called when production class_implementation is entered.
func (s *BaseObjCListener) EnterClass_implementation(ctx *Class_implementationContext) {
	node := opts.TreeData{Name: "Class_implementation"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitClass_implementation is called when production class_implementation is exited.
func (s *BaseObjCListener) ExitClass_implementation(ctx *Class_implementationContext) {}

// EnterCategory_implementation is called when production category_implementation is entered.
func (s *BaseObjCListener) EnterCategory_implementation(ctx *Category_implementationContext) {
	node := opts.TreeData{Name: "Category_implementation"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitCategory_implementation is called when production category_implementation is exited.
func (s *BaseObjCListener) ExitCategory_implementation(ctx *Category_implementationContext) {}

// EnterProtocol_declaration is called when production protocol_declaration is entered.
func (s *BaseObjCListener) EnterProtocol_declaration(ctx *Protocol_declarationContext) {
	node := opts.TreeData{Name: "Protocol_declaration"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitProtocol_declaration is called when production protocol_declaration is exited.
func (s *BaseObjCListener) ExitProtocol_declaration(ctx *Protocol_declarationContext) {}

// EnterProtocol_declaration_list is called when production protocol_declaration_list is entered.
func (s *BaseObjCListener) EnterProtocol_declaration_list(ctx *Protocol_declaration_listContext) {
	node := opts.TreeData{Name: "Protocol_declaration_list"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitProtocol_declaration_list is called when production protocol_declaration_list is exited.
func (s *BaseObjCListener) ExitProtocol_declaration_list(ctx *Protocol_declaration_listContext) {}

// EnterClass_declaration_list is called when production class_declaration_list is entered.
func (s *BaseObjCListener) EnterClass_declaration_list(ctx *Class_declaration_listContext) {
	node := opts.TreeData{Name: "Class_declaration_list"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitClass_declaration_list is called when production class_declaration_list is exited.
func (s *BaseObjCListener) ExitClass_declaration_list(ctx *Class_declaration_listContext) {}

// EnterClass_list is called when production class_list is entered.
func (s *BaseObjCListener) EnterClass_list(ctx *Class_listContext) {
	node := opts.TreeData{Name: "Class_list"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitClass_list is called when production class_list is exited.
func (s *BaseObjCListener) ExitClass_list(ctx *Class_listContext) {}

// EnterProtocol_reference_list is called when production protocol_reference_list is entered.
func (s *BaseObjCListener) EnterProtocol_reference_list(ctx *Protocol_reference_listContext) {
	node := opts.TreeData{Name: "Protocol_reference_list"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitProtocol_reference_list is called when production protocol_reference_list is exited.
func (s *BaseObjCListener) ExitProtocol_reference_list(ctx *Protocol_reference_listContext) {}

// EnterProtocol_list is called when production protocol_list is entered.
func (s *BaseObjCListener) EnterProtocol_list(ctx *Protocol_listContext) {
	node := opts.TreeData{Name: "Protocol_list"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitProtocol_list is called when production protocol_list is exited.
func (s *BaseObjCListener) ExitProtocol_list(ctx *Protocol_listContext) {}

// EnterProperty_declaration is called when production property_declaration is entered.
func (s *BaseObjCListener) EnterProperty_declaration(ctx *Property_declarationContext) {
	node := opts.TreeData{Name: "Property_declaration"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitProperty_declaration is called when production property_declaration is exited.
func (s *BaseObjCListener) ExitProperty_declaration(ctx *Property_declarationContext) {}

// EnterProperty_attributes_declaration is called when production property_attributes_declaration is entered.
func (s *BaseObjCListener) EnterProperty_attributes_declaration(ctx *Property_attributes_declarationContext) {
	node := opts.TreeData{Name: "Property_attributes_declaration"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitProperty_attributes_declaration is called when production property_attributes_declaration is exited.
func (s *BaseObjCListener) ExitProperty_attributes_declaration(ctx *Property_attributes_declarationContext) {}

// EnterProperty_attributes_list is called when production property_attributes_list is entered.
func (s *BaseObjCListener) EnterProperty_attributes_list(ctx *Property_attributes_listContext) {
	node := opts.TreeData{Name: "Property_attributes_list"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitProperty_attributes_list is called when production property_attributes_list is exited.
func (s *BaseObjCListener) ExitProperty_attributes_list(ctx *Property_attributes_listContext) {}

// EnterProperty_attribute is called when production property_attribute is entered.
func (s *BaseObjCListener) EnterProperty_attribute(ctx *Property_attributeContext) {
	node := opts.TreeData{Name: "Property_attribute"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitProperty_attribute is called when production property_attribute is exited.
func (s *BaseObjCListener) ExitProperty_attribute(ctx *Property_attributeContext) {}

// EnterClass_name is called when production class_name is entered.
func (s *BaseObjCListener) EnterClass_name(ctx *Class_nameContext) {
	node := opts.TreeData{Name: "Class_name"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitClass_name is called when production class_name is exited.
func (s *BaseObjCListener) ExitClass_name(ctx *Class_nameContext) {}

// EnterSuperclass_name is called when production superclass_name is entered.
func (s *BaseObjCListener) EnterSuperclass_name(ctx *Superclass_nameContext) {
	node := opts.TreeData{Name: "Superclass_name"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitSuperclass_name is called when production superclass_name is exited.
func (s *BaseObjCListener) ExitSuperclass_name(ctx *Superclass_nameContext) {}

// EnterCategory_name is called when production category_name is entered.
func (s *BaseObjCListener) EnterCategory_name(ctx *Category_nameContext) {
	node := opts.TreeData{Name: "Category_name"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitCategory_name is called when production category_name is exited.
func (s *BaseObjCListener) ExitCategory_name(ctx *Category_nameContext) {}

// EnterProtocol_name is called when production protocol_name is entered.
func (s *BaseObjCListener) EnterProtocol_name(ctx *Protocol_nameContext) {
	node := opts.TreeData{Name: "Protocol_name"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitProtocol_name is called when production protocol_name is exited.
func (s *BaseObjCListener) ExitProtocol_name(ctx *Protocol_nameContext) {}

// EnterInstance_variables is called when production instance_variables is entered.
func (s *BaseObjCListener) EnterInstance_variables(ctx *Instance_variablesContext) {
	node := opts.TreeData{Name: "Instance_variables"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitInstance_variables is called when production instance_variables is exited.
func (s *BaseObjCListener) ExitInstance_variables(ctx *Instance_variablesContext) {}

// EnterVisibility_specification is called when production visibility_specification is entered.
func (s *BaseObjCListener) EnterVisibility_specification(ctx *Visibility_specificationContext) {
	node := opts.TreeData{Name: "Visibility_specification"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitVisibility_specification is called when production visibility_specification is exited.
func (s *BaseObjCListener) ExitVisibility_specification(ctx *Visibility_specificationContext) {}

// EnterInterface_declaration_list is called when production interface_declaration_list is entered.
func (s *BaseObjCListener) EnterInterface_declaration_list(ctx *Interface_declaration_listContext) {
	node := opts.TreeData{Name: "Interface_declaration_list"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitInterface_declaration_list is called when production interface_declaration_list is exited.
func (s *BaseObjCListener) ExitInterface_declaration_list(ctx *Interface_declaration_listContext) {}

// EnterClass_method_declaration is called when production class_method_declaration is entered.
func (s *BaseObjCListener) EnterClass_method_declaration(ctx *Class_method_declarationContext) {
	node := opts.TreeData{Name: "Class_method_declaration"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitClass_method_declaration is called when production class_method_declaration is exited.
func (s *BaseObjCListener) ExitClass_method_declaration(ctx *Class_method_declarationContext) {}

// EnterInstance_method_declaration is called when production instance_method_declaration is entered.
func (s *BaseObjCListener) EnterInstance_method_declaration(ctx *Instance_method_declarationContext) {
	node := opts.TreeData{Name: "Instance_method_declaration"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitInstance_method_declaration is called when production instance_method_declaration is exited.
func (s *BaseObjCListener) ExitInstance_method_declaration(ctx *Instance_method_declarationContext) {}

// EnterMethod_declaration is called when production method_declaration is entered.
func (s *BaseObjCListener) EnterMethod_declaration(ctx *Method_declarationContext) {
	node := opts.TreeData{Name: "Method_declaration"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitMethod_declaration is called when production method_declaration is exited.
func (s *BaseObjCListener) ExitMethod_declaration(ctx *Method_declarationContext) {}

// EnterImplementation_definition_list is called when production implementation_definition_list is entered.
func (s *BaseObjCListener) EnterImplementation_definition_list(ctx *Implementation_definition_listContext) {
	node := opts.TreeData{Name: "Implementation_definition_list"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitImplementation_definition_list is called when production implementation_definition_list is exited.
func (s *BaseObjCListener) ExitImplementation_definition_list(ctx *Implementation_definition_listContext) {}

// EnterClass_method_definition is called when production class_method_definition is entered.
func (s *BaseObjCListener) EnterClass_method_definition(ctx *Class_method_definitionContext) {
	node := opts.TreeData{Name: "Class_method_definition"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitClass_method_definition is called when production class_method_definition is exited.
func (s *BaseObjCListener) ExitClass_method_definition(ctx *Class_method_definitionContext) {}

// EnterInstance_method_definition is called when production instance_method_definition is entered.
func (s *BaseObjCListener) EnterInstance_method_definition(ctx *Instance_method_definitionContext) {
	node := opts.TreeData{Name: "Instance_method_definition"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitInstance_method_definition is called when production instance_method_definition is exited.
func (s *BaseObjCListener) ExitInstance_method_definition(ctx *Instance_method_definitionContext) {}

// EnterMethod_definition is called when production method_definition is entered.
func (s *BaseObjCListener) EnterMethod_definition(ctx *Method_definitionContext) {
	node := opts.TreeData{Name: "Method_definition"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitMethod_definition is called when production method_definition is exited.
func (s *BaseObjCListener) ExitMethod_definition(ctx *Method_definitionContext) {}

// EnterMethod_selector is called when production method_selector is entered.
func (s *BaseObjCListener) EnterMethod_selector(ctx *Method_selectorContext) {
	node := opts.TreeData{Name: "Method_selector"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitMethod_selector is called when production method_selector is exited.
func (s *BaseObjCListener) ExitMethod_selector(ctx *Method_selectorContext) {}

// EnterKeyword_declarator is called when production keyword_declarator is entered.
func (s *BaseObjCListener) EnterKeyword_declarator(ctx *Keyword_declaratorContext) {
	node := opts.TreeData{Name: "Keyword_declarator"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitKeyword_declarator is called when production keyword_declarator is exited.
func (s *BaseObjCListener) ExitKeyword_declarator(ctx *Keyword_declaratorContext) {}

// EnterSelector is called when production selector is entered.
func (s *BaseObjCListener) EnterSelector(ctx *SelectorContext) {
	node := opts.TreeData{Name: "Selector"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitSelector is called when production selector is exited.
func (s *BaseObjCListener) ExitSelector(ctx *SelectorContext) {}

// EnterMethod_type is called when production method_type is entered.
func (s *BaseObjCListener) EnterMethod_type(ctx *Method_typeContext) {
	node := opts.TreeData{Name: "Method_type"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitMethod_type is called when production method_type is exited.
func (s *BaseObjCListener) ExitMethod_type(ctx *Method_typeContext) {}

// EnterProperty_implementation is called when production property_implementation is entered.
func (s *BaseObjCListener) EnterProperty_implementation(ctx *Property_implementationContext) {
	node := opts.TreeData{Name: "Property_implementation"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitProperty_implementation is called when production property_implementation is exited.
func (s *BaseObjCListener) ExitProperty_implementation(ctx *Property_implementationContext) {}

// EnterProperty_synthesize_list is called when production property_synthesize_list is entered.
func (s *BaseObjCListener) EnterProperty_synthesize_list(ctx *Property_synthesize_listContext) {
	node := opts.TreeData{Name: "Property_synthesize_list"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitProperty_synthesize_list is called when production property_synthesize_list is exited.
func (s *BaseObjCListener) ExitProperty_synthesize_list(ctx *Property_synthesize_listContext) {}

// EnterProperty_synthesize_item is called when production property_synthesize_item is entered.
func (s *BaseObjCListener) EnterProperty_synthesize_item(ctx *Property_synthesize_itemContext) {
	node := opts.TreeData{Name: "Property_synthesize_item"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitProperty_synthesize_item is called when production property_synthesize_item is exited.
func (s *BaseObjCListener) ExitProperty_synthesize_item(ctx *Property_synthesize_itemContext) {}

// EnterBlock_type is called when production block_type is entered.
func (s *BaseObjCListener) EnterBlock_type(ctx *Block_typeContext) {
	node := opts.TreeData{Name: "Block_type"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitBlock_type is called when production block_type is exited.
func (s *BaseObjCListener) ExitBlock_type(ctx *Block_typeContext) {}

// EnterType_specifier is called when production type_specifier is entered.
func (s *BaseObjCListener) EnterType_specifier(ctx *Type_specifierContext) {
	node := opts.TreeData{Name: "Type_specifier"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitType_specifier is called when production type_specifier is exited.
func (s *BaseObjCListener) ExitType_specifier(ctx *Type_specifierContext) {}

// EnterType_qualifier is called when production type_qualifier is entered.
func (s *BaseObjCListener) EnterType_qualifier(ctx *Type_qualifierContext) {
	node := opts.TreeData{Name: "Type_qualifier"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitType_qualifier is called when production type_qualifier is exited.
func (s *BaseObjCListener) ExitType_qualifier(ctx *Type_qualifierContext) {}

// EnterProtocol_qualifier is called when production protocol_qualifier is entered.
func (s *BaseObjCListener) EnterProtocol_qualifier(ctx *Protocol_qualifierContext) {
	node := opts.TreeData{Name: "Protocol_qualifier"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitProtocol_qualifier is called when production protocol_qualifier is exited.
func (s *BaseObjCListener) ExitProtocol_qualifier(ctx *Protocol_qualifierContext) {}

// EnterPrimary_expression is called when production primary_expression is entered.
func (s *BaseObjCListener) EnterPrimary_expression(ctx *Primary_expressionContext) {
	node := opts.TreeData{Name: "Primary_expression"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitPrimary_expression is called when production primary_expression is exited.
func (s *BaseObjCListener) ExitPrimary_expression(ctx *Primary_expressionContext) {}

// EnterDictionary_pair is called when production dictionary_pair is entered.
func (s *BaseObjCListener) EnterDictionary_pair(ctx *Dictionary_pairContext) {
	node := opts.TreeData{Name: "Dictionary_pair"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitDictionary_pair is called when production dictionary_pair is exited.
func (s *BaseObjCListener) ExitDictionary_pair(ctx *Dictionary_pairContext) {}

// EnterDictionary_expression is called when production dictionary_expression is entered.
func (s *BaseObjCListener) EnterDictionary_expression(ctx *Dictionary_expressionContext) {
	node := opts.TreeData{Name: "Dictionary_expression"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitDictionary_expression is called when production dictionary_expression is exited.
func (s *BaseObjCListener) ExitDictionary_expression(ctx *Dictionary_expressionContext) {}

// EnterArray_expression is called when production array_expression is entered.
func (s *BaseObjCListener) EnterArray_expression(ctx *Array_expressionContext) {
	node := opts.TreeData{Name: "Array_expression"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitArray_expression is called when production array_expression is exited.
func (s *BaseObjCListener) ExitArray_expression(ctx *Array_expressionContext) {}

// EnterBox_expression is called when production box_expression is entered.
func (s *BaseObjCListener) EnterBox_expression(ctx *Box_expressionContext) {
	node := opts.TreeData{Name: "Box_expressio"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitBox_expression is called when production box_expression is exited.
func (s *BaseObjCListener) ExitBox_expression(ctx *Box_expressionContext) {}

// EnterBlock_parameters is called when production block_parameters is entered.
func (s *BaseObjCListener) EnterBlock_parameters(ctx *Block_parametersContext) {
	node := opts.TreeData{Name: "Block_parameters"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitBlock_parameters is called when production block_parameters is exited.
func (s *BaseObjCListener) ExitBlock_parameters(ctx *Block_parametersContext) {}

// EnterBlock_expression is called when production block_expression is entered.
func (s *BaseObjCListener) EnterBlock_expression(ctx *Block_expressionContext) {
	node := opts.TreeData{Name: "Block_expression"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitBlock_expression is called when production block_expression is exited.
func (s *BaseObjCListener) ExitBlock_expression(ctx *Block_expressionContext) {}

// EnterMessage_expression is called when production message_expression is entered.
func (s *BaseObjCListener) EnterMessage_expression(ctx *Message_expressionContext) {
	node := opts.TreeData{Name: "Message_expression"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitMessage_expression is called when production message_expression is exited.
func (s *BaseObjCListener) ExitMessage_expression(ctx *Message_expressionContext) {}

// EnterReceiver is called when production receiver is entered.
func (s *BaseObjCListener) EnterReceiver(ctx *ReceiverContext) {
	node := opts.TreeData{Name: "Receiver"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitReceiver is called when production receiver is exited.
func (s *BaseObjCListener) ExitReceiver(ctx *ReceiverContext) {}

// EnterMessage_selector is called when production message_selector is entered.
func (s *BaseObjCListener) EnterMessage_selector(ctx *Message_selectorContext) {
	node := opts.TreeData{Name: "Message_selector"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitMessage_selector is called when production message_selector is exited.
func (s *BaseObjCListener) ExitMessage_selector(ctx *Message_selectorContext) {}

// EnterKeyword_argument is called when production keyword_argument is entered.
func (s *BaseObjCListener) EnterKeyword_argument(ctx *Keyword_argumentContext) {
	node := opts.TreeData{Name: "Keyword_argument"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitKeyword_argument is called when production keyword_argument is exited.
func (s *BaseObjCListener) ExitKeyword_argument(ctx *Keyword_argumentContext) {}

// EnterSelector_expression is called when production selector_expression is entered.
func (s *BaseObjCListener) EnterSelector_expression(ctx *Selector_expressionContext) {
	node := opts.TreeData{Name: "Selector_expression"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitSelector_expression is called when production selector_expression is exited.
func (s *BaseObjCListener) ExitSelector_expression(ctx *Selector_expressionContext) {}

// EnterSelector_name is called when production selector_name is entered.
func (s *BaseObjCListener) EnterSelector_name(ctx *Selector_nameContext) {
	node := opts.TreeData{Name: "Selector_name"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitSelector_name is called when production selector_name is exited.
func (s *BaseObjCListener) ExitSelector_name(ctx *Selector_nameContext) {}

// EnterProtocol_expression is called when production protocol_expression is entered.
func (s *BaseObjCListener) EnterProtocol_expression(ctx *Protocol_expressionContext) {
	node := opts.TreeData{Name: "Protocol_expression"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitProtocol_expression is called when production protocol_expression is exited.
func (s *BaseObjCListener) ExitProtocol_expression(ctx *Protocol_expressionContext) {}

// EnterEncode_expression is called when production encode_expression is entered.
func (s *BaseObjCListener) EnterEncode_expression(ctx *Encode_expressionContext) {
	node := opts.TreeData{Name: "Encode_expression"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitEncode_expression is called when production encode_expression is exited.
func (s *BaseObjCListener) ExitEncode_expression(ctx *Encode_expressionContext) {}

// EnterType_variable_declarator is called when production type_variable_declarator is entered.
func (s *BaseObjCListener) EnterType_variable_declarator(ctx *Type_variable_declaratorContext) {
	node := opts.TreeData{Name: "Type_variable_declarator"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitType_variable_declarator is called when production type_variable_declarator is exited.
func (s *BaseObjCListener) ExitType_variable_declarator(ctx *Type_variable_declaratorContext) {}

// EnterTry_statement is called when production try_statement is entered.
func (s *BaseObjCListener) EnterTry_statement(ctx *Try_statementContext) {
	node := opts.TreeData{Name: "Try_statement"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitTry_statement is called when production try_statement is exited.
func (s *BaseObjCListener) ExitTry_statement(ctx *Try_statementContext) {}

// EnterCatch_statement is called when production catch_statement is entered.
func (s *BaseObjCListener) EnterCatch_statement(ctx *Catch_statementContext) {
	node := opts.TreeData{Name: "Catch_statement"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitCatch_statement is called when production catch_statement is exited.
func (s *BaseObjCListener) ExitCatch_statement(ctx *Catch_statementContext) {}

// EnterFinally_statement is called when production finally_statement is entered.
func (s *BaseObjCListener) EnterFinally_statement(ctx *Finally_statementContext) {
	node := opts.TreeData{Name: "Finally_statement"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitFinally_statement is called when production finally_statement is exited.
func (s *BaseObjCListener) ExitFinally_statement(ctx *Finally_statementContext) {}

// EnterThrow_statement is called when production throw_statement is entered.
func (s *BaseObjCListener) EnterThrow_statement(ctx *Throw_statementContext) {
	node := opts.TreeData{Name: "Throw_statemen"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitThrow_statement is called when production throw_statement is exited.
func (s *BaseObjCListener) ExitThrow_statement(ctx *Throw_statementContext) {}

// EnterTry_block is called when production try_block is entered.
func (s *BaseObjCListener) EnterTry_block(ctx *Try_blockContext) {
	node := opts.TreeData{Name: "Try_block"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitTry_block is called when production try_block is exited.
func (s *BaseObjCListener) ExitTry_block(ctx *Try_blockContext) {}

// EnterSynchronized_statement is called when production synchronized_statement is entered.
func (s *BaseObjCListener) EnterSynchronized_statement(ctx *Synchronized_statementContext) {
	node := opts.TreeData{Name: "Synchronized_statement"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitSynchronized_statement is called when production synchronized_statement is exited.
func (s *BaseObjCListener) ExitSynchronized_statement(ctx *Synchronized_statementContext) {}

// EnterAutorelease_statement is called when production autorelease_statement is entered.
func (s *BaseObjCListener) EnterAutorelease_statement(ctx *Autorelease_statementContext) {
	node := opts.TreeData{Name: "Autorelease_statement"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitAutorelease_statement is called when production autorelease_statement is exited.
func (s *BaseObjCListener) ExitAutorelease_statement(ctx *Autorelease_statementContext) {}

// EnterFunction_definition is called when production function_definition is entered.
func (s *BaseObjCListener) EnterFunction_definition(ctx *Function_definitionContext) {
	node := opts.TreeData{Name: "Function_definition"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitFunction_definition is called when production function_definition is exited.
func (s *BaseObjCListener) ExitFunction_definition(ctx *Function_definitionContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseObjCListener) EnterDeclaration(ctx *DeclarationContext) {
	node := opts.TreeData{Name: "Declaration"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseObjCListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterDeclaration_specifiers is called when production declaration_specifiers is entered.
func (s *BaseObjCListener) EnterDeclaration_specifiers(ctx *Declaration_specifiersContext) {
	node := opts.TreeData{Name: "Declaration_specifier"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitDeclaration_specifiers is called when production declaration_specifiers is exited.
func (s *BaseObjCListener) ExitDeclaration_specifiers(ctx *Declaration_specifiersContext) {}

// EnterArc_behaviour_specifier is called when production arc_behaviour_specifier is entered.
func (s *BaseObjCListener) EnterArc_behaviour_specifier(ctx *Arc_behaviour_specifierContext) {
	node := opts.TreeData{Name: "Arc_behaviour_specifier"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitArc_behaviour_specifier is called when production arc_behaviour_specifier is exited.
func (s *BaseObjCListener) ExitArc_behaviour_specifier(ctx *Arc_behaviour_specifierContext) {}

// EnterStorage_class_specifier is called when production storage_class_specifier is entered.
func (s *BaseObjCListener) EnterStorage_class_specifier(ctx *Storage_class_specifierContext) {
	node := opts.TreeData{Name: "Storage_class_specifier"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitStorage_class_specifier is called when production storage_class_specifier is exited.
func (s *BaseObjCListener) ExitStorage_class_specifier(ctx *Storage_class_specifierContext) {}

// EnterInit_declarator_list is called when production init_declarator_list is entered.
func (s *BaseObjCListener) EnterInit_declarator_list(ctx *Init_declarator_listContext) {
	node := opts.TreeData{Name: "Init_declarator_list"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitInit_declarator_list is called when production init_declarator_list is exited.
func (s *BaseObjCListener) ExitInit_declarator_list(ctx *Init_declarator_listContext) {}

// EnterInit_declarator is called when production init_declarator is entered.
func (s *BaseObjCListener) EnterInit_declarator(ctx *Init_declaratorContext) {
	node := opts.TreeData{Name: "Init_declarator"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitInit_declarator is called when production init_declarator is exited.
func (s *BaseObjCListener) ExitInit_declarator(ctx *Init_declaratorContext) {}

// EnterStruct_or_union_specifier is called when production struct_or_union_specifier is entered.
func (s *BaseObjCListener) EnterStruct_or_union_specifier(ctx *Struct_or_union_specifierContext) {
	node := opts.TreeData{Name: "Struct_or_union_specifier"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitStruct_or_union_specifier is called when production struct_or_union_specifier is exited.
func (s *BaseObjCListener) ExitStruct_or_union_specifier(ctx *Struct_or_union_specifierContext) {}

// EnterStruct_declaration is called when production struct_declaration is entered.
func (s *BaseObjCListener) EnterStruct_declaration(ctx *Struct_declarationContext) {
	node := opts.TreeData{Name: "Struct_declaration"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitStruct_declaration is called when production struct_declaration is exited.
func (s *BaseObjCListener) ExitStruct_declaration(ctx *Struct_declarationContext) {}

// EnterSpecifier_qualifier_list is called when production specifier_qualifier_list is entered.
func (s *BaseObjCListener) EnterSpecifier_qualifier_list(ctx *Specifier_qualifier_listContext) {
	node := opts.TreeData{Name: "Specifier_qualifier_list"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitSpecifier_qualifier_list is called when production specifier_qualifier_list is exited.
func (s *BaseObjCListener) ExitSpecifier_qualifier_list(ctx *Specifier_qualifier_listContext) {}

// EnterStruct_declarator_list is called when production struct_declarator_list is entered.
func (s *BaseObjCListener) EnterStruct_declarator_list(ctx *Struct_declarator_listContext) {
	node := opts.TreeData{Name: "Struct_declarator_list"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitStruct_declarator_list is called when production struct_declarator_list is exited.
func (s *BaseObjCListener) ExitStruct_declarator_list(ctx *Struct_declarator_listContext) {}

// EnterStruct_declarator is called when production struct_declarator is entered.
func (s *BaseObjCListener) EnterStruct_declarator(ctx *Struct_declaratorContext) {
	node := opts.TreeData{Name: "Struct_declarator"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitStruct_declarator is called when production struct_declarator is exited.
func (s *BaseObjCListener) ExitStruct_declarator(ctx *Struct_declaratorContext) {}

// EnterEnum_specifier is called when production enum_specifier is entered.
func (s *BaseObjCListener) EnterEnum_specifier(ctx *Enum_specifierContext) {
	node := opts.TreeData{Name: "Enum_specifie"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitEnum_specifier is called when production enum_specifier is exited.
func (s *BaseObjCListener) ExitEnum_specifier(ctx *Enum_specifierContext) {}

// EnterEnumerator_list is called when production enumerator_list is entered.
func (s *BaseObjCListener) EnterEnumerator_list(ctx *Enumerator_listContext) {
	node := opts.TreeData{Name: "Enumerator_list"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitEnumerator_list is called when production enumerator_list is exited.
func (s *BaseObjCListener) ExitEnumerator_list(ctx *Enumerator_listContext) {}

// EnterEnumerator is called when production enumerator is entered.
func (s *BaseObjCListener) EnterEnumerator(ctx *EnumeratorContext) {
	node := opts.TreeData{Name: "Enumerator"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitEnumerator is called when production enumerator is exited.
func (s *BaseObjCListener) ExitEnumerator(ctx *EnumeratorContext) {}

// EnterPointer is called when production pointer is entered.
func (s *BaseObjCListener) EnterPointer(ctx *PointerContext) {
	node := opts.TreeData{Name: "Pointer"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitPointer is called when production pointer is exited.
func (s *BaseObjCListener) ExitPointer(ctx *PointerContext) {}

// EnterDeclarator is called when production declarator is entered.
func (s *BaseObjCListener) EnterDeclarator(ctx *DeclaratorContext) {
	node := opts.TreeData{Name: "Declarator"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitDeclarator is called when production declarator is exited.
func (s *BaseObjCListener) ExitDeclarator(ctx *DeclaratorContext) {}

// EnterDirect_declarator is called when production direct_declarator is entered.
func (s *BaseObjCListener) EnterDirect_declarator(ctx *Direct_declaratorContext) {
	node := opts.TreeData{Name: "Direct_declarator"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitDirect_declarator is called when production direct_declarator is exited.
func (s *BaseObjCListener) ExitDirect_declarator(ctx *Direct_declaratorContext) {}

// EnterDeclarator_suffix is called when production declarator_suffix is entered.
func (s *BaseObjCListener) EnterDeclarator_suffix(ctx *Declarator_suffixContext) {
	node := opts.TreeData{Name: "Declarator_suffix"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitDeclarator_suffix is called when production declarator_suffix is exited.
func (s *BaseObjCListener) ExitDeclarator_suffix(ctx *Declarator_suffixContext) {}

// EnterParameter_list is called when production parameter_list is entered.
func (s *BaseObjCListener) EnterParameter_list(ctx *Parameter_listContext) {
	node := opts.TreeData{Name: "Parameter_list"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitParameter_list is called when production parameter_list is exited.
func (s *BaseObjCListener) ExitParameter_list(ctx *Parameter_listContext) {}

// EnterParameter_declaration is called when production parameter_declaration is entered.
func (s *BaseObjCListener) EnterParameter_declaration(ctx *Parameter_declarationContext) {
	node := opts.TreeData{Name: "Parameter_declaration"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitParameter_declaration is called when production parameter_declaration is exited.
func (s *BaseObjCListener) ExitParameter_declaration(ctx *Parameter_declarationContext) {}

// EnterInitializer is called when production initializer is entered.
func (s *BaseObjCListener) EnterInitializer(ctx *InitializerContext) {
	node := opts.TreeData{Name: "Initializer"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitInitializer is called when production initializer is exited.
func (s *BaseObjCListener) ExitInitializer(ctx *InitializerContext) {}

// EnterType_name is called when production type_name is entered.
func (s *BaseObjCListener) EnterType_name(ctx *Type_nameContext) {
	node := opts.TreeData{Name: "Type_name"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitType_name is called when production type_name is exited.
func (s *BaseObjCListener) ExitType_name(ctx *Type_nameContext) {}

// EnterAbstract_declarator is called when production abstract_declarator is entered.
func (s *BaseObjCListener) EnterAbstract_declarator(ctx *Abstract_declaratorContext) {
	node := opts.TreeData{Name: "Abstract_declarator"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitAbstract_declarator is called when production abstract_declarator is exited.
func (s *BaseObjCListener) ExitAbstract_declarator(ctx *Abstract_declaratorContext) {}

// EnterAbstract_declarator_suffix is called when production abstract_declarator_suffix is entered.
func (s *BaseObjCListener) EnterAbstract_declarator_suffix(ctx *Abstract_declarator_suffixContext) {
	node := opts.TreeData{Name: "Abstract_declarator_suffi"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitAbstract_declarator_suffix is called when production abstract_declarator_suffix is exited.
func (s *BaseObjCListener) ExitAbstract_declarator_suffix(ctx *Abstract_declarator_suffixContext) {}

// EnterParameter_declaration_list is called when production parameter_declaration_list is entered.
func (s *BaseObjCListener) EnterParameter_declaration_list(ctx *Parameter_declaration_listContext) {
	node := opts.TreeData{Name: "Parameter_declaration_list"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitParameter_declaration_list is called when production parameter_declaration_list is exited.
func (s *BaseObjCListener) ExitParameter_declaration_list(ctx *Parameter_declaration_listContext) {}

// EnterStatement_list is called when production statement_list is entered.
func (s *BaseObjCListener) EnterStatement_list(ctx *Statement_listContext) {
	node := opts.TreeData{Name: "Statement_list"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitStatement_list is called when production statement_list is exited.
func (s *BaseObjCListener) ExitStatement_list(ctx *Statement_listContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseObjCListener) EnterStatement(ctx *StatementContext) {
	node := opts.TreeData{Name: "Statement"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitStatement is called when production statement is exited.
func (s *BaseObjCListener) ExitStatement(ctx *StatementContext) {}

// EnterLabeled_statement is called when production labeled_statement is entered.
func (s *BaseObjCListener) EnterLabeled_statement(ctx *Labeled_statementContext) {
	node := opts.TreeData{Name: "Labeled_statement"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitLabeled_statement is called when production labeled_statement is exited.
func (s *BaseObjCListener) ExitLabeled_statement(ctx *Labeled_statementContext) {}

// EnterCompound_statement is called when production compound_statement is entered.
func (s *BaseObjCListener) EnterCompound_statement(ctx *Compound_statementContext) {
	node := opts.TreeData{Name: "Compound_statement"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitCompound_statement is called when production compound_statement is exited.
func (s *BaseObjCListener) ExitCompound_statement(ctx *Compound_statementContext) {}

// EnterSelection_statement is called when production selection_statement is entered.
func (s *BaseObjCListener) EnterSelection_statement(ctx *Selection_statementContext) {
	node := opts.TreeData{Name: "Selection_statement"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitSelection_statement is called when production selection_statement is exited.
func (s *BaseObjCListener) ExitSelection_statement(ctx *Selection_statementContext) {}

// EnterFor_in_statement is called when production for_in_statement is entered.
func (s *BaseObjCListener) EnterFor_in_statement(ctx *For_in_statementContext) {
	node := opts.TreeData{Name: "For_in_statement"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitFor_in_statement is called when production for_in_statement is exited.
func (s *BaseObjCListener) ExitFor_in_statement(ctx *For_in_statementContext) {}

// EnterFor_statement is called when production for_statement is entered.
func (s *BaseObjCListener) EnterFor_statement(ctx *For_statementContext) {
	node := opts.TreeData{Name: "For_statement"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitFor_statement is called when production for_statement is exited.
func (s *BaseObjCListener) ExitFor_statement(ctx *For_statementContext) {}

// EnterWhile_statement is called when production while_statement is entered.
func (s *BaseObjCListener) EnterWhile_statement(ctx *While_statementContext) {
	node := opts.TreeData{Name: "While_statement"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitWhile_statement is called when production while_statement is exited.
func (s *BaseObjCListener) ExitWhile_statement(ctx *While_statementContext) {}

// EnterDo_statement is called when production do_statement is entered.
func (s *BaseObjCListener) EnterDo_statement(ctx *Do_statementContext) {
	node := opts.TreeData{Name: "Do_statement"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitDo_statement is called when production do_statement is exited.
func (s *BaseObjCListener) ExitDo_statement(ctx *Do_statementContext) {}

// EnterIteration_statement is called when production iteration_statement is entered.
func (s *BaseObjCListener) EnterIteration_statement(ctx *Iteration_statementContext) {
	node := opts.TreeData{Name: "Iteration_statement"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitIteration_statement is called when production iteration_statement is exited.
func (s *BaseObjCListener) ExitIteration_statement(ctx *Iteration_statementContext) {}

// EnterJump_statement is called when production jump_statement is entered.
func (s *BaseObjCListener) EnterJump_statement(ctx *Jump_statementContext) {
	node := opts.TreeData{Name: "Jump_statement"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitJump_statement is called when production jump_statement is exited.
func (s *BaseObjCListener) ExitJump_statement(ctx *Jump_statementContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseObjCListener) EnterExpression(ctx *ExpressionContext) {
	node := opts.TreeData{Name: "Expression"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitExpression is called when production expression is exited.
func (s *BaseObjCListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAssignment_expression is called when production assignment_expression is entered.
func (s *BaseObjCListener) EnterAssignment_expression(ctx *Assignment_expressionContext) {
	node := opts.TreeData{Name: "Assignment_expression"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitAssignment_expression is called when production assignment_expression is exited.
func (s *BaseObjCListener) ExitAssignment_expression(ctx *Assignment_expressionContext) {}

// EnterAssignment_operator is called when production assignment_operator is entered.
func (s *BaseObjCListener) EnterAssignment_operator(ctx *Assignment_operatorContext) {
	node := opts.TreeData{Name: "Assignment_operator"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitAssignment_operator is called when production assignment_operator is exited.
func (s *BaseObjCListener) ExitAssignment_operator(ctx *Assignment_operatorContext) {}

// EnterConditional_expression is called when production conditional_expression is entered.
func (s *BaseObjCListener) EnterConditional_expression(ctx *Conditional_expressionContext) {
	node := opts.TreeData{Name: "Conditional_expression"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitConditional_expression is called when production conditional_expression is exited.
func (s *BaseObjCListener) ExitConditional_expression(ctx *Conditional_expressionContext) {}

// EnterConstant_expression is called when production constant_expression is entered.
func (s *BaseObjCListener) EnterConstant_expression(ctx *Constant_expressionContext) {
	node := opts.TreeData{Name: "Constant_expression"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitConstant_expression is called when production constant_expression is exited.
func (s *BaseObjCListener) ExitConstant_expression(ctx *Constant_expressionContext) {}

// EnterLogical_or_expression is called when production logical_or_expression is entered.
func (s *BaseObjCListener) EnterLogical_or_expression(ctx *Logical_or_expressionContext) {
	node := opts.TreeData{Name: "Logical_or_expression"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitLogical_or_expression is called when production logical_or_expression is exited.
func (s *BaseObjCListener) ExitLogical_or_expression(ctx *Logical_or_expressionContext) {}

// EnterLogical_and_expression is called when production logical_and_expression is entered.
func (s *BaseObjCListener) EnterLogical_and_expression(ctx *Logical_and_expressionContext) {
	node := opts.TreeData{Name: "Logical_and_expression"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitLogical_and_expression is called when production logical_and_expression is exited.
func (s *BaseObjCListener) ExitLogical_and_expression(ctx *Logical_and_expressionContext) {}

// EnterInclusive_or_expression is called when production inclusive_or_expression is entered.
func (s *BaseObjCListener) EnterInclusive_or_expression(ctx *Inclusive_or_expressionContext) {
	node := opts.TreeData{Name: "Inclusive_or_expression"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitInclusive_or_expression is called when production inclusive_or_expression is exited.
func (s *BaseObjCListener) ExitInclusive_or_expression(ctx *Inclusive_or_expressionContext) {}

// EnterExclusive_or_expression is called when production exclusive_or_expression is entered.
func (s *BaseObjCListener) EnterExclusive_or_expression(ctx *Exclusive_or_expressionContext) {
	node := opts.TreeData{Name: "Exclusive_or_expression"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitExclusive_or_expression is called when production exclusive_or_expression is exited.
func (s *BaseObjCListener) ExitExclusive_or_expression(ctx *Exclusive_or_expressionContext) {}

// EnterAnd_expression is called when production and_expression is entered.
func (s *BaseObjCListener) EnterAnd_expression(ctx *And_expressionContext) {
	node := opts.TreeData{Name: "And_expression"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitAnd_expression is called when production and_expression is exited.
func (s *BaseObjCListener) ExitAnd_expression(ctx *And_expressionContext) {}

// EnterEquality_expression is called when production equality_expression is entered.
func (s *BaseObjCListener) EnterEquality_expression(ctx *Equality_expressionContext) {
	node := opts.TreeData{Name: "Equality_expression"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitEquality_expression is called when production equality_expression is exited.
func (s *BaseObjCListener) ExitEquality_expression(ctx *Equality_expressionContext) {}

// EnterRelational_expression is called when production relational_expression is entered.
func (s *BaseObjCListener) EnterRelational_expression(ctx *Relational_expressionContext) {
	node := opts.TreeData{Name: "Relational_expression"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitRelational_expression is called when production relational_expression is exited.
func (s *BaseObjCListener) ExitRelational_expression(ctx *Relational_expressionContext) {}

// EnterShift_expression is called when production shift_expression is entered.
func (s *BaseObjCListener) EnterShift_expression(ctx *Shift_expressionContext) {
	node := opts.TreeData{Name: "Shift_expression"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitShift_expression is called when production shift_expression is exited.
func (s *BaseObjCListener) ExitShift_expression(ctx *Shift_expressionContext) {}

// EnterAdditive_expression is called when production additive_expression is entered.
func (s *BaseObjCListener) EnterAdditive_expression(ctx *Additive_expressionContext) {
	node := opts.TreeData{Name: "Additive_expression"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitAdditive_expression is called when production additive_expression is exited.
func (s *BaseObjCListener) ExitAdditive_expression(ctx *Additive_expressionContext) {}

// EnterMultiplicative_expression is called when production multiplicative_expression is entered.
func (s *BaseObjCListener) EnterMultiplicative_expression(ctx *Multiplicative_expressionContext) {
	node := opts.TreeData{Name: "Multiplicative_expression"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitMultiplicative_expression is called when production multiplicative_expression is exited.
func (s *BaseObjCListener) ExitMultiplicative_expression(ctx *Multiplicative_expressionContext) {}

// EnterCast_expression is called when production cast_expression is entered.
func (s *BaseObjCListener) EnterCast_expression(ctx *Cast_expressionContext) {
	node := opts.TreeData{Name: "Cast_expression"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitCast_expression is called when production cast_expression is exited.
func (s *BaseObjCListener) ExitCast_expression(ctx *Cast_expressionContext) {}

// EnterUnary_expression is called when production unary_expression is entered.
func (s *BaseObjCListener) EnterUnary_expression(ctx *Unary_expressionContext) {
	node := opts.TreeData{Name: "Unary_expression"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitUnary_expression is called when production unary_expression is exited.
func (s *BaseObjCListener) ExitUnary_expression(ctx *Unary_expressionContext) {}

// EnterUnary_operator is called when production unary_operator is entered.
func (s *BaseObjCListener) EnterUnary_operator(ctx *Unary_operatorContext) {
	node := opts.TreeData{Name: "Unary_operator"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitUnary_operator is called when production unary_operator is exited.
func (s *BaseObjCListener) ExitUnary_operator(ctx *Unary_operatorContext) {}

// EnterPostfix_expression is called when production postfix_expression is entered.
func (s *BaseObjCListener) EnterPostfix_expression(ctx *Postfix_expressionContext) {
	node := opts.TreeData{Name: "Postfix_expression"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitPostfix_expression is called when production postfix_expression is exited.
func (s *BaseObjCListener) ExitPostfix_expression(ctx *Postfix_expressionContext) {}

// EnterArgument_expression_list is called when production argument_expression_list is entered.
func (s *BaseObjCListener) EnterArgument_expression_list(ctx *Argument_expression_listContext) {
	node := opts.TreeData{Name: "Argument_expression_list"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitArgument_expression_list is called when production argument_expression_list is exited.
func (s *BaseObjCListener) ExitArgument_expression_list(ctx *Argument_expression_listContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseObjCListener) EnterIdentifier(ctx *IdentifierContext) {
	node := opts.TreeData{Name: "Identifier"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseObjCListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseObjCListener) EnterConstant(ctx *ConstantContext) {
	node := opts.TreeData{Name: "Constant"}
	s.current.Children = append(s.current.Children, &node)
	s.current = &node
	s.nodes = append(s.nodes, &node)
}

// ExitConstant is called when production constant is exited.
func (s *BaseObjCListener) ExitConstant(ctx *ConstantContext) {}
