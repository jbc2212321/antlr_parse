// Code generated from /Users/bytedance/go/src/code.byted.org/lark-approval/ai_sdk/clickhouse_parse/DP.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // DP

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by DPParser.
type DPVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by DPParser#queryStmt.
	VisitQueryStmt(ctx *QueryStmtContext) interface{}

	// Visit a parse tree produced by DPParser#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by DPParser#ctes.
	VisitCtes(ctx *CtesContext) interface{}

	// Visit a parse tree produced by DPParser#namedQuery.
	VisitNamedQuery(ctx *NamedQueryContext) interface{}

	// Visit a parse tree produced by DPParser#columnAliases.
	VisitColumnAliases(ctx *ColumnAliasesContext) interface{}

	// Visit a parse tree produced by DPParser#selectUnionStmt.
	VisitSelectUnionStmt(ctx *SelectUnionStmtContext) interface{}

	// Visit a parse tree produced by DPParser#selectStmtWithParens.
	VisitSelectStmtWithParens(ctx *SelectStmtWithParensContext) interface{}

	// Visit a parse tree produced by DPParser#selectStmt.
	VisitSelectStmt(ctx *SelectStmtContext) interface{}

	// Visit a parse tree produced by DPParser#withClause.
	VisitWithClause(ctx *WithClauseContext) interface{}

	// Visit a parse tree produced by DPParser#topClause.
	VisitTopClause(ctx *TopClauseContext) interface{}

	// Visit a parse tree produced by DPParser#fromClause.
	VisitFromClause(ctx *FromClauseContext) interface{}

	// Visit a parse tree produced by DPParser#arrayJoinClause.
	VisitArrayJoinClause(ctx *ArrayJoinClauseContext) interface{}

	// Visit a parse tree produced by DPParser#windowClause.
	VisitWindowClause(ctx *WindowClauseContext) interface{}

	// Visit a parse tree produced by DPParser#prewhereClause.
	VisitPrewhereClause(ctx *PrewhereClauseContext) interface{}

	// Visit a parse tree produced by DPParser#whereClause.
	VisitWhereClause(ctx *WhereClauseContext) interface{}

	// Visit a parse tree produced by DPParser#groupByClause.
	VisitGroupByClause(ctx *GroupByClauseContext) interface{}

	// Visit a parse tree produced by DPParser#havingClause.
	VisitHavingClause(ctx *HavingClauseContext) interface{}

	// Visit a parse tree produced by DPParser#orderByClause.
	VisitOrderByClause(ctx *OrderByClauseContext) interface{}

	// Visit a parse tree produced by DPParser#limitByClause.
	VisitLimitByClause(ctx *LimitByClauseContext) interface{}

	// Visit a parse tree produced by DPParser#limitClause.
	VisitLimitClause(ctx *LimitClauseContext) interface{}

	// Visit a parse tree produced by DPParser#tealimitClause.
	VisitTealimitClause(ctx *TealimitClauseContext) interface{}

	// Visit a parse tree produced by DPParser#settingsClause.
	VisitSettingsClause(ctx *SettingsClauseContext) interface{}

	// Visit a parse tree produced by DPParser#JoinExprOp.
	VisitJoinExprOp(ctx *JoinExprOpContext) interface{}

	// Visit a parse tree produced by DPParser#JoinExprTable.
	VisitJoinExprTable(ctx *JoinExprTableContext) interface{}

	// Visit a parse tree produced by DPParser#JoinExprParens.
	VisitJoinExprParens(ctx *JoinExprParensContext) interface{}

	// Visit a parse tree produced by DPParser#JoinOpInner.
	VisitJoinOpInner(ctx *JoinOpInnerContext) interface{}

	// Visit a parse tree produced by DPParser#JoinOpLeftRight.
	VisitJoinOpLeftRight(ctx *JoinOpLeftRightContext) interface{}

	// Visit a parse tree produced by DPParser#JoinOpFull.
	VisitJoinOpFull(ctx *JoinOpFullContext) interface{}

	// Visit a parse tree produced by DPParser#joinOpCross.
	VisitJoinOpCross(ctx *JoinOpCrossContext) interface{}

	// Visit a parse tree produced by DPParser#joinConstraintClause.
	VisitJoinConstraintClause(ctx *JoinConstraintClauseContext) interface{}

	// Visit a parse tree produced by DPParser#sampleClause.
	VisitSampleClause(ctx *SampleClauseContext) interface{}

	// Visit a parse tree produced by DPParser#limitExpr.
	VisitLimitExpr(ctx *LimitExprContext) interface{}

	// Visit a parse tree produced by DPParser#tealimitExpr.
	VisitTealimitExpr(ctx *TealimitExprContext) interface{}

	// Visit a parse tree produced by DPParser#orderExprList.
	VisitOrderExprList(ctx *OrderExprListContext) interface{}

	// Visit a parse tree produced by DPParser#orderExpr.
	VisitOrderExpr(ctx *OrderExprContext) interface{}

	// Visit a parse tree produced by DPParser#ratioExpr.
	VisitRatioExpr(ctx *RatioExprContext) interface{}

	// Visit a parse tree produced by DPParser#settingExprList.
	VisitSettingExprList(ctx *SettingExprListContext) interface{}

	// Visit a parse tree produced by DPParser#settingExpr.
	VisitSettingExpr(ctx *SettingExprContext) interface{}

	// Visit a parse tree produced by DPParser#windowExpr.
	VisitWindowExpr(ctx *WindowExprContext) interface{}

	// Visit a parse tree produced by DPParser#winPartitionByClause.
	VisitWinPartitionByClause(ctx *WinPartitionByClauseContext) interface{}

	// Visit a parse tree produced by DPParser#winOrderByClause.
	VisitWinOrderByClause(ctx *WinOrderByClauseContext) interface{}

	// Visit a parse tree produced by DPParser#winFrameClause.
	VisitWinFrameClause(ctx *WinFrameClauseContext) interface{}

	// Visit a parse tree produced by DPParser#frameStart.
	VisitFrameStart(ctx *FrameStartContext) interface{}

	// Visit a parse tree produced by DPParser#frameBetween.
	VisitFrameBetween(ctx *FrameBetweenContext) interface{}

	// Visit a parse tree produced by DPParser#winFrameBound.
	VisitWinFrameBound(ctx *WinFrameBoundContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnTypeExprSimple.
	VisitColumnTypeExprSimple(ctx *ColumnTypeExprSimpleContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnTypeExprNested.
	VisitColumnTypeExprNested(ctx *ColumnTypeExprNestedContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnTypeExprEnum.
	VisitColumnTypeExprEnum(ctx *ColumnTypeExprEnumContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnTypeExprComplex.
	VisitColumnTypeExprComplex(ctx *ColumnTypeExprComplexContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnTypeExprParam.
	VisitColumnTypeExprParam(ctx *ColumnTypeExprParamContext) interface{}

	// Visit a parse tree produced by DPParser#columnExprList.
	VisitColumnExprList(ctx *ColumnExprListContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnsExprAsterisk.
	VisitColumnsExprAsterisk(ctx *ColumnsExprAsteriskContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnsExprSubquery.
	VisitColumnsExprSubquery(ctx *ColumnsExprSubqueryContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnsExprColumn.
	VisitColumnsExprColumn(ctx *ColumnsExprColumnContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnExprTernaryOp.
	VisitColumnExprTernaryOp(ctx *ColumnExprTernaryOpContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnExprAlias.
	VisitColumnExprAlias(ctx *ColumnExprAliasContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnExprExtract.
	VisitColumnExprExtract(ctx *ColumnExprExtractContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnExprNegate.
	VisitColumnExprNegate(ctx *ColumnExprNegateContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnExprSubquery.
	VisitColumnExprSubquery(ctx *ColumnExprSubqueryContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnExprLiteral.
	VisitColumnExprLiteral(ctx *ColumnExprLiteralContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnExprBraceParam.
	VisitColumnExprBraceParam(ctx *ColumnExprBraceParamContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnExprArray.
	VisitColumnExprArray(ctx *ColumnExprArrayContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnExprSubstring.
	VisitColumnExprSubstring(ctx *ColumnExprSubstringContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnExprCast.
	VisitColumnExprCast(ctx *ColumnExprCastContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnExprOr.
	VisitColumnExprOr(ctx *ColumnExprOrContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnExprPrecedence1.
	VisitColumnExprPrecedence1(ctx *ColumnExprPrecedence1Context) interface{}

	// Visit a parse tree produced by DPParser#ColumnExprPrecedence2.
	VisitColumnExprPrecedence2(ctx *ColumnExprPrecedence2Context) interface{}

	// Visit a parse tree produced by DPParser#ColumnExprPrecedence3.
	VisitColumnExprPrecedence3(ctx *ColumnExprPrecedence3Context) interface{}

	// Visit a parse tree produced by DPParser#ColumnExprInterval.
	VisitColumnExprInterval(ctx *ColumnExprIntervalContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnExprIsNull.
	VisitColumnExprIsNull(ctx *ColumnExprIsNullContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnExprWinFunctionTarget.
	VisitColumnExprWinFunctionTarget(ctx *ColumnExprWinFunctionTargetContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnExprTrim.
	VisitColumnExprTrim(ctx *ColumnExprTrimContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnExprTuple.
	VisitColumnExprTuple(ctx *ColumnExprTupleContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnExprArrayAccess.
	VisitColumnExprArrayAccess(ctx *ColumnExprArrayAccessContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnExprBetween.
	VisitColumnExprBetween(ctx *ColumnExprBetweenContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnExprParens.
	VisitColumnExprParens(ctx *ColumnExprParensContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnExprTimestamp.
	VisitColumnExprTimestamp(ctx *ColumnExprTimestampContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnExprAnd.
	VisitColumnExprAnd(ctx *ColumnExprAndContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnExprTupleAccess.
	VisitColumnExprTupleAccess(ctx *ColumnExprTupleAccessContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnExprCase.
	VisitColumnExprCase(ctx *ColumnExprCaseContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnExprDate.
	VisitColumnExprDate(ctx *ColumnExprDateContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnExprNot.
	VisitColumnExprNot(ctx *ColumnExprNotContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnExprWinFunction.
	VisitColumnExprWinFunction(ctx *ColumnExprWinFunctionContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnExprIdentifier.
	VisitColumnExprIdentifier(ctx *ColumnExprIdentifierContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnExprFunction.
	VisitColumnExprFunction(ctx *ColumnExprFunctionContext) interface{}

	// Visit a parse tree produced by DPParser#ColumnExprAsterisk.
	VisitColumnExprAsterisk(ctx *ColumnExprAsteriskContext) interface{}

	// Visit a parse tree produced by DPParser#columnArgList.
	VisitColumnArgList(ctx *ColumnArgListContext) interface{}

	// Visit a parse tree produced by DPParser#columnArgExpr.
	VisitColumnArgExpr(ctx *ColumnArgExprContext) interface{}

	// Visit a parse tree produced by DPParser#columnLambdaExpr.
	VisitColumnLambdaExpr(ctx *ColumnLambdaExprContext) interface{}

	// Visit a parse tree produced by DPParser#columnIdentifier.
	VisitColumnIdentifier(ctx *ColumnIdentifierContext) interface{}

	// Visit a parse tree produced by DPParser#nestedIdentifier.
	VisitNestedIdentifier(ctx *NestedIdentifierContext) interface{}

	// Visit a parse tree produced by DPParser#TableExprIdentifier.
	VisitTableExprIdentifier(ctx *TableExprIdentifierContext) interface{}

	// Visit a parse tree produced by DPParser#TableExprSubquery.
	VisitTableExprSubquery(ctx *TableExprSubqueryContext) interface{}

	// Visit a parse tree produced by DPParser#TableExprAlias.
	VisitTableExprAlias(ctx *TableExprAliasContext) interface{}

	// Visit a parse tree produced by DPParser#TableExprFunction.
	VisitTableExprFunction(ctx *TableExprFunctionContext) interface{}

	// Visit a parse tree produced by DPParser#tableFunctionExpr.
	VisitTableFunctionExpr(ctx *TableFunctionExprContext) interface{}

	// Visit a parse tree produced by DPParser#tableIdentifier.
	VisitTableIdentifier(ctx *TableIdentifierContext) interface{}

	// Visit a parse tree produced by DPParser#tableArgList.
	VisitTableArgList(ctx *TableArgListContext) interface{}

	// Visit a parse tree produced by DPParser#tableArgExpr.
	VisitTableArgExpr(ctx *TableArgExprContext) interface{}

	// Visit a parse tree produced by DPParser#databaseIdentifier.
	VisitDatabaseIdentifier(ctx *DatabaseIdentifierContext) interface{}

	// Visit a parse tree produced by DPParser#floatingLiteral.
	VisitFloatingLiteral(ctx *FloatingLiteralContext) interface{}

	// Visit a parse tree produced by DPParser#numberLiteral.
	VisitNumberLiteral(ctx *NumberLiteralContext) interface{}

	// Visit a parse tree produced by DPParser#literal.
	VisitLiteral(ctx *LiteralContext) interface{}

	// Visit a parse tree produced by DPParser#interval.
	VisitInterval(ctx *IntervalContext) interface{}

	// Visit a parse tree produced by DPParser#keyword.
	VisitKeyword(ctx *KeywordContext) interface{}

	// Visit a parse tree produced by DPParser#keywordForAlias.
	VisitKeywordForAlias(ctx *KeywordForAliasContext) interface{}

	// Visit a parse tree produced by DPParser#alias.
	VisitAlias(ctx *AliasContext) interface{}

	// Visit a parse tree produced by DPParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by DPParser#identifierOrNull.
	VisitIdentifierOrNull(ctx *IdentifierOrNullContext) interface{}

	// Visit a parse tree produced by DPParser#enumValue.
	VisitEnumValue(ctx *EnumValueContext) interface{}
}
