FROM ubuntu:20.04
COPY main /
RUN apt-get update && apt-get install -y \
    ncat iproute2 socat nmap openssh-client mtr iputils-ping wget curl vim \
 && rm -rf /var/lib/apt/lists/*
CMD /main
