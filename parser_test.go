package main

import (
	"testing"
)

func TestNumberParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"",
			"",
		},
		{
			"1",
			"1",
		},
	}

	for _, tt := range tests {
		l := NewLexer(tt.input)
		p := NewParser(l)
		actual := p.Parse()

		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}

func TestPrefixParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"-1",
			"(-1)",
		},
		{
			"--1",
			"(-(-1))",
		},
	}

	for _, tt := range tests {
		l := NewLexer(tt.input)
		p := NewParser(l)
		actual := p.Parse()

		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}

func TestInfixParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"1 + 2",
			"(1 + 2)",
		},
		{
			"1 + -2",
			"(1 + (-2))",
		},
		{
			"1 + 2 + 3",
			"((1 + 2) + 3)",
		},
		{
			"1 + 2 * 3 + 4 / 5 - 6",
			"(((1 + (2 * 3)) + (4 / 5)) - 6)",
		},
	}

	for _, tt := range tests {
		l := NewLexer(tt.input)
		p := NewParser(l)
		actual := p.Parse()

		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}

func TestPostfixParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"1!",
			"(1!)",
		},
		{
			"1 + 2!",
			"(1 + (2!))",
		},
		{
			"1 + 2! + 3",
			"((1 + (2!)) + 3)",
		},
	}

	for _, tt := range tests {
		l := NewLexer(tt.input)
		p := NewParser(l)
		actual := p.Parse()

		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}

func TestParenParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"(1 + 2)",
			"(1 + 2)",
		},
		{
			"(1 + 2) + 3",
			"((1 + 2) + 3)",
		},
		{
			"1 + (2 + 3) + 4",
			"((1 + (2 + 3)) + 4)",
		},
		{
			"(5 + 5) * 2",
			"((5 + 5) * 2)",
		},
		{
			"2 / (5 + 5)",
			"(2 / (5 + 5))",
		},
		{
			"(5 + 5) * 2 * (5 + 5)",
			"(((5 + 5) * 2) * (5 + 5))",
		},
		{
			"1 + (2 + 3) * 4",
			"(1 + ((2 + 3) * 4))",
		},
		{
			"-(5 + 5)",
			"(-(5 + 5))",
		},
	}

	for _, tt := range tests {
		l := NewLexer(tt.input)
		p := NewParser(l)
		actual := p.Parse()

		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}
