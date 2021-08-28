FROM nginx
RUN apt update && \
    apt install -y certbot && \
    apt clean
COPY nginx.conf /etc/nginx/nginx.conf
COPY conf.d/default.conf /etc/nginx/conf.d/default.conf
COPY html /usr/share/nginx/html