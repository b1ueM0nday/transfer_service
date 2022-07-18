MOCKS_DESTINATION=internal/mocks
.PHONY: mocks
# put the files with interfaces you'd like to mock in prerequisites
# wildcards are allowed
mocks: internal/client/api.go internal/client/client.go internal/client/logs/log.go internal/repository/repo.go internal/gg/gg.go internal/gg/gg_flow.go
	@echo "Generating mocks..."
	@rm -rf $(MOCKS_DESTINATION)
	@for file in $^; do mockgen -source=$$file -destination=$(MOCKS_DESTINATION)/$$file; done