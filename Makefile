#!/usr/bin/make --no-print-directory --jobs=1 --environment-overrides -f

VERSION_TAGS += STRCASES
STRCASES_MK_SUMMARY := go-corelibs/strcases
STRCASES_MK_VERSION := v0.0.0

include CoreLibs.mk
