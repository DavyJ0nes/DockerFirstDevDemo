# Starting Image
FROM ubuntu

# Commands to Run when building image
RUN apt-get update && \
    apt-get install -y iputils-ping

# Command to run as default when starting container
ENTRYPOINT exec bash
