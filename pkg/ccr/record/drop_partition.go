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

type DropPartition struct {
	TableId int64  `json:"tableId"`
	Sql     string `json:"sql"`
	IsTemp  bool   `json:"isTempPartition"`
}

func NewDropPartitionFromJson(data string) (*DropPartition, error) {
	var dropPartition DropPartition
	err := json.Unmarshal([]byte(data), &dropPartition)
	if err != nil {
		return nil, xerror.Wrap(err, xerror.Normal, "unmarshal drop partition error")
	}

	if dropPartition.Sql == "" {
		// TODO: fallback to create sql from other fields
		return nil, xerror.Errorf(xerror.Normal, "drop partition sql is empty")
	}

	if dropPartition.TableId == 0 {
		return nil, xerror.Errorf(xerror.Normal, "table id not found")
	}

	return &dropPartition, nil
}
