# LM Tech Test

## Specification

### Preface

Do not use any extended functionality of the framework to solve this problem. (For example, don't use the string methods
that have functionality such as IndexOf, substring, regular expression classes, LINQ etc). The test should take no
longer than an hour for you to complete, please submit the working source code to solve the problem along with any
supporting code that you might have used in testing.

### Problem

We need a way of finding all the occurrences of a particular set of characters in a string. It should

* Accept two strings as input. One called 'textToSearch' and one called 'subtext'
* The solution should match the subtext against the textToSearch, outputting the positions of the beginning of each
  match for the subtext within the textToSearch.
* Multiple matches are possible
* Matching is case-insensitive
* If no matches have been found, no output is generated

### Sample Acceptance Criteria

**Text:**

`textToSearch = "Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!"`

**Expected Result:**

| Subtext    | Result               |
| -----------|:---------------------|
| Peter      | [1, 20, 75]          |
| peter      | [1, 20, 75]          |
| pick       | [30, 58]             |
| pi         | [30, 37, 43, 51, 58] |
| z          | No Output          |
| Peterz     | No Output          |

### Notes on the solution

The specs say:

> Do not use any extended functionality of the framework to solve this problem

But when it came to making things case-insensitive, that became tricky due to Unicode. I don't know the logic for
converting an uppercase Unicode character to a lowercase off the top of my head. So I would just be replicating the
logic from [here](https://go.googlesource.com/go/+/go1.16.3/src/unicode/letter.go#207). Which didn't seem like the right
thing to do, so I used the [unicode.ToLower()](https://pkg.go.dev/unicode#ToLower) function. I hope that is OK.

Also, the specs were a bit unclear about what to do when it came to empty strings. So I just made my own decision and
documented it via tests.

I also added some CI/CD to run the tests against multiple versions of Go and different operating systems. You can 
access it by clicking on this badge -> [![Continuous Integration](https://github.com/ebh/lm-tech-test/actions/workflows/ci.yml/badge.svg)](https://github.com/ebh/lm-tech-test/actions/workflows/ci.yml)
