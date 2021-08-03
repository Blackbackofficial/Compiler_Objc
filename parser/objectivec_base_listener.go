// Code generated from ObjectiveC.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // ObjectiveC

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseObjectiveCListener is a complete listener for a parse tree produced by ObjectiveCParser.
type BaseObjectiveCListener struct{}

var _ ObjectiveCListener = &BaseObjectiveCListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseObjectiveCListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseObjectiveCListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseObjectiveCListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseObjectiveCListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterTranslation_unit is called when production translation_unit is entered.
func (s *BaseObjectiveCListener) EnterTranslation_unit(ctx *Translation_unitContext) {}

// ExitTranslation_unit is called when production translation_unit is exited.
func (s *BaseObjectiveCListener) ExitTranslation_unit(ctx *Translation_unitContext) {}

// EnterExternal_declaration is called when production external_declaration is entered.
func (s *BaseObjectiveCListener) EnterExternal_declaration(ctx *External_declarationContext) {}

// ExitExternal_declaration is called when production external_declaration is exited.
func (s *BaseObjectiveCListener) ExitExternal_declaration(ctx *External_declarationContext) {}

// EnterPreprocessor_declaration is called when production preprocessor_declaration is entered.
func (s *BaseObjectiveCListener) EnterPreprocessor_declaration(ctx *Preprocessor_declarationContext) {}

// ExitPreprocessor_declaration is called when production preprocessor_declaration is exited.
func (s *BaseObjectiveCListener) ExitPreprocessor_declaration(ctx *Preprocessor_declarationContext) {}

// EnterFile_specification is called when production file_specification is entered.
func (s *BaseObjectiveCListener) EnterFile_specification(ctx *File_specificationContext) {}

// ExitFile_specification is called when production file_specification is exited.
func (s *BaseObjectiveCListener) ExitFile_specification(ctx *File_specificationContext) {}

// EnterMacro_specification is called when production macro_specification is entered.
func (s *BaseObjectiveCListener) EnterMacro_specification(ctx *Macro_specificationContext) {}

// ExitMacro_specification is called when production macro_specification is exited.
func (s *BaseObjectiveCListener) ExitMacro_specification(ctx *Macro_specificationContext) {}

// EnterClass_interface is called when production class_interface is entered.
func (s *BaseObjectiveCListener) EnterClass_interface(ctx *Class_interfaceContext) {}

// ExitClass_interface is called when production class_interface is exited.
func (s *BaseObjectiveCListener) ExitClass_interface(ctx *Class_interfaceContext) {}

// EnterCategory_interface is called when production category_interface is entered.
func (s *BaseObjectiveCListener) EnterCategory_interface(ctx *Category_interfaceContext) {}

// ExitCategory_interface is called when production category_interface is exited.
func (s *BaseObjectiveCListener) ExitCategory_interface(ctx *Category_interfaceContext) {}

// EnterClass_implementation is called when production class_implementation is entered.
func (s *BaseObjectiveCListener) EnterClass_implementation(ctx *Class_implementationContext) {}

// ExitClass_implementation is called when production class_implementation is exited.
func (s *BaseObjectiveCListener) ExitClass_implementation(ctx *Class_implementationContext) {}

// EnterCategory_implementation is called when production category_implementation is entered.
func (s *BaseObjectiveCListener) EnterCategory_implementation(ctx *Category_implementationContext) {}

// ExitCategory_implementation is called when production category_implementation is exited.
func (s *BaseObjectiveCListener) ExitCategory_implementation(ctx *Category_implementationContext) {}

// EnterProtocol_declaration is called when production protocol_declaration is entered.
func (s *BaseObjectiveCListener) EnterProtocol_declaration(ctx *Protocol_declarationContext) {}

// ExitProtocol_declaration is called when production protocol_declaration is exited.
func (s *BaseObjectiveCListener) ExitProtocol_declaration(ctx *Protocol_declarationContext) {}

// EnterProtocol_declaration_list is called when production protocol_declaration_list is entered.
func (s *BaseObjectiveCListener) EnterProtocol_declaration_list(ctx *Protocol_declaration_listContext) {}

// ExitProtocol_declaration_list is called when production protocol_declaration_list is exited.
func (s *BaseObjectiveCListener) ExitProtocol_declaration_list(ctx *Protocol_declaration_listContext) {}

// EnterClass_declaration_list is called when production class_declaration_list is entered.
func (s *BaseObjectiveCListener) EnterClass_declaration_list(ctx *Class_declaration_listContext) {}

