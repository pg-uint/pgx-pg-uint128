<?php

declare(strict_types=1);

require_once __DIR__ . '/Type.php';

class TypeConfig
{
    /**
     * @param array<Type> $scanTypes
     */
    public function __construct(
        public readonly Type $type,
        public readonly array $scanTypes = [],
    ) {
    }
}
