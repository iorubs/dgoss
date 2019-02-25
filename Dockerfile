FROM docker:18

LABEL maintainer "ruben.vasconcelos3@mail.dcu.ie"

RUN apk upgrade \
    && apk add bash

ENV GOSS_URL https://github.com/aelsabbahy/goss/releases/download/v0.3.6/goss-linux-amd64
ENV DGOSS_URL https://raw.githubusercontent.com/aelsabbahy/goss/master/extras/dgoss/dgoss
RUN wget $GOSS_URL -O /usr/local/bin/goss \
    && wget $DGOSS_URL -O /usr/local/bin/dgoss \
    && chmod +rx /usr/local/bin/goss /usr/local/bin/dgoss

# Set cp stategy because we don't expect every image to contain goss binaries.
ENV GOSS_FILES_STRATEGY cp

WORKDIR /src
