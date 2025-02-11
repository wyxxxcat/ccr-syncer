from itertools import count
import random
import string
from time import sleep
import mysql.connector
from mysql.connector import Error
import logging

class DDLGenerator:
    def __init__(self, executor):
        self.executor = executor
        self.table_columns = {}  
        self._init_table_columns() 
        self.partitioned_tables = {}  # {table_name: partition_config}
        self.col_counter = 0
        self.table_partition_type = {}  # {table_name: partition_type}

    def _init_table_columns(self):
        for table in self.executor.get_table_names():
            self.table_columns[table] = self.executor.get_table_columns(table)
            self.table_partition_type = self.executor.get_partition_type(table)

    def generate_add_partition(self):
        if not self.partitioned_tables:
            return None

        table_name = random.choice(list(self.partitioned_tables.keys()))
        table_partition_type = self.table_partition_type.get(table_name)
        if table_partition_type == "NONE":
            return None
        config = self.partitioned_tables[table_name]

        partition_name = f"p{len(self.partitioned_tables[table_name].get('partitions', [])) + 1}"
        if config["data_type"] == "INT":
            new_value_1 = config["first_value"] + random.randint(0, 1000)
            new_value_2 = config["last_value"] + random.randint(1001, 2000)
            if table_partition_type == "RANGE":
                partition_def = f"VALUES [({new_value_1}), ({new_value_2}))"
            else:
                partition_def = f"VALUES (({new_value_1}, {new_value_2}))"
            self.partitioned_tables[table_name]["last_value"] = new_value_2
        else:  
            new_value_1 = f"'{2023 + random.randint(1,5)}-01-01'"
            new_value_2 = f"'{2023 + random.randint(6,10)}-01-01'"
            if table_partition_type == "RANGE":
                 partition_def = f"VALUES [({new_value_1}), (new_value_2))"
            else:
                new_value = f"'{2023 + random.randint(1,5)}-01-01'"
                partition_def = f"VALUES (({new_value}))"
            new_value = f"'{2023 + random.randint(1,5)}-01-01'"
            partition_def = f"VALUES LESS THAN ({new_value})"

        return f"ALTER TABLE {table_name} ADD PARTITION (PARTITION {partition_name} {partition_def});"

    def generate_add_column(self):
        if not self.executor.get_table_names():
            return None
        table_name = random.choice(list(self.executor.get_table_names()))
        new_col = f"new_col_{''.join(random.choices(string.ascii_lowercase, k=3))}"
        data_type = random.choice(["INT", "VARCHAR(100)", "DATE"])
        if table_name not in self.table_columns:
            self.table_columns[table_name] = []
        self.table_columns[table_name].append(new_col)
        return f"ALTER TABLE {table_name} ADD COLUMN {new_col} {data_type};"

    def generate_truncate_table(self):
        if not self.executor.get_table_names():
            return None
        table_name = random.choice(list(self.executor.get_table_names()))
        return f"TRUNCATE TABLE {table_name};"

    def generate_alter_table_rollup(self):
        if not self.executor.get_table_names():
            return None

        valid_tables = [t for t in self.executor.get_table_names() if len(self.table_columns.get(t, [])) >= 2]
        if not valid_tables:
            return None

        table_name = random.choice(valid_tables)
        available_columns = self.table_columns[table_name]

        rollup_columns = random.sample(available_columns, k=random.randint(2, min(4, len(available_columns))))

        rollup_name = f"r_{''.join(random.choices(string.digits, k=3))}"

        return f"ALTER TABLE {table_name} ADD ROLLUP {rollup_name} ({', '.join(rollup_columns)});"

    def generate_alter_table_rename(self):
        if not self.executor.get_table_names():
            return None
        old_table_name = random.choice(list(self.executor.get_table_names()))
        new_table_name = random.choice(list(self.executor.get_table_names()))
        random_prefix = ''.join(random.choices(string.ascii_lowercase, k=3))
        return f"ALTER TABLE {old_table_name} RENAME {random_prefix}_{new_table_name};"

    def generate_alter_table_replace(self):
        if not self.executor.get_table_names():
            return None
        table_name = random.choice(list(self.executor.get_table_names()))
        replaced_table = random.choice(list(self.executor.get_table_names()))
        return f"ALTER TABLE {table_name} REPLACE WITH TABLE {replaced_table};"

    def generate_alter_table_order_by(self):
        if not self.executor.get_table_names():
            return None


        table_name = random.choice(list(self.executor.get_table_names()))
        if table_name not in self.table_columns:
            return None
        all_columns = self.table_columns[table_name]

        first_column = all_columns[0]
        remaining_columns = all_columns[1:]
        random.shuffle(remaining_columns)
        ordered_columns = [first_column] + remaining_columns

        from_clause = ""
        if random.random() < 0.3 and self.partitioned_tables.get(table_name):
            from_clause = f" FROM {random.choice(list(self.partitioned_tables[table_name]))}"

        return f"ALTER TABLE {table_name} ORDER BY ({', '.join(ordered_columns)}){from_clause};"
    
    def generate_create_index(self):
        index_name = ''.join(random.choices(string.ascii_lowercase, k=3)) + "_index"
        table_name = random.choice(list(self.executor.get_table_names()))
        if table_name not in self.table_columns:
            return None
        available_columns = self.table_columns[table_name]

        index_column  = random.sample(available_columns, k=random.randint(1, min(2, len(available_columns))))
        bitmap = random.choice(["USING BITMAP", ""])
        return f"CREATE INDEX {index_name} ON {table_name} ({', '.join(index_column)}) {bitmap};"

    def generate_create_view(self):
        view_name = ''.join(random.choices(string.ascii_lowercase, k=3)) + "_view"
        table_name = random.choice(list(self.executor.get_table_names()))
        if table_name not in self.table_columns:
            return None
        available_columns = self.table_columns[table_name]
        view_column = random.sample(available_columns, k=random.randint(1, min(2, len(available_columns))))
        return f"CREATE VIEW {view_name} AS SELECT {', '.join(view_column)} FROM {table_name};"

    def generate_random_ddl(self):
        operation = random.choices(
            # order by error too easily
            ["ADD_COLUMN", "ADD_PARTITION", "ADD_ROLLUP", "RENAME", "REPLACE", "CREATE INDEX", "CREATE VIEW"],
            weights=[0.3, 0.2, 0.1, 0.1, 0.1, 0.2, 0.2],
            k=1
        )[0]

        if operation == "ADD_COLUMN":
            return self.generate_add_column()
        elif operation == "ADD_PARTITION":
            return self.generate_add_partition()
        elif operation == "ADD_ROLLUP":
            return self.generate_alter_table_rollup()
        elif operation == "ORDER_BY":
            return self.generate_alter_table_order_by()
        elif operation == "RENAME":
            return self.generate_alter_table_rename()
        elif operation == "REPLACE":
            return self.generate_alter_table_replace()
        elif operation == "CREATE INDEX":
            return self.generate_create_index()
        elif operation == "CREATE VIEW":
            return self.generate_create_view()
        return None

