FROM node:22-alpine

WORKDIR /cache_modules

COPY package.json package-lock.json ./
RUN npm install

WORKDIR /app
COPY . .

