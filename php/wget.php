<?php

class Wget
{
    private $url = 'ftp.epgservice.ru';
    private $login = '';
    private $password = '';
    private $filename = 'TV_Pack.xml';

    public function __construct(string $login, string $pass)
    {
        $this->login = $login;
        $this->password = $pass;
    }

    /**
     * @throws \Exception
     */
    public function parse()
    {
        $uploadDir = 'upload';
        $filename = date("Ymd_H") . '.xml';
        $today = date('Y-m-d');

        if (!file_exists($uploadDir)) {
            mkdir($uploadDir, 0777);
        }

        $filePath = sprintf('%s/%s', $uploadDir, $filename);
        if (file_exists($filePath)) {
            $xmlFile = file_get_contents($filePath);

            line(sprintf('Actual file: %s is exist, continue working with him', $filename));
        } else {

            line(sprintf('Actual file not found. Starting download file: %s', $filename), 0);
            line(sprintf(' --start: %s', date('H:i:s', time())), 0);

            $xmlFile = file_get_contents('ftp://' . $this->login . ':' . $this->password . '@' . $this->url . '/' . $this->filename);
            @chmod($uploadDir . '/' . $filename, 0777);
            file_put_contents($uploadDir . '/' . $filename, $xmlFile);

            line(sprintf(' --end: %s', date('H:i:s', time())));
        }

        line(sprintf('Load file in: %s', date('H:i:s', time())), 0);
        $xml = simplexml_load_string($xmlFile);
        echo sprintf(' --loaded in: %s', date('H:i:s', time())), PHP_EOL;

        $program = [];

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

            if (!isset($program[$programDate])) $program[$programDate] = [];
            if (!isset($program[$programDate][$id])) $program[$programDate][$id] = [];

            if ($programDate >= $today) {
                $timeStart = $start->format('Y-m-d H:i:s');
                $timeStop = $stop->format('Y-m-d H:i:s');

                $program[$programDate][$id][] = [
                    'timestart' => $timeStart,
                    'timestop'  => $timeStop,
                    'title'     => $title,
                    'desc'      => $desc,
                    'rating'    => $rating,
                ];
            }
        }

        unset($xmlFile);
        unset ($xml);

        line(sprintf('Complete parsing XML in: %s', date('H:i:s', time())));
        return $program;
    }

    private function zeropad($num, $lim)
    {
        return (strlen($num) >= $lim) ? $num : $this->zeropad("0" . $num, $lim);
    }
}

function line($content, $countBreaks = 1)
{
    for ($i = 1; $i <= $countBreaks; ++$i) {
        $content .= PHP_EOL;
    }

    echo ' [PARSER] --> ', $content;
}
