#/repo

Contains all the code that interact to database.

Using CQRS (Command and Query Responsibility Segregation) pattern at application level pattern,
so we distinguish the code responsibility to 2 main roles which is:
1. Command, to manipulate data in database (the Create, Update, and Delete part of CRUD)
2. Query, to retrieve data from database (the Retrieve part from CRUD)

Command, using **entity** model (/entity) to manipulate the data at database.
Query, using **domain** model (/apps/domain) to retrieve data from database.