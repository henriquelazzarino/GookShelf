# GookShelf
## Description
GookShelf is a library application that uses the Gin framework for routing, Firebase for authentication, and JWT for access control. The application manages three main entities: Books, Users, and Publishers. There are three types of personas with different access levels: Admin, Librarian, and User.

## Entities
### Book
- name (string): Name of the book.
- description (string): Description of the book.
- author (string): Author of the book.
- publisher (Publisher): Publisher of the book.
- isFree (boolean): Indicates if the book is available for free.
- releaseDate (date): Release date of the book.
- minimumAge (int): Recommended minimum age for reading the book.
### User
- name (string): Name of the user.
- picture (file): User's picture.
- age (int): Age of the user.
- bookedBooks (list<Book>): List of books reserved by the user.
### Publisher
- name (string): Name of the publisher.
- books (list<Book>): List of books published by the publisher.

## Endpoints
### Books
- POST - /books
    - Description: Create a new book.
    - Accessible by: Admin, Librarian
- GET (All) - /books
    - Description: Retrieve the list of all books.
    - Accessible by: Admin, Librarian, User
- GET (One) - /books/{id}
    - Description: Retrieve details of a specific book.
    - Accessible by: Admin, Librarian, User
- PUT - /books/{id}
    - Description: Update details of a specific book.
    - Accessible by: Admin, Librarian
- DELETE - /books/{id}
    - Description: Delete a specific book.
    - Accessible by: Admin, Librarian
### Users
- POST - /user
    - Description: Create a new user.
    - Accessible by: Admin
- GET (All) - /user
    - Description: Retrieve the list of all users.
    - Accessible by: Admin, User
- GET (One) - /user/{id}
    - Description: Retrieve details of a specific user.
    - Accessible by: Admin, User
- PUT - /user/{id}
    - Description: Update details of a specific user.
    - Accessible by: Admin, User (only their own account)
- DELETE - /user/{id}
    - Description: Delete a specific user.
    - Accessible by: Admin, User (only their own account)
### Publishers
- POST - /publisher
    - Description: Create a new publisher.
    - Accessible by: Admin
- GET (All) - /publisher
    - Description: Retrieve the list of all publishers.
    - Accessible by: Admin
- GET (One) - /publisher/{id}
    - Description: Retrieve details of a specific publisher.
    - Accessible by: Admin
- PUT - /publisher/{id}
    - Description: Update details of a specific publisher.
    - Accessible by: Admin
- DELETE - /publisher/{id}
    - Description: Delete a specific publisher.
    - Accessible by: Admin

## Business Rules
### Admin:
- Has full permissions to create, read, update, and delete all resources: books, users, and publishers.
### Librarian:
- Has permission to create, read, update, and delete books.
- Cannot access or manipulate users and publishers.
### User:
- Can read the list of books and details of a specific book.
- Can read details of all users.
- Can update and delete only their own user.

## How to Run
1. Clone the repository.
2. Run `go mod tidy` to install the dependencies.
3. Create a `.env` file in the config directory with the following environment variables:
    - `PORT`: Port number for the server.
    - `JWT_SECRET`: Secret key for JWT.
4. Create a Firebase project and download the service account key.
    - 4.1. Go to the Firebase console and create a new project.
    - 4.2. Go to the project settings and Service Accounts.
    - 4.3. Download the service account key and save it as `serviceAccountKey.json` in the config directory.
    - 4.4. Enable Firebase Realtime Database and Firebase Storage.
5. Run `go run main.go` to start the server.
6. Access the API using the base URL `http://localhost:{PORT}`.

## Technologies
- Go
    - Gin
- Firebase
    - Realtime Database
    - Storage
- JWT
