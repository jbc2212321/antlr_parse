// Code generated from /Users/bytedance/go/src/code.byted.org/lark-approval/ai_sdk/clickhouse_parse/Time.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Time

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by TimeParser.
type TimeVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by TimeParser#start.
	VisitStart(ctx *StartContext) interface{}

	// Visit a parse tree produced by TimeParser#ToDateTimeOrZero.
	VisitToDateTimeOrZero(ctx *ToDateTimeOrZeroContext) interface{}

	// Visit a parse tree produced by TimeParser#Compare.
	VisitCompare(ctx *CompareContext) interface{}

	// Visit a parse tree produced by TimeParser#Clause.
	VisitClause(ctx *ClauseContext) interface{}

	// Visit a parse tree produced by TimeParser#ToDateTimeOrZeroWithTimeZone.
	VisitToDateTimeOrZeroWithTimeZone(ctx *ToDateTimeOrZeroWithTimeZoneContext) interface{}

	// Visit a parse tree produced by TimeParser#DateAdd.
	VisitDateAdd(ctx *DateAddContext) interface{}

	// Visit a parse tree produced by TimeParser#Date.
	VisitDate(ctx *DateContext) interface{}
}
