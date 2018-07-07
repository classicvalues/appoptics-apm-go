// Copyright (C) 2017 Librato, Inc. All rights reserved.

package agent

import (
	"os"
	"regexp"
	"strings"
	"unicode/utf8"
)

type configBuilder struct {
	name         ConfName
	defaultValue string
	builders     []initFunc
	// TODO: validation func
}

type conf struct {
	initialized bool
	items       map[ConfName]string
}

type initFunc func(n ConfName) string

// Environment variable reader. Empty string is considered as invalid so just use os.Getenv()
// to ignore empty environment variables
var envVar initFunc = func(n ConfName) string {
	return strings.ToLower(os.Getenv(string(n)))
}

// A valid service key is something like 'service_token:service_name'.
// The service_token should be of 64 characters long and the size of
// service_name is larger than 0 but up to 255 characters.
var IsValidServiceKey = regexp.MustCompile(`^[a-zA-Z0-9]{64}:\S{1,255}$`).MatchString

// Default values
const (
	defaultSSLCollector       = "collector.appoptics.com:443"
	defaultServiceKey         = ""
	defaultLogLevel           = "WARN"
	defaultTrustedPath        = ""
	defaultCollectorUDP       = "127.0.0.1:7831"
	defaultReporter           = "ssl"
	defaultTracingMode        = "always"
	defaultPrependDomain      = "false"
	defaultHostnameAlias      = ""
	defaultInsecureSkipVerify = "false"
	defaultHistogramPrecision = ""
)

var cb = []configBuilder{
	{AppOpticsCollector, defaultSSLCollector, []initFunc{envVar}},
	{AppOpticsServiceKey, defaultServiceKey, []initFunc{envVar}},
	{AppOpticsLogLevel, defaultLogLevel, []initFunc{envVar}},
	{AppOpticsTrustedPath, defaultTrustedPath, []initFunc{envVar}},
	{AppOpticsCollectorUDP, defaultCollectorUDP, []initFunc{envVar}},
	{AppOpticsReporter, defaultReporter, []initFunc{envVar}},
	{AppOpticsTracingMode, defaultTracingMode, []initFunc{envVar}},
	{AppOpticsPrependDomain, defaultPrependDomain, []initFunc{envVar}},
	{AppOpticsHostnameAlias, defaultHostnameAlias, []initFunc{envVar}},
	{AppOpticsInsecureSkipVerify, defaultInsecureSkipVerify, []initFunc{envVar}},
	{AppOpticsHistogramPrecision, defaultHistogramPrecision, []initFunc{envVar}},
}

// The package variable to store all configurations, which is read only after initialized.
var agentConf = &conf{
	initialized: false,
	items:       make(map[ConfName]string),
}

func initConf(cf *conf) {
	Info("Initializing the AppOptics agent.")
	for _, item := range cb {
		k := item.name
		v := ""
		l := len(item.builders) - 1
		for i := l; i >= 0; i-- {
			v = item.builders[i](k)
			if v != "" {
				val := v
				if k == "APPOPTICS_SERVICE_KEY" {
					if IsValidServiceKey(val) {
						val = maskServiceKey(val)
					} else {
						val = v
					}
				}
				Warningf("non-default configuration used %v=%v", k, val)

				break
			}
		}
		if v == "" {
			v = item.defaultValue
		}
		cf.items[k] = v
	}
	cf.initialized = true
}

// maskServiceKey masks the middle part of the token and returns the masked service key
// For example, the key "ae38315f6116585d64d82ec2455aa3ec61e02fee25d286f74ace9e4fea189217:go"
// will be masked to "ae38********************************************************9217:go"
func maskServiceKey(validKey string) string {
	var sep = ":"
	var hLen, tLen = 4, 4
	var mask = "*"

	s := strings.Split(validKey, sep)
	tk := s[0]

	if len(tk) <= hLen+tLen {
		return validKey
	}

	tk = tk[0:4] + strings.Repeat(mask, utf8.RuneCountInString(tk)-hLen-tLen) + tk[len(tk)-4:]

	return tk + sep + s[1]
}
