<?php

require_once("src/ParserCards.php");

$options = getopt("o:r:", ["output:", "resource:"]);

if (isset($options["output"])) {
    $output = $options["output"];
} else {
    $output = sprintf("mejor_%s.csv", date("Y-m-d"));
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