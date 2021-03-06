gophersql
=========

Database/SQL talk for Denver Gophers. Drivers and code examples are on the [package page](http://golang.org/pkg/database/sql/).

### Recommended Drivers

* [Postgres](https://github.com/lib/pq)
* [SQLite3](https://github.com/mattn/go-sqlite3)
* [MySQL](https://github.com/ziutek/mymysql)


Examples
--------

Import syntax, the empty `_` import

Building the connection string

Knowing the parameterization scheme

`Exec` versus `Query`

Scanning into pointers

Closing rows

The scanner interface

Cross-platform issues

Talk `struct` with the database!


Issues
------

### Zero-initialization

Default `INSERT` values and dealing with `NULL` values with pointers or Null types.


### Closing rows and duplicate connections with :memory:

The case of the disappearing table.


### Scanning into an interface

Strings are scanned as `[]byte`, resulting in unexpected `json` output.


### No interface for connections and transactions

Build one to test instances requiring a connection


Packages
--------

A list of popular projects

* [sqlx](https://github.com/jmoiron/sqlx) - provides a super-set of the `database/sql` library
* [Goose](https://bitbucket.org/liamstask/goose) - database migrations
* [Go Relational Persistence](https://github.com/coopernurse/gorp) - talk `struct` with the database
* [GORM](https://github.com/jinzhu/gorm) - talk `struct` with the database

And my own: [Aspect](https://github.com/aodin/aspect) - cross-dialect query construction, schema declarations, and talking a variety of types with the database (including structs and slices)