// ExitClass_declaration_list is called when production class_declaration_list is exited.
func (s *BaseObjectiveCListener) ExitClass_declaration_list(ctx *Class_declaration_listContext) {}

// EnterClass_list is called when production class_list is entered.
func (s *BaseObjectiveCListener) EnterClass_list(ctx *Class_listContext) {}

// ExitClass_list is called when production class_list is exited.
func (s *BaseObjectiveCListener) ExitClass_list(ctx *Class_listContext) {}

// EnterProtocol_reference_list is called when production protocol_reference_list is entered.
func (s *BaseObjectiveCListener) EnterProtocol_reference_list(ctx *Protocol_reference_listContext) {}

// ExitProtocol_reference_list is called when production protocol_reference_list is exited.
func (s *BaseObjectiveCListener) ExitProtocol_reference_list(ctx *Protocol_reference_listContext) {}

// EnterProtocol_list is called when production protocol_list is entered.
func (s *BaseObjectiveCListener) EnterProtocol_list(ctx *Protocol_listContext) {}

// ExitProtocol_list is called when production protocol_list is exited.
func (s *BaseObjectiveCListener) ExitProtocol_list(ctx *Protocol_listContext) {}

// EnterClass_name is called when production class_name is entered.
func (s *BaseObjectiveCListener) EnterClass_name(ctx *Class_nameContext) {}

// ExitClass_name is called when production class_name is exited.
func (s *BaseObjectiveCListener) ExitClass_name(ctx *Class_nameContext) {}

// EnterSuperclass_name is called when production superclass_name is entered.
func (s *BaseObjectiveCListener) EnterSuperclass_name(ctx *Superclass_nameContext) {}

// ExitSuperclass_name is called when production superclass_name is exited.
func (s *BaseObjectiveCListener) ExitSuperclass_name(ctx *Superclass_nameContext) {}

// EnterCategory_name is called when production category_name is entered.
func (s *BaseObjectiveCListener) EnterCategory_name(ctx *Category_nameContext) {}

// ExitCategory_name is called when production category_name is exited.
func (s *BaseObjectiveCListener) ExitCategory_name(ctx *Category_nameContext) {}

// EnterProtocol_name is called when production protocol_name is entered.
func (s *BaseObjectiveCListener) EnterProtocol_name(ctx *Protocol_nameContext) {}

// ExitProtocol_name is called when production protocol_name is exited.
func (s *BaseObjectiveCListener) ExitProtocol_name(ctx *Protocol_nameContext) {}

// EnterInstance_variables is called when production instance_variables is entered.
func (s *BaseObjectiveCListener) EnterInstance_variables(ctx *Instance_variablesContext) {}

// ExitInstance_variables is called when production instance_variables is exited.
func (s *BaseObjectiveCListener) ExitInstance_variables(ctx *Instance_variablesContext) {}

// EnterInstance_variable_declaration is called when production instance_variable_declaration is entered.
func (s *BaseObjectiveCListener) EnterInstance_variable_declaration(ctx *Instance_variable_declarationContext) {}

// ExitInstance_variable_declaration is called when production instance_variable_declaration is exited.
func (s *BaseObjectiveCListener) ExitInstance_variable_declaration(ctx *Instance_variable_declarationContext) {}

// EnterVisibility_specification is called when production visibility_specification is entered.
func (s *BaseObjectiveCListener) EnterVisibility_specification(ctx *Visibility_specificationContext) {}

// ExitVisibility_specification is called when production visibility_specification is exited.
func (s *BaseObjectiveCListener) ExitVisibility_specification(ctx *Visibility_specificationContext) {}

// EnterInterface_declaration_list is called when production interface_declaration_list is entered.
func (s *BaseObjectiveCListener) EnterInterface_declaration_list(ctx *Interface_declaration_listContext) {}

// ExitInterface_declaration_list is called when production interface_declaration_list is exited.
func (s *BaseObjectiveCListener) ExitInterface_declaration_list(ctx *Interface_declaration_listContext) {}

// EnterClass_method_declaration is called when production class_method_declaration is entered.
func (s *BaseObjectiveCListener) EnterClass_method_declaration(ctx *Class_method_declarationContext) {}

// ExitClass_method_declaration is called when production class_method_declaration is exited.
func (s *BaseObjectiveCListener) ExitClass_method_declaration(ctx *Class_method_declarationContext) {}

// EnterInstance_method_declaration is called when production instance_method_declaration is entered.
func (s *BaseObjectiveCListener) EnterInstance_method_declaration(ctx *Instance_method_declarationContext) {}

