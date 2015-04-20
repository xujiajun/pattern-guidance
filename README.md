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
- &nbsp;&nbsp;[2.1、概念](#what-is-singleton)
- &nbsp;&nbsp;[2.2、特点](#feature-singleton)
- &nbsp;&nbsp;[2.3、适合场景](#when-use-singleton)
- &nbsp;&nbsp;[2.4、问题](#issue-singleton)
- &nbsp;&nbsp;[2.5、解决](#solution-singleton)
- &nbsp;&nbsp;[2.6、效果总结](#result-singleton)
- [3、工厂方法模式(FactoryMethod)](#factoryMethod)
- &nbsp;&nbsp;[3.1、概念](#what-is-factoryMethod)
- &nbsp;&nbsp;[3.2、适合场景](#when-use-factoryMethod)
- &nbsp;&nbsp;[3.3、问题](#issue-factoryMethod)
- &nbsp;&nbsp;[3.4、解决](#solution-factoryMethod)
- &nbsp;&nbsp;[3.5、效果总结](#result-factoryMethod)
- [4、抽象工厂模式(AbstractFactory)](#abstractFactory)
- &nbsp;&nbsp;[4.1、概念](#what-is-abstractFactory)
- &nbsp;&nbsp;[4.2、适合场景](#when-use-abstractFactory)
- &nbsp;&nbsp;[4.3、问题](#issue-abstractFactory)
- &nbsp;&nbsp;[4.4、解决](#solution-abstractFactory)
- &nbsp;&nbsp;[4.5、效果总结](#result-abstractFactory)
- [5、原型模式(prototype)](#prototype)
- &nbsp;&nbsp;[5.1、概念](#what-is-prototype)
- &nbsp;&nbsp;[5.2、适合场景](#when-use-prototype)
- &nbsp;&nbsp;[5.3、问题](#issue-prototype)
- &nbsp;&nbsp;[5.4、解决](#solution-prototype)
- &nbsp;&nbsp;[5.5、效果总结](#result-prototype)

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

<h2 id="factoryMethod">工厂方法模式</h2>

<h5 id="what-is-factoryMethod">3.1、概念</h5>

它是一种实现了“工厂”概念的面向对象设计模式。定义一个创建对象的接口，但让实现这个接口的类来决定实例化哪个类。工厂方法让类的实例化推迟到子类中进行。

<h5 id="when-use-factoryMethod">3.2、适合场景</h5>

1、当一个类不知道它所必须创建的对象的类的时候

2、当一个类希望由它的子类来指定它所创建的对象的时候

3、当类将创建对象的职责委托给多个帮助子类的某一个，并且你希望将哪一个帮助子类是代理者这一信息局部化的时候

<h5 id="issue-factoryMethod">3.3、问题</h5>

假设现在有一个关于个人事务管理的项目，其功能之一是管理Appointment对象。需要一个BloggsCal的格式来和他们进行交流预约相

关数据。同时业务团队提醒将来会有更加多的数据格式。

在接口级别上，我们可以立即定义两个类。其一是需要一个数据编码器来把Appointment对象转换成专有的格式,这里我们命名成ApptEncoder这个类；另外需要一个管理类来获得编码器,并使用编码器与第三方进行通信，我们将这个管理类命名成CommsManager类。

用模式的角色来描述：这个CommsManager类为创建者（creator）,而ApptEncoder是产品（product)

那么，如何得到一个真正的具体的ApptEncoder对象？


```php

abstract class ApptEncoder
{
    abstract function encode();
}

class BlogApptEncoder extends ApptEncoder
{
    function encode()
    {
        return "Appointment data encoded in BloggsCal\n";
    }
}

class CommsManager
{
    function getApptEncoder()
    {
        return new BlogApptEncoder();
    }
}
```
当我们的业务团队说要加一个格式MegaCal时。我们调整代码:

```php
class CommsManager
{
    const BLOGGS = 1;
    const MEGA = 2;
    private $mode = 1;
    
    function __construct($mode)
    {
        $this->mode = $mode;
    }
    
    function getApptEncoder()
    {
        switch($this->mode)
        {
            case(self::MEGA):
                return new MegaApptEncoder();
            default:
                return new BlogApptEncoder();
        }
    }
}

class MegaApptEncoder extends ApptEncoder
{
    function encode()
    {
        return "Appointment data encoded in MegaCal\n";
    }
}

//client
$comms = new CommsManger(CommsManger::MEGA);
$apptEncoder = $comms->getApptEncoder();
echo $apptEncoder->encode();//Appointment data encoded in MegaCal
```
不要觉得这样已经很完美了。如果现在使用的协议要求提供页眉和页脚数据来描述每次预约,情况又如何。请看代码:

```php

class CommsManger
{
    const BLOGGS = 1;
    const MEGA = 2;
    private $mode = 1;
    
    function __construct($mode)
    {
        $this->mode = $mode;
    }
    
    function getHeaderText()
    {
        switch($this->mode)
        {
            case (self::MEGA):
                return "Mega ..header";
            default:
                return "Blogger header";
        }
    }
    
    function getApptEncoder()
    {
        switch($this->mode)
        {
            case(self::MEGA):
                return new MegaApptEncoder();
            default:
                return new BlogApptEncoder();
        }
    }
}
```

从上面代码发现没有，如果我还要加页脚的方法，oh my god！工作量上去了，而且bad smell。

总结下问题:

1、在代码运行时，我们才知道要生成的对象类型(BlogsApptEncoder或者MegaApptEncoder);

2、我们需要能够相对轻松地加入一些新的产品类型

3、每一个产品类型都要可定制的功能，如上面的页眉页脚。

<h5 id="solution-factoryMethod">3.4、解决</h5>

直接上代码:

```php
abstract class ApptEncoder
{
    abstract function encode();
}

class BloggsApptEncoder extends ApptEncoder
{
    function encode()
    {
        return "Appointment data encode in BLoggsCal format\n";
    }
}

class MegaApptEncoder extends ApptEncoder
{
    function encode()
    {
        return "Appointment data encode in MegaCal format\n";
    }
}

abstract class CommsManager
{
    abstract function getHeaderText();
    abstract function getApptEncoder();
    abstract function getFooterText();
}

class BloggsCommsManger extends CommManager
{
    fucntion getHeaderText()
    {
        return "Bloggs.. hearder\n";
    }

    function getApptEncoder()
    {
        return new BloggsApptEncoder();
    }
    
    function getFooterText()
    {
        return "Bloggs.. footer\n";
    }
}

class MegaCommsManger extends CommManager
{
    fucntion getHeaderText()
    {
        return "Mega.. hearder\n";
    }

    function getApptEncoder()
    {
        return new MegaApptEncoder();
    }
    
    function getFooterText()
    {
        return "Mega.. footer\n";
    }
}
```
<h5 id="result-factoryMethod">3.5、效果总结</h5>

我们从上面的实现来看,创建类与产品层次结构式非常相似的。这是使用工厂方法模式常见的结果。这形成一种特殊的代码重复吧。

另一方面，该模式可能会鼓励不必要的子类化。

<h2 id="abstractFactory">4、抽象工厂模式(AbstractFactory</h2>

<h5 id="what-is-abstractFactory">4.1、概念</h5>

提供一个创建一系列相关或相互依赖对象的接口，而无需指定他们具体的类。

<h5 id="when-use-abstractFactory">4.2、适合场景</h5>
比如支持多种观感标准的用户界面工具箱（Kit）。

<h5 id="issue-abstractFactory">4.3、问题</h5>

在上面的示例中我们只关注“预约”的功能。如果想扩展一下功能，使其能够处理“待办事宜”以及联系人,那么将会遇到新的问题。我们需要一个可以同时处理一组相关实现的架构。

<h5 id="solution-abstractFactory">4.4、实现</h5>

直接上代码:

 ```php
 
abstract class CommsManger
{
    abstract function getHeaderText();
    abstract function getApptEncoder();
    abstract function getTtdEncoder();//待办事宜
    abstract function getContactEncoder();//联系人
    abstract function getFooterText();
}

class MegaBloggsCommsManager extends  CommsManger
{
    function getHeaderText()
    {
        return "MegaCal header\n";
    }
    
    function getApptEncoder()
    {
        return new MegaApptEncoder();
    }
    
    function getTtdEncoder()
    {
        return new MegaTtdEncoder();
    }
    
    function getContactEncoder()
    {
        return new MegaContactEncoder();
    }
    
    function getFooterText()
    {
        return "MegaCal footer";
    }
}

class BloggsCommsManager extends  CommsManger
{
    function getHeaderText()
    {
        return "BloggsCal header\n";
    }
    
    function getApptEncoder()
    {
        return new BloggsApptEncoder();
    }
    
    function getTtdEncoder()
    {
        return new BloggsTtdEncoder();
    }
    
    function getContactEncoder()
    {
        return new BloggsContactEncoder();
    }
    
    function getFooterText()
    {
        return "BloggsCal footer";
    }
}

abstract class ApptEncoder
{
    abstract function encode();
}

class MegaApptEncoder extends ApptEncoder
{
    public function encode()
    {
        return "Appointment data encode in MegaCal format\n";
    }
}

class BloggsApptEncoder extends ApptEncoder
{
    public function encode()
    {
        return "Appointment data encode in BloggCal format\n";
    } 
}

abstract class TtdEncoder
{
    abstract function encode();
}

class MegaTtdEncoder extends TtdEncoder
{
    public function encode()
    {
        return "Todo data encode in MegaCal format\n";
    }
}

class BloggsTtdEncoder extends TtdEncoder
{
    public function encode()
    {
        return "Todo data encode in BloggCal format\n";
    } 
}

abstract class ContactEncoder
{
    abstract function encode();
}

class MegaContactEncoder extends ContactEncoder
{
    public function encode()
    {
        return "Contact data encode in MegaCal format\n";
    }
}

class BloggsContactEncoder extends ContactEncoder
{
    public function encode()
    {
        return "Contact data encode in BloggCal format\n";
    } 
}
 ```
 
 <h5 id="result-abstractFactory">4.5、结果</h5>

那么这个模式带来了什么？

1、首先将系统和实现的细节分离出来。我们可以在实例中添加或移除任意一个编码格式而不影响系统

2、对系统中功能相关的元素强制进行组合。因此，比如通过事宜BloggsCommsManager,可以确保只使用与BloggsCal相关的类。

3、缺点：添加新产品将会很麻烦。不仅要创建新产品的具体实现，为了支持他，我们必须修改抽象创建者和它的每一个具体的实现
 
 <h2 id="prototype">5、原型模式</h2>
 
 <h5 id="what-is-prototype">5.1、概念</h5>
 
 原型模式是创建型模式的一种,其特点在于通过“复制”一个已经存在的实例来返回新的实例,而不是新建实例。被复制的实例就是我们所称的“原型”，这个原型是可定制的。
 
 <h5 id="when-use-prototype">5.2、适合场景</h5>
 
 原型模式多用于创建复杂的或者耗时的实例，因为这种情况下，复制一个已经存在的实例使程序运行更高效；或者创建值相等，只是命名不一样的同类数据。
 
 <h5 id="issue-prototype">5.3、问题</h5>
 
假设有一款‘文明’风格的游戏，可在区块组成的格子中操作战斗单元。区块分:海洋（sea）、平原（plain）、forest（森林）
我们允许用户允许用户在完全的不同的环境里选择

 <h5 id="solution-prototype">5.4、解决</h5>
 
 
 
