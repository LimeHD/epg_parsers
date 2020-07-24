<?php

require_once("src/ParserCards.php");
require_once("src/Help.php");

$options = getopt("o:r:h", ["output:", "resource:", "help"]);

if (isset($options["help"])) {
    Help::Print();
    exit;
}

if (isset($options["output"])) {
    $output = $options["output"];
} else {
    $output = sprintf("%s.csv",basename(dirname(__FILE__))) ;
}

$urlKeepEpgIds = $options["resource"];

$version = "v.0.0.1";
echo "Парсер stv_parser $version" . PHP_EOL . PHP_EOL;

$parser = new ParserCards($urlKeepEpgIds);

$epgsId = $parser->getAllEpgIDs();
echo "Собрал все epg_id для парсера c ресурса $urlKeepEpgIds" . PHP_EOL;
echo "Начинаю парсить телепрограммы" . PHP_EOL;
$data = $parser->readEpgFromJson($epgsId);
echo "Cпарсил телепрограммы" . PHP_EOL;

$parser->save($data, $output);
echo "Сохраняю данные в $output" .PHP_EOL;
echo "Работа парсера завершена" . PHP_EOL;