<?php ini_set('memory_limit','512M');

require_once 'wget.php';
$options = getopt("l:p:h:o::f::d::");

(function($options) {
    $wget = new Wget(new Options($options));
    $days = $wget->parse();

    $tsvStrings = [];

    foreach ($days as $key => $day) {
        foreach ($day as $channel => $programms) {
            foreach ($programms as $programm) {
                $tsvStrings[] = toTSV($channel, $programm);
            }
        }
    }

    toFile($wget->output, $tsvStrings);
})($options);

function toTSV(int $id, array $programms) : array {
    return [
        $programms['timestart'],
        $programms['timestop'],
        $id,
        $programms['title'],
        $programms['desc']
    ];
}

function toFile(string $output, array $programms) {
    $fp = fopen($output, 'w');
    $header = [
        'datetime_start', 'datetime_finish', 'channel', 'title', 'description'
    ];

    fputcsv($fp, $header, "\t");
    foreach ($programms as $programm) {
        fputcsv($fp, $programm, "\t");
    }

    fclose($fp);
}
