FROM node:12

ARG APP_NAME
ENV APP_NAME=$APP_NAME

# Create app directory
RUN mkdir -p /app

WORKDIR /app

# Install app dependencies
# A wildcard is used to ensure both package.json AND package-lock.json are copied
# where available (npm@5+)

COPY workspace/$APP_NAME/package.json ./

COPY workspace/$APP_NAME/yarn.lock ./

RUN yarn install

# copy app source
COPY ../workspace/$APP_NAME .

#RUN mv .env.example .env

EXPOSE 8081
CMD [ "node", "server.js" ]