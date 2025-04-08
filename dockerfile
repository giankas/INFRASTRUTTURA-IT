# Stage 1: Build del backend Go
FROM golang:1.19-alpine AS backend-builder

WORKDIR /app

# Copia i file Go
COPY backend/go.mod backend/go.sum ./
RUN go mod download

# Copia il codice Go
COPY backend ./

WORKDIR /app/backend

# Compila il progetto Go
RUN go build -o out main.go


# Stage 2: Build del frontend Angular
FROM node:16-alpine AS frontend-builder

WORKDIR /frontend

# Copia i file del frontend (Angular)
COPY frontend/package.json frontend/package-lock.json ./
RUN npm install

COPY frontend ./

# Costruisci il progetto Angular
RUN npm run build --prod


# Stage 3: Setup Runtime (per eseguire l'app)
FROM nginx:alpine

# Copia il build del frontend (Angular) nella directory Nginx
COPY --from=frontend-builder /frontend/dist /usr/share/nginx/html

# Copia il binario del backend dal primo stage
COPY --from=backend-builder /app/backend/out /root/

# Comando per avviare il backend e servire il frontend
CMD ["/root/out"]
