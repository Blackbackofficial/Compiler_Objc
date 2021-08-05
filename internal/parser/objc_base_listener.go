// Code generated from ObjC.g4 by ANTLR 4.9.2. DO NOT EDIT.

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

var _ ObjCListener = &BaseObjCListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseObjCListener) VisitTerminal(node antlr.TerminalNode) {}

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

// EnterExternal_declaration is called when production external_declaration is entered.
func (s *BaseObjCListener) EnterExternal_declaration(ctx *External_declarationContext) {}

// ExitExternal_declaration is called when production external_declaration is exited.
func (s *BaseObjCListener) ExitExternal_declaration(ctx *External_declarationContext) {}

// EnterPreprocessor_declaration is called when production preprocessor_declaration is entered.
func (s *BaseObjCListener) EnterPreprocessor_declaration(ctx *Preprocessor_declarationContext) {}

// ExitPreprocessor_declaration is called when production preprocessor_declaration is exited.
func (s *BaseObjCListener) ExitPreprocessor_declaration(ctx *Preprocessor_declarationContext) {}

// EnterClass_interface is called when production class_interface is entered.
func (s *BaseObjCListener) EnterClass_interface(ctx *Class_interfaceContext) {}

// ExitClass_interface is called when production class_interface is exited.
func (s *BaseObjCListener) ExitClass_interface(ctx *Class_interfaceContext) {}

// EnterCategory_interface is called when production category_interface is entered.
func (s *BaseObjCListener) EnterCategory_interface(ctx *Category_interfaceContext) {}

// ExitCategory_interface is called when production category_interface is exited.
func (s *BaseObjCListener) ExitCategory_interface(ctx *Category_interfaceContext) {}

// EnterClass_implementation is called when production class_implementation is entered.
func (s *BaseObjCListener) EnterClass_implementation(ctx *Class_implementationContext) {}

// ExitClass_implementation is called when production class_implementation is exited.
func (s *BaseObjCListener) ExitClass_implementation(ctx *Class_implementationContext) {}

// EnterCategory_implementation is called when production category_implementation is entered.
func (s *BaseObjCListener) EnterCategory_implementation(ctx *Category_implementationContext) {}

// ExitCategory_implementation is called when production category_implementation is exited.
func (s *BaseObjCListener) ExitCategory_implementation(ctx *Category_implementationContext) {}

// EnterProtocol_declaration is called when production protocol_declaration is entered.
func (s *BaseObjCListener) EnterProtocol_declaration(ctx *Protocol_declarationContext) {}

// ExitProtocol_declaration is called when production protocol_declaration is exited.
func (s *BaseObjCListener) ExitProtocol_declaration(ctx *Protocol_declarationContext) {}

// EnterProtocol_declaration_list is called when production protocol_declaration_list is entered.
func (s *BaseObjCListener) EnterProtocol_declaration_list(ctx *Protocol_declaration_listContext) {}

// ExitProtocol_declaration_list is called when production protocol_declaration_list is exited.
func (s *BaseObjCListener) ExitProtocol_declaration_list(ctx *Protocol_declaration_listContext) {}

// EnterClass_declaration_list is called when production class_declaration_list is entered.
func (s *BaseObjCListener) EnterClass_declaration_list(ctx *Class_declaration_listContext) {}

// ExitClass_declaration_list is called when production class_declaration_list is exited.
func (s *BaseObjCListener) ExitClass_declaration_list(ctx *Class_declaration_listContext) {}

// EnterClass_list is called when production class_list is entered.
func (s *BaseObjCListener) EnterClass_list(ctx *Class_listContext) {}

// ExitClass_list is called when production class_list is exited.
func (s *BaseObjCListener) ExitClass_list(ctx *Class_listContext) {}

// EnterProtocol_reference_list is called when production protocol_reference_list is entered.
func (s *BaseObjCListener) EnterProtocol_reference_list(ctx *Protocol_reference_listContext) {}

// ExitProtocol_reference_list is called when production protocol_reference_list is exited.
func (s *BaseObjCListener) ExitProtocol_reference_list(ctx *Protocol_reference_listContext) {}

// EnterProtocol_list is called when production protocol_list is entered.
func (s *BaseObjCListener) EnterProtocol_list(ctx *Protocol_listContext) {}

// ExitProtocol_list is called when production protocol_list is exited.
func (s *BaseObjCListener) ExitProtocol_list(ctx *Protocol_listContext) {}

// EnterProperty_declaration is called when production property_declaration is entered.
func (s *BaseObjCListener) EnterProperty_declaration(ctx *Property_declarationContext) {}

// ExitProperty_declaration is called when production property_declaration is exited.
func (s *BaseObjCListener) ExitProperty_declaration(ctx *Property_declarationContext) {}

// EnterProperty_attributes_declaration is called when production property_attributes_declaration is entered.
func (s *BaseObjCListener) EnterProperty_attributes_declaration(ctx *Property_attributes_declarationContext) {}

