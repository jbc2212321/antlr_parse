// Code generated from /Users/bytedance/go/src/code.byted.org/lark-approval/ai_sdk/clickhouse_parse/DP.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // DP

import "github.com/antlr4-go/antlr/v4"

// BaseDPListener is a complete listener for a parse tree produced by DPParser.
type BaseDPListener struct{}

var _ DPListener = &BaseDPListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseDPListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseDPListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseDPListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseDPListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQueryStmt is called when production queryStmt is entered.
func (s *BaseDPListener) EnterQueryStmt(ctx *QueryStmtContext) {}

// ExitQueryStmt is called when production queryStmt is exited.
func (s *BaseDPListener) ExitQueryStmt(ctx *QueryStmtContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseDPListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseDPListener) ExitQuery(ctx *QueryContext) {}

// EnterCtes is called when production ctes is entered.
func (s *BaseDPListener) EnterCtes(ctx *CtesContext) {}

// ExitCtes is called when production ctes is exited.
func (s *BaseDPListener) ExitCtes(ctx *CtesContext) {}

// EnterNamedQuery is called when production namedQuery is entered.
func (s *BaseDPListener) EnterNamedQuery(ctx *NamedQueryContext) {}

// ExitNamedQuery is called when production namedQuery is exited.
func (s *BaseDPListener) ExitNamedQuery(ctx *NamedQueryContext) {}

// EnterColumnAliases is called when production columnAliases is entered.
func (s *BaseDPListener) EnterColumnAliases(ctx *ColumnAliasesContext) {}

// ExitColumnAliases is called when production columnAliases is exited.
func (s *BaseDPListener) ExitColumnAliases(ctx *ColumnAliasesContext) {}

// EnterSelectUnionStmt is called when production selectUnionStmt is entered.
func (s *BaseDPListener) EnterSelectUnionStmt(ctx *SelectUnionStmtContext) {}

// ExitSelectUnionStmt is called when production selectUnionStmt is exited.
func (s *BaseDPListener) ExitSelectUnionStmt(ctx *SelectUnionStmtContext) {}

// EnterSelectStmtWithParens is called when production selectStmtWithParens is entered.
func (s *BaseDPListener) EnterSelectStmtWithParens(ctx *SelectStmtWithParensContext) {}

// ExitSelectStmtWithParens is called when production selectStmtWithParens is exited.
func (s *BaseDPListener) ExitSelectStmtWithParens(ctx *SelectStmtWithParensContext) {}

// EnterSelectStmt is called when production selectStmt is entered.
func (s *BaseDPListener) EnterSelectStmt(ctx *SelectStmtContext) {}

// ExitSelectStmt is called when production selectStmt is exited.
func (s *BaseDPListener) ExitSelectStmt(ctx *SelectStmtContext) {}

// EnterWithClause is called when production withClause is entered.
func (s *BaseDPListener) EnterWithClause(ctx *WithClauseContext) {}

// ExitWithClause is called when production withClause is exited.
func (s *BaseDPListener) ExitWithClause(ctx *WithClauseContext) {}

// EnterTopClause is called when production topClause is entered.
func (s *BaseDPListener) EnterTopClause(ctx *TopClauseContext) {}

// ExitTopClause is called when production topClause is exited.
func (s *BaseDPListener) ExitTopClause(ctx *TopClauseContext) {}

// EnterFromClause is called when production fromClause is entered.
func (s *BaseDPListener) EnterFromClause(ctx *FromClauseContext) {}

// ExitFromClause is called when production fromClause is exited.
func (s *BaseDPListener) ExitFromClause(ctx *FromClauseContext) {}

// EnterArrayJoinClause is called when production arrayJoinClause is entered.
func (s *BaseDPListener) EnterArrayJoinClause(ctx *ArrayJoinClauseContext) {}

// ExitArrayJoinClause is called when production arrayJoinClause is exited.
func (s *BaseDPListener) ExitArrayJoinClause(ctx *ArrayJoinClauseContext) {}

// EnterWindowClause is called when production windowClause is entered.
func (s *BaseDPListener) EnterWindowClause(ctx *WindowClauseContext) {}

// ExitWindowClause is called when production windowClause is exited.
func (s *BaseDPListener) ExitWindowClause(ctx *WindowClauseContext) {}

// EnterPrewhereClause is called when production prewhereClause is entered.
func (s *BaseDPListener) EnterPrewhereClause(ctx *PrewhereClauseContext) {}

