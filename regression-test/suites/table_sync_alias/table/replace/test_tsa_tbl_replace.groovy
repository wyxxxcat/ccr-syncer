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

suite("test_tsa_tbl_replace") {
    def helper = new GroovyShell(new Binding(['suite': delegate]))
            .evaluate(new File("${context.config.suitePath}/../common", "helper.groovy"))

    if (!helper.is_version_supported([30003, 20108, 20016])) {
        // at least doris 3.0.3, 2.1.8 and doris 2.0.16
        def version = helper.upstream_version()
        logger.info("skip this suite because version is not supported, upstream version ${version}")
        return
    }

    def oldTableName = "tbl_old_" + helper.randomSuffix()
    def newTableName = "tbl_new_" + helper.randomSuffix()
    def aliasTableName = "tbl_alias_" + helper.randomSuffix()

    helper.set_alias(aliasTableName)

    def exist = { res -> Boolean
        return res.size() != 0
    }
    def notExist = { res -> Boolean
        return res.size() == 0
    }

    logger.info("=== Create both table ===")
    sql """
        CREATE TABLE if NOT EXISTS ${oldTableName}
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
        CREATE TABLE if NOT EXISTS ${newTableName}
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

    helper.ccrJobDelete(oldTableName)
    helper.ccrJobCreate(oldTableName)

    assertTrue(helper.checkRestoreFinishTimesOf("${oldTableName}", 60))

    sql "INSERT INTO ${oldTableName} VALUES (1, 100), (100, 1), (2, 200), (200, 2)"
    assertTrue(helper.checkSelectTimesOf("SELECT * FROM ${aliasTableName}", 4, 60))

    logger.info(" ==== replace with swap ==== ")

    sql "INSERT INTO ${newTableName} VALUES (3, 300), (300, 3)"  // o:n, 4:2
    sql "INSERT INTO ${oldTableName} VALUES (3, 300), (300, 3)"  // o:n, 6:2
    sql "ALTER TABLE ${oldTableName} REPLACE WITH TABLE ${newTableName} PROPERTIES (\"swap\"=\"true\")"  // o:n, 2:6
    sql "INSERT INTO ${oldTableName} VALUES (4, 400)"            // o:n, 3:6
    sql "INSERT INTO ${newTableName} VALUES (4, 400)"            // o:n, 3:7

    assertTrue(helper.checkSelectTimesOf("SELECT * FROM ${aliasTableName}", 3, 60))

    logger.info(" ==== replace without swap ==== ")

    sql "INSERT INTO ${newTableName} VALUES (5, 500), (500, 5)"  // o:n, 3:9
    sql "INSERT INTO ${oldTableName} VALUES (5, 500), (500, 5)"  // o:n, 5:9
    sql "ALTER TABLE ${oldTableName} REPLACE WITH TABLE ${newTableName} PROPERTIES (\"swap\"=\"false\")"  // o:n, 9:0
    sql "INSERT INTO ${oldTableName} VALUES (6, 600)"            // o:n, 10:0

    assertTrue(helper.checkSelectTimesOf("SELECT * FROM ${aliasTableName}", 10, 60))
}

