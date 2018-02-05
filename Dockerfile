FROM alpine
ADD bin/wiki /
ADD bin/templates /templates
RUN mkdir -p /pages
CMD ["./wiki"]