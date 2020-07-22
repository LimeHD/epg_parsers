<?php


class Datamapper
{
    /**
     * @param string $file
     * @param string $from
     * @return array
     * @throws Exception
     */
    public static function readXMLtoArray(string $file, string $from = '', $countOfDays = 3) : array
    {
        if (FileHelper::isExist($file) === false) {
            throw new Exception("Файла не существует или указан неверный путь к файлу");
        }

        if (FileHelper::is($file, 'xml') === false) {
            throw new Exception("Кажется вы пытаетесь загрузить недопустимый файл, пожалуйста выберите файл с расширением .xml");
        }

        $xml = simplexml_load_string(file_get_contents($file));
        $today = date('Y-m-d');
        $loopStartDay = 0;

        if (static::isValisDate($from)) {
            Fmt::info(sprintf('Указана дата %s начинаю сканировать с этой даты', $from));
            $today =  date('Y-m-d', strtotime($from));

            $loopStartDay = static::countDaysBetweenDates(date('Y-m-d'),  $today);
        }

        $datastructure = [
            'count' => 0,
            'items' => []
        ];

        foreach ($xml->programme as $row) {
            $attr = $row->attributes();

            $id = (string)$attr['channel'];

            $start = DateTime::createFromFormat('YmdHis O', (string)$attr['start']);
            $stop = DateTime::createFromFormat('YmdHis O', (string)$attr['stop']);
            // приведение к MSK для нового поставщика
            $start->setTimezone(new \DateTimeZone('Europe/Moscow'));
            $stop->setTimezone(new \DateTimeZone('Europe/Moscow'));

            $subTitle = (isset($row->{'sub-title'})) ? '. ' . (string)$row->{'sub-title'} : '';
            $title = sprintf('%s%s', (string) $row->title, $subTitle);
            $desc = (string)$row->desc;
            $rating = isset($row->rating) && isset($row->rating->value) ? (int)$row->rating->value : null;

            $programDate = $start->format('Y-m-d');

            if ($programDate >= $today) {
                $timeStart = $start->format('Y-m-d H:i:s');
                $timeStop = $stop->format('Y-m-d H:i:s');

                $day = $start->format('Y-m-d');

                if (!isset($datastructure['items'][$day])) {
                    $datastructure['items'][$day] = [];
                }

                if (!isset($datastructure['items'][$day][$id])) {
                    $datastructure['items'][$day][$id] = [];
                }

                $datastructure['items'][$day][$id][] = [
                    'epg_id'        => $id,
                    'time_zone'     => sprintf("UTC%s", $start->format('P')),
                    'timestart'     => $timeStart,
                    'date'          => $day,
                    'timestop'      => $timeStop,
                    'title'         => $title,
                    'desc'          => $desc,
                    'rating'        => $rating,
                ];
                $datastructure['count']++;
            }
        }

        unset($xml);

        $cut = [
            'items' => []
        ];

        for ($i = $loopStartDay; $i <= $countOfDays; $i++) {
            $day = date("Y-m-d", strtotime("+" . $i . " day"));

            if (!isset($cut['items'][$day])) {
                $cut['items'][$day] = [];
            }

            $cut['items'][$day] = $datastructure['items'][$day];
        }

        unset($datastructure);

        return $cut;
    }

    /**
     * @param string $file
     * @return array
     * @throws Exception
     */
    public static function readTSVtoArray(string $file) : array
    {
        if (FileHelper::isExist($file) === false) {
            throw new Exception("Файла не существует или указан неверный путь к файлу");
        }

        if (FileHelper::is($file, 'xml') === false) {
            throw new Exception("Кажется вы пытаетесь загрузить недопустимый файл, пожалуйста выберите файл с расширением .xml");
        }

        $fp = fopen($file, 'r');

        $datastructure = [
            'count' => 0,
            'items' => []
        ];

        while (!feof($fp))
        {
            $line = fgets($fp, 2048);
            $item = str_getcsv($line, "\t");

            $start = DateTime::createFromFormat(DATE_RFC3339, $item[0]);
            $stop = DateTime::createFromFormat(DATE_RFC3339, $item[1]);

            if (!$start || !$stop) {
                Fmt::warning("Invalid datetime");

                continue;
            }

            $datastructure['items'][] = [
                'epg_id'    => 0,
                'time_zone' => sprintf("UTC%s", $start->getTimezone()->getName()),
                'timestart' => $start->format('Y-m-d H:i:s'),
                'timestop'  => $stop->format('Y-m-d H:i:s'),
                'title'     => $item[3],
                'desc'      => $item[5],
                'date'      => $start->format('Y-m-d'),
            ];
            $datastructure['count']++;
        }

        fclose($fp);
        return $datastructure;
    }

    private static function isValisDate(string $date, $format = 'Y-m-d') : bool
    {
        $d = DateTime::createFromFormat($format, $date);
        return $d && $d->format($format) === $date;
    }

    private static function countDaysBetweenDates(string $a, string $b) : int
    {
        $timeDiff = strtotime($b) - strtotime($a);
        return $timeDiff / 86400;
    }

    /**
     * @param array $items
     * @return array
     */
    public static function batches(array $items) : array
    {
        return array_chunk(array_values($items), 10);
    }

    /**
     * @param array $items
     * @return string
     */
    public static function implode(array $items) : string
    {
        $string = 'Stored identifiers in database:';

        foreach ($items as $k => $item) {
            $string .= PHP_EOL . "{$k} batch: " . implode(',', $item);
        }

        return $string;
    }

    /**
     * @param array $items
     * @return array
     */
    public static function groupByDay(array $items) : array
    {
        $grouped = [];

        foreach ($items as $item) {
            // чтобы не конвертировать по 100500 раз просто выдергиваем дату из даты и времени
            $day = explode(' ', $item['timestart'])[0];
            $grouped[$day][] = $item;
        }

        return $grouped;
    }

    /**
     * @param array $a
     * @param array $b
     * @return bool
     */
    public static function isEqualMaps(array $a, array $b) : bool
    {
        $diff = count($a) != count($b);

        if (!$diff) {
            // simple deep check is equal array
            foreach ($a as $k => $v) {
                if (count(array_diff($v, $b[$k]))) {
                    $diff = true;
                    break;
                }
            }
        }

        return !$diff;
    }

    /**
     * @param array $batches
     * @return int
     */
    public static function innerCount(array $batches) : int
    {
        $counter = 0;

        foreach ($batches as $batch) {
            $counter += count($batch);
        }

        return $counter;
    }
}