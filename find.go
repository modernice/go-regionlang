package regionlang

import (
	"golang.org/x/text/language"
)

// AllTags returns the built-in language tags from the
// golang.org/x/text/language package.
func AllTags() []language.Tag {
	return allTags[:]
}

// Find returns the base language for the given region (country).
//
// If `tags` are provided, only the provided tags are considered. Otherwise all
// built-in languages are considered (see AllTags). When a region matches
// multiple of the provided language tags, the last matched tag is returned.
//
// Example:
//	base, confidence := regionlang.Find("at") // Find language for Austria
//	base.String() == "de" // German is the Austrian language
//	conf == language.Exact
//
// Find searches through the allowed tags (f.tags) and compares the regions of
// the tags with the provided region. If the region of a language tag matches
// with the provided region, the base language of the tag is returned. In cases
// where the provided region matches multiple language tags, the last match with
// the highest confidence is returned.
//
// If Find fails to parse the region, language.Make(region).Base() is returned.
func Find(region string, tags ...language.Tag) (language.Base, language.Confidence) {
	return NewFinder(tags...).Find(region)
}

// Finder finds base languages for regions (countries).
// Use NewFinder to create a Finder.
type Finder struct {
	tags []language.Tag
}

// NewFinder returns a Finder for the given language tags.
func NewFinder(tags ...language.Tag) *Finder {
	if len(tags) == 0 {
		tags = allTags[:]
	}
	return &Finder{tags: tags}
}

// Find returns the base language for the given region (country).
//
// Example:
//	base, confidence := f.Find("at") // Find language for Austria
//	base.String() == "de" // German is the Austrian language
//	conf == language.Exact
//
// Find searches through the allowed tags (f.tags) and compares the regions of
// the tags with the provided region. If the region of a language tag matches
// with the provided region, the base language of the tag is returned. In cases
// where the provided region matches multiple language tags, the last match with
// the highest confidence is returned.
//
// If Find fails to parse the region, language.Make(region).Base() is returned.
func (f *Finder) Find(region string) (language.Base, language.Confidence) {
	reg, err := language.ParseRegion(region)
	if err != nil {
		return language.Make(region).Base()
	}

	var bestMatch language.Base
	var bestConfidence language.Confidence

	for _, tag := range f.tags {
		tagRegion, conf := tag.Region()
		if conf == language.No {
			continue
		}

		if reg == tagRegion {
			base, conf := tag.Base()
			if conf >= bestConfidence {
				bestMatch = base
				bestConfidence = conf
			}
		}
	}

	return bestMatch, bestConfidence
}

// All language tags. Keep this updated.
var allTags = [...]language.Tag{
	language.Afrikaans,
	language.Amharic,
	language.Arabic,
	language.ModernStandardArabic,
	language.Azerbaijani,
	language.Bulgarian,
	language.Bengali,
	language.Catalan,
	language.Czech,
	language.Danish,
	language.German,
	language.Greek,
	language.English,
	language.AmericanEnglish,
	language.BritishEnglish,
	language.Spanish,
	language.EuropeanSpanish,
	language.LatinAmericanSpanish,
	language.Estonian,
	language.Persian,
	language.Finnish,
	language.Filipino,
	language.French,
	language.CanadianFrench,
	language.Gujarati,
	language.Hebrew,
	language.Hindi,
	language.Croatian,
	language.Hungarian,
	language.Armenian,
	language.Indonesian,
	language.Icelandic,
	language.Italian,
	language.Japanese,
	language.Georgian,
	language.Kazakh,
	language.Khmer,
	language.Kannada,
	language.Korean,
	language.Kirghiz,
	language.Lao,
	language.Lithuanian,
	language.Latvian,
	language.Macedonian,
	language.Malayalam,
	language.Mongolian,
	language.Marathi,
	language.Malay,
	language.Burmese,
	language.Nepali,
	language.Dutch,
	language.Norwegian,
	language.Punjabi,
	language.Polish,
	language.Portuguese,
	language.BrazilianPortuguese,
	language.EuropeanPortuguese,
	language.Romanian,
	language.Russian,
	language.Sinhala,
	language.Slovak,
	language.Slovenian,
	language.Albanian,
	language.Serbian,
	language.SerbianLatin,
	language.Swedish,
	language.Swahili,
	language.Tamil,
	language.Telugu,
	language.Thai,
	language.Turkish,
	language.Ukrainian,
	language.Urdu,
	language.Uzbek,
	language.Vietnamese,
	language.Chinese,
	language.SimplifiedChinese,
	language.TraditionalChinese,
	language.Zulu,
}
