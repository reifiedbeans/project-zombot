NAME    = pzombot
BINDIR  = bin
FLAGS   = -trimpath

.PHONY: clean

all: bot

bot:
	go build $(FLAGS) -o $(BINDIR)/$(NAME) cmd/bot/main.go

clean:
	rm -rf $(BINDIR)/$(NAME)
