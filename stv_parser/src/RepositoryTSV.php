<?php 

require_once("RepositoryInterface.php");

/**
 * Класс для для работы с репозиторием типа TSV
 */
class RepositoryTSV implements RepositoryInterface
{
    private $fileName;

    public function __construct($fileName = null)
    {
        if (is_null($fileName)) {
            $this->fileName = sprintf("mejor_%s",  date('Y_m_d'));
            return;
        }
        
        $this->fileName = $fileName;
    }


    /**
     * Сохраняем данные epg в tsv
     */
    public function save($rawData) :void
    {
        $fp = fopen($this->fileName, 'w+');

        $head = ['datetime_start', 'datetime_finish', 'channel', 'title','channel_logo_url', 'description'];
        fputcsv($fp, $head, "\t"); 

        foreach ($rawData as $line) {
            $fields = [
                $line['timestart'],
                $line['timestop'],
                $line['channel'],
                $line['titleEs'],
                '',
                $line['descriptionEs']
            ];

            fputcsv($fp, $fields, "\t");
        }

        fclose($fp);
    }
    
}