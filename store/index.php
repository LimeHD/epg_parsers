<?php ini_set('memory_limit','2048M');

require_once 'vendor/autoload.php';
require_once 'src/StorageInterface.php';
require_once 'src/Fmt.php';
require_once 'src/Datamapper.php';
require_once 'src/Query.php';
require_once 'src/Storage.php';
require_once 'src/FileHelper.php';

$options    = getopt("h:l:d:p:i:g");
$connection = new Query($options);
$storage    = new Storage($connection);

$days       = Datamapper::readXMLtoArray($options['i']);
$groupedDb  = $storage->getEqualItemsFor($days['items']);

foreach ($days['items'] as $key => $day) {
    foreach ($day as $channel => $programs) {

        if (!Datamapper::isEqualMaps($programs, $groupedDb[$key][$channel])) {
            {
                Fmt::info("Delete deprecated epg...");
                $deleteIds = $storage->deleteDay($key, $channel);
            }

            {
                $storedIds = $storage->store($programs);
                Fmt::info(sprintf("Total rows inserted to database: %d", Datamapper::innerCount($storedIds)));
                Fmt::info(Datamapper::implode($storedIds));
            }

            continue;
        }

        Fmt::info(sprintf("Not deprecated items in day: %s and broadcast id: %d", $key, $channel));
    }
}

