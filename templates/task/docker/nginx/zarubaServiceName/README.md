# Nginx + Certbot

This is Nginx image with certbot.

To build the image you need to invoke `zaruba please buildZarubaServiceNameImage`.

To run as container you need to invoke `zaruba pleae runZarubaServiceName`.

# Configuration

Before building the image, you probably want to set some environment variables:

* `ZARUBA_SERVICE_NAME_BACKEND_ADDR`: Backend address (e.g: `localhost:3000`). Leave this blank if you only want to serve static files.
* `ZARUBA_SERVICE_NAME_SERVER_NAME`: Your server name (e.g: `stalchmst.com`)

Beside the variables, you probably also want to modify:

* `conf.d/default.conf.gotmpl`. This file is used to generate `conf.d/default.conf` that will be copied to `/etc/nginx/conf.d/default.conf`. This basic template already handled several use cases based on your environment variables.
* `html` folder. This folder contains all your static resources and will be mounted to `/usr/share/nginx/html`
* `nginx.conf.gotmpl`. This is is used to generate `nginx.conf` that will be copied to `/etc/nginx/nginx.conf`. Only modify this file if necessary.
* `letsencrypt`. In case of you already have `letsencrypt` in your previous system, you can copy `/etc/letsencrypt` content here. Otherwise, read the next section.

# Generating SSL Key

Let's say your `ZARUBA_SERVICE_NAME_BACKEND_ADDR` is `example.com` and your email is `john@gmail.com`

To generate ssl key for https, you need to perform:

```sh
docker exec -it zarubaContainerName bash

# staging (make sure if command really work)
certbot certonly --nginx --email john@gmail.com --agree-tos --no-eff-email --staging -d example.com -d www.example.com

# real command (if you are sure that the command should work)
certbot certonly --nginx --email john@gmail.com --agree-tos --no-eff-email --force-renewal -d example.com -d www.example.com
```

The certbot will create an `acme-challenge` and access it through `http` protocol.

After finishing the challenge, it will generate SSL key in your `letsencrypt` directory.

To load your SSL, you just need to rebuild/rerun `zarubaServiceName` again:

```sh
zaruba please runZarubaServiceName
```

# Renew SSL Key

You can run this command:

```sh
docker exec -it zarubaContainerName bash

# renew once
certbot renew --nginx

# renew using crontab
crontab -e
# Add this:
# 0 */12 * * * root certbot -q renew --nginx
```

# Adding static files

Adding static files is very easy. Just put your files in `html` directory. It will be loaded automatically, so you don't need to rerun `zarubaServiceName` again.

# Adding/changing backend URL

The default template only support one backend URL. To set this, you need to set `ZARUBA_SERVICE_NAME_BACKEND_ADDR` and rebuild/rerun `zarubaServiceName`:

```sh
zaruba please runZarubaServiceName
```