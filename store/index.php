<?php 
ini_set('memory_limit','3072M');

require_once 'vendor/autoload.php';
require_once 'src/StorageInterface.php';
require_once 'src/Fmt.php';
require_once 'src/Datamapper.php';
require_once 'src/Query.php';
require_once 'src/Storage.php';
require_once 'src/FileHelper.php';

Fmt::info("Основной парсер EPG на PHP");
Fmt::info("Версия 0.0.1");
Fmt::info("Выделено памяти: " . ini_get('memory_limit'));
Fmt::info("Инициализирую соединение с базой данных...");

$options    = getopt("h:l:d:p:i:g:f:c:u:");
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

$updatedIdMap = [];

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

                    $updatedIdMap[$channel] = 1;
                }
            }

            continue;
        }

        Fmt::info(sprintf("Нет устаревшей телепрограммы для даты %s и для ЕПГ id: %d", $key, $channel));
    }
}

if (count($updatedIdMap) > 0) {
    Fmt::info("Обнаружены изменения в телепрграмме, фиксирую изменения для следующих EPG...");
    Fmt::info(implode(',', array_keys($updatedIdMap)));

    $storage->setAsAffectedEpgSection(array_keys($updatedIdMap));
    $status = updateHashSum($options['u']);

    if ($status !== true) {
        Fmt::warning(sprintf("Не удалось обновить хешсумму! %s", $status));
    }

} else {
    Fmt::info("Парсер завершил работу. Нет изменений в телепрограмме");
}

function updateHashSum($url) : bool
{
    $result = json_decode(file_get_contents($url), true);

    return $result['result'];
}

