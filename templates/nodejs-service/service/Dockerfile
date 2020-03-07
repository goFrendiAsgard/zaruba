FROM node:lts
LABEL SERVICE servicename

# define environments
ENV SERVICENAME_HTTP_PORT 3010

# Create and set the workdir
RUN mkdir -p /app
WORKDIR /app

# copy current directory to workdir
ADD . /app
RUN npm install --production

CMD npm start
