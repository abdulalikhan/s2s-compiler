// Code generated from MiniImp.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MiniImp

import "github.com/antlr4-go/antlr/v4"

type BaseMiniImpVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseMiniImpVisitor) VisitTruth(ctx *TruthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniImpVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniImpVisitor) VisitSimpleExpr(ctx *SimpleExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniImpVisitor) VisitTerm(ctx *TermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniImpVisitor) VisitFactor(ctx *FactorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniImpVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniImpVisitor) VisitSelect(ctx *SelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniImpVisitor) VisitIterat(ctx *IteratContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniImpVisitor) VisitSet(ctx *SetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniImpVisitor) VisitWrite(ctx *WriteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniImpVisitor) VisitRead(ctx *ReadContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniImpVisitor) VisitDecl(ctx *DeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniImpVisitor) VisitVariable(ctx *VariableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniImpVisitor) VisitAsStr(ctx *AsStrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniImpVisitor) VisitAsInt(ctx *AsIntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniImpVisitor) VisitStmts(ctx *StmtsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniImpVisitor) VisitDecls(ctx *DeclsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniImpVisitor) VisitScope(ctx *ScopeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniImpVisitor) VisitInit(ctx *InitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseMiniImpVisitor) VisitProg(ctx *ProgContext) interface{} {
	return v.VisitChildren(ctx)
}
