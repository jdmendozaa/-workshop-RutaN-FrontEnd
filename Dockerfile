# Parte 1

FROM node:latest as node

WORKDIR /app

COPY . .

RUN npm install

RUN npm run build --prod

# Parte 2

FROM nginx:alpine

COPY --from=node /app/dist/praxis-fe /usr/share/nginx/html
