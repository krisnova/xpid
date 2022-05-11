# =========================================================================== #
#            MIT License Copyright (c) 2022 Kris Nóva <kris@nivenly.com>      #
#                                                                             #
#                 ┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓                 #
#                 ┃   ███╗   ██╗ ██████╗ ██╗   ██╗ █████╗   ┃                 #
#                 ┃   ████╗  ██║██╔═████╗██║   ██║██╔══██╗  ┃                 #
#                 ┃   ██╔██╗ ██║██║██╔██║██║   ██║███████║  ┃                 #
#                 ┃   ██║╚██╗██║████╔╝██║╚██╗ ██╔╝██╔══██║  ┃                 #
#                 ┃   ██║ ╚████║╚██████╔╝ ╚████╔╝ ██║  ██║  ┃                 #
#                 ┃   ╚═╝  ╚═══╝ ╚═════╝   ╚═══╝  ╚═╝  ╚═╝  ┃                 #
#                 ┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛                 #
#                                                                             #
#                        This machine kills fascists.                         #
#                                                                             #
# =========================================================================== #

TARGET  := xpid-bpf
CFLAGS  ?= -I /usr/include
LDFLAGS ?=
LIBS     = -l bpf
STYLE    = Google

all: autogen probes build ## Build everything

.PHONY: clean
clean: ## Clean objects
	rm -vf $(TARGET)
	rm -vf *.o
	rm -vf *.ll
	rm -vf $(PROBE)

format: ## Format the code
	@echo "  ->  Formatting code"
	@clang-format -i -style=$(STYLE) *.c *.h

build: ## Build the userspace program
	@echo "  ->  Building"
	clang $(CFLAGS) $(LDFLAGS) $(LIBS) -o $(TARGET) xpid-bpf.c -Wl,

install: ## Install the target, and probe
	mkdir -p /usr/share/xpid
	cp $(TARGET) /usr/bin/$(TARGET)
	cp -v probe.*.o /usr/share/xpid/

.PHONY: probes
probes: probe-tracepoints ## Compile eBPF probes
	@echo "  ->  Building probe"

autogen:
	bpftool btf dump file /sys/kernel/btf/vmlinux format c > vmlinux.h

probe-tracepoints: probe = probe.tracepoints
probe-tracepoints:
	clang -S \
	    -target bpf \
	    -D __BPF_TRACING__ \
	    $(CFLAGS) \
	    -Wall \
	    -Werror \
	    -O2 -emit-llvm -c -g $(probe).c
	llc -march=bpf -filetype=obj -o $(probe).o $(probe).ll

.PHONY: help
help:  ## Show help messages for make targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(firstword $(MAKEFILE_LIST)) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}'