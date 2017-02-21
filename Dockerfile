FROM scratch
EXPOSE 9090
COPY /ifconfigme /ifconfigme
CMD ["/ifconfigme"]
