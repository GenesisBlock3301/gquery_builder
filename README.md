
| Concept       | Example                         | How we handle it in builder                                                                               |
| ------------- | ------------------------------- | --------------------------------------------------------------------------------------------------------- |
| **QueryType** | SELECT, INSERT, UPDATE, DELETE  | Enum (`QueryType`) because it determines the **type of query**                                            |
| **Clause**    | FROM, WHERE, GROUP BY, ORDER BY | Usually **hardcoded in Build()** or handled by methods, because they are **parts of a query**, not a type |


Done:
- Select
- InsertInto
- Update
- Delete
- From
- Where
- DropTable
- DropDatabase

# Go SQL Builder

A simple **SQL query builder in Go** inspired by Python’s `pypika` library.  
This project demonstrates **builder pattern**, **interfaces**, and **polymorphism** in Go for constructing SQL queries dynamically.

---

## Features

- **Fluent Builder Pattern**
    - Chain methods to build queries iteratively.
    - Example: `Query{}.From("customers").Select("*").Where("id = 1")`.

- **Supports Multiple SQL Types**
    - `SELECT`
    - `INSERT`
    - `UPDATE`
    - `DELETE`
    - `CREATE TABLE`
    - `DROP TABLE` / `DROP DATABASE`

- **Interfaces for Polymorphism**
    - All query builders implement a `Builder` interface with `Build()` method.
    - Enables functions to accept **any builder** type (`SELECT`, `CREATE`, `DROP`) seamlessly.

- **Extensible**
    - Easy to add new builders like `CreateIndexBuilder`, `DropUserBuilder`.
    - Any new builder only needs to implement `Build()` to integrate with existing functions.

- **Chainable Query Building**
    - Methods return the builder itself (`*QueryBuilder`) to allow fluent chaining.
    - Example:
      ```go
      q := Query{}
      sql := q.From("customers").Select("*").Where("id = 1").Build()
      ```

- **Flexible Print / Execution**
    - `PrintSQL(b Builder)` can print **any query** implementing the `Builder` interface.
    - Example:
      ```go
      var b Builder = q.CreateTable("users")
      PrintSQL(b)
      ```

---

## Example Usage

```go
func main() {
    q := Query{}

    // SELECT
    var b Builder = q.From("customers").Select("*").Where("id = 1")
    PrintSQL(b) // Output: SELECT * FROM customers WHERE id = 1

    // CREATE TABLE
    b = q.CreateTable("users")
    PrintSQL(b) // Output: CREATE TABLE users (...)

    // DROP DATABASE
    b = q.DropDatabase("testdb")
    PrintSQL(b) // Output: DROP DATABASE testdb
}
