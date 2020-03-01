FROM node:lts
LABEL SERVICE servicename

# define environments
ENV port 3000

# Create and set the workdir
RUN mkdir -p /app
WORKDIR /app

# copy current directory to workdir
ADD . /app
RUN npm install --production

CMD npm start
