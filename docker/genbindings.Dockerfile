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
        libutf8proc-dev \
        libutf8proc2 \
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
        cmake \
        build-essential && \
    apt-get clean

RUN mkdir -p /usr/local/src/scintilla && \
    git clone 'https://github.com/mirror/scintilla.git' /usr/local/src/scintilla && \
    git -C /usr/local/src/scintilla checkout rel-5-5-2

RUN \
    cd /usr/local/src/scintilla/qt/ScintillaEditBase && \
    qmake && \
    make && \
    cd /usr/local/src/scintilla/qt/ScintillaEdit && \
    python3 WidgetGen.py && \
    qmake && \
    make

# Custom pkg-config definitions

RUN mkdir -p /usr/local/lib/pkgconfig

RUN 	echo 'includedir=/usr/include/x86_64-linux-gnu/qt5/Qsci/' \
	'\n' \
	'\nName: QScintilla' \
	'\nDescription: Qt5 port of the Scintilla source code editing widget' \
	'\nURL: http://www.riverbankcomputing.co.uk/software/qscintilla' \
	'\nVersion: 2.13.3' \
	'\nRequires: Qt5Widgets, Qt5PrintSupport' \
	'\nLibs: -lqscintilla2_qt5' \
	'\nCflags: -I${includedir}' \
	> /usr/local/lib/pkgconfig/QScintilla.pc

RUN echo 'includedir=/usr/include/x86_64-linux-gnu/qt6/Qsci/' \
	'\n' \
	'\nName: QScintilla6' \
	'\nDescription: Qt6 port of the Scintilla source code editing widget' \
	'\nURL: http://www.riverbankcomputing.co.uk/software/qscintilla' \
	'\nVersion: 2.13.3' \
	'\nRequires: Qt6Widgets, Qt6PrintSupport' \
	'\nLibs: -lqscintilla2_qt6' \
	'\nCflags: -I${includedir}' \
	> /usr/local/lib/pkgconfig/QScintilla6.pc

RUN echo 'srcdir=/usr/local/src/scintilla/' \
	'\n' \
	'\nName: ScintillaEdit' \
	'\nDescription: Scintilla upstream Qt port' \
	'\nURL: https://www.scintilla.org/' \
	'\nVersion: 5.5.2' \
	'\nRequires: Qt5Widgets' \
	'\nLibs: -L${srcdir}/bin -lScintillaEdit' \
	'\nCflags: -include stdint.h -I${srcdir}/qt/ScintillaEdit -I${srcdir}/qt/ScintillaEditBase -I${srcdir}/include -I${srcdir}/src' \
	> /usr/local/lib/pkgconfig/ScintillaEdit.pc

RUN wget -q https://go.dev/dl/go1.23.1.linux-amd64.tar.gz -O /tmp/golang.tar.gz && \
    tar -C /usr/local -xzf /tmp/golang.tar.gz

COPY qtermwidget-win qtermwidget-win

RUN cd /qtermwidget-win && mkdir build && cd build && \
   cmake .. && make && make install && \
  cd / && rm -rf qtermwidget-win

ENV PATH=/usr/local/go/bin:$PATH

ENV GOFLAGS=-buildvcs=false
