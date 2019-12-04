package quoad

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCommitMessage(t *testing.T) {
	tests := map[string]Commit{
		"chore: testing\n":                                                  Commit{Category: "chore", Scope: "", Heading: "testing"},
		"feat(ci): ci test\n":                                               Commit{Category: "feat", Scope: "ci", Heading: "ci test"},
		"feat(ci)!: ci test\n":                                              Commit{Category: "feat", Scope: "ci", Heading: "ci test", Breaking: true},
		"merge master in something\n":                                       Commit{Scope: "", Heading: "merge master in something\n", Breaking: false},
		"chore: test\n\nsomething more here":                                Commit{Category: "chore", Scope: "", Heading: "test", Body: "something more here"},
		"chore: test\n\nsomething more here\nRefs: #12":                     Commit{Category: "chore", Scope: "", Heading: "test", Body: "something more here", Issues: []int{12}},
		"chore: test\n\nsomething more here\n\tRefs: #12":                   Commit{Category: "chore", Scope: "", Heading: "test", Body: "something more here", Issues: []int{12}},
		"chore: test\n\nsomething more here\n\t Refs: #12":                  Commit{Category: "chore", Scope: "", Heading: "test", Body: "something more here", Issues: []int{12}},
		"chore: test\n\nsomething more here\nRefs: #12\nRefs: #13":          Commit{Category: "chore", Scope: "", Heading: "test", Body: "something more here", Issues: []int{12, 13}},
		"chore: test\n\nsomething more here\nRefs: #12, #13":                Commit{Category: "chore", Scope: "", Heading: "test", Body: "something more here", Issues: []int{12, 13}},
		"chore: test\n\nsomething more here\nRefs: #12 and #13":             Commit{Category: "chore", Scope: "", Heading: "test", Body: "something more here", Issues: []int{12, 13}},
		"chore: add something\n":                                            Commit{Category: "chore", Heading: "add something"},
		"chore(ci): added new CI stuff\n":                                   Commit{Category: "chore", Scope: "ci", Heading: "added new CI stuff"},
		"feat: added a new feature\n":                                       Commit{Category: "feat", Heading: "added a new feature"},
		"fix!: breaking change\n":                                           Commit{Category: "fix", Breaking: true, Heading: "breaking change"},
		"fix(security)!: breaking\n":                                        Commit{Category: "fix", Scope: "security", Breaking: true, Heading: "breaking"},
		"fix!!: breaking\n":                                                 Commit{Heading: "fix!!: breaking\n"},
		"fix(security)(stuff): should break\n":                              Commit{Category: "fix", Scope: "security(stuff)", Heading: "should break"},
		"chore:really close\n":                                              Commit{Heading: "chore:really close\n"},
		"perf(): nope\n":                                                    Commit{Heading: "perf(): nope\n"},
		"chore(: bad\n":                                                     Commit{Heading: "chore(: bad\n"},
		": nope\n":                                                          Commit{Heading: "nope"},
		"fix tests\n":                                                       Commit{Heading: "fix tests\n"},
		"test(full): a heading\n\nbody is here\nit can have multiple lines": Commit{Category: "test", Scope: "full", Heading: "a heading", Body: "body is here\nit can have multiple lines"},
	}

	for test, expected := range tests {
		err := ParseCommitMessage(test)
		assert.Equal(t, expected, err)
	}
}
