//     Copyright (C) 2020-2021, IrineSistiana
//
//     This file is part of mosdns.
//
//     mosdns is free software: you can redistribute it and/or modify
//     it under the terms of the GNU General Public License as published by
//     the Free Software Foundation, either version 3 of the License, or
//     (at your option) any later version.
//
//     mosdns is distributed in the hope that it will be useful,
//     but WITHOUT ANY WARRANTY; without even the implied warranty of
//     MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
//     GNU General Public License for more details.
//
//     You should have received a copy of the GNU General Public License
//     along with this program.  If not, see <https://www.gnu.org/licenses/>.

package upstream

const (
	// DNS header size is 12. It is also the minimum length of a valid dns msg.
	headerSize = 12
)

func setMsgId(m []byte, id uint16) {
	m[0] = byte(id >> 8)
	m[1] = byte(id)
}

func getMsgId(m []byte) uint16 {
	return uint16(m[0])<<8 + uint16(m[1])
}

func isTruncated(m []byte) bool {
	return m[3]&1<<1 != 0
}
