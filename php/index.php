<?php ini_set('memory_limit','512M');

require_once 'wget.php';
$options = getopt("l:p:");

(function($options) {
    $wget = new Wget($options['l'], $options['p']);
    $days = $wget->parse();

    $tsvStrings = [];

    foreach ($days as $key => $day) {
        foreach ($day as $channel => $programms) {
            foreach ($programms as $programm) {
                $tsvStrings[] = toTSV($channel, $programm);
            }
        }
    }

    toFile($tsvStrings);
})($options);

function toTSV(int $id, array $programms) : string {
    return sprintf("%s\t%s\t%d\t%s\t%s",
        $programms['timestart'],
        $programms['timestop'],
        $id,
        $programms['title'],
        $programms['desc']
    );
}

function toFile(array $programms) {
    file_put_contents('export.csv', "datetime_start\tdatetime_finish\tchannel\ttitle\tdescription\n");
    file_put_contents('export.csv', implode(PHP_EOL, $programms), FILE_APPEND);
}
