# fiter
Обходит директорию и выводит информацию о файле: 
Полный путь, время последнего изменения и просмотра

Использование:

Обход директории /dir и заполнение файла output.csv.

    ./fiter -o=output.csv /dir

Обход директории /dir и вывод информации на экран

    ./fiter /dir

Результаты бенчмарков:

Тестовая среда: директория - 52 объекта; 4,82 ГБ

## BenchmarkMain
Тестовая среда: директория - 52 объекта; 4,82 ГБ

42775 B/op - байты, выделенные за операцию

543 allocs/op - количество аллокаций памяти за операцию

1.751s - время выполнения

## Hyperfine
Тестовая среда: директория - 12361 объектов; 955 МБ
### Старт
hyperfine --warmup 3 './fiter -o=output.csv /Users/artem/Documents'
### Итог

Time (mean ± σ):     162.3 ms ±   5.5 ms
 
[User: 45.6 ms, System: 143.8 ms]

Range (min … max):   155.9 ms … 172.8 ms

18 runs

## Time
Тестовая среда: директория - 52 объекта; 4,82 ГБ

### Старт
./fiter -o=output.csv /Users/artem/go/src/fiter/test_env
### Итог
0.00s user 0.01s system 79% cpu 0.014 total

## Как можно уменьшить потребление ресурсов?:

Реализация функции WalkDir в path/filepath выполняет копирование всех наименований из директории в память и их лексографическую сортировку -> необходимо больше ресурсов.

Решением может стать написание своей рекурсивной функции обхода.

Однако предварительные тесты показали, что уменьшение количества аллокаций незначительное (10-20 allocs/op) и к тому же постоянно скачет (возможно это погрешность).