class DDLExecutor:
    def __init__(self, host, port, user, password, database=None):
        self.host = host
        self.user = user
        self.password = password
        self.database = database
        self.connection = None
        self.port = port

    def connect(self):
        try:
            self.connection = mysql.connector.connect(
                host=self.host,
                port=self.port,
                user=self.user,
                password=self.password,
                database=self.database
            )
            print("Successfully connected to database!")
        except Error as e:
            print(f"Error connecting to MySQL: {e}")

    def get_table_names(self):
        if not self.connection:
            print("Not connected to database!")
            return set()

        cursor = self.connection.cursor()
        try:
            cursor.execute("SHOW TABLES")
            tables = cursor.fetchall()
            if not tables:
                print("No tables found in the database!")
                return set()
            
            table_names = set()
            for table in tables:
                table_name = table[0]
                cursor.execute(f"SHOW CREATE TABLE {table_name}")
                result = cursor.fetchone()
                if result is not None:
                    create_table_sql = result[1]
                    if "CREATE TABLE" in create_table_sql:
                        table_names.add(table_name)
            
            return table_names
        except Error as e:
            print(f"Error fetching tables: {e}")
            return set()
        finally:
            cursor.close()

    def execute_ddl(self, ddl):
        if not self.connection:
            print("Not connected to database!")
            return False
        cursor = self.connection.cursor()
        try:
            cursor.execute(ddl)
            self.connection.commit()
            return "Success"
        except Exception as e:
            print(f"Execute Err: {str(e)}")
            return f"Failed: {str(e)}"
        finally:
            cursor.close()
            if self.connection:
                self.connection.rollback()

    def close(self):
        if self.connection:
            self.connection.close()
            print("Connection closed.")

    def get_table_columns(self, table_name):
        if not self.connection:
            return []

        cursor = self.connection.cursor()
        try:
            cursor.execute(f"DESCRIBE {table_name}")
            result = cursor.fetchall()
            return [row[0] for row in result] if result else []
        except Error as e:
            print(f"Get Table Column Err: {e}")
            return []
        finally:
            cursor.close()

    def get_partition_type(self, table_name):
        if not self.connection:
            return []

        cursor = self.connection.cursor()
        try:
            cursor.execute(f"SHOW CREATE TABLE {table_name}")
            result = cursor.fetchall()
            if str(result).find("PARTITION BY RANGE") != -1:
                return "RANGE"
            elif str(result).find("PARTITION BY LIST") != -1:
                return "LIST"
            else:
                return "NONE"
        except Error as e:
            print(f"Get Partition Type Err: {e}")
            return []
        finally:
            cursor.close()

if __name__ == "__main__":
    logging.basicConfig(
        filename='rand_ddl.log',
        level=logging.INFO,
        format='%(message)s',
        filemode='w'
    )

    DB_CONFIG = {
        "host": "127.0.0.1",
        "port": "9330",
        "user": "root",
        "password": "",
        "database": "db"
    }

    executor = DDLExecutor(**DB_CONFIG)
    executor.connect()

    table_names = executor.get_table_names()
    if not table_names:
        print("No tables found in the database!")
        exit()

    generator = DDLGenerator(executor)
    count = 50

    for i in range(count):
        ddl = generator.generate_random_ddl()
        if ddl:
            print("\n" + "="*50 + "\n")
            print("Generated DDL:")
            logging.info("\n" + "="*50 + "\n")
            logging.info("Generated DDL:\n" + ddl + "\n")
            print(ddl)

            result = executor.execute_ddl(ddl)
            logging.info("Execution Result: " + str(result) + "\n")
            print("Execution Result:", result)

            sleep(1)

    executor.close()