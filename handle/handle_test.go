// SPDX-FileCopyrightText: 2021 Sebastian Crane <seabass-labrax@gmx.com>
// SPDX-License-Identifier: AGPL-3.0-or-later

package handle

import (
	"math/rand"
	"testing"
)

type TestHandleLineValues struct {
	line     string
	seed     int64
	expected string
}

func TestHandleLine(t *testing.T) {

	var testCases [3]TestHandleLineValues

	testCases[0] = TestHandleLineValues{
		line:     "“The chances against anything manlike on Mars are a million to one,” he said.",
		seed:     33,
		expected: "“The cehancs ansiagt aiyhnntg mialkne on Mars are a mlilion to one,” he siad.",
	}

	testCases[1] = TestHandleLineValues{
		line:     "When the night wind howls in the chimney cowls, and the bat in the moonlight flies,",
		seed:     19,
		expected: "When the nghit wnid hwlos in the cnmeihy cwlos, and the bat in the mnhooligt felis,",
	}

	testCases[2] = TestHandleLineValues{
		line:     "",
		seed:     12,
		expected: "",
	}

	for _, testCase := range testCases {

		actual := HandleLine(testCase.line, *rand.New(rand.NewSource(testCase.seed)))

		if actual != testCase.expected {
			t.Errorf("Test failed: expected '%s', got '%s'", testCase.expected, actual)
		}
	}
}
