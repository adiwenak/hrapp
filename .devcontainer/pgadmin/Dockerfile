FROM dpage/pgadmin4:7.4

# Copy in the servers.json to avoid having to mount the file path,
# which requires everyone to update their Docker Desktop
# resource sharing.
COPY ./servers.json /pgadmin4/servers.json
