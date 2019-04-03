// Copyright 2019 Andrei Tudor Călin
//
// Permission to use, copy, modify, and/or distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

package mem_test

import (
	"fmt"
	"log"

	"acln.ro/mem"
)

/*
IPv4Header is the Internet Protocol version 4 header.

RFC 791 illustrates the header as:

0                   1                   2                   3
0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|Version|  IHL  |Type of Service|          Total Length         |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|         Identification        |Flags|      Fragment Offset    |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|  Time to Live |    Protocol   |         Header Checksum       |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                       Source Address                          |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                    Destination Address                        |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|                    Options                    |    Padding    |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
*/
var ipv4Header = mem.Layout{
	Fields: []mem.Field{
		{
			Name: "version",
			Bits: 4,
		},
		{
			Name: "ihl",
			Bits: 4,
		},
		// 0     1     2     3     4     5     6     7
		// +-----+-----+-----+-----+-----+-----+-----+-----+
		// |                 |     |     |     |     |     |
		// |   PRECEDENCE    |  D  |  T  |  R  |  0  |  0  |
		// |                 |     |     |     |     |     |
		// +-----+-----+-----+-----+-----+-----+-----+-----+
		{
			Name: "tos",
			Desc: "type of service",
			Bits: 8,
			Layout: mem.Layout{
				Fields: []mem.Field{
					{
						Name: "precedence",
						Bits: 3,
					},
					{
						Name: "delay",
						Bits: 1,
					},
					{
						Name: "throughput",
						Bits: 1,
					},
					{
						Name: "reliability",
						Bits: 1,
					},
					{
						Name: mem.Reserved,
						Bits: 2,
					},
				},
			},
		},
		{
			Name: "tot_len",
			Desc: "total length of the datagram, in octets",
			Bits: 16,
		},
		{
			Name: "id",
			Bits: 16,
		},
		{
			Name: "frag_off",
			Bits: 16,
			Layout: mem.Layout{
				Fields: []mem.Field{
					{
						Name: "flags",
						Bits: 3,
					},
					{
						Name: "offset",
						Bits: 13,
					},
				},
			},
		},
		{
			Name: "ttl",
			Desc: "time to live",
			Bits: 8,
		},
		{
			Name: "protocol",
			Desc: "protocol used in data portion",
			Bits: 8,
		},
		{
			Name: "check",
			Desc: "header checksum",
			Bits: 16,
		},
		{
			Name: "saddr",
			Desc: "source address",
			Bits: 32,
		},
		{
			Name: "daddr",
			Desc: "destination address",
			Bits: 32,
		},
	},
}

func Example() {
	if err := ipv4Header.Init(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(ipv4Header.Offsetof["protocol"]) // Output: 72
}
