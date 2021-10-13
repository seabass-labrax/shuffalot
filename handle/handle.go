// SPDX-FileCopyrightText: 2021 Sebastian Crane <seabass-labrax@gmx.com>
// SPDX-License-Identifier: AGPL-3.0-or-later

package handle

import (
	"shuffalot/shuffle"
	"shuffalot/split"
)

func HandleLine(text string, seed int64) string {
	var output string

	if len(text) == 0 {
		return ""
	}

	for _, fragment := range split.SplitString(text) {
		if fragment.Shuffle == true {
			output += shuffle.ShuffleWord(fragment.Value, seed)
		} else {
			output += fragment.Value
		}
	}

	return output
}
