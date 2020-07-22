<?php

use Pixie\Connection;
use Pixie\QueryBuilder\QueryBuilderHandler;

class Query
{
    /**
     * @var QueryBuilderHandler
     */
    public $builder;

    public function __construct(array $options)
    {
        try {
            $this->builder = $this->createQueryBuilder($options);
        } catch (Exception $e) {
            Fmt::fatal($e->getMessage());

            exit;
        }
    }

    /**
     * @param array $options
     * @return \Pixie\QueryBuilder\QueryBuilderHandler
     * @throws \Pixie\Exception
     */
    public function createQueryBuilder(array $options) : QueryBuilderHandler
    {
        $config = [
            'driver'    => 'mysql',
            'host'      => $options['h'],
            'database'  => $options['d'],
            'username'  => $options['l'],
            'password'  => $options['p'],
            'charset'   => 'utf8',
            'collation' => 'utf8_unicode_ci',
            'prefix'    => '',
        ];

        $connection = new Connection('mysql', $config);
        return new QueryBuilderHandler($connection);
    }
}
