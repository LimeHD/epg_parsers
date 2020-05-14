<?php


class Datamapper
{
    /**
     * @param string $file
     * @return array
     */
    public static function readTSVtoArray(string $file) : array
    {
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
                'start_at'  => $start->format('Y-m-d H:i:s'),
                'finish_at' => $stop->format('Y-m-d H:i:s'),
                'title'     => $item[3],
                'detail'    => $item[5],
                'rating'    => null
            ];
            $datastructure['count']++;
        }

        fclose($fp);
        return $datastructure;
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
            $day = explode(' ', $item['start_at'])[0];
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
