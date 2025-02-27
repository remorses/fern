<?php

namespace Seed;

use Seed\Core\Json\JsonSerializableType;
use Seed\Traits\BaseResource;
use Seed\Core\Json\JsonProperty;

class Account extends JsonSerializableType
{
    use BaseResource;

    /**
     * @var string $resourceType
     */
    #[JsonProperty('resource_type')]
    public string $resourceType;

    /**
     * @var string $name
     */
    #[JsonProperty('name')]
    public string $name;

    /**
     * @var ?Patient $patient
     */
    #[JsonProperty('patient')]
    public ?Patient $patient;

    /**
     * @var ?Practitioner $practitioner
     */
    #[JsonProperty('practitioner')]
    public ?Practitioner $practitioner;

    /**
     * @param array{
     *   resourceType: string,
     *   name: string,
     *   patient?: ?Patient,
     *   practitioner?: ?Practitioner,
     *   id: string,
     *   relatedResources: array<Account|Patient|Practitioner|Script>,
     *   memo: Memo,
     * } $values
     */
    public function __construct(
        array $values,
    ) {
        $this->resourceType = $values['resourceType'];
        $this->name = $values['name'];
        $this->patient = $values['patient'] ?? null;
        $this->practitioner = $values['practitioner'] ?? null;
        $this->id = $values['id'];
        $this->relatedResources = $values['relatedResources'];
        $this->memo = $values['memo'];
    }
}
