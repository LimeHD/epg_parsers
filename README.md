# Парсеры программ передач (EPG)

1. Парсеры могут писаться на любых языках.
2. Парсеры запусть периодически (запускалка, по-умолчанию по cron-у).
3. Для каждого сайта свой парсер.
4. У парсеров `стандартизированный вывод`.
5. Парсеры имеют логирование.
6. Если парсер не может толком распарсить, он пишет ошибку в лог и завершается с ошибкой (status > 0)

# Парсеры

- [digea_parser](https://www.digea.gr)
- [ept_parser](//program.ert.gr)

# Как собрать парсеры

- `make build`

# Как запустить

```
./digea_parser --format csv --output ./output
./ept_parser --format csv --output ./output
```

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
Fri, 17 Apr 2020 16:53:15 +0300;-ALERT-;https://www.digea.gr//images/channel-icons/ALERT.jpg;[K8] Ανατρεπτικό Δελτίο;2020-04-18T00:00:00+03:00;2020-04-18T01:00:00+03:00
Fri, 17 Apr 2020 16:53:15 +0300;-ALERT-;https://www.digea.gr//images/channel-icons/ALERT.jpg;[K8] Κόντρα Και Ρήξη - Τηλεπώληση;2020-04-18T01:00:00+03:00;2020-04-18T02:30:00+03:00
Fri, 17 Apr 2020 16:53:15 +0300;-VOLCANO-;https://www.digea.gr//images/channel-icons/VOLCANO TV EPG.jpg;Volcano;2020-04-17T09:00:00+03:00;2020-04-17T15:00:00+03:00
Fri, 17 Apr 2020 16:53:15 +0300;-VOLCANO-;https://www.digea.gr//images/channel-icons/VOLCANO TV EPG.jpg;Volcano;2020-04-17T15:00:00+03:00;2020-04-17T21:00:00+03:00
Fri, 17 Apr 2020 16:53:15 +0300;-VOLCANO-;https://www.digea.gr//images/channel-icons/VOLCANO TV EPG.jpg;Volcano;2020-04-17T21:00:00+03:00;2020-04-17T03:00:00+03:00
Fri, 17 Apr 2020 16:53:15 +0300;-STAR-;https://www.digea.gr//images/channel-icons/star.png;[K8] Ντετέκτιβ Μονκ - 3ος Κύκλος;2020-04-17T18:00:00+03:00;2020-04-17T18:15:00+03:00
Fri, 17 Apr 2020 16:53:15 +0300;-STAR-;https://www.digea.gr//images/channel-icons/star.png;[K8] Η Τρομερή Γκίλυ (The Great Gilly Hopkins);2020-04-17T18:15:00+03:00;2020-04-17T18:30:00+03:00
Fri, 17 Apr 2020 16:53:15 +0300;-STAR-;https://www.digea.gr//images/channel-icons/star.png;[K8] Η Κιβωτός Του Νώε (Noah's Ark);2020-04-17T18:30:00+03:00;2020-04-17T19:15:00+03:00
Fri, 17 Apr 2020 16:53:15 +0300;-STAR-;https://www.digea.gr//images/channel-icons/star.png;[K8] Μου Λείπεις Ήδη (Miss You Already);2020-04-17T19:15:00+03:00;2020-04-17T20:20:00+03:00
Fri, 17 Apr 2020 16:53:15 +0300;-STAR-;https://www.digea.gr//images/channel-icons/star.png;[K] Star News;2020-04-17T20:20:00+03:00;2020-04-17T22:50:00+03:00
Fri, 17 Apr 2020 16:53:15 +0300;-STAR-;https://www.digea.gr//images/channel-icons/star.png;[K8] Lion;2020-04-17T22:50:00+03:00;2020-04-18T00:00:00+03:00
Fri, 17 Apr 2020 16:53:15 +0300;-STAR-;https://www.digea.gr//images/channel-icons/star.png;[K12] Μια Θάλασσα Από Δέντρα (The Sea Of Trees);2020-04-18T00:00:00+03:00;2020-04-18T02:30:00+03:00
Fri, 17 Apr 2020 16:53:15 +0300;-STAR-;https://www.digea.gr//images/channel-icons/star.png;[K12] Η Σφαίρα (Sphere);2020-04-18T02:30:00+03:00;2020-04-17T05:00:00+03:00
Fri, 17 Apr 2020 16:53:15 +0300;-STAR-;https://www.digea.gr//images/channel-icons/star.png;[K12] Η Κατάρα Της Σφίγγας (Sphinx);2020-04-17T05:00:00+03:00;2020-04-17T06:30:00+03:00
Fri, 17 Apr 2020 16:53:15 +0300;-STAR-;https://www.digea.gr//images/channel-icons/star.png;[K8] Psych - 6ος Κύκλος;2020-04-17T06:30:00+03:00;2020-04-17T07:45:00+03:00
Fri, 17 Apr 2020 16:53:15 +0300;-EURO CHANNEL-;https://www.digea.gr//images/channel-icons/SUPER TV EPG.jpg;Euro Channel;2020-04-17T09:00:00+03:00;2020-04-17T15:00:00+03:00
Fri, 17 Apr 2020 16:53:15 +0300;-EURO CHANNEL-;https://www.digea.gr//images/channel-icons/SUPER TV EPG.jpg;Euro Channel;2020-04-17T15:00:00+03:00;2020-04-17T21:00:00+03:00
Fri, 17 Apr 2020 16:53:15 +0300;-EURO CHANNEL-;https://www.digea.gr//images/channel-icons/SUPER TV EPG.jpg;Euro Channel;2020-04-17T21:00:00+03:00;2020-04-17T03:00:00+03:00
```

## Ядро которое собирает спарсенные данные и складывает в базу.

### Список всех зон для инициализации локальных временных зон

- [List of tz database time zones](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)
