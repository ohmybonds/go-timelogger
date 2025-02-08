package timelogger

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

// ANSI-коды для цветов
const (
	Reset  = "\033[0m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Red    = "\033[31m"
)

// Пороговые значения для подсветки
const (
	ThresholdGreen  = 5 * time.Millisecond
	ThresholdYellow = 20 * time.Millisecond
)

// StepTime логирует время выполнения с автоматическим уровнем вложенности
func StepTime(message string, start time.Time) time.Time {
	now := time.Now()
	elapsed := now.Sub(start)

	// Определяем цвет по длительности выполнения
	color := getColor(elapsed)

	// Выводим результат с цветом
	fmt.Printf("%s%-40s: %v%s\n", color, message, elapsed, Reset)

	return now
}

// getIndentLevel определяет уровень вложенности по стеку вызовов
func getIndentLevel() int {
	depth := 0
	for i := 2; i < 10; i++ { // Начинаем со 2-го уровня, чтобы исключить сам StepTime
		_, file, _, ok := runtime.Caller(i)
		if !ok || !strings.Contains(file, "timelogger") { // Прерываем, если стек вызова не относится к нашему пакету
			break
		}
		depth++
	}
	return depth
}

// getColor возвращает цвет в зависимости от времени выполнения
func getColor(elapsed time.Duration) string {
	switch {
	case elapsed >= ThresholdYellow:
		return Red
	case elapsed >= ThresholdGreen:
		return Yellow
	default:
		return Green
	}
}
