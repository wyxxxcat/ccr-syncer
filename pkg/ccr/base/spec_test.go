// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License
package base_test

import (
	"reflect"
	"testing"

	"github.com/selectdb/ccr_syncer/pkg/ccr/base"
	"github.com/selectdb/ccr_syncer/pkg/ccr/record"
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

func TestReplaceAndEscapeComment(t *testing.T) {
	type TestCase struct {
		origin, expect string
	}

	testCases := []TestCase{
		{"CREATE TABLE `t` (\n  `test` int NOT NULL COMMENT '[\"0000-01-01\", \"9999-12-31\"]'", "CREATE TABLE `t` (\n  `test` int NOT NULL COMMENT \"[\\\"0000-01-01\\\", \\\"9999-12-31\\\"]\""},
		{"CREATE TABLE `t` (\n  `test` int NOT NULL COMMENT 'xxx\"test\"'", "CREATE TABLE `t` (\n  `test` int NOT NULL COMMENT \"xxx\\\"test\\\"\""},
		{"CREATE TABLE `t` (\n  `test1` int NOT NULL COMMENT 'xxx\"test1\"', `test2` int NOT NULL COMMENT 'xxx\"test2\"'", "CREATE TABLE `t` (\n  `test1` int NOT NULL COMMENT \"xxx\\\"test1\\\"\", `test2` int NOT NULL COMMENT \"xxx\\\"test2\\\"\""},
		{"CREATE TABLE `t` (\n  `test1` int NOT NULL COMMENT '涓\ue161浆杩愯緭鍗曚俊鎭'", "CREATE TABLE `t` (\n  `test1` int NOT NULL COMMENT \"涓浆杩愯緭鍗曚俊鎭\""},
	}

	for i, c := range testCases {
		if actual := base.ReplaceAndEscapeComment(c.origin); actual != c.expect {
			t.Errorf("case %d failed, expect %s, but got %s", i, c.expect, actual)
		}
	}
}

func TestCheckModifyTablePropertySql(t *testing.T) {

	type TestCase struct {
		origin record.ModifyTableProperty
		expect map[string]string
	}
	testCases := []TestCase{
		{
			origin: record.ModifyTableProperty{
				DbId:      0,
				TableId:   0,
				TableName: "test_table_0",
				Properties: map[string]string{
					"replication_num":          "3",
					"storage_policy":           "policy_0",
					"dynamic_partition.enable": "true",
					"compaction_policy":        "time_series",
				},
				Sql: "SET (\"replication_num\"=\"3\", \"storage_policy\"=\"policy_0\", \"dynamic_partition.enable\"=\"true\", \"compaction_policy\"=\"time_series\")",
			},
			expect: map[string]string{
				"compaction_policy": "time_series",
			},
		},
		{
			origin: record.ModifyTableProperty{
				DbId:      1,
				TableId:   1,
				TableName: "test_table_1",
				Properties: map[string]string{
					"colocate_with": "group1",
					"bucket_num":    "10",
					"binlog.enable": "true",
				},
				Sql: "SET (\"colocate_with\"=\"group1\", \"bucket_num\"=\"10\", \"binlog.enable\"=\"true\")",
			},
			expect: map[string]string{
				"bucket_num": "10",
			},
		},
		{
			origin: record.ModifyTableProperty{
				DbId:      2,
				TableId:   2,
				TableName: "test_table_2",
				Properties: map[string]string{
					"colocate_with": "group2",
					"binlog.enable": "true",
				},
				Sql: "SET (\"colocate_with\"=\"group1\", \"binlog.enable\"=\"true\")",
			},
			expect: map[string]string{},
		},
	}
	for i, c := range testCases {
		actual := base.FilterUnsupportedProperties(&c.origin)

		if !reflect.DeepEqual(actual, c.expect) {
			t.Errorf("case %d failed, expect %v, but got %v", i, c.expect, actual)
		}
	}
}

func TestHandleDefaultValue(t *testing.T) {
	type TestCase struct {
		origin record.ModifyTableAddOrDropColumns
		expect string
	}

	testCases := []TestCase{
		{
			origin: record.ModifyTableAddOrDropColumns{
				DbId:    0,
				TableId: 0,
				RawSql:  "ALTER TABLE `t` ADD COLUMN `test` DATETIME NOT NULL DEFAULT \"CURRENT_TIMESTAMP\" COMMENT 'test'",
				IndexSchemaMap: map[int64][]record.ColumnSchema{
					0: {
						{Name: "test", Type: record.ColumnType{Type: "DATETIME"}, DefaultValue: "CURRENT_TIMESTAMP"},
					},
				},
			},
			expect: "ALTER TABLE `t` ADD COLUMN `test` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'test'",
		},
		{
			origin: record.ModifyTableAddOrDropColumns{
				DbId:    0,
				TableId: 0,
				RawSql:  "ALTER TABLE `t` ADD COLUMN `test` BITMAP NOT NULL DEFAULT \"BITMAP_EMPTY_DEFAULT_VALUE\" COMMENT 'test'",
				IndexSchemaMap: map[int64][]record.ColumnSchema{
					0: {
						{Name: "test", Type: record.ColumnType{Type: "BITMAP"}, DefaultValue: "BITMAP_EMPTY_DEFAULT_VALUE"},
					},
				},
			},
			expect: "ALTER TABLE `t` ADD COLUMN `test` BITMAP NOT NULL DEFAULT BITMAP_EMPTY_DEFAULT_VALUE COMMENT 'test'",
		},
		{
			origin: record.ModifyTableAddOrDropColumns{
				DbId:    0,
				TableId: 0,
				RawSql:  "ALTER TABLE `t` ADD COLUMN `test` VARCHAR(10) NOT NULL DEFAULT \"BITMAP_EMPTY_DEFAULT_VALUE\" COMMENT 'test'",
				IndexSchemaMap: map[int64][]record.ColumnSchema{
					0: {
						{Name: "test", Type: record.ColumnType{Type: "VARCHAR", Len: 10}, DefaultValue: "BITMAP_EMPTY_DEFAULT_VALUE"},
					},
				},
			},
			expect: "ALTER TABLE `t` ADD COLUMN `test` VARCHAR(10) NOT NULL DEFAULT \"BITMAP_EMPTY_DEFAULT_VALUE\" COMMENT 'test'",
		},
		{
			origin: record.ModifyTableAddOrDropColumns{
				DbId:    0,
				TableId: 0,
				RawSql:  "ALTER TABLE `t` ADD COLUMN `test` VARCHAR(10) NOT NULL DEFAULT \"CURRENT_TIMESTAMP\" COMMENT 'test'",
				IndexSchemaMap: map[int64][]record.ColumnSchema{
					0: {
						{Name: "test", Type: record.ColumnType{Type: "VARCHAR", Len: 10}, DefaultValue: "CURRENT_TIMESTAMP"},
					},
				},
			},
			expect: "ALTER TABLE `t` ADD COLUMN `test` VARCHAR(10) NOT NULL DEFAULT \"CURRENT_TIMESTAMP\" COMMENT 'test'",
		},
	}

	for i, c := range testCases {
		if actual := base.HandleSchemaChangeDefaultValue(c.origin.RawSql, &c.origin); actual != c.expect {
			t.Errorf("case %d failed, expect %s, but got %s", i, c.expect, actual)
		}
	}
}
