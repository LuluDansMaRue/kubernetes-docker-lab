FROM node:10.15.1-alpine

ENV APP_PATH /data

# Create folder
RUN mkdir ${APP_PATH}

# Set workdir
WORKDIR ${APP_PATH}

RUN ls -la

# Install dependencies
RUN npm install

# Run the app
CMD ["npm", "run", "serve"]