// ExitInstance_method_declaration is called when production instance_method_declaration is exited.
func (s *BaseObjectiveCListener) ExitInstance_method_declaration(ctx *Instance_method_declarationContext) {}

// EnterMethod_declaration is called when production method_declaration is entered.
func (s *BaseObjectiveCListener) EnterMethod_declaration(ctx *Method_declarationContext) {}

// ExitMethod_declaration is called when production method_declaration is exited.
func (s *BaseObjectiveCListener) ExitMethod_declaration(ctx *Method_declarationContext) {}

// EnterImplementation_definition_list is called when production implementation_definition_list is entered.
func (s *BaseObjectiveCListener) EnterImplementation_definition_list(ctx *Implementation_definition_listContext) {}

// ExitImplementation_definition_list is called when production implementation_definition_list is exited.
func (s *BaseObjectiveCListener) ExitImplementation_definition_list(ctx *Implementation_definition_listContext) {}

// EnterClass_method_definition is called when production class_method_definition is entered.
func (s *BaseObjectiveCListener) EnterClass_method_definition(ctx *Class_method_definitionContext) {}

// ExitClass_method_definition is called when production class_method_definition is exited.
func (s *BaseObjectiveCListener) ExitClass_method_definition(ctx *Class_method_definitionContext) {}

// EnterInstance_method_definition is called when production instance_method_definition is entered.
func (s *BaseObjectiveCListener) EnterInstance_method_definition(ctx *Instance_method_definitionContext) {}

// ExitInstance_method_definition is called when production instance_method_definition is exited.
func (s *BaseObjectiveCListener) ExitInstance_method_definition(ctx *Instance_method_definitionContext) {}

// EnterMethod_definition is called when production method_definition is entered.
func (s *BaseObjectiveCListener) EnterMethod_definition(ctx *Method_definitionContext) {}

// ExitMethod_definition is called when production method_definition is exited.
func (s *BaseObjectiveCListener) ExitMethod_definition(ctx *Method_definitionContext) {}

// EnterMethod_selector is called when production method_selector is entered.
func (s *BaseObjectiveCListener) EnterMethod_selector(ctx *Method_selectorContext) {}

// ExitMethod_selector is called when production method_selector is exited.
func (s *BaseObjectiveCListener) ExitMethod_selector(ctx *Method_selectorContext) {}

// EnterKeyword_declarator is called when production keyword_declarator is entered.
func (s *BaseObjectiveCListener) EnterKeyword_declarator(ctx *Keyword_declaratorContext) {}

// ExitKeyword_declarator is called when production keyword_declarator is exited.
func (s *BaseObjectiveCListener) ExitKeyword_declarator(ctx *Keyword_declaratorContext) {}

// EnterSelector is called when production selector is entered.
func (s *BaseObjectiveCListener) EnterSelector(ctx *SelectorContext) {}

// ExitSelector is called when production selector is exited.
func (s *BaseObjectiveCListener) ExitSelector(ctx *SelectorContext) {}

// EnterMethod_type is called when production method_type is entered.
func (s *BaseObjectiveCListener) EnterMethod_type(ctx *Method_typeContext) {}

// ExitMethod_type is called when production method_type is exited.
func (s *BaseObjectiveCListener) ExitMethod_type(ctx *Method_typeContext) {}

// EnterType_specifier is called when production type_specifier is entered.
func (s *BaseObjectiveCListener) EnterType_specifier(ctx *Type_specifierContext) {}

// ExitType_specifier is called when production type_specifier is exited.
func (s *BaseObjectiveCListener) ExitType_specifier(ctx *Type_specifierContext) {}

// EnterType_qualifier is called when production type_qualifier is entered.
func (s *BaseObjectiveCListener) EnterType_qualifier(ctx *Type_qualifierContext) {}

// ExitType_qualifier is called when production type_qualifier is exited.
func (s *BaseObjectiveCListener) ExitType_qualifier(ctx *Type_qualifierContext) {}

// EnterProtocol_qualifier is called when production protocol_qualifier is entered.
func (s *BaseObjectiveCListener) EnterProtocol_qualifier(ctx *Protocol_qualifierContext) {}

// ExitProtocol_qualifier is called when production protocol_qualifier is exited.
func (s *BaseObjectiveCListener) ExitProtocol_qualifier(ctx *Protocol_qualifierContext) {}

// EnterPrimary_expression is called when production primary_expression is entered.
func (s *BaseObjectiveCListener) EnterPrimary_expression(ctx *Primary_expressionContext) {}

// ExitPrimary_expression is called when production primary_expression is exited.
func (s *BaseObjectiveCListener) ExitPrimary_expression(ctx *Primary_expressionContext) {}

