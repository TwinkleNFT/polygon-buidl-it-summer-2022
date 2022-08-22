####################################################################################################
## Builder
####################################################################################################
FROM golang:alpine AS builder

# RUN apt update 
RUN update-ca-certificates

# Create appuser
ENV USER=myapp
ENV UID=10001

RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "/nonexistent" \
    --shell "/sbin/nologin" \
    --no-create-home \
    --uid "${UID}" \
    "${USER}"


WORKDIR /myapp

COPY ./ .

RUN env GOOS=linux CGO_ENABLED=0 go build -o dist/myapp

####################################################################################################
## Final image
####################################################################################################
FROM alpine 
# Import from builder.
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group


WORKDIR /myapp

# Copy our build
COPY --from=builder /myapp/dist/myapp ./

# Use an unprivileged user.
USER myapp:myapp

EXPOSE 3000

CMD ["/myapp/myapp"]
