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
        return $this->builder()->table('epg');
    }

    /**
     * @param array $items
     * @return array
     */
    public function store(array $items) : array
    {
        $insertIds = [];

        foreach (Datamapper::batches($items) as $batch) {
            try {
                $insertIds[] = $this->getBroadcasterStorage()->insert($batch);
            } catch (Exception $e) {
                // todo to bugsnag
            }
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
        $dt = date('Y-m-d', strtotime($day));
        $this->builder()->transaction(function (QueryBuilderHandler $db) use($day, $dt, $id) {
            $db->table('epg')
                ->where('epg_id', '=', $id)
                ->where('date', '=', $dt)
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

                $dt = date('Y-m-d', strtotime($day));
                $rows = $this->find()
                    ->where('epg_id', '=', $id)
                    ->where('date', '=', $dt)
                    ->get();

                $dbItems[$day][$id] = json_decode(json_encode($rows), true);
            }
        }

        return $dbItems;
    }
}
