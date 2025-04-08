@echo off

REM Creazione directory principale del progetto
mkdir project
cd project

echo Creazione struttura di directory...
mkdir backend
mkdir backend\controllers
mkdir backend\routes
mkdir backend\middleware
mkdir frontend
mkdir database

echo Creazione file backend vuoti...
echo. > backend\main.go
echo. > backend\routes\routes.go
echo. > backend\middleware\auth.go
echo. > backend\controllers\auth.go
echo. > backend\controllers\ticket.go
echo. > backend\controllers\ecommerce.go
echo. > backend\controllers\domain.go
echo. > backend\controllers\contact.go
echo. > backend\controllers\service.go

echo Creazione file di configurazione (Procfile backend/frontend)...
echo web: go run main.go > backend\Procfile
echo web: serve -s dist/angular > frontend\Procfile

echo Creazione script SQL di inizializzazione database...
(
    echo CREATE TABLE users (
    echo    id SERIAL PRIMARY KEY,
    echo    name VARCHAR(100),
    echo    email VARCHAR(100)
    echo );
    echo.
    echo CREATE TABLE tickets (
    echo    id SERIAL PRIMARY KEY,
    echo    user_id INT REFERENCES users(id),
    echo    subject VARCHAR(100),
    echo    status VARCHAR(20)
    echo );
) > database\init.sql

echo Creazione file .gitignore...
(
    echo # Dipendenze Node/Angular
    echo node_modules/
    echo frontend/angular/dist/
    echo
    echo # File di log e output build
    echo logs/
    echo *.log
    echo npm-debug.log*
    echo yarn-debug.log*
    echo yarn-error.log*
    echo dist/
    echo
    echo # File compilati Go
    echo *.exe
    echo *.dll
    echo *.so
    echo *.dylib
    echo *.test
    echo *.out
    echo
    echo # Configurazioni environment e IDE
    echo .env
    echo .vscode/
    echo .idea/
) > .gitignore

echo Creazione README.md con istruzioni Railway...
(
    echo # Project Setup for Railway Deployment
    echo Questo progetto ^e predisposto per il deploy sulla piattaforma Railway. Contiene un backend Go, un frontend Angular e un database PostgreSQL.
    echo.
    echo ## Istruzioni per il Deploy
    echo 1. Collega questo repository a un nuovo progetto su Railway.
    echo 2. Aggiungi un database **PostgreSQL** come plug-in nel progetto Railway.
    echo 3. Effettua il deploy: Railway avvier^a automaticamente il backend e il frontend secondo quanto definito in `railway.json` e nei Procfile.
    echo 4. Opzionale Esegui lo script SQL in `database/init.sql` sul database per creare le tabelle iniziali.
) > README.md

echo Creazione file railway.json per configurazione Railway...
(
    echo {
    echo   "$schema": "https://railway.app/railway.schema.json",
    echo   "services": [
    echo     {
    echo       "serviceName": "backend",
    echo       "path": "backend",
    echo       "startCommand": "go run main.go",
    echo       "envVars": {
    echo         "PORT": "8080"
    echo       }
    echo     },
    echo     {
    echo       "serviceName": "frontend",
    echo       "path": "frontend/angular",
    echo       "buildCommand": "npm install && npm run build",
    echo       "startCommand": "npx serve -s dist/angular -l $PORT"
    echo     },
    echo     {
    echo       "serviceName": "database",
    echo       "startCommand": "postgres",
    echo       "envVars": {
    echo         "POSTGRES_USER": "postgres",
    echo         "POSTGRES_PASSWORD": "postgres",
    echo         "POSTGRES_DB": "postgres"
    echo       }
    echo     }
    echo   ]
    echo }
) > railway.json

echo Creazione file docker-compose.yml per sviluppo locale...
(
    echo version: '3.9'
    echo services:
    echo   backend:
    echo     image: "golang:1.20"
    echo     working_dir: /app
    echo     volumes:
    echo       - ./backend:/app
    echo     command: "go run main.go"
    echo     ports:
    echo       - "8080:8080"
    echo     environment:
    echo       - PGHOST=database
    echo       - PGUSER=postgres
    echo       - PGPASSWORD=postgres
    echo       - PGDATABASE=postgres
    echo       - PGPORT=5432
    echo   frontend:
    echo     image: "node:18"
    echo     working_dir: /app
    echo     volumes:
    echo       - ./frontend/angular/dist/angular:/app
    echo     command: "npx serve -s . -l 5000"
    echo     ports:
    echo       - "5000:5000"
    echo   database:
    echo     image: "postgres:13"
    echo     environment:
    echo       - POSTGRES_USER=postgres
    echo       - POSTGRES_PASSWORD=postgres
    echo       - POSTGRES_DB=postgres
    echo     ports:
    echo       - "5432:5432"
) > docker-compose.yml

echo Installazione Angular CLI globalmente...
npm install -g @angular/cli

echo Creazione progetto Angular in .\frontend ...
cd frontend
ng new angular --defaults --skip-git

echo Installazione di serve nella applicazione Angular...
cd angular
npm install serve

echo Inizializzazione repository Git e primo commit...
cd ..
cd ..
git init
git add .
git commit -m "Initial commit"