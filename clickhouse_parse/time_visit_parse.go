package clickhouse_parse

import (
	"context"

	"github.com/antlr4-go/antlr/v4"

	timeparser "code.byted.org/lark-approval/ai_sdk/clickhouse_parse/time_gen"
)

type TimeVisitor struct {
	timeparser.BaseTimeVisitor
	CompareType string
	Date        string
	Clause      []string
	DateAdd
}

type DateAdd struct {
	DateType string
	Number   string
}

func NewTimeVisitor() *TimeVisitor {
	return &TimeVisitor{}
}

func (v *TimeVisitor) visitRule(node antlr.RuleNode) interface{} {
	if node == nil {
		return nil
	}
	return node.Accept(v)
}

func (v *TimeVisitor) VisitStart(ctx *timeparser.StartContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *TimeVisitor) VisitCompare(ctx *timeparser.CompareContext) interface{} {
	for _, expressionContext := range ctx.AllExpression() {
		v.visitRule(expressionContext)
	}
	if ctx.GetOp() != nil {
		v.CompareType = ctx.GetOp().GetText()
	}
	return nil
}

func (v *TimeVisitor) VisitToDateTimeOrZero(ctx *timeparser.ToDateTimeOrZeroContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *TimeVisitor) VisitToDateTimeOrZeroWithTimeZone(ctx *timeparser.ToDateTimeOrZeroWithTimeZoneContext) interface{} {
	return v.visitRule(ctx.Expression())
}

func (v *TimeVisitor) VisitDateAdd(ctx *timeparser.DateAddContext) interface{} {
	v.DateAdd.DateType = ctx.DateType().GetText()
	if ctx.NUMBER() != nil {
		v.DateAdd.Number = ctx.NUMBER().GetText()
	}
	return v.visitRule(ctx.Expression())
}

func (v *TimeVisitor) VisitDate(ctx *timeparser.DateContext) interface{} {
	v.Date = ctx.Date().GetText()
	return nil

}

func (v *TimeVisitor) VisitClause(ctx *timeparser.ClauseContext) interface{} {
	v.Clause = append(v.Clause, ctx.Clause().GetText())
	return nil
}

func (v *TimeVisitor) VisitTime(ctx context.Context, input string) {

	is := antlr.NewInputStream(input)

	// Create the Lexer
	lexer := timeparser.NewTimeLexer(is)
	tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := timeparser.NewTimeParser(tokens)

	p.Start_().Accept(v)
}
