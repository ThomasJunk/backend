# SPDX-License-Identifier: MIT
# © 2021 Thomas Junk
GOCMD=go
GOBUILD=$(GOCMD) build
NAME=backend
PATHTOMAIN=./cmd/$(NAME)
all:	build
build:
	$(GOBUILD) -o $(NAME) $(PATHTOMAIN)/main.go
