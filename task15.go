package main

import (
	"fmt"
	"os"
	"strings"
)

// В изначальной реализации задумка автора не очень понятна, каждый вызов функции
// перезаписывает глобальную переменную. При этом ссылка на v нигде не сохраняется
// Если это попытка создания буфера для динамической строки,
// то удобнее использовать слайс или StringBuilder, если задача в конкатенации.
// В крайнем случае можно сделать глобальной переменную v для реализации
// простейшего линейного аллокатора. Также если функция возвращает string, то
// v[:100] берёт 100 байт, что может обрезать символ юникод символ.
var justString string

func createHugeString(size int) string {
	var builder strings.Builder
	builder.Grow(size)
	for i := 0; i < size; i++ {
		builder.WriteRune('ф')
		builder.WriteRune('a')
	}
	return builder.String()
}

func someFunc(size int, capacity int) {
	v := createHugeString(capacity)
	justString = string([]rune(v)[:size])
}

func task15() {
	someFunc(100, 1 << 10)
	fmt.Fprintln(os.Stdout, justString)
}
