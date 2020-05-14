<?php

class Fmt
{
    public static function info(string $message) : void
    {
        echo sprintf("INFO %s", $message), PHP_EOL;
    }

    public static function warning(string $message) : void
    {
        echo sprintf("WARNING %s", $message), PHP_EOL;
    }

    public static function fatal(string $message) : void
    {
        echo sprintf("ERROR %s", $message), PHP_EOL;
    }
}
