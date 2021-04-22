FROM golang:1.16.3-buster AS builder

WORKDIR .

COPY . /opt/build/

# Download all the dependencies
RUN cd /opt/build/ && make compileLinux

from debian 

COPY --from=builder /opt/build/bin/linux/jwtExample ./jwtExample

# This container exposes port 8080 to the outside world
EXPOSE 9000 

# Run the executable 
CMD ["./jwtExample"]