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
        return $this->builder()->table('broadcasters');
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

    /**
     * @return QueryBuilderHandler
     */
    public function find() : QueryBuilderHandler
    {
        return $this->getBroadcasterStorage();
    }

    /**
     * @return QueryBuilderHandler
     */
    public function builder() : QueryBuilderHandler
    {
        return $this->storageConnection;
    }

    public function deleteDay(string $day)
    {
        $tomorrow = date('Y-m-d', strtotime('+1 day', strtotime($day)));

        $this->builder()->transaction(function (QueryBuilderHandler $db) use($day, $tomorrow) {
            $db->table('broadcasters')
                ->where('start_at', '>=', $day)
                ->where('start_at', '<', $tomorrow)
                ->delete();
        });
    }

    public function getEqualItemsFor(array $items) : array
    {
        $dbItems = [];

        foreach ($items as $day => $_) {
            $tomorrow = date('Y-m-d', strtotime('+1 day', strtotime($day)));
            $rows = $this->find()
                ->where('start_at', '>=', $day)
                ->where('start_at', '<', $tomorrow)
                ->get();

            $dbItems[$day] = json_decode(json_encode($rows), true);

        }

        return $dbItems;
    }
}
