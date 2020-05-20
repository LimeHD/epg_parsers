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
        return $this->builder()->table('broadcasts');
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

    public function deleteDay(string $day, int $id)
    {
        $tomorrow = date('Y-m-d', strtotime('+1 day', strtotime($day)));

        $this->builder()->transaction(function (QueryBuilderHandler $db) use($day, $tomorrow, $id) {
            $db->table('broadcasts')
                ->where('epg_id', '=', $id)
                ->where('start_at', '>=', $day)
                ->where('start_at', '<', $tomorrow)
                ->delete();
        });
    }

    public function getEqualItemsFor(array $items) : array
    {
        $dbItems = [];

        foreach ($items as $day => $broadcasts) {
            if (!isset($dbItems[$day])) {
                $dbItems[$day] = [];
            }

            foreach ($broadcasts as $id => $_) {
                if (!isset($dbItems[$day][$id])) {
                    $dbItems[$day][$id] = [];
                }

                $tomorrow = date('Y-m-d', strtotime('+1 day', strtotime($day)));
                $rows = $this->find()
                    ->where('epg_id', '=', $id)
                    ->where('start_at', '>=', $day)
                    ->where('start_at', '<', $tomorrow)
                    ->get();

                $dbItems[$day][$id] = json_decode(json_encode($rows), true);
            }
        }

        return $dbItems;
    }
}
