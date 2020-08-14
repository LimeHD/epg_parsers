<?php

require_once("stv_parser/src/RepositoryInterface.php");

class RepositoryTSV implements RepositoryInterface
{
    private $fileName;

    public function __construct(string $fileName)
    {
        $this->fileName = $fileName;
    }

    public function save($rawData) :void
    {
        $fp = fopen($this->fileName, 'w+');

        fputcsv($fp, [
            'datetime_start',
            'datetime_finish',
            'channel',
            'title',
            'channel_logo_url',
            'description',
            'available_archive',
            'geo_regions'
        ], "\t"); 

        foreach($rawData as $program) {
            $writable = fputcsv($fp, $program, "\t");
            if ($writable == false) {
                echo "Ошибка записи строки в файл. " . json_encode($program) . PHP_EOL;
                exit;
            }
        }

        fclose($fp);
    }
}