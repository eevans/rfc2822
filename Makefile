
GOROOT ?= /usr/share/go

include $(GOROOT)/src/Make.inc

TARG=rfc2822
GOFILES=\
	rfc2822.go

include $(GOROOT)/src/Make.pkg
