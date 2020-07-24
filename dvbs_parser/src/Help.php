<?php

Class Help
{
    /**
     * Данные для help
     */
    private static function data()
    {
        return [
            "required" => 
                [
                    [
                        "short" => "r",
                        "full"  => "resource",
                        "desc"  => "Ресурс с которого парсим список epg id",
                    ],
                    [
                        "short" => "o",
                        "full"  => "output",
                        "desc"  => "Файл в который выводим спарсенные данные"
                    ]
                ],
        ];
    }

    /** 
     * Выводит help
     */
    public static function  Print()
    {
        $params = static::data();

        foreach ($params as $type => $data) {
            if ($type == "required") {
                echo "Обязательные аргументы для длинных параметров являются обязательными и для коротких параметров." . PHP_EOL . PHP_EOL;
            }

            foreach ($data  as $args) {
                echo sprintf("-%s   --%s    %s", $args['short'], $args['full'], $args['desc']) . PHP_EOL;
            }
        }

        echo PHP_EOL;
    }
}