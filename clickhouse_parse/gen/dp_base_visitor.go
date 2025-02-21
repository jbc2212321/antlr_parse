// Code generated from /Users/bytedance/go/src/code.byted.org/lark-approval/ai_sdk/clickhouse_parse/DP.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // DP

import "github.com/antlr4-go/antlr/v4"

type BaseDPVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseDPVisitor) VisitQueryStmt(ctx *QueryStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitQuery(ctx *QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitCtes(ctx *CtesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitNamedQuery(ctx *NamedQueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnAliases(ctx *ColumnAliasesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitSelectUnionStmt(ctx *SelectUnionStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitSelectStmtWithParens(ctx *SelectStmtWithParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitSelectStmt(ctx *SelectStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitWithClause(ctx *WithClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitTopClause(ctx *TopClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitFromClause(ctx *FromClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitArrayJoinClause(ctx *ArrayJoinClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitWindowClause(ctx *WindowClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitPrewhereClause(ctx *PrewhereClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitWhereClause(ctx *WhereClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitGroupByClause(ctx *GroupByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitHavingClause(ctx *HavingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitOrderByClause(ctx *OrderByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitLimitByClause(ctx *LimitByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitLimitClause(ctx *LimitClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitTealimitClause(ctx *TealimitClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitSettingsClause(ctx *SettingsClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitJoinExprOp(ctx *JoinExprOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitJoinExprTable(ctx *JoinExprTableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitJoinExprParens(ctx *JoinExprParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitJoinOpInner(ctx *JoinOpInnerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitJoinOpLeftRight(ctx *JoinOpLeftRightContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitJoinOpFull(ctx *JoinOpFullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitJoinOpCross(ctx *JoinOpCrossContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitJoinConstraintClause(ctx *JoinConstraintClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitSampleClause(ctx *SampleClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitLimitExpr(ctx *LimitExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitTealimitExpr(ctx *TealimitExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitOrderExprList(ctx *OrderExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitOrderExpr(ctx *OrderExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitRatioExpr(ctx *RatioExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitSettingExprList(ctx *SettingExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitSettingExpr(ctx *SettingExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitWindowExpr(ctx *WindowExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitWinPartitionByClause(ctx *WinPartitionByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitWinOrderByClause(ctx *WinOrderByClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitWinFrameClause(ctx *WinFrameClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitFrameStart(ctx *FrameStartContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitFrameBetween(ctx *FrameBetweenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitWinFrameBound(ctx *WinFrameBoundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnTypeExprSimple(ctx *ColumnTypeExprSimpleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnTypeExprNested(ctx *ColumnTypeExprNestedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnTypeExprEnum(ctx *ColumnTypeExprEnumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnTypeExprComplex(ctx *ColumnTypeExprComplexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnTypeExprParam(ctx *ColumnTypeExprParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnExprList(ctx *ColumnExprListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnsExprAsterisk(ctx *ColumnsExprAsteriskContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnsExprSubquery(ctx *ColumnsExprSubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnsExprColumn(ctx *ColumnsExprColumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnExprTernaryOp(ctx *ColumnExprTernaryOpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnExprAlias(ctx *ColumnExprAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnExprExtract(ctx *ColumnExprExtractContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnExprNegate(ctx *ColumnExprNegateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnExprSubquery(ctx *ColumnExprSubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnExprLiteral(ctx *ColumnExprLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnExprBraceParam(ctx *ColumnExprBraceParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnExprArray(ctx *ColumnExprArrayContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnExprSubstring(ctx *ColumnExprSubstringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnExprCast(ctx *ColumnExprCastContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnExprOr(ctx *ColumnExprOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnExprPrecedence1(ctx *ColumnExprPrecedence1Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnExprPrecedence2(ctx *ColumnExprPrecedence2Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnExprPrecedence3(ctx *ColumnExprPrecedence3Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnExprInterval(ctx *ColumnExprIntervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnExprIsNull(ctx *ColumnExprIsNullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnExprWinFunctionTarget(ctx *ColumnExprWinFunctionTargetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnExprTrim(ctx *ColumnExprTrimContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnExprTuple(ctx *ColumnExprTupleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnExprArrayAccess(ctx *ColumnExprArrayAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnExprBetween(ctx *ColumnExprBetweenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnExprParens(ctx *ColumnExprParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnExprTimestamp(ctx *ColumnExprTimestampContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnExprAnd(ctx *ColumnExprAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnExprTupleAccess(ctx *ColumnExprTupleAccessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnExprCase(ctx *ColumnExprCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnExprDate(ctx *ColumnExprDateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnExprNot(ctx *ColumnExprNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnExprWinFunction(ctx *ColumnExprWinFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnExprIdentifier(ctx *ColumnExprIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnExprFunction(ctx *ColumnExprFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnExprAsterisk(ctx *ColumnExprAsteriskContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnArgList(ctx *ColumnArgListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnArgExpr(ctx *ColumnArgExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnLambdaExpr(ctx *ColumnLambdaExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitColumnIdentifier(ctx *ColumnIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitNestedIdentifier(ctx *NestedIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitTableExprIdentifier(ctx *TableExprIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitTableExprSubquery(ctx *TableExprSubqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitTableExprAlias(ctx *TableExprAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitTableExprFunction(ctx *TableExprFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitTableFunctionExpr(ctx *TableFunctionExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitTableIdentifier(ctx *TableIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitTableArgList(ctx *TableArgListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitTableArgExpr(ctx *TableArgExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitDatabaseIdentifier(ctx *DatabaseIdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitFloatingLiteral(ctx *FloatingLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitNumberLiteral(ctx *NumberLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitLiteral(ctx *LiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitInterval(ctx *IntervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitKeyword(ctx *KeywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitKeywordForAlias(ctx *KeywordForAliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitAlias(ctx *AliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitIdentifierOrNull(ctx *IdentifierOrNullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseDPVisitor) VisitEnumValue(ctx *EnumValueContext) interface{} {
	return v.VisitChildren(ctx)
}