// EnterMessage_expression is called when production message_expression is entered.
func (s *BaseObjectiveCListener) EnterMessage_expression(ctx *Message_expressionContext) {}

// ExitMessage_expression is called when production message_expression is exited.
func (s *BaseObjectiveCListener) ExitMessage_expression(ctx *Message_expressionContext) {}

// EnterReceiver is called when production receiver is entered.
func (s *BaseObjectiveCListener) EnterReceiver(ctx *ReceiverContext) {}

// ExitReceiver is called when production receiver is exited.
func (s *BaseObjectiveCListener) ExitReceiver(ctx *ReceiverContext) {}

// EnterMessage_selector is called when production message_selector is entered.
func (s *BaseObjectiveCListener) EnterMessage_selector(ctx *Message_selectorContext) {}

// ExitMessage_selector is called when production message_selector is exited.
func (s *BaseObjectiveCListener) ExitMessage_selector(ctx *Message_selectorContext) {}

// EnterKeyword_argument is called when production keyword_argument is entered.
func (s *BaseObjectiveCListener) EnterKeyword_argument(ctx *Keyword_argumentContext) {}

// ExitKeyword_argument is called when production keyword_argument is exited.
func (s *BaseObjectiveCListener) ExitKeyword_argument(ctx *Keyword_argumentContext) {}

// EnterSelector_expression is called when production selector_expression is entered.
func (s *BaseObjectiveCListener) EnterSelector_expression(ctx *Selector_expressionContext) {}

// ExitSelector_expression is called when production selector_expression is exited.
func (s *BaseObjectiveCListener) ExitSelector_expression(ctx *Selector_expressionContext) {}

// EnterSelector_name is called when production selector_name is entered.
func (s *BaseObjectiveCListener) EnterSelector_name(ctx *Selector_nameContext) {}

// ExitSelector_name is called when production selector_name is exited.
func (s *BaseObjectiveCListener) ExitSelector_name(ctx *Selector_nameContext) {}

// EnterProtocol_expression is called when production protocol_expression is entered.
func (s *BaseObjectiveCListener) EnterProtocol_expression(ctx *Protocol_expressionContext) {}

// ExitProtocol_expression is called when production protocol_expression is exited.
func (s *BaseObjectiveCListener) ExitProtocol_expression(ctx *Protocol_expressionContext) {}

// EnterEncode_expression is called when production encode_expression is entered.
func (s *BaseObjectiveCListener) EnterEncode_expression(ctx *Encode_expressionContext) {}

// ExitEncode_expression is called when production encode_expression is exited.
func (s *BaseObjectiveCListener) ExitEncode_expression(ctx *Encode_expressionContext) {}

// EnterException_declarator is called when production exception_declarator is entered.
func (s *BaseObjectiveCListener) EnterException_declarator(ctx *Exception_declaratorContext) {}

// ExitException_declarator is called when production exception_declarator is exited.
func (s *BaseObjectiveCListener) ExitException_declarator(ctx *Exception_declaratorContext) {}

// EnterTry_statement is called when production try_statement is entered.
func (s *BaseObjectiveCListener) EnterTry_statement(ctx *Try_statementContext) {}

// ExitTry_statement is called when production try_statement is exited.
func (s *BaseObjectiveCListener) ExitTry_statement(ctx *Try_statementContext) {}

// EnterCatch_statement is called when production catch_statement is entered.
func (s *BaseObjectiveCListener) EnterCatch_statement(ctx *Catch_statementContext) {}

// ExitCatch_statement is called when production catch_statement is exited.
func (s *BaseObjectiveCListener) ExitCatch_statement(ctx *Catch_statementContext) {}

// EnterFinally_statement is called when production finally_statement is entered.
func (s *BaseObjectiveCListener) EnterFinally_statement(ctx *Finally_statementContext) {}

// ExitFinally_statement is called when production finally_statement is exited.
func (s *BaseObjectiveCListener) ExitFinally_statement(ctx *Finally_statementContext) {}

// EnterThrow_statement is called when production throw_statement is entered.
func (s *BaseObjectiveCListener) EnterThrow_statement(ctx *Throw_statementContext) {}

// ExitThrow_statement is called when production throw_statement is exited.
func (s *BaseObjectiveCListener) ExitThrow_statement(ctx *Throw_statementContext) {}

// EnterTry_block is called when production try_block is entered.
func (s *BaseObjectiveCListener) EnterTry_block(ctx *Try_blockContext) {}

