<?php

namespace Seed\Submission\Types;

use Seed\Core\Json\JsonSerializableType;
use Seed\Core\Json\JsonProperty;

class TraceResponseV2 extends JsonSerializableType
{
    /**
     * @var string $submissionId
     */
    #[JsonProperty('submissionId')]
    public string $submissionId;

    /**
     * @var int $lineNumber
     */
    #[JsonProperty('lineNumber')]
    public int $lineNumber;

    /**
     * @var TracedFile $file
     */
    #[JsonProperty('file')]
    public TracedFile $file;

    /**
     * @var mixed $returnValue
     */
    #[JsonProperty('returnValue')]
    public mixed $returnValue;

    /**
     * @var ?ExpressionLocation $expressionLocation
     */
    #[JsonProperty('expressionLocation')]
    public ?ExpressionLocation $expressionLocation;

    /**
     * @var StackInformation $stack
     */
    #[JsonProperty('stack')]
    public StackInformation $stack;

    /**
     * @var ?string $stdout
     */
    #[JsonProperty('stdout')]
    public ?string $stdout;

    /**
     * @param array{
     *   submissionId: string,
     *   lineNumber: int,
     *   file: TracedFile,
     *   returnValue?: mixed,
     *   expressionLocation?: ?ExpressionLocation,
     *   stack: StackInformation,
     *   stdout?: ?string,
     * } $values
     */
    public function __construct(
        array $values,
    ) {
        $this->submissionId = $values['submissionId'];
        $this->lineNumber = $values['lineNumber'];
        $this->file = $values['file'];
        $this->returnValue = $values['returnValue'] ?? null;
        $this->expressionLocation = $values['expressionLocation'] ?? null;
        $this->stack = $values['stack'];
        $this->stdout = $values['stdout'] ?? null;
    }
}