// ExitProperty_attributes_declaration is called when production property_attributes_declaration is exited.
func (s *BaseObjCListener) ExitProperty_attributes_declaration(ctx *Property_attributes_declarationContext) {}

// EnterProperty_attributes_list is called when production property_attributes_list is entered.
func (s *BaseObjCListener) EnterProperty_attributes_list(ctx *Property_attributes_listContext) {}

// ExitProperty_attributes_list is called when production property_attributes_list is exited.
func (s *BaseObjCListener) ExitProperty_attributes_list(ctx *Property_attributes_listContext) {}

// EnterProperty_attribute is called when production property_attribute is entered.
func (s *BaseObjCListener) EnterProperty_attribute(ctx *Property_attributeContext) {}

// ExitProperty_attribute is called when production property_attribute is exited.
func (s *BaseObjCListener) ExitProperty_attribute(ctx *Property_attributeContext) {}

// EnterClass_name is called when production class_name is entered.
func (s *BaseObjCListener) EnterClass_name(ctx *Class_nameContext) {}

// ExitClass_name is called when production class_name is exited.
func (s *BaseObjCListener) ExitClass_name(ctx *Class_nameContext) {}

// EnterSuperclass_name is called when production superclass_name is entered.
func (s *BaseObjCListener) EnterSuperclass_name(ctx *Superclass_nameContext) {}

// ExitSuperclass_name is called when production superclass_name is exited.
func (s *BaseObjCListener) ExitSuperclass_name(ctx *Superclass_nameContext) {}

// EnterCategory_name is called when production category_name is entered.
func (s *BaseObjCListener) EnterCategory_name(ctx *Category_nameContext) {}

// ExitCategory_name is called when production category_name is exited.
func (s *BaseObjCListener) ExitCategory_name(ctx *Category_nameContext) {}

// EnterProtocol_name is called when production protocol_name is entered.
func (s *BaseObjCListener) EnterProtocol_name(ctx *Protocol_nameContext) {}

// ExitProtocol_name is called when production protocol_name is exited.
func (s *BaseObjCListener) ExitProtocol_name(ctx *Protocol_nameContext) {}

// EnterInstance_variables is called when production instance_variables is entered.
func (s *BaseObjCListener) EnterInstance_variables(ctx *Instance_variablesContext) {}

// ExitInstance_variables is called when production instance_variables is exited.
func (s *BaseObjCListener) ExitInstance_variables(ctx *Instance_variablesContext) {}

// EnterVisibility_specification is called when production visibility_specification is entered.
func (s *BaseObjCListener) EnterVisibility_specification(ctx *Visibility_specificationContext) {}

// ExitVisibility_specification is called when production visibility_specification is exited.
func (s *BaseObjCListener) ExitVisibility_specification(ctx *Visibility_specificationContext) {}

// EnterInterface_declaration_list is called when production interface_declaration_list is entered.
func (s *BaseObjCListener) EnterInterface_declaration_list(ctx *Interface_declaration_listContext) {}

// ExitInterface_declaration_list is called when production interface_declaration_list is exited.
func (s *BaseObjCListener) ExitInterface_declaration_list(ctx *Interface_declaration_listContext) {}

// EnterClass_method_declaration is called when production class_method_declaration is entered.
func (s *BaseObjCListener) EnterClass_method_declaration(ctx *Class_method_declarationContext) {}

// ExitClass_method_declaration is called when production class_method_declaration is exited.
func (s *BaseObjCListener) ExitClass_method_declaration(ctx *Class_method_declarationContext) {}

// EnterInstance_method_declaration is called when production instance_method_declaration is entered.
func (s *BaseObjCListener) EnterInstance_method_declaration(ctx *Instance_method_declarationContext) {}

// ExitInstance_method_declaration is called when production instance_method_declaration is exited.
func (s *BaseObjCListener) ExitInstance_method_declaration(ctx *Instance_method_declarationContext) {}

// EnterMethod_declaration is called when production method_declaration is entered.
func (s *BaseObjCListener) EnterMethod_declaration(ctx *Method_declarationContext) {}

// ExitMethod_declaration is called when production method_declaration is exited.
func (s *BaseObjCListener) ExitMethod_declaration(ctx *Method_declarationContext) {}

// EnterImplementation_definition_list is called when production implementation_definition_list is entered.
func (s *BaseObjCListener) EnterImplementation_definition_list(ctx *Implementation_definition_listContext) {}

// ExitImplementation_definition_list is called when production implementation_definition_list is exited.
func (s *BaseObjCListener) ExitImplementation_definition_list(ctx *Implementation_definition_listContext) {}

// EnterClass_method_definition is called when production class_method_definition is entered.
func (s *BaseObjCListener) EnterClass_method_definition(ctx *Class_method_definitionContext) {}

// ExitClass_method_definition is called when production class_method_definition is exited.
func (s *BaseObjCListener) ExitClass_method_definition(ctx *Class_method_definitionContext) {}

// EnterInstance_method_definition is called when production instance_method_definition is entered.
func (s *BaseObjCListener) EnterInstance_method_definition(ctx *Instance_method_definitionContext) {}

