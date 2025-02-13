package visitors

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	"miniimpplus/parser"
)

// PythonVisitor implements parser.MiniImpVisitor and antlr.ParseTreeVisitor
type PythonVisitor struct {
	*parser.BaseMiniImpVisitor
	builder         strings.Builder
	currIndentation int
}

// NewPythonVisitor creates a new instance of PythonVisitor
func NewPythonVisitor() *PythonVisitor {
	return &PythonVisitor{
		BaseMiniImpVisitor: &parser.BaseMiniImpVisitor{},
	}
}

func (pv *PythonVisitor) Generate(ctx *parser.ProgContext) string {
	return pv.VisitProg(ctx).(string)
}

// indent returns the current indentation
func (pv *PythonVisitor) indent() string {
	return strings.Repeat("  ", pv.currIndentation)
}

func (pv *PythonVisitor) Visit(tree antlr.ParseTree) interface{} {
	return tree.Accept(pv)
}

func (pv *PythonVisitor) VisitAsInt(ctx *parser.AsIntContext) interface{} {
	expr := pv.Visit(ctx.Expr())
	return "int(" + expr.(string) + ")"
}

func (pv *PythonVisitor) VisitAsStr(ctx *parser.AsStrContext) interface{} {
	expr := pv.Visit(ctx.Expr())
	return "str(" + expr.(string) + ")"
}

func (pv *PythonVisitor) VisitDecl(ctx *parser.DeclContext) interface{} {
	return pv.Visit(ctx.Variable())
}

func (pv *PythonVisitor) VisitStmt(ctx *parser.StmtContext) interface{} {
	child := ctx.GetChild(0).(antlr.ParseTree) // explicitly cast here
	return pv.Visit(child)
}

// VisitErrorNode is needed for implementing antlr.ParseTreeVisitor
func (pv *PythonVisitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	return node.GetText()
}

// VisitTerminal returns the text of a terminal node
func (pv *PythonVisitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	return node.GetText()
}

func (pv *PythonVisitor) VisitTerm(ctx *parser.TermContext) interface{} {
	//fmt.Println("Visiting Term:", ctx.GetText())
	//fmt.Println("ctx.ChildCount:", ctx.GetChildCount())
	var term strings.Builder
	firstChild := ctx.GetChild(0)

	//fmt.Printf("First child type: %T\n", firstChild)

	if factorCtx, ok := firstChild.(*parser.FactorContext); ok {
		visited := pv.Visit(factorCtx)
		if visited != nil {
			term.WriteString(visited.(string))
		} else {
			fmt.Println("Error: pv.Visit(factorCtx) returned NIL")
		}
	} else {
		fmt.Println("Unexpected node type:", reflect.TypeOf(firstChild))
	}

	// process any remaining children here
	for i := 1; i < ctx.GetChildCount(); i++ {
		child := ctx.GetChild(i)
		//fmt.Printf("Child %d type: %T\n", i, child)

		term.WriteString(" ")

		switch c := child.(type) {
		case antlr.TerminalNode:
			term.WriteString(c.GetText())
		case antlr.ParserRuleContext:
			visited := pv.Visit(c)
			if visited != nil {
				term.WriteString(visited.(string))
			} else {
				fmt.Println("Error: pv.Visit(child) returned NIL")
			}
		default:
			fmt.Println("Unexpected child node type:", reflect.TypeOf(child))
		}
	}

	return term.String()
}

func (pv *PythonVisitor) VisitFactor(ctx *parser.FactorContext) interface{} {
	//fmt.Println("Visiting Factor:", ctx.GetText())

	// 1 - if Factor contains an expression -> visit it
	if ctx.Expr() != nil {
		//fmt.Println("Factor contains Expr:", ctx.Expr().GetText())
		visited := pv.Visit(ctx.Expr())
		if visited != nil {
			//fmt.Println("Visit(Expr) returned:", visited)
			return visited
		} else {
			fmt.Println("Error: Visit(Expr) returned NIL")
		}
	}

	// 2 - if Factor contains a TerminalNode (e.g. String/Number) - return its text
	if ctx.GetChildCount() == 1 {
		if truthCtx, ok := ctx.GetChild(0).(*parser.TruthContext); ok {
			//fmt.Println("Factor contains TruthContext:", truthCtx.GetText())
			return pv.Visit(truthCtx)
		} else if term, ok := ctx.GetChild(0).(antlr.TerminalNode); ok {
			//fmt.Println("Factor contains TerminalNode:", term.GetText())
			return term.GetText()
		} else {
			fmt.Println("Error: Unexpected Factor child type:", reflect.TypeOf(ctx.GetChild(0)))
		}
	}

	fmt.Println("Factor did not match any expected structure")
	return nil
}

