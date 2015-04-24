<h2 id="composite">6、组合模式</h2>

<h5 id="what-is-composite">6.1、概念</h5>

将一组对象组合为可像单个对象一样被使用的结构

<h5 id="composite-roles">6.2、角色</h5>

1、Component 是组合中的对象声明接口，在适当的情况下，实现所有类共有接口的默认行为。声明一个接口用于访问和管理Component子部件。

2、Leaf 在组合中表示叶子结点对象，叶子结点没有子结点。

3、Composite 定义有枝节点行为，用来存储子部件，在Component接口中实现与子部件有关操作，如增加(add)和删除(remove)等。

<h5 id="when-use-composite">6.3、适合场景</h5>
 
1、你想表示对象的部分-整体层次结构

2、你希望用户忽略组合对象与单个对象的不同，用户将统一地使用组合结构中的所有对象。
 
 <h5 id="issue-composite">6.4、问题</h5>
 
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

<h5 id="solution-composite">6.5、解决</h5>

组合模式定义了一个单根继承体系，使具有截然不同职责的集合可以并肩工作。

我们重写Unit抽象类

```php
abstract class Unit
{
   abstract function addUnit(Unit $unit);
   abstract function removeUnit(Unit $unit);
   abstract function bombardStrength();
}
```

可以看到，我们为所有的Unit对象设计了基本功能。现在看看一个组合对象是如何实现这些抽象方法的:

```php

<?php 

class Army extends Unit
{
    private $units = array();

    function addUnit(Unit $unit)
    {
        if (in_array($unit,$this->units,true)) {
            return;
        }

        $this->units[] = $unit;
    }

    function removeUnit(Unit $unit)
    {
        $units = array();

        foreach ($this->units as $thisUnit) {
            if ($thisUnit != $unit) {
                $unit[] = $thisUnit;
            }
        }
    }

    function bombardStrength()
    {
        $ret  = 0;

        foreach ($this->units as $unit) {
            $ret += $unit->bombardStrength();
        }

        return $ret;
    }
}


```

组合模式的一个问题是如何实现add和remove方法。一般组合模式会在抽象类中添加add和remove方法。这可以确保模式中的所有类都

是共享同一个接口，但这同时也意味着局部必须实现这些方法:

```php

class unitException extends Exception {}

class Archer extends Unit
{

     function addUnit(Unit $unit)
     {
        throw new Exception(get_class($this)." is a leaf");
     }

     function removeUnit(Unit $unit)
     {
        throw new Exception(get_class($this)." is a leaf");
     }

     function bombardStrength()
     {
        return 4;
     }
}

```

如上所示,我们不希望用到addUnit、removeUnit,所以用到的时候,我们抛出异常来处理。

我们修改下Unit类的addUnit和removeUnit,要不然我们所有的局部类的add和remove都要改。

```php
abstract class Unit
{
    abstract function bombardStrength();

    function addUnit(Unit $unit)
    {
        throw new Exception(get_class($this)." is a leaf");
    }

    function removeUnit(Unit $unit)
    {
       throw new Exception(get_class($this)." is a leaf");
    }
}
```
这样做好处可以去除局部类中的重复的代码,但是组合类不再需要强制性实现addUnit()和removeUnit()方法了。

```php
//client
$mainArmy->addUnit(new Archer());
$mainArmy->addUnit(new Archer());

$subArmy = new Army();
$subArmy->addUnit(new Archer());
$mainArmy->addUnit($subArmy);
//输出总的攻击强度
echo $mainArmy->bombardStrength();//输出12
```

<h5 id="result-composite">6.6、效果总结</h5>

不知道看到现在,你有没有发现什么问题。Archer这个类明明不需要add和remove的方法，为什么还有这两个冗余方法。这就是组合模
式的原则便是局部类和组合类具有同样的接口。

那我们可不可以剥离Unit的add和remove方法:

```php
abstract class Unit 
{
    function getComposite()
    {
        return null;
    }

    abstract function bombardStrength();
}
```

再抽象出来一个类,甚至实现部分方法:

```php
abstract class CompositeUnit extends Unit
{
    private $units = array();

    function getComposite()
    {
        return $this;
    }

    protected function units()
    {
        return $this->units;
    }

   function removeUnit(Unit $unit)
   {
       $units = array();

       foreach ($this->units as $thisUnit) {
           if ($thisUnit != $unit) {
               $unit[] = $thisUnit;
           }
       }
   }

   function addUnit(Unit $unit)
   {
       if (in_array($unit,$this->units,true)) {
           return;
       }

       $this->units[] = $unit;
   }
}

```

注意：我们增加了`getComposite`这个方法，目的是区分是否是compositeUnit本身。


讲讲组合模式的缺陷,简化的前提是使所有的类都继承自同一个基类。简化的好处有时候会导致对象类型安全上的代价:假如有个Cavalry对象, 游戏规则规定不能将马匹放在运兵船上,这个时候没有自动化的方式强制执行规则,怎么办？

```php
function addUnit(Unit $unit)
{
  function addUnit(Unit $unit)
  {
  //手动判断下
  if( $unit instanceof Cavalry)
  {
    throw new UnitException('不要将马匹放在运兵船!');
  }
  super::addUnit($unit);
}
```
这样模型会变得复杂。特殊对象越多,用组合模式显得弊大于利。

另外,如果一个对象树中有大量的对象。比如上面的例子（Army类调用计算攻击强度方法）,可能一个简单的调用,可能导致系统奔溃。

总的来说,如果你想像对待单个对象一样对待组合对象，那么组合对象十分游泳。另一方面,组合模式又依赖于它组成部分的简单性。随着我们引入复杂的规则，代码会越来越难维护。