// ExitInstance_method_definition is called when production instance_method_definition is exited.
func (s *BaseObjCListener) ExitInstance_method_definition(ctx *Instance_method_definitionContext) {}

// EnterMethod_definition is called when production method_definition is entered.
func (s *BaseObjCListener) EnterMethod_definition(ctx *Method_definitionContext) {}

// ExitMethod_definition is called when production method_definition is exited.
func (s *BaseObjCListener) ExitMethod_definition(ctx *Method_definitionContext) {}

// EnterMethod_selector is called when production method_selector is entered.
func (s *BaseObjCListener) EnterMethod_selector(ctx *Method_selectorContext) {}

// ExitMethod_selector is called when production method_selector is exited.
func (s *BaseObjCListener) ExitMethod_selector(ctx *Method_selectorContext) {}

// EnterKeyword_declarator is called when production keyword_declarator is entered.
func (s *BaseObjCListener) EnterKeyword_declarator(ctx *Keyword_declaratorContext) {}

// ExitKeyword_declarator is called when production keyword_declarator is exited.
func (s *BaseObjCListener) ExitKeyword_declarator(ctx *Keyword_declaratorContext) {}

// EnterSelector is called when production selector is entered.
func (s *BaseObjCListener) EnterSelector(ctx *SelectorContext) {}

// ExitSelector is called when production selector is exited.
func (s *BaseObjCListener) ExitSelector(ctx *SelectorContext) {}

// EnterMethod_type is called when production method_type is entered.
func (s *BaseObjCListener) EnterMethod_type(ctx *Method_typeContext) {}

// ExitMethod_type is called when production method_type is exited.
func (s *BaseObjCListener) ExitMethod_type(ctx *Method_typeContext) {}

// EnterProperty_implementation is called when production property_implementation is entered.
func (s *BaseObjCListener) EnterProperty_implementation(ctx *Property_implementationContext) {}

// ExitProperty_implementation is called when production property_implementation is exited.
func (s *BaseObjCListener) ExitProperty_implementation(ctx *Property_implementationContext) {}

// EnterProperty_synthesize_list is called when production property_synthesize_list is entered.
func (s *BaseObjCListener) EnterProperty_synthesize_list(ctx *Property_synthesize_listContext) {}

// ExitProperty_synthesize_list is called when production property_synthesize_list is exited.
func (s *BaseObjCListener) ExitProperty_synthesize_list(ctx *Property_synthesize_listContext) {}

// EnterProperty_synthesize_item is called when production property_synthesize_item is entered.
func (s *BaseObjCListener) EnterProperty_synthesize_item(ctx *Property_synthesize_itemContext) {}

// ExitProperty_synthesize_item is called when production property_synthesize_item is exited.
func (s *BaseObjCListener) ExitProperty_synthesize_item(ctx *Property_synthesize_itemContext) {}

// EnterBlock_type is called when production block_type is entered.
func (s *BaseObjCListener) EnterBlock_type(ctx *Block_typeContext) {}

// ExitBlock_type is called when production block_type is exited.
func (s *BaseObjCListener) ExitBlock_type(ctx *Block_typeContext) {}

// EnterType_specifier is called when production type_specifier is entered.
func (s *BaseObjCListener) EnterType_specifier(ctx *Type_specifierContext) {}

// ExitType_specifier is called when production type_specifier is exited.
func (s *BaseObjCListener) ExitType_specifier(ctx *Type_specifierContext) {}

// EnterType_qualifier is called when production type_qualifier is entered.
func (s *BaseObjCListener) EnterType_qualifier(ctx *Type_qualifierContext) {}

// ExitType_qualifier is called when production type_qualifier is exited.
func (s *BaseObjCListener) ExitType_qualifier(ctx *Type_qualifierContext) {}

// EnterProtocol_qualifier is called when production protocol_qualifier is entered.
func (s *BaseObjCListener) EnterProtocol_qualifier(ctx *Protocol_qualifierContext) {}

// ExitProtocol_qualifier is called when production protocol_qualifier is exited.
func (s *BaseObjCListener) ExitProtocol_qualifier(ctx *Protocol_qualifierContext) {}

// EnterPrimary_expression is called when production primary_expression is entered.
func (s *BaseObjCListener) EnterPrimary_expression(ctx *Primary_expressionContext) {}

// ExitPrimary_expression is called when production primary_expression is exited.
func (s *BaseObjCListener) ExitPrimary_expression(ctx *Primary_expressionContext) {}

// EnterDictionary_pair is called when production dictionary_pair is entered.
func (s *BaseObjCListener) EnterDictionary_pair(ctx *Dictionary_pairContext) {}

// ExitDictionary_pair is called when production dictionary_pair is exited.
func (s *BaseObjCListener) ExitDictionary_pair(ctx *Dictionary_pairContext) {}

// EnterDictionary_expression is called when production dictionary_expression is entered.
func (s *BaseObjCListener) EnterDictionary_expression(ctx *Dictionary_expressionContext) {}

