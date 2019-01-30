<?php

require  __DIR__.'/../../vendor/autoload.php';

use Wsdl2PhpGenerator\Generator;

$generator = new Generator();
$generator->generate(new \Wsdl2PhpGenerator\Config([
    'inputFile' => 'emutation.wsdl',
    'outputDir' => 'dist/InteropFibre/Emutation/Generated',
    'namespaceName' => 'InteropFibre\Emutation\Generated'
]));

