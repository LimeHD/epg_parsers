<?php

class DownloadCredentials
{
    private $url = null;
    private $login = null;
    private $password = null;
    private $filename = 'TV_Pack.xml';

    public function __construct(string $url, string $login, string $pass, string $filename)
    {
        $this->url = $url;
        $this->login = $login;
        $this->password = $pass;
        $this->filename = $filename;
    }

    public function getUrl() : string
    {
        return $this->url;
    }

    public function getLogin() : string
    {
        return $this->login;
    }

    public function getPassword() : string
    {
        return $this->password;
    }

    public function getFilename() : string
    {
        return $this->filename;
    }
}