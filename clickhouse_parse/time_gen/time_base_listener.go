// Code generated from /Users/bytedance/go/src/code.byted.org/lark-approval/ai_sdk/clickhouse_parse/Time.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Time

import "github.com/antlr4-go/antlr/v4"

// BaseTimeListener is a complete listener for a parse tree produced by TimeParser.
type BaseTimeListener struct{}

var _ TimeListener = &BaseTimeListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseTimeListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseTimeListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseTimeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseTimeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseTimeListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseTimeListener) ExitStart(ctx *StartContext) {}

// EnterToDateTimeOrZero is called when production ToDateTimeOrZero is entered.
func (s *BaseTimeListener) EnterToDateTimeOrZero(ctx *ToDateTimeOrZeroContext) {}

// ExitToDateTimeOrZero is called when production ToDateTimeOrZero is exited.
func (s *BaseTimeListener) ExitToDateTimeOrZero(ctx *ToDateTimeOrZeroContext) {}

// EnterCompare is called when production Compare is entered.
func (s *BaseTimeListener) EnterCompare(ctx *CompareContext) {}

// ExitCompare is called when production Compare is exited.
func (s *BaseTimeListener) ExitCompare(ctx *CompareContext) {}

// EnterClause is called when production Clause is entered.
func (s *BaseTimeListener) EnterClause(ctx *ClauseContext) {}

// ExitClause is called when production Clause is exited.
func (s *BaseTimeListener) ExitClause(ctx *ClauseContext) {}

// EnterToDateTimeOrZeroWithTimeZone is called when production ToDateTimeOrZeroWithTimeZone is entered.
func (s *BaseTimeListener) EnterToDateTimeOrZeroWithTimeZone(ctx *ToDateTimeOrZeroWithTimeZoneContext) {
}

// ExitToDateTimeOrZeroWithTimeZone is called when production ToDateTimeOrZeroWithTimeZone is exited.
func (s *BaseTimeListener) ExitToDateTimeOrZeroWithTimeZone(ctx *ToDateTimeOrZeroWithTimeZoneContext) {
}

// EnterDateAdd is called when production DateAdd is entered.
func (s *BaseTimeListener) EnterDateAdd(ctx *DateAddContext) {}

// ExitDateAdd is called when production DateAdd is exited.
func (s *BaseTimeListener) ExitDateAdd(ctx *DateAddContext) {}

// EnterDate is called when production Date is entered.
func (s *BaseTimeListener) EnterDate(ctx *DateContext) {}

// ExitDate is called when production Date is exited.
func (s *BaseTimeListener) ExitDate(ctx *DateContext) {}
