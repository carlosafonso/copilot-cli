// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cli

import (
	"errors"
	"regexp"
	"unicode"
)

var (
	errValueEmpty              = errors.New("value must not be empty")
	errValueTooLong            = errors.New("value must not exceed 255 characters")
	errValueNotAlphanumeric    = errors.New("value must be alphanumeric: [A-Za-z0-9]")
	errValueNotAString         = errors.New("value must be a string")
	errValueFirstCharNotLetter = errors.New("value must start with letter")
)

func validateProjectName(val interface{}) error {
	return basicNameValidation(val)
}

func validateApplicationName(val interface{}) error {
	return basicNameValidation(val)
}

func validateEnvironmentName(val interface{}) error {
	return basicNameValidation(val)
}

func basicNameValidation(val interface{}) error {
	s, ok := val.(string)
	if !ok {
		return errValueNotAString
	}
	if s == "" {
		return errValueEmpty
	}
	if len(s) > 255 {
		return errValueTooLong
	}
	if !isAlphanumeric(s) {
		return errValueNotAlphanumeric
	}
	if !startsWithLetter(s) {
		return errValueFirstCharNotLetter
	}

	return nil
}

func isAlphanumeric(s string) bool {
	for _, r := range s {
		if !(unicode.IsLetter(r) || unicode.IsNumber(r)) {
			return false
		}
	}
	return true
}

func startsWithLetter(s string) bool {
	valid, err := regexp.MatchString(`^[a-zA-Z]`, s)
	if err != nil {
		return false // bubble up error?
	}

	return valid
}