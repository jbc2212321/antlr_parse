// Code generated from /Users/bytedance/go/src/code.byted.org/lark-approval/ai_sdk/clickhouse_parse/DP.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // DP

import "github.com/antlr4-go/antlr/v4"

// DPListener is a complete listener for a parse tree produced by DPParser.
type DPListener interface {
	antlr.ParseTreeListener

	// EnterQueryStmt is called when entering the queryStmt production.
	EnterQueryStmt(c *QueryStmtContext)

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterCtes is called when entering the ctes production.
	EnterCtes(c *CtesContext)

	// EnterNamedQuery is called when entering the namedQuery production.
	EnterNamedQuery(c *NamedQueryContext)

	// EnterColumnAliases is called when entering the columnAliases production.
	EnterColumnAliases(c *ColumnAliasesContext)

	// EnterSelectUnionStmt is called when entering the selectUnionStmt production.
	EnterSelectUnionStmt(c *SelectUnionStmtContext)

	// EnterSelectStmtWithParens is called when entering the selectStmtWithParens production.
	EnterSelectStmtWithParens(c *SelectStmtWithParensContext)

	// EnterSelectStmt is called when entering the selectStmt production.
	EnterSelectStmt(c *SelectStmtContext)

	// EnterWithClause is called when entering the withClause production.
	EnterWithClause(c *WithClauseContext)

	// EnterTopClause is called when entering the topClause production.
	EnterTopClause(c *TopClauseContext)

	// EnterFromClause is called when entering the fromClause production.
	EnterFromClause(c *FromClauseContext)

	// EnterArrayJoinClause is called when entering the arrayJoinClause production.
	EnterArrayJoinClause(c *ArrayJoinClauseContext)

	// EnterWindowClause is called when entering the windowClause production.
	EnterWindowClause(c *WindowClauseContext)

	// EnterPrewhereClause is called when entering the prewhereClause production.
	EnterPrewhereClause(c *PrewhereClauseContext)

	// EnterWhereClause is called when entering the whereClause production.
	EnterWhereClause(c *WhereClauseContext)

	// EnterGroupByClause is called when entering the groupByClause production.
	EnterGroupByClause(c *GroupByClauseContext)

	// EnterHavingClause is called when entering the havingClause production.
	EnterHavingClause(c *HavingClauseContext)

	// EnterOrderByClause is called when entering the orderByClause production.
	EnterOrderByClause(c *OrderByClauseContext)

	// EnterLimitByClause is called when entering the limitByClause production.
	EnterLimitByClause(c *LimitByClauseContext)

	// EnterLimitClause is called when entering the limitClause production.
	EnterLimitClause(c *LimitClauseContext)

	// EnterTealimitClause is called when entering the tealimitClause production.
	EnterTealimitClause(c *TealimitClauseContext)

	// EnterSettingsClause is called when entering the settingsClause production.
	EnterSettingsClause(c *SettingsClauseContext)

	// EnterJoinExprOp is called when entering the JoinExprOp production.
	EnterJoinExprOp(c *JoinExprOpContext)

	// EnterJoinExprTable is called when entering the JoinExprTable production.
	EnterJoinExprTable(c *JoinExprTableContext)

	// EnterJoinExprParens is called when entering the JoinExprParens production.
	EnterJoinExprParens(c *JoinExprParensContext)

	// EnterJoinOpInner is called when entering the JoinOpInner production.
	EnterJoinOpInner(c *JoinOpInnerContext)

	// EnterJoinOpLeftRight is called when entering the JoinOpLeftRight production.
	EnterJoinOpLeftRight(c *JoinOpLeftRightContext)

	// EnterJoinOpFull is called when entering the JoinOpFull production.
	EnterJoinOpFull(c *JoinOpFullContext)

	// EnterJoinOpCross is called when entering the joinOpCross production.
	EnterJoinOpCross(c *JoinOpCrossContext)

	// EnterJoinConstraintClause is called when entering the joinConstraintClause production.
	EnterJoinConstraintClause(c *JoinConstraintClauseContext)

	// EnterSampleClause is called when entering the sampleClause production.
	EnterSampleClause(c *SampleClauseContext)

	// EnterLimitExpr is called when entering the limitExpr production.
	EnterLimitExpr(c *LimitExprContext)

	// EnterTealimitExpr is called when entering the tealimitExpr production.
	EnterTealimitExpr(c *TealimitExprContext)

	// EnterOrderExprList is called when entering the orderExprList production.
	EnterOrderExprList(c *OrderExprListContext)

	// EnterOrderExpr is called when entering the orderExpr production.
	EnterOrderExpr(c *OrderExprContext)

	// EnterRatioExpr is called when entering the ratioExpr production.
	EnterRatioExpr(c *RatioExprContext)

	// EnterSettingExprList is called when entering the settingExprList production.
	EnterSettingExprList(c *SettingExprListContext)

	// EnterSettingExpr is called when entering the settingExpr production.
	EnterSettingExpr(c *SettingExprContext)

	// EnterWindowExpr is called when entering the windowExpr production.
	EnterWindowExpr(c *WindowExprContext)

	// EnterWinPartitionByClause is called when entering the winPartitionByClause production.
	EnterWinPartitionByClause(c *WinPartitionByClauseContext)

	// EnterWinOrderByClause is called when entering the winOrderByClause production.
	EnterWinOrderByClause(c *WinOrderByClauseContext)

	// EnterWinFrameClause is called when entering the winFrameClause production.
	EnterWinFrameClause(c *WinFrameClauseContext)

	// EnterFrameStart is called when entering the frameStart production.
	EnterFrameStart(c *FrameStartContext)

	// EnterFrameBetween is called when entering the frameBetween production.
	EnterFrameBetween(c *FrameBetweenContext)

	// EnterWinFrameBound is called when entering the winFrameBound production.
	EnterWinFrameBound(c *WinFrameBoundContext)

	// EnterColumnTypeExprSimple is called when entering the ColumnTypeExprSimple production.
	EnterColumnTypeExprSimple(c *ColumnTypeExprSimpleContext)

	// EnterColumnTypeExprNested is called when entering the ColumnTypeExprNested production.
	EnterColumnTypeExprNested(c *ColumnTypeExprNestedContext)

	// EnterColumnTypeExprEnum is called when entering the ColumnTypeExprEnum production.
	EnterColumnTypeExprEnum(c *ColumnTypeExprEnumContext)

	// EnterColumnTypeExprComplex is called when entering the ColumnTypeExprComplex production.
	EnterColumnTypeExprComplex(c *ColumnTypeExprComplexContext)

	// EnterColumnTypeExprParam is called when entering the ColumnTypeExprParam production.
	EnterColumnTypeExprParam(c *ColumnTypeExprParamContext)

	// EnterColumnExprList is called when entering the columnExprList production.
	EnterColumnExprList(c *ColumnExprListContext)

	// EnterColumnsExprAsterisk is called when entering the ColumnsExprAsterisk production.
	EnterColumnsExprAsterisk(c *ColumnsExprAsteriskContext)

	// EnterColumnsExprSubquery is called when entering the ColumnsExprSubquery production.
	EnterColumnsExprSubquery(c *ColumnsExprSubqueryContext)

	// EnterColumnsExprColumn is called when entering the ColumnsExprColumn production.
	EnterColumnsExprColumn(c *ColumnsExprColumnContext)

	// EnterColumnExprTernaryOp is called when entering the ColumnExprTernaryOp production.
	EnterColumnExprTernaryOp(c *ColumnExprTernaryOpContext)

	// EnterColumnExprAlias is called when entering the ColumnExprAlias production.
	EnterColumnExprAlias(c *ColumnExprAliasContext)

	// EnterColumnExprExtract is called when entering the ColumnExprExtract production.
	EnterColumnExprExtract(c *ColumnExprExtractContext)

	// EnterColumnExprNegate is called when entering the ColumnExprNegate production.
	EnterColumnExprNegate(c *ColumnExprNegateContext)

	// EnterColumnExprSubquery is called when entering the ColumnExprSubquery production.
	EnterColumnExprSubquery(c *ColumnExprSubqueryContext)

	// EnterColumnExprLiteral is called when entering the ColumnExprLiteral production.
	EnterColumnExprLiteral(c *ColumnExprLiteralContext)

	// EnterColumnExprBraceParam is called when entering the ColumnExprBraceParam production.
	EnterColumnExprBraceParam(c *ColumnExprBraceParamContext)

	// EnterColumnExprArray is called when entering the ColumnExprArray production.
	EnterColumnExprArray(c *ColumnExprArrayContext)

	// EnterColumnExprSubstring is called when entering the ColumnExprSubstring production.
	EnterColumnExprSubstring(c *ColumnExprSubstringContext)

	// EnterColumnExprCast is called when entering the ColumnExprCast production.
	EnterColumnExprCast(c *ColumnExprCastContext)

	// EnterColumnExprOr is called when entering the ColumnExprOr production.
	EnterColumnExprOr(c *ColumnExprOrContext)

	// EnterColumnExprPrecedence1 is called when entering the ColumnExprPrecedence1 production.
	EnterColumnExprPrecedence1(c *ColumnExprPrecedence1Context)

	// EnterColumnExprPrecedence2 is called when entering the ColumnExprPrecedence2 production.
	EnterColumnExprPrecedence2(c *ColumnExprPrecedence2Context)

	// EnterColumnExprPrecedence3 is called when entering the ColumnExprPrecedence3 production.
	EnterColumnExprPrecedence3(c *ColumnExprPrecedence3Context)

	// EnterColumnExprInterval is called when entering the ColumnExprInterval production.
	EnterColumnExprInterval(c *ColumnExprIntervalContext)

	// EnterColumnExprIsNull is called when entering the ColumnExprIsNull production.
	EnterColumnExprIsNull(c *ColumnExprIsNullContext)

	// EnterColumnExprWinFunctionTarget is called when entering the ColumnExprWinFunctionTarget production.
	EnterColumnExprWinFunctionTarget(c *ColumnExprWinFunctionTargetContext)

	// EnterColumnExprTrim is called when entering the ColumnExprTrim production.
	EnterColumnExprTrim(c *ColumnExprTrimContext)

	// EnterColumnExprTuple is called when entering the ColumnExprTuple production.
	EnterColumnExprTuple(c *ColumnExprTupleContext)

	// EnterColumnExprArrayAccess is called when entering the ColumnExprArrayAccess production.
	EnterColumnExprArrayAccess(c *ColumnExprArrayAccessContext)

	// EnterColumnExprBetween is called when entering the ColumnExprBetween production.
	EnterColumnExprBetween(c *ColumnExprBetweenContext)

	// EnterColumnExprParens is called when entering the ColumnExprParens production.
	EnterColumnExprParens(c *ColumnExprParensContext)

	// EnterColumnExprTimestamp is called when entering the ColumnExprTimestamp production.
	EnterColumnExprTimestamp(c *ColumnExprTimestampContext)

	// EnterColumnExprAnd is called when entering the ColumnExprAnd production.
	EnterColumnExprAnd(c *ColumnExprAndContext)

	// EnterColumnExprTupleAccess is called when entering the ColumnExprTupleAccess production.
	EnterColumnExprTupleAccess(c *ColumnExprTupleAccessContext)

	// EnterColumnExprCase is called when entering the ColumnExprCase production.
	EnterColumnExprCase(c *ColumnExprCaseContext)

	// EnterColumnExprDate is called when entering the ColumnExprDate production.
	EnterColumnExprDate(c *ColumnExprDateContext)

	// EnterColumnExprNot is called when entering the ColumnExprNot production.
	EnterColumnExprNot(c *ColumnExprNotContext)

	// EnterColumnExprWinFunction is called when entering the ColumnExprWinFunction production.
	EnterColumnExprWinFunction(c *ColumnExprWinFunctionContext)

	// EnterColumnExprIdentifier is called when entering the ColumnExprIdentifier production.
	EnterColumnExprIdentifier(c *ColumnExprIdentifierContext)

	// EnterColumnExprFunction is called when entering the ColumnExprFunction production.
	EnterColumnExprFunction(c *ColumnExprFunctionContext)

	// EnterColumnExprAsterisk is called when entering the ColumnExprAsterisk production.
	EnterColumnExprAsterisk(c *ColumnExprAsteriskContext)

	// EnterColumnArgList is called when entering the columnArgList production.
	EnterColumnArgList(c *ColumnArgListContext)

	// EnterColumnArgExpr is called when entering the columnArgExpr production.
	EnterColumnArgExpr(c *ColumnArgExprContext)

	// EnterColumnLambdaExpr is called when entering the columnLambdaExpr production.
	EnterColumnLambdaExpr(c *ColumnLambdaExprContext)

	// EnterColumnIdentifier is called when entering the columnIdentifier production.
	EnterColumnIdentifier(c *ColumnIdentifierContext)

	// EnterNestedIdentifier is called when entering the nestedIdentifier production.
	EnterNestedIdentifier(c *NestedIdentifierContext)

	// EnterTableExprIdentifier is called when entering the TableExprIdentifier production.
	EnterTableExprIdentifier(c *TableExprIdentifierContext)

	// EnterTableExprSubquery is called when entering the TableExprSubquery production.
	EnterTableExprSubquery(c *TableExprSubqueryContext)

	// EnterTableExprAlias is called when entering the TableExprAlias production.
	EnterTableExprAlias(c *TableExprAliasContext)

	// EnterTableExprFunction is called when entering the TableExprFunction production.
	EnterTableExprFunction(c *TableExprFunctionContext)

	// EnterTableFunctionExpr is called when entering the tableFunctionExpr production.
	EnterTableFunctionExpr(c *TableFunctionExprContext)

	// EnterTableIdentifier is called when entering the tableIdentifier production.
	EnterTableIdentifier(c *TableIdentifierContext)

	// EnterTableArgList is called when entering the tableArgList production.
	EnterTableArgList(c *TableArgListContext)

	// EnterTableArgExpr is called when entering the tableArgExpr production.
	EnterTableArgExpr(c *TableArgExprContext)

	// EnterDatabaseIdentifier is called when entering the databaseIdentifier production.
	EnterDatabaseIdentifier(c *DatabaseIdentifierContext)

	// EnterFloatingLiteral is called when entering the floatingLiteral production.
	EnterFloatingLiteral(c *FloatingLiteralContext)

	// EnterNumberLiteral is called when entering the numberLiteral production.
	EnterNumberLiteral(c *NumberLiteralContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterInterval is called when entering the interval production.
	EnterInterval(c *IntervalContext)

	// EnterKeyword is called when entering the keyword production.
	EnterKeyword(c *KeywordContext)

	// EnterKeywordForAlias is called when entering the keywordForAlias production.
	EnterKeywordForAlias(c *KeywordForAliasContext)

	// EnterAlias is called when entering the alias production.
	EnterAlias(c *AliasContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterIdentifierOrNull is called when entering the identifierOrNull production.
	EnterIdentifierOrNull(c *IdentifierOrNullContext)

	// EnterEnumValue is called when entering the enumValue production.
	EnterEnumValue(c *EnumValueContext)

	// ExitQueryStmt is called when exiting the queryStmt production.
	ExitQueryStmt(c *QueryStmtContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitCtes is called when exiting the ctes production.
	ExitCtes(c *CtesContext)

	// ExitNamedQuery is called when exiting the namedQuery production.
	ExitNamedQuery(c *NamedQueryContext)

	// ExitColumnAliases is called when exiting the columnAliases production.
	ExitColumnAliases(c *ColumnAliasesContext)

	// ExitSelectUnionStmt is called when exiting the selectUnionStmt production.
	ExitSelectUnionStmt(c *SelectUnionStmtContext)

	// ExitSelectStmtWithParens is called when exiting the selectStmtWithParens production.
	ExitSelectStmtWithParens(c *SelectStmtWithParensContext)

	// ExitSelectStmt is called when exiting the selectStmt production.
	ExitSelectStmt(c *SelectStmtContext)

	// ExitWithClause is called when exiting the withClause production.
	ExitWithClause(c *WithClauseContext)

	// ExitTopClause is called when exiting the topClause production.
	ExitTopClause(c *TopClauseContext)

	// ExitFromClause is called when exiting the fromClause production.
	ExitFromClause(c *FromClauseContext)

	// ExitArrayJoinClause is called when exiting the arrayJoinClause production.
	ExitArrayJoinClause(c *ArrayJoinClauseContext)

	// ExitWindowClause is called when exiting the windowClause production.
	ExitWindowClause(c *WindowClauseContext)

	// ExitPrewhereClause is called when exiting the prewhereClause production.
	ExitPrewhereClause(c *PrewhereClauseContext)

	// ExitWhereClause is called when exiting the whereClause production.
	ExitWhereClause(c *WhereClauseContext)

	// ExitGroupByClause is called when exiting the groupByClause production.
	ExitGroupByClause(c *GroupByClauseContext)

	// ExitHavingClause is called when exiting the havingClause production.
	ExitHavingClause(c *HavingClauseContext)

	// ExitOrderByClause is called when exiting the orderByClause production.
	ExitOrderByClause(c *OrderByClauseContext)

	// ExitLimitByClause is called when exiting the limitByClause production.
	ExitLimitByClause(c *LimitByClauseContext)

	// ExitLimitClause is called when exiting the limitClause production.
	ExitLimitClause(c *LimitClauseContext)

	// ExitTealimitClause is called when exiting the tealimitClause production.
	ExitTealimitClause(c *TealimitClauseContext)

	// ExitSettingsClause is called when exiting the settingsClause production.
	ExitSettingsClause(c *SettingsClauseContext)

	// ExitJoinExprOp is called when exiting the JoinExprOp production.
	ExitJoinExprOp(c *JoinExprOpContext)

	// ExitJoinExprTable is called when exiting the JoinExprTable production.
	ExitJoinExprTable(c *JoinExprTableContext)

	// ExitJoinExprParens is called when exiting the JoinExprParens production.
	ExitJoinExprParens(c *JoinExprParensContext)

	// ExitJoinOpInner is called when exiting the JoinOpInner production.
	ExitJoinOpInner(c *JoinOpInnerContext)

	// ExitJoinOpLeftRight is called when exiting the JoinOpLeftRight production.
	ExitJoinOpLeftRight(c *JoinOpLeftRightContext)

	// ExitJoinOpFull is called when exiting the JoinOpFull production.
	ExitJoinOpFull(c *JoinOpFullContext)

	// ExitJoinOpCross is called when exiting the joinOpCross production.
	ExitJoinOpCross(c *JoinOpCrossContext)

	// ExitJoinConstraintClause is called when exiting the joinConstraintClause production.
	ExitJoinConstraintClause(c *JoinConstraintClauseContext)

	// ExitSampleClause is called when exiting the sampleClause production.
	ExitSampleClause(c *SampleClauseContext)

	// ExitLimitExpr is called when exiting the limitExpr production.
	ExitLimitExpr(c *LimitExprContext)

	// ExitTealimitExpr is called when exiting the tealimitExpr production.
	ExitTealimitExpr(c *TealimitExprContext)

	// ExitOrderExprList is called when exiting the orderExprList production.
	ExitOrderExprList(c *OrderExprListContext)

	// ExitOrderExpr is called when exiting the orderExpr production.
	ExitOrderExpr(c *OrderExprContext)

	// ExitRatioExpr is called when exiting the ratioExpr production.
	ExitRatioExpr(c *RatioExprContext)

	// ExitSettingExprList is called when exiting the settingExprList production.
	ExitSettingExprList(c *SettingExprListContext)

	// ExitSettingExpr is called when exiting the settingExpr production.
	ExitSettingExpr(c *SettingExprContext)

	// ExitWindowExpr is called when exiting the windowExpr production.
	ExitWindowExpr(c *WindowExprContext)

	// ExitWinPartitionByClause is called when exiting the winPartitionByClause production.
	ExitWinPartitionByClause(c *WinPartitionByClauseContext)

	// ExitWinOrderByClause is called when exiting the winOrderByClause production.
	ExitWinOrderByClause(c *WinOrderByClauseContext)

	// ExitWinFrameClause is called when exiting the winFrameClause production.
	ExitWinFrameClause(c *WinFrameClauseContext)

	// ExitFrameStart is called when exiting the frameStart production.
	ExitFrameStart(c *FrameStartContext)

	// ExitFrameBetween is called when exiting the frameBetween production.
	ExitFrameBetween(c *FrameBetweenContext)

	// ExitWinFrameBound is called when exiting the winFrameBound production.
	ExitWinFrameBound(c *WinFrameBoundContext)

	// ExitColumnTypeExprSimple is called when exiting the ColumnTypeExprSimple production.
	ExitColumnTypeExprSimple(c *ColumnTypeExprSimpleContext)

	// ExitColumnTypeExprNested is called when exiting the ColumnTypeExprNested production.
	ExitColumnTypeExprNested(c *ColumnTypeExprNestedContext)

	// ExitColumnTypeExprEnum is called when exiting the ColumnTypeExprEnum production.
	ExitColumnTypeExprEnum(c *ColumnTypeExprEnumContext)

	// ExitColumnTypeExprComplex is called when exiting the ColumnTypeExprComplex production.
	ExitColumnTypeExprComplex(c *ColumnTypeExprComplexContext)

	// ExitColumnTypeExprParam is called when exiting the ColumnTypeExprParam production.
	ExitColumnTypeExprParam(c *ColumnTypeExprParamContext)

	// ExitColumnExprList is called when exiting the columnExprList production.
	ExitColumnExprList(c *ColumnExprListContext)

	// ExitColumnsExprAsterisk is called when exiting the ColumnsExprAsterisk production.
	ExitColumnsExprAsterisk(c *ColumnsExprAsteriskContext)

	// ExitColumnsExprSubquery is called when exiting the ColumnsExprSubquery production.
	ExitColumnsExprSubquery(c *ColumnsExprSubqueryContext)

	// ExitColumnsExprColumn is called when exiting the ColumnsExprColumn production.
	ExitColumnsExprColumn(c *ColumnsExprColumnContext)

	// ExitColumnExprTernaryOp is called when exiting the ColumnExprTernaryOp production.
	ExitColumnExprTernaryOp(c *ColumnExprTernaryOpContext)

	// ExitColumnExprAlias is called when exiting the ColumnExprAlias production.
	ExitColumnExprAlias(c *ColumnExprAliasContext)

	// ExitColumnExprExtract is called when exiting the ColumnExprExtract production.
	ExitColumnExprExtract(c *ColumnExprExtractContext)

	// ExitColumnExprNegate is called when exiting the ColumnExprNegate production.
	ExitColumnExprNegate(c *ColumnExprNegateContext)

	// ExitColumnExprSubquery is called when exiting the ColumnExprSubquery production.
	ExitColumnExprSubquery(c *ColumnExprSubqueryContext)

	// ExitColumnExprLiteral is called when exiting the ColumnExprLiteral production.
	ExitColumnExprLiteral(c *ColumnExprLiteralContext)

	// ExitColumnExprBraceParam is called when exiting the ColumnExprBraceParam production.
	ExitColumnExprBraceParam(c *ColumnExprBraceParamContext)

	// ExitColumnExprArray is called when exiting the ColumnExprArray production.
	ExitColumnExprArray(c *ColumnExprArrayContext)

	// ExitColumnExprSubstring is called when exiting the ColumnExprSubstring production.
	ExitColumnExprSubstring(c *ColumnExprSubstringContext)

	// ExitColumnExprCast is called when exiting the ColumnExprCast production.
	ExitColumnExprCast(c *ColumnExprCastContext)

	// ExitColumnExprOr is called when exiting the ColumnExprOr production.
	ExitColumnExprOr(c *ColumnExprOrContext)

	// ExitColumnExprPrecedence1 is called when exiting the ColumnExprPrecedence1 production.
	ExitColumnExprPrecedence1(c *ColumnExprPrecedence1Context)

	// ExitColumnExprPrecedence2 is called when exiting the ColumnExprPrecedence2 production.
	ExitColumnExprPrecedence2(c *ColumnExprPrecedence2Context)

	// ExitColumnExprPrecedence3 is called when exiting the ColumnExprPrecedence3 production.
	ExitColumnExprPrecedence3(c *ColumnExprPrecedence3Context)

	// ExitColumnExprInterval is called when exiting the ColumnExprInterval production.
	ExitColumnExprInterval(c *ColumnExprIntervalContext)

	// ExitColumnExprIsNull is called when exiting the ColumnExprIsNull production.
	ExitColumnExprIsNull(c *ColumnExprIsNullContext)

	// ExitColumnExprWinFunctionTarget is called when exiting the ColumnExprWinFunctionTarget production.
	ExitColumnExprWinFunctionTarget(c *ColumnExprWinFunctionTargetContext)

	// ExitColumnExprTrim is called when exiting the ColumnExprTrim production.
	ExitColumnExprTrim(c *ColumnExprTrimContext)

	// ExitColumnExprTuple is called when exiting the ColumnExprTuple production.
	ExitColumnExprTuple(c *ColumnExprTupleContext)

	// ExitColumnExprArrayAccess is called when exiting the ColumnExprArrayAccess production.
	ExitColumnExprArrayAccess(c *ColumnExprArrayAccessContext)

	// ExitColumnExprBetween is called when exiting the ColumnExprBetween production.
	ExitColumnExprBetween(c *ColumnExprBetweenContext)

	// ExitColumnExprParens is called when exiting the ColumnExprParens production.
	ExitColumnExprParens(c *ColumnExprParensContext)

	// ExitColumnExprTimestamp is called when exiting the ColumnExprTimestamp production.
	ExitColumnExprTimestamp(c *ColumnExprTimestampContext)

	// ExitColumnExprAnd is called when exiting the ColumnExprAnd production.
	ExitColumnExprAnd(c *ColumnExprAndContext)

	// ExitColumnExprTupleAccess is called when exiting the ColumnExprTupleAccess production.
	ExitColumnExprTupleAccess(c *ColumnExprTupleAccessContext)

	// ExitColumnExprCase is called when exiting the ColumnExprCase production.
	ExitColumnExprCase(c *ColumnExprCaseContext)

	// ExitColumnExprDate is called when exiting the ColumnExprDate production.
	ExitColumnExprDate(c *ColumnExprDateContext)

	// ExitColumnExprNot is called when exiting the ColumnExprNot production.
	ExitColumnExprNot(c *ColumnExprNotContext)

	// ExitColumnExprWinFunction is called when exiting the ColumnExprWinFunction production.
	ExitColumnExprWinFunction(c *ColumnExprWinFunctionContext)

	// ExitColumnExprIdentifier is called when exiting the ColumnExprIdentifier production.
	ExitColumnExprIdentifier(c *ColumnExprIdentifierContext)

	// ExitColumnExprFunction is called when exiting the ColumnExprFunction production.
	ExitColumnExprFunction(c *ColumnExprFunctionContext)

	// ExitColumnExprAsterisk is called when exiting the ColumnExprAsterisk production.
	ExitColumnExprAsterisk(c *ColumnExprAsteriskContext)

	// ExitColumnArgList is called when exiting the columnArgList production.
	ExitColumnArgList(c *ColumnArgListContext)

	// ExitColumnArgExpr is called when exiting the columnArgExpr production.
	ExitColumnArgExpr(c *ColumnArgExprContext)

	// ExitColumnLambdaExpr is called when exiting the columnLambdaExpr production.
	ExitColumnLambdaExpr(c *ColumnLambdaExprContext)

	// ExitColumnIdentifier is called when exiting the columnIdentifier production.
	ExitColumnIdentifier(c *ColumnIdentifierContext)

	// ExitNestedIdentifier is called when exiting the nestedIdentifier production.
	ExitNestedIdentifier(c *NestedIdentifierContext)

	// ExitTableExprIdentifier is called when exiting the TableExprIdentifier production.
	ExitTableExprIdentifier(c *TableExprIdentifierContext)

	// ExitTableExprSubquery is called when exiting the TableExprSubquery production.
	ExitTableExprSubquery(c *TableExprSubqueryContext)

	// ExitTableExprAlias is called when exiting the TableExprAlias production.
	ExitTableExprAlias(c *TableExprAliasContext)

	// ExitTableExprFunction is called when exiting the TableExprFunction production.
	ExitTableExprFunction(c *TableExprFunctionContext)

	// ExitTableFunctionExpr is called when exiting the tableFunctionExpr production.
	ExitTableFunctionExpr(c *TableFunctionExprContext)

	// ExitTableIdentifier is called when exiting the tableIdentifier production.
	ExitTableIdentifier(c *TableIdentifierContext)

	// ExitTableArgList is called when exiting the tableArgList production.
	ExitTableArgList(c *TableArgListContext)

	// ExitTableArgExpr is called when exiting the tableArgExpr production.
	ExitTableArgExpr(c *TableArgExprContext)

	// ExitDatabaseIdentifier is called when exiting the databaseIdentifier production.
	ExitDatabaseIdentifier(c *DatabaseIdentifierContext)

	// ExitFloatingLiteral is called when exiting the floatingLiteral production.
	ExitFloatingLiteral(c *FloatingLiteralContext)

	// ExitNumberLiteral is called when exiting the numberLiteral production.
	ExitNumberLiteral(c *NumberLiteralContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitInterval is called when exiting the interval production.
	ExitInterval(c *IntervalContext)

	// ExitKeyword is called when exiting the keyword production.
	ExitKeyword(c *KeywordContext)

	// ExitKeywordForAlias is called when exiting the keywordForAlias production.
	ExitKeywordForAlias(c *KeywordForAliasContext)

	// ExitAlias is called when exiting the alias production.
	ExitAlias(c *AliasContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitIdentifierOrNull is called when exiting the identifierOrNull production.
	ExitIdentifierOrNull(c *IdentifierOrNullContext)

	// ExitEnumValue is called when exiting the enumValue production.
	ExitEnumValue(c *EnumValueContext)
}
