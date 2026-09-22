// Copyright IBM Corp. 2014, 2026
// SPDX-License-Identifier: MPL-2.0

package version

import "testing"

func TestNewStrictSemverValid(t *testing.T) {
	valid := []string{
		"0.0.4",
		"1.2.3",
		"10.20.30",
		"1.1.2-prerelease+meta",
		"1.0.0-alpha.beta.1",
		"2.0.0+build.1848",
		"1.2.3----RC-SNAPSHOT.12.9.1--.12+788",
	}

	for _, version := range valid {
		t.Run(version, func(t *testing.T) {
			if _, err := NewStrictSemver(version); err != nil {
				t.Fatalf("NewStrictSemver(%q) returned an error: %v", version, err)
			}
		})
	}
}

func TestNewStrictSemverInvalid(t *testing.T) {
	invalid := []string{
		"1.2.-3",
		"1.2.03",
		"1.2.3.4.5.6",
		"1.2.3-preview.01",
		"1.2.3.alpha",
		"1.2.3-preview!",
		"1.2.3-preview..1",
		"1.2.3+one+two",
		"1.2.3+one.2!",
		"v 1.2.3",
		"1.2.3-",
		"1.2.3+",
		"1.2.3.",
		"1..2.3",
		"1",
		"1.2",
		"v1.2.3",
	}

	for _, version := range invalid {
		t.Run(version, func(t *testing.T) {
			if _, err := NewStrictSemver(version); err == nil {
				t.Fatalf("NewStrictSemver(%q) accepted an invalid SemVer", version)
			}
		})
	}
}

func TestNewStrictSemverPreservesLegacyParsing(t *testing.T) {
	legacyOnly := []string{
		"1.2.03",
		"1.2.3.4.5.6",
		"1.2.3-preview.01",
	}

	for _, version := range legacyOnly {
		t.Run(version, func(t *testing.T) {
			if _, err := NewVersion(version); err != nil {
				t.Fatalf("NewVersion(%q) changed: %v", version, err)
			}
			if _, err := NewStrictSemver(version); err == nil {
				t.Fatalf("NewStrictSemver(%q) accepted a legacy-only form", version)
			}
		})
	}
}
