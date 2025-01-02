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

	"github.com/selectdb/ccr_syncer/pkg/xerror"
)

type ReplacePartitionRecord struct {
	DbId           int64    `json:"dbId"`
	DbName         string   `json:"dbName"`
	TableId        int64    `json:"tblId"`
	TableName      string   `json:"tblName"`
	Partitions     []string `json:"partitions"`
	TempPartitions []string `json:"tempPartitions"`
	StrictRange    bool     `json:"strictRange"`
	UseTempName    bool     `json:"useTempPartitionName"`
}

func NewReplacePartitionFromJson(data string) (*ReplacePartitionRecord, error) {
	var replacePartition ReplacePartitionRecord
	err := json.Unmarshal([]byte(data), &replacePartition)
	if err != nil {
		return nil, xerror.Wrap(err, xerror.Normal, "unmarshal replace partition error")
	}

	if len(replacePartition.TempPartitions) == 0 {
		return nil, xerror.Errorf(xerror.Normal, "the temp partitions of the replace partition record is empty")
	}

	if replacePartition.TableId == 0 {
		return nil, xerror.Errorf(xerror.Normal, "table id not found")
	}

	if replacePartition.TableName == "" {
		return nil, xerror.Errorf(xerror.Normal, "table name is empty")
	}

	return &replacePartition, nil
}
