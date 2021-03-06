<?php

use \XMLDocument as XMLDocument;
use \DOMDocument as DOMDocument;

class Parser
{
    /**
     * @var XMLReader
     */
    private $xmlReader;

    /**
     * @var DOMDocument
     */
    private $dom;

    /**
     * @var string
     */
    private $xmlFile;

    public function __construct($filePath) {
        $this->xmlFile = $filePath;
        $this->xmlReader = new XMLReader;
        $this->dom = new DOMDocument;
    }

    /**
     * @param string $xmlFile Путь к xml файлу
     * @return array
     */
    public function parserXML() :array
    {
        $program = [];
        
        $this->xmlReader->open($this->xmlFile);

        while ($this->xmlReader->read()) {
            if ($this->xmlReader->nodeType == XMLReader::ELEMENT && $this->xmlReader->name == 'programme') {
                $list = $this->domNodeAsArray($this->xmlReader);
                
                $start = DateTime::createFromFormat('YmdHis O', (string)$list['start']);
                $stop = DateTime::createFromFormat('YmdHis O', (string)$list['stop']);

                $programDate = $start->format('Y-m-d');
                $start = $start->setTimezone(new \DateTimeZone('Europe/Madrid'));
                $stop = $stop->setTimezone(new \DateTimeZone('Europe/Madrid'));
                $timeStart = $start->format(DateTime::RFC3339_EXTENDED); 
                $timeStop = $stop->format(DateTime::RFC3339_EXTENDED);
                $channel = $list['channel'];
                
                $program[] = [
                    'channel'           => $channel,
                    'timestart'         => $timeStart,
                    'timestop'          => $timeStop,                    
                    'titleEs'           => $this->replaceKyrilicTextToNull($list['titleEs']),
                    'titleCa'           => $this->replaceKyrilicTextToNull($list['titleCa']),
                    'descriptionEs'     => $this->replaceKyrilicTextToNull($list['descriptionEs']),
                    'descriptionCa'     => $this->replaceKyrilicTextToNull($list['descriptionCa']),
                ];
            }
        }

        return $program;
    }

    /**
     * Строим новое дерево из ноды 
     * 
     * @param XMLReader $reader
     * @return array
     */
    private function domNodeAsArray(XMLReader $reader): array
    {
        $list = [];
        $genre = [];
        $dom = simplexml_import_dom($this->dom->importNode($reader->expand(), true));
        $attributes = $dom->attributes();
        
        $list['channel'] = (string)$attributes['channel'];
        $list['start'] = (string)$attributes['start'];
        $list['stop'] = (string)$attributes['stop'];
        $list['titleEs'] = (string)$dom->title;
        $list['titleCa'] = (string)$dom->title;
        $list['descriptionEs'] = (string)$dom->desc;
        $list['descriptionCa'] = (string)$dom->desc;
        
        return $list;
    }

    /**
     * Ловим кириллицу
     * 
     * @param string $str
     * @return string|null
     */
    public function replaceKyrilicTextToNull(string $str): ?string
    {   
        preg_match('/(*UTF8)[а-яА-Я]/i', $str, $matches);

        if (isset($matches[0])) {
            return null;
        }
        
        return $str;
    }

}
