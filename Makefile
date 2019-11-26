GIT_COMMIT := $(shell git rev-parse HEAD)
VERSION := $(shell cat ver)

LDFLAGS := "-X simplesurance/github-cli/internal/command.buildCommit=$(GIT_COMMIT) \
	    -X simplesurance/github-cli/internal/command.buildVersion=$(VERSION)"

TARFLAGS := --sort=name --mtime='2018-01-01 00:00:00' --owner=0 --group=0 --numeric-owner
DOCKER_REPO := simplesurance/github-cli
DOCKER_ARG_TAGS := -t $(DOCKER_REPO):latest -t $(DOCKER_REPO):$(VERSION)

.PHONY: all
all:
	go build -ldflags=$(LDFLAGS) -o github-cli cmd/github-cli/main.go

.PHONY: clean
clean:
	@rm -rf github-cli dist/

.PHONY: dist/darwin_amd64/github-cli
dist/darwin_amd64/github-cli:
	$(info * building $@)
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build \
		-ldflags=$(LDFLAGS) -o "$@" cmd/github-cli/main.go

	$(info * creating $(@D)/github-cli-darwin_amd64-$(VERSION).tar.xz)
	@tar $(TARFLAGS) -C $(@D) -cJf $(@D)/github-cli-darwin_amd64-$(VERSION).tar.xz $(@F)

	$(info * creating $(@D)/github-cli-darwin_amd64-$(VERSION).tar.xz.sha256)
	@(cd $(@D) && sha256sum github-cli-darwin_amd64-$(VERSION).tar.xz > github-cli-darwin_amd64-$(VERSION).tar.xz.sha256)

.PHONY: dist/linux_amd64/github-cli
dist/linux_amd64/github-cli:
	$(info * building $@)
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
		-ldflags=$(LDFLAGS) -o "$@" cmd/github-cli/main.go

	$(info * creating $(@D)/github-cli-linux_amd64-$(VERSION).tar.xz)
	@tar $(TARFLAGS) -C $(@D) -cJf $(@D)/github-cli-linux_amd64-$(VERSION).tar.xz $(@F)

	$(info * creating $(@D)/github-cli-linux_amd64-$(VERSION).tar.xz.sha256)
	@(cd $(@D) && sha256sum github-cli-linux_amd64-$(VERSION).tar.xz > github-cli-linux_amd64-$(VERSION).tar.xz.sha256)

.PHONY: docker_image
docker_image: dist/linux_amd64/github-cli
	@mkdir -p docker/files
	@cp dist/linux_amd64/github-cli docker/files/
	( cd docker && docker build  $(DOCKER_ARG_TAGS) . )

.PHONY: dirty_worktree_check
dirty_worktree_check:
	@if ! git diff-files --quiet || git ls-files --other --directory --exclude-standard | grep ".*" > /dev/null ; then \
		echo "remove untracked files and changed files in repository before creating a release, see 'git status'"; \
		exit 1; \
		fi

.PHONY: release
release: clean dirty_worktree_check dist/linux_amd64/github-cli dist/darwin_amd64/github-cli docker_image
	@echo
	@echo next steps:
	@echo - git tag v$(VERSION)
	@echo - git push --tags
	@echo - upload $(ls dist/*/*.tar.xz) files
	@echo - "push docker image - docker push $(DOCKER_REPO):latest && docker push $(DOCKER_REPO):$(VERSION)"