// ExitDictionary_expression is called when production dictionary_expression is exited.
func (s *BaseObjCListener) ExitDictionary_expression(ctx *Dictionary_expressionContext) {}

// EnterArray_expression is called when production array_expression is entered.
func (s *BaseObjCListener) EnterArray_expression(ctx *Array_expressionContext) {}

// ExitArray_expression is called when production array_expression is exited.
func (s *BaseObjCListener) ExitArray_expression(ctx *Array_expressionContext) {}

// EnterBox_expression is called when production box_expression is entered.
func (s *BaseObjCListener) EnterBox_expression(ctx *Box_expressionContext) {}

// ExitBox_expression is called when production box_expression is exited.
func (s *BaseObjCListener) ExitBox_expression(ctx *Box_expressionContext) {}

// EnterBlock_parameters is called when production block_parameters is entered.
func (s *BaseObjCListener) EnterBlock_parameters(ctx *Block_parametersContext) {}

// ExitBlock_parameters is called when production block_parameters is exited.
func (s *BaseObjCListener) ExitBlock_parameters(ctx *Block_parametersContext) {}

// EnterBlock_expression is called when production block_expression is entered.
func (s *BaseObjCListener) EnterBlock_expression(ctx *Block_expressionContext) {}

// ExitBlock_expression is called when production block_expression is exited.
func (s *BaseObjCListener) ExitBlock_expression(ctx *Block_expressionContext) {}

// EnterMessage_expression is called when production message_expression is entered.
func (s *BaseObjCListener) EnterMessage_expression(ctx *Message_expressionContext) {}

// ExitMessage_expression is called when production message_expression is exited.
func (s *BaseObjCListener) ExitMessage_expression(ctx *Message_expressionContext) {}

// EnterReceiver is called when production receiver is entered.
func (s *BaseObjCListener) EnterReceiver(ctx *ReceiverContext) {}

// ExitReceiver is called when production receiver is exited.
func (s *BaseObjCListener) ExitReceiver(ctx *ReceiverContext) {}

// EnterMessage_selector is called when production message_selector is entered.
func (s *BaseObjCListener) EnterMessage_selector(ctx *Message_selectorContext) {}

// ExitMessage_selector is called when production message_selector is exited.
func (s *BaseObjCListener) ExitMessage_selector(ctx *Message_selectorContext) {}

// EnterKeyword_argument is called when production keyword_argument is entered.
func (s *BaseObjCListener) EnterKeyword_argument(ctx *Keyword_argumentContext) {}

// ExitKeyword_argument is called when production keyword_argument is exited.
func (s *BaseObjCListener) ExitKeyword_argument(ctx *Keyword_argumentContext) {}

// EnterSelector_expression is called when production selector_expression is entered.
func (s *BaseObjCListener) EnterSelector_expression(ctx *Selector_expressionContext) {}

// ExitSelector_expression is called when production selector_expression is exited.
func (s *BaseObjCListener) ExitSelector_expression(ctx *Selector_expressionContext) {}

// EnterSelector_name is called when production selector_name is entered.
func (s *BaseObjCListener) EnterSelector_name(ctx *Selector_nameContext) {}

// ExitSelector_name is called when production selector_name is exited.
func (s *BaseObjCListener) ExitSelector_name(ctx *Selector_nameContext) {}

// EnterProtocol_expression is called when production protocol_expression is entered.
func (s *BaseObjCListener) EnterProtocol_expression(ctx *Protocol_expressionContext) {}

// ExitProtocol_expression is called when production protocol_expression is exited.
func (s *BaseObjCListener) ExitProtocol_expression(ctx *Protocol_expressionContext) {}

// EnterEncode_expression is called when production encode_expression is entered.
func (s *BaseObjCListener) EnterEncode_expression(ctx *Encode_expressionContext) {}

// ExitEncode_expression is called when production encode_expression is exited.
func (s *BaseObjCListener) ExitEncode_expression(ctx *Encode_expressionContext) {}

// EnterType_variable_declarator is called when production type_variable_declarator is entered.
func (s *BaseObjCListener) EnterType_variable_declarator(ctx *Type_variable_declaratorContext) {}

// ExitType_variable_declarator is called when production type_variable_declarator is exited.
func (s *BaseObjCListener) ExitType_variable_declarator(ctx *Type_variable_declaratorContext) {}

// EnterTry_statement is called when production try_statement is entered.
func (s *BaseObjCListener) EnterTry_statement(ctx *Try_statementContext) {}

// ExitTry_statement is called when production try_statement is exited.
func (s *BaseObjCListener) ExitTry_statement(ctx *Try_statementContext) {}

// EnterCatch_statement is called when production catch_statement is entered.
func (s *BaseObjCListener) EnterCatch_statement(ctx *Catch_statementContext) {}

// ExitCatch_statement is called when production catch_statement is exited.
func (s *BaseObjCListener) ExitCatch_statement(ctx *Catch_statementContext) {}

