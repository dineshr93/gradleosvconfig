
ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))
TEST_DIR := ~/asterix2/AddressBookSync
CONFIG_FILE := ~/asterix2/gradleosvconfig/gradleconfig.txt
SOURCING_SCRIPT_FILE := ~/devenvs/asterix2-int-devenv-1.7.0/main/env_setup.sh

EXE_NAME := goc
BIN := bin
exe :=${BIN}/${EXE_NAME}

ifeq ($(OS),Windows_NT)
	SHELL := powershell.exe
	.SHELLFLAGS := -NoProfile -Command
	RM_F_CMD = Remove-Item -erroraction silentlycontinue -Force
    RM_RF_CMD = ${RM_F_CMD} -Recurse
	exe =${BIN}/${EXE_NAME}.exe
	HELP_CMD = Select-String "^[a-zA-Z_-]+:.*?\#\# .*$$" "./Makefile" | Foreach-Object { $$_data = $$_.matches -split ":.*?\#\# "; $$obj = New-Object PSCustomObject; Add-Member -InputObject $$obj -NotePropertyName ('Command') -NotePropertyValue $$_data[0]; Add-Member -InputObject $$obj -NotePropertyName ('Description') -NotePropertyValue $$_data[1]; $$obj } | Format-Table -HideTableHeaders @{Expression={ $$e = [char]27; "$$e[36m$$($$_.Command)$${e}[0m" }}, Description
else
	SHELL := bash
	RM_F_CMD = rm -f
	RM_RF_CMD = ${RM_F_CMD} -r
	exe =${BIN}/${EXE_NAME}
	HELP_CMD = grep -E '^[a-zA-Z_-]+:.*?\#\# .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?\#\# "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
endif

all: clean build test
.DEFAULT_GOAL := help
.PHONY: clean build test all git help

all: $(PROJ_DIR) ## performs clean build and test
clean: $@ ## Move back the files to original state in test folder
build: $@ ## Generate the windows and linux builds for sep
test: $@ ## Separates the test folder
git: $@ ## commits and push the changes if commit msg m is given without spaces ex m=added_files

prep:
	cd ${TEST_DIR}
	git clean -fd /home/dinesh/asterix2/AddressBookSync
	git restore /home/dinesh/asterix2/AddressBookSync/build.gradle

build:
	echo "Compiling for every OS and Platform"
	set GOOS=windows
	set GOARCH=arm64
	go build -o ${BIN}/${EXE_NAME}.exe main.go
	set GOOS=linux
	set GOARCH=arm64
	go build -o ${BIN}/${EXE_NAME} main.go


test: clean build
	echo "===========Testing==============="
	${exe} ${PROJ_DIR} ${CONFIG_FILE} ${TEST_DIR} ${SOURCING_SCRIPT_FILE}

del:
	${RM_RF_CMD} bin/*


clean: del


git:
	git status
	git add .
	git status
	git commit -m ${m}
	git push

help: ## Show this help
	@${HELP_CMD}