FROM alpine:3.12.0

RUN mkdir /project && mkdir /project/data

COPY bin/ /project/

WORKDIR /project

EXPOSE 46699

CMD ["/project/paste_together", "--listen=:46699"]
