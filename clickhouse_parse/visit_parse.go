package clickhouse_parse

import (
	"context"
	"strings"

	"github.com/antlr4-go/antlr/v4"

	parser "code.byted.org/lark-approval/ai_sdk/clickhouse_parse/gen"
)

type Visitor struct {
	parser.BaseDPVisitor
	SelectCol     []string
	SelectClause  []string
	WhereCol      []string
	WhereClause   []string
	OrderByClause []string
	TableName     []string
	Identifier    []string
	Function      []string
	TableCol      map[string]bool
	HasTable      bool
}

func NewVisitor() *Visitor {
	return &Visitor{
		TableCol: map[string]bool{},
	}
}

func (v *Visitor) VisitQueryStmt(ctx *parser.QueryStmtContext) interface{} {
	return v.visitRule(ctx.Query())

}

func (v *Visitor) VisitQuery(ctx *parser.QueryContext) interface{} {
	return v.visitRule(ctx.SelectStmt())
}

// select
func (v *Visitor) VisitSelectStmt(ctx *parser.SelectStmtContext) interface{} {
	v.visitRule(ctx.ColumnExprList())
	v.visitRule(ctx.FromClause())
	v.visitRule(ctx.WhereClause())
	v.visitRule(ctx.OrderByClause())
	return nil
}
func (v *Visitor) VisitColumnExprList(ctx *parser.ColumnExprListContext) interface{} {
	for idx, _ := range ctx.AllColumnsExpr() {
		v.visitRule(ctx.ColumnsExpr(idx))
	}

	return v.VisitChildren(ctx)
}

// 所有列字段
func (v *Visitor) VisitColumnsExprColumn(ctx *parser.ColumnsExprColumnContext) interface{} {

	if len(ctx.ColumnExpr().GetChildren()) == 1 {

	}

	return v.visitRule(ctx.ColumnExpr())
}

func (v *Visitor) VisitColumnExprFunction(ctx *parser.ColumnExprFunctionContext) interface{} {
	s := ""
	for _, child := range ctx.GetChildren() {
		if tree, ok := child.(antlr.ParseTree); ok {
			s += " " + tree.GetText()
		}
	}
	if !v.HasTable {
		v.SelectClause = append(v.SelectClause, s)
	}
	v.visitRule(ctx.ColumnArgList())
	return v.visitRule(ctx.Identifier())
}

func (v *Visitor) VisitColumnsExprAsterisk(ctx *parser.ColumnsExprAsteriskContext) interface{} {
	v.SelectClause = append(v.SelectClause, ctx.GetText())
	return nil
	//return v.visitRule(ctx.TableIdentifier())
}

