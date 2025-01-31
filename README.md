Parser for Cassandra CQL3 query language.

# Installation

```
go get github.com/shachibista/gocqlparse
```

# Implemented Statements

| Statement                       | Status             |
|---------------------------------|--------------------|
| `SELECT`                        |                    |
| `INSERT`                        |                    |
| `UPDATE`                        |                    |
| `BATCH`                         |                    |
| `DELETE`                        |                    |
| `USE`                           | :white_check_mark: |
| `TRUNCATE`                      |                    |
| `CREATE KEYSPACE`               |                    |
| `CREATE TABLE`                  | :white_check_mark: |
| `CREATE INDEX`                  |                    |
| `DROP KEYSPACE`                 |                    |
| `DROP TABLE`                    |                    |
| `DROP INDEX`                    |                    |
| `ALTER TABLE`                   |                    |
| `ALTER KEYSPACE`                |                    |
| `GRANT PERMISSIONS`             |                    |
| `REVOKE PERMISSIONS`            |                    |
| `LIST PERMISSIONS`              |                    |
| `CREATE USER`                   |                    |
| `ALTER USER`                    |                    |
| `DROP USER`                     |                    |
| `LIST USERS`                    |                    |
| `CREATE TRIGGER`                |                    |
| `DROP TRIGGER`                  |                    |
| `CREATE TYPE`                   | :white_check_mark: |
| `ALTER TYPE`                    |                    |
| `DROP TYPE`                     |                    |
| `CREATE FUNCTION`               |                    |
| `DROP FUNCTION`                 |                    |
| `CREATE AGGREGATE`              |                    |
| `DROP AGGREGATE`                |                    |
| `CREATE ROLE`                   |                    |
| `ALTER ROLE`                    |                    |
| `DROP ROLE`                     |                    |
| `LIST ROLES`                    |                    |
| `GRANT ROLE`                    |                    |
| `REVOKE ROLE`                   |                    |
| `CREATE MATERIALIZED VIEW`      |                    |
| `DROP MATERIALIZED VIEW`        |                    |
| `ALTER MATERIALIZED VIEW`       |                    |
| `DESCRIBE`                      |                    |
| `ADD IDENTITY`                  |                    |
| `DROP IDENTITY`                 |                    |
| `LIST SUPER USERS`              |                    |

