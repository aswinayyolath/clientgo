FROM ubuntu

COPY ./clientgo ./clientgo

ENTRYPOINT [ "./clientgo" ]