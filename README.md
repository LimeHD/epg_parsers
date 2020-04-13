# Парсеры программ передач (EPG)

1. Парсеры могут писаться на любых языках.
2. Парсеры запусть периодически (запускалка, по-умолчанию по cron-у).
3. Для каждого сайта свой парсер.
4. У парсеров `стандартизированный вывод`.
5. Парсеры имеют логирование.
6. Если парсер не может толком распарсить, он пишет ошибку в лог и завершается с ошибкой (status > 0)

# Как запустить парсер?

Пример

# Список парсеров

- [Digea](https://www.digea.gr)
- [EPT](https://program.ert.gr)

# Вывод парсера

`CSV`

1 - строка - одна передача.

* Канал (название, строка)
* Передача (название, строка)
* Дата и время начала
* Дата и время конца
* Описание передачи (текст)

Пример (10 строк):


## Ядро которое собирает спарсенные данные и складывает в базу.

### Список всех зон для инициализации локальных временных зон

- [List of tz database time zones](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)
