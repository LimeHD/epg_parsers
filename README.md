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

* [x] После запуска парсера писать в STDOUT название парсера, время и что он
  делает. Желательно показывать какой-то прогресс точками, например вводить
  точку после каждого запроса или тп. Выводить куда пишет (в какой файл)
* [x] Сменить формат вывода: 1. Первой строкой выводить заголовок. 2. Сменить
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
2020-04-17T09:00:00+03:00;2020-04-17T09:30:00+03:00;-STAR K.E.-;;https://www.digea.gr//images/channel-icons/STAR KENTIKHS ELLADAS.jpg;[K] Be Μy Guest;
2020-04-17T09:30:00+03:00;2020-04-17T10:00:00+03:00;-STAR K.E.-;;https://www.digea.gr//images/channel-icons/STAR KENTIKHS ELLADAS.jpg;[K] Παιδικό Πρόγραμμα;
2020-04-17T10:00:00+03:00;2020-04-17T11:00:00+03:00;-STAR K.E.-;;https://www.digea.gr//images/channel-icons/STAR KENTIKHS ELLADAS.jpg;[K] 24 Ώρες Ρεπορτάζ (Ε);
2020-04-17T11:00:00+03:00;2020-04-17T11:30:00+03:00;-STAR K.E.-;;https://www.digea.gr//images/channel-icons/STAR KENTIKHS ELLADAS.jpg;[K] Όρεξη Να ΄Χεις (Ε);
2020-04-17T11:30:00+03:00;2020-04-17T12:00:00+03:00;-STAR K.E.-;;https://www.digea.gr//images/channel-icons/STAR KENTIKHS ELLADAS.jpg;[K] Αποκαθήλωση (Μ. Παρασκευή);
2020-04-17T12:00:00+03:00;2020-04-17T14:30:00+03:00;-STAR K.E.-;;https://www.digea.gr//images/channel-icons/STAR KENTIKHS ELLADAS.jpg;[K] The Jesus Life;
2020-04-17T14:30:00+03:00;2020-04-17T17:30:00+03:00;-STAR K.E.-;;https://www.digea.gr//images/channel-icons/STAR KENTIKHS ELLADAS.jpg;[K] Μεσημβρινό Δελτίο Ειδήσεων;
2020-04-17T17:30:00+03:00;2020-04-17T18:30:00+03:00;-STAR K.E.-;;https://www.digea.gr//images/channel-icons/STAR KENTIKHS ELLADAS.jpg;[K] In Style;
2020-04-17T18:30:00+03:00;2020-04-17T19:30:00+03:00;-STAR K.E.-;;https://www.digea.gr//images/channel-icons/STAR KENTIKHS ELLADAS.jpg;[K] Παιδικό Πρόγραμμα;
2020-04-17T19:30:00+03:00;2020-04-17T20:55:00+03:00;-STAR K.E.-;;https://www.digea.gr//images/channel-icons/STAR KENTIKHS ELLADAS.jpg;[K] Σύντομα Γεγονότα;
2020-04-17T20:55:00+03:00;2020-04-17T21:00:00+03:00;-STAR K.E.-;;https://www.digea.gr//images/channel-icons/STAR KENTIKHS ELLADAS.jpg;[K] Ακολουθία Επιταφίου;
```

## Ядро которое собирает спарсенные данные и складывает в базу.

### Список всех зон для инициализации локальных временных зон

Не понял зачем это и для чего [dapi]

- [List of tz database time zones](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)
