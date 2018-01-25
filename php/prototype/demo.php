<?php

class sea
{
}

class EarthSea extends sea
{
}

class MarsSea extends sea
{
}

class Plains
{
}

class EarthPlains extends Plains
{
}

class MarsPlain extends Plains
{
}

class Forest
{
}

class EarthForest extends Forest
{
}

class MarsForest extends Forest
{
}

class TerrainFactory
{
    private $sea;
    private $forest;
    private $plains;

    function __construct(Sea $sea, Plains $plains, Forest $forest)
    {
        $this->sea = $sea;
        $this->plains = $plains;
        $this->forest = $forest;
    }

    function getSea()
    {
        return clone $this->sea;
    }

    function getPlain()
    {
        return clone $this->plains;
    }

    function getForeast()
    {
        return clone $this->forest;
    }
}

$factory = new TerrainFactory(new EarthSea(), new EarthPlains(), new  EarthForest);

var_dump($factory->getSea());
var_dump($factory->getPlain());
var_dump($factory->getForeast());

//输出
//class EarthSea#5 (0) {
//}
//class EarthPlains#5 (0) {
//}
//class EarthForest#5 (0) {
//}