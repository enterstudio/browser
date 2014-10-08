##
# Provides targets to build, test and run the Veyron Browser application.
#
# make  # Builds the project.
# make test  # Runs unit and integration tests.
# make start  # Starts the services and http server needed to run the application at http://localhost:9000
# make clean  # Deleted all build, testing and other artifacts.
#
# Note: :; at the beginning of commands is a work-around for an issue in MacOS version of GNU `make` where
# `make` may not invoke shell to run a command if command is deemed simple enough causing environment variables
# like PATH that are modified here not to be used.
# :; tricks make to assume command is not simple and needs to invoke shell.
# see http://stackoverflow.com/questions/21708839/problems-setting-path-in-makefile for details.
##

PATH:=$(VEYRON_ROOT)/environment/cout/node/bin:$(PATH)
PATH:=node_modules/.bin:$(PATH)
ifndef TMPDIR
	export TMPDIR:=/tmp
endif
TMPDIR:=$(TMPDIR)/veyron_browser

# All JS and CSS files except build.js and third party.
BROWSERIFY_FILES = $(shell find src -name "*.js" -o -name "*.css")
BROWSERIFY_OPTIONS = --transform ./css-transform --debug

# All Go and VDL files.
GO_FILES = $(shell find go -name "*.go")
VDL_FILES = $(shell find go -name "*.vdl")

# Builds everything.
all: directories public/bundle.js public/bundle.html public/platform.js public/platform.js.map public/polymer.js.map

# Creating the bundle JS file.
public/bundle.js: $(BROWSERIFY_FILES) node_modules
	:;jshint src # lint all src JavaScript files.
	:;browserify src/app.js $(BROWSERIFY_OPTIONS) $< | exorcist $@.map > $@ # Browserify and generate map file.

# Creating the bundle HTML file.
public/bundle.html: web-component-dependencies.html node_modules bower_components
	:;vulcanize --output public/bundle.html web-component-dependencies.html --inline

# Copies the web components platform file.
public/platform.js: bower_components
	cp bower_components/platform/platform.js public/platform.js

public/platform.js.map: bower_components
	cp bower_components/platform/platform.js.map public/platform.js.map

public/polymer.js.map: bower_components
	cp bower_components/polymer/polymer.js.map public/polymer.js.map

# Install what we need from NPM.
node_modules: package.json
	:;npm prune
	:;npm install
	touch node_modules

# Install non-JS dependencies from bower.
bower_components: bower.json node_modules
	:;bower prune
	:;bower install
	touch bower_components

# PHONY targets:

# Run unit and integration tests.
test: all
	:;jshint test # lint all test JavaScript files.
	:;./scripts/services/run-tests.sh

# Continuously watch for changes to .js, .html or .css files.
# Rebundles the appropriate bundles when local files change.
watch:
	watch -n 1 make

# Continuously reruns the tests as they change.
watch-test:
	:;PROVA_WATCH=true ./scripts/services/run-tests.sh

# Serves the needed daemons and starts a server at http://localhost:9000
# CTRL-C to stop
start: all
	:;./scripts/services/run-webapp.sh

# Create needed directories like temp.
directories:
	mkdir -p $(TMPDIR)

# Clean all build artifacts.
clean:
	rm -f public/bundle.*
	rm -f public/platform.*
	rm -f public/polymer.*
	rm -rf node_modules
	rm -rf go/bin
	rm -rf bower_components
	rm -rf $(TMPDIR)

.PHONY: start clean watch test watch-test directories
