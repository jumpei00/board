FROM node:16.14.2 as react-build

WORKDIR /app

COPY package.json yarn.lock ./
RUN yarn

COPY . ./
RUN yarn build

FROM nginx:latest

WORKDIR /etc/nginx
COPY ./nginx/nginx.conf /etc/nginx/nginx.conf

COPY --from=react-build /app/build /usr/share/nginx/html
CMD ["nginx", "-g", "daemon off;"]