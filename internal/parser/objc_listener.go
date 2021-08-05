// Code generated from ObjC.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // ObjC

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ObjCListener is a complete listener for a parse tree produced by ObjCParser.
type ObjCListener interface {
	antlr.ParseTreeListener

	// EnterTranslation_unit is called when entering the translation_unit production.
	EnterTranslation_unit(c *Translation_unitContext)

	// EnterExternal_declaration is called when entering the external_declaration production.
	EnterExternal_declaration(c *External_declarationContext)

	// EnterPreprocessor_declaration is called when entering the preprocessor_declaration production.
	EnterPreprocessor_declaration(c *Preprocessor_declarationContext)

	// EnterClass_interface is called when entering the class_interface production.
	EnterClass_interface(c *Class_interfaceContext)

	// EnterCategory_interface is called when entering the category_interface production.
	EnterCategory_interface(c *Category_interfaceContext)

	// EnterClass_implementation is called when entering the class_implementation production.
	EnterClass_implementation(c *Class_implementationContext)

	// EnterCategory_implementation is called when entering the category_implementation production.
	EnterCategory_implementation(c *Category_implementationContext)

	// EnterProtocol_declaration is called when entering the protocol_declaration production.
	EnterProtocol_declaration(c *Protocol_declarationContext)

	// EnterProtocol_declaration_list is called when entering the protocol_declaration_list production.
	EnterProtocol_declaration_list(c *Protocol_declaration_listContext)

	// EnterClass_declaration_list is called when entering the class_declaration_list production.
	EnterClass_declaration_list(c *Class_declaration_listContext)

	// EnterClass_list is called when entering the class_list production.
	EnterClass_list(c *Class_listContext)

	// EnterProtocol_reference_list is called when entering the protocol_reference_list production.
	EnterProtocol_reference_list(c *Protocol_reference_listContext)

	// EnterProtocol_list is called when entering the protocol_list production.
	EnterProtocol_list(c *Protocol_listContext)

	// EnterProperty_declaration is called when entering the property_declaration production.
	EnterProperty_declaration(c *Property_declarationContext)

	// EnterProperty_attributes_declaration is called when entering the property_attributes_declaration production.
	EnterProperty_attributes_declaration(c *Property_attributes_declarationContext)

	// EnterProperty_attributes_list is called when entering the property_attributes_list production.
	EnterProperty_attributes_list(c *Property_attributes_listContext)

	// EnterProperty_attribute is called when entering the property_attribute production.
	EnterProperty_attribute(c *Property_attributeContext)

	// EnterClass_name is called when entering the class_name production.
	EnterClass_name(c *Class_nameContext)

	// EnterSuperclass_name is called when entering the superclass_name production.
	EnterSuperclass_name(c *Superclass_nameContext)

	// EnterCategory_name is called when entering the category_name production.
	EnterCategory_name(c *Category_nameContext)

	// EnterProtocol_name is called when entering the protocol_name production.
	EnterProtocol_name(c *Protocol_nameContext)

	// EnterInstance_variables is called when entering the instance_variables production.
	EnterInstance_variables(c *Instance_variablesContext)

	// EnterVisibility_specification is called when entering the visibility_specification production.
	EnterVisibility_specification(c *Visibility_specificationContext)

	// EnterInterface_declaration_list is called when entering the interface_declaration_list production.
	EnterInterface_declaration_list(c *Interface_declaration_listContext)

	// EnterClass_method_declaration is called when entering the class_method_declaration production.
	EnterClass_method_declaration(c *Class_method_declarationContext)

	// EnterInstance_method_declaration is called when entering the instance_method_declaration production.
	EnterInstance_method_declaration(c *Instance_method_declarationContext)

	// EnterMethod_declaration is called when entering the method_declaration production.
	EnterMethod_declaration(c *Method_declarationContext)

	// EnterImplementation_definition_list is called when entering the implementation_definition_list production.
	EnterImplementation_definition_list(c *Implementation_definition_listContext)

	// EnterClass_method_definition is called when entering the class_method_definition production.
	EnterClass_method_definition(c *Class_method_definitionContext)

	// EnterInstance_method_definition is called when entering the instance_method_definition production.
	EnterInstance_method_definition(c *Instance_method_definitionContext)

	// EnterMethod_definition is called when entering the method_definition production.
	EnterMethod_definition(c *Method_definitionContext)

	// EnterMethod_selector is called when entering the method_selector production.
	EnterMethod_selector(c *Method_selectorContext)

	// EnterKeyword_declarator is called when entering the keyword_declarator production.
	EnterKeyword_declarator(c *Keyword_declaratorContext)

	// EnterSelector is called when entering the selector production.
	EnterSelector(c *SelectorContext)

	// EnterMethod_type is called when entering the method_type production.
	EnterMethod_type(c *Method_typeContext)

	// EnterProperty_implementation is called when entering the property_implementation production.
	EnterProperty_implementation(c *Property_implementationContext)

	// EnterProperty_synthesize_list is called when entering the property_synthesize_list production.
	EnterProperty_synthesize_list(c *Property_synthesize_listContext)

	// EnterProperty_synthesize_item is called when entering the property_synthesize_item production.
	EnterProperty_synthesize_item(c *Property_synthesize_itemContext)

	// EnterBlock_type is called when entering the block_type production.
	EnterBlock_type(c *Block_typeContext)

	// EnterType_specifier is called when entering the type_specifier production.
	EnterType_specifier(c *Type_specifierContext)

	// EnterType_qualifier is called when entering the type_qualifier production.
	EnterType_qualifier(c *Type_qualifierContext)

	// EnterProtocol_qualifier is called when entering the protocol_qualifier production.
	EnterProtocol_qualifier(c *Protocol_qualifierContext)

	// EnterPrimary_expression is called when entering the primary_expression production.
	EnterPrimary_expression(c *Primary_expressionContext)

	// EnterDictionary_pair is called when entering the dictionary_pair production.
	EnterDictionary_pair(c *Dictionary_pairContext)

	// EnterDictionary_expression is called when entering the dictionary_expression production.
	EnterDictionary_expression(c *Dictionary_expressionContext)

	// EnterArray_expression is called when entering the array_expression production.
	EnterArray_expression(c *Array_expressionContext)

	// EnterBox_expression is called when entering the box_expression production.
	EnterBox_expression(c *Box_expressionContext)

	// EnterBlock_parameters is called when entering the block_parameters production.
	EnterBlock_parameters(c *Block_parametersContext)

	// EnterBlock_expression is called when entering the block_expression production.
	EnterBlock_expression(c *Block_expressionContext)

	// EnterMessage_expression is called when entering the message_expression production.
	EnterMessage_expression(c *Message_expressionContext)

	// EnterReceiver is called when entering the receiver production.
	EnterReceiver(c *ReceiverContext)

	// EnterMessage_selector is called when entering the message_selector production.
	EnterMessage_selector(c *Message_selectorContext)

	// EnterKeyword_argument is called when entering the keyword_argument production.
	EnterKeyword_argument(c *Keyword_argumentContext)

	// EnterSelector_expression is called when entering the selector_expression production.
	EnterSelector_expression(c *Selector_expressionContext)

	// EnterSelector_name is called when entering the selector_name production.
	EnterSelector_name(c *Selector_nameContext)

	// EnterProtocol_expression is called when entering the protocol_expression production.
	EnterProtocol_expression(c *Protocol_expressionContext)

	// EnterEncode_expression is called when entering the encode_expression production.
	EnterEncode_expression(c *Encode_expressionContext)

	// EnterType_variable_declarator is called when entering the type_variable_declarator production.
	EnterType_variable_declarator(c *Type_variable_declaratorContext)

	// EnterTry_statement is called when entering the try_statement production.
	EnterTry_statement(c *Try_statementContext)

	// EnterCatch_statement is called when entering the catch_statement production.
	EnterCatch_statement(c *Catch_statementContext)

	// EnterFinally_statement is called when entering the finally_statement production.
	EnterFinally_statement(c *Finally_statementContext)

	// EnterThrow_statement is called when entering the throw_statement production.
	EnterThrow_statement(c *Throw_statementContext)

	// EnterTry_block is called when entering the try_block production.
	EnterTry_block(c *Try_blockContext)

	// EnterSynchronized_statement is called when entering the synchronized_statement production.
	EnterSynchronized_statement(c *Synchronized_statementContext)

	// EnterAutorelease_statement is called when entering the autorelease_statement production.
	EnterAutorelease_statement(c *Autorelease_statementContext)

	// EnterFunction_definition is called when entering the function_definition production.
	EnterFunction_definition(c *Function_definitionContext)

	// EnterDeclaration is called when entering the declaration production.
	EnterDeclaration(c *DeclarationContext)

	// EnterDeclaration_specifiers is called when entering the declaration_specifiers production.
	EnterDeclaration_specifiers(c *Declaration_specifiersContext)

	// EnterArc_behaviour_specifier is called when entering the arc_behaviour_specifier production.
	EnterArc_behaviour_specifier(c *Arc_behaviour_specifierContext)

	// EnterStorage_class_specifier is called when entering the storage_class_specifier production.
	EnterStorage_class_specifier(c *Storage_class_specifierContext)

	// EnterInit_declarator_list is called when entering the init_declarator_list production.
	EnterInit_declarator_list(c *Init_declarator_listContext)

	// EnterInit_declarator is called when entering the init_declarator production.
	EnterInit_declarator(c *Init_declaratorContext)

	// EnterStruct_or_union_specifier is called when entering the struct_or_union_specifier production.
	EnterStruct_or_union_specifier(c *Struct_or_union_specifierContext)

	// EnterStruct_declaration is called when entering the struct_declaration production.
	EnterStruct_declaration(c *Struct_declarationContext)

	// EnterSpecifier_qualifier_list is called when entering the specifier_qualifier_list production.
	EnterSpecifier_qualifier_list(c *Specifier_qualifier_listContext)

	// EnterStruct_declarator_list is called when entering the struct_declarator_list production.
	EnterStruct_declarator_list(c *Struct_declarator_listContext)

	// EnterStruct_declarator is called when entering the struct_declarator production.
	EnterStruct_declarator(c *Struct_declaratorContext)

	// EnterEnum_specifier is called when entering the enum_specifier production.
	EnterEnum_specifier(c *Enum_specifierContext)

	// EnterEnumerator_list is called when entering the enumerator_list production.
	EnterEnumerator_list(c *Enumerator_listContext)

	// EnterEnumerator is called when entering the enumerator production.
	EnterEnumerator(c *EnumeratorContext)

	// EnterPointer is called when entering the pointer production.
	EnterPointer(c *PointerContext)

	// EnterDeclarator is called when entering the declarator production.
	EnterDeclarator(c *DeclaratorContext)

	// EnterDirect_declarator is called when entering the direct_declarator production.
	EnterDirect_declarator(c *Direct_declaratorContext)

	// EnterDeclarator_suffix is called when entering the declarator_suffix production.
	EnterDeclarator_suffix(c *Declarator_suffixContext)

	// EnterParameter_list is called when entering the parameter_list production.
	EnterParameter_list(c *Parameter_listContext)

	// EnterParameter_declaration is called when entering the parameter_declaration production.
	EnterParameter_declaration(c *Parameter_declarationContext)

	// EnterInitializer is called when entering the initializer production.
	EnterInitializer(c *InitializerContext)

	// EnterType_name is called when entering the type_name production.
	EnterType_name(c *Type_nameContext)

	// EnterAbstract_declarator is called when entering the abstract_declarator production.
	EnterAbstract_declarator(c *Abstract_declaratorContext)

	// EnterAbstract_declarator_suffix is called when entering the abstract_declarator_suffix production.
	EnterAbstract_declarator_suffix(c *Abstract_declarator_suffixContext)

	// EnterParameter_declaration_list is called when entering the parameter_declaration_list production.
	EnterParameter_declaration_list(c *Parameter_declaration_listContext)

	// EnterStatement_list is called when entering the statement_list production.
	EnterStatement_list(c *Statement_listContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterLabeled_statement is called when entering the labeled_statement production.
	EnterLabeled_statement(c *Labeled_statementContext)

	// EnterCompound_statement is called when entering the compound_statement production.
	EnterCompound_statement(c *Compound_statementContext)

	// EnterSelection_statement is called when entering the selection_statement production.
	EnterSelection_statement(c *Selection_statementContext)

	// EnterFor_in_statement is called when entering the for_in_statement production.
	EnterFor_in_statement(c *For_in_statementContext)

	// EnterFor_statement is called when entering the for_statement production.
	EnterFor_statement(c *For_statementContext)

	// EnterWhile_statement is called when entering the while_statement production.
	EnterWhile_statement(c *While_statementContext)

	// EnterDo_statement is called when entering the do_statement production.
	EnterDo_statement(c *Do_statementContext)

	// EnterIteration_statement is called when entering the iteration_statement production.
	EnterIteration_statement(c *Iteration_statementContext)

	// EnterJump_statement is called when entering the jump_statement production.
	EnterJump_statement(c *Jump_statementContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterAssignment_expression is called when entering the assignment_expression production.
	EnterAssignment_expression(c *Assignment_expressionContext)

	// EnterAssignment_operator is called when entering the assignment_operator production.
	EnterAssignment_operator(c *Assignment_operatorContext)

	// EnterConditional_expression is called when entering the conditional_expression production.
	EnterConditional_expression(c *Conditional_expressionContext)

	// EnterConstant_expression is called when entering the constant_expression production.
	EnterConstant_expression(c *Constant_expressionContext)

	// EnterLogical_or_expression is called when entering the logical_or_expression production.
	EnterLogical_or_expression(c *Logical_or_expressionContext)

	// EnterLogical_and_expression is called when entering the logical_and_expression production.
	EnterLogical_and_expression(c *Logical_and_expressionContext)

	// EnterInclusive_or_expression is called when entering the inclusive_or_expression production.
	EnterInclusive_or_expression(c *Inclusive_or_expressionContext)

	// EnterExclusive_or_expression is called when entering the exclusive_or_expression production.
	EnterExclusive_or_expression(c *Exclusive_or_expressionContext)

	// EnterAnd_expression is called when entering the and_expression production.
	EnterAnd_expression(c *And_expressionContext)

	// EnterEquality_expression is called when entering the equality_expression production.
	EnterEquality_expression(c *Equality_expressionContext)

	// EnterRelational_expression is called when entering the relational_expression production.
	EnterRelational_expression(c *Relational_expressionContext)

	// EnterShift_expression is called when entering the shift_expression production.
	EnterShift_expression(c *Shift_expressionContext)

	// EnterAdditive_expression is called when entering the additive_expression production.
	EnterAdditive_expression(c *Additive_expressionContext)

	// EnterMultiplicative_expression is called when entering the multiplicative_expression production.
	EnterMultiplicative_expression(c *Multiplicative_expressionContext)

	// EnterCast_expression is called when entering the cast_expression production.
	EnterCast_expression(c *Cast_expressionContext)

	// EnterUnary_expression is called when entering the unary_expression production.
	EnterUnary_expression(c *Unary_expressionContext)

	// EnterUnary_operator is called when entering the unary_operator production.
	EnterUnary_operator(c *Unary_operatorContext)

	// EnterPostfix_expression is called when entering the postfix_expression production.
	EnterPostfix_expression(c *Postfix_expressionContext)

	// EnterArgument_expression_list is called when entering the argument_expression_list production.
	EnterArgument_expression_list(c *Argument_expression_listContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// ExitTranslation_unit is called when exiting the translation_unit production.
	ExitTranslation_unit(c *Translation_unitContext)

	// ExitExternal_declaration is called when exiting the external_declaration production.
	ExitExternal_declaration(c *External_declarationContext)

	// ExitPreprocessor_declaration is called when exiting the preprocessor_declaration production.
	ExitPreprocessor_declaration(c *Preprocessor_declarationContext)

	// ExitClass_interface is called when exiting the class_interface production.
	ExitClass_interface(c *Class_interfaceContext)

	// ExitCategory_interface is called when exiting the category_interface production.
	ExitCategory_interface(c *Category_interfaceContext)

	// ExitClass_implementation is called when exiting the class_implementation production.
	ExitClass_implementation(c *Class_implementationContext)

	// ExitCategory_implementation is called when exiting the category_implementation production.
	ExitCategory_implementation(c *Category_implementationContext)

	// ExitProtocol_declaration is called when exiting the protocol_declaration production.
	ExitProtocol_declaration(c *Protocol_declarationContext)

	// ExitProtocol_declaration_list is called when exiting the protocol_declaration_list production.
	ExitProtocol_declaration_list(c *Protocol_declaration_listContext)

	// ExitClass_declaration_list is called when exiting the class_declaration_list production.
	ExitClass_declaration_list(c *Class_declaration_listContext)

	// ExitClass_list is called when exiting the class_list production.
	ExitClass_list(c *Class_listContext)

	// ExitProtocol_reference_list is called when exiting the protocol_reference_list production.
	ExitProtocol_reference_list(c *Protocol_reference_listContext)

	// ExitProtocol_list is called when exiting the protocol_list production.
	ExitProtocol_list(c *Protocol_listContext)

	// ExitProperty_declaration is called when exiting the property_declaration production.
	ExitProperty_declaration(c *Property_declarationContext)

	// ExitProperty_attributes_declaration is called when exiting the property_attributes_declaration production.
	ExitProperty_attributes_declaration(c *Property_attributes_declarationContext)

	// ExitProperty_attributes_list is called when exiting the property_attributes_list production.
	ExitProperty_attributes_list(c *Property_attributes_listContext)

	// ExitProperty_attribute is called when exiting the property_attribute production.
	ExitProperty_attribute(c *Property_attributeContext)

	// ExitClass_name is called when exiting the class_name production.
	ExitClass_name(c *Class_nameContext)

	// ExitSuperclass_name is called when exiting the superclass_name production.
	ExitSuperclass_name(c *Superclass_nameContext)

	// ExitCategory_name is called when exiting the category_name production.
	ExitCategory_name(c *Category_nameContext)

	// ExitProtocol_name is called when exiting the protocol_name production.
	ExitProtocol_name(c *Protocol_nameContext)

	// ExitInstance_variables is called when exiting the instance_variables production.
	ExitInstance_variables(c *Instance_variablesContext)

	// ExitVisibility_specification is called when exiting the visibility_specification production.
	ExitVisibility_specification(c *Visibility_specificationContext)

	// ExitInterface_declaration_list is called when exiting the interface_declaration_list production.
	ExitInterface_declaration_list(c *Interface_declaration_listContext)

	// ExitClass_method_declaration is called when exiting the class_method_declaration production.
	ExitClass_method_declaration(c *Class_method_declarationContext)

	// ExitInstance_method_declaration is called when exiting the instance_method_declaration production.
	ExitInstance_method_declaration(c *Instance_method_declarationContext)

	// ExitMethod_declaration is called when exiting the method_declaration production.
	ExitMethod_declaration(c *Method_declarationContext)

	// ExitImplementation_definition_list is called when exiting the implementation_definition_list production.
	ExitImplementation_definition_list(c *Implementation_definition_listContext)

	// ExitClass_method_definition is called when exiting the class_method_definition production.
	ExitClass_method_definition(c *Class_method_definitionContext)

	// ExitInstance_method_definition is called when exiting the instance_method_definition production.
	ExitInstance_method_definition(c *Instance_method_definitionContext)

	// ExitMethod_definition is called when exiting the method_definition production.
	ExitMethod_definition(c *Method_definitionContext)

	// ExitMethod_selector is called when exiting the method_selector production.
	ExitMethod_selector(c *Method_selectorContext)

	// ExitKeyword_declarator is called when exiting the keyword_declarator production.
	ExitKeyword_declarator(c *Keyword_declaratorContext)

	// ExitSelector is called when exiting the selector production.
	ExitSelector(c *SelectorContext)

	// ExitMethod_type is called when exiting the method_type production.
	ExitMethod_type(c *Method_typeContext)

	// ExitProperty_implementation is called when exiting the property_implementation production.
	ExitProperty_implementation(c *Property_implementationContext)

	// ExitProperty_synthesize_list is called when exiting the property_synthesize_list production.
	ExitProperty_synthesize_list(c *Property_synthesize_listContext)

	// ExitProperty_synthesize_item is called when exiting the property_synthesize_item production.
	ExitProperty_synthesize_item(c *Property_synthesize_itemContext)

	// ExitBlock_type is called when exiting the block_type production.
	ExitBlock_type(c *Block_typeContext)

	// ExitType_specifier is called when exiting the type_specifier production.
	ExitType_specifier(c *Type_specifierContext)

	// ExitType_qualifier is called when exiting the type_qualifier production.
	ExitType_qualifier(c *Type_qualifierContext)

	// ExitProtocol_qualifier is called when exiting the protocol_qualifier production.
	ExitProtocol_qualifier(c *Protocol_qualifierContext)

	// ExitPrimary_expression is called when exiting the primary_expression production.
	ExitPrimary_expression(c *Primary_expressionContext)

	// ExitDictionary_pair is called when exiting the dictionary_pair production.
	ExitDictionary_pair(c *Dictionary_pairContext)

	// ExitDictionary_expression is called when exiting the dictionary_expression production.
	ExitDictionary_expression(c *Dictionary_expressionContext)

	// ExitArray_expression is called when exiting the array_expression production.
	ExitArray_expression(c *Array_expressionContext)

	// ExitBox_expression is called when exiting the box_expression production.
	ExitBox_expression(c *Box_expressionContext)

	// ExitBlock_parameters is called when exiting the block_parameters production.
	ExitBlock_parameters(c *Block_parametersContext)

	// ExitBlock_expression is called when exiting the block_expression production.
	ExitBlock_expression(c *Block_expressionContext)

	// ExitMessage_expression is called when exiting the message_expression production.
	ExitMessage_expression(c *Message_expressionContext)

	// ExitReceiver is called when exiting the receiver production.
	ExitReceiver(c *ReceiverContext)

	// ExitMessage_selector is called when exiting the message_selector production.
	ExitMessage_selector(c *Message_selectorContext)

	// ExitKeyword_argument is called when exiting the keyword_argument production.
	ExitKeyword_argument(c *Keyword_argumentContext)

	// ExitSelector_expression is called when exiting the selector_expression production.
	ExitSelector_expression(c *Selector_expressionContext)

	// ExitSelector_name is called when exiting the selector_name production.
	ExitSelector_name(c *Selector_nameContext)

	// ExitProtocol_expression is called when exiting the protocol_expression production.
	ExitProtocol_expression(c *Protocol_expressionContext)

	// ExitEncode_expression is called when exiting the encode_expression production.
	ExitEncode_expression(c *Encode_expressionContext)

	// ExitType_variable_declarator is called when exiting the type_variable_declarator production.
	ExitType_variable_declarator(c *Type_variable_declaratorContext)

	// ExitTry_statement is called when exiting the try_statement production.
	ExitTry_statement(c *Try_statementContext)

	// ExitCatch_statement is called when exiting the catch_statement production.
	ExitCatch_statement(c *Catch_statementContext)

	// ExitFinally_statement is called when exiting the finally_statement production.
	ExitFinally_statement(c *Finally_statementContext)

	// ExitThrow_statement is called when exiting the throw_statement production.
	ExitThrow_statement(c *Throw_statementContext)

	// ExitTry_block is called when exiting the try_block production.
	ExitTry_block(c *Try_blockContext)

	// ExitSynchronized_statement is called when exiting the synchronized_statement production.
	ExitSynchronized_statement(c *Synchronized_statementContext)

	// ExitAutorelease_statement is called when exiting the autorelease_statement production.
	ExitAutorelease_statement(c *Autorelease_statementContext)

	// ExitFunction_definition is called when exiting the function_definition production.
	ExitFunction_definition(c *Function_definitionContext)

	// ExitDeclaration is called when exiting the declaration production.
	ExitDeclaration(c *DeclarationContext)

	// ExitDeclaration_specifiers is called when exiting the declaration_specifiers production.
	ExitDeclaration_specifiers(c *Declaration_specifiersContext)

	// ExitArc_behaviour_specifier is called when exiting the arc_behaviour_specifier production.
	ExitArc_behaviour_specifier(c *Arc_behaviour_specifierContext)

	// ExitStorage_class_specifier is called when exiting the storage_class_specifier production.
	ExitStorage_class_specifier(c *Storage_class_specifierContext)

	// ExitInit_declarator_list is called when exiting the init_declarator_list production.
	ExitInit_declarator_list(c *Init_declarator_listContext)

	// ExitInit_declarator is called when exiting the init_declarator production.
	ExitInit_declarator(c *Init_declaratorContext)

	// ExitStruct_or_union_specifier is called when exiting the struct_or_union_specifier production.
	ExitStruct_or_union_specifier(c *Struct_or_union_specifierContext)

	// ExitStruct_declaration is called when exiting the struct_declaration production.
	ExitStruct_declaration(c *Struct_declarationContext)

	// ExitSpecifier_qualifier_list is called when exiting the specifier_qualifier_list production.
	ExitSpecifier_qualifier_list(c *Specifier_qualifier_listContext)

	// ExitStruct_declarator_list is called when exiting the struct_declarator_list production.
	ExitStruct_declarator_list(c *Struct_declarator_listContext)

	// ExitStruct_declarator is called when exiting the struct_declarator production.
	ExitStruct_declarator(c *Struct_declaratorContext)

	// ExitEnum_specifier is called when exiting the enum_specifier production.
	ExitEnum_specifier(c *Enum_specifierContext)

	// ExitEnumerator_list is called when exiting the enumerator_list production.
	ExitEnumerator_list(c *Enumerator_listContext)

	// ExitEnumerator is called when exiting the enumerator production.
	ExitEnumerator(c *EnumeratorContext)

	// ExitPointer is called when exiting the pointer production.
	ExitPointer(c *PointerContext)

	// ExitDeclarator is called when exiting the declarator production.
	ExitDeclarator(c *DeclaratorContext)

	// ExitDirect_declarator is called when exiting the direct_declarator production.
	ExitDirect_declarator(c *Direct_declaratorContext)

	// ExitDeclarator_suffix is called when exiting the declarator_suffix production.
	ExitDeclarator_suffix(c *Declarator_suffixContext)

	// ExitParameter_list is called when exiting the parameter_list production.
	ExitParameter_list(c *Parameter_listContext)

	// ExitParameter_declaration is called when exiting the parameter_declaration production.
	ExitParameter_declaration(c *Parameter_declarationContext)

	// ExitInitializer is called when exiting the initializer production.
	ExitInitializer(c *InitializerContext)

	// ExitType_name is called when exiting the type_name production.
	ExitType_name(c *Type_nameContext)

	// ExitAbstract_declarator is called when exiting the abstract_declarator production.
	ExitAbstract_declarator(c *Abstract_declaratorContext)

	// ExitAbstract_declarator_suffix is called when exiting the abstract_declarator_suffix production.
	ExitAbstract_declarator_suffix(c *Abstract_declarator_suffixContext)

	// ExitParameter_declaration_list is called when exiting the parameter_declaration_list production.
	ExitParameter_declaration_list(c *Parameter_declaration_listContext)

	// ExitStatement_list is called when exiting the statement_list production.
	ExitStatement_list(c *Statement_listContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitLabeled_statement is called when exiting the labeled_statement production.
	ExitLabeled_statement(c *Labeled_statementContext)

	// ExitCompound_statement is called when exiting the compound_statement production.
	ExitCompound_statement(c *Compound_statementContext)

	// ExitSelection_statement is called when exiting the selection_statement production.
	ExitSelection_statement(c *Selection_statementContext)

	// ExitFor_in_statement is called when exiting the for_in_statement production.
	ExitFor_in_statement(c *For_in_statementContext)

	// ExitFor_statement is called when exiting the for_statement production.
	ExitFor_statement(c *For_statementContext)

	// ExitWhile_statement is called when exiting the while_statement production.
	ExitWhile_statement(c *While_statementContext)

	// ExitDo_statement is called when exiting the do_statement production.
	ExitDo_statement(c *Do_statementContext)

	// ExitIteration_statement is called when exiting the iteration_statement production.
	ExitIteration_statement(c *Iteration_statementContext)

	// ExitJump_statement is called when exiting the jump_statement production.
	ExitJump_statement(c *Jump_statementContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitAssignment_expression is called when exiting the assignment_expression production.
	ExitAssignment_expression(c *Assignment_expressionContext)

	// ExitAssignment_operator is called when exiting the assignment_operator production.
	ExitAssignment_operator(c *Assignment_operatorContext)

	// ExitConditional_expression is called when exiting the conditional_expression production.
	ExitConditional_expression(c *Conditional_expressionContext)

	// ExitConstant_expression is called when exiting the constant_expression production.
	ExitConstant_expression(c *Constant_expressionContext)

	// ExitLogical_or_expression is called when exiting the logical_or_expression production.
	ExitLogical_or_expression(c *Logical_or_expressionContext)

	// ExitLogical_and_expression is called when exiting the logical_and_expression production.
	ExitLogical_and_expression(c *Logical_and_expressionContext)

	// ExitInclusive_or_expression is called when exiting the inclusive_or_expression production.
	ExitInclusive_or_expression(c *Inclusive_or_expressionContext)

	// ExitExclusive_or_expression is called when exiting the exclusive_or_expression production.
	ExitExclusive_or_expression(c *Exclusive_or_expressionContext)

	// ExitAnd_expression is called when exiting the and_expression production.
	ExitAnd_expression(c *And_expressionContext)

	// ExitEquality_expression is called when exiting the equality_expression production.
	ExitEquality_expression(c *Equality_expressionContext)

	// ExitRelational_expression is called when exiting the relational_expression production.
	ExitRelational_expression(c *Relational_expressionContext)

	// ExitShift_expression is called when exiting the shift_expression production.
	ExitShift_expression(c *Shift_expressionContext)

	// ExitAdditive_expression is called when exiting the additive_expression production.
	ExitAdditive_expression(c *Additive_expressionContext)

	// ExitMultiplicative_expression is called when exiting the multiplicative_expression production.
	ExitMultiplicative_expression(c *Multiplicative_expressionContext)

	// ExitCast_expression is called when exiting the cast_expression production.
	ExitCast_expression(c *Cast_expressionContext)

	// ExitUnary_expression is called when exiting the unary_expression production.
	ExitUnary_expression(c *Unary_expressionContext)

	// ExitUnary_operator is called when exiting the unary_operator production.
	ExitUnary_operator(c *Unary_operatorContext)

	// ExitPostfix_expression is called when exiting the postfix_expression production.
	ExitPostfix_expression(c *Postfix_expressionContext)

	// ExitArgument_expression_list is called when exiting the argument_expression_list production.
	ExitArgument_expression_list(c *Argument_expression_listContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)
}
