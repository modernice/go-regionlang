package regionlang_test

import (
	"fmt"
	"testing"

	"github.com/modernice/go-regionlang"
	"golang.org/x/text/language"
)

func TestFind(t *testing.T) {
	tests := []struct {
		tags   []language.Tag
		region string
		want   func() (language.Base, language.Confidence)
	}{
		{
			region: "de",
			want: func() (language.Base, language.Confidence) {
				return language.German.Base()
			},
		},
		{
			region: "de-DE",
			want: func() (language.Base, language.Confidence) {
				return language.German.Base()
			},
		},
		{
			region: "de-CH",
			want: func() (language.Base, language.Confidence) {
				return language.German.Base()
			},
		},
		{
			region: "de-AT",
			want: func() (language.Base, language.Confidence) {
				return language.German.Base()
			},
		},
		{
			region: "fr",
			want: func() (language.Base, language.Confidence) {
				return language.French.Base()
			},
		},
		{
			region: "fr-FR",
			want: func() (language.Base, language.Confidence) {
				return language.French.Base()
			},
		},
		{
			region: "fr-BE",
			want: func() (language.Base, language.Confidence) {
				return language.French.Base()
			},
		},
		{
			region: "fr-CA",
			want: func() (language.Base, language.Confidence) {
				return language.CanadianFrench.Base()
			},
		},
		{
			region: "en",
			want: func() (language.Base, language.Confidence) {
				return language.English.Base()
			},
		},
		{
			region: "en-US",
			want: func() (language.Base, language.Confidence) {
				return language.AmericanEnglish.Base()
			},
		},
		{
			region: "en-CA",
			want: func() (language.Base, language.Confidence) {
				return language.English.Base()
			},
		},
		{
			region: "be",
			tags:   append(regionlang.AllTags(), language.Make("fr-BE")),
			want: func() (language.Base, language.Confidence) {
				return language.French.Base()
			},
		},
		// "en-CA" should override the default "fr-CA" match.
		{
			region: "ca",
			tags:   append(regionlang.AllTags(), language.Make("en-CA")),
			want: func() (language.Base, language.Confidence) {
				return language.English.Base()
			},
		},
		// "en-CA" should be overriden by the default "fr-CA" match.
		{
			region: "ca",
			tags:   append([]language.Tag{language.Make("en-CA")}, regionlang.AllTags()...),
			want: func() (language.Base, language.Confidence) {
				return language.CanadianFrench.Base()
			},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%v (%v)", tt.region, tt.tags), func(t *testing.T) {
			wantBase, wantConf := tt.want()
			base, conf := regionlang.Find(tt.region, tt.tags...)

			if wantConf != conf {
				t.Fatalf("Find(%q) should return confidence %q; got %q", tt.region, wantConf, conf)
			}

			if wantBase != base {
				t.Fatalf("Find(%q) should return language base %q; got %q", tt.region, wantBase, base)
			}

			if conf != language.Exact {
				t.Fatalf("confidence: %v", conf)
			}
		})
	}
}
