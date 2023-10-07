#FROM node:20-alpine3.17 as builder
FROM node:20-alpine3.17
LABEL authors="xederro"

RUN mkdir /app
COPY . /app
WORKDIR /app

RUN npm install
RUN npm run build
RUN chmod +x /app/build/index.js

#FROM node:20-alpine3.17
#RUN mkdir /app
#COPY --from=builder /app/build /app

CMD [ "node", "./build" ]