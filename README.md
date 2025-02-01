Parser for Cassandra CQL3 query language.

# Installation

```
go get github.com/shachibista/gocqlparse
```

# Implemented Statements

| Statement                       | Status             | Difficulty
|---------------------------------|--------------------|-----------
| `SELECT`                        | :large_orange_diamond: | Hard
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
| `ALTER KEYSPACE`                | :white_check_mark: |
| `GRANT PERMISSIONS`             |                    | Moderate
| `REVOKE PERMISSIONS`            |                    | Moderate
| `LIST PERMISSIONS`              |                    | Moderate
| `CREATE USER`                   | :white_check_mark: |
| `ALTER USER`                    | :white_check_mark: |
| `DROP USER`                     | :white_check_mark: |
| `LIST USERS`                    | :white_check_mark: |
| `CREATE TRIGGER`                | :white_check_mark: |
| `DROP TRIGGER`                  | :white_check_mark: |
| `CREATE TYPE`                   | :white_check_mark: |
| `ALTER TYPE`                    |                    | Moderate
| `DROP TYPE`                     | :white_check_mark: |
| `CREATE FUNCTION`               |                    | Moderate
| `DROP FUNCTION`                 |                    | Moderate
| `CREATE AGGREGATE`              |                    | Moderate
| `DROP AGGREGATE`                |                    | Moderate
| `CREATE ROLE`                   |                    | Moderate
| `ALTER ROLE`                    |                    | Moderate
| `DROP ROLE`                     | :white_check_mark: |
| `LIST ROLES`                    | :white_check_mark: |
| `GRANT ROLE`                    | :white_check_mark: |
| `REVOKE ROLE`                   | :white_check_mark: |
| `CREATE MATERIALIZED VIEW`      |                    | Hard
| `DROP MATERIALIZED VIEW`        | :white_check_mark: |
| `ALTER MATERIALIZED VIEW`       | :white_check_mark: |
| `DESCRIBE`                      |                    | Moderate
| `ADD IDENTITY`                  | :white_check_mark: |
| `DROP IDENTITY`                 | :white_check_mark: |
| `LIST SUPER USERS`              | :white_check_mark: |
