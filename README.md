# Парсеры программ передач (EPG)

1. Парсеры могут писаться на любых языках.
2. Парсеры запусть периодически (запускалка, по-умолчанию по cron-у).
3. Для каждого сайта свой парсер.
4. У парсеров `стандартизированный вывод`.
5. Парсеры имеют логирование.
6. Если парсер не может толком распарсить, он пишет ошибку в лог и завершается с ошибкой (status > 0)

# Парсеры

## Скидывают в CSV

- `./bin/digea_parser` парсит https://www.digea.gr
- `./bin/ept_parser` парсит http://program.ert.gr
- `./bin/star_parser` парсит https://www.star.gr
- `./bin/download_stv && ./bin/stv_parser` парсит xml от поставщика s-tv.ru
- `./bin/dvbs_parser` парсит с dvbs
- `./bin/graella_parser` парсит с источника от испанцев
- `./bin/alfaomega` парсит с https://alfaomega.tv/canal-tv/programul-tv
- `./bin/kazakh` парсит .xlsx файл Kazakh TV

## Складывают прямо в базу

- `./bin/download_tv_pack && ./bin/tv_pack_parser` парсит скачанную стандартную XML и скидывает в базу (в
  таблицу `epg`)

### Ручной запуск стандартного парсера

- `php store/index.php -h {MYSQL_HOST} -l {MYSQL_USER} -p {MYSQL_PASSWORD} -d {MYSQL_DATABASE} -i ./TV_Pack.xml -f 2020-07-19 -g ./log/log.log -c 2`, **где** опциональные параметры:
- [x] `-f` - с какой даты начинать сканировать: default `today`, 
- [x] `-c` - сколько дней от текущего дня добавлять в БД: default `5`, 
- [x] `-g` - куда складывать логи: default `null`.

# Сборка

- `make build`

# Пример запуска парсера

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

Вывод в формате CSV, разделитель - табуляция. 1 - строка - одна передача.

Следующие строки:

* Дата и время начала (RFC3339) с таймзоной
* Дата и время конца (RFC3339) с таймзоной
* Канал (название, строка)
* Передача (название, строка)
* Лого канала (URL, строка)
* Описание передачи (текст)
* Доступность архива передачи (число 1 или 0, доступно на данный момент только для Испании)
* Доступность передачи (число 1 или 0, доступно на данный момент только для Испании)

Обязательный заголовок:

```csv
datetime_start	datetime_finish	channel	title	channel_logo_url	description available_archive geo_regions
```

Пример (10 строк):

```csv
2020-04-17T09:00:00+03:00	2020-04-17T09:30:00+03:00	-STAR K.E.-		https://www.digea.gr//images/channel-icons/STAR KENTIKHS ELLADAS.jpg	[K] Be Μy Guest	
2020-04-17T09:30:00+03:00	2020-04-17T10:00:00+03:00	-STAR K.E.-		https://www.digea.gr//images/channel-icons/STAR KENTIKHS ELLADAS.jpg	[K] Παιδικό Πρόγραμμα	
2020-04-17T10:00:00+03:00	2020-04-17T11:00:00+03:00	-STAR K.E.-		https://www.digea.gr//images/channel-icons/STAR KENTIKHS ELLADAS.jpg	[K] 24 Ώρες Ρεπορτάζ (Ε)	
2020-04-17T11:00:00+03:00	2020-04-17T11:30:00+03:00	-STAR K.E.-		https://www.digea.gr//images/channel-icons/STAR KENTIKHS ELLADAS.jpg	[K] Όρεξη Να ΄Χεις (Ε)	
2020-04-17T11:30:00+03:00	2020-04-17T12:00:00+03:00	-STAR K.E.-		https://www.digea.gr//images/channel-icons/STAR KENTIKHS ELLADAS.jpg	[K] Αποκαθήλωση (Μ. Παρασκευή)	
2020-04-17T12:00:00+03:00	2020-04-17T14:30:00+03:00	-STAR K.E.-		https://www.digea.gr//images/channel-icons/STAR KENTIKHS ELLADAS.jpg	[K] The Jesus Life	
2020-04-17T14:30:00+03:00	2020-04-17T17:30:00+03:00	-STAR K.E.-		https://www.digea.gr//images/channel-icons/STAR KENTIKHS ELLADAS.jpg	[K] Μεσημβρινό Δελτίο Ειδήσεων	
2020-04-17T17:30:00+03:00	2020-04-17T18:30:00+03:00	-STAR K.E.-		https://www.digea.gr//images/channel-icons/STAR KENTIKHS ELLADAS.jpg	[K] In Style	
2020-04-17T18:30:00+03:00	2020-04-17T19:30:00+03:00	-STAR K.E.-		https://www.digea.gr//images/channel-icons/STAR KENTIKHS ELLADAS.jpg	[K] Παιδικό Πρόγραμμα	
2020-04-17T19:30:00+03:00	2020-04-17T20:55:00+03:00	-STAR K.E.-		https://www.digea.gr//images/channel-icons/STAR KENTIKHS ELLADAS.jpg	[K] Σύντομα Γεγονότα	
2020-04-17T20:55:00+03:00	2020-04-17T21:00:00+03:00	-STAR K.E.-		https://www.digea.gr//images/channel-icons/STAR KENTIKHS ELLADAS.jpg	[K] Ακολουθία Επιταφίου	
```

## Ядро которое собирает спарсенные данные и складывает в базу.

### Список всех зон для инициализации локальных временных зон

Не понял зачем это и для чего [dapi]

- [List of tz database time zones](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)
