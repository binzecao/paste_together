FROM alpine:3.12.0

#RUN mkdir /project

COPY bin/paste_together /project

EXPOSE 46699

CMD ["/project/paste_together", "--listen=:46699"]

