// Code generated from ObjC.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // ObjC

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by ObjCParser.
type ObjCVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ObjCParser#translation_unit.
	VisitTranslation_unit(ctx *Translation_unitContext) interface{}

	// Visit a parse tree produced by ObjCParser#external_declaration.
	VisitExternal_declaration(ctx *External_declarationContext) interface{}

	// Visit a parse tree produced by ObjCParser#preprocessor_declaration.
	VisitPreprocessor_declaration(ctx *Preprocessor_declarationContext) interface{}

	// Visit a parse tree produced by ObjCParser#class_interface.
	VisitClass_interface(ctx *Class_interfaceContext) interface{}

	// Visit a parse tree produced by ObjCParser#category_interface.
	VisitCategory_interface(ctx *Category_interfaceContext) interface{}

	// Visit a parse tree produced by ObjCParser#class_implementation.
	VisitClass_implementation(ctx *Class_implementationContext) interface{}

	// Visit a parse tree produced by ObjCParser#category_implementation.
	VisitCategory_implementation(ctx *Category_implementationContext) interface{}

	// Visit a parse tree produced by ObjCParser#protocol_declaration.
	VisitProtocol_declaration(ctx *Protocol_declarationContext) interface{}

	// Visit a parse tree produced by ObjCParser#protocol_declaration_list.
	VisitProtocol_declaration_list(ctx *Protocol_declaration_listContext) interface{}

	// Visit a parse tree produced by ObjCParser#class_declaration_list.
	VisitClass_declaration_list(ctx *Class_declaration_listContext) interface{}

	// Visit a parse tree produced by ObjCParser#class_list.
	VisitClass_list(ctx *Class_listContext) interface{}

	// Visit a parse tree produced by ObjCParser#protocol_reference_list.
	VisitProtocol_reference_list(ctx *Protocol_reference_listContext) interface{}

	// Visit a parse tree produced by ObjCParser#protocol_list.
	VisitProtocol_list(ctx *Protocol_listContext) interface{}

	// Visit a parse tree produced by ObjCParser#property_declaration.
	VisitProperty_declaration(ctx *Property_declarationContext) interface{}

	// Visit a parse tree produced by ObjCParser#property_attributes_declaration.
	VisitProperty_attributes_declaration(ctx *Property_attributes_declarationContext) interface{}

	// Visit a parse tree produced by ObjCParser#property_attributes_list.
	VisitProperty_attributes_list(ctx *Property_attributes_listContext) interface{}

	// Visit a parse tree produced by ObjCParser#property_attribute.
	VisitProperty_attribute(ctx *Property_attributeContext) interface{}

	// Visit a parse tree produced by ObjCParser#class_name.
	VisitClass_name(ctx *Class_nameContext) interface{}

	// Visit a parse tree produced by ObjCParser#superclass_name.
	VisitSuperclass_name(ctx *Superclass_nameContext) interface{}

	// Visit a parse tree produced by ObjCParser#category_name.
	VisitCategory_name(ctx *Category_nameContext) interface{}

	// Visit a parse tree produced by ObjCParser#protocol_name.
	VisitProtocol_name(ctx *Protocol_nameContext) interface{}

	// Visit a parse tree produced by ObjCParser#instance_variables.
	VisitInstance_variables(ctx *Instance_variablesContext) interface{}

	// Visit a parse tree produced by ObjCParser#visibility_specification.
	VisitVisibility_specification(ctx *Visibility_specificationContext) interface{}

	// Visit a parse tree produced by ObjCParser#interface_declaration_list.
	VisitInterface_declaration_list(ctx *Interface_declaration_listContext) interface{}

	// Visit a parse tree produced by ObjCParser#class_method_declaration.
	VisitClass_method_declaration(ctx *Class_method_declarationContext) interface{}

	// Visit a parse tree produced by ObjCParser#instance_method_declaration.
	VisitInstance_method_declaration(ctx *Instance_method_declarationContext) interface{}

	// Visit a parse tree produced by ObjCParser#method_declaration.
	VisitMethod_declaration(ctx *Method_declarationContext) interface{}

	// Visit a parse tree produced by ObjCParser#implementation_definition_list.
	VisitImplementation_definition_list(ctx *Implementation_definition_listContext) interface{}

	// Visit a parse tree produced by ObjCParser#class_method_definition.
	VisitClass_method_definition(ctx *Class_method_definitionContext) interface{}

	// Visit a parse tree produced by ObjCParser#instance_method_definition.
	VisitInstance_method_definition(ctx *Instance_method_definitionContext) interface{}

	// Visit a parse tree produced by ObjCParser#method_definition.
	VisitMethod_definition(ctx *Method_definitionContext) interface{}

	// Visit a parse tree produced by ObjCParser#method_selector.
	VisitMethod_selector(ctx *Method_selectorContext) interface{}

	// Visit a parse tree produced by ObjCParser#keyword_declarator.
	VisitKeyword_declarator(ctx *Keyword_declaratorContext) interface{}

	// Visit a parse tree produced by ObjCParser#selector.
	VisitSelector(ctx *SelectorContext) interface{}

	// Visit a parse tree produced by ObjCParser#method_type.
	VisitMethod_type(ctx *Method_typeContext) interface{}

	// Visit a parse tree produced by ObjCParser#property_implementation.
	VisitProperty_implementation(ctx *Property_implementationContext) interface{}

	// Visit a parse tree produced by ObjCParser#property_synthesize_list.
	VisitProperty_synthesize_list(ctx *Property_synthesize_listContext) interface{}

	// Visit a parse tree produced by ObjCParser#property_synthesize_item.
	VisitProperty_synthesize_item(ctx *Property_synthesize_itemContext) interface{}

	// Visit a parse tree produced by ObjCParser#block_type.
	VisitBlock_type(ctx *Block_typeContext) interface{}

	// Visit a parse tree produced by ObjCParser#type_specifier.
	VisitType_specifier(ctx *Type_specifierContext) interface{}

	// Visit a parse tree produced by ObjCParser#type_qualifier.
	VisitType_qualifier(ctx *Type_qualifierContext) interface{}

	// Visit a parse tree produced by ObjCParser#protocol_qualifier.
	VisitProtocol_qualifier(ctx *Protocol_qualifierContext) interface{}

	// Visit a parse tree produced by ObjCParser#primary_expression.
	VisitPrimary_expression(ctx *Primary_expressionContext) interface{}

	// Visit a parse tree produced by ObjCParser#dictionary_pair.
	VisitDictionary_pair(ctx *Dictionary_pairContext) interface{}

	// Visit a parse tree produced by ObjCParser#dictionary_expression.
	VisitDictionary_expression(ctx *Dictionary_expressionContext) interface{}

	// Visit a parse tree produced by ObjCParser#array_expression.
	VisitArray_expression(ctx *Array_expressionContext) interface{}

	// Visit a parse tree produced by ObjCParser#box_expression.
	VisitBox_expression(ctx *Box_expressionContext) interface{}

	// Visit a parse tree produced by ObjCParser#block_parameters.
	VisitBlock_parameters(ctx *Block_parametersContext) interface{}

	// Visit a parse tree produced by ObjCParser#block_expression.
	VisitBlock_expression(ctx *Block_expressionContext) interface{}

	// Visit a parse tree produced by ObjCParser#message_expression.
	VisitMessage_expression(ctx *Message_expressionContext) interface{}

	// Visit a parse tree produced by ObjCParser#receiver.
	VisitReceiver(ctx *ReceiverContext) interface{}

	// Visit a parse tree produced by ObjCParser#message_selector.
	VisitMessage_selector(ctx *Message_selectorContext) interface{}

	// Visit a parse tree produced by ObjCParser#keyword_argument.
	VisitKeyword_argument(ctx *Keyword_argumentContext) interface{}

	// Visit a parse tree produced by ObjCParser#selector_expression.
	VisitSelector_expression(ctx *Selector_expressionContext) interface{}

	// Visit a parse tree produced by ObjCParser#selector_name.
	VisitSelector_name(ctx *Selector_nameContext) interface{}

	// Visit a parse tree produced by ObjCParser#protocol_expression.
	VisitProtocol_expression(ctx *Protocol_expressionContext) interface{}

	// Visit a parse tree produced by ObjCParser#encode_expression.
	VisitEncode_expression(ctx *Encode_expressionContext) interface{}

	// Visit a parse tree produced by ObjCParser#type_variable_declarator.
	VisitType_variable_declarator(ctx *Type_variable_declaratorContext) interface{}

	// Visit a parse tree produced by ObjCParser#try_statement.
	VisitTry_statement(ctx *Try_statementContext) interface{}

	// Visit a parse tree produced by ObjCParser#catch_statement.
	VisitCatch_statement(ctx *Catch_statementContext) interface{}

	// Visit a parse tree produced by ObjCParser#finally_statement.
	VisitFinally_statement(ctx *Finally_statementContext) interface{}

	// Visit a parse tree produced by ObjCParser#throw_statement.
	VisitThrow_statement(ctx *Throw_statementContext) interface{}

	// Visit a parse tree produced by ObjCParser#try_block.
	VisitTry_block(ctx *Try_blockContext) interface{}

	// Visit a parse tree produced by ObjCParser#synchronized_statement.
	VisitSynchronized_statement(ctx *Synchronized_statementContext) interface{}

	// Visit a parse tree produced by ObjCParser#autorelease_statement.
	VisitAutorelease_statement(ctx *Autorelease_statementContext) interface{}

	// Visit a parse tree produced by ObjCParser#function_definition.
	VisitFunction_definition(ctx *Function_definitionContext) interface{}

	// Visit a parse tree produced by ObjCParser#declaration.
	VisitDeclaration(ctx *DeclarationContext) interface{}

	// Visit a parse tree produced by ObjCParser#declaration_specifiers.
	VisitDeclaration_specifiers(ctx *Declaration_specifiersContext) interface{}

	// Visit a parse tree produced by ObjCParser#arc_behaviour_specifier.
	VisitArc_behaviour_specifier(ctx *Arc_behaviour_specifierContext) interface{}

	// Visit a parse tree produced by ObjCParser#storage_class_specifier.
	VisitStorage_class_specifier(ctx *Storage_class_specifierContext) interface{}

	// Visit a parse tree produced by ObjCParser#init_declarator_list.
	VisitInit_declarator_list(ctx *Init_declarator_listContext) interface{}

	// Visit a parse tree produced by ObjCParser#init_declarator.
	VisitInit_declarator(ctx *Init_declaratorContext) interface{}

	// Visit a parse tree produced by ObjCParser#struct_or_union_specifier.
	VisitStruct_or_union_specifier(ctx *Struct_or_union_specifierContext) interface{}

	// Visit a parse tree produced by ObjCParser#struct_declaration.
	VisitStruct_declaration(ctx *Struct_declarationContext) interface{}

	// Visit a parse tree produced by ObjCParser#specifier_qualifier_list.
	VisitSpecifier_qualifier_list(ctx *Specifier_qualifier_listContext) interface{}

	// Visit a parse tree produced by ObjCParser#struct_declarator_list.
	VisitStruct_declarator_list(ctx *Struct_declarator_listContext) interface{}

	// Visit a parse tree produced by ObjCParser#struct_declarator.
	VisitStruct_declarator(ctx *Struct_declaratorContext) interface{}

	// Visit a parse tree produced by ObjCParser#enum_specifier.
	VisitEnum_specifier(ctx *Enum_specifierContext) interface{}

	// Visit a parse tree produced by ObjCParser#enumerator_list.
	VisitEnumerator_list(ctx *Enumerator_listContext) interface{}

	// Visit a parse tree produced by ObjCParser#enumerator.
	VisitEnumerator(ctx *EnumeratorContext) interface{}

	// Visit a parse tree produced by ObjCParser#pointer.
	VisitPointer(ctx *PointerContext) interface{}

	// Visit a parse tree produced by ObjCParser#declarator.
	VisitDeclarator(ctx *DeclaratorContext) interface{}

	// Visit a parse tree produced by ObjCParser#direct_declarator.
	VisitDirect_declarator(ctx *Direct_declaratorContext) interface{}

	// Visit a parse tree produced by ObjCParser#declarator_suffix.
	VisitDeclarator_suffix(ctx *Declarator_suffixContext) interface{}

	// Visit a parse tree produced by ObjCParser#parameter_list.
	VisitParameter_list(ctx *Parameter_listContext) interface{}

	// Visit a parse tree produced by ObjCParser#parameter_declaration.
	VisitParameter_declaration(ctx *Parameter_declarationContext) interface{}

	// Visit a parse tree produced by ObjCParser#initializer.
	VisitInitializer(ctx *InitializerContext) interface{}

	// Visit a parse tree produced by ObjCParser#type_name.
	VisitType_name(ctx *Type_nameContext) interface{}

	// Visit a parse tree produced by ObjCParser#abstract_declarator.
	VisitAbstract_declarator(ctx *Abstract_declaratorContext) interface{}

	// Visit a parse tree produced by ObjCParser#abstract_declarator_suffix.
	VisitAbstract_declarator_suffix(ctx *Abstract_declarator_suffixContext) interface{}

	// Visit a parse tree produced by ObjCParser#parameter_declaration_list.
	VisitParameter_declaration_list(ctx *Parameter_declaration_listContext) interface{}

	// Visit a parse tree produced by ObjCParser#statement_list.
	VisitStatement_list(ctx *Statement_listContext) interface{}

	// Visit a parse tree produced by ObjCParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by ObjCParser#labeled_statement.
	VisitLabeled_statement(ctx *Labeled_statementContext) interface{}

	// Visit a parse tree produced by ObjCParser#compound_statement.
	VisitCompound_statement(ctx *Compound_statementContext) interface{}

	// Visit a parse tree produced by ObjCParser#selection_statement.
	VisitSelection_statement(ctx *Selection_statementContext) interface{}

	// Visit a parse tree produced by ObjCParser#for_in_statement.
	VisitFor_in_statement(ctx *For_in_statementContext) interface{}

	// Visit a parse tree produced by ObjCParser#for_statement.
	VisitFor_statement(ctx *For_statementContext) interface{}

	// Visit a parse tree produced by ObjCParser#while_statement.
	VisitWhile_statement(ctx *While_statementContext) interface{}

	// Visit a parse tree produced by ObjCParser#do_statement.
	VisitDo_statement(ctx *Do_statementContext) interface{}

	// Visit a parse tree produced by ObjCParser#iteration_statement.
	VisitIteration_statement(ctx *Iteration_statementContext) interface{}

	// Visit a parse tree produced by ObjCParser#jump_statement.
	VisitJump_statement(ctx *Jump_statementContext) interface{}

	// Visit a parse tree produced by ObjCParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by ObjCParser#assignment_expression.
	VisitAssignment_expression(ctx *Assignment_expressionContext) interface{}

	// Visit a parse tree produced by ObjCParser#assignment_operator.
	VisitAssignment_operator(ctx *Assignment_operatorContext) interface{}

	// Visit a parse tree produced by ObjCParser#conditional_expression.
	VisitConditional_expression(ctx *Conditional_expressionContext) interface{}

	// Visit a parse tree produced by ObjCParser#constant_expression.
	VisitConstant_expression(ctx *Constant_expressionContext) interface{}

	// Visit a parse tree produced by ObjCParser#logical_or_expression.
	VisitLogical_or_expression(ctx *Logical_or_expressionContext) interface{}

	// Visit a parse tree produced by ObjCParser#logical_and_expression.
	VisitLogical_and_expression(ctx *Logical_and_expressionContext) interface{}

	// Visit a parse tree produced by ObjCParser#inclusive_or_expression.
	VisitInclusive_or_expression(ctx *Inclusive_or_expressionContext) interface{}

	// Visit a parse tree produced by ObjCParser#exclusive_or_expression.
	VisitExclusive_or_expression(ctx *Exclusive_or_expressionContext) interface{}

	// Visit a parse tree produced by ObjCParser#and_expression.
	VisitAnd_expression(ctx *And_expressionContext) interface{}

	// Visit a parse tree produced by ObjCParser#equality_expression.
	VisitEquality_expression(ctx *Equality_expressionContext) interface{}

	// Visit a parse tree produced by ObjCParser#relational_expression.
	VisitRelational_expression(ctx *Relational_expressionContext) interface{}

	// Visit a parse tree produced by ObjCParser#shift_expression.
	VisitShift_expression(ctx *Shift_expressionContext) interface{}

	// Visit a parse tree produced by ObjCParser#additive_expression.
	VisitAdditive_expression(ctx *Additive_expressionContext) interface{}

	// Visit a parse tree produced by ObjCParser#multiplicative_expression.
	VisitMultiplicative_expression(ctx *Multiplicative_expressionContext) interface{}

	// Visit a parse tree produced by ObjCParser#cast_expression.
	VisitCast_expression(ctx *Cast_expressionContext) interface{}

	// Visit a parse tree produced by ObjCParser#unary_expression.
	VisitUnary_expression(ctx *Unary_expressionContext) interface{}

	// Visit a parse tree produced by ObjCParser#unary_operator.
	VisitUnary_operator(ctx *Unary_operatorContext) interface{}

	// Visit a parse tree produced by ObjCParser#postfix_expression.
	VisitPostfix_expression(ctx *Postfix_expressionContext) interface{}

	// Visit a parse tree produced by ObjCParser#argument_expression_list.
	VisitArgument_expression_list(ctx *Argument_expression_listContext) interface{}

	// Visit a parse tree produced by ObjCParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by ObjCParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}
}
