// SPDX-License-Identifier: AGPL-3.0-only
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published
// by the Free Software Foundation, version 3.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <https://www.gnu.org/licenses/>

package wiki

import (
	"strconv"
	"strings"
)

var _ interface {
	Error() string
	Unwrap() error
} = (*SyntaxError)(nil)

type SyntaxError struct {
	Err     error
	Line    string
	Lino    int
	infobox string
}

func (w *SyntaxError) ReadableError() string {
	lines := strings.Split(w.infobox, "\n")
	show := lines[max(w.Lino-3, 0) : w.Lino-1]

	var buf strings.Builder
	for _, line := range show {
		buf.WriteString(line)
		buf.WriteRune('\n')
	}

	buf.WriteString(strings.Repeat("^", len(lines[w.Lino-1])))
	buf.WriteRune('\n')

	buf.WriteString(w.Err.Error())
	return buf.String()
}

func (p *SyntaxError) Error() string {
	return p.Err.Error() + " line: " + strconv.Itoa(p.Lino) + " " + strconv.Quote(p.Line)
}

func (p *SyntaxError) Unwrap() error {
	return p.Err
}

func wrapError(err error, lino int, line string, infobox string) error {
	return &SyntaxError{
		Line:    line,
		Lino:    lino,
		Err:     err,
		infobox: infobox,
	}
}
