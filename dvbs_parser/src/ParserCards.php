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

            $program[] = $array['items'];

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
        $head = ['datetime_start', 'datetime_finish', 'channel', 'title','channel_logo_url', 'description'];
        fputcsv($fp, $head, "\t"); 
        
        foreach($program as $lines) { 
            foreach ($lines as $line)  {
                $title = '';
                $desc = '';

                $esTitle = isset($line['title']['es']) ? $line['title']['es'] : null;
                $caTitle = isset($line['title']['ca']) ? $line['title']['ca'] : null;

                if(!is_null($esTitle)) {
                    $title = trim($esTitle);
                } elseif (!is_null($caTitle)) {
                    $title = trim($esTitle);
                }
                
                $esDesc = isset($line['desc']['es']) ? $line['desc']['es'] : null;
                $caDesc = isset($line['desc']['ca']) ? $line['desc']['ca'] : null;

                if(!is_null($esDesc)) {
                    $desc = trim($esDesc);
                } elseif (!is_null($caDesc)) {
                    $desc = trim($caDesc);
                }

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