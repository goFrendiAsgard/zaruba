# -- Building stage
FROM node:19 as builder

WORKDIR /app

COPY package*.json ./

RUN npm install 

COPY . .

RUN npm run build

# -- Running stage
FROM node:19

WORKDIR /app

COPY package*.json ./
COPY --from=builder /app/dist /app/dist

EXPOSE 8080
CMD [ "npm", "start" ]