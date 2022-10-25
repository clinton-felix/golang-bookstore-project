# golang-bookstore-project

## This is a bookstore management system built with golang, connected to a local mysql database

### FEATURES:

> Get List of all books in the Database
> Create a Book
> Get a Book by ID
> Update a Book
> Delete a Book from the Database

### How To Use

1. Ensure that you have mysql workspace running in your background.

2. run "git pull https://github.com/clinton-felix/golang-bookstore-project"

3. Update the func Connect in the app.go file, by inserting your MYSQL root password

4. Run "go run main.go" in your terminal. This will start a server at Localhost port: 9010

5. open postman or Thunder Client in vscode and test the API route calls to the Database

## Available Route calls:

> Get Books: /book/ :(Method: GET)
> Create a Book: /book/ :(Method: POST)
> Get a Book by ID: /book/{id}/ :(Method: GET)
> Update a Book: /book/{id}/ :(Method: PUT)
> Delete a Book: /book/{id}/ :(Method: DELETE)
