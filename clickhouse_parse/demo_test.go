package clickhouse_parse

import (
	parser "code.byted.org/lark-approval/ai_sdk/clickhouse_parse/gen"
	"context"
	"fmt"
	"strings"
	"testing"
)

type sqlListener struct {
	*parser.BaseDPListener

	stack []int
}

func TestParse(t *testing.T) {
	//expre := "SELECT  *,instance_id,count(distinct task_id) FROM  lark_approval_dashboard.form_analysis_demo WHERE process_name='aaa' AND task_assignee_id = '8f1d7d5dcd6898269a0f7a0f3b12ae7af8a447c7'  AND (toInt64(form_map ['采购数量']) > 5 OR toInt64(form_map ['采购数量'])<10) AND toDateTimeOrZero(instance_end_time) < '2023-12-31'"
	expre := "select instance_id,task_id from lark_nl2sql_task where (start_user = '57f83a891da4d8451f9555b3874d0d68c4ffe37b'  OR assagine = '57f83a891da4d8451f9555b3874d0d68c4ffe37b') AND toDateTimeOrZero(instance_start_time,'America/New_York') >= toDateTimeOrZero(date_add(DAY, -7, toDateTimeOrZero('2024-04-12'))) Order by task_end_time desc, instance_start_time "
	v := NewVisitor()

	var LarkNl2sqlTaskCol = []string{
		"process_name", "instance_start_time", "instance_end_time",
		"task_start_time", "task_end_time", "instance_id", "task_id",
		"instance_status", "start_user", "task_approval_all_dept_ids_now", "assignee",
	}
	for _, s := range LarkNl2sqlTaskCol {
		v.TableCol[s] = false
	}

	v.VisitSQL(context.Background(), expre)
	fmt.Println(v.OrderByClause)

	if v.TableCol["process_name"] {
		for _, clause := range v.WhereClause {
			if strings.Contains(clause, "process_name") {
				processQuery := strings.Split(clause, "=")[1]

				processQuery = strings.ReplaceAll(processQuery, `'`, "")
				fmt.Println(processQuery)
			}
		}
	}
	fmt.Println(v.WhereClause)
}

func GenerateSQLFromVisit(ctx context.Context, tenantID, userID int64, visit *Visitor) string {
	sql := ""
	sql += "SELECT "
	for idx, s := range visit.SelectClause {
		sql += s
		if idx < len(visit.SelectClause)-1 {
			sql += ", "
		}
	}
	if len(visit.TableName) == 0 {
		return ""
	}
	return sql
}
