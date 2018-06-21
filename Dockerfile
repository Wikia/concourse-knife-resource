FROM debian:stretch-slim

RUN apt-get update && \
    apt-get install -y curl jq && \
    curl -O https://packages.chef.io/files/stable/chefdk/3.0.36/debian/9/chefdk_3.0.36-1_amd64.deb && \
    dpkg -i chefdk_3.0.36-1_amd64.deb && \
    rm -rf chefdk_3.0.36-1_amd64.deb && \
    apt-get purge -y curl && \
    apt-get autoremove -y && \
    rm -rf /var/lib/apt/lists/*

RUN useradd -m -s /bin/bash chef

ADD assets/ /opt/resource/
RUN chmod +x /opt/resource/*

RUN mkdir /root/.chef

COPY knife.rb /root/.chef/knife.rb
COPY credentials /root/.chef/credentials
