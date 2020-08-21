FROM node:latest
LABEL SERVICE servicename

# define environments
ENV SERVICE_NAME servicename

# Create app directory
WORKDIR /app

# Install app dependencies
COPY package.json package-lock.json ./
RUN npm ci --only=production

# Bundle app source
COPY . .
RUN npm run build

EXPOSE 3000
CMD npm start
