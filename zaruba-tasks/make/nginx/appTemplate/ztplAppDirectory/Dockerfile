FROM docker.io/bitnami/nginx:1.21.6

# become root to install certbot
USER 0
RUN apt update && \
    apt install certbot -y && \
    apt-get autoremove -yqq --purge && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# stop become root
USER 1001

COPY html /opt/bitnami/nginx/html
COPY server_blocks/my_server_block.conf /opt/bitnami/nginx/conf/server_blocks/my_server_block.conf 