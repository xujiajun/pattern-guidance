<h2>7、装饰者模式</h2>

今天不谈概念。直接来提出我们的问题。

<h5>7.1、问题</h5>

我们回到之前的游戏中Tile（区域）:

```php
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
```

我们先定义了一个Tile类.Tile表示部队单元所在的区域。每一个Tile对象都有明确的特征。

继续,我们希望给地表钻石的分布以及污染造成的破坏建模。我们可以这么做:

```php
class DiamondPlain extends Plains
{
    function getWealthFactor()
    {
        return parent::getWealthFactor() + 2;
    }
}

class PollutePlain extends Plains
{
    function getWealthFactory()
    {
        return parent::getWealthFactory() - 3;
    }
}
```

ok，我们现在要获取钻石地的收益:

```php
$tile = new DiamondPlain();
echo $tile->getWealthFactor();//输出4
```

似乎很完美了,现在问题来了,如果既要获取含有钻石又被污染的对象呢？

我们只能再加一个"既要获取含有钻石又被污染的类",这样，功能定义完全依赖继承体系，会导致类的数量很多。而且代码会产生重复。

<h5>7.2、实现</h5>

```php
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
        return $this->tile->getWealthFactor()+2;
    }
}

//client

$tile = new DiamondDecotor(new Plains());
echo $tile->getWealthFactor();//输出4
```
