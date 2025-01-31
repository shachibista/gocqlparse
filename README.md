Parser for Cassandra CQL3 query language.

# Installation

```
go get github.com/shachibista/gocqlparse
```

# Implemented Statements

| Statement                       | Status             | Difficulty
|---------------------------------|--------------------|-----------
| `SELECT`                        |                    | Hard
| `INSERT`                        |                    | Moderate
| `UPDATE`                        |                    | Moderate
| `BATCH`                         |                    | Hard
| `DELETE`                        |                    | Moderate
| `USE`                           | :white_check_mark: |
| `TRUNCATE`                      | :white_check_mark: |
| `CREATE KEYSPACE`               | :white_check_mark: |
| `CREATE TABLE`                  | :white_check_mark: |
| `CREATE INDEX`                  |                    | Moderate
| `DROP KEYSPACE`                 | :white_check_mark: |
| `DROP TABLE`                    | :white_check_mark: |
| `DROP INDEX`                    | :white_check_mark: |
| `ALTER TABLE`                   |                    | Hard
| `ALTER KEYSPACE`                |                    | Easy
| `GRANT PERMISSIONS`             |                    | Moderate
| `REVOKE PERMISSIONS`            |                    | Moderate
| `LIST PERMISSIONS`              |                    | Moderate
| `CREATE USER`                   |                    | Easy
| `ALTER USER`                    |                    | Easy
| `DROP USER`                     | :white_check_mark: |
| `LIST USERS`                    | :white_check_mark: |
| `CREATE TRIGGER`                |                    | Easy
| `DROP TRIGGER`                  |                    | Easy
| `CREATE TYPE`                   | :white_check_mark: |
| `ALTER TYPE`                    |                    | Moderate
| `DROP TYPE`                     | :white_check_mark: |
| `CREATE FUNCTION`               |                    | Moderate
| `DROP FUNCTION`                 |                    | Moderate
| `CREATE AGGREGATE`              |                    | Moderate
| `DROP AGGREGATE`                |                    | Moderate
| `CREATE ROLE`                   |                    | Moderate
| `ALTER ROLE`                    |                    | Moderate
| `DROP ROLE`                     |                    | Easy
| `LIST ROLES`                    | :white_check_mark: |
| `GRANT ROLE`                    | :white_check_mark: |
| `REVOKE ROLE`                   | :white_check_mark: |
| `CREATE MATERIALIZED VIEW`      |                    | Hard
| `DROP MATERIALIZED VIEW`        | :white_check_mark: |
| `ALTER MATERIALIZED VIEW`       |                    | Easy
| `DESCRIBE`                      |                    | Moderate
| `ADD IDENTITY`                  | :white_check_mark: |
| `DROP IDENTITY`                 | :white_check_mark: |
| `LIST SUPER USERS`              | :white_check_mark: |
