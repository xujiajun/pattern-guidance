<h2 id="composite">6、组合模式</h2>

<h5 id="what-is-composite">6.1、概念</h5>

将一组对象组合为可像单个对象一样被使用的结构

<h5 id="composite-roles">角色</h5>

1、Component 是组合中的对象声明接口，在适当的情况下，实现所有类共有接口的默认行为。声明一个接口用于访问和管理Component子部件。

2、Leaf 在组合中表示叶子结点对象，叶子结点没有子结点。

3、Composite 定义有枝节点行为，用来存储子部件，在Component接口中实现与子部件有关操作，如增加(add)和删除(remove)等。

<h5 id="when-use-composite">6.2、适合场景</h5>
 
1、你想表示对象的部分-整体层次结构

2、你希望用户忽略组合对象与单个对象的不同，用户将统一地使用组合结构中的所有对象。
 
 <h5 id="issue-composite">6.3、问题</h5>
 
 我们回到上回虚构的场景，以“文明”的游戏为基础，设计一个系统。玩家可以在由大量区块组成的地图上移动战斗单元。独立的单元可以被组合来一起移动。ok，我先来定义战斗单元：
 
 ```php
 abstract class Unit
 {
    abstract function bombardStrength();//抽象一个攻击强度类
 }

class Archer extends Unit
{
    function bombardStrength()
    {
        return 4;
    }
}

class lasterCanonUnit extends Unit
{
    function bombardStrength()
    {
       return 44;
    }
}
```

定义一个独立组合战斗的单元

```php
class Army
{
  private $units = array();
  
  function addUnit( Unit $unit)
  {
    array_push($unit);
  }
  
  function bombardStrength()
  {
    $ret = 0;
    foreach($this->units as $unit)
    {
      $ret += $unit->bombardStrength();
    }
  }
  
}
```

Arm类有一个addUnit()方法，用户接受Unit对象。Unit对象被保存在属性数组$units中。同时我们在Army的bombardStrength()方法中通过一个简单的迭代遍历所聚合的Unit对象的bombardStrength()方法，计算总的攻击强度。

如果问题一直如此简单，那么这样的模型还是令人满意的。但是当我们加入新的需求的话，会发生什么呢？

比如现在军队应该可以与其他军队进行合并。同时，每个军队都会有它自己的ID，这样军队以后还能从整编中解散出来。比如今天有两支军队要并肩作战，但是当国内发生叛乱时，可能一支军队要调回。因此，我们不能仅仅将军队与军队合并，还要像添加Unit对象一样添加Army对象:

那我们修改bombardStrength()方法，遍历所有的军队(army)以及部队单位(unit)：

```
function bombardStrength()
{
  $ret = 0;
  
  foreach($this->units as $unit)
  {
    $ret += $unit->bombardStrength();
  }
  
  foreach($this->armies as $army)
  {
    $ret += $army->bombardStrength();
  }
  
  return $ret;
}
```

此时，这个新的需求（添加和抽取军队）还不算太复杂。但是记住，我们还需要对防御、移动等方法（上面省略）等做相似的修改。

现在需求方提出新要求：运兵船可以支持最多10个战斗单元以改进它们在某些地形上的活动范围。显然，运兵船与用于组合的Army对

象相似，但是它也有一些自己的特性。虽然我们可以通过改进Army类来处理运兵船对象。但是我们知道以后会有更多对部队进行组合

的需求。显然，我们需要一个更加灵活的模型
