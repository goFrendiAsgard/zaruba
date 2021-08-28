FROM nginx
COPY nginx.conf /etc/nginx/conf.d/nginx.conf
RUN apt update && \
    apt install -y certbot && \
    apt clean
COPY html /usr/share/nginx/html