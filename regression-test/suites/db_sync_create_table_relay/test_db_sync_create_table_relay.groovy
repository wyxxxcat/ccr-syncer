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

suite("test_db_sync_create_table_relay") {
    def helper = new GroovyShell(new Binding(['suite': delegate]))
            .evaluate(new File("${context.config.suitePath}/../common", "helper.groovy"))

    def tableName = "test_db_sync_create_table_relay_table_1"

    def dbNameOrigin = context.dbName
    def dbNameRelay = "TEST_" + context.dbName

    def test_num = 0
    def insert_num = 10

    def exist = { res -> Boolean
        return res.size() != 0
    }

    helper.enableDbBinlog()

    sql "DROP TABLE IF EXISTS ${dbNameOrigin}.${tableName}"
    sql "DROP TABLE IF EXISTS ${dbNameRelay}.${tableName}"

    sql """
        CREATE TABLE if NOT EXISTS ${tableName}
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

    helper.ccrJobDelete(tableName, dbNameOrigin)
    helper.ccrJobCreate(tableName, dbNameOrigin)

    assertTrue(helper.checkRestoreFinishTimesOf("${tableName}", 60))

    for (int index = 0; index < insert_num; index++) {
        sql """
            INSERT INTO ${tableName} VALUES (${test_num}, ${index})
            """
    }

    sql "sync"

    logger.info("=== Test 1: Resume and check table ===")

    helper.ccrJobResume(tableName, dbNameOrigin)

    assertTrue(helper.checkShowTimesOf(""" SHOW TABLES LIKE "${tableName}" """, exist, 60, "target"))

    helper.ccrJobPause()

    logger.info("=== Test 2: Delete old job and create new job ===")

    helper.ccrJobDelete(tableName, dbNameOrigin)

    helper.ccrJobCreate(tableName, dbNameRelay)

    assertTrue(helper.checkRestoreFinishTimesOf("${tableName}", 60))

    sql "sync"

    helper.ccrJobPause()
    
    logger.info("=== Test 3: Resume and Check new table ===")

    helper.ccrJobResume(tableName, dbNameRelay)

    assertTrue(helper.checkShowTimesOf(""" SHOW TABLES LIKE "${tableName}" """, exist, 60, "target"))

    result = sql "select * from ${dbNameRelay}.${tableName}"

    assertEquals(insert_num, result.size())
}