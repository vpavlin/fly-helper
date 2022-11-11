FROM registry.access.redhat.com/ubi9/ubi:9.0.0-1640

WORKDIR /opt/flyhelper/

ADD https://github.com/vpavlin/fly-helper/releases/latest/download/flyhelper /bin/flyhelper
#ADD flyhelper /bin/flyhelper

ADD entrypoint.sh .
RUN chmod +x entrypoint.sh /bin/flyhelper

ENTRYPOINT ./entrypoint.sh