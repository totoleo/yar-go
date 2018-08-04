<?php
$t1 = microtime(true);
$client = new Yar_Client('http://zee.test:80');
$arguments =  ['A'=>4,'B'=>5];
$data = $client->__call("add",$arguments);
var_dump([$data,microtime(true)-$t1]);