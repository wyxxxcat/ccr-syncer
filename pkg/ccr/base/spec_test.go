package base_test

import (
	"testing"

	"github.com/selectdb/ccr_syncer/pkg/ccr/base"
)

func TestAddDBPrefixToCreateTableOrViewSql(t *testing.T) {
	type TestCase struct {
		origin, expect string
	}

	testCases := []TestCase{
		{"CREATE VIEW `v` AS SELECT * FROM t", "CREATE VIEW `target_db`.`v` AS SELECT * FROM t"},
		{"CREATE VIEW `target_db`.`v` AS SELECT * FROM t", "CREATE VIEW `target_db`.`v` AS SELECT * FROM t"},
		{" CREATE VIEW `v` AS SELECT * FROM t", "CREATE VIEW `target_db`.`v` AS SELECT * FROM t"},
		{"CREATE TABLE `v` (...", "CREATE TABLE `target_db`.`v` (..."},
		{"CREATE VIEW `view_test_746794472` AS SELECT `internal`.`TEST_regression_test_db_sync_mv_basic$.`tbl_duplicate_0_746794472`.`user_id` AS `k1`, `internal`.`TEST_regression_test_db_sync_mv_basic`.`tbl_duplicate_0_746794472`.`name` AS `name`, SUM(`internal`.`TEST_regression_te$t_db_sync_mv_basic`.`tbl_duplicate_0_746794472`.`age`) AS `v1` FROM `internal`.`TEST_regression_test_db_sync_mv_basic`.`tbl_duplicate_0_746794472` GROUP BY k1,`internal`.`TEST_regression_test_db_sync_mv_basic`.`tbl_duplicate_0_746794472`.`name`",
			"CREATE VIEW `target_db`.`view_test_746794472` AS SELECT `internal`.`TEST_regression_test_db_sync_mv_basic$.`tbl_duplicate_0_746794472`.`user_id` AS `k1`, `internal`.`TEST_regression_test_db_sync_mv_basic`.`tbl_duplicate_0_746794472`.`name` AS `name`, SUM(`internal`.`TEST_regression_te$t_db_sync_mv_basic`.`tbl_duplicate_0_746794472`.`age`) AS `v1` FROM `internal`.`TEST_regression_test_db_sync_mv_basic`.`tbl_duplicate_0_746794472` GROUP BY k1,`internal`.`TEST_regression_test_db_sync_mv_basic`.`tbl_duplicate_0_746794472`.`name`"},
	}

	for i, c := range testCases {
		if actual := base.AddDBPrefixToCreateTableOrViewSql("target_db", c.origin); actual != c.expect {
			t.Errorf("case %d failed, expect %s, but got %s", i, c.expect, actual)
		}
	}
}

func TestReplaceDBNameForCreateSql(t *testing.T) {
	type TestCase struct {
		origin, expect string
	}

	testCases := []TestCase{
		{"CREATE VIEW `target_db`.`v` AS (select `origin_db`.`func`(`internal`.`target_db`.`create_view_table1`.`id`) as `c1`,abs(`internal`.`target_db`.`create_view_table1`.`id`) from `internal`.`target_db`.`create_view_table1`)", "CREATE VIEW `target_db`.`v` AS (select `target_db`.`func`(`internal`.`target_db`.`create_view_table1`.`id`) as `c1`,abs(`internal`.`target_db`.`create_view_table1`.`id`) from `internal`.`target_db`.`create_view_table1`)"},
		{"CREATE VIEW `target_db`.`v` AS (select `origin_db_not_replace`.`func`(`internal`.`target_db`.`create_view_table1`.`id`) as `c1`,abs(`internal`.`target_db`.`create_view_table1`.`id`) from `internal`.`target_db`.`create_view_table1`)", "CREATE VIEW `target_db`.`v` AS (select `origin_db_not_replace`.`func`(`internal`.`target_db`.`create_view_table1`.`id`) as `c1`,abs(`internal`.`target_db`.`create_view_table1`.`id`) from `internal`.`target_db`.`create_view_table1`)"},
		{"CREATE VIEW `target_db`.`v` AS (select `origin_db`.`func`(`internal`.`target_db`.`origin_db`.`id`) as `c1`,abs(`internal`.`target_db`.`origin_db`.`id`) from `internal`.`target_db`.`origin_db`)", "CREATE VIEW `target_db`.`v` AS (select `target_db`.`func`(`internal`.`target_db`.`origin_db`.`id`) as `c1`,abs(`internal`.`target_db`.`origin_db`.`id`) from `internal`.`target_db`.`origin_db`)"},
	}

	for i, c := range testCases {
		if actual := base.ReplaceDBNameForCreateSql(c.origin, "origin_db", "target_db"); actual != c.expect {
			t.Errorf("case %d failed, expect %s, but got %s", i, c.expect, actual)
		}
	}
}