// VisitChildren visits all children of a node and concatenates their results
func (pv *PythonVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	var sb strings.Builder
	childCount := node.GetChildCount()
	for i := 0; i < childCount; i++ {
		child := node.GetChild(i).(antlr.ParseTree)
		res := pv.Visit(child)
		if str, ok := res.(string); ok {
			sb.WriteString(str)
		}
	}
	return sb.String()
}

// VisitProg converts the whole program
func (pv *PythonVisitor) VisitProg(ctx *parser.ProgContext) interface{} {
	progName := ctx.Init().Identifier().GetText()
	pv.builder.WriteString(fmt.Sprintf("def %s()", progName))

	scopeCtx := ctx.Scope()
	if _, ok := interface{}(pv).(parser.MiniImpVisitor); !ok {
		fmt.Println("Error: PythonVisitor does NOT implement MiniImpVisitor")
	}
	//fmt.Println(reflect.TypeOf(pv))
	//fmt.Println(pv)
	//fmt.Printf("pv is: %T\n", pv)
	//fmt.Printf("Embedded visitor inside pv: %T\n", pv.BaseMiniImpVisitor)

	mainScope := pv.Visit(scopeCtx)
	scopeStr := ""
	if mainScope != nil {
		scopeStr = mainScope.(string)
	}
	//pv.builder.WriteString(scopeStr)
	pv.builder.WriteString(":" + scopeStr)

	pv.builder.WriteString(fmt.Sprintf("\n%s()", progName))
	return pv.builder.String()
}

// VisitScope processes a scope block; handling declarations and statements
func (pv *PythonVisitor) VisitScope(ctx *parser.ScopeContext) interface{} {
	//fmt.Println("Visiting Scope...")
	var sb strings.Builder
	//sb.WriteString(":")
	pv.currIndentation++

	if ctx.Decls() != nil {
		for _, decl := range ctx.Decls().AllDecl() {
			sb.WriteString("\n" + pv.indent() + pv.Visit(decl).(string))
		}
	}

	if ctx.Stmts() != nil {
		for _, stmt := range ctx.Stmts().AllStmt() {
			//sb.WriteString("\n" + pv.indent() + pv.Visit(stmt).(string))
			stmtStr := pv.Visit(stmt).(string)
			lines := strings.Split(stmtStr, "\n")
			for _, line := range lines {
				sb.WriteString("\n" + pv.indent() + line)
			}
		}
	}

	pv.currIndentation--
	return sb.String()
}

// VisitWrite converts a write statement to a Python print statement
func (pv *PythonVisitor) VisitWrite(ctx *parser.WriteContext) interface{} {
	//fmt.Println("Visiting Write...")

	if ctx.Expr() == nil {
		//fmt.Println("ctx.Expr() is NIL in VisitWrite")
		return "print()"
	}

	//fmt.Println("ctx.Expr():", ctx.Expr().GetText())
	expr := pv.Visit(ctx.Expr()).(string)

	//fmt.Printf("pv.Visit(ctx.Expr()) returned: %#v\n", expr)
	return fmt.Sprintf("print(%s)", expr)
}

// VisitRead converts a read statement to a Python input call
func (pv *PythonVisitor) VisitRead(ctx *parser.ReadContext) interface{} {
	prompt := `""` // default empty prompt
	if ctx.Expr() != nil {
		prompt = pv.Visit(ctx.Expr()).(string)
	}
	return fmt.Sprintf("input(%s)", prompt)
}

// VisitSet converts a set statement (assignment) to Python
func (pv *PythonVisitor) VisitSet(ctx *parser.SetContext) interface{} {
	ident := ctx.Identifier().GetText()
	expr := pv.Visit(ctx.Expr()).(string)
	return fmt.Sprintf("%s = %s", ident, expr)
}

// VisitVariable converts a variable declaration
func (pv *PythonVisitor) VisitVariable(ctx *parser.VariableContext) interface{} {
	ident := ctx.Identifier().GetText()
	value := "None" // default is `None` if no value
	if ctx.Expr() != nil {
		value = pv.Visit(ctx.Expr()).(string)
	}
	return fmt.Sprintf("%s = %s", ident, value)
}

