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

suite("test_ts_tbl_rename") {
    def helper = new GroovyShell(new Binding(['suite': delegate]))
            .evaluate(new File("${context.config.suitePath}/../common", "helper.groovy"))

    def tableName = "tbl_" + helper.randomSuffix()
    def newTableName = "NEW_${tableName}"
    def dbName = context.dbName
    def test_num = 0
    def insert_num = 5

    def exist = { res -> Boolean
        return res.size() != 0
    }
    def notExist = { res -> Boolean
        return res.size() == 0
    }

    sql "DROP TABLE IF EXISTS ${dbName}.${tableName}"
    sql "DROP TABLE IF EXISTS TEST_${dbName}.${tableName}"

    helper.enableDbBinlog()
    sql """
        CREATE TABLE if NOT EXISTS ${tableName}
        (
            `test` INT,
            `id` INT
        )
        ENGINE=OLAP
        UNIQUE KEY(`test`, `id`)
        PARTITION BY RANGE(`test`)
        (
            PARTITION `less100` VALUES LESS THAN ("100")
        )
        DISTRIBUTED BY HASH(id) BUCKETS 1
        PROPERTIES (
            "replication_allocation" = "tag.location.default: 1",
            "binlog.enable" = "true"
        )
    """

    helper.ccrJobDelete(tableName)
    helper.ccrJobCreate(tableName)

    assertTrue(helper.checkRestoreFinishTimesOf("${tableName}", 30))

    logger.info("=== Test 0: Common insert case ===")
    for (int index = 0; index < insert_num; index++) {
        sql """
            INSERT INTO ${tableName} VALUES (${test_num}, ${index})
            """
    }
    sql "sync"
    assertTrue(helper.checkSelectTimesOf("SELECT * FROM ${tableName} WHERE test=${test_num}",
                                  insert_num, 30))

    logger.info("=== Test 1: Check old table exist and new table not exist ===")

    assertTrue(helper.checkShowTimesOf("SHOW TABLES LIKE '${tableName}'", exist, 30, "sql"))

    assertTrue(helper.checkShowTimesOf("SHOW TABLES LIKE '${newTableName}'", notExist, 30, "sql"))

    assertTrue(helper.checkShowTimesOf("SHOW TABLES LIKE '${tableName}'", exist, 30, "target_sql"))

    assertTrue(helper.checkShowTimesOf("SHOW TABLES LIKE '${newTableName}'", notExist, 30, "target_sql"))

    logger.info("=== Test 2: Rename table case and insert data ===")
    test_num = 1
    sql "ALTER TABLE ${tableName} RENAME ${newTableName}"

    for (int index = 0; index < insert_num; index++) {
        sql """
            INSERT INTO ${newTableName} VALUES (${test_num}, ${index})
            """
    }
    sql "sync"

    logger.info("=== Test 3: Check new table exist and old table not exist ===")

    assertTrue(helper.checkShowTimesOf("SHOW TABLES LIKE '${tableName}'", notExist, 30, "sql"))

    assertTrue(helper.checkShowTimesOf("SHOW TABLES LIKE '${newTableName}'", exist, 30, "sql"))

    assertTrue(helper.checkShowTimesOf("SHOW TABLES LIKE '${tableName}'", notExist, 30, "target_sql"))

    assertTrue(helper.checkShowTimesOf("SHOW TABLES LIKE '${newTableName}'", exist, 30, "target_sql"))

    assertTrue(helper.checkSelectTimesOf("SELECT * FROM ${newTableName} WHERE test=${test_num}",
                                  insert_num, 30))
}
