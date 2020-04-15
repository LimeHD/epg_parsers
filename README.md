# Парсеры программ передач (EPG)

1. Парсеры могут писаться на любых языках.
2. Парсеры запусть периодически (запускалка, по-умолчанию по cron-у).
3. Для каждого сайта свой парсер.
4. У парсеров `стандартизированный вывод`.
5. Парсеры имеют логирование.
6. Если парсер не может толком распарсить, он пишет ошибку в лог и завершается с ошибкой (status > 0)

# Как собрать проект

??

# Как запустить парсер

???
 - without build: `go run . -format {format}`, format is optional one of [`csv`, `json`] default format `csv`
 - with build `go build && ./main -format {format}` _(позвеж немного изменится)_

# Список парсеров

- [Digea](https://www.digea.gr)
- [EPT](https://program.ert.gr)

# Вывод парсера

### CSV, separated by ";"

1 - строка - одна передача.

* День (дата [yyyymmdd], строка)
* Канал (название, строка)
* Лого канала (URL, строка)
* Передача (название, строка)
* Дата и время начала
* Дата и время конца
* Описание передачи (текст)

Пример (10 строк):

```csv
20200417;-KPHTH TV-;/images/channel-icons/KRHTH TV EPG.png;[K] Gormiti;2020-04-17T09:30:00+03:00;2020-04-17T10:00:00+03:00
20200417;-KPHTH TV-;/images/channel-icons/KRHTH TV EPG.png;[K] Sissi - Η Μικρή Πριγκίπισσα;2020-04-17T10:00:00+03:00;2020-04-17T10:30:00+03:00
20200417;-KPHTH TV-;/images/channel-icons/KRHTH TV EPG.png;[K8] Ακολουθία Των Μεγάλων Ωρών & Αποκαθήλωση;2020-04-17T10:30:00+03:00;2020-04-17T11:00:00+03:00
20200417;-KPHTH TV-;/images/channel-icons/KRHTH TV EPG.png;[K8] Ύμνοι Μεγάλης Εβδομάδας - Κρητική Χορωδία;2020-04-17T11:00:00+03:00;2020-04-17T14:00:00+03:00
20200417;-KPHTH TV-;/images/channel-icons/KRHTH TV EPG.png;[K] Πάνω Από Όλα  Η Υγεία Μας;2020-04-17T14:00:00+03:00;2020-04-17T14:30:00+03:00
20200417;-KPHTH TV-;/images/channel-icons/KRHTH TV EPG.png;[K8] Ιστορίες Γης;2020-04-17T14:30:00+03:00;2020-04-17T15:00:00+03:00
20200417;-KPHTH TV-;/images/channel-icons/KRHTH TV EPG.png;[K8] Πας Μαγειρεύοντας;2020-04-17T15:00:00+03:00;2020-04-17T15:30:00+03:00
20200417;-KPHTH TV-;/images/channel-icons/KRHTH TV EPG.png;[K8] Πας Μαγειρεύοντας;2020-04-17T15:30:00+03:00;2020-04-17T16:00:00+03:00
20200417;-KPHTH TV-;/images/channel-icons/KRHTH TV EPG.png;[K8] Κάνουμε Πράξη Την Αγάπη;2020-04-17T16:00:00+03:00;2020-04-17T16:30:00+03:00
20200417;-KPHTH TV-;/images/channel-icons/KRHTH TV EPG.png;[K8] Bake My Day;2020-04-17T16:30:00+03:00;2020-04-17T17:00:00+03:00
20200417;-KPHTH TV-;/images/channel-icons/KRHTH TV EPG.png;[K] Ειδήσεις - Σύντομο Δελτίο;2020-04-17T17:00:00+03:00;2020-04-17T17:30:00+03:00
20200417;-KPHTH TV-;/images/channel-icons/KRHTH TV EPG.png;[K8] Καλό Μεσημέρι;2020-04-17T17:30:00+03:00;2020-04-17T18:00:00+03:00
20200417;-KPHTH TV-;/images/channel-icons/KRHTH TV EPG.png;[K8] Biblical Mysteries;2020-04-17T18:00:00+03:00;2020-04-17T19:00:00+03:00
20200417;-KPHTH TV-;/images/channel-icons/KRHTH TV EPG.png;[K] Ταξίδια Στον Όμορφο Κόσμο;2020-04-17T19:00:00+03:00;2020-04-17T20:00:00+03:00
20200417;-KPHTH TV-;/images/channel-icons/KRHTH TV EPG.png;[K8] Η Κρήτη TV Θυμάται;2020-04-17T20:00:00+03:00;2020-04-17T20:15:00+03:00
20200417;-KPHTH TV-;/images/channel-icons/KRHTH TV EPG.png;[K8] Hit The Road;2020-04-17T20:15:00+03:00;2020-04-17T20:30:00+03:00
20200417;-KPHTH TV-;/images/channel-icons/KRHTH TV EPG.png;[K] Ειδήσεις - Κεντρικό Δελτίο;2020-04-17T20:30:00+03:00;2020-04-17T21:00:00+03:00
20200417;-KPHTH TV-;/images/channel-icons/KRHTH TV EPG.png;[K] Η Ακολουθία Tου Επιταφίου Θρήνου;2020-04-17T21:00:00+03:00;2020-04-17T22:00:00+03:00
20200417;-KPHTH TV-;/images/channel-icons/KRHTH TV EPG.png;[K] Βiblical Mysteries;2020-04-17T22:00:00+03:00;2020-04-18T01:00:00+03:00
20200417;-KPHTH TV-;/images/channel-icons/KRHTH TV EPG.png;[K] Ξένη Ταινία;2020-04-18T01:00:00+03:00;2020-04-18T02:00:00+03:00
20200417;-SYROS TV1 -;/images/channel-icons/TV 1 SYROS EPG.jpg;[K] Παιδική Ζώνη;2020-04-17T09:00:00+03:00;2020-04-17T12:00:00+03:00
20200417;-SYROS TV1 -;/images/channel-icons/TV 1 SYROS EPG.jpg;[K8] Χρώματα Ελλάδας;2020-04-17T12:00:00+03:00;2020-04-17T13:00:00+03:00
20200417;-SYROS TV1 -;/images/channel-icons/TV 1 SYROS EPG.jpg;[K8] Ίασις;2020-04-17T13:00:00+03:00;2020-04-17T14:00:00+03:00
20200417;-SYROS TV1 -;/images/channel-icons/TV 1 SYROS EPG.jpg;[K8] Ντοκιμαντέρ;2020-04-17T14:00:00+03:00;2020-04-17T15:00:00+03:00
20200417;-SYROS TV1 -;/images/channel-icons/TV 1 SYROS EPG.jpg;[K8] Άρωμα Ελλάδας;2020-04-17T15:00:00+03:00;2020-04-17T17:00:00+03:00
20200417;-SYROS TV1 -;/images/channel-icons/TV 1 SYROS EPG.jpg;[K8] Τηλεαγορές;2020-04-17T17:00:00+03:00;2020-04-17T19:00:00+03:00
20200417;-SYROS TV1 -;/images/channel-icons/TV 1 SYROS EPG.jpg;[K8] Μπροστά Από Τις Κάμερες;2020-04-17T19:00:00+03:00;2020-04-17T21:00:00+03:00
20200417;-SYROS TV1 -;/images/channel-icons/TV 1 SYROS EPG.jpg;[K8] Ζωές Παράλληλες;2020-04-17T21:00:00+03:00;2020-04-17T22:00:00+03:00
20200417;-SYROS TV1 -;/images/channel-icons/TV 1 SYROS EPG.jpg;[K8] Ειδήσεις;2020-04-17T22:00:00+03:00;2020-04-17T23:00:00+03:00
20200417;-SYROS TV1 -;/images/channel-icons/TV 1 SYROS EPG.jpg;[K8] Σεργιάνι Στην Κρήτη;2020-04-17T23:00:00+03:00;2020-04-18T00:00:00+03:00
20200417;-SYROS TV1 -;/images/channel-icons/TV 1 SYROS EPG.jpg;[K8] Ανατρεπτικό Δελτίο;2020-04-18T00:00:00+03:00;2020-04-18T01:00:00+03:00
20200417;-SYROS TV1 -;/images/channel-icons/TV 1 SYROS EPG.jpg;[K8] Κόντρα Και Ρήξη;2020-04-18T01:00:00+03:00;2020-04-18T02:20:00+03:00
20200417;-SYROS TV1 -;/images/channel-icons/TV 1 SYROS EPG.jpg;[K8] Ειδήσεις;2020-04-18T02:20:00+03:00;2020-04-17T03:20:00+03:00
20200417;-SYROS TV1 -;/images/channel-icons/TV 1 SYROS EPG.jpg;[K8] Μη Μου Πεις;2020-04-17T03:20:00+03:00;2020-04-17T04:15:00+03:00
20200417;-SYROS TV1 -;/images/channel-icons/TV 1 SYROS EPG.jpg;[K12] Ζώνη Ξένων Ταινιών;2020-04-17T04:15:00+03:00;2020-04-17T06:15:00+03:00
```

## Ядро которое собирает спарсенные данные и складывает в базу.

### Список всех зон для инициализации локальных временных зон

- [List of tz database time zones](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)
