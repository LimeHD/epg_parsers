<?php ini_set('memory_limit','2048M');

require_once 'vendor/autoload.php';
require_once 'StorageInterface.php';
require_once 'Fmt.php';
require_once 'Datamapper.php';
require_once 'Query.php';
require_once 'Storage.php';
require_once 'FileHelper.php';

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

        Fmt::info(sprintf("Not deprecated items in day: %s and channel %d", $key, $channel));
    }
}

