<?php

require_once("src/Parser.php");
require_once("src/RepositoryTSV.php");
require_once("src/ArgHelper.php");

echo "Graella парсер v0.0.1" . PHP_EOL . PHP_EOL;
$options = getopt('f:o:',['filesdir:', 'output:']);
list($files, $outputFile) = ArgHelper::argChecker($options);

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
