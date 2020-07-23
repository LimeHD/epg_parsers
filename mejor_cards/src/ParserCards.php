<?php

class ParserCards
{
    private $urlKeepEpgIds;

    private $spanishChannelList = [];

    public function __construct($urlKeepEpgIds)
    {
        $this->urlKeepEpgIds = $urlKeepEpgIds;
    }

    /**
     * Парсим телепрограмму и возвращаем его ввиде массива
     * @return array
     */
    public function readEpgFromJson($channels = []) :array
    {
        $program = [];

        if (sizeof($this->spanishChannelList) == 0) {
            $this->spanishChannelList = $channels;
        }
        
        foreach ($this->spanishChannelList as $channelId) {
            $data = file_get_contents('http://hls.mejor.tv:81/api/v1/epg?epg=' . $channelId);
            $data = substr_replace($data, '', -4, 1);
            
            $array = json_decode($data, true);

            if (!isset($array['items']) || sizeof($array['items']) == 0) {
                continue;
            }

            $program[] = $array;
        }

        return $program;
    }

    /**
     * Сохраняем в csv
     * @return void
     */
    public function save($program, $output) :void
    {
        $fp = fopen($output, "w+");
        $fields = [];

        foreach($program[0]["items"] as $line) {

            $title = (sizeof($line['title']) > 0) ? $line['title']['ca'] : '';
            $desc = (sizeof($line['desc']) > 0) ? $line['desc']['ca'] : '';
            $start = DateTime::createFromFormat('U', (string)$line['start_ut']);
            $stop = DateTime::createFromFormat('U', (string)$line['stop_ut']);

            $fields = [
                $start->format(DateTime::RFC3339_EXTENDED),
                $stop->format(DateTime::RFC3339_EXTENDED),
                $line['channel'],
                $title,
                '',
                $desc
            ];
            fputcsv($fp, $fields, "\t");
        }

        fclose($fp);
    }

    /**
     * Получаем список всех epg_id
     * @return array 
     */
    public function getAllEpgIDs() :array
    {
        $epgIDs = [];

        $data = file_get_contents($this->urlKeepEpgIds);
        $dataInArray = json_decode($data);
        foreach ($dataInArray->channels  as $channel) {
            $epgIDs[] = $channel->epg_id;
        }

        return $epgIDs;
    }
}