<?php

require_once("src/ParserCards.php");

$parser = new ParserCards();
$data = $parser->readEpgFromJson(['a01p']);
$parser->save($data);