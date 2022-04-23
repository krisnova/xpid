# =========================================================================== #

CC     = clang
CFLAGS = -ltcap -lpcap
DEPS   = xp.c
OBJ    = xp.o

%.o: %.c $(DEPS)
	$(CC) -c -o $@ $< $(CFLAGS)

tlsdump: $(OBJ) ## Generate the main tlsdump program
	$(CC) -o $@ $^ $(CFLAGS)

clean:  ## Show help messages for make targets
		rm -rf *.o
		rm -rf xp

install:  ## Install tlsdump to /usr/local/bin
		chmod +x xp
		cp xp /usr/bin/xp

.PHONY: help
help:  ## Show help messages for make targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}'