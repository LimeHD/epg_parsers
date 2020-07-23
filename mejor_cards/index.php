<?php

require_once("src/ParserCards.php");

$options = getopt("o:r:", ["output:resource:"]);

if (isset($options["output"])) {
    $output = $options["output"];
} else {
    $output = sprintf("mejor_%s.csv", date("Y-m-d"));
}

$urlKeepEpgIds = $output["resource"];

$parser = new ParserCards($urlKeepEpgIds);
$data = $parser->readEpgFromJson(["a01p"]);
$parser->save($data, $output);