// EnterFinally_statement is called when production finally_statement is entered.
func (s *BaseObjCListener) EnterFinally_statement(ctx *Finally_statementContext) {}

// ExitFinally_statement is called when production finally_statement is exited.
func (s *BaseObjCListener) ExitFinally_statement(ctx *Finally_statementContext) {}

// EnterThrow_statement is called when production throw_statement is entered.
func (s *BaseObjCListener) EnterThrow_statement(ctx *Throw_statementContext) {}

// ExitThrow_statement is called when production throw_statement is exited.
func (s *BaseObjCListener) ExitThrow_statement(ctx *Throw_statementContext) {}

// EnterTry_block is called when production try_block is entered.
func (s *BaseObjCListener) EnterTry_block(ctx *Try_blockContext) {}

// ExitTry_block is called when production try_block is exited.
func (s *BaseObjCListener) ExitTry_block(ctx *Try_blockContext) {}

// EnterSynchronized_statement is called when production synchronized_statement is entered.
func (s *BaseObjCListener) EnterSynchronized_statement(ctx *Synchronized_statementContext) {}

// ExitSynchronized_statement is called when production synchronized_statement is exited.
func (s *BaseObjCListener) ExitSynchronized_statement(ctx *Synchronized_statementContext) {}

// EnterAutorelease_statement is called when production autorelease_statement is entered.
func (s *BaseObjCListener) EnterAutorelease_statement(ctx *Autorelease_statementContext) {}

// ExitAutorelease_statement is called when production autorelease_statement is exited.
func (s *BaseObjCListener) ExitAutorelease_statement(ctx *Autorelease_statementContext) {}

// EnterFunction_definition is called when production function_definition is entered.
func (s *BaseObjCListener) EnterFunction_definition(ctx *Function_definitionContext) {}

