FROM node:22.8-alpine as builder

WORKDIR /app

COPY ./package*.json ./
RUN npm install

COPY . .
RUN npm run build

# ---------------------------

# Run application

FROM nginx:1.27-alpine

RUN adduser -D -g 'www' www

COPY --from=builder /app/dist /usr/share/nginx/html

CMD nginx -g 'daemon off;'
