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

function toTSV(int $id, array $programms) : string {
    return sprintf("%s\t%s\t%d\t%s\t%s",
        $programms['timestart'],
        $programms['timestop'],
        $id,
        $programms['title'],
        $programms['desc']
    );
}

function toFile(string $output, array $programms) {
    file_put_contents($output, "datetime_start\tdatetime_finish\tchannel\ttitle\tdescription\n");
    file_put_contents($output, implode(PHP_EOL, $programms), FILE_APPEND);
}