// ExitFunction_definition is called when production function_definition is exited.
func (s *BaseObjCListener) ExitFunction_definition(ctx *Function_definitionContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseObjCListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseObjCListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterDeclaration_specifiers is called when production declaration_specifiers is entered.
func (s *BaseObjCListener) EnterDeclaration_specifiers(ctx *Declaration_specifiersContext) {}

// ExitDeclaration_specifiers is called when production declaration_specifiers is exited.
func (s *BaseObjCListener) ExitDeclaration_specifiers(ctx *Declaration_specifiersContext) {}

// EnterArc_behaviour_specifier is called when production arc_behaviour_specifier is entered.
func (s *BaseObjCListener) EnterArc_behaviour_specifier(ctx *Arc_behaviour_specifierContext) {}

// ExitArc_behaviour_specifier is called when production arc_behaviour_specifier is exited.
func (s *BaseObjCListener) ExitArc_behaviour_specifier(ctx *Arc_behaviour_specifierContext) {}

// EnterStorage_class_specifier is called when production storage_class_specifier is entered.
func (s *BaseObjCListener) EnterStorage_class_specifier(ctx *Storage_class_specifierContext) {}

// ExitStorage_class_specifier is called when production storage_class_specifier is exited.
func (s *BaseObjCListener) ExitStorage_class_specifier(ctx *Storage_class_specifierContext) {}

// EnterInit_declarator_list is called when production init_declarator_list is entered.
func (s *BaseObjCListener) EnterInit_declarator_list(ctx *Init_declarator_listContext) {}

// ExitInit_declarator_list is called when production init_declarator_list is exited.
func (s *BaseObjCListener) ExitInit_declarator_list(ctx *Init_declarator_listContext) {}

// EnterInit_declarator is called when production init_declarator is entered.
func (s *BaseObjCListener) EnterInit_declarator(ctx *Init_declaratorContext) {}

// ExitInit_declarator is called when production init_declarator is exited.
func (s *BaseObjCListener) ExitInit_declarator(ctx *Init_declaratorContext) {}

// EnterStruct_or_union_specifier is called when production struct_or_union_specifier is entered.
func (s *BaseObjCListener) EnterStruct_or_union_specifier(ctx *Struct_or_union_specifierContext) {}

// ExitStruct_or_union_specifier is called when production struct_or_union_specifier is exited.
func (s *BaseObjCListener) ExitStruct_or_union_specifier(ctx *Struct_or_union_specifierContext) {}

// EnterStruct_declaration is called when production struct_declaration is entered.
func (s *BaseObjCListener) EnterStruct_declaration(ctx *Struct_declarationContext) {}

// ExitStruct_declaration is called when production struct_declaration is exited.
func (s *BaseObjCListener) ExitStruct_declaration(ctx *Struct_declarationContext) {}

// EnterSpecifier_qualifier_list is called when production specifier_qualifier_list is entered.
func (s *BaseObjCListener) EnterSpecifier_qualifier_list(ctx *Specifier_qualifier_listContext) {}

// ExitSpecifier_qualifier_list is called when production specifier_qualifier_list is exited.
func (s *BaseObjCListener) ExitSpecifier_qualifier_list(ctx *Specifier_qualifier_listContext) {}

// EnterStruct_declarator_list is called when production struct_declarator_list is entered.
func (s *BaseObjCListener) EnterStruct_declarator_list(ctx *Struct_declarator_listContext) {}

// ExitStruct_declarator_list is called when production struct_declarator_list is exited.
func (s *BaseObjCListener) ExitStruct_declarator_list(ctx *Struct_declarator_listContext) {}

// EnterStruct_declarator is called when production struct_declarator is entered.
func (s *BaseObjCListener) EnterStruct_declarator(ctx *Struct_declaratorContext) {}

// ExitStruct_declarator is called when production struct_declarator is exited.
func (s *BaseObjCListener) ExitStruct_declarator(ctx *Struct_declaratorContext) {}

// EnterEnum_specifier is called when production enum_specifier is entered.
func (s *BaseObjCListener) EnterEnum_specifier(ctx *Enum_specifierContext) {}

// ExitEnum_specifier is called when production enum_specifier is exited.
func (s *BaseObjCListener) ExitEnum_specifier(ctx *Enum_specifierContext) {}

// EnterEnumerator_list is called when production enumerator_list is entered.
func (s *BaseObjCListener) EnterEnumerator_list(ctx *Enumerator_listContext) {}

// ExitEnumerator_list is called when production enumerator_list is exited.
func (s *BaseObjCListener) ExitEnumerator_list(ctx *Enumerator_listContext) {}

// EnterEnumerator is called when production enumerator is entered.
func (s *BaseObjCListener) EnterEnumerator(ctx *EnumeratorContext) {}

// ExitEnumerator is called when production enumerator is exited.
func (s *BaseObjCListener) ExitEnumerator(ctx *EnumeratorContext) {}

// EnterPointer is called when production pointer is entered.
func (s *BaseObjCListener) EnterPointer(ctx *PointerContext) {}

// ExitPointer is called when production pointer is exited.
func (s *BaseObjCListener) ExitPointer(ctx *PointerContext) {}

// EnterDeclarator is called when production declarator is entered.
func (s *BaseObjCListener) EnterDeclarator(ctx *DeclaratorContext) {}

// ExitDeclarator is called when production declarator is exited.
func (s *BaseObjCListener) ExitDeclarator(ctx *DeclaratorContext) {}

// EnterDirect_declarator is called when production direct_declarator is entered.
func (s *BaseObjCListener) EnterDirect_declarator(ctx *Direct_declaratorContext) {}

// ExitDirect_declarator is called when production direct_declarator is exited.
func (s *BaseObjCListener) ExitDirect_declarator(ctx *Direct_declaratorContext) {}

// EnterDeclarator_suffix is called when production declarator_suffix is entered.
func (s *BaseObjCListener) EnterDeclarator_suffix(ctx *Declarator_suffixContext) {}

// ExitDeclarator_suffix is called when production declarator_suffix is exited.
func (s *BaseObjCListener) ExitDeclarator_suffix(ctx *Declarator_suffixContext) {}

// EnterParameter_list is called when production parameter_list is entered.
func (s *BaseObjCListener) EnterParameter_list(ctx *Parameter_listContext) {}

// ExitParameter_list is called when production parameter_list is exited.
func (s *BaseObjCListener) ExitParameter_list(ctx *Parameter_listContext) {}

// EnterParameter_declaration is called when production parameter_declaration is entered.
func (s *BaseObjCListener) EnterParameter_declaration(ctx *Parameter_declarationContext) {}

// ExitParameter_declaration is called when production parameter_declaration is exited.
func (s *BaseObjCListener) ExitParameter_declaration(ctx *Parameter_declarationContext) {}

// EnterInitializer is called when production initializer is entered.
func (s *BaseObjCListener) EnterInitializer(ctx *InitializerContext) {}

// ExitInitializer is called when production initializer is exited.
func (s *BaseObjCListener) ExitInitializer(ctx *InitializerContext) {}

// EnterType_name is called when production type_name is entered.
func (s *BaseObjCListener) EnterType_name(ctx *Type_nameContext) {}

// ExitType_name is called when production type_name is exited.
func (s *BaseObjCListener) ExitType_name(ctx *Type_nameContext) {}

// EnterAbstract_declarator is called when production abstract_declarator is entered.
func (s *BaseObjCListener) EnterAbstract_declarator(ctx *Abstract_declaratorContext) {}

// ExitAbstract_declarator is called when production abstract_declarator is exited.
func (s *BaseObjCListener) ExitAbstract_declarator(ctx *Abstract_declaratorContext) {}

// EnterAbstract_declarator_suffix is called when production abstract_declarator_suffix is entered.
func (s *BaseObjCListener) EnterAbstract_declarator_suffix(ctx *Abstract_declarator_suffixContext) {}

// ExitAbstract_declarator_suffix is called when production abstract_declarator_suffix is exited.
func (s *BaseObjCListener) ExitAbstract_declarator_suffix(ctx *Abstract_declarator_suffixContext) {}

// EnterParameter_declaration_list is called when production parameter_declaration_list is entered.
func (s *BaseObjCListener) EnterParameter_declaration_list(ctx *Parameter_declaration_listContext) {}

// ExitParameter_declaration_list is called when production parameter_declaration_list is exited.
func (s *BaseObjCListener) ExitParameter_declaration_list(ctx *Parameter_declaration_listContext) {}

// EnterStatement_list is called when production statement_list is entered.
func (s *BaseObjCListener) EnterStatement_list(ctx *Statement_listContext) {}

// ExitStatement_list is called when production statement_list is exited.
func (s *BaseObjCListener) ExitStatement_list(ctx *Statement_listContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseObjCListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseObjCListener) ExitStatement(ctx *StatementContext) {}

// EnterLabeled_statement is called when production labeled_statement is entered.
func (s *BaseObjCListener) EnterLabeled_statement(ctx *Labeled_statementContext) {}

// ExitLabeled_statement is called when production labeled_statement is exited.
func (s *BaseObjCListener) ExitLabeled_statement(ctx *Labeled_statementContext) {}

// EnterCompound_statement is called when production compound_statement is entered.
func (s *BaseObjCListener) EnterCompound_statement(ctx *Compound_statementContext) {}

// ExitCompound_statement is called when production compound_statement is exited.
func (s *BaseObjCListener) ExitCompound_statement(ctx *Compound_statementContext) {}

// EnterSelection_statement is called when production selection_statement is entered.
func (s *BaseObjCListener) EnterSelection_statement(ctx *Selection_statementContext) {}

// ExitSelection_statement is called when production selection_statement is exited.
func (s *BaseObjCListener) ExitSelection_statement(ctx *Selection_statementContext) {}

// EnterFor_in_statement is called when production for_in_statement is entered.
func (s *BaseObjCListener) EnterFor_in_statement(ctx *For_in_statementContext) {}

// ExitFor_in_statement is called when production for_in_statement is exited.
func (s *BaseObjCListener) ExitFor_in_statement(ctx *For_in_statementContext) {}

// EnterFor_statement is called when production for_statement is entered.
func (s *BaseObjCListener) EnterFor_statement(ctx *For_statementContext) {}

// ExitFor_statement is called when production for_statement is exited.
func (s *BaseObjCListener) ExitFor_statement(ctx *For_statementContext) {}

// EnterWhile_statement is called when production while_statement is entered.
func (s *BaseObjCListener) EnterWhile_statement(ctx *While_statementContext) {}

// ExitWhile_statement is called when production while_statement is exited.
func (s *BaseObjCListener) ExitWhile_statement(ctx *While_statementContext) {}

// EnterDo_statement is called when production do_statement is entered.
func (s *BaseObjCListener) EnterDo_statement(ctx *Do_statementContext) {}

// ExitDo_statement is called when production do_statement is exited.
func (s *BaseObjCListener) ExitDo_statement(ctx *Do_statementContext) {}

// EnterIteration_statement is called when production iteration_statement is entered.
func (s *BaseObjCListener) EnterIteration_statement(ctx *Iteration_statementContext) {}

// ExitIteration_statement is called when production iteration_statement is exited.
func (s *BaseObjCListener) ExitIteration_statement(ctx *Iteration_statementContext) {}

// EnterJump_statement is called when production jump_statement is entered.
func (s *BaseObjCListener) EnterJump_statement(ctx *Jump_statementContext) {}

// ExitJump_statement is called when production jump_statement is exited.
func (s *BaseObjCListener) ExitJump_statement(ctx *Jump_statementContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseObjCListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseObjCListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAssignment_expression is called when production assignment_expression is entered.
func (s *BaseObjCListener) EnterAssignment_expression(ctx *Assignment_expressionContext) {}

// ExitAssignment_expression is called when production assignment_expression is exited.
func (s *BaseObjCListener) ExitAssignment_expression(ctx *Assignment_expressionContext) {}

// EnterAssignment_operator is called when production assignment_operator is entered.
func (s *BaseObjCListener) EnterAssignment_operator(ctx *Assignment_operatorContext) {}

// ExitAssignment_operator is called when production assignment_operator is exited.
func (s *BaseObjCListener) ExitAssignment_operator(ctx *Assignment_operatorContext) {}

// EnterConditional_expression is called when production conditional_expression is entered.
func (s *BaseObjCListener) EnterConditional_expression(ctx *Conditional_expressionContext) {}

// ExitConditional_expression is called when production conditional_expression is exited.
func (s *BaseObjCListener) ExitConditional_expression(ctx *Conditional_expressionContext) {}

// EnterConstant_expression is called when production constant_expression is entered.
func (s *BaseObjCListener) EnterConstant_expression(ctx *Constant_expressionContext) {}

// ExitConstant_expression is called when production constant_expression is exited.
func (s *BaseObjCListener) ExitConstant_expression(ctx *Constant_expressionContext) {}

// EnterLogical_or_expression is called when production logical_or_expression is entered.
func (s *BaseObjCListener) EnterLogical_or_expression(ctx *Logical_or_expressionContext) {}

// ExitLogical_or_expression is called when production logical_or_expression is exited.
func (s *BaseObjCListener) ExitLogical_or_expression(ctx *Logical_or_expressionContext) {}

// EnterLogical_and_expression is called when production logical_and_expression is entered.
func (s *BaseObjCListener) EnterLogical_and_expression(ctx *Logical_and_expressionContext) {}

// ExitLogical_and_expression is called when production logical_and_expression is exited.
func (s *BaseObjCListener) ExitLogical_and_expression(ctx *Logical_and_expressionContext) {}

// EnterInclusive_or_expression is called when production inclusive_or_expression is entered.
func (s *BaseObjCListener) EnterInclusive_or_expression(ctx *Inclusive_or_expressionContext) {}

// ExitInclusive_or_expression is called when production inclusive_or_expression is exited.
func (s *BaseObjCListener) ExitInclusive_or_expression(ctx *Inclusive_or_expressionContext) {}

// EnterExclusive_or_expression is called when production exclusive_or_expression is entered.
func (s *BaseObjCListener) EnterExclusive_or_expression(ctx *Exclusive_or_expressionContext) {}

// ExitExclusive_or_expression is called when production exclusive_or_expression is exited.
func (s *BaseObjCListener) ExitExclusive_or_expression(ctx *Exclusive_or_expressionContext) {}

// EnterAnd_expression is called when production and_expression is entered.
func (s *BaseObjCListener) EnterAnd_expression(ctx *And_expressionContext) {}

// ExitAnd_expression is called when production and_expression is exited.
func (s *BaseObjCListener) ExitAnd_expression(ctx *And_expressionContext) {}

// EnterEquality_expression is called when production equality_expression is entered.
func (s *BaseObjCListener) EnterEquality_expression(ctx *Equality_expressionContext) {}

// ExitEquality_expression is called when production equality_expression is exited.
func (s *BaseObjCListener) ExitEquality_expression(ctx *Equality_expressionContext) {}

// EnterRelational_expression is called when production relational_expression is entered.
func (s *BaseObjCListener) EnterRelational_expression(ctx *Relational_expressionContext) {}

// ExitRelational_expression is called when production relational_expression is exited.
func (s *BaseObjCListener) ExitRelational_expression(ctx *Relational_expressionContext) {}

// EnterShift_expression is called when production shift_expression is entered.
func (s *BaseObjCListener) EnterShift_expression(ctx *Shift_expressionContext) {}

// ExitShift_expression is called when production shift_expression is exited.
func (s *BaseObjCListener) ExitShift_expression(ctx *Shift_expressionContext) {}

// EnterAdditive_expression is called when production additive_expression is entered.
func (s *BaseObjCListener) EnterAdditive_expression(ctx *Additive_expressionContext) {}

// ExitAdditive_expression is called when production additive_expression is exited.
func (s *BaseObjCListener) ExitAdditive_expression(ctx *Additive_expressionContext) {}

// EnterMultiplicative_expression is called when production multiplicative_expression is entered.
func (s *BaseObjCListener) EnterMultiplicative_expression(ctx *Multiplicative_expressionContext) {}

// ExitMultiplicative_expression is called when production multiplicative_expression is exited.
func (s *BaseObjCListener) ExitMultiplicative_expression(ctx *Multiplicative_expressionContext) {}

// EnterCast_expression is called when production cast_expression is entered.
func (s *BaseObjCListener) EnterCast_expression(ctx *Cast_expressionContext) {}

// ExitCast_expression is called when production cast_expression is exited.
func (s *BaseObjCListener) ExitCast_expression(ctx *Cast_expressionContext) {}

// EnterUnary_expression is called when production unary_expression is entered.
func (s *BaseObjCListener) EnterUnary_expression(ctx *Unary_expressionContext) {}

// ExitUnary_expression is called when production unary_expression is exited.
func (s *BaseObjCListener) ExitUnary_expression(ctx *Unary_expressionContext) {}

// EnterUnary_operator is called when production unary_operator is entered.
func (s *BaseObjCListener) EnterUnary_operator(ctx *Unary_operatorContext) {}

// ExitUnary_operator is called when production unary_operator is exited.
func (s *BaseObjCListener) ExitUnary_operator(ctx *Unary_operatorContext) {}

// EnterPostfix_expression is called when production postfix_expression is entered.
func (s *BaseObjCListener) EnterPostfix_expression(ctx *Postfix_expressionContext) {}

// ExitPostfix_expression is called when production postfix_expression is exited.
func (s *BaseObjCListener) ExitPostfix_expression(ctx *Postfix_expressionContext) {}

// EnterArgument_expression_list is called when production argument_expression_list is entered.
func (s *BaseObjCListener) EnterArgument_expression_list(ctx *Argument_expression_listContext) {}

// ExitArgument_expression_list is called when production argument_expression_list is exited.
func (s *BaseObjCListener) ExitArgument_expression_list(ctx *Argument_expression_listContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseObjCListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseObjCListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseObjCListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseObjCListener) ExitConstant(ctx *ConstantContext) {}
