<?php

require_once("Help.php");

class ArgHelper
{
    /**
     * Проверяет наличие и корректность параметров
     * @var array Ассоциативный массив параметров
     * @return array
     */
    public static function argChecker(array $options) :array
    {
        if (isset($options['help'])) {
            Help::Print();
            exit;
        }

        if (!isset($options['filesdir'])) {
            echo "Аргумент filesdir пустой или отсутствует" . PHP_EOL;
            exit;
        }

        if (!file_exists($options['filesdir'])) {
            echo "Директории '{$options['filesdir']}/' не существует или не была передана в аргументах" . PHP_EOL;
            exit;
        }

        if (!isset($options["output"])) {
            $options["output"] = sprintf("%s.csv",basename(dirname(__FILE__))) ;
        }
        $res = array_values($options);

        return $res;
    }
}