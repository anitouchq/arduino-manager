package main

import "testing"

func TestParseBoardListExtractsPortsAndFQBN(t *testing.T) {
	output := "Port Protocol Type Board Name FQBN Core\n" +
		"COM4 serial Serial Port (USB) Arduino Uno arduino:avr:uno arduino:avr\n" +
		"/dev/ttyACM0 serial Serial Port Arduino Mega arduino:avr:mega arduino:avr\n"

	got := ParseBoardList(output)
	if len(got) != 2 {
		t.Fatalf("ParseBoardList() returned %d boards, want 2", len(got))
	}
	if got[0].Port != "COM4" || got[0].FQBN != "arduino:avr:uno" {
		t.Fatalf("first board = %#v, want COM4 / arduino:avr:uno", got[0])
	}
	if got[1].Port != "/dev/ttyACM0" || got[1].FQBN != "arduino:avr:mega" {
		t.Fatalf("second board = %#v, want /dev/ttyACM0 / arduino:avr:mega", got[1])
	}
}

func TestParseBoardListReturnsEmptyWhenNoBoardsFound(t *testing.T) {
	if got := ParseBoardList("No boards found.\n"); len(got) != 0 {
		t.Fatalf("ParseBoardList() returned %d boards, want 0", len(got))
	}
}
