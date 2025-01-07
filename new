# Knex Migration with Docker

This project demonstrates the use of Knex.js for database migrations and seeding with Docker. It includes a simple Node.js application for handling migrations, seeds, and connecting to a database. The project is fully containerized using Docker and Docker Compose.

## Project Structure

The project repository includes the following files and directories:

- **docker-compose.yml**: Defines the services required for the setup, including the database and the Node.js application that handles migrations and seeds.
- **Dockerfile**: The Dockerfile for the main application, used to build and run the Node.js application inside a container.
- **knex/**: Contains the Node.js application and all related files.
  - **db/**: Directory that holds the database-related files.
  - **Dockerfile**: A Dockerfile used specifically for setting up the Node.js application inside a container.
  - **migrations/**: Directory where database migration files are stored.
  - **node_modules/**: Directory where the project's npm dependencies are installed.
  - **package.json**: Contains project dependencies and scripts for migrations, seeds, and other configurations.
  - **package-lock.json**: Automatically generated file that locks down the versions of dependencies.
  - **seeds/**: Directory where database seed files are stored.
  - **server.js**: The main application file for running Knex.js migrations and seed operations.

## Services in `docker-compose.yml`

The `docker-compose.yml` file defines the following services:

1. **Database**: The database service where Knex.js will connect and execute migrations/seeds.
2. **Node.js App**: The service that runs the Node.js application using Docker.

### Example `docker-compose.yml`:

```yaml
version: '3'

services:
  db:
    image: postgres:13
    environment:
      POSTGRES_USER: example
      POSTGRES_PASSWORD: example
      POSTGRES_DB: exampledb
    volumes:
      - ./knex/db:/var/lib/postgresql/data
    ports:
      - "5432:5432"

  knex-app:
    build: ./knex
    environment:
      DB_HOST: db
      DB_USER: example
      DB_PASSWORD: example
      DB_NAME: exampledb
    depends_on:
      - db
    command: ["npm", "run", "migrate"]

