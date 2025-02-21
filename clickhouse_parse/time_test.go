package clickhouse_parse

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"
)

func TestTimeParse(t *testing.T) {
	//expre := " toDateTimeOrZero(instance_start_time,'America/Chicago') >= toDateTimeOrZero(date_add(DAY, -7, toDateTimeOrZero('2024-12-12')))"
	expre := " toDateTimeOrZero(instance_start_time) >= toDateTimeOrZero('2024-12-01')"
	v := NewTimeVisitor()
	startT := time.Now()
	v.VisitTime(context.Background(), expre)
	costT := time.Since(startT)
	fmt.Println(costT)
	fmt.Println(v.Clause)
	s := expre
	condition := make([]string, 0)
	if strings.Contains(s, "toDateTimeOrZero") {
		timeVisit := NewTimeVisitor()
		timeVisit.VisitTime(context.Background(), s)
		clause := ""
		if len(timeVisit.Clause) > 0 {
			clause = timeMap[timeVisit.Clause[0]]
		}
		if clause == "" {

		}
		if len(timeVisit.DateAdd.DateType) > 0 {
			condition = append(condition, clause+"在最近"+timeVisit.Number+timeTypeMap[timeVisit.DateType])
		} else {
			condition = append(condition, clause+"在"+timeVisit.Date+"之后")
		}
	}
	fmt.Println(condition)
}

var timeMap = map[string]string{
	"task_end_time":       "任务完成时间",
	"task_start_time":     "任务开始时间",
	"instance_start_time": "单据创建时间",
	"instance_end_time":   "单据结束时间",
}

var timeTypeMap = map[string]string{
	"YEAR": "年",
	"DAY":  "天",
	"WEEK": "周",
}
