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
certbot certonly 
```

Also modify `my_server_block` to use the certificate.

The ssl certificates should be located under `/etc/letsencrypt/live/<domain.com>`

# Renew certificate


```sh
# in container's shell as root
certbot renew --dry-run 
certbot renew 
```
