<?php

require_once("src/ParserCards.php");

$options = getopt("o:r:", ["output:", "resource:"]);

if (isset($options["output"])) {
    $output = $options["output"];
} else {
    $output = sprintf("mejor_%s.csv", date("Y-m-d"));
}

$urlKeepEpgIds = $options["resource"];

$parser = new ParserCards($urlKeepEpgIds);
$epgsId = $parser->getAllEpgIDs();
$data = $parser->readEpgFromJson($epgsId);
$parser->save($data, $output);