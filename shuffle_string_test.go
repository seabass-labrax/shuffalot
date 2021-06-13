// SPDX-FileCopyrightText: 2021 Sebastian Crane <seabass-labrax@gmx.com>
// SPDX-License-Identifier: AGPL-3.0-or-later

package shuffle_string

import ("testing")

type TestShuffleWordValues struct {
	word string
	seed int64
	expected string
}

func TestShuffleWord (t *testing.T) {

	var testCases [3]TestShuffleWordValues

	testCases [0] = TestShuffleWordValues {word: "international", seed: 10, expected: "itnenanariotl"}
	testCases [1] = TestShuffleWordValues {word: "circumnavigation", seed: 37, expected: "ccmatoivgrniauin"}
	testCases [2] = TestShuffleWordValues {word: "communication", seed: 94, expected: "cniocimutoamn"}

	for _, testCase := range testCases {

		actual := ShuffleWord (testCase.word, testCase.seed)

		if actual != testCase.expected {
			t.Errorf ("Test failed: expected '%s', got '%s'", testCase.expected, actual)
		}
	}
}
