APP     := justlive
DIST    := dist
LDFLAGS := -s -w
TARGETS := linux-amd64 linux-arm64 windows-amd64 windows-arm64 darwin-amd64 darwin-arm64

# Build helpers for "os-arch" targets
os    = $(word 1,$(subst -, ,$1))
arch  = $(word 2,$(subst -, ,$1))
ext   = $(if $(findstring windows,$1),.exe)
winld = $(if $(findstring windows,$1),-H windowsgui)

.PHONY: build run vet dist clean $(TARGETS)

# Production build for this machine
build:
	CGO_ENABLED=0 go build -trimpath -ldflags "$(LDFLAGS)" -o $(DIST)/$(APP) .

run:
	go run .

vet:
	CGO_ENABLED=0 go vet ./...

# Production builds for every OS / arch
dist: $(TARGETS)

$(TARGETS):
	CGO_ENABLED=0 GOOS=$(call os,$@) GOARCH=$(call arch,$@) \
		go build -trimpath -ldflags "$(LDFLAGS) $(call winld,$@)" \
		-o $(DIST)/$(APP)-$@$(call ext,$@) .

clean:
	rm -rf $(DIST)
