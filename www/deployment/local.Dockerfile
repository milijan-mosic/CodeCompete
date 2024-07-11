FROM node:22-alpine AS build

WORKDIR /code
RUN npm update
COPY package*.json ./
RUN npm ci
COPY . .

FROM node:22-alpine

WORKDIR /code
COPY --from=build /code/node_modules ./node_modules
COPY --from=build /code .
CMD npm run dev
