<?php
// GENERATED CODE -- DO NOT EDIT!

namespace Lightweight;

/**
 * 健身房
 */
class GymClient extends \Grpc\BaseStub {

    /**
     * @param string $hostname hostname
     * @param array $opts channel options
     * @param \Grpc\Channel $channel (optional) re-use channel object
     */
    public function __construct($hostname, $opts, $channel = null) {
        parent::__construct($hostname, $opts, $channel);
    }

    /**
     * @param \Lightweight\Person $argument input argument
     * @param array $metadata metadata
     * @param array $options call options
     */
    public function BodyBuilding(\Lightweight\Person $argument,
      $metadata = [], $options = []) {
        return $this->_simpleRequest('/lightweight.Gym/BodyBuilding',
        $argument,
        ['\Lightweight\Reply', 'decode'],
        $metadata, $options);
    }

}
