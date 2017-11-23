FROM r.xiaozhou.net/alpine/base:latest
MAINTAINER Timothy Ye <yexiaozhou2003@gmail.com>
WORKDIR /

ADD ./app/views/index.html /app/views/index.html
ADD biturl /biturl
RUN chmod +x /biturl


EXPOSE 8000
ENTRYPOINT ["./biturl"]

