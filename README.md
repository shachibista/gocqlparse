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
| `DROP KEYSPACE`                 |                    | Easy
| `DROP TABLE`                    | :white_check_mark: |
| `DROP INDEX`                    |                    | Easy
| `ALTER TABLE`                   |                    | Hard
| `ALTER KEYSPACE`                |                    | Easy
| `GRANT PERMISSIONS`             |                    | Moderate
| `REVOKE PERMISSIONS`            |                    | Moderate
| `LIST PERMISSIONS`              |                    | Moderate
| `CREATE USER`                   |                    | Easy
| `ALTER USER`                    |                    | Easy
| `DROP USER`                     |                    | Easy
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
| `LIST ROLES`                    |                    | Easy
| `GRANT ROLE`                    |                    | Easy
| `REVOKE ROLE`                   |                    | Easy
| `CREATE MATERIALIZED VIEW`      |                    | Hard
| `DROP MATERIALIZED VIEW`        |                    | Easy
| `ALTER MATERIALIZED VIEW`       |                    | Easy
| `DESCRIBE`                      |                    | Moderate
| `ADD IDENTITY`                  |                    | Easy
| `DROP IDENTITY`                 |                    | Easy
| `LIST SUPER USERS`              |                    | Easy

