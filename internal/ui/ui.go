package ui

import (

    "github.com/fatih/color"
)

func Success(msg string) {
    color.Green("✓ %s", msg)
}

func LoggerError(err error) {
    color.Red("✗ %v", err)
}

func Error(msg string) {
    color.Red("✗ %v", msg)
}

func Info(msg string) {
    color.Cyan("ℹ %s", msg)
}

func Warn(msg string) {
    color.Yellow("⚠ %s", msg)
}

