// Copyright 2018 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testparsers

import (
	"bytes"

	. "github.com/cilium/cilium/proxylib/proxylib"

	log "github.com/sirupsen/logrus"
)

//
// Line parser used for testing
//

type LineParserFactory struct{}

var lineParserFactory *LineParserFactory

func init() {
	log.Info("init(): Registering lineParserFactory")
	RegisterParserFactory("test.lineparser", lineParserFactory)
}

type LineParser struct {
	connection *Connection
	inserted   bool
}

func (p *LineParserFactory) Create(connection *Connection) Parser {
	log.Infof("LineParserFactory: Create: %v", connection)
	return &LineParser{connection: connection}
}

func getLine(data [][]byte, offset int) ([]byte, bool) {
	var line bytes.Buffer
	for i, s := range data {
		index := bytes.IndexByte(s[offset:], '\n')
		if index < 0 {
			line.Write(s[offset:])
		} else {
			log.Infof("getLine: unit: %d offset: %d length: %d index: %d", i, offset, len(s), index)
			line.Write(s[offset : offset+index+1])
			return line.Bytes(), true
		}
		offset = 0
	}
	return line.Bytes(), false
}

//
// Parses individual lines that must start with one of:
// "PASS" the line is passed
// "DROP" the line is dropped
// "INJECT" the line is injected in reverse direction
// "INSERT" the line is injected in current direction
//
func (p *LineParser) OnData(reply, endStream bool, data [][]byte, offset int) (OpType, int) {
	line, ok := getLine(data, offset)
	line_len := len(line)

	if p.inserted {
		p.inserted = false
		return DROP, line_len
	}

	if !reply {
		log.Infof("LineParser: Request: %s", line)
	} else {
		log.Infof("LineParser: Response: %s", line)
	}

	if !ok {
		if line_len > 0 {
			// Partial line received, but no newline, ask for more
			return MORE, 1
		} else {
			// Nothing received, don't know if more will be coming; do nothing
			return NOP, 0
		}
	}

	if bytes.HasPrefix(line, []byte("PASS")) {
		return PASS, line_len
	}
	if bytes.HasPrefix(line, []byte("DROP")) {
		return DROP, line_len
	}
	if bytes.HasPrefix(line, []byte("INJECT")) {
		// Inject line in the reverse direction
		p.connection.Inject(!reply, []byte(line))
		// Drop the INJECT line in the current direction
		return DROP, line_len
	}
	if bytes.HasPrefix(line, []byte("INSERT")) {
		// Inject the line in the current direction
		p.connection.Inject(reply, []byte(line))
		// Drop the INJECT line in the current direction
		p.inserted = true
		return INJECT, line_len
	}

	return ERROR, int(ERROR_INVALID_FRAME_TYPE)
}
