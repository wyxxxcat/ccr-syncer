// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License
package ccr_test

import (
	"testing"

	"github.com/selectdb/ccr_syncer/pkg/ccr"
)

func TestIsSessionVariableRequired(t *testing.T) {
	tests := []string{
		"If you want to specify column names, please `set enable_nereids_planner=true`",
		"set enable_variant_access_in_original_planner = true in session variable",
		"Please enable the session variable 'enable_projection' through `set enable_projection = true",
		"agg state not enable, need set enable_agg_state=true",
		"which is greater than 38 is disabled by default. set enable_decimal256 = true to enable it",
		"if we have a column with decimalv3 type and set enable_decimal_conversion = false",
	}
	for i, test := range tests {
		if !ccr.IsSessionVariableRequired(test) {
			t.Errorf("test %d failed", i)
		}
	}
}
