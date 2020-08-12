<?php

require_once("src/Parser.php");
require_once("src/RepositoryTSV.php");
require_once("src/ArgHelper.php");

echo "Graella парсер v0.0.2" . PHP_EOL . PHP_EOL;
$options = getopt('f:o:h',['filesdir:', 'output:', 'help']);
list($files, $outputFile) = ArgHelper::argChecker($options);
$chlist = ['324', 'TV3', 'ES3'];

echo "Получаю список файлов" . PHP_EOL;
$filesInDir = new FilesystemIterator($files, FilesystemIterator::KEY_AS_FILENAME | FilesystemIterator::SKIP_DOTS);

$parser = new Parser();

if ($filesInDir->getType() != 'file') {
    echo "Не удалось найти файлы в папке $files" . PHP_EOL;
    exit;
}

echo "Начинаю парсить файлы" . PHP_EOL;
foreach ($filesInDir as $file) {
    $fileData = explode('_', $file);
    $ch = $fileData[1];

    if (!$file->isReadable()) {
        echo "Файл $file недоступен на чтение";
        exit;
    }

    if (sizeof($chlist) > 0) {
        if (!in_array($ch, $chlist)) {
            continue;
        }
    }
    $parser->parseXML($file);
}

$data = $parser->getResult();

if (sizeof($data) == 0) {
    echo "Не удалось спарсить файлы. Результирующий массив пустой" . PHP_EOL;
    exit;
}

echo "Сохраняю в CSV $outputFile" . PHP_EOL;
$repository = new RepositoryTSV($outputFile);
$repository->save($data);

echo "Работа парсера завершена" . PHP_EOL;
