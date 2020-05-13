<?php

require_once 'vendor/autoload.php';
require_once 'StorageInterface.php';
require_once 'Fmt.php';
require_once 'Datamapper.php';
require_once 'Query.php';
require_once 'Storage.php';

$options    = getopt("h:l:d:p:i:g");
$connection = new Query($options);
$storage    = new Storage($connection);

$common = Datamapper::readTSVtoArray($options['i']);
$storedIds = $storage->store($common['items']);

$count = $common['count'];
Fmt::info("Всего строк обработано: {$count}");
Fmt::info(Datamapper::implode($storedIds));



