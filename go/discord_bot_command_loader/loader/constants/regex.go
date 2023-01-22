package constants

import (
	"regexp"
)

var (
	DataRegex     = regexp.MustCompile(`var (\w+)Data`)
	RunRegex      = regexp.MustCompile(`func (\w+)Run\(`)
	PackageRegex  = regexp.MustCompile(`(?s)package (\w+)`)
	FileNameRegex = regexp.MustCompile(`(?i)(\w+)\.go`)
	DefaultImport = "github.com/vyrekxd/bot/core"
)
