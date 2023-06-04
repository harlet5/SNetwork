FROM node:19.8.1-slim

WORKDIR /app

# Install dependencies
COPY Frontend/package*.json ./
RUN npm install

# Copy frontend source code
COPY Frontend/. .

# Build frontend
RUN npm run build

# Expose port and start server
EXPOSE 8080
CMD ["npm", "run", "dev"]
