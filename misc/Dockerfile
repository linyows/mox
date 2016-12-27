FROM golang:1.6

RUN apt-get update -qq && apt-get install -y -qq wget unzip supervisor
RUN wget -q https://github.com/linyows/mox/releases/download/v0.2.1/linux_amd64.zip && \
    unzip -q linux_amd64.zip && \
    rm -rf linux_amd64.zip && \
    mv mox /usr/bin/
ADD supervisord.conf /etc/supervisor/conf.d/mox.conf

ENV MOX_CONF="/mox/mox.conf" \
    MOX_ROOT="/mox/res" \
    MOX_ADDR="0.0.0.0:3000"

RUN mkdir /mox
WORKDIR /mox
ADD ../testdata/rest /mox

CMD ["supervisord", "-n"]
