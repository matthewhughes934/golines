# Changelog

## 0.12.0 - 2023-07-30

Now maintained at <https://github.com/matthewhughes934/golines>

### Fixed

  - Pass filenames to pre-commit hook

## 0.11.0 - 2022-08-06

### Fixed

  - Fix tag alignment with multi-byte characters (\#78)

### Changed

  - Skip writing files with unchanged content (\#77)
  - Regenerate generated graph code (\#69)

## 0.10.0 - 2022-05-29

### Added

  - Add pre-commit hook config (\#64)

### Changed

  - Update `github.com/dave/dst` to `v0.27.0` (\#68)
  - **Breaking** Require Go `1.18+` (\#68)

## 0.9.0 - 2022-03-19

### Changed

  - Upgrade `x/sys` for go1.18 (\#58)

### Fixed

  - Fix unable to install under Windows (\#55)

## 0.8.0 - 2022-03-04

### Changed

  - Improve comment reflow (\#47)
  - Remove Python dependency (\#51)

### Fixed

  - Don't add colors unless output goes to a terminal (\#50)

## 0.7.0 - 2021-12-17

### Fixed

  - Fix shortening of inline interface definitions inside function declarations
    (\#46)

### Added

  - Add goland instructions to readme (\#43)

## 0.6.0 - 2021-11-17

### Changed

  - Update to Go 1.17

## 0.5.0 - 2021-06-26

### Fixed

  - fix formatting for single var declaration (\#35)

## 0.4.0 - 2021-04-23

### Fixed

  - Fix formatting for very long binary expressions (\#31)

## 0.3.0 - 2021-04-11

### Changed

  - Improve handling of chained methods (\#28)

### Added

  - Add `--chain-split-dots` argument

## 0.2.0 - 2021-03-30

### Changed

  - Improve generated files detection heuristic (\#26)

## 0.1.0 - 2021-02-22

Initial release
