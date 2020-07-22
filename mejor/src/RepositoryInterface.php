<?php

/**
 * Интерфейс для репозиториев
 */
interface RepositoryInterface
{
    public function save($rawData);
}