# Use the official Node.js image as base
FROM node:18

# Set the working directory in the container
WORKDIR /app

# Install the Firebase CLI globally
RUN npm install -g firebase-tools

# Expose the ports used by the emulator
EXPOSE 8080

# Start the Firestore emulator
CMD ["firebase", "emulators:start", "--only", "firestore"]
