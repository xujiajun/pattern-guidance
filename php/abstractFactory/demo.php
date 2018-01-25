<?php

abstract class CommsManger
{
    abstract function getHeaderText();

    abstract function getApptEncoder();

    abstract function getTtdEncoder();//待办事宜

    abstract function getContactEncoder();//联系人

    abstract function getFooterText();
}

class MegaBloggsCommsManager extends CommsManger
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

class BloggsCommsManager extends CommsManger
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