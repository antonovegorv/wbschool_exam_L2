Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее устройство слайсов и что происходит при передачи их в качестве аргументов функции.

```go
package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}
```

Ответ:
```
Вывод программы:
[3 2 3]

Пояснение:
Для того, чтобы понять, почему мы получили такой результат, нужно вспомнить о том, как устроены срезы внутри. Внутри себя срезы хранят ссылку на массив, длину среза (кол-во элементов) и ёмкость (максимальное кол-во элементов). Если мы превысим ёмкость среза, то в таком случае у нас произойдет выделение нового массива, на который будет указывать уже новый срез.

```
