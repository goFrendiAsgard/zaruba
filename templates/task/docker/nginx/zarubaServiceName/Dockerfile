FROM nginx
COPY nginx.conf /etc/nginx/conf.d/nginx.conf
RUN apk add certbot certbot-nginx
COPY html /usr/share/nginx/html