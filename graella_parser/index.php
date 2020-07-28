<?php

require_once("src/Parser.php");
require_once("src/RepositoryTSV.php");

echo "Graella парсер v0.0.1" . PHP_EOL . PHP_EOL;
$options = getopt('f:o:',['filesdir:', 'output:']);
if (!isset($options['filesdir'])) {
    echo "Аргумент filesdir пустой или отсутствует" . PHP_EOL;
    exit;
}

$files = $options['filesdir'];

if (!file_exists($files)) {
    echo 'Директории $files не существует' . PHP_EOL;
    exit;
}

$outputFile = $options["output"];
echo "Получаю список файлов" . PHP_EOL;
$filesInDir = new FilesystemIterator($files, FilesystemIterator::KEY_AS_FILENAME | FilesystemIterator::SKIP_DOTS);

$parser = new Parser();

echo "Начинаю парсить файлы" . PHP_EOL;
foreach ($filesInDir as $file) {
    $parser->parseXML($file);
}

$data = $parser->getResult();
echo "Сохраняю в CSV" . PHP_EOL;
$repository = new RepositoryTSV($outputFile);
$repository->save($data);

echo "Работа парсера завершена" . PHP_EOL;
