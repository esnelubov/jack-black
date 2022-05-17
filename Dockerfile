FROM golang:1.17
WORKDIR /app
COPY . .
RUN rm .env
RUN mv .docker.env .env
RUN make -f Makefile build install-migrate
CMD cmd/backend
