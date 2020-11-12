package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// inputText
type inputText struct {
	text string
}

func (i *inputText) getText() string {
	return i.text
}

func (i *inputText) append(input string) {
	var b strings.Builder
	b.WriteString(i.text)
	b.WriteString(input)
	i.text = b.String()
}

func (i *inputText) createSnapshot() *snapshot {
	return &snapshot{
		i.text,
	}
}

func (i *inputText) restoreSnapshot(s *snapshot) {
	i.text = s.getText()
}

// snapshot
type snapshot struct {
	text string
}

func (s *snapshot) getText() string {
	return s.text
}

// snapshotHolder
type snapshotHolder struct {
	snapshots []*snapshot
}

func (s *snapshotHolder) popSnapshot() *snapshot {
	tail := s.snapshots[len(s.snapshots)-1]
	s.snapshots = s.snapshots[:len(s.snapshots)-1]
	return tail
}

func (s *snapshotHolder) pushSnapshot(newSnap *snapshot) {
	s.snapshots = append(s.snapshots, newSnap)
}

// main
func main() {
	data := &inputText{}
	snapshotHolder := &snapshotHolder{}

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if line == ":list" {
			fmt.Println(data.getText())
		} else if line == ":undo" {
			newSnap := snapshotHolder.popSnapshot()
			data.restoreSnapshot(newSnap)
		} else {
			snapshotHolder.pushSnapshot(data.createSnapshot())
			data.append(line)
		}
	}
}