// ExitTry_block is called when production try_block is exited.
func (s *BaseObjectiveCListener) ExitTry_block(ctx *Try_blockContext) {}

// EnterSynchronized_statement is called when production synchronized_statement is entered.
func (s *BaseObjectiveCListener) EnterSynchronized_statement(ctx *Synchronized_statementContext) {}

// ExitSynchronized_statement is called when production synchronized_statement is exited.
func (s *BaseObjectiveCListener) ExitSynchronized_statement(ctx *Synchronized_statementContext) {}

// EnterFunction_definition is called when production function_definition is entered.
func (s *BaseObjectiveCListener) EnterFunction_definition(ctx *Function_definitionContext) {}

// ExitFunction_definition is called when production function_definition is exited.
func (s *BaseObjectiveCListener) ExitFunction_definition(ctx *Function_definitionContext) {}

// EnterDeclaration is called when production declaration is entered.
func (s *BaseObjectiveCListener) EnterDeclaration(ctx *DeclarationContext) {}

// ExitDeclaration is called when production declaration is exited.
func (s *BaseObjectiveCListener) ExitDeclaration(ctx *DeclarationContext) {}

// EnterDeclaration_specifiers is called when production declaration_specifiers is entered.
func (s *BaseObjectiveCListener) EnterDeclaration_specifiers(ctx *Declaration_specifiersContext) {}

// ExitDeclaration_specifiers is called when production declaration_specifiers is exited.
func (s *BaseObjectiveCListener) ExitDeclaration_specifiers(ctx *Declaration_specifiersContext) {}

// EnterStorage_class_specifier is called when production storage_class_specifier is entered.
func (s *BaseObjectiveCListener) EnterStorage_class_specifier(ctx *Storage_class_specifierContext) {}

// ExitStorage_class_specifier is called when production storage_class_specifier is exited.
func (s *BaseObjectiveCListener) ExitStorage_class_specifier(ctx *Storage_class_specifierContext) {}

// EnterInit_declarator_list is called when production init_declarator_list is entered.
func (s *BaseObjectiveCListener) EnterInit_declarator_list(ctx *Init_declarator_listContext) {}

// ExitInit_declarator_list is called when production init_declarator_list is exited.
func (s *BaseObjectiveCListener) ExitInit_declarator_list(ctx *Init_declarator_listContext) {}

// EnterInit_declarator is called when production init_declarator is entered.
func (s *BaseObjectiveCListener) EnterInit_declarator(ctx *Init_declaratorContext) {}

// ExitInit_declarator is called when production init_declarator is exited.
func (s *BaseObjectiveCListener) ExitInit_declarator(ctx *Init_declaratorContext) {}

// EnterStruct_or_union_specifier is called when production struct_or_union_specifier is entered.
func (s *BaseObjectiveCListener) EnterStruct_or_union_specifier(ctx *Struct_or_union_specifierContext) {}

// ExitStruct_or_union_specifier is called when production struct_or_union_specifier is exited.
func (s *BaseObjectiveCListener) ExitStruct_or_union_specifier(ctx *Struct_or_union_specifierContext) {}

// EnterStruct_declaration is called when production struct_declaration is entered.
func (s *BaseObjectiveCListener) EnterStruct_declaration(ctx *Struct_declarationContext) {}

// ExitStruct_declaration is called when production struct_declaration is exited.
func (s *BaseObjectiveCListener) ExitStruct_declaration(ctx *Struct_declarationContext) {}

// EnterSpecifier_qualifier_list is called when production specifier_qualifier_list is entered.
func (s *BaseObjectiveCListener) EnterSpecifier_qualifier_list(ctx *Specifier_qualifier_listContext) {}

// ExitSpecifier_qualifier_list is called when production specifier_qualifier_list is exited.
func (s *BaseObjectiveCListener) ExitSpecifier_qualifier_list(ctx *Specifier_qualifier_listContext) {}

// EnterStruct_declarator_list is called when production struct_declarator_list is entered.
func (s *BaseObjectiveCListener) EnterStruct_declarator_list(ctx *Struct_declarator_listContext) {}

// ExitStruct_declarator_list is called when production struct_declarator_list is exited.
func (s *BaseObjectiveCListener) ExitStruct_declarator_list(ctx *Struct_declarator_listContext) {}

// EnterStruct_declarator is called when production struct_declarator is entered.
func (s *BaseObjectiveCListener) EnterStruct_declarator(ctx *Struct_declaratorContext) {}

// ExitStruct_declarator is called when production struct_declarator is exited.
func (s *BaseObjectiveCListener) ExitStruct_declarator(ctx *Struct_declaratorContext) {}

