// SPDX-FileCopyrightText: 2021 Sebastian Crane <seabass-labrax@gmx.com>
// SPDX-License-Identifier: AGPL-3.0-or-later

package shuffle

import (
	"math/rand"
	"time"
)

func ShuffleWord(word string, seed ...int64) string {

	if len(seed) > 0 {
		rand.Seed(seed[0])
	} else {
		rand.Seed(time.Now().UnixNano())
	}

	runes := []rune(word)

	// Fisher-Yates shuffle of elements
	// do not shuffle first and last element of slice
	for i := len(runes) - 2; i > 0; i-- {
		randomNumber := rand.Intn(i) + 1
		runes[randomNumber], runes[i] = runes[i], runes[randomNumber]
	}

	return string(runes)
}
