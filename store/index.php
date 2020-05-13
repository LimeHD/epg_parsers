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
$grouped = Datamapper::groupByDay($common['items']);
$groupedDb = $storage->getEqualItemsFor($grouped);

foreach ($grouped as $day => $group) {
    if (!Datamapper::isEqualMaps($grouped[$day], $groupedDb[$day])) {
        Fmt::info(sprintf('Different program in day: %s', $day));

        // todo delete deprecated tv program
        {
            Fmt::info("Delete deprecated epg...");
            $deleteIds = $storage->deleteDay($day);
        }

        // todo save new program
        {
            $storedIds = $storage->store($grouped[$day]);
            $count = $common['count'];
            Fmt::info("Total rows read from file: {$count}");
            Fmt::info(sprintf("Total rows inserted to database: %d", count($storedIds)));
            Fmt::info(Datamapper::implode($storedIds));
        }
    }
}

