# psGoWebApp

Pluralsight course.

http://www.pluralsight.com/courses/creating-web-applications-go

Reminding myself how to use Git and Github at the same time.

## Known Issues during Module 8: Persisting Data

 1. All data is hard-coded.
 2. Only one category (Juices) has been implemented.
 3. Stand_locator doesn't implement the postcode finder. Just returns hardcoded data.
 4. Tests only added for a few types.
 
## Observations from this learning project

The name of the source file has no bearing on how the contents is referred to in other source files. Any public methods will be visible to a consumer once the package is imported.

Thanks to http://nathanleclaire.com/blog/2014/08/09/dont-get-bitten-by-pointer-vs-non-pointer-method-receivers-in-golang/, I understand the reason for **func (this \*struct) Fn()** + **func (this type) Fn()** and when to chose between them. Hoorah!

MVC has proven hard to follow in this project. Naming conventions (matching  concepts have the same name in model/package/controller/converter folders) may be a large part of the issue.

### SQL Lite

SQLite go library chosen: https://github.com/mattn/go-sqlite3

64-bit MINGW build that actually worked with CGO was here: http://sourceforge.net/projects/mingwbuilds/?source=typ_redirect

SQLite manager recommendation: http://stackoverflow.com/questions/835069/which-sqlite-administration-console-do-you-recommend 

Database structure:

    CREATE TABLE "member" ("id" INTEGER PRIMARY KEY  AUTOINCREMENT  NOT NULL  UNIQUE , "email" VARCHAR, "password" VARCHAR, "first_name" VARCHAR)

    CREATE TABLE "session" ("id" INTEGER PRIMARY KEY  AUTOINCREMENT  NOT NULL  UNIQUE , "session_id" VARCHAR NOT NULL , "member_id" INTEGER NOT NULL )

    CREATE TABLE sqlite_sequence(name,seq)
