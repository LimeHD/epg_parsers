<?php

require_once 'vendor/autoload.php';
require_once 'Fmt.php';
require_once 'Datamapper.php';
require_once 'Query.php';

$options = getopt("h:l:d:p:i:g");
$connection = new Query($options);

$common = Datamapper::readTSVtoArray($options['i']);
$connection->store($common['items']);

