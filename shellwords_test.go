package shellwords

import (
    "testing"
)

func TestNext1(t *testing.T) {
    var tok string
    var hasTok bool

    sp := New("this is a test")

    tok, hasTok = sp.Next()
    if !((tok == "this") && (hasTok)) { t.Fail() }
    tok, hasTok = sp.Next()
    if !((tok == "is") && (hasTok)) { t.Fail() }
    tok, hasTok = sp.Next()
    if !((tok == "a") && (hasTok)) { t.Fail() }
    tok, hasTok = sp.Next()
    if !((tok == "test") && (hasTok)) { t.Fail() }
    tok, hasTok = sp.Next()
    if !((tok == "") && (! hasTok)) { t.Fail() }
}

func TestNext2(t *testing.T) {
    var tok string
    var hasTok bool

    sp := New("   this   is  another test")

    tok, hasTok = sp.Next()
    if !((tok == "this") && (hasTok)) { t.Fail() }
    tok, hasTok = sp.Next()
    if !((tok == "is") && (hasTok)) { t.Fail() }
    tok, hasTok = sp.Next()
    if !((tok == "another") && (hasTok)) { t.Fail() }
    tok, hasTok = sp.Next()
    if !((tok == "test") && (hasTok)) { t.Fail() }
    tok, hasTok = sp.Next()
    if !((tok == "") && (! hasTok)) { t.Fail() }
}

func TestNext3(t *testing.T) {
    var tok string
    var hasTok bool

    sp := New("\tthis   is     the  third test           ")

    tok, hasTok = sp.Next()
    if !((tok == "this") && (hasTok)) { t.Fail() }
    tok, hasTok = sp.Next()
    if !((tok == "is") && (hasTok)) { t.Fail() }
    tok, hasTok = sp.Next()
    if !((tok == "the") && (hasTok)) { t.Fail() }
    tok, hasTok = sp.Next()
    if !((tok == "third") && (hasTok)) { t.Fail() }
    tok, hasTok = sp.Next()
    if !((tok == "test") && (hasTok)) { t.Fail() }
    tok, hasTok = sp.Next()
    if !((tok == "") && (! hasTok)) { t.Fail() }
}

func TestNext4(t *testing.T) {
    var tok string
    var hasTok bool

    sp := New("  \t    this   is  \n   the  \t  forth  \r\n   test          \n")

    tok, hasTok = sp.Next()
    if !((tok == "this") && (hasTok)) { t.Fail() }
    tok, hasTok = sp.Next()
    if !((tok == "is") && (hasTok)) { t.Fail() }
    tok, hasTok = sp.Next()
    if !((tok == "the") && (hasTok)) { t.Fail() }
    tok, hasTok = sp.Next()
    if !((tok == "forth") && (hasTok)) { t.Fail() }
    tok, hasTok = sp.Next()
    if !((tok == "test") && (hasTok)) { t.Fail() }
    tok, hasTok = sp.Next()
    if !((tok == "") && (! hasTok)) { t.Fail() }
}

func TestSplit1(t *testing.T) {
    sp := New("this is a test").Split()

    if !((len(sp) == 4) &&
        (sp[0] == "this") &&
        (sp[1] == "is") &&
        (sp[2] == "a") &&
        (sp[3] == "test")) {
        t.Fail()
    }
}

func TestSplit2(t *testing.T) {
    s := New("\t\t\tthis\t  is\n another\t\t test \n\t")

    tok, hasTok := s.Next()
    if !((tok == "this") && (hasTok)) { t.Fail() }

    sp := s.Split()

    if !((len(sp) == 3) &&
        (sp[0] == "is") &&
        (sp[1] == "another") &&
        (sp[2] == "test")) {
        t.Fail()
    }
}

func TestSplit3(t *testing.T) {
    sp := New("\t\t\t \n \t\t \n\t").Split()

    if !((len(sp) == 0)) {
        t.Fail()
    }
}

func TestStrings1(t *testing.T) {
    sp := New("this 'has got' some \"string values\"").Split()

    if !((len(sp) == 4) &&
        (sp[0] == "this") &&
        (sp[1] == "has got") &&
        (sp[2] == "some") &&
        (sp[3] == "string values")) {
        t.Fail()
    }
}

func TestStrings2(t *testing.T) {
    sp := New("strings '  can start and end  ' with \"\twhitespace\ncharacters\"").Split()

    if !((len(sp) == 4) &&
        (sp[0] == "strings") &&
        (sp[1] == "  can start and end  ") &&
        (sp[2] == "with") &&
        (sp[3] == "\twhitespace\ncharacters")) {
        t.Fail()
    }
}

func TestStrings3(t *testing.T) {
    sp := New("nested '\"string characters\"' are \"'supported'\"").Split()

    if !((len(sp) == 4) &&
        (sp[0] == "nested") &&
        (sp[1] == "\"string characters\"") &&
        (sp[2] == "are") &&
        (sp[3] == "'supported'")) {
        t.Fail()
    }
}

func TestStrings4(t *testing.T) {
    sp := New("empty '' tokens \"\"").Split()

    if !((len(sp) == 4) &&
        (sp[0] == "empty") &&
        (sp[1] == "") &&
        (sp[2] == "tokens") &&
        (sp[3] == "")) {
        t.Fail()
    }
}
