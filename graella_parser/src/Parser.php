<?php

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
     * Собираем все телепрограммы
     */
    private $programs = [];

    public function __construct() {
        $this->xmlReader = new XMLReader;
        $this->dom = new DOMDocument;
    }

    /**
     * Парсит xml файл и возвращает массив
     * @param string Путь к файлу xml
     * @return void
     */
    public function parseXML(string $filePath) :void
    {
        $openedXmlFile = $this->xmlReader->open($filePath);
        if (!$openedXmlFile) {
            echo "XmlReader: не удалось открыть файл $filePath" . PHP_EOL;
        }
        
        while ($this->xmlReader->read()) {
            $fields = [];

            if ($this->xmlReader->name == 'Graella') {
                $channel  = $this->xmlReader->getAttribute('Canal');
            }

            if ($this->xmlReader->name == 'Data') {
                $nowDate  = $this->xmlReader->getAttribute('Dia');
            }

            if ($this->xmlReader->nodeType == XMLReader::ELEMENT && $this->xmlReader->name == 'Programa') {
                $dom = simplexml_import_dom($this->dom->importNode($this->xmlReader->expand(), true));
                
                $startRawFormat = DateTime::createFromFormat('d/m/Y H:i:s', $nowDate . ' ' . $dom->HIni);
                $start = $startRawFormat->format(DateTime::RFC3339_EXTENDED);

                $stopRawFormat = DateTime::createFromFormat('d/m/Y H:i:s', $nowDate . ' ' . $dom->HFi);
                $stop = $stopRawFormat->format(DateTime::RFC3339_EXTENDED);

                $title = str_replace('"',"", (string)$dom->TitProg[0]);
                $desc = str_replace(["\n", '"'], [" ", ""], (string)$dom->Sinopsi[0]);
                
                $derecho = $this->checkArchive((string)$dom->Publicacion->Derecho);
                

                $geolocalizacion = $this->setGeoRegions((string)$dom->Publicacion->Geolocalizacion);
                
                $fields = [
                    'start' => $start,
                    'stop'  => $stop,
                    'channel' => $channel,
                    'title' => $title,
                    'logo'  => '',
                    'desc'  => $desc,
                    'available_archive' => $derecho,
                    'geo_regions' => $geolocalizacion
                ];
    
                $this->programs[] = $fields;
            }
        }
    }

    /**
     * Получаем данные из @var program , который является результатом работы парсера
     * @return array
     */
    public function getResult() :array
    {
        return $this->programs;
    }

    /**
     * Проверяем и приводим значение к булевому из каталанского si/no
     * Ну а если ничего нет, то возвращаем 0
     * 
     * @var string
     * @return int
     */
    private function checkArchive($resolution) :int
    {
        if ($resolution == "") {
            return 0;
        }

        if ($resolution == "NO") {
            return 0;
        }

        return 1;
    }

    /**
     * Указываем в каких регионах доступна трансляция
     * Si -- по всей Испании, No -- только по Каталонии
     * Указываем конкретно регион , а не число, для совместимости с другими парсерами
     * @var string
     * @return string
     */
    private function setGeoRegions($available) :string
    {
        if ($available == "") {
            return 'ES,CT';
        }

        if ($available == "NO") {
            return 'CT';
        }

        return "ES,CT";
    }
}
