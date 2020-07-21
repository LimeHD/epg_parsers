<?php

require_once("src/ParserCards.php");

$options = getopt("o:", ["output:"]);

if (isset($options['output'])) {
    $output = $options["output"];
} else {
    $output = sprintf("mejor_%s.csv", date("Y-m-d"));
}

$parser = new ParserCards();
$data = $parser->readEpgFromJson(['a01p']);
$parser->save($data, $output);