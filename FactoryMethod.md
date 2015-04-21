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