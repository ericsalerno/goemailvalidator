FROM golang

# Copy local package to the containers workspace
ADD . /go/src/github.com/ericsalerno/goemailvalidator

# Build package in container
RUN go install github.com/ericsalerno/goemailvalidator

# Download blacklist.conf file
ADD https://raw.githubusercontent.com/martenson/disposable-email-domains/master/disposable_email_blacklist.conf /go/bin/blacklist.conf

# Set container entrypoint to compiled binary
ENTRYPOINT /go/bin/goemailvalidator

# Expose port 8081
EXPOSE 8081