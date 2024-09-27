# syntax=docker/dockerfile:1

FROM golang:1.23.1

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY mock_external/ ./mock_external
COPY service/ ./service
COPY api/ ./api
COPY song_detail_scheme.yml ./
COPY scripts/exec.sh ./


RUN cd mock_external && go mod download
RUN cd service && go mod download
RUN cd api && go mod download

COPY go* ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /mock_external ./mock_external/cmd 
RUN CGO_ENABLED=0 GOOS=linux go build -o /service ./service/cmd 

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/reference/dockerfile/#expose
EXPOSE 8080

# Run
CMD ["./exec.sh"]

