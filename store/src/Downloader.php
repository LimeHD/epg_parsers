<?php

class Downloader
{
    public static function download(DownloadCredentials $credentials, bool $dev = false) : void
    {
        if (FileHelper::isExist($credentials->getFilename())) {
            if ($dev) {
                echo PHP_EOL, 'drop old XML file';
            }

            unlink($credentials->getFilename());
        }

        if ($dev) {
            echo PHP_EOL, 'run download XML file';
        }

        $xmlFile = file_get_contents(static::makeFtpLink([
            $credentials->getLogin(),
            $credentials->getPassword(),
            $credentials->getUrl(),
            $credentials->getFilename()
        ]));
        @chmod($credentials->getFilename(), 0777);
        file_put_contents($credentials->getFilename(), $xmlFile);
    }

    private static function makeFtpLink(array $opts) : string
    {
        return sprintf("ftp://%s:%s@%s/%s", ...$opts);
    }
}