<?php

abstract class Tile
{
    abstract function getWealthFactor();//获取收益
}


class Plains extends Tile
{
    private $wealthfactor = 2;

    function getWealthFactor()
    {
        return $this->wealthfactor;
    }
}

abstract class TileDecorator extends Tile
{
    protected $tile;

    function __construct(Tile $tile)
    {
        $this->tile = $tile;
    }
}

class DiamondDecotor extends TileDecorator
{
    function getWealthFactor()
    {
        return $this->tile->getWealthFactor() + 2;
    }
}

//client

$tile = new DiamondDecotor(new Plains());
echo $tile->getWealthFactor();//输出4