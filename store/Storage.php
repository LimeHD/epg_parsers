<?php

use Pixie\QueryBuilder\QueryBuilderHandler;

class Storage implements StorageInterface
{
    /**
     * @var QueryBuilderHandler
     */
    public $storageConnection;

    public function __construct(Query $builder)
    {
        $this->storageConnection = $builder->builder;
    }

    /**
     * @return QueryBuilderHandler
     */
    public function getBroadcasterStorage() : QueryBuilderHandler
    {
        return $this->storageConnection->table('broadcasters');
    }

    /**
     * @param array $items
     * @return array
     */
    public function store(array $items) : array
    {
        $insertIds = [];

        foreach (Datamapper::batches($items) as $batch) {
            $insertIds[] = $this->getBroadcasterStorage()->insert($batch);
        }

        return $insertIds;
    }
}
