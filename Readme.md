# Test task for Clario (Weather App 2)

## Task
Знайти 2-3 відкриті API, які надають інформацію про погоду.
Написати CLI додаток, який приймає країну та місто, та на виході віддає інформацію про погоду за поточний день.
Вивести результати тієї API, яка найшвидше віддала відповідь.

## Description
The project uses a simple layered architecture.
Basic ideas: separation of levels of logic, clear interfaces, dependencies directed only "upward".

The project does not use any dependencies, even the mocks for the tests are self-created because the project is simple and small.

Go version 1.23.2.

## Getting Started
1. Build the project:

    ```bash
    make build
    ```

2. Run the project:

    ```bash
    make run
    ```

3. Run tests on the project:
   
    ```bash
    make test
    ```