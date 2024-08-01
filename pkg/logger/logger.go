package logger

import (
	"context"
	"fmt"
)

func Info(ctx context.Context, message string, fields ...interface{}) {
	fmt.Println(fmt.Sprintf(message, fields...))
}

func Debug(ctx context.Context, message string, fields ...interface{}) {
	fmt.Println(fmt.Sprintf(message, fields...))
}

func Warn(ctx context.Context, message string, fields ...interface{}) {
	fmt.Println(fmt.Sprintf(message, fields...))
}

func Error(ctx context.Context, message string, fields ...interface{}) {
	fmt.Println(fmt.Sprintf(message, fields...))
}
