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