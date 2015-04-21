 <h2 id="prototype">5、原型模式</h2>
 
 <h5 id="what-is-prototype">5.1、概念</h5>
 
 原型模式是创建型模式的一种,其特点在于通过“复制”一个已经存在的实例来返回新的实例,而不是新建实例。被复制的实例就是我们所称的“原型”，这个原型是可定制的。
 
 <h5 id="when-use-prototype">5.2、适合场景</h5>
 
 原型模式多用于创建复杂的或者耗时的实例，因为这种情况下，复制一个已经存在的实例使程序运行更高效；或者创建值相等，只是命名不一样的同类数据。
 
 <h5 id="issue-prototype">5.3、问题</h5>
 
假设有一款‘文明’风格的游戏，可在区块组成的格子中操作战斗单元。区块分:海洋（sea）、平原（plain）、forest（森林）
我们允许用户允许用户在完全的不同的环境里选择

 <h5 id="solution-prototype">5.4、解决</h5>
 
 ```php
 <?php 

class sea{}

class EarthSea extends sea{}

class MarsSea extends sea{}

class Plains{}

class EarthPlains extends Plains{}

class MarsPlain extends Plains{}

class Forest{}

class EarthForest extends Forest{}

class MarsForest extends Forest{}

class TerrainFactory
{
    private $sea;
    private $forest;
    private $plains;

    function __construct(Sea $sea, Plains $plains,Forest $forest)
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
 ```