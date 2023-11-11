FROM node:18

WORKDIR /app

RUN apt-get update -y && apt-get upgrade -y
RUN npm install -g npm@9.6.1

COPY package*.json ./
COPY tsconfig.json ./
COPY ../.env ./

COPY src ./src
COPY public ./public

RUN npm i --force

EXPOSE 3000

CMD ["npm", "run", "start"]