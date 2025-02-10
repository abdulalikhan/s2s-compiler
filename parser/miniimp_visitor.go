// Code generated from MiniImp.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // MiniImp

import (
	"github.com/antlr4-go/antlr/v4"
	"sort"
)

// A complete Visitor for a parse tree produced by MiniImpParser.
type MiniImpVisitor interface {
	antlr.ParseTreeVisitor
	sort.IntSlice()

	// Visit a parse tree produced by MiniImpParser#truth.
	VisitTruth(ctx *TruthContext) interface{}

	// Visit a parse tree produced by MiniImpParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by MiniImpParser#simpleExpr.
	VisitSimpleExpr(ctx *SimpleExprContext) interface{}

	// Visit a parse tree produced by MiniImpParser#term.
	VisitTerm(ctx *TermContext) interface{}

	// Visit a parse tree produced by MiniImpParser#factor.
	VisitFactor(ctx *FactorContext) interface{}

	// Visit a parse tree produced by MiniImpParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by MiniImpParser#select.
	VisitSelect(ctx *SelectContext) interface{}

	// Visit a parse tree produced by MiniImpParser#iterat.
	VisitIterat(ctx *IteratContext) interface{}

	// Visit a parse tree produced by MiniImpParser#set.
	VisitSet(ctx *SetContext) interface{}

	// Visit a parse tree produced by MiniImpParser#write.
	VisitWrite(ctx *WriteContext) interface{}

	// Visit a parse tree produced by MiniImpParser#read.
	VisitRead(ctx *ReadContext) interface{}

	// Visit a parse tree produced by MiniImpParser#decl.
	VisitDecl(ctx *DeclContext) interface{}

	// Visit a parse tree produced by MiniImpParser#variable.
	VisitVariable(ctx *VariableContext) interface{}

	// Visit a parse tree produced by MiniImpParser#asStr.
	VisitAsStr(ctx *AsStrContext) interface{}

	// Visit a parse tree produced by MiniImpParser#asInt.
	VisitAsInt(ctx *AsIntContext) interface{}

	// Visit a parse tree produced by MiniImpParser#stmts.
	VisitStmts(ctx *StmtsContext) interface{}

	// Visit a parse tree produced by MiniImpParser#decls.
	VisitDecls(ctx *DeclsContext) interface{}

	// Visit a parse tree produced by MiniImpParser#scope.
	VisitScope(ctx *ScopeContext) interface{}

	// Visit a parse tree produced by MiniImpParser#init.
	VisitInit(ctx *InitContext) interface{}

	// Visit a parse tree produced by MiniImpParser#prog.
	VisitProg(ctx *ProgContext) interface{}
}
