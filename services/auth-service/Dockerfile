
FROM node:14-alpine

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

# Build the TypeScript code
RUN npm run build

# Expose the port on which the application will run
EXPOSE 3000

# Start the application
CMD ["npm", "start"]
