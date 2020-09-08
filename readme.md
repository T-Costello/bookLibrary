This is a simple go example that represents a Book Library API.  
---
Gin is an extension to go's built in http libraries.  
Gorm is an object relational mapper to an in memory database.

###### Getting Started:
1.  Depending on your environment dependencies may download automatically.  If not, enter the following:

    - go get github.com/gin-gonic/gin   
    - go get github.com/jinzhu/gorm
    - go get github.com/mattn/go-sqlite3

2.  Build the project 

3.  Run main.go

4.  Endpoints
    - Get all books ```GET localhost:8080/books```
    - Get book ```GET localhost:8080/books/2```
    - Create book ```POST localhost:8080/books```
    ```
    {
      "title": "All or Nothing",
      "author": "Gambler Swanson"
    }
    ```
    - Update book ```PATCH localhost:8080/books```
    ```
    {
      "title": "The Infinite Game"
    }
    ```
    - Delete book ```DELETE localhost:8080/books/2```


