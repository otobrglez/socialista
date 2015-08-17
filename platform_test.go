package socialista

import (
	"testing"
)

func TestTwitter(t *testing.T) {
	GetStatsForPlatform("https://www.kickstarter.com/projects/elanlee/exploding-kittens", "twitter")
}

func TestFacebook(t *testing.T) {
	GetStatsForPlatform("https://www.kickstarter.com/projects/elanlee/exploding-kittens", "facebook")
}

func TestLinkedin(t *testing.T) {
	GetStatsForPlatform("https://www.kickstarter.com/projects/elanlee/exploding-kittens", "linkedin")
}

func TestPintarest(t *testing.T) {
	GetStatsForPlatform("https://www.kickstarter.com/projects/elanlee/exploding-kittens", "pintarest")
}