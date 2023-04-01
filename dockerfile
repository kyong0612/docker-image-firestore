# Use the official Node.js image as base
FROM google/cloud-sdk:latest

# Set the working directory in the container
WORKDIR /app

# Install Java(11 or higher)
RUN apt-get update && \
    apt-get install -y openjdk-11-jre-headless && \
    rm -rf /var/lib/apt/lists/*

COPY firebase.json /app


EXPOSE 8080

CMD ["gcloud", "emulators", "firestore", "start", "--host-port=0.0.0.0:8080"]