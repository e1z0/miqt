FROM debian:bookworm

#RUN sed -i 's|deb.debian.org/debian|archive.debian.org/debian|g' /etc/apt/sources.list \
# && sed -i 's|deb.debian.org/debian-security|archive.debian.org/debian-security|g' /etc/apt/sources.list \
# && echo 'Acquire::Check-Valid-Until "false";' > /etc/apt/apt.conf.d/99no-check-valid-until \
# && apt-get update

RUN DEBIAN_FRONTEND=noninteractive apt-get update && \
    apt-get install -qyy gnupg2 ca-certificates wget && \
    apt-get clean

RUN DEBIAN_FRONTEND=noninteractive apt-get update && \
    apt-get install --no-install-recommends -qyy \
        qtbase5-dev \
        qtmultimedia5-dev \
        qtpdf5-dev \
        qtscript5-dev \
        libqt5svg5-dev \
        libqt5webkit5-dev \
        qtwebengine5-dev \
        qt6-base-dev \
        qt6-charts-dev \
        qt6-declarative-dev \
        qt6-multimedia-dev \
        qt6-pdf-dev \
        qt6-svg-dev \
        qt6-webengine-dev \
        libqscintilla2-qt5-dev \
        libqscintilla2-qt6-dev \
        clang \
        git \
        ca-certificates \
        pkg-config \
        build-essential && \
    apt-get clean

RUN wget -q https://go.dev/dl/go1.23.1.linux-amd64.tar.gz -O /tmp/golang.tar.gz && \
    tar -C /usr/local -xzf /tmp/golang.tar.gz

COPY qtermwidget-win qtermwidget-win

RUN apt-get install -qyy cmake

RUN cd /qtermwidget-win && mkdir build && cd build && \
   cmake .. && make && make install && \
  cd / && rm -rf qtermwidget-win

ENV PATH=/usr/local/go/bin:$PATH

#ENV GOFLAGS=-buildvcs=false
