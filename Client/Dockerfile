# pull the base image
FROM node:lts-alpine

# set working directory
WORKDIR /app

# install app dependencies
COPY Client/package.json ./

RUN yarn

COPY Client ./

RUN yarn build

# 포트 노출
EXPOSE 3000

# 앱 실행
CMD ["yarn", "start"]