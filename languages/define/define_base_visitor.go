// Code generated from Define.g4 by ANTLR 4.7.2. DO NOT EDIT.

package parser // Define

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseDefineVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseDefineVisitor) VisitCompilationUnit(ctx *CompilationUnitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDefineVisitor) VisitSymbolDelcaration(ctx *SymbolDelcarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDefineVisitor) VisitNormalDeclarations(ctx *NormalDeclarationsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDefineVisitor) VisitDefineDeclaration(ctx *DefineDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDefineVisitor) VisitDefineExpress(ctx *DefineExpressContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDefineVisitor) VisitSymbolKey(ctx *SymbolKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDefineVisitor) VisitSymbolValue(ctx *SymbolValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDefineVisitor) VisitDefineBody(ctx *DefineBodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDefineVisitor) VisitTemplateData(ctx *TemplateDataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDefineVisitor) VisitSystemDeclaration(ctx *SystemDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDefineVisitor) VisitDefineKey(ctx *DefineKeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDefineVisitor) VisitDefineValue(ctx *DefineValueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDefineVisitor) VisitModuleDeclaration(ctx *ModuleDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDefineVisitor) VisitModuleDefine(ctx *ModuleDefineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDefineVisitor) VisitModuleAttributes(ctx *ModuleAttributesContext) interface{} {
	return v.VisitChildren(ctx)
}
