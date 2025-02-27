<?php

namespace Seed\V2\V3\Problem;

use Seed\Core\Json\JsonSerializableType;
use Seed\Core\Json\JsonProperty;

class GetFunctionSignatureRequest extends JsonSerializableType
{
    /**
     * @var mixed $functionSignature
     */
    #[JsonProperty('functionSignature')]
    public mixed $functionSignature;

    /**
     * @param array{
     *   functionSignature: mixed,
     * } $values
     */
    public function __construct(
        array $values,
    ) {
        $this->functionSignature = $values['functionSignature'];
    }
}
