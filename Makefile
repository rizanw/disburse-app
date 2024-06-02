.PHONY: build
build:
	go build -v -o bin/disburse-app cmd/*.go

.PHONY: build
run:
	@echo " >> build disburse-app"
	@make build
	@echo " >> disburse-app built."
	@echo " >> executing disburse-app"
	@./bin/disburse-app
	@echo " >> disburse-app is running"

.PHONY: testf
testf:
	@make test | grep "FAIL" || echo "ALL tests passed"