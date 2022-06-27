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
Slice представляет собой своеобразную "упаковку" над массивом, является структурой, которая хранит в себе размер,
вместимость и ссылку на исходный массив. Все изменения, осуществляемые в срезе, отображаются и в исходном массиве,
до того момента, как срез не перестанет указывать на исходный массив. Такое может произойти, если при appende размер
массива станет больше его вместимости. В таком случае копируются все значения из старого массива и записывается в новый

В данном случае программа на экран выведет 3 2 3. Дело в функции modifySlice, в аргументах которого мы прокидываем
ЗНАЧЕНИЕ среза, а не его самого. Что происходит внутри функции:

1. Изменяем в исходном массиве первый элемент на тройку, все ок
2. Добавляем еще один элемент в срез. В этот момент выделятся новый массив и срезу присваивается ссылка на новосозданный
массив, внутри функции к исхоному массиву отношение мы больше не имеем. Новый массив имеет вместимости в два раза больше
вместимости прошлого массива
3. Изменяем второй элемент массива на 5, на исходный массив это никакак не сказывается, поскольку нет на него ссылки
4. Добавляем еще один элемент к срезу, на исходный массив это так же никиках не сказывается. Новый массив не создается

Для того, чтобы исправить эту "оплошность", вместо передачи среза по значение следует его передать по ссылке
```