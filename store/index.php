<?php ini_set('memory_limit','2048M');

require_once 'vendor/autoload.php';
require_once 'src/StorageInterface.php';
require_once 'src/Fmt.php';
require_once 'src/Datamapper.php';
require_once 'src/Query.php';
require_once 'src/Storage.php';
require_once 'src/FileHelper.php';

Fmt::info("Инициализирую соединение с базой данных...");

$options    = getopt("h:l:d:p:i:g:f:c:");
$connection = new Query($options);
$storage    = new Storage($connection);

Fmt::info("Начинаю читать файл XML...");

$days = Datamapper::readXMLtoArray(
    $options['i'],
    $options['f'] ?? '',
    $options['c'] ?? 5
);

Fmt::info("Получаю аналогичные данные из базы данных...");

$groupedDb  = $storage->getEqualItemsFor($days['items']);

Fmt::info("Начинаю процесс сравнения данных и их обновления...");

foreach ($days['items'] as $key => $day) {
    foreach ($day as $channel => $programs) {

        if (!Datamapper::isEqualMaps($programs, $groupedDb[$key][$channel])) {
            {
                Fmt::info(sprintf("Обнаружены устаревшие данные в БД для даты %s и для ЕПГ id: %d, удаляю...", $key, $channel));
                $deleteIds = $storage->deleteDay($key, $channel);
            }

            {
                $storedIds = $storage->store($programs);

                if ($storedIds && count($storedIds)) {
                    Fmt::info(sprintf("Добавляю новые данные..., количество вставленных строк составляет: %d", Datamapper::innerCount($storedIds)));
                    Fmt::info(Datamapper::implode($storedIds));
                }
            }

            continue;
        }

        Fmt::info(sprintf("Нет устаревшей телепрограммы для даты %s и для ЕПГ id: %d", $key, $channel));
    }
}

