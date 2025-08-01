
# Change Log
All notable changes to this project will be documented in this file.
 
The format is based on [Keep a Changelog](http://keepachangelog.com/)
and this project adheres to [Semantic Versioning](http://semver.org/).

# Unreleased

### ⚠ BREAKING CHANGES

### Added

### Changed

### Fixed
- [#113](https://github.com/cludden/protoc-gen-go-temporal/pull/113) fix Ensures workflowIdConflictPolicy parameter is properly set when not UNSPECIFIED



# [1.17.1](https://github.com/cludden/protoc-gen-go-temporal/releases/tag/v1.17.1) - 2025-07-23

### Fixed
- [#112](https://github.com/cludden/protoc-gen-go-temporal/pull/112) fix nexus generation for multiple services



# [1.17.0](https://github.com/cludden/protoc-gen-go-temporal/releases/tag/v1.17.0) - 2025-07-23

### ⚠ BREAKING CHANGES
- Signal with start cross-namespace (XNS) methods now utilize dedicated option types instead of using the workflow options. This allows for more flexibility in configuring the signal with start behavior, but may require changes to existing code that uses these methods.

### Added
- [#103](https://github.com/cludden/protoc-gen-go-temporal/pull/103) add support for experimental update with start
- [#105](https://github.com/cludden/protoc-gen-go-temporal/pull/105) add WithDataConverter helper to activity, child workflow options
- [#106](https://github.com/cludden/protoc-gen-go-temporal/pull/106) allow update validation to be configured by each workflow
- [#109](https://github.com/cludden/protoc-gen-go-temporal/pull/109) add support for urfave/cli/v3
- [#111](https://github.com/cludden/protoc-gen-go-temporal/pull/111) add experimental nexus support

### Changed
- [#108](https://github.com/cludden/protoc-gen-go-temporal/pull/108) allow overriding default task queue at runtime

### Fixed
- [#103](https://github.com/cludden/protoc-gen-go-temporal/pull/103) support parsing files for signals, updates
- [#103](https://github.com/cludden/protoc-gen-go-temporal/pull/103) fix input collision handling for signals, updates
- [#103](https://github.com/cludden/protoc-gen-go-temporal/pull/103) fix race condition in xns activity cancellation
- [#108](https://github.com/cludden/protoc-gen-go-temporal/pull/108) fix default cli task queue for workflows with custom task queue



# [1.16.2](https://github.com/cludden/protoc-gen-go-temporal/releases/tag/v1.16.2) - 2025-06-03

### Changed
- [#100](https://github.com/cludden/protoc-gen-go-temporal/pull/100) remove go version from generated files (@szykes)



# [1.16.1](https://github.com/cludden/protoc-gen-go-temporal/releases/tag/v1.16.1) - 2025-03-13

### Changed
- [#99](https://github.com/cludden/protoc-gen-go-temporal/pull/99) revert minimum required go version to 1.21.5



# [1.16.0](https://github.com/cludden/protoc-gen-go-temporal/releases/tag/v1.16.0) - 2025-03-11

### ⚠ BREAKING CHANGES
- The existing usage of the `iancoleman/strcase` third-party library has been
  replaced with a first-party implementation with better acronym handling. This 
  may lead to changes in generated cli command and flag names. To compensate for
  any potential issues, the protobuf options for workflows, queries, signals, 
  updates, and message fields have been expanded to support user-defined 
  overrides.

### Added
- [#89](https://github.com/cludden/protoc-gen-go-temporal/pull/89) add option to ignore camel case renames for specific acronyms
- [#94](https://github.com/cludden/protoc-gen-go-temporal/pull/94) improve acronym behavior, expand cli options
- [#95](https://github.com/cludden/protoc-gen-go-temporal/pull/95) add support for opaque, hybrid apis

### Changed
- [#97](https://github.com/cludden/protoc-gen-go-temporal/pull/97) prevent unnecessary code generation for activity only services

### Fixed
- [#95](https://github.com/cludden/protoc-gen-go-temporal/pull/95) fix unmarshal cli flag prefixing



# [1.15.0](https://github.com/cludden/protoc-gen-go-temporal/releases/tag/v1.15.0) - 2025-01-28

### ⚠ BREAKING CHANGES

### Added
- [#87] add support for [Editions](https://protobuf.dev/programming-guides/editions/) and [Proto3 optional](https://protobuf.dev/programming-guides/editions/#field-labels)

### Changed

### Fixed



# [1.14.5](https://github.com/cludden/protoc-gen-go-temporal/releases/tag/v1.14.5) - 2025-01-06

### Fixed

- [#85](https://github.com/cludden/protoc-gen-go-temporal/pull/85) ensure updateID is passed through test client (@nishkrishnan)



# [1.14.4](https://github.com/cludden/protoc-gen-go-temporal/releases/tag/v1.14.4) - 2024-12-19

### Fixed

- [#80](https://github.com/cludden/protoc-gen-go-temporal/pull/80) fix cross-package signal-with-start
- [#84](https://github.com/cludden/protoc-gen-go-temporal/pull/84) fix cross-package cli imports, repeated string flags 



# [1.14.3](https://github.com/cludden/protoc-gen-go-temporal/releases/tag/v1.14.3) - 2024-08-09

### Fixed
- [#76](https://github.com/cludden/protoc-gen-go-temporal/pull/76) add schedule example



# [1.14.2](https://github.com/cludden/protoc-gen-go-temporal/releases/tag/v1.14.2) - 2024-07-01

### Fixed

- [#74](https://github.com/cludden/protoc-gen-go-temporal/pull/74) fix support for external messages in cli generation



# [1.14.1](https://github.com/cludden/protoc-gen-go-temporal/releases/tag/v1.14.1) - 2024-06-20

### Fixed

- [#73](https://github.com/cludden/protoc-gen-go-temporal/pull/73) default to WorkflowUpdateStageCompleted if update options WaitForStage unspecified



# [1.14.0](https://github.com/cludden/protoc-gen-go-temporal/releases/tag/v1.14.0) - 2024-06-20

### ⚠ BREAKING CHANGES

- [#72](https://github.com/cludden/protoc-gen-go-temporal/pull/72) upgrade go.temporal.io/sdk to [v1.27.0](https://github.com/temporalio/sdk-go/releases/tag/v1.27.0)



# [1.13.3](https://github.com/cludden/protoc-gen-go-temporal/releases/tag/v1.13.3) - 2024-06-13
 
### Fixed

- [#71](https://github.com/cludden/protoc-gen-go-temporal/pull/71) fix activity non_retryable_error_types



# [1.13.2](https://github.com/cludden/protoc-gen-go-temporal/releases/tag/v1.13.2) - 2024-05-31
 
### Fixed

- [#69](https://github.com/cludden/protoc-gen-go-temporal/pull/69) support external messages as rpc parameters



# [1.13.1](https://github.com/cludden/protoc-gen-go-temporal/releases/tag/v1.13.1) - 2024-05-14
 
### Fixed

- [#68](https://github.com/cludden/protoc-gen-go-temporal/pull/68) prevent xns cancellation propagation on worker close



# [1.13.0](https://github.com/cludden/protoc-gen-go-temporal/releases/tag/v1.13.0) - 2024-05-03

## Added

- [#62](https://github.com/cludden/protoc-gen-go-temporal/pull/62) add individual option override methods
 
## Fixed

- [#65](https://github.com/cludden/protoc-gen-go-temporal/pull/65) wrap expression evaluation in local activities inside workflow contexts ([Patch Version 64](https://cludden.github.io/protoc-gen-go-temporal/docs/guides/patches#pv_64-expression-evaluation-local-activity))
- [#66](https://github.com/cludden/protoc-gen-go-temporal/pull/66) fix cancellation propagation in xns activities
 


# [1.12.0](https://github.com/cludden/protoc-gen-go-temporal/releases/tag/v1.12.0) - 2024-04-19
 
## Added

- [0182d7b](https://github.com/cludden/protoc-gen-go-temporal/commit/0182d7bec153fb71636592bbf3a266937fe8bc97) add generated WorkflowFunction helpers
- [#57](https://github.com/cludden/protoc-gen-go-temporal/pull/57) add missing WaitForCancellation for activity options
 
## Changed
  
- [#60](https://github.com/cludden/protoc-gen-go-temporal/pull/60) add additional details to expression evaluation errors
 
## Fixed
 
- [84342c6](https://github.com/cludden/protoc-gen-go-temporal/commit/84342c6e9d6907bf080666572b100561964a4715) support brackets in bloblang expressions