// 字段最细粒度、处理count等
func (v *Visitor) VisitColumnArgList(ctx *parser.ColumnArgListContext) interface{} {
	for i, _ := range ctx.AllColumnArgExpr() {

		//v.WhereCol = append(v.WhereCol, ctx.ColumnArgExpr(i).GetText())
		v.visitRule(ctx.ColumnArgExpr(i))
	}

	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitColumnExprInterval(ctx *parser.ColumnExprIntervalContext) interface{} {
	return v.visitRule(ctx.ColumnExpr())
}

func (v *Visitor) VisitColumnIdentifier(ctx *parser.ColumnIdentifierContext) interface{} {

	return v.visitRule(ctx.NestedIdentifier())
}

func (v *Visitor) VisitDatabaseIdentifier(ctx *parser.DatabaseIdentifierContext) interface{} {
	return v.visitRule(ctx.Identifier())
}

func (v *Visitor) VisitColumnExprIdentifier(ctx *parser.ColumnExprIdentifierContext) interface{} {
	return v.visitRule(ctx.ColumnIdentifier())
}

func (v *Visitor) visitRule(node antlr.RuleNode) interface{} {
	if node == nil {
		return nil
	}
	return node.Accept(v)
}

func (v *Visitor) VisitFromClause(ctx *parser.FromClauseContext) interface{} {

	return v.visitRule(ctx.JoinExpr())
}

func (v *Visitor) VisitJoinExprTable(ctx *parser.JoinExprTableContext) interface{} {
	v.TableName = append(v.TableName, ctx.GetText())
	v.HasTable = true
	return v.VisitChildren(ctx)
}

func (v *Visitor) VisitWhereClause(ctx *parser.WhereClauseContext) interface{} {
	return v.visitRule(ctx.ColumnExpr())
}
func (v *Visitor) VisitColumnExprParens(ctx *parser.ColumnExprParensContext) interface{} {
	return v.visitRule(ctx.ColumnExpr())
}

func (v *Visitor) VisitColumnExprAnd(ctx *parser.ColumnExprAndContext) interface{} {

	for i, _ := range ctx.AllColumnExpr() {
		v.visitRule(ctx.ColumnExpr(i))
	}
	return nil
}

func (v *Visitor) VisitColumnExprOr(ctx *parser.ColumnExprOrContext) interface{} {
	var list []string
	for i, _ := range ctx.AllColumnExpr() {
		list = append(list, ctx.ColumnExpr(i).GetText())
		//fmt.Println(ctx.ColumnExpr(i).GetText())
		//v.visitRule(ctx.ColumnExpr(i))
	}
	v.WhereClause = append(v.WhereClause, strings.Join(list, " OR "))
	return nil
}

// where查询表达式
func (v *Visitor) VisitColumnExprPrecedence3(ctx *parser.ColumnExprPrecedence3Context) interface{} {
	v.WhereClause = append(v.WhereClause, ctx.GetText())
	// 目前只需要知道字段名字
	for i, _ := range ctx.AllColumnExpr() {
		if i == 0 {
			v.visitRule(ctx.ColumnExpr(i))

		}

	}
	return nil
}

func (v *Visitor) VisitColumnArgExpr(ctx *parser.ColumnArgExprContext) interface{} {
	return v.visitRule(ctx.ColumnExpr())
}

func (v *Visitor) VisitColumnExprArrayAccess(ctx *parser.ColumnExprArrayAccessContext) interface{} {
	for i, _ := range ctx.AllColumnExpr() {
		// 中括号解析
		if i == 0 {
			//v.WhereCol = append(v.WhereCol, ctx.ColumnExpr(i).GetText())
		}
		v.visitRule(ctx.ColumnExpr(i))
	}
	return nil
}

func (v *Visitor) VisitNestedIdentifier(ctx *parser.NestedIdentifierContext) interface{} {
	for i, _ := range ctx.AllIdentifier() {
		v.visitRule(ctx.Identifier(i))
	}
	return nil
}

func (v *Visitor) VisitIdentifier(ctx *parser.IdentifierContext) interface{} {
	text := ctx.GetText()
	if _, ok := v.TableCol[text]; ok {
		v.TableCol[text] = true
		if v.HasTable {
			v.WhereCol = append(v.WhereCol, text)
		} else {
			v.SelectCol = append(v.SelectCol, text)
			v.SelectClause = append(v.SelectClause, text)

		}

	} else {
		v.Function = append(v.Function, text)
	}
	v.Identifier = append(v.Identifier, text)
	return nil
}

func (v *Visitor) VisitOrderByClause(ctx *parser.OrderByClauseContext) interface{} {
	return v.visitRule(ctx.OrderExprList())
}

func (v *Visitor) VisitOrderExprList(ctx *parser.OrderExprListContext) interface{} {
	for _, exprContext := range ctx.AllOrderExpr() {
		orderbyStr := exprContext.ColumnExpr().GetText()
		if exprContext.DESC() != nil {
			orderbyStr += " " + exprContext.DESC().GetText()
		}
		if exprContext.ASCENDING() != nil {
			orderbyStr += " " + exprContext.ASCENDING().GetText()
		}
		v.visitRule(exprContext.ColumnExpr())
		v.OrderByClause = append(v.OrderByClause, orderbyStr)
	}

	return nil
}

func (v *Visitor) VisitSQL(ctx context.Context, input string) {

	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := parser.NewDPLexer(is)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewDPParser(tokens)

	p.QueryStmt().Accept(v)
}
