// SPDX-FileCopyrightText: 2021 Sebastian Crane <seabass-labrax@gmx.com>
// SPDX-License-Identifier: AGPL-3.0-or-later

package shuffle_string

type StringFragment struct {
	value   string
	shuffle bool
}

func GetDelimeters() []rune {
	return []rune{
		'\u0020', // space
		'\u0021', // exclamation mark
		'\u0023', // hash symbol
		'\u0025', // percent sign
		'\u0026', // ampersand
		'\u0028', // left bracket (parentheses)
		'\u0029', // right bracket (parentheses)
		'\u002a', // asterisk
		'\u002c', // comma
		'\u002f', // forward slash
		'\u003a', // colon
		'\u003b', // semicolon
		'\u003f', // question mark
		'\u005c', // back slash
		'\u2010', // hyphen
		'\u2018', // left single quote
		'\u2019', // right single quote
		'\u201c', // left double quote
		'\u201d', // right double quote
		'\u2026', // ellipsis
		'\u002e', // full stop
	}
}

func IsDelimeter(character rune) bool {
	delimeters := GetDelimeters()
	for i := 0; i < len(delimeters); i++ {
		if character == delimeters[i] {
			return true
		}
	}
	return false
}

func SplitString(s string) []StringFragment {
	var fragments []StringFragment
	currentFragment := 0

	runes := []rune(s)

	// create first fragment
	fragments = append(fragments, StringFragment{
		value:   string(runes[0]),
		shuffle: !IsDelimeter(runes[0]),
	})

	for i := 1; i < len(runes); i++ {

		if IsDelimeter(runes[i]) != fragments[currentFragment].shuffle {
			// character should be added to current fragment
			fragments[currentFragment].value = fragments[currentFragment].value + string(runes[i])
		} else {
			// character should be added to a new fragment
			fragments = append(fragments, StringFragment{value: string(runes[i]), shuffle: !IsDelimeter(runes[i])})
			currentFragment = currentFragment + 1
		}
	}

	return fragments
}
