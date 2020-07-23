<?php

require_once("vendor/autoload.php");
require_once("src/Parser.php");
require_once("src/RepositoryTSV.php");

$options = getopt("f:h:",["format:","output:"]);

$outputFile = $options["output"];

$parser = new Parser($outputFile);
$epgData = $parser->parserXML();

if ($options["format"] == "csv") {
    $repository = new RepositoryTSV("stv_epg.xml");
    $repository->save($epgData);
}

