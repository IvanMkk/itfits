FROM node:18.16.0

COPY package*.json ./

RUN npm install
COPY . .
RUN npm run build
RUN npm install -g serve
EXPOSE 8080

CMD ["npm", "run", "serve"]
