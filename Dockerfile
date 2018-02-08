FROM alpine
ADD /wiki_unix /
ADD /templates /templates
CMD ["./wiki_unix"]