<?php

class Options
{
    public $h;
    public $l;
    public $p;
    public $o;
    public $f;
    public $d;

    /**
     * Options constructor.
     * @param array $options
     * @throws Exception
     */
    public function __construct(array $options)
    {
        if (!isset($options['h']) || !isset($options['l']) || !isset($options['p'])) {
            throw new \Exception('Вы забыли указать логин, пароль или хост. -l {login} -p {password} -h {ftp_host}');
        }

        $this->h = $options['h'];
        $this->l = $options['l'];
        $this->p = $options['p'];

        $root = dirname(dirname(__FILE__));
        $this->o = sprintf("%s/tmp/output/export.csv", $root);
        $this->d = sprintf("%s/tmp/upload", $root);
        $this->f = 'TV_Pack.xml';

        if (isset($options['o'])) {
            $this->o = $options['o'];
        }

        // todo: не понятно почему не работает...
        /*if (isset($options['f'])) {
            $this->f = $options['f'];
        }*/

        if (isset($options['d'])) {
            $this->d = $options['d'];
        }
    }
}

class Wget
{
    private $url = '';
    private $login;
    private $password = '';
    private $filename;
    private $upload;
    public $output;

    public function __construct(Options $options)
    {
        $this->login = $options->l;
        $this->password = $options->p;
        $this->url = $options->h;
        $this->filename = $options->f;
        $this->upload = $options->d;
        $this->output = $options->o;
    }

    /**
     * @throws \Exception
     */
    public function parse()
    {
        $filename = date("Ymd_H") . '.xml';
        $today = date('Y-m-d');

        if (!file_exists($this->upload)) {
            mkdir($this->upload, 0777);
        }

        $filePath = sprintf('%s/%s', $this->upload, $filename);
        if (file_exists($filePath)) {
            $xmlFile = file_get_contents($filePath);

            line(sprintf('Actual file: %s is exist, continue working with him', $filename));
        } else {

            line(sprintf('Actual file not found. Starting download file: %s', $filename), 0);
            line(sprintf(' --start: %s', date('H:i:s', time())), 0);

            $xmlFile = file_get_contents('ftp://' . $this->login . ':' . $this->password . '@' . $this->url . '/' . $this->filename);
            @chmod($this->upload . '/' . $filename, 0777);
            file_put_contents($this->upload . '/' . $filename, $xmlFile);

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
