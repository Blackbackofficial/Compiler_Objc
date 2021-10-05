// Code generated from ObjC.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // ObjC

import (
	"fmt"
	"github.com/antlr/antlr4/runtime/Go/antlr"
	//"github.com/go-echarts/go-echarts/v2/opts"
)

type BaseObjCVisitor struct {
	*antlr.BaseParseTreeVisitor
}
func NewBaseObjCVisitor() *BaseObjCVisitor {
	y := BaseObjCVisitor {
		//Tree:  Tree{},
		//nodes: []*opts.TreeData{},
		//Flags: Flags{},
	}
	return &y
}

func (v *BaseObjCVisitor) VisitTranslation_unit(ctx *Translation_unitContext) interface{} {
	fmt.Print(ctx.GetText() + "uwfhuwhefhwioefuwehfwefuwfjweifjhwefyguifjwkeofjhuiwegfyuwehfjowkefowgef")
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitExternal_declaration(ctx *External_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitPreprocessor_declaration(ctx *Preprocessor_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitClass_interface(ctx *Class_interfaceContext) interface{} {
	fmt.Print(ctx.GetText() + "uwfhuwhefhwioefuwehfwefuwfjweifjhwefyguifjwkeofjhuiwegfyuwehfjowkefowgef")
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitCategory_interface(ctx *Category_interfaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitClass_implementation(ctx *Class_implementationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitCategory_implementation(ctx *Category_implementationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitProtocol_declaration(ctx *Protocol_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitProtocol_declaration_list(ctx *Protocol_declaration_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitClass_declaration_list(ctx *Class_declaration_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitClass_list(ctx *Class_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitProtocol_reference_list(ctx *Protocol_reference_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitProtocol_list(ctx *Protocol_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitProperty_declaration(ctx *Property_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitProperty_attributes_declaration(ctx *Property_attributes_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitProperty_attributes_list(ctx *Property_attributes_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitProperty_attribute(ctx *Property_attributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitClass_name(ctx *Class_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitSuperclass_name(ctx *Superclass_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitCategory_name(ctx *Category_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitProtocol_name(ctx *Protocol_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitInstance_variables(ctx *Instance_variablesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitVisibility_specification(ctx *Visibility_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitInterface_declaration_list(ctx *Interface_declaration_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitClass_method_declaration(ctx *Class_method_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitInstance_method_declaration(ctx *Instance_method_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitMethod_declaration(ctx *Method_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitImplementation_definition_list(ctx *Implementation_definition_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitClass_method_definition(ctx *Class_method_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitInstance_method_definition(ctx *Instance_method_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitMethod_definition(ctx *Method_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitMethod_selector(ctx *Method_selectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitKeyword_declarator(ctx *Keyword_declaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitSelector(ctx *SelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitMethod_type(ctx *Method_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitProperty_implementation(ctx *Property_implementationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitProperty_synthesize_list(ctx *Property_synthesize_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitProperty_synthesize_item(ctx *Property_synthesize_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitBlock_type(ctx *Block_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitType_specifier(ctx *Type_specifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitType_qualifier(ctx *Type_qualifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitProtocol_qualifier(ctx *Protocol_qualifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitPrimary_expression(ctx *Primary_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitDictionary_pair(ctx *Dictionary_pairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitDictionary_expression(ctx *Dictionary_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitArray_expression(ctx *Array_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitBox_expression(ctx *Box_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitBlock_parameters(ctx *Block_parametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitBlock_expression(ctx *Block_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitMessage_expression(ctx *Message_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitReceiver(ctx *ReceiverContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitMessage_selector(ctx *Message_selectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitKeyword_argument(ctx *Keyword_argumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitSelector_expression(ctx *Selector_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitSelector_name(ctx *Selector_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitProtocol_expression(ctx *Protocol_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitEncode_expression(ctx *Encode_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitType_variable_declarator(ctx *Type_variable_declaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitTry_statement(ctx *Try_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitCatch_statement(ctx *Catch_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitFinally_statement(ctx *Finally_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitThrow_statement(ctx *Throw_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitTry_block(ctx *Try_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitSynchronized_statement(ctx *Synchronized_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitAutorelease_statement(ctx *Autorelease_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitFunction_definition(ctx *Function_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitDeclaration(ctx *DeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitDeclaration_specifiers(ctx *Declaration_specifiersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitArc_behaviour_specifier(ctx *Arc_behaviour_specifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitStorage_class_specifier(ctx *Storage_class_specifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitInit_declarator_list(ctx *Init_declarator_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitInit_declarator(ctx *Init_declaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitStruct_or_union_specifier(ctx *Struct_or_union_specifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitStruct_declaration(ctx *Struct_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitSpecifier_qualifier_list(ctx *Specifier_qualifier_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitStruct_declarator_list(ctx *Struct_declarator_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitStruct_declarator(ctx *Struct_declaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitEnum_specifier(ctx *Enum_specifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitEnumerator_list(ctx *Enumerator_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitEnumerator(ctx *EnumeratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitPointer(ctx *PointerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitDeclarator(ctx *DeclaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitDirect_declarator(ctx *Direct_declaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitDeclarator_suffix(ctx *Declarator_suffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitParameter_list(ctx *Parameter_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitParameter_declaration(ctx *Parameter_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitInitializer(ctx *InitializerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitType_name(ctx *Type_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitAbstract_declarator(ctx *Abstract_declaratorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitAbstract_declarator_suffix(ctx *Abstract_declarator_suffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitParameter_declaration_list(ctx *Parameter_declaration_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitStatement_list(ctx *Statement_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitLabeled_statement(ctx *Labeled_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitCompound_statement(ctx *Compound_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitSelection_statement(ctx *Selection_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitFor_in_statement(ctx *For_in_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitFor_statement(ctx *For_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitWhile_statement(ctx *While_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitDo_statement(ctx *Do_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitIteration_statement(ctx *Iteration_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitJump_statement(ctx *Jump_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	fmt.Print(ctx.GetText() + "uwfhuwhefhwioefuwehfwefuwfjweifjhwefyguifjwkeofjhuiwegfyuwehfjowkefowgef")
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitAssignment_expression(ctx *Assignment_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitAssignment_operator(ctx *Assignment_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitConditional_expression(ctx *Conditional_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitConstant_expression(ctx *Constant_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitLogical_or_expression(ctx *Logical_or_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitLogical_and_expression(ctx *Logical_and_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitInclusive_or_expression(ctx *Inclusive_or_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitExclusive_or_expression(ctx *Exclusive_or_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitAnd_expression(ctx *And_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitEquality_expression(ctx *Equality_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitRelational_expression(ctx *Relational_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitShift_expression(ctx *Shift_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitAdditive_expression(ctx *Additive_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitMultiplicative_expression(ctx *Multiplicative_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitCast_expression(ctx *Cast_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitUnary_expression(ctx *Unary_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitUnary_operator(ctx *Unary_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitPostfix_expression(ctx *Postfix_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitArgument_expression_list(ctx *Argument_expression_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseObjCVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}
