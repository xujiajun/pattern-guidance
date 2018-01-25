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

    public function setProperty($key, $val)
    {
        $this->props[$key] = $val;
    }

    public function getProperty($key)
    {
        return $this->props[$key];
    }
}

$pref = Preference::getInstance();
$pref->setProperty("name", "xujiajun");
unset($pref);
$pref2 = Preference::getInstance();
echo $pref2->getProperty("name");//输出xujiajun  key为name的属性值并没有丢失。