// EnterEnum_specifier is called when production enum_specifier is entered.
func (s *BaseObjectiveCListener) EnterEnum_specifier(ctx *Enum_specifierContext) {}

// ExitEnum_specifier is called when production enum_specifier is exited.
func (s *BaseObjectiveCListener) ExitEnum_specifier(ctx *Enum_specifierContext) {}

// EnterEnumerator_list is called when production enumerator_list is entered.
func (s *BaseObjectiveCListener) EnterEnumerator_list(ctx *Enumerator_listContext) {}

// ExitEnumerator_list is called when production enumerator_list is exited.
func (s *BaseObjectiveCListener) ExitEnumerator_list(ctx *Enumerator_listContext) {}

// EnterEnumerator is called when production enumerator is entered.
func (s *BaseObjectiveCListener) EnterEnumerator(ctx *EnumeratorContext) {}

// ExitEnumerator is called when production enumerator is exited.
func (s *BaseObjectiveCListener) ExitEnumerator(ctx *EnumeratorContext) {}

// EnterDeclarator is called when production declarator is entered.
func (s *BaseObjectiveCListener) EnterDeclarator(ctx *DeclaratorContext) {}

// ExitDeclarator is called when production declarator is exited.
func (s *BaseObjectiveCListener) ExitDeclarator(ctx *DeclaratorContext) {}

// EnterDirect_declarator is called when production direct_declarator is entered.
func (s *BaseObjectiveCListener) EnterDirect_declarator(ctx *Direct_declaratorContext) {}

// ExitDirect_declarator is called when production direct_declarator is exited.
func (s *BaseObjectiveCListener) ExitDirect_declarator(ctx *Direct_declaratorContext) {}

// EnterDeclarator_suffix is called when production declarator_suffix is entered.
func (s *BaseObjectiveCListener) EnterDeclarator_suffix(ctx *Declarator_suffixContext) {}

// ExitDeclarator_suffix is called when production declarator_suffix is exited.
func (s *BaseObjectiveCListener) ExitDeclarator_suffix(ctx *Declarator_suffixContext) {}

// EnterParameter_list is called when production parameter_list is entered.
func (s *BaseObjectiveCListener) EnterParameter_list(ctx *Parameter_listContext) {}

// ExitParameter_list is called when production parameter_list is exited.
func (s *BaseObjectiveCListener) ExitParameter_list(ctx *Parameter_listContext) {}

// EnterParameter_declaration is called when production parameter_declaration is entered.
func (s *BaseObjectiveCListener) EnterParameter_declaration(ctx *Parameter_declarationContext) {}

// ExitParameter_declaration is called when production parameter_declaration is exited.
func (s *BaseObjectiveCListener) ExitParameter_declaration(ctx *Parameter_declarationContext) {}

// EnterInitializer is called when production initializer is entered.
func (s *BaseObjectiveCListener) EnterInitializer(ctx *InitializerContext) {}

// ExitInitializer is called when production initializer is exited.
func (s *BaseObjectiveCListener) ExitInitializer(ctx *InitializerContext) {}

// EnterType_name is called when production type_name is entered.
func (s *BaseObjectiveCListener) EnterType_name(ctx *Type_nameContext) {}

// ExitType_name is called when production type_name is exited.
func (s *BaseObjectiveCListener) ExitType_name(ctx *Type_nameContext) {}

// EnterAbstract_declarator is called when production abstract_declarator is entered.
func (s *BaseObjectiveCListener) EnterAbstract_declarator(ctx *Abstract_declaratorContext) {}

// ExitAbstract_declarator is called when production abstract_declarator is exited.
func (s *BaseObjectiveCListener) ExitAbstract_declarator(ctx *Abstract_declaratorContext) {}

// EnterAbstract_declarator_suffix is called when production abstract_declarator_suffix is entered.
func (s *BaseObjectiveCListener) EnterAbstract_declarator_suffix(ctx *Abstract_declarator_suffixContext) {}

// ExitAbstract_declarator_suffix is called when production abstract_declarator_suffix is exited.
func (s *BaseObjectiveCListener) ExitAbstract_declarator_suffix(ctx *Abstract_declarator_suffixContext) {}

// EnterParameter_declaration_list is called when production parameter_declaration_list is entered.
func (s *BaseObjectiveCListener) EnterParameter_declaration_list(ctx *Parameter_declaration_listContext) {}

// ExitParameter_declaration_list is called when production parameter_declaration_list is exited.
func (s *BaseObjectiveCListener) ExitParameter_declaration_list(ctx *Parameter_declaration_listContext) {}