// VisitSelect converts an if-then-else statement
func (pv *PythonVisitor) VisitSelect(ctx *parser.SelectContext) interface{} {
	var sb strings.Builder
	condition := strings.TrimSpace(pv.Visit(ctx.Expr()).(string))
	ifScope := pv.Visit(ctx.Scope(0))
	var elseScope interface{}
	if len(ctx.AllScope()) > 1 {
		elseScope = pv.Visit(ctx.Scope(1))
	}
	sb.WriteString("if " + condition + ":" + ifScope.(string))
	if elseScope != nil {
		//sb.WriteString("\n" + pv.indent() + "else:" + elseScope.(string))
		sb.WriteString("\n" + "else:" + elseScope.(string))
	}
	return sb.String()
}

// VisitIterat converts a while loop
func (pv *PythonVisitor) VisitIterat(ctx *parser.IteratContext) interface{} {
	condition := strings.TrimSpace(pv.Visit(ctx.Expr()).(string))
	if condition == "" {
		// default to infinite loop
		condition = "True"
	}

	// identify the correct child index for the loop body
	var body strings.Builder
	// hasBody is used for tracking if there are statements inside the loop
	hasBody := false

	for i := 0; i < ctx.GetChildCount(); i++ {
		child := ctx.GetChild(i)

		// condition-related children must be skipped to prevent duplication
		// here 0 = "while", 1 = condition
		if i == 1 {
			continue
		}

		// skip `while`, condition and `begin`
		if _, ok := child.(antlr.TerminalNode); ok {
			continue
		}

		stmt := pv.Visit(child.(antlr.ParseTree))
		if stmtStr, ok := stmt.(string); ok {
			if !hasBody {
				hasBody = true
			}
			body.WriteString("\n" + pv.indent() + stmtStr)
		} else {
			fmt.Println("Error: Unexpected statement type in while loop:", reflect.TypeOf(stmt))
		}
	}

	//fmt.Println("DEBUG: while condition:", condition)
	// only add a colon when there is a body - else use `pass`
	if !hasBody {
		return fmt.Sprintf("while %s:\n%s pass", condition, pv.indent())
	}

	return fmt.Sprintf("while %s:%s", condition, strings.TrimPrefix(body.String(), "\n"))
}

func (pv *PythonVisitor) VisitExpr(ctx *parser.ExprContext) interface{} {
	//fmt.Println("Visiting Expr:", ctx.GetText())
	var expr strings.Builder
	expr.WriteString(pv.Visit(ctx.GetChild(0).(antlr.ParseTree)).(string))
	for i := 1; i < ctx.GetChildCount(); i++ {
		expr.WriteString(" ")
		expr.WriteString(pv.Visit(ctx.GetChild(i).(antlr.ParseTree)).(string))
	}
	return expr.String()
}

func (pv *PythonVisitor) VisitInt(ctx *parser.InitContext) string {
	ident := pv.Visit(ctx.Identifier()).(string)
	return "def " + ident + "():" + "\n"
}

func (pv *PythonVisitor) VisitTruth(ctx *parser.TruthContext) interface{} {
	//fmt.Println("Visiting Truth:", ctx.GetText())
	var expr strings.Builder

	childCount := ctx.GetChildCount()
	for i := 0; i < childCount; i++ {
		child := ctx.GetChild(i).(antlr.ParseTree)
		part, ok := pv.Visit(child).(string)

		if !ok {
			fmt.Println("Error: Visit returned NIL or non-string for child:", child.GetText())
			continue
		}

		switch part {
		case "true":
			expr.WriteString("True")
		case "false":
			expr.WriteString("False")
		case "is":
			if i+2 < childCount {
				left := pv.Visit(ctx.GetChild(i + 1).(antlr.ParseTree)).(string)
				right := pv.Visit(ctx.GetChild(i + 2).(antlr.ParseTree)).(string)
				expr.WriteString(fmt.Sprintf("%s == %s", left, right))
				// skip next two children since they were processed
				i += 2
			} else {
				fmt.Println("Warning: Malformed 'is' expression")
			}
		case "not":
			if i+1 < childCount {
				right := pv.Visit(ctx.GetChild(i + 1).(antlr.ParseTree)).(string)
				expr.WriteString(fmt.Sprintf(" not %s", right))
				// skip next child since it is already processed
				i++
			} else {
				fmt.Println("Warning: Malformed 'not' expression")
			}
		case "or", "and":
			if i > 0 && i+1 < childCount {
				right := pv.Visit(ctx.GetChild(i + 1).(antlr.ParseTree)).(string)
				expr.WriteString(fmt.Sprintf(" %s %s", part, right))
				// skip next child
				i++
			} else {
				fmt.Println("Warning: Malformed logical operator usage:", part)
			}
		default:
			// append any remaining valid part
			expr.WriteString(part)
		}

		// spacing should only be added if/when necessary
		if i < childCount-1 {
			expr.WriteString(" ")
		}
	}
	return strings.TrimSpace(expr.String())
}
