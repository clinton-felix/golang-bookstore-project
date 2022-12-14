# golang-bookstore-project

## This is a bookstore management system built with golang, connected to a local mysql database

### FEATURES:

1. Get List of all books in the Database
2. Create a Book
3. Get a Book by ID
4. Update a Book
5. Delete a Book from the Database

### How To Use

1. Ensure that you have mysql workspace running in your background.

2. run "git pull https://github.com/clinton-felix/golang-bookstore-project"

3. Update the func Connect in the app.go file, by inserting your MYSQL root password

4. Run "go run main.go" in your terminal. This will start a server at Localhost port: 9010

5. open postman or Thunder Client in vscode and test the API route calls to the Database

## Available Route calls:

1. Get Books: /book/ :(Method: GET)

2. Create a Book: /book/ :(Method: POST)

3. Get a Book by ID: /book/{id}/ :(Method: GET)

4. Update a Book: /book/{id}/ :(Method: PUT)

5. Delete a Book: /book/{id}/ :(Method: DELETE)
