# Guinx

A web application with Go backend and SvelteKit frontend.

## Prerequisites

- Docker and Docker Compose installed on your machine
- Git for cloning the repository

## Quick Start with Docker

1. Clone the repository:
```bash
git clone https://github.com/yourusername/guinx.git
cd guinx
```


2. Environment Setup

You don't need PostgreSQL installed on your machine as it runs in a Docker container. However, you need to set up the environment variables:

Create `.env` file in the project root:


Database Configuration
```
DB_USER=guinx # PostgreSQL username
DB_PASSWORD=choose_secure_pwd # Choose a secure password
DB_NAME=guinx # Database name
DB_PORT=5433 # Using 5433 to avoid conflicts with local PostgreSQL
```
Security
```
JWT_SECRET=generate_secure_key # Used for token signing (see below)
```
API Configuration
```
BACKEND_PORT=8080 # Backend API port
FRONTEND_PORT=3000 # Frontend port
```

### About Environment Variables

#### JWT Secret
The JWT_SECRET is used to sign authentication tokens. Generate a secure key using one of these methods:
Method 1: Using openssl (recommended)
```bash
openssl rand -base64 32
```
Method 2: Using node
```bash
node -e "console.log(require('crypto').randomBytes(32).toString('base64'));"
```
Method 3: Using python
```bash
python -c "import secrets; print(secrets.token_urlsafe(32))"
```

Copy the generated string as your JWT_SECRET.

#### Database Password
Choose a secure password for your database. You can generate one using:
```bash
openssl rand -base64 32
```


#### Port Configuration
- The database runs on port 5433 to avoid conflicts with any local PostgreSQL installation
- Backend API runs on port 8080
- Frontend runs on port 3000

Make sure these ports are available on your machine.

### Database Access

While you don't need PostgreSQL installed locally, you might want to access the database for development. You can:

1. Use the Docker container's PostgreSQL:

3. Build and run with Docker Compose:

```bash
docker compose up --build
```
Connect to database container
```bash
docker-compose exec db psql -U guinx -d guinx
```
Or from your local machine (if you have psql installed)
```bash
psql -h localhost -p 5433 -U guinx -d guinx
```

2. Use a database management tool like:
   - [pgAdmin](https://www.pgadmin.org/)
   - [DBeaver](https://dbeaver.io/)
   - [TablePlus](https://tableplus.com/)

Connection details:
Host: localhost
Port: 5433
Database: guinx
Username: guinx
Password: (the one you set in .env)


### Security Notes

1. Never commit your `.env` file to version control
2. Use strong, unique passwords for production
3. Change default ports in production
4. Consider using a secrets management service in production


The application will be available at:
- Frontend: http://localhost:3000
- Backend API: http://localhost:8080
- Database: localhost:5433 (PostgreSQL)

## Docker Commands

### Basic Commands

Start the application
```bash
docker-compose up
```

Start in detached mode (background)
```bash
docker-compose up -d
```

Stop the application
```bash
docker-compose down
```

View logs
```bash
docker-compose logs
```

View logs for specific service
```bash
docker-compose logs frontend
docker-compose logs backend
docker-compose logs db
```

Rebuild containers
```bash
docker-compose up --build
```

### Maintenance Commands

Remove all containers and volumes
```bash
docker-compose down -v
```
View all containers
```bash
docker ps -a
```

Restart a specific service
```bash
docker-compose restart frontend
docker-compose restart backend
```
## Project Structure

```
guinx/
├── frontend/ # SvelteKit frontend application
├── backend/ # Go backend application
├── docker-compose.yml # Docker Compose configuration
└── .env # Environment variables
```


## Development without Docker

### Frontend
```bash
cd frontend
npm install
npm run dev
```

### Backend
```bash
cd backend
go mod download
go run cmd/server/main.go
```


## Environment Variables

### Backend Variables
- `DATABASE_URL`: PostgreSQL connection string
- `JWT_SECRET`: Secret key for JWT tokens
- `PORT`: Server port (default: 8080)

### Frontend Variables
- `VITE_API_URL`: Backend API URL

## API Documentation

The backend API is available at `http://localhost:8080` with the following endpoints:

### Authentication
- `POST /api/auth/login`: User login
- `POST /api/auth/register`: User registration

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request
