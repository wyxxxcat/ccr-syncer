package record

import (
	"encoding/json"
	"fmt"
	"regexp"

	"github.com/selectdb/ccr_syncer/pkg/xerror"
)

type CreateMTMV struct {
	DbId    int64  `json:"dbId"`
	TableId int64  `json:"tableId"`
	Sql     string `json:"sql"`

	// Below fields was added in doris 2.0.3: https://github.com/apache/doris/pull/26901
	DbName   string `json:"dbName"`
	MtmvName string `json:"tableName"`
}

func NewCreateMTMVFromJson(data string) (*CreateMTMV, error) {
	var createMTMV CreateMTMV
	err := json.Unmarshal([]byte(data), &createMTMV)
	if err != nil {
		return nil, xerror.Wrap(err, xerror.Normal, "unmarshal create mtmv error")
	}

	if createMTMV.Sql == "" {
		return nil, xerror.Errorf(xerror.Normal, "create mtmv sql is empty")
	}

	if createMTMV.TableId == 0 {
		return nil, xerror.Errorf(xerror.Normal, "mtmv id not found")
	}

	return &createMTMV, nil
}

func (c *CreateMTMV) IsCreateMv() bool {
	viewRegex := regexp.MustCompile(`(?i)^CREATE(\s+)MATERIALIZED VIEW`)
	return viewRegex.MatchString(c.Sql)
}

// String
func (c *CreateMTMV) String() string {
	return fmt.Sprintf("CreateMTMV: DbId: %d, DbName: %s, TableId: %d, MtmvName: %s, Sql: %s",
		c.DbId, c.DbName, c.TableId, c.MtmvName, c.Sql)
}
