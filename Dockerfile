FROM node:latest as node
WORKDIR /usr/src/app/front
COPY . .
RUN npm install
CMD npm run start
EXPOSE 4200
