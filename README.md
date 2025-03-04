# Amartha-Loan-Service
This is a loan engine service project where user can request for funding loan. Employee can approve a loan request by providing employee id and picture of proof. Approved loan request can be invested by investors. Once the investments have reached the treshold amount requested on the loan request, it will send link to agreement letter to all investors. 

This project implements PostgreSQL database and Swagger for API documentation. This project is broken down into 4 layers:

1. Router: handles HTTP requests and routing to the appropriate APIs.
2. Controller: handles interaction between router and service layer. Handles request and validates inputs before calling the service layer.
3. Service: handles business logic where the data is processed before interacting with repository layer (if needed).
4. Repository: handles database interactions.

This file structure is implemented to ensure clean code practice, maintainability, and scalability by separating the components. 

Several adjustments were made for this assignment purposes:

1. Borrower ID, employee/field officer ID, and investor ID are inputted manually instead of parsing form JWT token after authentication.
2. Field validator picture proof and agreement letter are represented by a mock file link instead of using actual file download link form a Cloud Object Storage (COS) where the files should be stored.
3. Email distribution after an investment is fully invested is replaced with a function call to simulate the email sending process.
4. .env file is kept instead of ignored/deleted for ease of development and testing process.

## Getting Started (How to Run the Program)

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

* Clone the repository and navigate to the project directory.
* Install all its dependencies by running the command below.

    ```bash
    go mod tidy
    ```
* Run command below to generate the Swagger documentation
    ```bash
    swag init
    ```
* Run the app.
    ```bash
    go run main.go
    ```
* Input the link below into your browser

    ```bash
    localhost:3000/swagger/index.html#/
    ```
* The Swagger UI should appear like this
  ![Swagger](https://github.com/user-attachments/assets/e6a9913d-cbcf-446f-9310-9292fa95f8c7)

## PostgreSQL Database

* This project implements PostgreSQL to store its data. Several tables are designed to accomodate the project specifications. Below is the ERD of the project:

![ERD](https://github.com/user-attachments/assets/d6f3e6d0-023d-4659-ae98-37c48db82d99)

* We also used DBeaver as the database tool. To install PostgreSQL, setting up the database, and install DBeaver please follow this link

https://medium.com/@zum.hatice/how-to-create-a-postgresql-db-and-connect-in-windows-b26eaa48c7fb

* The database connection takes parameter from the .env file. We can adjust the parameters so it matched with the credentials we inputted during the creation of the database.

![env](https://github.com/user-attachments/assets/a488c82e-1bf4-401f-982c-9f633262a096)