// ExitPrewhereClause is called when production prewhereClause is exited.
func (s *BaseDPListener) ExitPrewhereClause(ctx *PrewhereClauseContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BaseDPListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BaseDPListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterGroupByClause is called when production groupByClause is entered.
func (s *BaseDPListener) EnterGroupByClause(ctx *GroupByClauseContext) {}

// ExitGroupByClause is called when production groupByClause is exited.
func (s *BaseDPListener) ExitGroupByClause(ctx *GroupByClauseContext) {}

// EnterHavingClause is called when production havingClause is entered.
func (s *BaseDPListener) EnterHavingClause(ctx *HavingClauseContext) {}

// ExitHavingClause is called when production havingClause is exited.
func (s *BaseDPListener) ExitHavingClause(ctx *HavingClauseContext) {}

// EnterOrderByClause is called when production orderByClause is entered.
func (s *BaseDPListener) EnterOrderByClause(ctx *OrderByClauseContext) {}

// ExitOrderByClause is called when production orderByClause is exited.
func (s *BaseDPListener) ExitOrderByClause(ctx *OrderByClauseContext) {}

// EnterLimitByClause is called when production limitByClause is entered.
func (s *BaseDPListener) EnterLimitByClause(ctx *LimitByClauseContext) {}

// ExitLimitByClause is called when production limitByClause is exited.
func (s *BaseDPListener) ExitLimitByClause(ctx *LimitByClauseContext) {}

// EnterLimitClause is called when production limitClause is entered.
func (s *BaseDPListener) EnterLimitClause(ctx *LimitClauseContext) {}

// ExitLimitClause is called when production limitClause is exited.
func (s *BaseDPListener) ExitLimitClause(ctx *LimitClauseContext) {}

// EnterTealimitClause is called when production tealimitClause is entered.
func (s *BaseDPListener) EnterTealimitClause(ctx *TealimitClauseContext) {}

// ExitTealimitClause is called when production tealimitClause is exited.
func (s *BaseDPListener) ExitTealimitClause(ctx *TealimitClauseContext) {}

// EnterSettingsClause is called when production settingsClause is entered.
func (s *BaseDPListener) EnterSettingsClause(ctx *SettingsClauseContext) {}

// ExitSettingsClause is called when production settingsClause is exited.
func (s *BaseDPListener) ExitSettingsClause(ctx *SettingsClauseContext) {}

// EnterJoinExprOp is called when production JoinExprOp is entered.
func (s *BaseDPListener) EnterJoinExprOp(ctx *JoinExprOpContext) {}

// ExitJoinExprOp is called when production JoinExprOp is exited.
func (s *BaseDPListener) ExitJoinExprOp(ctx *JoinExprOpContext) {}

// EnterJoinExprTable is called when production JoinExprTable is entered.
func (s *BaseDPListener) EnterJoinExprTable(ctx *JoinExprTableContext) {}

// ExitJoinExprTable is called when production JoinExprTable is exited.
func (s *BaseDPListener) ExitJoinExprTable(ctx *JoinExprTableContext) {}

// EnterJoinExprParens is called when production JoinExprParens is entered.
func (s *BaseDPListener) EnterJoinExprParens(ctx *JoinExprParensContext) {}

// ExitJoinExprParens is called when production JoinExprParens is exited.
func (s *BaseDPListener) ExitJoinExprParens(ctx *JoinExprParensContext) {}

// EnterJoinOpInner is called when production JoinOpInner is entered.
func (s *BaseDPListener) EnterJoinOpInner(ctx *JoinOpInnerContext) {}

// ExitJoinOpInner is called when production JoinOpInner is exited.
func (s *BaseDPListener) ExitJoinOpInner(ctx *JoinOpInnerContext) {}

// EnterJoinOpLeftRight is called when production JoinOpLeftRight is entered.
func (s *BaseDPListener) EnterJoinOpLeftRight(ctx *JoinOpLeftRightContext) {}

// ExitJoinOpLeftRight is called when production JoinOpLeftRight is exited.
func (s *BaseDPListener) ExitJoinOpLeftRight(ctx *JoinOpLeftRightContext) {}

// EnterJoinOpFull is called when production JoinOpFull is entered.
func (s *BaseDPListener) EnterJoinOpFull(ctx *JoinOpFullContext) {}

// ExitJoinOpFull is called when production JoinOpFull is exited.
func (s *BaseDPListener) ExitJoinOpFull(ctx *JoinOpFullContext) {}

// EnterJoinOpCross is called when production joinOpCross is entered.
func (s *BaseDPListener) EnterJoinOpCross(ctx *JoinOpCrossContext) {}

// ExitJoinOpCross is called when production joinOpCross is exited.
func (s *BaseDPListener) ExitJoinOpCross(ctx *JoinOpCrossContext) {}

// EnterJoinConstraintClause is called when production joinConstraintClause is entered.
func (s *BaseDPListener) EnterJoinConstraintClause(ctx *JoinConstraintClauseContext) {}

// ExitJoinConstraintClause is called when production joinConstraintClause is exited.
func (s *BaseDPListener) ExitJoinConstraintClause(ctx *JoinConstraintClauseContext) {}

// EnterSampleClause is called when production sampleClause is entered.
func (s *BaseDPListener) EnterSampleClause(ctx *SampleClauseContext) {}

// ExitSampleClause is called when production sampleClause is exited.
func (s *BaseDPListener) ExitSampleClause(ctx *SampleClauseContext) {}

// EnterLimitExpr is called when production limitExpr is entered.
func (s *BaseDPListener) EnterLimitExpr(ctx *LimitExprContext) {}

// ExitLimitExpr is called when production limitExpr is exited.
func (s *BaseDPListener) ExitLimitExpr(ctx *LimitExprContext) {}

// EnterTealimitExpr is called when production tealimitExpr is entered.
func (s *BaseDPListener) EnterTealimitExpr(ctx *TealimitExprContext) {}

// ExitTealimitExpr is called when production tealimitExpr is exited.
func (s *BaseDPListener) ExitTealimitExpr(ctx *TealimitExprContext) {}

// EnterOrderExprList is called when production orderExprList is entered.
func (s *BaseDPListener) EnterOrderExprList(ctx *OrderExprListContext) {}

// ExitOrderExprList is called when production orderExprList is exited.
func (s *BaseDPListener) ExitOrderExprList(ctx *OrderExprListContext) {}

// EnterOrderExpr is called when production orderExpr is entered.
func (s *BaseDPListener) EnterOrderExpr(ctx *OrderExprContext) {}

// ExitOrderExpr is called when production orderExpr is exited.
func (s *BaseDPListener) ExitOrderExpr(ctx *OrderExprContext) {}

// EnterRatioExpr is called when production ratioExpr is entered.
func (s *BaseDPListener) EnterRatioExpr(ctx *RatioExprContext) {}

// ExitRatioExpr is called when production ratioExpr is exited.
func (s *BaseDPListener) ExitRatioExpr(ctx *RatioExprContext) {}

// EnterSettingExprList is called when production settingExprList is entered.
func (s *BaseDPListener) EnterSettingExprList(ctx *SettingExprListContext) {}

// ExitSettingExprList is called when production settingExprList is exited.
func (s *BaseDPListener) ExitSettingExprList(ctx *SettingExprListContext) {}

// EnterSettingExpr is called when production settingExpr is entered.
func (s *BaseDPListener) EnterSettingExpr(ctx *SettingExprContext) {}

// ExitSettingExpr is called when production settingExpr is exited.
func (s *BaseDPListener) ExitSettingExpr(ctx *SettingExprContext) {}

// EnterWindowExpr is called when production windowExpr is entered.
func (s *BaseDPListener) EnterWindowExpr(ctx *WindowExprContext) {}

// ExitWindowExpr is called when production windowExpr is exited.
func (s *BaseDPListener) ExitWindowExpr(ctx *WindowExprContext) {}

// EnterWinPartitionByClause is called when production winPartitionByClause is entered.
func (s *BaseDPListener) EnterWinPartitionByClause(ctx *WinPartitionByClauseContext) {}

// ExitWinPartitionByClause is called when production winPartitionByClause is exited.
func (s *BaseDPListener) ExitWinPartitionByClause(ctx *WinPartitionByClauseContext) {}

// EnterWinOrderByClause is called when production winOrderByClause is entered.
func (s *BaseDPListener) EnterWinOrderByClause(ctx *WinOrderByClauseContext) {}

// ExitWinOrderByClause is called when production winOrderByClause is exited.
func (s *BaseDPListener) ExitWinOrderByClause(ctx *WinOrderByClauseContext) {}

// EnterWinFrameClause is called when production winFrameClause is entered.
func (s *BaseDPListener) EnterWinFrameClause(ctx *WinFrameClauseContext) {}

// ExitWinFrameClause is called when production winFrameClause is exited.
func (s *BaseDPListener) ExitWinFrameClause(ctx *WinFrameClauseContext) {}

// EnterFrameStart is called when production frameStart is entered.
func (s *BaseDPListener) EnterFrameStart(ctx *FrameStartContext) {}

// ExitFrameStart is called when production frameStart is exited.
func (s *BaseDPListener) ExitFrameStart(ctx *FrameStartContext) {}

// EnterFrameBetween is called when production frameBetween is entered.
func (s *BaseDPListener) EnterFrameBetween(ctx *FrameBetweenContext) {}

// ExitFrameBetween is called when production frameBetween is exited.
func (s *BaseDPListener) ExitFrameBetween(ctx *FrameBetweenContext) {}

// EnterWinFrameBound is called when production winFrameBound is entered.
func (s *BaseDPListener) EnterWinFrameBound(ctx *WinFrameBoundContext) {}

// ExitWinFrameBound is called when production winFrameBound is exited.
func (s *BaseDPListener) ExitWinFrameBound(ctx *WinFrameBoundContext) {}

// EnterColumnTypeExprSimple is called when production ColumnTypeExprSimple is entered.
func (s *BaseDPListener) EnterColumnTypeExprSimple(ctx *ColumnTypeExprSimpleContext) {}

// ExitColumnTypeExprSimple is called when production ColumnTypeExprSimple is exited.
func (s *BaseDPListener) ExitColumnTypeExprSimple(ctx *ColumnTypeExprSimpleContext) {}

// EnterColumnTypeExprNested is called when production ColumnTypeExprNested is entered.
func (s *BaseDPListener) EnterColumnTypeExprNested(ctx *ColumnTypeExprNestedContext) {}

// ExitColumnTypeExprNested is called when production ColumnTypeExprNested is exited.
func (s *BaseDPListener) ExitColumnTypeExprNested(ctx *ColumnTypeExprNestedContext) {}

// EnterColumnTypeExprEnum is called when production ColumnTypeExprEnum is entered.
func (s *BaseDPListener) EnterColumnTypeExprEnum(ctx *ColumnTypeExprEnumContext) {}

// ExitColumnTypeExprEnum is called when production ColumnTypeExprEnum is exited.
func (s *BaseDPListener) ExitColumnTypeExprEnum(ctx *ColumnTypeExprEnumContext) {}

// EnterColumnTypeExprComplex is called when production ColumnTypeExprComplex is entered.
func (s *BaseDPListener) EnterColumnTypeExprComplex(ctx *ColumnTypeExprComplexContext) {}

// ExitColumnTypeExprComplex is called when production ColumnTypeExprComplex is exited.
func (s *BaseDPListener) ExitColumnTypeExprComplex(ctx *ColumnTypeExprComplexContext) {}

// EnterColumnTypeExprParam is called when production ColumnTypeExprParam is entered.
func (s *BaseDPListener) EnterColumnTypeExprParam(ctx *ColumnTypeExprParamContext) {}

// ExitColumnTypeExprParam is called when production ColumnTypeExprParam is exited.
func (s *BaseDPListener) ExitColumnTypeExprParam(ctx *ColumnTypeExprParamContext) {}

// EnterColumnExprList is called when production columnExprList is entered.
func (s *BaseDPListener) EnterColumnExprList(ctx *ColumnExprListContext) {}

// ExitColumnExprList is called when production columnExprList is exited.
func (s *BaseDPListener) ExitColumnExprList(ctx *ColumnExprListContext) {}

// EnterColumnsExprAsterisk is called when production ColumnsExprAsterisk is entered.
func (s *BaseDPListener) EnterColumnsExprAsterisk(ctx *ColumnsExprAsteriskContext) {}

// ExitColumnsExprAsterisk is called when production ColumnsExprAsterisk is exited.
func (s *BaseDPListener) ExitColumnsExprAsterisk(ctx *ColumnsExprAsteriskContext) {}

// EnterColumnsExprSubquery is called when production ColumnsExprSubquery is entered.
func (s *BaseDPListener) EnterColumnsExprSubquery(ctx *ColumnsExprSubqueryContext) {}

// ExitColumnsExprSubquery is called when production ColumnsExprSubquery is exited.
func (s *BaseDPListener) ExitColumnsExprSubquery(ctx *ColumnsExprSubqueryContext) {}

// EnterColumnsExprColumn is called when production ColumnsExprColumn is entered.
func (s *BaseDPListener) EnterColumnsExprColumn(ctx *ColumnsExprColumnContext) {}

// ExitColumnsExprColumn is called when production ColumnsExprColumn is exited.
func (s *BaseDPListener) ExitColumnsExprColumn(ctx *ColumnsExprColumnContext) {}

// EnterColumnExprTernaryOp is called when production ColumnExprTernaryOp is entered.
func (s *BaseDPListener) EnterColumnExprTernaryOp(ctx *ColumnExprTernaryOpContext) {}

// ExitColumnExprTernaryOp is called when production ColumnExprTernaryOp is exited.
func (s *BaseDPListener) ExitColumnExprTernaryOp(ctx *ColumnExprTernaryOpContext) {}

// EnterColumnExprAlias is called when production ColumnExprAlias is entered.
func (s *BaseDPListener) EnterColumnExprAlias(ctx *ColumnExprAliasContext) {}

// ExitColumnExprAlias is called when production ColumnExprAlias is exited.
func (s *BaseDPListener) ExitColumnExprAlias(ctx *ColumnExprAliasContext) {}

// EnterColumnExprExtract is called when production ColumnExprExtract is entered.
func (s *BaseDPListener) EnterColumnExprExtract(ctx *ColumnExprExtractContext) {}

// ExitColumnExprExtract is called when production ColumnExprExtract is exited.
func (s *BaseDPListener) ExitColumnExprExtract(ctx *ColumnExprExtractContext) {}

// EnterColumnExprNegate is called when production ColumnExprNegate is entered.
func (s *BaseDPListener) EnterColumnExprNegate(ctx *ColumnExprNegateContext) {}

// ExitColumnExprNegate is called when production ColumnExprNegate is exited.
func (s *BaseDPListener) ExitColumnExprNegate(ctx *ColumnExprNegateContext) {}

// EnterColumnExprSubquery is called when production ColumnExprSubquery is entered.
func (s *BaseDPListener) EnterColumnExprSubquery(ctx *ColumnExprSubqueryContext) {}

// ExitColumnExprSubquery is called when production ColumnExprSubquery is exited.
func (s *BaseDPListener) ExitColumnExprSubquery(ctx *ColumnExprSubqueryContext) {}

// EnterColumnExprLiteral is called when production ColumnExprLiteral is entered.
func (s *BaseDPListener) EnterColumnExprLiteral(ctx *ColumnExprLiteralContext) {}

// ExitColumnExprLiteral is called when production ColumnExprLiteral is exited.
func (s *BaseDPListener) ExitColumnExprLiteral(ctx *ColumnExprLiteralContext) {}

// EnterColumnExprBraceParam is called when production ColumnExprBraceParam is entered.
func (s *BaseDPListener) EnterColumnExprBraceParam(ctx *ColumnExprBraceParamContext) {}

// ExitColumnExprBraceParam is called when production ColumnExprBraceParam is exited.
func (s *BaseDPListener) ExitColumnExprBraceParam(ctx *ColumnExprBraceParamContext) {}

// EnterColumnExprArray is called when production ColumnExprArray is entered.
func (s *BaseDPListener) EnterColumnExprArray(ctx *ColumnExprArrayContext) {}

// ExitColumnExprArray is called when production ColumnExprArray is exited.
func (s *BaseDPListener) ExitColumnExprArray(ctx *ColumnExprArrayContext) {}

// EnterColumnExprSubstring is called when production ColumnExprSubstring is entered.
func (s *BaseDPListener) EnterColumnExprSubstring(ctx *ColumnExprSubstringContext) {}

// ExitColumnExprSubstring is called when production ColumnExprSubstring is exited.
func (s *BaseDPListener) ExitColumnExprSubstring(ctx *ColumnExprSubstringContext) {}

// EnterColumnExprCast is called when production ColumnExprCast is entered.
func (s *BaseDPListener) EnterColumnExprCast(ctx *ColumnExprCastContext) {}

// ExitColumnExprCast is called when production ColumnExprCast is exited.
func (s *BaseDPListener) ExitColumnExprCast(ctx *ColumnExprCastContext) {}

// EnterColumnExprOr is called when production ColumnExprOr is entered.
func (s *BaseDPListener) EnterColumnExprOr(ctx *ColumnExprOrContext) {}

// ExitColumnExprOr is called when production ColumnExprOr is exited.
func (s *BaseDPListener) ExitColumnExprOr(ctx *ColumnExprOrContext) {}

// EnterColumnExprPrecedence1 is called when production ColumnExprPrecedence1 is entered.
func (s *BaseDPListener) EnterColumnExprPrecedence1(ctx *ColumnExprPrecedence1Context) {}

// ExitColumnExprPrecedence1 is called when production ColumnExprPrecedence1 is exited.
func (s *BaseDPListener) ExitColumnExprPrecedence1(ctx *ColumnExprPrecedence1Context) {}

// EnterColumnExprPrecedence2 is called when production ColumnExprPrecedence2 is entered.
func (s *BaseDPListener) EnterColumnExprPrecedence2(ctx *ColumnExprPrecedence2Context) {}

// ExitColumnExprPrecedence2 is called when production ColumnExprPrecedence2 is exited.
func (s *BaseDPListener) ExitColumnExprPrecedence2(ctx *ColumnExprPrecedence2Context) {}

// EnterColumnExprPrecedence3 is called when production ColumnExprPrecedence3 is entered.
func (s *BaseDPListener) EnterColumnExprPrecedence3(ctx *ColumnExprPrecedence3Context) {}

// ExitColumnExprPrecedence3 is called when production ColumnExprPrecedence3 is exited.
func (s *BaseDPListener) ExitColumnExprPrecedence3(ctx *ColumnExprPrecedence3Context) {}

// EnterColumnExprInterval is called when production ColumnExprInterval is entered.
func (s *BaseDPListener) EnterColumnExprInterval(ctx *ColumnExprIntervalContext) {}

// ExitColumnExprInterval is called when production ColumnExprInterval is exited.
func (s *BaseDPListener) ExitColumnExprInterval(ctx *ColumnExprIntervalContext) {}

// EnterColumnExprIsNull is called when production ColumnExprIsNull is entered.
func (s *BaseDPListener) EnterColumnExprIsNull(ctx *ColumnExprIsNullContext) {}

// ExitColumnExprIsNull is called when production ColumnExprIsNull is exited.
func (s *BaseDPListener) ExitColumnExprIsNull(ctx *ColumnExprIsNullContext) {}

// EnterColumnExprWinFunctionTarget is called when production ColumnExprWinFunctionTarget is entered.
func (s *BaseDPListener) EnterColumnExprWinFunctionTarget(ctx *ColumnExprWinFunctionTargetContext) {}

// ExitColumnExprWinFunctionTarget is called when production ColumnExprWinFunctionTarget is exited.
func (s *BaseDPListener) ExitColumnExprWinFunctionTarget(ctx *ColumnExprWinFunctionTargetContext) {}

// EnterColumnExprTrim is called when production ColumnExprTrim is entered.
func (s *BaseDPListener) EnterColumnExprTrim(ctx *ColumnExprTrimContext) {}

// ExitColumnExprTrim is called when production ColumnExprTrim is exited.
func (s *BaseDPListener) ExitColumnExprTrim(ctx *ColumnExprTrimContext) {}

// EnterColumnExprTuple is called when production ColumnExprTuple is entered.
func (s *BaseDPListener) EnterColumnExprTuple(ctx *ColumnExprTupleContext) {}

// ExitColumnExprTuple is called when production ColumnExprTuple is exited.
func (s *BaseDPListener) ExitColumnExprTuple(ctx *ColumnExprTupleContext) {}

// EnterColumnExprArrayAccess is called when production ColumnExprArrayAccess is entered.
func (s *BaseDPListener) EnterColumnExprArrayAccess(ctx *ColumnExprArrayAccessContext) {}

// ExitColumnExprArrayAccess is called when production ColumnExprArrayAccess is exited.
func (s *BaseDPListener) ExitColumnExprArrayAccess(ctx *ColumnExprArrayAccessContext) {}

// EnterColumnExprBetween is called when production ColumnExprBetween is entered.
func (s *BaseDPListener) EnterColumnExprBetween(ctx *ColumnExprBetweenContext) {}

// ExitColumnExprBetween is called when production ColumnExprBetween is exited.
func (s *BaseDPListener) ExitColumnExprBetween(ctx *ColumnExprBetweenContext) {}

// EnterColumnExprParens is called when production ColumnExprParens is entered.
func (s *BaseDPListener) EnterColumnExprParens(ctx *ColumnExprParensContext) {}

// ExitColumnExprParens is called when production ColumnExprParens is exited.
func (s *BaseDPListener) ExitColumnExprParens(ctx *ColumnExprParensContext) {}

// EnterColumnExprTimestamp is called when production ColumnExprTimestamp is entered.
func (s *BaseDPListener) EnterColumnExprTimestamp(ctx *ColumnExprTimestampContext) {}

// ExitColumnExprTimestamp is called when production ColumnExprTimestamp is exited.
func (s *BaseDPListener) ExitColumnExprTimestamp(ctx *ColumnExprTimestampContext) {}

// EnterColumnExprAnd is called when production ColumnExprAnd is entered.
func (s *BaseDPListener) EnterColumnExprAnd(ctx *ColumnExprAndContext) {}

// ExitColumnExprAnd is called when production ColumnExprAnd is exited.
func (s *BaseDPListener) ExitColumnExprAnd(ctx *ColumnExprAndContext) {}

// EnterColumnExprTupleAccess is called when production ColumnExprTupleAccess is entered.
func (s *BaseDPListener) EnterColumnExprTupleAccess(ctx *ColumnExprTupleAccessContext) {}

// ExitColumnExprTupleAccess is called when production ColumnExprTupleAccess is exited.
func (s *BaseDPListener) ExitColumnExprTupleAccess(ctx *ColumnExprTupleAccessContext) {}

// EnterColumnExprCase is called when production ColumnExprCase is entered.
func (s *BaseDPListener) EnterColumnExprCase(ctx *ColumnExprCaseContext) {}

// ExitColumnExprCase is called when production ColumnExprCase is exited.
func (s *BaseDPListener) ExitColumnExprCase(ctx *ColumnExprCaseContext) {}

// EnterColumnExprDate is called when production ColumnExprDate is entered.
func (s *BaseDPListener) EnterColumnExprDate(ctx *ColumnExprDateContext) {}

// ExitColumnExprDate is called when production ColumnExprDate is exited.
func (s *BaseDPListener) ExitColumnExprDate(ctx *ColumnExprDateContext) {}

// EnterColumnExprNot is called when production ColumnExprNot is entered.
func (s *BaseDPListener) EnterColumnExprNot(ctx *ColumnExprNotContext) {}

// ExitColumnExprNot is called when production ColumnExprNot is exited.
func (s *BaseDPListener) ExitColumnExprNot(ctx *ColumnExprNotContext) {}

// EnterColumnExprWinFunction is called when production ColumnExprWinFunction is entered.
func (s *BaseDPListener) EnterColumnExprWinFunction(ctx *ColumnExprWinFunctionContext) {}

// ExitColumnExprWinFunction is called when production ColumnExprWinFunction is exited.
func (s *BaseDPListener) ExitColumnExprWinFunction(ctx *ColumnExprWinFunctionContext) {}

// EnterColumnExprIdentifier is called when production ColumnExprIdentifier is entered.
func (s *BaseDPListener) EnterColumnExprIdentifier(ctx *ColumnExprIdentifierContext) {}

// ExitColumnExprIdentifier is called when production ColumnExprIdentifier is exited.
func (s *BaseDPListener) ExitColumnExprIdentifier(ctx *ColumnExprIdentifierContext) {}

// EnterColumnExprFunction is called when production ColumnExprFunction is entered.
func (s *BaseDPListener) EnterColumnExprFunction(ctx *ColumnExprFunctionContext) {}

// ExitColumnExprFunction is called when production ColumnExprFunction is exited.
func (s *BaseDPListener) ExitColumnExprFunction(ctx *ColumnExprFunctionContext) {}

// EnterColumnExprAsterisk is called when production ColumnExprAsterisk is entered.
func (s *BaseDPListener) EnterColumnExprAsterisk(ctx *ColumnExprAsteriskContext) {}

// ExitColumnExprAsterisk is called when production ColumnExprAsterisk is exited.
func (s *BaseDPListener) ExitColumnExprAsterisk(ctx *ColumnExprAsteriskContext) {}

// EnterColumnArgList is called when production columnArgList is entered.
func (s *BaseDPListener) EnterColumnArgList(ctx *ColumnArgListContext) {}

// ExitColumnArgList is called when production columnArgList is exited.
func (s *BaseDPListener) ExitColumnArgList(ctx *ColumnArgListContext) {}

// EnterColumnArgExpr is called when production columnArgExpr is entered.
func (s *BaseDPListener) EnterColumnArgExpr(ctx *ColumnArgExprContext) {}

// ExitColumnArgExpr is called when production columnArgExpr is exited.
func (s *BaseDPListener) ExitColumnArgExpr(ctx *ColumnArgExprContext) {}

// EnterColumnLambdaExpr is called when production columnLambdaExpr is entered.
func (s *BaseDPListener) EnterColumnLambdaExpr(ctx *ColumnLambdaExprContext) {}

// ExitColumnLambdaExpr is called when production columnLambdaExpr is exited.
func (s *BaseDPListener) ExitColumnLambdaExpr(ctx *ColumnLambdaExprContext) {}

// EnterColumnIdentifier is called when production columnIdentifier is entered.
func (s *BaseDPListener) EnterColumnIdentifier(ctx *ColumnIdentifierContext) {}

// ExitColumnIdentifier is called when production columnIdentifier is exited.
func (s *BaseDPListener) ExitColumnIdentifier(ctx *ColumnIdentifierContext) {}

// EnterNestedIdentifier is called when production nestedIdentifier is entered.
func (s *BaseDPListener) EnterNestedIdentifier(ctx *NestedIdentifierContext) {}

// ExitNestedIdentifier is called when production nestedIdentifier is exited.
func (s *BaseDPListener) ExitNestedIdentifier(ctx *NestedIdentifierContext) {}

// EnterTableExprIdentifier is called when production TableExprIdentifier is entered.
func (s *BaseDPListener) EnterTableExprIdentifier(ctx *TableExprIdentifierContext) {}

// ExitTableExprIdentifier is called when production TableExprIdentifier is exited.
func (s *BaseDPListener) ExitTableExprIdentifier(ctx *TableExprIdentifierContext) {}

// EnterTableExprSubquery is called when production TableExprSubquery is entered.
func (s *BaseDPListener) EnterTableExprSubquery(ctx *TableExprSubqueryContext) {}

// ExitTableExprSubquery is called when production TableExprSubquery is exited.
func (s *BaseDPListener) ExitTableExprSubquery(ctx *TableExprSubqueryContext) {}

// EnterTableExprAlias is called when production TableExprAlias is entered.
func (s *BaseDPListener) EnterTableExprAlias(ctx *TableExprAliasContext) {}

// ExitTableExprAlias is called when production TableExprAlias is exited.
func (s *BaseDPListener) ExitTableExprAlias(ctx *TableExprAliasContext) {}

// EnterTableExprFunction is called when production TableExprFunction is entered.
func (s *BaseDPListener) EnterTableExprFunction(ctx *TableExprFunctionContext) {}

// ExitTableExprFunction is called when production TableExprFunction is exited.
func (s *BaseDPListener) ExitTableExprFunction(ctx *TableExprFunctionContext) {}

// EnterTableFunctionExpr is called when production tableFunctionExpr is entered.
func (s *BaseDPListener) EnterTableFunctionExpr(ctx *TableFunctionExprContext) {}

// ExitTableFunctionExpr is called when production tableFunctionExpr is exited.
func (s *BaseDPListener) ExitTableFunctionExpr(ctx *TableFunctionExprContext) {}

// EnterTableIdentifier is called when production tableIdentifier is entered.
func (s *BaseDPListener) EnterTableIdentifier(ctx *TableIdentifierContext) {}

// ExitTableIdentifier is called when production tableIdentifier is exited.
func (s *BaseDPListener) ExitTableIdentifier(ctx *TableIdentifierContext) {}

// EnterTableArgList is called when production tableArgList is entered.
func (s *BaseDPListener) EnterTableArgList(ctx *TableArgListContext) {}

// ExitTableArgList is called when production tableArgList is exited.
func (s *BaseDPListener) ExitTableArgList(ctx *TableArgListContext) {}

// EnterTableArgExpr is called when production tableArgExpr is entered.
func (s *BaseDPListener) EnterTableArgExpr(ctx *TableArgExprContext) {}

// ExitTableArgExpr is called when production tableArgExpr is exited.
func (s *BaseDPListener) ExitTableArgExpr(ctx *TableArgExprContext) {}

// EnterDatabaseIdentifier is called when production databaseIdentifier is entered.
func (s *BaseDPListener) EnterDatabaseIdentifier(ctx *DatabaseIdentifierContext) {}

// ExitDatabaseIdentifier is called when production databaseIdentifier is exited.
func (s *BaseDPListener) ExitDatabaseIdentifier(ctx *DatabaseIdentifierContext) {}

// EnterFloatingLiteral is called when production floatingLiteral is entered.
func (s *BaseDPListener) EnterFloatingLiteral(ctx *FloatingLiteralContext) {}

// ExitFloatingLiteral is called when production floatingLiteral is exited.
func (s *BaseDPListener) ExitFloatingLiteral(ctx *FloatingLiteralContext) {}

// EnterNumberLiteral is called when production numberLiteral is entered.
func (s *BaseDPListener) EnterNumberLiteral(ctx *NumberLiteralContext) {}

// ExitNumberLiteral is called when production numberLiteral is exited.
func (s *BaseDPListener) ExitNumberLiteral(ctx *NumberLiteralContext) {}

// EnterLiteral is called when production literal is entered.
func (s *BaseDPListener) EnterLiteral(ctx *LiteralContext) {}

// ExitLiteral is called when production literal is exited.
func (s *BaseDPListener) ExitLiteral(ctx *LiteralContext) {}

// EnterInterval is called when production interval is entered.
func (s *BaseDPListener) EnterInterval(ctx *IntervalContext) {}

// ExitInterval is called when production interval is exited.
func (s *BaseDPListener) ExitInterval(ctx *IntervalContext) {}

// EnterKeyword is called when production keyword is entered.
func (s *BaseDPListener) EnterKeyword(ctx *KeywordContext) {}

// ExitKeyword is called when production keyword is exited.
func (s *BaseDPListener) ExitKeyword(ctx *KeywordContext) {}

// EnterKeywordForAlias is called when production keywordForAlias is entered.
func (s *BaseDPListener) EnterKeywordForAlias(ctx *KeywordForAliasContext) {}

// ExitKeywordForAlias is called when production keywordForAlias is exited.
func (s *BaseDPListener) ExitKeywordForAlias(ctx *KeywordForAliasContext) {}

// EnterAlias is called when production alias is entered.
func (s *BaseDPListener) EnterAlias(ctx *AliasContext) {}

// ExitAlias is called when production alias is exited.
func (s *BaseDPListener) ExitAlias(ctx *AliasContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseDPListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseDPListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterIdentifierOrNull is called when production identifierOrNull is entered.
func (s *BaseDPListener) EnterIdentifierOrNull(ctx *IdentifierOrNullContext) {}

// ExitIdentifierOrNull is called when production identifierOrNull is exited.
func (s *BaseDPListener) ExitIdentifierOrNull(ctx *IdentifierOrNullContext) {}

// EnterEnumValue is called when production enumValue is entered.
func (s *BaseDPListener) EnterEnumValue(ctx *EnumValueContext) {}

// ExitEnumValue is called when production enumValue is exited.
func (s *BaseDPListener) ExitEnumValue(ctx *EnumValueContext) {}
