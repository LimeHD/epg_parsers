<?php


class FileHelper
{
    public static function isExist(string $file) : bool
    {
        return file_exists($file);
    }

    public static function is(string $file, string $extention) : bool
    {
        $info = pathinfo($file);
        return $info['extension'] === $extention;
    }
}
