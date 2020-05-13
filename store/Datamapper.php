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
                'rating'    => 0
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
}
