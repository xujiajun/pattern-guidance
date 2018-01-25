<?php

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

class BloggsCommsManger extends CommsManager
{
    function getHeaderText()
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

class MegaCommsManger extends CommsManager
{
    function getHeaderText()
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

$comms = new MegaCommsManger();
$apptEncoder = $comms->getApptEncoder();
echo $apptEncoder->encode();//Appointment data encoded in MegaCal