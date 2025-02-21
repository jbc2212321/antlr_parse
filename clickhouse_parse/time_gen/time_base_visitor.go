// Code generated from /Users/bytedance/go/src/code.byted.org/lark-approval/ai_sdk/clickhouse_parse/Time.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Time

import "github.com/antlr4-go/antlr/v4"

type BaseTimeVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseTimeVisitor) VisitStart(ctx *StartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTimeVisitor) VisitToDateTimeOrZero(ctx *ToDateTimeOrZeroContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTimeVisitor) VisitCompare(ctx *CompareContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTimeVisitor) VisitClause(ctx *ClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTimeVisitor) VisitToDateTimeOrZeroWithTimeZone(ctx *ToDateTimeOrZeroWithTimeZoneContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTimeVisitor) VisitDateAdd(ctx *DateAddContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseTimeVisitor) VisitDate(ctx *DateContext) interface{} {
	return v.VisitChildren(ctx)
}
