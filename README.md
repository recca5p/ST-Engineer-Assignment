# Full Stack Project with Next.js Frontend and Go Backend

This repository contains a full stack application with a Next.js frontend and a Go backend.

## Prerequisites

- Node.js 18 or later
- Go (latest version recommended)
- Docker
- PostgreSQL
- [golang-migrate](https://github.com/golang-migrate/migrate)
- [sqlc](https://sqlc.dev/)
- [mockgen](https://github.com/golang/mock)

## Frontend Setup

1. Clone the repository:
   ```
   git clone https://github.com/recca5p/ST-Engineer-Assignment.git
   cd ST-Engineer-Assignment/client
   ```

2. Install dependencies:
   ```
   yarn install
   ```

3. Run the development server:
   ```
   yarn dev
   ```

4. Open your browser and navigate to `http://localhost:3000` to view the application.

## Backend Setup

1. Working path
   ```
   cd ST-Engineer-Assignment/server
   ```
3. Set up the database:
   ```
   make infras
   ```
   This command uses Docker to set up a PostgreSQL instance.

4. Run database migrations:
   ```
   make migrateup
   ```

5. Seed the database with initial data (if applicable).

6. Start the server:
   ```
   make server
   ```

## Configuration

You can configure the application using the `app.env` file.

## Available Make Commands

- `make infras`: Start the Docker containers for the database
- `make migrateup`: Run all up migrations
- `make migrateup1`: Run one up migration
- `make migratedown`: Run all down migrations
- `make migratedown1`: Run one down migration
- `make new_migration name=migration_name`: Create a new migration file
- `make sqlc`: Generate Go code from SQL
- `make test`: Run tests
- `make server`: Start the Go server
- `make mock`: Generate mock for database interface

## Database Connection

The database connection string is set in the Makefile:

```
DB_URL=postgresql://myuser:mypassword@localhost:5432/demo?sslmode=disable
```

Ensure that this matches your local PostgreSQL setup.
## Continuous Integration and Continuous Deployment (CI/CD)

This project uses GitHub Actions for CI/CD. The workflow is defined in `.github/workflows/cicd-backend.yml`.

### Workflow Overview

The CI/CD pipeline consists of two main jobs: `build` and `deploy`.

1. **Build Job**:
   - Sets up Go environment
   - Builds and tests the application
   - Updates the `app.env` file with secrets
   - Builds a Docker image and pushes it to Docker Hub

2. **Deploy Job**:
   - Connects to an EC2 instance
   - Pulls the latest Docker image
   - Stops and removes the old container
   - Starts a new container with the updated image

### Triggering the Workflow

The CI/CD pipeline can be triggered in three ways:

1. **Push to main branch**: Any push to the `main` branch that includes changes in the `server/` directory or the workflow file itself will trigger the pipeline.

2. **Pull Request to main branch**: Opening or updating a pull request to the `main` branch with changes in the `server/` directory or the workflow file will trigger the pipeline.

3. **Manual Trigger**: You can manually trigger the workflow from the "Actions" tab in the GitHub repository.

### Testing the CI/CD Pipeline

To test the CI/CD pipeline:

1. Make changes to the backend code in the `server/` directory.
2. Commit and push your changes to the `main` branch, or create a pull request targeting the `main` branch.
3. Go to the "Actions" tab in your GitHub repository to see the workflow running.
4. Once the workflow completes successfully, you can verify the deployment by accessing the application at:

   ```
   http://13.250.17.15:9090
   ```

### Manual Trigger

To manually trigger the workflow:

1. Go to the "Actions" tab in your GitHub repository.
2. Select the "Build and deploy Backend to EC2" workflow.
3. Click on "Run workflow".
4. Choose the branch you want to run the workflow on and click "Run workflow".

### Environment Variables and Secrets

The workflow uses several secrets and environment variables:

- `DOCKER_USERNAME`: Your Docker Hub username
- `DOCKER_PASSWORD`: Your Docker Hub password
- `DB_STRING`: Database connection string
- `KEY`: Token symmetric key
- `EC2_KEY`: EC2 SSH key (base64 encoded)
- `EC2_HOST`: EC2 instance hostname or IP address

Ensure these secrets are properly set in your GitHub repository settings under "Secrets and variables" > "Actions".

## Test Coverage

Test coverage is an essential part of maintaining code quality and ensuring that your application behaves as expected. This project includes automated tests for the Go backend, and we use Go's built-in coverage tool to measure how much of our code is tested.

### Generating Coverage Reports

During the CI/CD pipeline, the coverage report is generated after running the tests. The steps are as follows:

1. **Run Tests**: The `go test` command is executed with the `-coverprofile` flag to create a coverage profile file (`coverage.out`).
2. **Generate HTML Report**: The coverage profile is then processed to create an HTML report (`coverage.html`), which provides a detailed view of which parts of the codebase are covered by tests.

### Coverage Reports in CI/CD

In the CI/CD workflow, the generated coverage reports are uploaded as artifacts. After a successful run of the workflow, you can download the coverage reports directly from the Actions tab in your GitHub repository.

### Analyzing Coverage

To view the coverage results:

1. After the workflow completes, navigate to the "Actions" tab in your GitHub repository.
2. Select the latest workflow run.
3. Under the "Artifacts" section, you’ll find the coverage reports available for download.
4. Download `coverage.html` and open it in your web browser to see a visual representation of test coverage across your codebase.

### Coverage Thresholds

For this project, we aim to maintain a high level of test coverage to ensure robustness. While there’s no strict coverage threshold defined, we recommend aiming for at least 80% coverage to have confidence in the stability of the application. You can adjust the coverage goals as necessary based on your team’s needs and the complexity of the codebase.

### Benefits of Coverage Reports

- **Identifying Uncovered Code**: The coverage report helps identify areas of the code that lack tests, allowing you to improve test coverage over time.
- **Improving Code Quality**: High test coverage often leads to better code quality and fewer bugs, as more scenarios are tested.
- **Confidence in Changes**: When making changes to the codebase, having comprehensive test coverage provides reassurance that existing functionality remains intact.

