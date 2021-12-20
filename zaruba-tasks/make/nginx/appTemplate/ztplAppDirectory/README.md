# Bitnami's Nginx + Certbot

This is a custom bitnami's nginx image. It contains custom configuration (`./my_server_block.cnf`) and some static files (`./html`)

# Ingredients

* [Bitnami's Nginx](https://github.com/bitnami/bitnami-docker-nginx)
* [Certbot](https://certbot.eff.org/)

We install certbot using root account.

# Running certbot

To run certbot you first need to access container's shell as root:

```sh
docker exec -u 0 -it <container-name> bash
```

# Generating certificate

```sh
# in container's shell as root
certbot certonly --dry-run
certbot certonly 
chmod -R 1001 /etc/letsencrypt
```

By default your webroot directory is `/opt/bitnami/nginx/html`

After invoking the command, your should find the ssl certificates under `/etc/letsencrypt/live/<domain.com>`

In order to use the certificate, you need to modify `my_server_block.conf`:

```conf
server {

    # -- SSL Configuration
    listen       443 ssl;
    listen  [::]:443 ssl;
    
    ssl_certificate      /etc/letsencrypt/live/<domain.com>/fullchain.pem;
    ssl_certificate_key  /etc/letsencrypt/live/<domain.com>/privkey.pem;
    
    ssl_session_cache    shared:SSL:1m;
    ssl_session_timeout  5m;
    
    ssl_ciphers  HIGH:!aNULL:!MD5;
    ssl_prefer_server_ciphers  on;

    # -- Non SSL Configuration
    # listen       443;
    # listen  [::]:443;

    # -- Serve static file 
    location / {
      root       /opt/bitnami/nginx/html;
      index      index.html index.htm;
      try_files  $uri $uri/ @backend;
    }

    # -- Reverse proxy
    location @backend {
        proxy_set_header X-Real-IP $remote_addr;
        proxy_set_header HOST $http_host;
        proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
        proxy_set_header X-NginX-Proxy true;
        proxy_pass       http://backend;
        proxy_redirect   off;
    }

}
```

# Renew certificate


```sh
# in container's shell as root
certbot renew --dry-run 
certbot renew 
```
