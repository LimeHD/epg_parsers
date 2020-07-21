<?php

class ParserCards
{

    private $spanishChannelList = [];

    public function readEpgFromJson($channels = [])
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

    public function save($program, $output)
    {
        $fp = fopen($output, "a+");
        $fields = [];

        foreach($program[0]["items"] as $line) {

            $title = (sizeof($line['title']) > 0) ? $line['title']['ca'] : '';
            $desc = (sizeof($line['desc']) > 0) ? $line['desc']['ca'] : '';

            $fields = [
                $line['start_ut'],
                $line['stop_ut'],
                $line['channel'],
                $title,
                '',
                $desc
            ];
            fputcsv($fp, $fields, "\t");
        }

        fclose($fp);
    }
}