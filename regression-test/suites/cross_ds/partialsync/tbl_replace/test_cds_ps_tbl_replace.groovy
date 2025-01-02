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
// under the License.

suite('test_cds_ps_tbl_replace') {
    def helper = new GroovyShell(new Binding(['suite': delegate]))
            .evaluate(new File("${context.config.suitePath}/../common", 'helper.groovy'))

    if (!helper.is_version_supported([30004, 20108, 20099])) {
        // at least doris 3.0.4, 2.1.8
        def version = helper.upstream_version()
        logger.info("skip this suite because version is not supported, upstream version ${version}")
        return
    }

    def tableNameA = 'tbl_a_' + helper.randomSuffix()
    def tableNameB = 'tbl_b_' + helper.randomSuffix()

    def exist = { res -> Boolean
        return res.size() != 0
    }
    def notExist = { res -> Boolean
        return res.size() == 0
    }

    sql """
        CREATE TABLE if NOT EXISTS ${tableNameA}
        (
            `test` INT,
            `id` INT
        )
        ENGINE=OLAP
        UNIQUE KEY(`test`, `id`)
        PARTITION BY RANGE(`id`)
        (
            PARTITION `p1` VALUES LESS THAN ("0"),
            PARTITION `p2` VALUES LESS THAN ("100"),
            PARTITION `p3` VALUES LESS THAN ("200"),
            PARTITION `p4` VALUES LESS THAN ("300"),
            PARTITION `p5` VALUES LESS THAN ("1000")
        )
        DISTRIBUTED BY HASH(id) BUCKETS AUTO
        PROPERTIES (
            "replication_allocation" = "tag.location.default: 1",
            "binlog.enable" = "true"
        )
    """
    sql """
        CREATE TABLE if NOT EXISTS ${tableNameB}
        (
            `test` INT,
            `id` INT
        )
        ENGINE=OLAP
        UNIQUE KEY(`test`, `id`)
        PARTITION BY RANGE(`id`)
        (
            PARTITION `p1` VALUES LESS THAN ("0"),
            PARTITION `p2` VALUES LESS THAN ("100"),
            PARTITION `p3` VALUES LESS THAN ("200"),
            PARTITION `p4` VALUES LESS THAN ("300"),
            PARTITION `p5` VALUES LESS THAN ("1000")
        )
        DISTRIBUTED BY HASH(id) BUCKETS AUTO
        PROPERTIES (
            "replication_allocation" = "tag.location.default: 1",
            "binlog.enable" = "true"
        )
    """

    helper.enableDbBinlog()
    helper.ccrJobDelete()
    helper.ccrJobCreate()

    assertTrue(helper.checkRestoreFinishTimesOf(tableNameA, 60))

    def first_job_progress = helper.get_job_progress()

    helper.ccrJobPause()

    logger.info(' === Add a key column to trigger a patrial sync === ')
    sql "ALTER TABLE ${tableNameA} ADD COLUMN `new_col` INT KEY DEFAULT \"0\""

    assertTrue(helper.checkShowTimesOf("""
                                SHOW ALTER TABLE COLUMN
                                FROM ${context.dbName}
                                WHERE TableName = "${tableNameA}" AND State = "FINISHED"
                                """,
                                exist, 30))

    sql "INSERT INTO ${tableNameA} VALUES (1, 100, 100), (100, 1, 1), (2, 200, 200), (200, 2, 2)"

    logger.info(' === Replace table A with table B, swap = false')
    sql """
        ALTER TABLE ${tableNameA} REPLACE WITH TABLE ${tableNameB} PROPERTIES ("swap"="false")
        """

    sql "INSERT INTO ${tableNameA} VALUES (5, 500)"

    assertTrue(helper.checkShowTimesOf("SHOW TABLES LIKE '${tableNameB}'", notExist, 60))

    helper.ccrJobResume()

    assertTrue(helper.checkSelectTimesOf("SELECT * FROM ${tableNameA}", 1, 60))

    // TODO: check tableNameB is not exists in the downstream
    // FIXME
    // [2024-12-24 06:22:50.124] DEBUG table commit seq map: map[17586:2252], table name mapping: map[17586:tbl_a_2064939174] job=test_cds_ps_tbl_replace line=ccr/job.go:524
    // [2024-12-24 06:22:50.124]  WARN partial sync table tbl_a_2064939174 id not match, force full sync. table id 17479, backup object id 17586 job=test_cds_ps_tbl_replace line=ccr/job.go:528
    // [2024-12-24 06:22:50.124]  INFO new snapshot, commitSeq: 2222 job=test_cds_ps_tbl_replace line=ccr/job.go:3116

    // // no fullsync are triggered
    // def last_job_progress = helper.get_job_progress()
    // assertTrue(last_job_progress.full_sync_start_at == first_job_progress.full_sync_start_at)
}



