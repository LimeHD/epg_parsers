# Парсеры программ передач (EPG)

1. Парсеры могут писаться на любых языках.
2. Парсеры запусть периодически (запускалка, по-умолчанию по cron-у).
3. Для каждого сайта свой парсер.
4. У парсеров `стандартизированный вывод`.
5. Парсеры имеют логирование.
6. Если парсер не может толком распарсить, он пишет ошибку в лог и завершается с ошибкой (status > 0)

# Парсеры

- `./bin/digea_parser` парсит https://www.digea.gr
- `./bin/ept_parser` парсет http://program.ert.gr

# Сборка

- `make build`

# Запуск

Собирает и выводит в файл digea_parser.csv

```
./bin/digea_parser # Собирает данные и сбрасывает в ./digea_parser.csv
./bin/digea_parser --output ./output/digea_parser.csv --bugsnag_api_key {key}
```

# TODO

* [ ] После запуска парсера писать в STDOUT название парсера, время и что он
  делает. Желательно показывать какой-то прогресс точками, например вводить
  точку после каждого запроса или тп. Выводить куда пишет (в какой файл)
* [ ] Сменить формат вывода: 1. Первой строкой выводить заголовок. 2. Сменить
  порядок и перечень колонок согласно описаниию ниже.
* [x] Аргумент `--output` вместо каталога ожидается что можно будет указвать
  конкретный файл куда выводить.
* [ ] zabbix (удача/не удача). (`--zabbix-server=zx.iptv2022.com --zabbix-port=10051 --zabbix-key=epg_parsers.digea --zabbix-host=epg.iptv2022.com`
* [x] bugsnag (`--bugsnag-api-key`)

# Вывод парсера

Вывод в формате CSV, separated by ";". 1 - строка - одна передача.

Следующие строки:

* Дата и время начала
* Дата и время конца
* Канал (название, строка)
* Передача (название, строка)
* Лого канала (URL, строка)
* Описание передачи (текст)

Обязательный заголовок:

```csv
datetime_start;datetime_finish;channel;broardcast;channel_logo_url;description
```

Пример (10 строк):

```csv
Fri, 17 Apr 2020 16:53:15 +0300;-ALERT-;https://www.digea.gr//images/channel-icons/ALERT.jpg;[K8] Κόντρα Και Ρήξη - Τηλεπώληση;2020-04-17T19:00:00+03:00;2020-04-17T20:00:00+03:00
Fri, 17 Apr 2020 16:53:15 +0300;-ALERT-;https://www.digea.gr//images/channel-icons/ALERT.jpg;[K8] El Kady;2020-04-17T20:00:00+03:00;2020-04-17T21:00:00+03:00
Fri, 17 Apr 2020 16:53:15 +0300;-ALERT-;https://www.digea.gr//images/channel-icons/ALERT.jpg;[K8] Φίλτρα Νερού Nanofilter.gr - Τηλεπώληση;2020-04-17T21:00:00+03:00;2020-04-17T22:00:00+03:00
Fri, 17 Apr 2020 16:53:15 +0300;-ALERT-;https://www.digea.gr//images/channel-icons/ALERT.jpg;[K8] Φίλτρα Νερού Nanofilter.gr - Τηλεπώληση;2020-04-17T22:00:00+03:00;2020-04-17T22:30:00+03:00
Fri, 17 Apr 2020 16:53:15 +0300;-ALERT-;https://www.digea.gr//images/channel-icons/ALERT.jpg;[K8] Φίλτρα Νερού Nanofilter.gr - Τηλεπώληση;2020-04-17T22:30:00+03:00;2020-04-17T23:00:00+03:00
Fri, 17 Apr 2020 16:53:15 +0300;-ALERT-;https://www.digea.gr//images/channel-icons/ALERT.jpg;[K8] Τα Νέα Της Αγοράς;2020-04-17T23:00:00+03:00;2020-04-18T00:00:00+03:00
```

## Ядро которое собирает спарсенные данные и складывает в базу.

### Список всех зон для инициализации локальных временных зон

Не понял зачем это и для чего [dapi]

- [List of tz database time zones](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)
