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
            'available_geolocal'
        ], "\t"); 

        foreach($rawData as $program) {
            fputcsv($fp, $program, "\t");
        }

        fclose($fp);
    }
}