// EnterStatement_list is called when production statement_list is entered.
func (s *BaseObjectiveCListener) EnterStatement_list(ctx *Statement_listContext) {}

// ExitStatement_list is called when production statement_list is exited.
func (s *BaseObjectiveCListener) ExitStatement_list(ctx *Statement_listContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseObjectiveCListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseObjectiveCListener) ExitStatement(ctx *StatementContext) {}

// EnterLabeled_statement is called when production labeled_statement is entered.
func (s *BaseObjectiveCListener) EnterLabeled_statement(ctx *Labeled_statementContext) {}

// ExitLabeled_statement is called when production labeled_statement is exited.
func (s *BaseObjectiveCListener) ExitLabeled_statement(ctx *Labeled_statementContext) {}

// EnterCompound_statement is called when production compound_statement is entered.
func (s *BaseObjectiveCListener) EnterCompound_statement(ctx *Compound_statementContext) {}

// ExitCompound_statement is called when production compound_statement is exited.
func (s *BaseObjectiveCListener) ExitCompound_statement(ctx *Compound_statementContext) {}

// EnterSelection_statement is called when production selection_statement is entered.
func (s *BaseObjectiveCListener) EnterSelection_statement(ctx *Selection_statementContext) {}

// ExitSelection_statement is called when production selection_statement is exited.
func (s *BaseObjectiveCListener) ExitSelection_statement(ctx *Selection_statementContext) {}

// EnterIteration_statement is called when production iteration_statement is entered.
func (s *BaseObjectiveCListener) EnterIteration_statement(ctx *Iteration_statementContext) {}

// ExitIteration_statement is called when production iteration_statement is exited.
func (s *BaseObjectiveCListener) ExitIteration_statement(ctx *Iteration_statementContext) {}

// EnterJump_statement is called when production jump_statement is entered.
func (s *BaseObjectiveCListener) EnterJump_statement(ctx *Jump_statementContext) {}

// ExitJump_statement is called when production jump_statement is exited.
func (s *BaseObjectiveCListener) ExitJump_statement(ctx *Jump_statementContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseObjectiveCListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseObjectiveCListener) ExitExpression(ctx *ExpressionContext) {}

// EnterAssignment_expression is called when production assignment_expression is entered.
func (s *BaseObjectiveCListener) EnterAssignment_expression(ctx *Assignment_expressionContext) {}

// ExitAssignment_expression is called when production assignment_expression is exited.
func (s *BaseObjectiveCListener) ExitAssignment_expression(ctx *Assignment_expressionContext) {}

// EnterAssignment_operator is called when production assignment_operator is entered.
func (s *BaseObjectiveCListener) EnterAssignment_operator(ctx *Assignment_operatorContext) {}

// ExitAssignment_operator is called when production assignment_operator is exited.
func (s *BaseObjectiveCListener) ExitAssignment_operator(ctx *Assignment_operatorContext) {}

// EnterConditional_expression is called when production conditional_expression is entered.
func (s *BaseObjectiveCListener) EnterConditional_expression(ctx *Conditional_expressionContext) {}

// ExitConditional_expression is called when production conditional_expression is exited.
func (s *BaseObjectiveCListener) ExitConditional_expression(ctx *Conditional_expressionContext) {}

// EnterConstant_expression is called when production constant_expression is entered.
func (s *BaseObjectiveCListener) EnterConstant_expression(ctx *Constant_expressionContext) {}

// ExitConstant_expression is called when production constant_expression is exited.
func (s *BaseObjectiveCListener) ExitConstant_expression(ctx *Constant_expressionContext) {}

// EnterLogical_or_expression is called when production logical_or_expression is entered.
func (s *BaseObjectiveCListener) EnterLogical_or_expression(ctx *Logical_or_expressionContext) {}

// ExitLogical_or_expression is called when production logical_or_expression is exited.
func (s *BaseObjectiveCListener) ExitLogical_or_expression(ctx *Logical_or_expressionContext) {}

// EnterLogical_and_expression is called when production logical_and_expression is entered.
func (s *BaseObjectiveCListener) EnterLogical_and_expression(ctx *Logical_and_expressionContext) {}

// ExitLogical_and_expression is called when production logical_and_expression is exited.
func (s *BaseObjectiveCListener) ExitLogical_and_expression(ctx *Logical_and_expressionContext) {}

// EnterInclusive_or_expression is called when production inclusive_or_expression is entered.
func (s *BaseObjectiveCListener) EnterInclusive_or_expression(ctx *Inclusive_or_expressionContext) {}

