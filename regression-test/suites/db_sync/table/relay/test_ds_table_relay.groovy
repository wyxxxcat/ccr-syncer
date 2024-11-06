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

suite("test_ds_table_relay") {
    def helper = new GroovyShell(new Binding(['suite': delegate]))
            .evaluate(new File("${context.config.suitePath}/../common", "helper.groovy"))

    def tableName = "test_db_sync_create_table_relay_table_1"
    def tableTmpName = "test_db_sync_table_tmp"

    def dbNameOrigin = context.dbName
    def dbNameRelay = "TEST_" + context.dbName
    def dbNameNew = "TEST_" + "TEST_" + context.dbName


    def test_num = 0
    def insert_num = 10

    def exist = { res -> Boolean
        return res.size() != 0
    }

    def notExist = { res -> Boolean
        return res.size() == 0
    }

    helper.enableDbBinlog()

    sql "CREATE DATABASE IF NOT EXISTS ${dbNameRelay}"
    sql "CREATE DATABASE IF NOT EXISTS ${dbNameNew}"

    sql "DROP TABLE IF EXISTS ${dbNameOrigin}.${tableName}"
    sql "DROP TABLE IF EXISTS ${dbNameRelay}.${tableName}"
    sql "DROP TABLE IF EXISTS ${dbNameNew}.${tableName}"
    sql "DROP TABLE IF EXISTS ${dbNameNew}.${tableTmpName}"


    sql """
        CREATE TABLE if NOT EXISTS ${dbNameOrigin}.${tableName}
        (
            `test` INT,
            `id` INT
        )
        UNIQUE KEY(`test`, `id`)
        DISTRIBUTED BY HASH(id) BUCKETS 1
        PROPERTIES (
            "replication_allocation" = "tag.location.default: 1",
            "binlog.enable" = "true"
        )
    """
    sql """
        CREATE TABLE if NOT EXISTS ${dbNameRelay}.${tableTmpName}
        (
            `test` INT,
            `id` INT
        )
        UNIQUE KEY(`test`, `id`)
        DISTRIBUTED BY HASH(id) BUCKETS 1
        PROPERTIES (
            "replication_allocation" = "tag.location.default: 1",
            "binlog.enable" = "true"
        )
    """

    helper.ccrJobDelete("", dbNameOrigin)
    helper.ccrJobDelete("", dbNameRelay)
    helper.ccrJobCreate("", dbNameRelay)
    helper.ccrJobCreate("", dbNameOrigin)

    assertTrue(helper.checkRestoreFinishTimesOf("${tableName}", 60))

    for (int index = 0; index < insert_num; index++) {
        sql """
            INSERT INTO ${dbNameOrigin}.${tableName} VALUES (${test_num}, ${index})
            """
    }
    for (int index = 0; index < insert_num; index++) {
        sql """
            INSERT INTO ${dbNameRelay}.${tableTmpName} VALUES (${test_num}, ${index})
            """
    }

    resultOrigin = sql "select * from ${dbNameOrigin}.${tableName}"

    assertEquals(resultOrigin.size(), insert_num)

    resultOrigin = sql "select * from ${dbNameRelay}.${tableTmpName}"

    assertEquals(resultOrigin.size(), insert_num)

    logger.info("=== Test 1: Check table exist ===")

    assertTrue(helper.checkShowTimesOf(""" select * from ${dbNameRelay}.${tableTmpName} """, exist, 60, "target"))

    assertTrue(helper.checkShowTimesOf(""" select * from ${dbNameNew}.${tableTmpName} """, exist, 60, "target"))

    assertTrue(helper.checkShowTimesOf(""" select * from ${dbNameOrigin}.${tableName} """, exist, 60, "target"))

    assertTrue(helper.checkShowTimesOf(""" select * from ${dbNameRelay}.${tableName} """, exist, 60, "target"))

    logger.info("=== Test 3: Check table not exist ===")

    sql "USE ${dbNameNew}"

    assertTrue(helper.checkShowTimesOf(""" SHOW TABLES LIKE "${tableName}" """, notExist, 60, "sql"))
}