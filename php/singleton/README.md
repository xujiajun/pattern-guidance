<h2 id="singleton">2、单例模式</h2>

<h5 id="what-is-singleton">2.1、概念</h5>
保证一个类仅有一个实例，并且提供一个访问它的全局访问点。单例又分懒汉、饿汉（PHP不支持）、登记式。

<h5 id="feature-singleton">2.2、特点</h5>

1、一个类只有一个实例

2、它必须自行创建这个实例

3、必须自行向整个系统提供这个实例

<h5 id="when-use-singleton">2.3、适合场景</h5>

1、当类只能有一个实例而且客户可以从一个众所周知的访问点访问它时

2、当这个唯一实例应该是通过子类化可扩展的。并且用户应该无需更改代码就能使用一个扩展的实例时。

<h5 id="issue-singleton">2.4、问题</h5>

假设现在有一个保存应用信息的类Preference,我们可以用它来保存如用户信息的字符串、文件路径等。这些信息在你每次调用时候都可能有所不同。类似一个“公告板”，它是可以被系统中其他无关对象设置和获取信息的中心。

但是问题来了，如果在对象中直接传递Preference,让原来并不使用Preference对象的类强制接受Preference对象，会产生耦合。

我们还需要保证系统中所有对象都使用的同一个Preference对象。

<h5 id="solution-singleton">2.5、解决</h5>

我们可以从强制控制对象的实例化开始。

```php

<?php 

class Preference
{
    private $props = array();

    private static $instance;
    
    //防止外界实例化对象
    private function __construct()
    {

    }
    //防止外界clone实例
    private function __clone()
    {

    }
    
    public static function getInstance()
    {
        if (empty(self::$instance)) {
            self::$instance = new Preference();
        }

        return self::$instance;
    }

    public function setProperty($key,$val)
    {
        $this->props[$key] = $val;
    }

    public function getProperty($key)
    {
        return $this->props[$key];
    }
}

$pref = Preference::getInstance();
$pref->setProperty("name","xujiajun");
unset($pref);
$pref2 = Preference::getInstance();
echo $pref2->getProperty("name");//输出xujiajun  key为name的属性值并没有丢失。

```
<h5 id="result-singleton">2.6、效果总结</h5>

因为单例在系统任何地方都可以被访问，所以它们可能导致很难调试的依赖关系。如果改变一个单例，那么所有使用该单例的类可能都会受到影响。虽然我们在每次声明一个特定类型的参数的方法时，也都创建了依赖关系。问题是当单例被使用时，依赖便会被隐藏在方法内部，并没有出现在方法的声明中,这使得系统的依赖关系更加难以追踪。

当然，适当的使用单例可以改进系统的设计。在OOP编程中，单例模式是一种全局变量的改进。你无法用错误类型的数据覆盖一个单例。这种保护尤其不支持命名空间的PHP版本中尤其重要。