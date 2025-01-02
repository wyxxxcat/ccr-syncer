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
package record

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"

	"github.com/selectdb/ccr_syncer/pkg/xerror"
)

type CreateTable struct {
	DbId    int64  `json:"dbId"`
	TableId int64  `json:"tableId"`
	Sql     string `json:"sql"`

	// Below fields was added in doris 2.0.3: https://github.com/apache/doris/pull/26901
	DbName    string `json:"dbName"`
	TableName string `json:"tableName"`

	// Below fields was added in doris 2.1.8/3.0.4: https://github.com/apache/doris/pull/44735
	TableType string `json:"tableType"`
}

func NewCreateTableFromJson(data string) (*CreateTable, error) {
	var createTable CreateTable
	err := json.Unmarshal([]byte(data), &createTable)
	if err != nil {
		return nil, xerror.Wrap(err, xerror.Normal, "unmarshal create table error")
	}

	if createTable.Sql == "" {
		// TODO: fallback to create sql from other fields
		return nil, xerror.Errorf(xerror.Normal, "create table sql is empty")
	}

	if createTable.TableId == 0 {
		return nil, xerror.Errorf(xerror.Normal, "table id not found")
	}

	return &createTable, nil
}

func (c *CreateTable) IsCreateView() bool {
	viewRegex := regexp.MustCompile(`(?i)^CREATE(\s+)VIEW`)
	return viewRegex.MatchString(c.Sql)
}

// String
func (c *CreateTable) String() string {
	return fmt.Sprintf("CreateTable: DbId: %d, DbName: %s, TableId: %d, TableName: %s, Sql: %s",
		c.DbId, c.DbName, c.TableId, c.TableName, c.Sql)
}

func (c *CreateTable) IsCreateTableWithInvertedIndex() bool {
	indexRegex := regexp.MustCompile(`INDEX (.*?) USING INVERTED`)
	return indexRegex.MatchString(c.Sql)
}

// Is asynchronous materialized view?
func (c *CreateTable) IsCreateMaterializedView() bool {
	if c.TableType == TableTypeMaterializedView {
		return true
	}

	return strings.Contains(c.Sql, "ENGINE=MATERIALIZED_VIEW")
}
