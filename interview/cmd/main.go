package main

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"sync"
)

// В RAM одна горутина около 4кб

//  В чем различия goroutine от потока системы?
// по простому - это более высокоуровневые потоки.
// К примеру, если нужно выполнять очень мелкие действия в отдельных потоках, то обычные потоки будут накладными - система будет дольше их создавать.
// для goroutines run-time может переиспользовать один и тот же поток, эмулируя работу настоящих потоков.
// В результате получается быстро и дешево.
// Проблема номер два - некоторые процессы могут создавать очень большое количество потоков (например, из за ошибки программиста).
// goroutines такого не должно случиться - они будут прост в очереди.

// Как огранить число потоков на систему при запуске Golang программы и возможно ли огранить их до 1 потока?

// ограничить число потоков можно указав runtime.GOMAXPROCS(1) здесь мы указываем количество ЛОГИЧЕСКИХ ЯДЕР ПРОЦЕССОРА, которые могут
// работать с прогой. Ограничение будет по количеству потоков, доступных одному ядру процессора
// если runtime.GOMAXPROCS(0) где (0) значение меньше 1  // изменений с предыдущей настройкой не будет сама функция возвращает предыдущую настройку

func oneThread() {
	fmt.Println("----------------ONE THREAD----------------")
	fmt.Println("runtime GOMAXPROCS : ", runtime.GOMAXPROCS(1))

	fmt.Println(runtime.NumCPU())
	//data := 10
	ch := make(chan int)

	go func() {
		i := 0
		for i = 0; i < 1000000000; i++ {
		}
		fmt.Println("boob")
		ch <- i

	}()

	fmt.Println("wow")

	x := <-ch // блокировка main goroutine, пока не будет прочитано из канала
	fmt.Println("bob")
	close(ch)

	fmt.Printf("x : %v\n", x)
}

// Что из себя представляет тип данных string в языке Golang? В Go строка в действительности является срезом (slice) байтов, доступным только для чтения.
// Можно ли изменить определенный символ в строке? в golang нельзя, так как строка в go это слайс байтов
// Что происходит при склеивании строк? ОТВЕТ: создание новой строки
// при итерации (for i := 0; i < len(str); i++{}) происходит итерирование по байтам
// при итерации range (for i, val := range str{}) происходит итерирование по рунам
// rune как псевдоним для типа int32
// byte как псевдоним для типа uint8
// кодировка UTF-8
func whatIsString() {
	fmt.Println("----------------whatIsString----------------")
	var str string = "прост"
	fmt.Println(reflect.TypeOf(str))

	for i := 0; i < len(str); i++ {
		fmt.Println(str[i])
	}

	for i, val := range str {
		fmt.Printf("i = %d, rune = %c\n", i, val)
	}
}

//Вытекающий вопрос — как эффективно склеивать множество строк? ОТВЕТ: через Strings.builder
func clueString() {
	var builder strings.Builder

	builder.WriteString("просто")
	builder.WriteString("Обычная")
	builder.WriteString("Строка")

	fmt.Println(builder.String())
}

// Что будет происходить при конкуррентной записи в map? Ошибка Будет dataRace map - потоконебезопасный тип данных : fatal error: concurrent map iteration and map write
// чтобы отловить такой баг, можно восользоваться директивой компиляции go run -race ./cmd/main.go
// Как можно решить эту проблему? использовать sync.Mutex (mu := &sync.Mutex{}) мьтекс необходим передавать по указателю в горутину
// А если еще плотнее, использовать RWMutex

func concurencyMap() {
	mp := make(map[string]string)
	//mp["key"] = "val0"
	//ch := make(chan int, 2)

	mu := &sync.Mutex{}

	go func(mp map[string]string, mu *sync.Mutex) {
		mu.Lock()
		defer mu.Unlock()
		mp["key"] = "valParallel"

	}(mp, mu)
	go func(mp map[string]string, mu *sync.Mutex) {
		mu.Lock()
		defer mu.Unlock()
		mp["key"] = "valParallel2"

	}(mp, mu)

	mu.Lock()
	fmt.Println(mp)
	mu.Unlock()
}

// Как устроен слайс и чем он отличается от массива? В отличии от массива slice имеет capacity (т.е. текущую вместимость)
// структура слайса такова {размер, указатель на беспрерывный кластер памяти, вместимость}

//Нужно ли лочить структуру мьютексом, если идет конкуррентная запись в разные поля структуры? - я думаю, что не нужно

func main() {
	// oneThread()
	whatIsString()
	clueString()
	concurencyMap()
}
