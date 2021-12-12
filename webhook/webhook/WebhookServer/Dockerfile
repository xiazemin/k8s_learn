FROM alpine
WORKDIR /
RUN mkdir -p /etc/webhook/certs/
ADD ./tmp/server-cert.pem /etc/webhook/certs/cert.pem
ADD ./tmp/server-key.pem /etc/webhook/certs/key.pem
COPY ./webhook webhook
RUN chmod +x webhook
CMD ["/bin/sh"]