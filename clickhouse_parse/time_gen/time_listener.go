// Code generated from /Users/bytedance/go/src/code.byted.org/lark-approval/ai_sdk/clickhouse_parse/Time.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Time

import "github.com/antlr4-go/antlr/v4"

// TimeListener is a complete listener for a parse tree produced by TimeParser.
type TimeListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterToDateTimeOrZero is called when entering the ToDateTimeOrZero production.
	EnterToDateTimeOrZero(c *ToDateTimeOrZeroContext)

	// EnterCompare is called when entering the Compare production.
	EnterCompare(c *CompareContext)

	// EnterClause is called when entering the Clause production.
	EnterClause(c *ClauseContext)

	// EnterToDateTimeOrZeroWithTimeZone is called when entering the ToDateTimeOrZeroWithTimeZone production.
	EnterToDateTimeOrZeroWithTimeZone(c *ToDateTimeOrZeroWithTimeZoneContext)

	// EnterDateAdd is called when entering the DateAdd production.
	EnterDateAdd(c *DateAddContext)

	// EnterDate is called when entering the Date production.
	EnterDate(c *DateContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitToDateTimeOrZero is called when exiting the ToDateTimeOrZero production.
	ExitToDateTimeOrZero(c *ToDateTimeOrZeroContext)

	// ExitCompare is called when exiting the Compare production.
	ExitCompare(c *CompareContext)

	// ExitClause is called when exiting the Clause production.
	ExitClause(c *ClauseContext)

	// ExitToDateTimeOrZeroWithTimeZone is called when exiting the ToDateTimeOrZeroWithTimeZone production.
	ExitToDateTimeOrZeroWithTimeZone(c *ToDateTimeOrZeroWithTimeZoneContext)

	// ExitDateAdd is called when exiting the DateAdd production.
	ExitDateAdd(c *DateAddContext)

	// ExitDate is called when exiting the Date production.
	ExitDate(c *DateContext)
}