// ExitInclusive_or_expression is called when production inclusive_or_expression is exited.
func (s *BaseObjectiveCListener) ExitInclusive_or_expression(ctx *Inclusive_or_expressionContext) {}

// EnterExclusive_or_expression is called when production exclusive_or_expression is entered.
func (s *BaseObjectiveCListener) EnterExclusive_or_expression(ctx *Exclusive_or_expressionContext) {}

// ExitExclusive_or_expression is called when production exclusive_or_expression is exited.
func (s *BaseObjectiveCListener) ExitExclusive_or_expression(ctx *Exclusive_or_expressionContext) {}

// EnterAnd_expression is called when production and_expression is entered.
func (s *BaseObjectiveCListener) EnterAnd_expression(ctx *And_expressionContext) {}

// ExitAnd_expression is called when production and_expression is exited.
func (s *BaseObjectiveCListener) ExitAnd_expression(ctx *And_expressionContext) {}

// EnterEquality_expression is called when production equality_expression is entered.
func (s *BaseObjectiveCListener) EnterEquality_expression(ctx *Equality_expressionContext) {}

// ExitEquality_expression is called when production equality_expression is exited.
func (s *BaseObjectiveCListener) ExitEquality_expression(ctx *Equality_expressionContext) {}

// EnterRelational_expression is called when production relational_expression is entered.
func (s *BaseObjectiveCListener) EnterRelational_expression(ctx *Relational_expressionContext) {}

// ExitRelational_expression is called when production relational_expression is exited.
func (s *BaseObjectiveCListener) ExitRelational_expression(ctx *Relational_expressionContext) {}

// EnterShift_expression is called when production shift_expression is entered.
func (s *BaseObjectiveCListener) EnterShift_expression(ctx *Shift_expressionContext) {}

// ExitShift_expression is called when production shift_expression is exited.
func (s *BaseObjectiveCListener) ExitShift_expression(ctx *Shift_expressionContext) {}

// EnterAdditive_expression is called when production additive_expression is entered.
func (s *BaseObjectiveCListener) EnterAdditive_expression(ctx *Additive_expressionContext) {}

// ExitAdditive_expression is called when production additive_expression is exited.
func (s *BaseObjectiveCListener) ExitAdditive_expression(ctx *Additive_expressionContext) {}

// EnterMultiplicative_expression is called when production multiplicative_expression is entered.
func (s *BaseObjectiveCListener) EnterMultiplicative_expression(ctx *Multiplicative_expressionContext) {}

// ExitMultiplicative_expression is called when production multiplicative_expression is exited.
func (s *BaseObjectiveCListener) ExitMultiplicative_expression(ctx *Multiplicative_expressionContext) {}

// EnterCast_expression is called when production cast_expression is entered.
func (s *BaseObjectiveCListener) EnterCast_expression(ctx *Cast_expressionContext) {}

// ExitCast_expression is called when production cast_expression is exited.
func (s *BaseObjectiveCListener) ExitCast_expression(ctx *Cast_expressionContext) {}

// EnterUnary_expression is called when production unary_expression is entered.
func (s *BaseObjectiveCListener) EnterUnary_expression(ctx *Unary_expressionContext) {}

// ExitUnary_expression is called when production unary_expression is exited.
func (s *BaseObjectiveCListener) ExitUnary_expression(ctx *Unary_expressionContext) {}

// EnterUnary_operator is called when production unary_operator is entered.
func (s *BaseObjectiveCListener) EnterUnary_operator(ctx *Unary_operatorContext) {}

// ExitUnary_operator is called when production unary_operator is exited.
func (s *BaseObjectiveCListener) ExitUnary_operator(ctx *Unary_operatorContext) {}

// EnterPostfix_expression is called when production postfix_expression is entered.
func (s *BaseObjectiveCListener) EnterPostfix_expression(ctx *Postfix_expressionContext) {}

// ExitPostfix_expression is called when production postfix_expression is exited.
func (s *BaseObjectiveCListener) ExitPostfix_expression(ctx *Postfix_expressionContext) {}

// EnterArgument_expression_list is called when production argument_expression_list is entered.
func (s *BaseObjectiveCListener) EnterArgument_expression_list(ctx *Argument_expression_listContext) {}

// ExitArgument_expression_list is called when production argument_expression_list is exited.
func (s *BaseObjectiveCListener) ExitArgument_expression_list(ctx *Argument_expression_listContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseObjectiveCListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseObjectiveCListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseObjectiveCListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseObjectiveCListener) ExitConstant(ctx *ConstantContext) {}
