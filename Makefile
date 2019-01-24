MACHINE := $(shell uname -m)

ifeq ($(MACHINE), x86_64)
libdir = /usr/lib64
endif
ifeq ($(MACHINE), i686)
libdir = /usr/lib
endif

libnss_etcd.so.2: cmd/libnss_etcd/*.go *.go *.h *.c
	go build -buildmode=c-shared -o $@ ./cmd/libnss_etcd
	chmod +x $@

build: libnss_etcd.so.2

install: libnss_etcd.so.2
	cp $< $(libdir)/$<

clean:
	rm -rf libnss_etcd.so.2 libnss_etcd.so.h

.PHONY: build install clean
