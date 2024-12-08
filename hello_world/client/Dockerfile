FROM node:22.8-alpine

RUN npm install -g http-server@14.1.1

WORKDIR /app

COPY ./package*.json ./
RUN npm install

COPY . .
RUN npm run build

CMD [ "http-server", "dist" ]
