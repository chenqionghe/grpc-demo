<?php
/**
 * 客户端使用示例
 * @author chenqionghe
 */

require dirname(__FILE__) . '/vendor/autoload.php';

//初始化要去的健身房
$client = new \Lightweight\GymClient('127.0.0.1:50051', [
    'credentials' => Grpc\ChannelCredentials::createInsecure(),
]);

//带上自己的卡和运动计划
$request = new \Lightweight\Person();
$request->setName("chenqionghe");
$request->setActions(['深蹲', '卧推', '硬拉']);

//去健身房健身
list($response, $status) = $client->BodyBuilding($request)->wait();

echo sprintf("code: %s, msg: %s \n", $response->getCode(), $response->getMsg());
