package cn.xujiajun.designPattern.creational.factory;

import cn.xujiajun.designPattern.creational.factory.model.Computer;
import cn.xujiajun.designPattern.creational.factory.model.PC;
import cn.xujiajun.designPattern.creational.factory.model.Server;

public class ComputerFactory {

    public static Computer getComputer(String type, String ram, String hdd, String cpu) {
        if ("PC".equalsIgnoreCase(type)) {
            return new PC(ram, hdd, cpu);
        } else if ("Server".equalsIgnoreCase(type)) {
            return new Server(ram, hdd, cpu);
        }

        return null;
    }
}