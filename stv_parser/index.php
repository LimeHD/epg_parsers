<?php

require_once("vendor/autoload.php");
require_once("src/Parser.php");
require_once("src/RepositoryTSV.php");
require_once("src/Help.php");

$version = "v0.0.1";

$options = getopt("fmt:f:o:h",["format:","file:","output:","help"]);

if (isset($options["help"])) {
    Help::Print();
    exit;
}

$outputFile = $options["output"];
$xmlFile = $options["file"];

echo "Парсер stv_parser $version" . PHP_EOL;

if (!file_exists($xmlFile)) {
    echo "Не могу найти файл $xmlFile" . PHP_EOL;
    exit;
}

echo "Начинаю парсить файл $xmlFile" . PHP_EOL;
$parser = new Parser($xmlFile);
$epgData = $parser->parserXML();

if ($options["format"] == "csv") {
    echo "Сохраняю данные в $outputFile" . PHP_EOL;
    $repository = new RepositoryTSV($outputFile);
    $repository->save($epgData);
}
 echo "Работа парсера завершена" . PHP_EOL;
