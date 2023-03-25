# Use the official Node.js image as base
FROM google/cloud-sdk:latest

# Set the working directory in the container
WORKDIR /app

# Install Java(11 or higher)
RUN apt-get update && \
    apt-get install -y openjdk-11-jre-headless && \
    rm -rf /var/lib/apt/lists/*

COPY firebase.json /app
# COPY .firebaserc /app

# Install the Firebase CLI globally
# RUN npm install -g firebase-tools

# Expose the ports used by the emulator
EXPOSE 8080
# EXPOSE 4000

# Start the Firestore emulator
# CMD ["firebase", "emulators:start", "--only", "firestore", "--project", "my-project-id"]
CMD ["gcloud", "beta", "emulators", "firestore", "start", "--host-port=localhost:8080"]