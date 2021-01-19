FROM golang:1.13.7

ENV GO111MODULE="on" \
    GOPROXY="https://goproxy.cn"

ADD . /home/micro-falsework

RUN cd /home/micro-falsework && go build .

#CMD /home/micro-falsework/micro-falsework
ENTRYPOINT /home/micro-falsework/micro-falsework --path=$0