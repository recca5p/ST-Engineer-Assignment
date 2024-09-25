# Use an official Node.js runtime as a parent image
FROM node:18-alpine

# Set the working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json files to the container
COPY package*.json ./

# Install dependencies
RUN yarn install --frozen-lockfile

# Copy the rest of the application to the container
COPY . .

# Build the Next.js application
RUN yarn run build

# Expose the port that the app will run on
EXPOSE 3000

# Start the Next.js application
CMD ["yarn", "run", "start"]