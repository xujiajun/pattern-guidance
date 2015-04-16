系统学习设计模式（PHP代码演示）
====
本文由 [xujiajun][xujiajun] 整理、编辑并在 [CC BY-SA 3.0][CC] 协议下发布，主要为了给自己和各位朋友作为系统学习设计模式的参考资料。更多请关注 [superu.org][superu]
[CC]: http://zh.wikipedia.org/wiki/Wikipedia:CC "Wiki: CC"
[xujiajun]:http://xujiajun.cn
[superu]:http://superu.org
- - - 
前言
----
如何编辑本文？[请点击这里](https://github.com/xujiajun/Learning-Pattern/edit/master/README.md)

如何fork本文？点击右上角"Fork" 按钮.

如何发pull Requests 作出贡献? [请点击这里] (https://github.com/xujiajun/Learning-Pattern/pulls)

<h2>目录</h2>

- [1、模式简介](#pattern)
- &nbsp;&nbsp;[1.1、什么是模式](#what-is-pattern)
- &nbsp;&nbsp;[1.2、模式概览](#pattern-overview)
- &nbsp;&nbsp;&nbsp;&nbsp;[1.2.1、命名](#name)
- &nbsp;&nbsp;&nbsp;&nbsp;[1.2.2、问题](#issue)
- &nbsp;&nbsp;&nbsp;&nbsp;[1.2.3、解决方案](#solution)
- &nbsp;&nbsp;&nbsp;&nbsp;[1.2.4、效果](#result)
- [2、单例模式(Singleton)](#singleton)
- &nbsp;&nbsp;[2.1、问题](#issue-singleton)
- &nbsp;&nbsp;[2.2、解决](#solution-singleton)
- &nbsp;&nbsp;[2.3、效果](#result-singleton)

<h2 id="pattern">1、模式简介</h2>

<h5 id="what-is-pattern">1.1、什么是模式</h5>

看看前辈们如何总结的:

`在软件世界中，每个开发机构就像一个部落，而模式就是对部落的某种共同记忆的一种表现。---Grady Booch,《J2EE核心模式》`

`模式便是特定环境下同类问题的一种解决方法  ---Gand of Four，《设计模式：可复用面向对象软件基础》`

正如上面所暗示的，设计模式便是分析过的问题和问题解决方案所阐述的优秀实践。注意模式本质是自下而上的。它们来源于实践非空洞的理论。

<h5 id="pattern-overview">1.2、模式概览</h5>

一个模式的核心由4个部分组成：命名、问题、解决方案和效果。

<h5 id="name">1.2.1、命名</h5>

命名非常重要。命名可以说丰富了程序员的语言,少许简短的文字便可表示相当复杂的问题和解决方案。命名必须兼顾简洁性和描述性。《设计模式》说道：找到一个好名字,成了我们在开发模式目录中最困难的部分之一。

Martin Flower 也赞同模式命名至关重要,因为模式的目的之一就是为开发者更加有效的交流提供词汇。

<h5 id="issue">1.2.2、问题</h5>

无论解决方案如何优雅，问题以及问题发生的环境都是一个模式的基础。找出问题比使用模式目录中的解决方案更难。

模式会小心谨慎地描述问题的空间（发生的条件和环境）。问题会被置于环境中简明扼要地描述。问题会被分解成不同的细节和各种表现。

<h5 id="solution">1.2.3、解决方案</h5>

解决方案最初和问题放在一起的。尽管代码有现成的,但是解决方案从来不是简单的复制哦，有可能实现上又上百种的细微差。

把一个比方:就像农作物播种操作，如果你盲目的遵循书本上的步骤,那么有可能到收货季节你有可能挨饿，以模式为基础，但又针
对各种情况随机应变的方案更加实用。虽然问题（农作物成长）的基本解决方案大致相同（播种、灌溉、收割）,但是实际场景会有
各种因素,比如土壤类别、地理位置、土地方位、当地有害物等等。

<h5 id="result">1.2.4、效果</h5>

在设计代码的时候，你所做的每一个决定都会带来不同的结果。当然我们总是希望得到令人满意的针对问题的解决方案。解决方案一旦部署，理想情况下它也许非常适合与其他模式一同工作。但是也要留心是否会带来风险。

<h2 id="singleton">2、单例模式</h2>

<h5 id="issue-singleton">2.1、问题</h5>

假设现在有一个保存应用信息的类Preference,我们可以用它来保存如用户信息的字符串、文件路径等。这些信息在你每次调用时候都可能有所不同。类似一个“公告板”，它是可以被系统中其他无关对象设置和获取信息的中心。

但是问题来了，如果在对象中直接传递Preference,让原来并不使用Preference对象的类强制接受Preference对象，会产生耦合。

我们还需要保证系统中所有对象都使用的同一个Preference对象。

<h5 id="solution-singleton">2.2、解决</h5>

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
另外,PHP中不支持饿汉式的单例模式。因为PHP不支持在类定义时给类的成员变量赋予非基本类型的值。如表达式，new操作等
